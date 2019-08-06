// Code generated for package policy by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../configs/vault/policies/powerssl-apiserver.hcl
// ../../../../configs/vault/policies/powerssl-controller.hcl
// ../../../../configs/vault/policies/powerssl-signer.hcl
package policy

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _powersslApiserverHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcc\xc1\x09\xc2\x50\x0c\x00\xd0\x7b\xa6\x08\x39\x0a\x92\x09\x9c\x44\x3c\x44\x0d\x1a\xfe\xe7\x1b\x92\xd8\x52\x4a\x77\xef\xa1\x03\x74\x80\xf7\x5c\xea\x8b\xe4\xcd\x38\xed\x33\xd8\x7f\xb3\x46\x66\xbf\x8a\x5b\x6a\x4c\x1a\x84\x2b\x20\xbe\xc4\xe5\x69\xdd\xca\x34\xf1\x86\x77\xfa\xfb\x5b\x4a\xe9\x01\x1b\xc0\x71\x54\xc8\x48\x2b\x6e\xba\x24\x5f\xce\xd9\x1e\x00\x00\xff\xff\x63\x24\xdc\xa5\x7a\x00\x00\x00")

func powersslApiserverHclBytes() ([]byte, error) {
	return bindataRead(
		_powersslApiserverHcl,
		"powerssl-apiserver.hcl",
	)
}

func powersslApiserverHcl() (*asset, error) {
	bytes, err := powersslApiserverHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "powerssl-apiserver.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _powersslControllerHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x51\x0a\x80\x30\x08\x00\xd0\x7f\x4f\x21\xfe\xc7\x4e\xd0\x49\xa2\x0f\x5b\x52\x92\x6c\x32\x8d\x3e\xa2\xbb\xf7\x9c\xf3\x44\xf2\x4b\x4b\xe8\xd1\x8a\xf7\x47\x46\x84\x4d\xb5\xb7\x1c\xdd\x4c\x06\xe1\x0b\x88\x95\x9d\x37\x35\x4d\x95\xc0\x19\x17\xba\x7d\xe7\x14\x5a\xe1\x03\xf8\x03\x00\x00\xff\xff\xbe\x70\x2e\xd4\x45\x00\x00\x00")

func powersslControllerHclBytes() ([]byte, error) {
	return bindataRead(
		_powersslControllerHcl,
		"powerssl-controller.hcl",
	)
}

func powersslControllerHcl() (*asset, error) {
	bytes, err := powersslControllerHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "powerssl-controller.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _powersslSignerHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x2c\xc9\x50\x50\x2a\xc8\xce\xd4\x2f\xce\x4c\xcf\xd3\x2f\xc8\x2f\x4f\x2d\x2a\x2e\xce\xd1\x05\xf1\x52\x8b\x94\x14\xaa\xb9\x14\x14\x92\x13\x0b\x12\x93\x32\x73\x32\x4b\x32\x53\x8b\x15\x6c\x15\xa2\x95\x4a\x0b\x52\x12\x4b\x52\x95\x62\xb9\x6a\xb9\xb8\x00\x01\x00\x00\xff\xff\xb4\xa4\x0b\xf2\x41\x00\x00\x00")

func powersslSignerHclBytes() ([]byte, error) {
	return bindataRead(
		_powersslSignerHcl,
		"powerssl-signer.hcl",
	)
}

func powersslSignerHcl() (*asset, error) {
	bytes, err := powersslSignerHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "powerssl-signer.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"powerssl-apiserver.hcl":  powersslApiserverHcl,
	"powerssl-controller.hcl": powersslControllerHcl,
	"powerssl-signer.hcl":     powersslSignerHcl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"powerssl-apiserver.hcl":  &bintree{powersslApiserverHcl, map[string]*bintree{}},
	"powerssl-controller.hcl": &bintree{powersslControllerHcl, map[string]*bintree{}},
	"powerssl-signer.hcl":     &bintree{powersslSignerHcl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
