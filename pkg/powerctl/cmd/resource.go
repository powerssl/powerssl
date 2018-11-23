package cmd

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/inflection"

	apiserverclient "powerssl.io/pkg/apiserver/client"
)

type ResourceHandler interface {
	Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error)
	Delete(client *apiserverclient.GRPCClient, name string) error
	Get(client *apiserverclient.GRPCClient, name string) (*Resource, error)
	List(client *apiserverclient.GRPCClient) ([]*Resource, error)
	Spec() interface{}
	Columns(resource *Resource) ([]string, []string)
}

type resources struct {
	m map[string]ResourceHandler
	s []string
}

func (r *resources) Add(resourceHandler ResourceHandler, aliases ...string) {
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

func (r *resources) Kinds() []string {
	return r.s
}

func (r *resources) Get(kind string) (ResourceHandler, error) {
	resourceHandler, ok := r.m[kind]
	if !ok {
		return nil, errors.New("unknown resource kind")
	}
	return resourceHandler, nil
}

var Resources = resources{m: make(map[string]ResourceHandler)}

type Resource struct {
	Kind string        `json:"kind,omitempty" yaml:"kind,omitempty"`
	Meta *ResourceMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
	Spec interface{}   `json:"spec,omitempty" yaml:"spec,omitempty"`
}

func (r *Resource) Create(client *apiserverclient.GRPCClient) (*Resource, error) {
	resourceHandler, err := Resources.Get(r.Kind)
	if err != nil {
		return nil, err
	}
	return resourceHandler.Create(client, r)
}

func (r *Resource) Delete(client *apiserverclient.GRPCClient) error {
	resourceHandler, err := Resources.Get(r.Kind)
	if err != nil {
		return err
	}
	return resourceHandler.Delete(client, r.Meta.UID)
}

func (r *Resource) Get(client *apiserverclient.GRPCClient) (*Resource, error) {
	resourceHandler, err := Resources.Get(r.Kind)
	if err != nil {
		return nil, err
	}
	return resourceHandler.Get(client, r.Meta.UID)
}

func (r *Resource) ToTable() ([]string, []string, error) {
	resourceHandler, err := Resources.Get(r.Kind)
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

type ResourceLoader struct {
	Resource
	Spec json.RawMessage `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type ResourceMeta struct {
	UID        string    `json:"uid,omitempty"        yaml:"uid,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty" yaml:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}

func resourcesFromArgs(args []string) (resources []*Resource, err error) {
	if strings.Contains(args[0], "/") {
		for _, res := range args {
			s := strings.Split(res, "/")
			resource := &Resource{
				Kind: s[0],
				Meta: &ResourceMeta{
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
				Meta: &ResourceMeta{
					UID: uid,
				},
			}
			resources = append(resources, resource)
		}
	}
	return resources, nil
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
