// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../db/migrations/20210130204142_create_acme_servers_table.down.sql
// ../../../db/migrations/20210130204142_create_acme_servers_table.up.sql
// ../../../db/migrations/20210202110208_create_acme_accounts_table.down.sql
// ../../../db/migrations/20210202110208_create_acme_accounts_table.up.sql
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

var __20210130204142_create_acme_servers_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\x48\x4c\xce\x4d\x8d\x2f\x4e\x2d\x2a\x4b\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x15\x02\x83\x08\x19\x00\x00\x00")

func _20210130204142_create_acme_servers_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210130204142_create_acme_servers_tableDownSql,
		"20210130204142_create_acme_servers_table.down.sql",
	)
}

func _20210130204142_create_acme_servers_tableDownSql() (*asset, error) {
	bytes, err := _20210130204142_create_acme_servers_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210130204142_create_acme_servers_table.down.sql", size: 25, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210130204142_create_acme_servers_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8e\x4b\x6a\xc3\x40\x10\x44\xf7\x3a\x45\x2d\x25\xc8\x0d\x72\x98\xa1\x33\xdd\x49\x1a\xcf\x8f\x9e\x1e\x99\xb9\xbd\xc1\x02\x6d\x84\x0d\xde\x16\xef\x15\x2f\x9a\x90\x0b\x9c\x7e\x92\x80\x62\x96\xd0\xc5\x76\xb1\x8e\x75\x01\x00\x65\x8c\xa1\x8c\x52\x1d\x65\xa4\xf4\xf5\x5c\x59\x7b\x4b\x34\x43\xa1\x2c\xd8\xc9\xe2\x3f\xd9\x05\x31\x89\x5e\x6d\x86\x61\xe9\x05\xa3\xc5\xe5\xcf\xc8\xb5\x96\x77\x57\x47\x23\x07\x72\xb8\x66\xe9\x4e\xb9\x9d\x08\x58\x7e\x69\x24\x47\xa9\xf7\x75\x3b\x84\xd1\xf8\x33\x81\x25\xc9\x55\x38\x03\x9a\x69\x26\x9b\xb8\xc9\xc4\xaa\xbc\x2d\xdb\xf7\xf2\x08\x00\x00\xff\xff\xb9\xc3\x5e\x64\x39\x01\x00\x00")

func _20210130204142_create_acme_servers_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210130204142_create_acme_servers_tableUpSql,
		"20210130204142_create_acme_servers_table.up.sql",
	)
}

func _20210130204142_create_acme_servers_tableUpSql() (*asset, error) {
	bytes, err := _20210130204142_create_acme_servers_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210130204142_create_acme_servers_table.up.sql", size: 313, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210202110208_create_acme_accounts_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\x28\x49\x4c\xca\x49\x55\x48\x4c\xce\x4d\x8d\x4f\x4c\x4e\xce\x2f\xcd\x2b\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xe6\x3a\xd1\xce\x1a\x00\x00\x00")

func _20210202110208_create_acme_accounts_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210202110208_create_acme_accounts_tableDownSql,
		"20210202110208_create_acme_accounts_table.down.sql",
	)
}

func _20210202110208_create_acme_accounts_tableDownSql() (*asset, error) {
	bytes, err := _20210202110208_create_acme_accounts_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210202110208_create_acme_accounts_table.down.sql", size: 26, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20210202110208_create_acme_accounts_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x5b\x6a\xed\x30\x0c\x45\xff\x33\x0a\x7d\x26\x70\x67\x70\x07\x63\x74\xa4\x9d\x56\xd4\x8f\x20\xcb\xa7\x64\xf6\x85\xd3\x07\x85\x42\x93\x40\xfd\x6b\xaf\xe5\xbd\x91\xc4\xc1\x01\x0a\xbe\x65\x10\x4b\x41\x62\x91\x36\x6a\x74\x9a\x27\xfa\xf5\x98\xd2\x18\xa6\x54\x5b\x50\x1d\x39\xff\x3b\x78\xff\xb0\x77\xf8\x1d\x9e\x3e\x59\xc7\x0a\x47\x15\xf4\xef\xd7\xfd\x48\xa5\xd6\xb7\xcc\x7b\xaa\x5c\x40\x77\x76\x79\x66\x3f\x9d\x23\x2c\xf2\x75\x4a\xd1\xc5\x6d\x0b\x6b\xf5\xfa\x8f\xf0\xd2\x53\x5b\x1f\xf5\x4c\x90\xf8\xc9\x01\xa5\x5b\x6b\x19\x5c\x4f\x7b\xa4\xd5\x60\x89\x7e\x39\xc0\xc7\x4c\xd3\xf0\x7c\x99\x7d\xdf\x10\x4d\x1c\x14\x56\xd0\x83\xcb\xf6\x05\x93\x62\xe5\x91\x83\x6a\x7b\x9d\x97\x23\xd5\xd8\xf4\xaf\x54\x8a\x8c\x9f\xaa\x13\x75\x36\xb7\xc2\xbe\xd3\x0b\x76\x9a\x4d\x97\x69\xf9\x3f\xbd\x05\x00\x00\xff\xff\xb1\x30\x5e\x94\x05\x03\x00\x00")

func _20210202110208_create_acme_accounts_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__20210202110208_create_acme_accounts_tableUpSql,
		"20210202110208_create_acme_accounts_table.up.sql",
	)
}

func _20210202110208_create_acme_accounts_tableUpSql() (*asset, error) {
	bytes, err := _20210202110208_create_acme_accounts_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20210202110208_create_acme_accounts_table.up.sql", size: 773, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
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
	"20210130204142_create_acme_servers_table.down.sql":  _20210130204142_create_acme_servers_tableDownSql,
	"20210130204142_create_acme_servers_table.up.sql":    _20210130204142_create_acme_servers_tableUpSql,
	"20210202110208_create_acme_accounts_table.down.sql": _20210202110208_create_acme_accounts_tableDownSql,
	"20210202110208_create_acme_accounts_table.up.sql":   _20210202110208_create_acme_accounts_tableUpSql,
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
	"20210130204142_create_acme_servers_table.down.sql":  &bintree{_20210130204142_create_acme_servers_tableDownSql, map[string]*bintree{}},
	"20210130204142_create_acme_servers_table.up.sql":    &bintree{_20210130204142_create_acme_servers_tableUpSql, map[string]*bintree{}},
	"20210202110208_create_acme_accounts_table.down.sql": &bintree{_20210202110208_create_acme_accounts_tableDownSql, map[string]*bintree{}},
	"20210202110208_create_acme_accounts_table.up.sql":   &bintree{_20210202110208_create_acme_accounts_tableUpSql, map[string]*bintree{}},
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
