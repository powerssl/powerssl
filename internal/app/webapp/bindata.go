// Package webapp Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../../../web/app/index.html
package webapp

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x5f\x53\xdb\x3a\x16\x7f\xef\xa7\x38\xd5\xde\x1d\x92\x01\xdb\x09\x01\x0a\xc1\xce\x4c\x96\x2d\x5d\x28\x6d\x13\x52\x0a\xb7\x6f\x8a\x25\xdb\x4a\x64\x49\x95\xe4\xfc\x69\x87\xef\xbe\x23\x3b\x26\x4e\x4b\xbb\xbd\xb3\x73\x1f\xc0\xd2\xd1\xd1\xf9\x7f\x7e\x47\x09\x5f\x12\x19\xdb\xb5\xa2\x90\xd9\x9c\x0f\x5e\x84\xee\x03\x1c\x8b\x34\x42\x54\xa0\xc1\x0b\x80\x30\xa3\x98\xb8\x05\x40\x98\x53\x8b\x21\xce\xb0\x36\xd4\x46\xa8\xb0\x89\x77\x8a\x9a\x47\x02\xe7\x34\x42\x0b\x46\x97\x4a\x6a\x8b\x20\x96\xc2\x52\x61\x23\xb4\x64\xc4\x66\x11\xa1\x0b\x16\x53\xaf\xdc\x1c\x00\x13\xcc\x32\xcc\x3d\x13\x63\x4e\xa3\xee\x01\x98\x4c\x33\x31\xf7\xac\xf4\x12\x66\x23\x21\x6b\xd1\x9c\x89\x39\x68\xca\x23\x64\xec\x9a\x53\x93\x51\x6a\x11\x64\x9a\x26\x11\xca\xac\x55\xa6\x1f\x04\x39\x5e\xc5\x44\xf8\x53\x29\xad\xb1\x1a\x2b\xb7\x89\x65\x1e\x3c\x11\x82\x23\xbf\xe3\x77\x82\xd8\x98\x2d\xcd\xcf\x99\xf0\x63\x63\x10\x30\x61\x69\xaa\x99\x5d\x47\xc8\x64\xb8\x77\x7a\xe4\xbd\x11\xc7\xbd\xd3\xa3\xd5\x97\x71\x17\xcb\xfb\x87\xe1\x7e\xe7\xf8\xf4\xf6\x61\xb4\x1a\xa5\x27\xc9\xfa\xe8\xea\x7e\xf1\xf1\x7d\xd6\x79\x7d\x78\xd2\x7b\xc8\x2f\xe3\x6b\x3e\x19\x2e\xd9\x9b\xf4\x72\x78\x1f\x90\x21\x9b\x9c\x5c\x3f\xe4\x08\x62\x2d\x8d\x91\x9a\xa5\x4c\x44\x08\x0b\x29\xd6\xb9\x2c\x4c\xed\x94\x65\x96\xd3\xc1\x48\x2e\xa9\x9e\x4c\x6e\x60\x58\xd8\x2c\x0c\x2a\x62\xc5\x50\xfa\x5a\xad\x01\xa6\x92\xac\x0f\xca\x1c\xc1\xb7\x0d\x09\x20\xa3\x2c\xcd\x6c\x1f\xba\x9d\xce\x3f\xcf\x9f\xa8\x39\xd6\x29\x13\x7d\xe8\x6c\x49\x72\x41\x75\xc2\xe5\xb2\x0f\x19\x23\x84\x8a\xed\x89\xc2\x84\x30\x91\x36\xb8\x1f\x37\x5f\xa5\x69\x43\x15\xc0\x32\x63\x96\x7a\x46\xe1\x98\xf6\xdd\xa1\xb7\xd4\x58\x9d\xff\x8c\xc1\xcb\xe5\x57\xef\x7f\x73\xc9\xdf\xe0\x79\x96\x43\x6a\x52\x12\xfb\x30\xd5\x14\xcf\x3d\x47\xf8\xde\x83\x7f\x70\x99\x32\xd1\x70\x62\x2a\xad\x95\xf9\x4e\x64\x38\x4d\xec\x0e\x41\x49\xc3\x2c\x93\xa2\x0f\x78\x6a\x24\x2f\x2c\xdd\x9e\xe9\x2a\xdc\x0d\x6e\x2b\x55\x1f\x3a\x6a\x75\x0e\x3b\xba\xc3\xe0\x29\x79\x61\x50\xb7\x4f\xe8\x72\x58\x37\x0b\x66\x02\xb4\xe4\x34\x42\x6e\x89\x20\xe6\xd8\x98\x08\xb9\x7e\xc1\x4c\x50\x8d\xa0\x94\x10\x21\xc2\x8c\xe2\x78\xdd\x07\x21\x05\x45\x75\x39\x84\x59\xb7\xbe\x92\x5b\xef\x18\x35\xea\x48\xa9\x30\xc8\xba\x1b\x3d\x81\x93\xbe\x59\x13\xb6\xf8\x4d\x35\xe7\x9b\x1a\xf2\x4a\xf7\xba\x34\xdf\xea\x55\xb5\x08\x4e\x31\x41\x83\xd0\x28\x2c\x6a\x52\x22\x85\xf5\x96\x65\x4d\x7a\x53\xc9\x09\x1a\x0c\xc1\xca\x39\x15\xc0\x0c\x60\x48\xa4\xce\x41\x26\x80\x41\x61\x63\x5c\xc2\xfc\x30\x70\xf7\x07\xf0\x6f\x09\x42\x5a\x30\x19\xd6\x14\xd6\xb2\xd0\x30\x1c\x5d\x55\x77\xfd\x30\x50\x5b\xed\x83\xd0\x95\xe5\x46\x9f\x2b\x0c\x13\x6b\xc9\x39\x9e\x72\x8a\x06\x61\x2c\x09\x1d\x84\xc1\xe6\xa3\x74\xf9\xbf\x79\x79\x5a\x58\x2b\xc5\x36\x0a\x6a\x0d\x53\x2b\xdc\x9f\x67\x68\x2c\x05\xc1\x7a\x8d\x80\x60\x8b\xbd\x98\x33\x35\x95\x58\x13\xcf\x62\x9d\x3a\xb4\x73\x62\xd1\xe0\x42\xaa\x75\x18\x54\x82\xb6\xe2\xc3\x80\xb0\x45\x23\xce\x8c\x44\xa8\xac\xbe\x6d\xe4\x58\xa2\x71\x4e\xc1\xe8\x38\x42\xdf\xbe\x81\xef\x1a\xfe\xee\xf6\x0a\x1e\x1f\x11\x54\xe8\x88\x5c\x1f\xa3\x4d\x57\xd7\xbb\xf2\xd6\x54\x6a\x42\x75\x84\x3a\xa8\x92\xcc\x8c\xa5\x82\xea\x4b\x77\x86\x06\x61\x50\xc9\xfe\xd1\x14\x13\x6b\xa6\x6c\xa5\xb3\x06\x4a\xe7\x86\x3f\xfb\x52\x50\xbd\x2e\x01\xb2\x5a\x7a\x3d\xbf\xe7\x77\x4b\x3c\x9c\xfd\x00\x87\x87\xc7\x27\xde\x65\xaa\x2e\xa6\xc1\xdb\xeb\x31\xbf\x79\x9f\x7c\x28\xce\xba\x16\xf7\x0e\x65\xf0\xfe\xdd\xe7\x15\xb7\xcb\x5b\x79\x3a\xb6\xf9\xfc\xdd\x2d\x19\x16\xa7\xd1\xcf\xa1\x2f\x0c\x2a\x93\x7e\x65\x1f\x11\x33\xe3\xc7\x5c\x16\x24\xe1\x58\xd3\xd2\x48\x3c\xc3\xab\x80\xb3\xa9\x09\x94\x54\x8a\x6a\x7f\x66\x82\xae\xdf\x3d\xf4\xcf\x82\x22\x27\x35\xf1\x79\xeb\x1d\x98\x0f\xd5\xfb\x69\x9a\x9d\xfd\x6b\xff\xcf\xee\xf8\xad\x5d\xf4\x6e\xc5\xab\xfb\x5e\x9e\x8e\x56\xd9\xdd\xd9\xdb\x60\x12\x8f\xcd\x70\xf4\x2a\xbb\x63\xd3\x87\xde\xd9\xec\x55\x82\xe7\x97\x23\x33\x5f\x3c\x14\x66\x91\xe0\xce\xf4\x68\xfc\x7f\x79\xf4\xbb\xa3\x69\xf6\xfd\x64\x7a\xde\x97\xeb\xcf\xb7\x27\x13\x45\x67\xd9\xd1\x5d\xe7\x90\x9c\xce\x3e\xd8\x93\xc5\xcd\xeb\xff\x24\x34\xb8\x1e\xbf\x61\xb7\xb7\x93\xf1\x78\x35\x49\x2e\xef\x15\xeb\xbe\xfb\x52\x7c\x22\xc3\xf5\xec\x0e\xeb\xe3\xfd\x57\x27\xa3\x4f\x17\xf9\x9f\xfc\xef\xcc\xce\x53\xd7\xb8\x04\x1d\xfa\x1d\xff\xa8\x41\xfa\x45\x71\xb1\xcf\x4b\xbc\x58\x5f\x1d\x7e\x3a\xd9\x7f\xfd\xf0\xa1\x7b\xb7\xbf\xfa\x7a\xf3\xa6\x77\xf5\x96\x13\xc5\x3e\x5e\x26\xbd\xee\x71\x8f\xe2\xb3\xaf\x6c\x3e\xfe\xab\xc5\x55\xf7\xdf\x1f\x2d\x22\xe3\x22\xa7\xc2\xb6\x7d\x4d\x31\x59\xb7\x92\x42\xc4\x0e\xe9\x5b\xed\xc6\x88\x58\x32\x41\xe4\xd2\xc7\x84\xbc\x5e\x50\x61\x6f\x36\x9d\xd6\xda\xcb\xa9\x31\x38\xa5\x7b\x07\x50\xdf\x83\x16\x75\x2c\xcd\xdb\x2c\xd9\x10\xfd\xca\x3c\x78\x19\xc1\xde\x6e\xb3\xef\xb5\x41\x53\x5b\x68\xd1\x1c\x67\x7f\xb4\x2a\x74\x69\xfb\x96\xae\x6c\x6b\x4f\x39\x34\x8f\x2d\x87\x6a\x82\x79\x1e\x26\x44\x43\x29\x68\x74\x35\x74\xeb\xc7\x47\x47\x2d\x6c\xe6\x55\x08\xbb\x07\xfb\x50\x69\x76\xe8\xd5\xfe\x4e\x38\x61\x0b\x7f\x8b\xf9\x6d\xdf\x64\x72\xd9\xfa\x9e\xa9\x1c\x44\x3f\x39\xab\x46\x29\x6a\xfb\x19\x23\xb4\x79\xfa\xd8\x3e\x7f\xf1\xa2\xc1\xe8\x3b\x50\x75\x7e\x48\xc9\x2d\x53\xad\xe6\x0b\xc2\x6a\x96\xa6\x54\xf7\x01\xc5\x9c\xc5\x73\x74\xd0\x38\x53\x1c\xc7\xd4\x65\xa7\x0f\xa8\x1a\xd3\xe8\x79\x1d\x4f\xd1\x37\xd4\x7e\xdc\x28\xd9\xe4\xa6\xbd\xf3\x5e\x79\xc6\x18\xe4\xac\x47\xed\x06\x13\x80\x8f\xad\xd5\x2d\x54\x62\x7e\x95\x34\xcc\xbd\xf2\x19\x86\x0e\xa0\x16\xbc\x7b\xe3\x49\x9c\x0b\x15\x6a\x06\xe3\x19\x33\x9d\xca\xda\xce\x5d\x03\x9d\x03\x2c\xa7\xb2\xb0\xcf\x97\xe2\xaf\x9d\x68\x26\xe8\xf1\xc0\x3d\x01\x3b\xcf\x5b\xb2\xc0\x1a\x9e\x7a\x10\x22\x10\x74\x09\x17\xf5\xfe\x7a\xf2\xa4\xa0\x11\xe3\x6d\xcb\x4a\xd1\x42\xa6\x88\x63\x6a\x0c\xda\x56\x7e\x8b\xfe\xe8\x49\x6d\xdd\x85\x54\x8c\x92\x97\xbb\x06\xee\x04\xe1\x27\xc5\xb3\xab\x94\x6a\x2d\xf5\x6f\xaa\xbc\xc4\x8c\xff\x15\x95\xbb\xab\x26\x6a\x84\x41\xf5\x46\x0b\x83\xea\xd7\xd0\x7f\x03\x00\x00\xff\xff\x14\xbc\x14\x6e\x1e\x0d\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
