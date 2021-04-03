// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../web/index.html
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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\x5d\x73\xda\x38\x14\x7d\xcf\xaf\x50\xf5\xd2\x87\xd4\x16\x0e\x84\x90\x5d\x9b\x19\x6f\xa7\xc9\x36\x5f\xe5\x63\x29\xb4\x6f\xb2\x2d\x63\x39\xb2\xa4\x48\xd7\x06\xa6\xd3\xff\xbe\x63\x0c\x81\x50\x32\x93\x99\xf8\x45\xd2\xd5\x3d\xf2\x39\x47\xd2\x95\xff\x21\x51\x31\xac\x34\x43\x19\x14\xa2\x7f\xe2\xd7\x0d\x12\x54\xce\x03\xcc\x24\xae\x03\x8c\x26\xfd\x13\x84\x10\xf2\x0b\x06\x14\xc5\x19\x35\x96\x41\x80\x4b\x48\x9d\x1e\x7e\x31\xa5\x24\x30\x09\x01\x5e\xf0\x04\xb2\x20\x61\x15\x8f\x99\xb3\x1e\x7c\x42\x5c\x72\xe0\x54\x38\x36\xa6\x82\x05\xde\x27\x64\x33\xc3\xe5\xa3\x03\xca\x49\x39\x04\x52\x61\x24\x69\xc1\x02\x5c\x71\xb6\xd0\xca\xc0\x76\xe9\x0f\x8e\x23\xb8\x7c\x44\x86\x89\x00\x5b\x58\x09\x66\x33\xc6\x00\xa3\xcc\xb0\x34\xc0\x19\x80\xb6\x7f\x11\x52\xd0\x65\x9c\x48\x37\x52\x0a\x2c\x18\xaa\xeb\x41\xac\x0a\xf2\x1c\x20\x1d\xb7\xe5\xb6\x48\x6c\xed\x2e\xe6\x16\x5c\xba\xb1\xb5\x18\x71\x09\x6c\x6e\x38\xac\x02\x6c\x33\xda\xee\x75\x9c\x6b\x79\xde\xee\x75\x96\x4f\x43\x8f\xaa\xe9\x2c\x3c\x6d\x9d\xf7\x46\xb3\xc1\x72\x30\xef\xa6\xab\xce\xd7\x69\xf5\xdf\x43\xd6\xfa\x72\xd6\x6d\xcf\x8a\xab\xf8\x46\x8c\xc3\x05\xbf\x9e\x5f\x85\x53\x92\x84\x7c\xdc\xbd\x99\x15\x18\xc5\x46\x59\xab\x0c\x9f\x73\x19\x60\x2a\x95\x5c\x15\xaa\xb4\xd8\x71\x36\xca\xd6\xb2\x1a\x15\x84\x5a\xcb\xe0\x28\xb3\x43\xdd\x1b\x30\x70\x10\xac\x3f\x50\x0b\x66\xc6\xe3\x3b\x14\x96\x90\xf9\xa4\x09\x36\x09\x36\x36\x5c\x43\x33\xa8\xbf\x8a\x1a\xb4\x60\x51\xa8\xf5\x64\xf4\x15\x05\xe8\xe3\xaf\x5f\xc8\x9d\x3e\x07\x7e\xff\xfe\xf8\x77\x03\x24\xfb\xc8\xda\xfe\x66\x8c\xac\x89\x77\x76\xc7\x2a\x61\x6e\xfe\x54\x32\xb3\x5a\xdb\xdc\x74\x9d\xb6\xdb\x76\xbd\x35\xf7\xfc\x0f\x53\xcf\xce\xbb\xce\xd5\x5c\x7f\x8e\xc8\xed\xcd\x50\xdc\x3d\xa4\xdf\xca\x4b\x0f\x68\xfb\x4c\x91\x87\xfb\x9f\x4b\x01\x8b\x91\xea\x0d\xa1\x78\xbc\x1f\x25\x61\xd9\x0b\x5e\x35\xb0\xbf\xa5\xf8\xec\xe4\x3e\xc3\xad\x95\x47\x29\xf5\xdf\xa8\x2e\x91\xb9\x75\x63\xa1\xca\x24\x15\xd4\xb0\xb5\x44\x9a\xd3\x25\x11\x3c\xb2\x44\x2b\xad\x99\x71\x73\x4b\x3c\xd7\x3b\x73\x2f\x49\x59\x24\xdb\xe0\x71\xed\xf5\x81\x0a\xf5\x43\x34\xcf\x2e\xff\x39\xfd\xe1\x0d\x6f\xa1\x6a\x8f\xe4\xc5\xb4\x5d\xcc\x07\xcb\x6c\x72\x79\x4b\xc6\xf1\xd0\x86\x83\x8b\x6c\xc2\xa3\x59\xfb\x32\xbf\x48\xe9\xe3\xd5\xc0\x3e\x56\xb3\xd2\x56\x29\x6d\x45\x9d\xe1\x3b\xfd\x78\xc9\xef\x8d\x46\xbc\xf5\x56\xe5\x87\x47\xf7\xb8\x05\x37\x3f\x47\xdd\xb1\x66\x79\xd6\x99\xb4\xce\x92\x5e\xfe\x0d\xba\xd5\xdd\x97\x7f\x53\x46\x6e\x86\xd7\x7c\x34\x1a\x0f\x87\xcb\x71\x7a\x35\xd5\xdc\xbb\x7f\x2a\xbf\x27\xe1\x2a\x9f\x50\x73\x7e\x7a\xd1\x1d\x7c\xff\x5c\xfc\x10\xef\xb4\xe0\x0f\x8a\x87\x2e\x1c\x03\x51\xad\x0f\x52\x7d\xd2\x14\x45\x3f\x52\xc9\xaa\x7f\xe2\x17\x94\x4b\x14\x0b\x6a\x6d\x80\xeb\x1a\x48\xb9\x64\x06\x23\xa3\x04\x0b\x70\x3d\xb9\xbd\xb3\x99\xb7\x4d\x2b\xc0\x39\xc7\x87\xb7\x37\xf3\x36\x79\xa9\x32\xc5\xee\xe2\xfa\x09\xaf\xb6\xb8\x7a\xc6\x99\x1b\x55\x6a\xbc\x4b\x68\x8a\x09\x8d\x98\x40\xa9\x32\x01\x2e\x2d\x33\x75\x2d\xc5\xfd\xc9\xa6\xe7\x93\xf5\xf4\x01\x84\x4b\x5d\x02\xa2\x25\xa8\x58\x15\x5a\x30\x60\x7b\xd8\x17\x7f\xac\x55\x19\x25\x30\xe2\xc9\x7e\x4a\xfd\x70\x04\x18\xd8\x12\xf6\xd8\xf8\x24\xe1\xd5\x3b\xd8\x6b\x6a\xed\x42\x99\x04\xf7\x07\x9b\xde\xdb\xd9\xc7\xa5\x31\x4c\x82\xf3\xbc\xc6\xeb\x2a\x76\x29\x8d\x8a\xdd\x6f\x5f\x55\x12\x95\x00\xea\x79\xa7\x23\x90\x28\x02\xe9\x68\xc3\x0b\x6a\x56\xdb\x75\x6c\x19\x15\x1c\x70\x7f\xbc\x6e\x7d\xd2\x80\x36\x1b\x4b\x9a\x9d\xf5\x49\x7d\x2a\xea\x76\x73\x82\x48\xf3\xfa\xfe\x1f\x00\x00\xff\xff\xf5\xb8\xb6\x33\x8e\x07\x00\x00")

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
