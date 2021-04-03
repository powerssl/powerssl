// Code generated for package migration by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../db/migrations/20210130204142_create_acme_servers_table.down.sql
// ../../db/migrations/20210130204142_create_acme_servers_table.up.sql
// ../../db/migrations/20210202110208_create_acme_accounts_table.down.sql
// ../../db/migrations/20210202110208_create_acme_accounts_table.up.sql
package migration

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

var __20210130204142_create_acme_servers_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x4d\xae\xc2\x30\x0c\x84\xf7\x39\x85\x97\xad\xf4\x6e\xf0\x0e\x13\x99\x64\x80\x88\xfc\xc9\x71\x8a\x72\x7b\xa4\x16\x2a\x21\x51\x16\x78\x69\x7f\x9f\x35\xe3\x04\xac\x20\xe5\x53\x04\xb1\x4b\xb0\x0d\xb2\x40\x9a\x99\x0c\x11\x51\xf0\xf4\x3e\xbd\xbf\x56\xb9\x28\xe5\x1e\xe3\xdf\x0a\xfa\xd0\x6a\xe4\x61\x33\x27\xac\xe7\x85\xc5\x5d\x59\x3e\x80\x02\xa7\x45\x86\xed\x12\xbf\x81\x21\x2b\x2e\xc2\x1a\x4a\xde\xbe\x1e\x81\x5b\x07\x6f\x59\x9f\x19\x35\x24\x34\xe5\x54\x77\x90\x3c\xce\xdc\xa3\x52\x2e\xf7\x69\xde\xb4\x5e\xfd\x2f\x9a\x47\xc4\x91\xb6\x47\xaa\x12\x12\xcb\xa0\x1b\x06\x4d\xc1\xcf\x66\xfe\x37\x8f\x00\x00\x00\xff\xff\x75\x4a\xaf\x7b\x6b\x01\x00\x00")

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

	info := bindataFileInfo{name: "20210130204142_create_acme_servers_table.up.sql", size: 363, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
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

var __20210202110208_create_acme_accounts_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\x4d\x6a\xc4\x30\x0c\x46\xf7\x39\x85\x96\x09\xf4\x06\x3d\x8c\xd1\xc8\x5f\x5a\x53\xff\x04\x49\x9e\x92\xdb\x17\x92\x76\x48\x99\x29\x0d\x8c\xb7\x7e\x4f\x7a\xd8\xa2\x60\x07\x39\x5f\x32\x88\xa5\x20\xb0\x48\xeb\xd5\x6d\x18\x07\x22\xa2\x14\xe9\xe1\xe9\xfd\xe7\xa6\x36\xa7\xda\x73\x7e\xd9\xf8\x6d\x86\x41\xaf\xd0\x70\x74\x37\x5e\x31\x43\x51\x05\x76\xe4\x6c\x37\x63\xb2\x25\xf3\x1a\x2a\x17\x1c\x37\x5d\x59\xe5\x9d\xf5\x6e\x93\x27\xcf\x78\x50\xf6\x17\x1f\x61\xa2\x69\xf1\xd4\xea\x29\xde\xa1\xc5\x42\x9b\xb7\xca\x24\x08\xfc\xa6\x40\xa4\x4b\x6b\x19\x5c\xef\x78\x69\xd5\x59\xdc\xce\xf6\x7c\x3f\x74\xe8\x9a\x4f\xf1\xfb\x57\xc5\xc0\xfe\x7b\xbe\xa7\x02\x73\x2e\xcb\x8d\xa7\x88\x99\x7b\x76\xaa\xed\x73\x9c\x76\xbb\x2f\xf1\x09\x3b\x22\xe3\x1f\xfb\xd6\xb9\x68\x2a\xac\x2b\x7d\x60\xa5\x31\xc5\x69\x98\x5e\x87\xaf\x00\x00\x00\xff\xff\xe1\x88\xaa\xef\x67\x02\x00\x00")

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

	info := bindataFileInfo{name: "20210202110208_create_acme_accounts_table.up.sql", size: 615, mode: os.FileMode(420), modTime: time.Unix(726710400, 0)}
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
