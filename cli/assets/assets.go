// Code generated by go-bindata.
// sources:
// assets/banner.txt
// schema/mashling_schema-0.2.json
// DO NOT EDIT!

package assets

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

var _assetsBannerTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x95\xbd\x8e\xe4\x30\x08\xc7\xeb\xcd\x53\xd0\x31\x2b\x1f\xd0\xf3\x2c\x23\xd1\x5c\xef\xc2\xad\x1f\xfe\xf4\xf7\x47\x26\x9b\x71\x32\xb7\x24\x8a\x12\xec\x1f\x60\x0c\x0e\xed\xc2\xcc\xac\x6e\xc6\xdb\x4b\x65\x5d\xa4\xab\x14\xef\xee\xb8\xcd\x36\x22\x55\x15\x1f\x02\x25\xa6\xf9\x40\xa0\x70\x17\x31\xdd\x88\x45\xa4\x4f\x34\x53\x18\x32\xdf\x48\xfa\x34\x8c\x08\x0f\x65\x37\x37\x26\x51\x57\xd2\x47\x61\xde\xc4\x4d\x94\x99\xe9\x57\xa0\x9a\x63\xb1\xea\xc6\x47\xb5\xc9\xd1\xb8\x8f\xf0\x4e\x3e\x1d\x6b\x20\x62\x39\x04\x4b\xfc\x03\xc5\xa0\x2d\xe0\x81\xc2\xb1\xf3\x2b\x16\x79\xf3\x60\xd7\x28\xb1\x39\xb7\x2d\x63\x96\x33\xba\xb0\x86\xad\xd4\x6d\x8f\xcb\x05\x89\x17\x11\xf3\xb7\xf8\xd8\x91\x49\x55\x9d\x09\x95\xb9\xff\x67\x83\x3f\xd2\x86\xda\x49\x2f\xe9\x95\x62\x0b\x90\x52\x29\x3b\xa9\x9e\x52\x4e\x0b\x71\xe5\x15\xcb\x80\xc7\x12\x73\x32\x5b\xa1\xf0\x2e\x4b\x9a\x34\x97\xec\x88\x3d\xe7\x24\x9a\xfd\xc4\xbb\x58\xb2\xce\x2f\x71\xd2\x52\x92\x70\xca\x19\x73\x45\xce\x21\x6b\x32\x49\x77\x3c\x79\x29\xa9\xe1\xcb\x35\x63\x43\x46\xf6\xf4\xc2\x00\x59\xc9\xf9\xca\xc0\x31\xfd\x57\x11\x10\xe7\xf2\xd9\x80\xdd\x44\x80\x2c\x5c\x1b\xb0\x69\xc0\x2f\x79\xb2\x52\x96\x3c\x4e\x84\x69\xe5\x06\xe7\x95\x7f\x7b\x15\x83\xdd\xac\x7e\xec\x81\x5c\x56\x8e\x5d\xd5\xce\xee\x3b\x39\x3a\xfb\xf7\x28\xb6\x2e\x69\x6b\x15\x39\x44\xdb\x5a\xe5\x3a\xd9\xa3\xea\x66\x9f\x32\xab\xb8\x8d\x22\x45\x87\xde\x81\xd8\xea\x56\x55\xaa\xad\xa1\x75\x96\x97\xab\xe2\xf0\xb9\x63\xd1\xa6\x49\xfa\x7f\xa1\xb1\x36\x52\x8e\xde\x55\xbb\x67\xbd\x94\x7c\x3c\xbd\xfc\xef\x58\xa9\xb0\xde\xaf\xb5\x15\x57\x3a\x1c\x2c\x9c\xd4\x7a\x65\x58\xb2\x0f\x21\x53\x1e\x31\x1f\x44\xe6\x0f\x49\x94\x91\x88\x7b\x03\xd6\x9c\x63\xc9\xfd\x68\x45\xce\x54\x01\x8b\xa8\x2b\x7e\x56\x2a\x3c\xe5\xde\x16\x1b\xfc\x9d\xc5\x17\xba\x2e\x1f\x83\x9b\x7b\xd1\x2c\xc8\xca\xf8\xae\x96\x7e\x7c\xf3\xb6\x7d\x7d\x7d\x2d\x4d\xc5\x78\xc6\x59\x7f\x45\x54\xaa\xed\xf9\x88\xef\x1f\xfa\x47\x7c\x37\x24\x28\x70\xc3\x1e\x5e\xf0\x5a\xa9\x46\x00\xec\xdf\x73\xa4\x39\x8f\x68\x9e\x2a\x71\x10\x53\xd0\x93\x8c\x82\xf1\x88\xa6\x7b\xd6\xe6\x91\xf7\x91\xf6\x6d\x98\xd9\xb9\xd7\xf5\x88\x4a\xcf\xc0\xc4\x7a\xba\x1e\x2d\x84\x3a\xa2\xae\x0d\x9c\xaa\xa8\xcf\x88\x3f\x51\x23\xc2\xfa\x77\xbf\xe6\x08\x90\xf6\x16\x76\x99\xc4\xf7\xac\x86\x75\x37\xff\x35\xbb\xf9\xde\xfe\x05\x00\x00\xff\xff\x37\x0c\xc8\x3c\x83\x09\x00\x00")

