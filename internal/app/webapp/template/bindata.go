// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../web/app/index.html
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\xc1\x6a\xdc\x30\x10\x86\xef\x79\x8a\xa9\x2e\xb9\xc4\x6b\x72\x28\x94\x54\x16\x2c\x29\x94\x40\x0e\x4b\x42\xe9\x79\x2c\xcd\xae\xd5\x68\x35\x42\x1a\x67\x31\x21\xef\x5e\x14\xaf\x77\xb7\x81\x1e\x6c\x69\x46\xdf\xfc\x1a\xfe\xb1\xf5\x17\xc7\x56\xa6\x44\x30\xc8\x3e\x98\x2b\x5d\x17\x08\x18\x77\x9d\xa2\xa8\xcc\x15\x80\x1e\x08\x5d\xdd\x00\xe8\x3d\x09\x82\x1d\x30\x17\x92\x4e\x8d\xb2\x6d\xbe\xa9\xcb\xa3\x88\x7b\xea\xd4\xab\xa7\x43\xe2\x2c\x0a\x2c\x47\xa1\x28\x9d\x3a\x78\x27\x43\xe7\xe8\xd5\x5b\x6a\x3e\x82\x1b\xf0\xd1\x8b\xc7\xd0\x14\x8b\x81\xba\xdb\x1b\x28\x43\xf6\xf1\xa5\x11\x6e\xb6\x5e\xba\xc8\x8b\xb4\x78\x09\x64\x36\x7c\xa0\xfc\xfc\xfc\x08\xeb\x94\x74\x3b\xe7\xe6\xf3\x62\xb3\x4f\x32\x07\x00\xaf\x98\x21\x55\xb6\x94\x00\x1d\xbc\x1d\xd3\x00\x98\xfc\xda\xb9\x7c\x07\xd7\x6f\x6f\xb0\x5a\x6f\x1e\x6a\x04\xef\xef\xd7\x37\x67\x64\x94\xe1\xd7\xd3\xc3\x82\xcc\xd1\xbf\xc8\x2e\x27\xfb\x9b\xfa\x33\xf5\xf3\x69\x73\x3f\x27\x2a\x78\xe4\xde\xbf\xcf\xad\xb5\x97\xbd\x1d\x1b\x85\x92\x6d\xa7\x5a\x2c\x85\xa4\xb4\xfd\x18\x5d\xa0\xd5\x9f\xa2\xcc\x25\xad\xdb\xc5\x76\xdd\xb3\x9b\x16\x93\xd1\x47\xc8\x1c\xa8\x53\x75\xab\xc0\x06\x2c\xa5\x53\xd5\x67\xf4\x91\xb2\x82\x22\x53\x3d\x76\xbe\xa4\x80\xd3\x1d\x44\x8e\xa4\x16\x6f\xf4\x70\xbb\x94\xec\xa5\xf9\xaa\x3e\x99\x3a\xdc\x9e\xc0\xb4\x70\x81\xd0\x29\xa3\x4b\xc2\xb8\xa4\xb6\x1c\xa5\x39\x90\xdf\x0d\xd2\xf4\x1c\x9c\x32\x6b\x10\x7e\xa1\x08\xbe\x00\xc2\x96\xf3\x1e\x78\x0b\x08\x09\x4b\x39\x70\x76\x2b\xdd\xd6\x7a\x03\x3f\x18\x22\x0b\x94\x01\x33\xc1\xc4\x63\x86\xf5\xe6\x61\xae\x5d\xe9\x36\x9d\x6f\x37\x3a\x65\x5a\xee\x4b\x99\x9a\x62\x33\x87\x80\x7d\x20\x65\xb4\x65\x47\x46\xb7\xc7\x25\xe5\x8f\xf7\x45\xf1\x69\x58\xba\x1f\x45\x38\x9e\x6d\x4a\x13\xf4\x12\xeb\xd3\x14\xb2\x1c\x1d\xe6\x49\x81\x43\xc1\xc6\x06\x9f\x7a\xc6\xec\x1a\xc1\xbc\xab\xdf\x77\xd5\x57\xe6\x9e\xd3\xa4\xdb\x59\xe8\xbf\xca\x81\x77\x3c\xca\x49\xdb\x61\xdc\x51\x56\xe6\xf1\x23\xfd\xb9\xfa\xd4\xac\x6e\xeb\x14\xe7\x71\xcf\x53\xd6\xed\xfc\x1f\xfe\x0d\x00\x00\xff\xff\xdb\x49\x90\x37\x98\x03\x00\x00")

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
