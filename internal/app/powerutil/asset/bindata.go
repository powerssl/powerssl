// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../../db/migrations/20210130204142_create.down.sql
// ../../../../db/migrations/20210130204142_create.up.sql
package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var __20210130204142_createDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x4c\xce\x4d\x8d\x2f\x4e\x2d\x2a\x4b\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x40\xfb\xa3\x94\x19\x00\x00\x00")

func _20210130204142_createDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210130204142_createDownSql,
		"20210130204142_create.down.sql",
	)
}

func _20210130204142_createDownSql() (*asset, error) {
	bytes, err := _20210130204142_createDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210130204142_create.down.sql", size: 25, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210130204142_createUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x6a\xc3\x30\x14\x85\xe1\xdd\x4f\x71\xc8\x64\x43\xbb\x75\xcb\xa4\x26\x0a\x15\xb5\xe5\x20\xcb\x6d\x32\x09\x61\x5d\x82\xc0\xb1\x8d\x24\x07\xf2\xf6\xa5\xa1\xf1\x52\x5a\xc8\xfe\x9d\x7b\xff\x8d\xe2\x4c\x73\xf0\x83\xe6\xb2\x11\xb5\x84\xd8\x41\xd6\x1a\xfc\x20\x1a\xdd\x60\x35\xcf\xde\x3d\x8f\x31\x4e\xab\x75\x96\xfd\x60\xcd\x5e\x4b\x0e\xdb\x9d\xc9\x44\x0a\x17\x0a\x11\x79\x06\x00\xde\xa1\x6d\xc5\x16\x5b\xbe\x63\x6d\xa9\xf1\x3d\x36\x27\x1a\x28\xd8\x44\xe6\xf2\x92\x17\x4f\x37\xe7\x7c\x9c\x7a\x7b\x35\x83\x3d\x13\x3e\x98\xda\xbc\x31\x75\xfb\x2a\xdb\xb2\xbc\x93\x40\x5d\x1a\xc3\xd5\xcc\xa1\xff\xc3\xf8\x21\xd1\x29\xd8\xe4\xc7\xe1\xbf\x53\x5d\x20\x9b\xc8\x19\x9b\xa0\x45\xc5\x1b\xcd\xaa\xfd\x42\x96\x56\x59\x7f\xde\xf3\xe6\xc9\xba\x87\x06\x8e\x7a\xfa\x3d\x58\x02\xf6\x4a\x54\x4c\x1d\xf1\xce\x8f\xc8\xbd\x2b\xb2\x62\x9d\x7d\x05\x00\x00\xff\xff\x58\x6d\x14\x11\x78\x01\x00\x00")

func _20210130204142_createUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210130204142_createUpSql,
		"20210130204142_create.up.sql",
	)
}

func _20210130204142_createUpSql() (*asset, error) {
	bytes, err := _20210130204142_createUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210130204142_create.up.sql", size: 376, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
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
	"20210130204142_create.down.sql": _20210130204142_createDownSql,
	"20210130204142_create.up.sql":   _20210130204142_createUpSql,
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
	"20210130204142_create.down.sql": &bintree{_20210130204142_createDownSql, map[string]*bintree{}},
	"20210130204142_create.up.sql":   &bintree{_20210130204142_createUpSql, map[string]*bintree{}},
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
