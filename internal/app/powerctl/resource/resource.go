package resource

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/jinzhu/inflection"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	goyaml "gopkg.in/yaml.v2"

	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

type Handler interface {
	Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error)
	Delete(client *apiserverclient.GRPCClient, name string) error
	Get(client *apiserverclient.GRPCClient, name string) (*Resource, error)
	List(client *apiserverclient.GRPCClient) ([]*Resource, error)
	Spec() interface{}
	Columns(resource *Resource) ([]string, []string)
	Describe(client *apiserverclient.GRPCClient, resource *Resource, output io.Writer) error
}

type resourcesStruct struct {
	m map[string]Handler
	s []string
}

func (r *resourcesStruct) Add(resourceHandler Handler, aliases ...string) {
	var kind string
	if t := reflect.TypeOf(resourceHandler); t.Kind() == reflect.Ptr {
		kind = t.Elem().Name()
	} else {
		kind = t.Name()
	}
	kind = strings.ToLower(kind)
	r.m[kind] = resourceHandler
	plural := inflection.Plural(kind)
	r.m[plural] = resourceHandler
	for _, alias := range aliases {
		r.m[alias] = resourceHandler
	}
	r.s = append(r.s, kind)
}

func (r *resourcesStruct) Kinds() []string {
	return r.s
}

func (r *resourcesStruct) Get(kind string) (Handler, error) {
	resourceHandler, ok := r.m[kind]
	if !ok {
		return nil, errors.New("unknown resource kind")
	}
	return resourceHandler, nil
}

var resources = resourcesStruct{m: make(map[string]Handler)}

type Resource struct {
	Kind string        `json:"kind,omitempty" yaml:"kind,omitempty"`
	Meta *resourceMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
	Spec interface{}   `json:"spec,omitempty" yaml:"spec,omitempty"`
}

func ResourcesFromFile(filename string) ([]*Resource, error) {
	if filename == "" {
		return nil, errors.New("please specify filename")
	}

	var in []byte
	if filename == "-" {
		var buf bytes.Buffer
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					buf.WriteString(line)
					break
				} else {
					return nil, err
				}
			}
			buf.WriteString(line)
		}
		in = buf.Bytes()
		if !json.Valid(in) {
			var err error
			in, err = yaml.YAMLToJSON(in)
			if err != nil {
				return nil, err
			}
		}
	} else {
		var err error
		in, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		dec := goyaml.NewDecoder(strings.NewReader(string(in)))
		for {
			var value interface{}
			err := dec.Decode(&value)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
		}

		switch filepath.Ext(filename) {
		case ".json":
		case ".yml", ".yaml":
			var err error
			in, err = yaml.YAMLToJSON(in)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("unknown input format")
		}
	}

	var resources []resourceLoader
	if json.Unmarshal(in, &resources) != nil {
		var resource resourceLoader
		if err := json.Unmarshal(in, &resource); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	out := make([]*Resource, len(resources))
	for i, r := range resources {
		resourceHandler, err := ResourceHandlerByKind(r.Kind)
		if err != nil {
			return nil, err
		}
		spec := resourceHandler.Spec()
		if r.Spec != nil {
			if err := json.Unmarshal(r.Spec, &spec); err != nil {
				return nil, fmt.Errorf("spec error: %s", err)
			}
		}
		out[i] = &Resource{
			Kind: r.Kind,
			Meta: r.Meta,
			Spec: spec,
		}
	}

	return out, nil
}

func (r *Resource) Create(client *apiserverclient.GRPCClient) (*Resource, error) {
	resourceHandler, err := resources.Get(r.Kind)
	if err != nil {
		return nil, err
	}
	return resourceHandler.Create(client, r)
}

func (r *Resource) Delete(client *apiserverclient.GRPCClient) error {
	resourceHandler, err := resources.Get(r.Kind)
	if err != nil {
		return err
	}
	return resourceHandler.Delete(client, r.Meta.UID)
}

func (r *Resource) Get(client *apiserverclient.GRPCClient) (*Resource, error) {
	resourceHandler, err := resources.Get(r.Kind)
	if err != nil {
		return nil, err
	}
	return resourceHandler.Get(client, r.Meta.UID)
}