func assetsBannerTxtBytes() ([]byte, error) {
	return bindataRead(
		_assetsBannerTxt,
		"assets/banner.txt",
	)
}

func assetsBannerTxt() (*asset, error) {
	bytes, err := assetsBannerTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/banner.txt", size: 2435, mode: os.FileMode(420), modTime: time.Unix(1508335574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaMashling_schema02Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4d\x6f\xdb\x3c\x0c\xbe\xe7\x57\x08\x6e\x8f\x6d\xf5\x1e\xde\x53\x8e\xdb\x69\xa7\x0e\xd8\x6e\x43\x10\x28\x36\x6d\x2b\xb3\x25\x4f\x52\x1a\x04\x45\xfe\xfb\x60\xf9\x23\x56\x22\xc9\xce\xa2\x04\x29\xd0\x1e\x7a\x20\x29\x92\x0f\xf3\x90\x94\xfc\x3e\x43\x08\xa1\xe8\x51\xc6\x39\x94\x24\x9a\xa3\x28\x57\xaa\x9a\x63\xbc\x96\x9c\x3d\x37\xd2\x17\x2e\x32\x9c\x08\x92\xaa\xe7\xff\xfe\xc7\x8d\xec\x21\x7a\x6a\x4e\xaa\x5d\x05\xf5\x31\xbe\x5a\x43\xac\x3a\xa9\x80\x3f\x1b\x2a\x20\x89\xe6\xe8\x97\x96\x68\x69\x49\x64\x5e\x50\x96\x2d\xdb\x68\x4f\x07\x55\x46\x14\x6c\xc9\x2e\xd2\x92\x45\xeb\xa5\x12\xbc\x02\xa1\x28\xc8\x68\x8e\xde\xdd\x7e\x86\x4a\x23\x29\xa9\x04\x65\x59\xd4\x2b\xf7\x96\x88\xce\xc3\x06\xa2\x5e\x6b\x45\xd6\x6b\x19\x29\xe1\xe8\x84\x96\xbf\x81\x90\x94\x33\x9b\x2a\xe6\x2c\xa5\xd9\x46\x10\x45\x39\x93\x36\x0b\x25\x68\x96\x81\xb0\xea\xe0\x0d\x98\x5a\xe6\x84\x25\x85\xd7\xa2\xa0\xec\xb7\x8c\x0c\xed\xe2\x08\x99\xa3\xda\x26\x36\x9b\x06\xf9\x4a\xde\xfd\xed\x3d\x55\x09\xea\x34\xa1\xb2\x2a\xc8\x6e\x19\x3e\xdd\xce\x33\x2d\x49\x16\xda\x35\xc8\x58\xd0\x4a\x05\xaf\xc6\x11\xbb\x9c\xbe\xa9\x82\xd2\xad\xd6\x26\x8f\x02\xd2\x3a\xfc\x03\x4e\x20\xa5\x8c\x6a\x8f\xd8\x08\x70\x9a\x93\x23\x2f\x03\x0f\x11\xa2\xeb\xfc\x31\x38\x47\x74\x0f\x0f\xc7\x08\x70\x23\x38\x4d\x6f\x5e\x0b\x4b\xed\xfd\xea\x40\xfa\x11\x15\x1e\x45\xeb\x3a\x2c\x84\x99\xc7\x49\x44\x92\x44\xc7\x26\xc5\xf7\xe1\x48\x4c\x49\x21\x61\x66\xba\x68\x8f\x46\x83\x84\xcd\x55\x65\xb6\xc7\x4d\x76\x8d\xf6\x69\x91\x4b\x50\x8a\xb2\xec\xee\xb6\x40\x7b\xe6\x23\x4c\xd3\xbe\x84\xa3\x5e\xad\x3f\x68\x6f\xe5\x60\x98\xa7\x37\xc6\xb2\x45\x27\xac\x3e\x95\xfc\x1b\xcf\x07\x57\xa6\xae\x15\x3f\x69\xfc\x49\xe3\x0f\x4c\x63\x73\xc7\x87\x27\xb3\x97\x98\x9c\xc1\x6b\x6a\x3d\xec\xf8\x29\xbc\x01\x07\x56\x29\x08\x60\x31\xd8\x6b\xba\x98\xc2\x8b\x8b\x12\x38\x6c\xc0\xc9\x19\xdc\x59\x03\x5f\xad\xdd\x2a\x22\x88\xe7\xee\x73\x6f\xcd\x66\x41\x70\x60\x57\xe0\x8a\xf7\x9c\xb9\xb0\x38\x13\x39\x4a\x14\x71\x78\xd0\x7a\x9a\xf8\xb4\xf5\xed\xd4\x4e\x6d\x47\x52\x23\x14\x36\xd3\xf2\x59\x20\x4b\x21\x9c\xc6\x8e\x2b\x31\x6a\x01\x4e\x8e\xe3\x63\xcd\x58\x9c\xe6\x22\x7f\x5e\x20\xb7\x3b\xd4\x72\x85\x6c\x0a\xd5\x7d\xa1\x92\x73\x8c\x33\xaa\xf2\xcd\xea\x25\xe6\x25\xfe\xf9\xed\xcb\xd7\xd7\x1f\x3c\x55\x5b\x22\x00\xa7\x05\xcf\xf8\x73\xcc\x99\x12\x74\x85\x57\x05\x5f\xe1\x92\x48\x05\x02\x93\xb8\x66\x5b\x6d\xb0\x1d\x3c\x32\xda\x6f\x5a\x2f\x6b\xe9\x9a\x5e\xc8\xda\x2a\x9e\x32\x4c\xdb\x4a\x76\xd7\x81\x16\x9c\x7e\xf8\x05\xdd\x6e\xbe\x0f\x51\x09\x95\x15\x51\x71\x0e\x17\x5e\xcc\xc6\x5f\x92\xe6\xdb\xce\x51\xfe\x09\xcf\xcd\x70\xe3\x71\x80\xfd\x66\x59\x7b\xe7\x21\x9a\x3c\x13\xb5\xa5\xf7\x73\x07\x72\x4f\x38\x74\xc6\x94\xd3\xb6\x74\x7c\x2c\xa0\x73\x66\x10\xf2\xcf\x21\x34\x04\x17\x3e\xf0\xb9\x93\xe2\x2a\x6d\x3f\x6b\xfe\xef\x67\x7f\x03\x00\x00\xff\xff\x39\x9c\x5e\x55\xca\x17\x00\x00")

func schemaMashling_schema02JsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaMashling_schema02Json,
		"schema/mashling_schema-0.2.json",
	)
}

func schemaMashling_schema02Json() (*asset, error) {
	bytes, err := schemaMashling_schema02JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/mashling_schema-0.2.json", size: 6090, mode: os.FileMode(420), modTime: time.Unix(1507736443, 0)}
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
	"assets/banner.txt":               assetsBannerTxt,
	"schema/mashling_schema-0.2.json": schemaMashling_schema02Json,
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
	"assets": &bintree{nil, map[string]*bintree{
		"banner.txt": &bintree{assetsBannerTxt, map[string]*bintree{}},
	}},
	"schema": &bintree{nil, map[string]*bintree{
		"mashling_schema-0.2.json": &bintree{schemaMashling_schema02Json, map[string]*bintree{}},
	}},
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
