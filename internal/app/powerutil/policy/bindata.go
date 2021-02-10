// Code generated for package policy by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../configs/vault/policies/powerssl-apiserver.hcl
// ../../../../configs/vault/policies/powerssl-controller.hcl
// ../../../../configs/vault/policies/powerssl-signer.hcl
// ../../../../configs/vault/policies/powerssl-worker.hcl
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

var _powersslApiserverHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0e\x40\x30\x10\x05\xd0\x7d\x4f\x31\x99\xbd\xf4\x04\x4e\x22\x16\x83\x09\x3f\x1a\x7e\x3a\xc5\x42\xdc\xdd\xa3\xb5\x4d\x94\x3b\x72\x60\x3d\x32\xcf\xc7\x6b\x44\xe9\x8c\x08\xaf\xb7\x57\x95\x37\x89\xcc\x46\x9b\x50\xd0\xe0\x21\xbd\x0c\x7a\x71\xb1\xe6\x3a\xa6\x2f\xfd\x01\x00\x00\xff\xff\xb7\xe0\x38\xa8\x43\x00\x00\x00")

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

var _powersslWorkerHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x2c\xc9\x50\x50\x2a\x29\x4a\xcc\x2b\xce\x2c\xd1\xcf\x4e\xad\x2c\xd6\xd7\x52\x52\xa8\xe6\x52\x50\x48\x4e\x2c\x48\x4c\xca\xcc\xc9\x2c\xc9\x4c\x2d\x56\xb0\x55\x88\x56\x2a\x2d\x48\x49\x2c\x49\x55\x8a\xe5\xaa\xe5\x02\x04\x00\x00\xff\xff\x4e\x74\x54\x86\x36\x00\x00\x00")

func powersslWorkerHclBytes() ([]byte, error) {
	return bindataRead(
		_powersslWorkerHcl,
		"powerssl-worker.hcl",
	)
}

func powersslWorkerHcl() (*asset, error) {
	bytes, err := powersslWorkerHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "powerssl-worker.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"powerssl-worker.hcl":     powersslWorkerHcl,
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
	"powerssl-worker.hcl":     &bintree{powersslWorkerHcl, map[string]*bintree{}},
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