func (r *Resource) ToTable() ([]string, []string, error) {
	resourceHandler, err := resources.Get(r.Kind)
	if err != nil {
		return nil, nil, err
	}
	rHeader, rColums := resourceHandler.Columns(r)
	var header []string
	header = append(header, "UID")
	header = append(header, rHeader...)
	var columns []string
	columns = append(columns, r.Meta.UID)
	columns = append(columns, rColums...)
	return header, columns, nil
}

func (r *Resource) Describe(client *apiserverclient.GRPCClient, output io.Writer) error {
	resourceHandler, err := resources.Get(r.Kind)
	if err != nil {
		return err
	}
	return resourceHandler.Describe(client, r, output)
}

type resourceLoader struct {
	Resource
	Spec json.RawMessage `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type resourceMeta struct {
	UID        string    `json:"uid,omitempty"        yaml:"uid,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty" yaml:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}

func ResourceFromArgs(args []string) (*Resource, error) {
	if len(args) == 1 && strings.Contains(args[0], "/") {
		s := strings.Split(args[0], "/")
		return &Resource{
			Kind: s[0],
			Meta: &resourceMeta{
				UID: s[1],
			},
		}, nil
	} else if len(args) == 2 && !strings.Contains(args[0], "/") && !strings.Contains(args[1], "/") {
		return &Resource{
			Kind: args[0],
			Meta: &resourceMeta{
				UID: args[1],
			},
		}, nil
	} else {
		return nil, errors.New("malformed")
	}
}

func ResourcesFromArgs(args []string) (resources []*Resource, err error) {
	if strings.Contains(args[0], "/") {
		for _, res := range args {
			s := strings.Split(res, "/")
			resource := &Resource{
				Kind: s[0],
				Meta: &resourceMeta{
					UID: s[1],
				},
			}
			resources = append(resources, resource)
		}
	} else {
		kind := args[0]
		if len(args) == 1 {
			return nil, errors.New("uid needed")
		}
		for _, uid := range args[1:] {
			if strings.Contains(uid, "/") {
				return nil, errors.New("malformed")
			}
			resource := &Resource{
				Kind: kind,
				Meta: &resourceMeta{
					UID: uid,
				},
			}
			resources = append(resources, resource)
		}
	}
	return resources, nil
}

func ResourceHandlerByKind(kind string) (Handler, error) {
	return resources.Get(kind)
}

func Kinds() []string {
	return resources.Kinds()
}

func listResource(listFunc func(pageToken string) ([]*Resource, string, error)) ([]*Resource, error) {
	var pageToken string
	var resources []*Resource
	for {
		res, nextPageToken, err := listFunc(pageToken)
		if err != nil {
			return nil, err
		}
		resources = append(resources, res...)
		if nextPageToken == "" {
			break
		}
		pageToken = nextPageToken
	}
	return resources, nil
}

func FormatResource(res interface{}, w io.Writer) (err error) {
	var out []byte
	switch viper.GetString("output") {
	case "json":
		if out, err = json.Marshal(res); err != nil {
			return err
		}
	case "table":
		var ok bool
		var resources []*Resource
		resources, ok = res.([]*Resource)
		if !ok {
			if r, ok := res.(*Resource); ok {
				resources = append(resources, r)
			}
		}
		resourceMap := make(map[string][]*Resource)
		for _, res := range resources {
			resourceMap[res.Kind] = append(resourceMap[res.Kind], res)
		}
		first := true
		for kind, resources := range resourceMap {
			if len(resourceMap) > 1 {
				if first {
					first = false
				} else {
					out = append(out, []byte(fmt.Sprintln(""))...)
				}
				out = append(out, []byte(fmt.Sprintln(kind))...)
			}
			buf := new(bytes.Buffer)
			table := tablewriter.NewWriter(buf)
			table.SetBorder(false)
			table.SetColumnSeparator(" ")
			table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			table.SetHeaderLine(false)
			for _, res := range resources {
				header, columns, err := res.ToTable()
				if err != nil {
					return err
				}
				table.SetHeader(header)
				table.Append(columns)
			}
			table.Render()
			scanner := bufio.NewScanner(buf)
			for scanner.Scan() {
				out = append(out, []byte(fmt.Sprintln(strings.TrimSpace(scanner.Text())))...)
			}
		}
	case "yaml":
		if out, err = yaml.Marshal(res); err != nil {
			return err
		}
	}
	_, _ = fmt.Fprintln(w, string(out))
	return nil
}
