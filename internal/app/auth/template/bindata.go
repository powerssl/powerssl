// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../web/auth/index.html
package template

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xdb\x73\x9b\xb8\x17\x7e\xef\x5f\xa1\xea\xa5\x0f\x29\xc8\xc4\x8e\xe3\xfc\x7e\xe0\x19\xb6\xd3\x64\x9b\x5b\x7d\x59\xd7\x6e\xdf\x04\x08\x23\x22\x24\x45\x3a\x60\x7b\x3a\xfd\xdf\x77\x30\xc1\xce\xc5\xd9\xcd\x6c\x9f\x38\xb7\x0f\x7d\xe7\xe3\xe8\xe0\xbf\x4f\x54\x0c\x1b\xcd\x50\x06\x85\x18\xbe\xf3\xeb\x07\x12\x54\x2e\x03\xcc\x24\x1e\xbe\x43\xc8\xcf\x18\x4d\x6a\x03\x21\xbf\x60\x40\x51\x9c\x51\x63\x19\x04\xb8\x84\xd4\x19\xe0\xc7\x29\x49\x0b\x16\xe0\x8a\xb3\x95\x56\x06\x30\x8a\x95\x04\x26\x21\xc0\x2b\x9e\x40\x16\x24\xac\xe2\x31\x73\xb6\xce\x47\xc4\x25\x07\x4e\x85\x63\x63\x2a\x58\xe0\x7d\x44\x36\x33\x5c\xde\x39\xa0\x9c\x94\x43\x20\x55\xfb\xea\xf7\x8e\x23\xb8\xbc\x43\x86\x89\x00\x5b\xd8\x08\x66\x33\xc6\x00\xa3\xcc\xb0\x34\xc0\x19\x80\xb6\xff\x23\xa4\xa0\xeb\x38\x91\x6e\xa4\x14\x58\x30\x54\xd7\x4e\xac\x0a\xb2\x0b\x90\x9e\xdb\x71\x3b\x24\xb6\x76\x1f\x73\x0b\x2e\xdd\xd8\x5a\x8c\xb8\x04\xb6\x34\x1c\x36\x01\xb6\x19\xed\x0e\x7a\xce\x85\x3c\xe9\x0e\x7a\xeb\xfb\xb1\x47\xd5\x7c\x11\x1e\x75\x4e\x06\x93\xc5\x68\x3d\x5a\xf6\xd3\x4d\xef\xcb\xbc\xfa\xeb\x36\xeb\x7c\x3e\xee\x77\x17\xc5\x79\x7c\x29\xa6\xe1\x8a\x5f\x2c\xcf\xc3\x39\x49\x42\x3e\xed\x5f\x2e\x0a\x8c\x62\xa3\xac\x55\x86\x2f\xb9\x0c\x30\x95\x4a\x6e\x0a\x55\x5a\xec\x38\x0f\x9d\xfd\x53\x5b\x84\x5a\xcb\xe0\x10\xd5\x07\x30\x70\x10\x6c\x38\x52\x2b\x66\xa6\xd3\x6b\x14\x96\x90\xf9\xa4\x09\x36\x05\x36\x36\x5c\x43\xe3\x20\x54\x51\x83\x56\x2c\x0a\xb5\x9e\x4d\xbe\xa0\x00\x7d\xf8\xf9\x13\xb9\xf3\x5d\xe0\xd7\xaf\x0f\xff\x6f\x60\xe4\x31\xae\x16\xbf\xf1\x91\x35\xf1\x5e\xec\x58\x25\xcc\xcd\xef\x4b\x66\x36\x5b\x91\x1b\xd3\xe9\xba\x5d\xd7\xdb\x12\xcd\x5f\x48\x7a\x7c\xd2\x77\xce\x97\xfa\x53\x44\xae\x2e\xc7\xe2\xfa\x36\xfd\x5a\x9e\x79\x40\xbb\xc7\x8a\xdc\xde\xfc\x58\x0b\x58\x4d\xd4\x60\x0c\xc5\xdd\xcd\x24\x09\xcb\x41\xf0\xaa\x7c\xc3\x96\xe2\x4e\xc7\xc7\x0c\x5b\xdd\x0e\x52\x1a\xbe\xb1\xbb\x44\xe6\xd6\x8d\x85\x2a\x93\x54\x50\xc3\xb6\x2d\xd2\x9c\xae\x89\xe0\x91\x25\x5a\x69\xcd\x8c\x9b\x5b\xe2\xb9\xde\xb1\x7b\x46\xca\x22\x69\x83\x87\x7b\xaf\xc7\x29\xd4\xb7\xd1\x32\x3b\xfb\xe3\xe8\xbb\x37\xbe\x82\xaa\x3b\x91\xa7\xf3\x6e\xb1\x1c\xad\xb3\xd9\xd9\x15\x99\xc6\x63\x1b\x8e\x4e\xb3\x19\x8f\x16\xdd\xb3\xfc\x34\xa5\x77\xe7\x23\x7b\x57\x2d\x4a\x5b\xa5\xb4\x13\xf5\xc6\xbf\xa9\xc7\x53\x7e\x6f\x14\xe2\xad\x77\x2a\x7f\x3e\xa7\x87\x25\xb8\xfc\x31\xe9\x4f\x35\xcb\xb3\xde\xac\x73\x9c\x0c\xf2\xaf\xd0\xaf\xae\x3f\xff\x99\x32\x72\x39\xbe\xe0\x93\xc9\x74\x3c\x5e\x4f\xd3\xf3\xb9\xe6\xde\xcd\x7d\xf9\x2d\x09\x37\xf9\x8c\x9a\x93\xa3\xd3\xfe\xe8\xdb\xa7\xe2\xbb\xf8\x4d\x09\x5e\x50\x7c\xae\xc2\x21\x10\xd5\xfa\x45\xa9\x4f\xda\xa5\xe8\x47\x2a\xd9\xb4\x2b\x90\x72\x89\x8c\x12\x2c\xc0\xb5\x89\x51\x2c\xa8\xb5\x01\xae\xb7\x20\xe5\x92\x19\xdc\xde\x45\x3f\xf3\xda\x64\x01\xce\x09\x7e\x7e\x8b\x33\x6f\x57\x99\x2a\x53\xb4\x0e\x42\x7e\xc2\xab\x16\x59\x67\x9c\xa5\x51\xa5\xc6\xfb\x82\x7a\xa9\xd0\x88\x09\x94\x2a\x13\xe0\xd2\x32\x53\xef\x64\x3c\x9c\x3d\x58\x3e\xd9\xa6\x9f\x00\xb8\xd4\x25\xa0\xfa\x2f\x10\x60\x60\x6b\xc0\x4f\x4e\xa8\xd9\x1b\x25\x30\xe2\xc9\xa3\x17\x22\x5a\x82\x8a\x55\xa1\x05\x03\xf6\xf8\xa0\x3d\x55\x92\xf0\xea\x3f\x33\xd7\xd4\xda\x95\x32\x09\x1e\x8e\x1e\xac\x7f\x63\xbe\x43\xbc\xce\x7e\x5f\xf2\x94\x7d\x5c\x1a\xc3\x24\x38\xfb\x43\x5f\xed\x22\x2a\x01\x94\x7c\x38\xd2\x96\x51\xc1\xf7\x72\x45\x20\x51\x04\xd2\xd1\x86\x17\xd4\x6c\xf0\x70\xba\xcd\xfb\xa4\x01\xed\x3e\x29\xd9\x7f\x53\x9f\xd4\x93\xd2\x8c\x54\x33\x49\x3e\x69\xfe\xc4\x7f\x07\x00\x00\xff\xff\xbc\x34\xa3\xfa\x9a\x07\x00\x00")

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
