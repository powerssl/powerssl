// Code generated by go-bindata.
// sources:
// ../../../web/app/index.html
// DO NOT EDIT!

package web

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6f\x53\xdb\x3c\x12\x7f\xdf\x4f\xb1\xd5\x3d\x37\x24\x03\xb6\x13\x02\x14\x82\x9d\x99\x1c\x57\x7a\x50\xda\x12\x28\x85\xa7\xef\x14\x6b\x1d\x2b\x91\x25\x55\x92\xf3\xa7\x1d\xbe\xfb\x8d\xe2\x84\x38\x2d\xed\xf5\x99\x9b\xe7\x05\x58\x5a\xad\xf6\xff\xfe\x56\x89\x5f\x32\x95\xba\x85\x46\xc8\x5d\x21\x7a\x2f\x62\xff\x01\x41\xe5\x28\x21\x28\x49\xef\x05\x40\x9c\x23\x65\x7e\x01\x10\x17\xe8\x28\xa4\x39\x35\x16\x5d\x42\x4a\x97\x05\xc7\xa4\x7e\x24\x69\x81\x09\x99\x72\x9c\x69\x65\x1c\x81\x54\x49\x87\xd2\x25\x64\xc6\x99\xcb\x13\x86\x53\x9e\x62\xb0\xdc\xec\x01\x97\xdc\x71\x2a\x02\x9b\x52\x81\x49\x7b\x0f\x6c\x6e\xb8\x9c\x04\x4e\x05\x19\x77\x89\x54\x6b\xd1\x82\xcb\x09\x18\x14\x09\xb1\x6e\x21\xd0\xe6\x88\x8e\x40\x6e\x30\x4b\x48\xee\x9c\xb6\xdd\x28\x2a\xe8\x3c\x65\x32\x1c\x2a\xe5\xac\x33\x54\xfb\x4d\xaa\x8a\xe8\x89\x10\x1d\x84\xad\xb0\x15\xa5\xd6\x6e\x68\x61\xc1\x65\x98\x5a\x4b\x80\x4b\x87\x23\xc3\xdd\x22\x21\x36\xa7\x9d\xe3\x83\xe0\x8d\x3c\xec\x1c\x1f\xcc\xbf\x0c\xda\x54\xdd\x3f\xf4\x77\x5b\x87\xc7\x37\x0f\xd7\xf3\xeb\xd1\x51\xb6\x38\xb8\xb8\x9f\x7e\x7c\x9f\xb7\x5e\xef\x1f\x75\x1e\x8a\xf3\xf4\x52\xdc\xf6\x67\xfc\xcd\xe8\xbc\x7f\x1f\xb1\x3e\xbf\x3d\xba\x7c\x28\x08\xa4\x46\x59\xab\x0c\x1f\x71\x99\x10\x2a\x95\x5c\x14\xaa\xb4\x6b\xa7\x1c\x77\x02\x7b\xd7\x6a\x86\xe6\xf6\xf6\x0a\xfa\xa5\xcb\xe3\xa8\x22\x56\x0c\x4b\x5f\xab\x35\xc0\x50\xb1\xc5\xde\x32\x47\xf0\x6d\x45\x02\xc8\x91\x8f\x72\xd7\x85\x76\xab\xf5\xcf\xd3\x27\x6a\x41\xcd\x88\xcb\x2e\xb4\x36\x24\x35\x45\x93\x09\x35\xeb\x42\xce\x19\x43\xb9\x39\xd1\x94\x31\x2e\x47\x35\xee\xc7\xd5\x57\x1b\xac\xa9\x02\x98\xe5\xdc\x61\x60\x35\x4d\xb1\xeb\x0f\x83\x99\xa1\xfa\xf4\x67\x0c\x41\xa1\xbe\x06\xff\x9b\x4b\xfd\x06\xcf\xb3\x1c\xca\xb0\x25\xb1\x0b\x43\x83\x74\x12\x78\xc2\xf7\x1e\xfc\x43\xa8\x11\x97\x35\x27\x86\xca\x39\x55\x6c\x45\x46\x60\xe6\xb6\x08\x5a\x59\xee\xb8\x92\x5d\xa0\x43\xab\x44\xe9\x70\x73\x66\xaa\x70\xd7\xb8\x9d\xd2\x5d\x68\xe9\xf9\x29\x6c\xe9\x8e\xa3\xa7\xe4\xc5\xd1\xba\x7d\x62\x9f\xc3\x75\xb3\x50\x2e\xc1\x28\x81\x09\xf1\x4b\x02\xa9\xa0\xd6\x26\xc4\xf7\x0b\xe5\x12\x0d\x81\xa5\x84\x84\x30\x6e\xb5\xa0\x8b\x2e\x48\x25\x91\xac\xcb\x21\xce\xdb\xeb\x2b\x85\x0b\x0e\x49\xad\x8e\xb4\x8e\xa3\xbc\xbd\xd2\x13\x79\xe9\xab\x35\xe3\xd3\xdf\x54\x73\xba\xaa\xa1\x60\xe9\x5e\x1b\x8b\x8d\x5e\xbd\x16\x21\x90\x32\xd2\x8b\xad\xa6\x72\x4d\xca\x94\x74\xc1\x6c\x59\x93\xc1\x50\x09\x46\x7a\x7d\x70\x6a\x82\x12\xb8\x05\x0a\x99\x32\x05\xa8\x0c\x28\x68\x6a\xad\x4f\x58\x18\x47\xfe\x7e\x0f\xfe\xad\x40\x2a\x07\x36\xa7\x06\x61\xa1\x4a\x03\xfd\xeb\x8b\xea\x6e\x18\x47\x7a\xa3\xbd\x17\xfb\xb2\x5c\xe9\xf3\x85\x61\x53\xa3\x84\xa0\x43\x81\xa4\x17\xa7\x8a\x61\x2f\x8e\x56\x1f\x6d\x96\xff\xeb\x97\x87\xa5\x73\x4a\x6e\xa2\xa0\x17\x30\x74\xd2\xff\x05\x16\x53\x25\x19\x35\x0b\x02\x8c\x3a\x1a\xa4\x82\xeb\xa1\xa2\x86\x05\x8e\x9a\x91\x47\x3b\x2f\x96\xf4\xce\x94\x5e\xc4\x51\x25\x68\x23\x3e\x8e\x18\x9f\xd6\xe2\xcc\x59\x42\x96\xd5\xb7\x89\x1c\xcf\x0c\x2d\x10\xac\x49\x13\xf2\xed\x1b\x84\xbe\xe1\xef\x6e\x2e\xe0\xf1\x91\x40\x85\x8e\xc4\xf7\x31\x59\x75\xf5\x7a\xb7\xbc\x35\x54\x86\xa1\x49\x48\x8b\x54\x92\xb9\x75\x28\xd1\x9c\xfb\x33\xd2\x8b\xa3\x4a\xf6\x8f\xa6\xd8\xd4\x70\xed\x2a\x9d\x6b\xa0\xf4\x6e\x84\xe3\x2f\x25\x9a\xc5\x12\x20\xab\x65\xd0\x09\x3b\x61\x7b\x89\x87\xe3\x1f\xe0\x70\xff\xf0\x28\x38\x1f\xe9\xb3\x61\xf4\xf6\x72\x20\xae\xde\x67\x1f\xca\x93\xb6\xa3\x9d\x7d\x15\xbd\x7f\xf7\x79\x2e\xdc\xec\x46\x1d\x0f\x5c\x31\x79\x77\xc3\xfa\xe5\x71\xf2\x73\xe8\x8b\xa3\xca\xa4\x5f\xd9\xc7\xe4\xd8\x86\xa9\x50\x25\xcb\x04\x35\xb8\x34\x92\x8e\xe9\x3c\x12\x7c\x68\x23\xad\xb4\x46\x13\x8e\x6d\xd4\x0e\xdb\xfb\xe1\x49\x54\x16\x6c\x4d\x7c\xde\x7a\x0f\xe6\x7d\xfd\x7e\x38\xca\x4f\xfe\xb5\xfb\x67\x7b\xf0\xd6\x4d\x3b\x37\xf2\xd5\x7d\xa7\x18\x5d\xcf\xf3\xbb\x93\xb7\xd1\x6d\x3a\xb0\xfd\xeb\x57\xf9\x1d\x1f\x3e\x74\x4e\xc6\xaf\x32\x3a\x39\xbf\xb6\x93\xe9\x43\x69\xa7\x19\x6d\x0d\x0f\x06\xff\x97\x47\xbf\x3b\x9a\xc6\xdf\x4f\xa6\xe7\x7d\xb9\xfc\x7c\x73\x74\xab\x71\x9c\x1f\xdc\xb5\xf6\xd9\xf1\xf8\x83\x3b\x9a\x5e\xbd\xfe\x4f\x86\xd1\xe5\xe0\x0d\xbf\xb9\xb9\x1d\x0c\xe6\xb7\xd9\xf9\xbd\xe6\xed\x77\x5f\xca\x4f\xac\xbf\x18\xdf\x51\x73\xb8\xfb\xea\xe8\xfa\xd3\x59\xf1\xa7\xf8\x3b\xb3\xf3\xd4\x35\x3e\x41\xfb\x61\x2b\x3c\xa8\x91\x7e\x51\x5c\xfc\xf3\x8c\x4e\x17\x17\xfb\x9f\x8e\x76\x5f\x3f\x7c\x68\xdf\xed\xce\xbf\x5e\xbd\xe9\x5c\xbc\x15\x4c\xf3\x8f\xe7\x59\xa7\x7d\xd8\x41\x7a\xf2\x95\x4f\x06\x7f\xb5\xb8\xd6\xfd\xf7\x47\x83\xa9\xb4\x2c\x50\xba\x66\x68\x90\xb2\x45\x23\x2b\x65\xea\x91\xbe\xd1\xac\x8d\x88\x19\x97\x4c\xcd\x42\xca\xd8\xeb\x29\x4a\x77\xb5\xea\xb4\xc6\x4e\x81\xd6\xd2\x11\xee\xec\xc1\xfa\x1e\x34\xd0\xb3\xd4\x6f\xf3\x6c\x45\x0c\x2b\xf3\xe0\x65\x02\x3b\xdb\xcd\xbe\xd3\x04\x83\xae\x34\xb2\x3e\xce\xfe\x68\x54\xe8\xd2\x0c\x1d\xce\x5d\x63\x47\x7b\x34\x4f\x9d\x80\x6a\x82\x05\x01\x2d\x5d\x1e\x54\x58\xba\x03\xbb\x50\xe9\xf0\x38\xd5\xfc\x4e\x0c\xe3\xd3\x70\x83\xee\xcd\xd0\xe6\x6a\xd6\xf8\x9e\x69\x39\x72\x7e\x72\x56\x0d\x4d\xd2\x0c\x73\xce\xb0\x7e\xfa\xd8\x3c\x7d\xf1\xa2\xc6\x18\x7a\xf8\xf4\x16\x2b\x25\x1c\xd7\x8d\xfa\x5b\xc1\x19\x3e\x1a\xa1\xe9\x02\x49\x05\x4f\x27\x64\xaf\x76\xa6\x05\x4d\xd1\xe7\xa1\x0b\xa4\x1a\xc8\xe4\x79\x1d\x4f\x71\xb6\xe8\x3e\xae\x94\xac\xb2\xd0\xdc\x7a\x99\x3c\x63\x0c\xf1\xd6\x93\x66\x8d\x09\x20\xa4\xce\x99\x06\x59\xa2\x7b\x95\x1e\x2a\x82\xe5\x83\x8b\xec\xc1\x5a\xf0\xf6\x8d\x27\x71\x3e\x54\xa4\x1e\x8c\x67\xcc\xf4\x2a\xd7\x76\x6e\x1b\xe8\x1d\xe0\x05\xaa\xd2\x3d\x5f\x74\xbf\x76\xa2\x9e\xa0\xc7\x3d\xff\xd8\x6b\x3d\x6f\xc9\x94\x1a\x78\xea\x36\x48\x40\xe2\x0c\xce\xd6\xfb\xcb\xdb\x27\x05\xb5\x18\x6f\x9a\x53\xc9\x06\xb1\x65\x9a\xa2\xb5\x64\x53\xe3\x0d\xfc\xd1\x93\xb5\x75\x67\x4a\x73\x64\x2f\xb7\x0d\xdc\x0a\xc2\x4f\x8a\x67\x5b\x29\x1a\xa3\xcc\x6f\xaa\x3c\xa7\x5c\xfc\x15\x95\xdb\xab\x3a\x3e\xc4\x51\xf5\x1a\x8b\xa3\xea\x77\xcf\x7f\x03\x00\x00\xff\xff\x72\xd1\xf6\x93\x08\x0d\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 3336, mode: os.FileMode(420), modTime: time.Unix(1551202530, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

