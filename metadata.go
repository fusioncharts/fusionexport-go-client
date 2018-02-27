// Code generated by go-bindata.
// sources:
// metadata/fusionexport-meta.json
// metadata/fusionexport-typings.json
// DO NOT EDIT!

package FusionExport

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

var _metadataFusionexportMetaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\xca\xcc\x2b\x28\x2d\x09\x0e\x73\x57\xb2\x52\x80\x88\x40\x44\x8b\x9d\x12\x8b\x53\xcd\x4c\x82\x52\x0b\x4b\x33\x8b\x52\x53\x94\xac\x14\x4a\x8a\x4a\x53\xc1\x0a\x6a\x75\x20\x3a\x93\x13\x73\x72\x92\x12\x93\xb3\xdd\x32\x73\x52\x03\x12\x4b\x32\x48\x37\x21\x25\xb1\x38\x23\x29\x3f\xb1\x28\xc5\x27\x3f\x3d\x9f\x74\xed\xf9\xa5\x25\x05\xa5\x25\x20\xeb\x5d\x52\xd3\x32\xf3\x32\x4b\x32\xf3\xf3\x48\x37\xa5\x28\xb5\x38\xbf\xb4\x28\x39\x95\x0c\x6f\x70\xd5\x02\x02\x00\x00\xff\xff\x68\x8b\x3b\x20\x45\x01\x00\x00")

func metadataFusionexportMetaJsonBytes() ([]byte, error) {
	return bindataRead(
		_metadataFusionexportMetaJson,
		"metadata/fusionexport-meta.json",
	)
}

func metadataFusionexportMetaJson() (*asset, error) {
	bytes, err := metadataFusionexportMetaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metadata/fusionexport-meta.json", size: 325, mode: os.FileMode(420), modTime: time.Unix(1519220983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _metadataFusionexportTypingsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x93\xc1\x4a\x03\x31\x10\x86\xef\x7d\x8a\x90\x73\x9f\xa0\x37\x5d\xad\x1e\x44\x84\x82\x82\xb7\xd9\xec\x74\x77\x30\x9b\x89\xb3\x13\xd9\x45\xfa\xee\xb2\x5d\xa9\x08\x2d\x24\x82\xd7\xf9\xf2\xe5\x4f\xc8\x9f\xcf\x95\x31\xc6\x58\xd7\x81\x68\xc5\x61\x4f\xad\xdd\x98\x65\x78\x04\x3a\x45\xb4\x1b\x63\x07\x15\x0a\xad\x3d\x82\xc3\x7a\x91\x28\xc4\xa4\xbb\xe7\xbb\x7c\x43\xb1\x8f\x1e\x14\xb7\xe4\xf1\x09\xb4\xcb\x37\x1d\x78\x5f\x83\x7b\x2b\x37\x61\x98\x82\xab\x20\x6a\x12\x3c\x6f\xd5\xcc\x1e\x21\xd8\xf5\x0f\x72\x1c\x3e\x50\x14\x65\xe6\xd7\x0b\xaf\x4e\xb3\x5f\xfb\xf7\x30\xbe\x00\xe9\x96\xe5\x3b\xe4\x76\x24\x3d\x1f\x44\x41\xb1\x45\xb9\x18\xf4\x98\xfa\x1a\xe5\x42\x4e\x03\x43\x57\x33\x48\xf3\xc0\x2d\xe7\x5f\xff\xa4\xdd\x23\x34\x33\x2d\x37\x77\xa9\xee\x4a\x65\x4e\x1a\x93\xce\xaf\x55\xd0\x8e\x85\x65\xae\x7e\x4f\xe0\x49\xa7\xbf\x1c\xe9\x06\xf7\x14\x48\x89\x43\xbe\x8d\x63\x64\xd1\xab\xe1\x95\xe2\x7f\xb4\x48\x70\xe0\x24\xae\xe4\x67\xac\x0e\x5f\x01\x00\x00\xff\xff\xfb\xbb\xd6\x9c\xbd\x03\x00\x00")

func metadataFusionexportTypingsJsonBytes() ([]byte, error) {
	return bindataRead(
		_metadataFusionexportTypingsJson,
		"metadata/fusionexport-typings.json",
	)
}

func metadataFusionexportTypingsJson() (*asset, error) {
	bytes, err := metadataFusionexportTypingsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metadata/fusionexport-typings.json", size: 957, mode: os.FileMode(420), modTime: time.Unix(1519220983, 0)}
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
	"metadata/fusionexport-meta.json":    metadataFusionexportMetaJson,
	"metadata/fusionexport-typings.json": metadataFusionexportTypingsJson,
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
	"metadata": &bintree{nil, map[string]*bintree{
		"fusionexport-meta.json":    &bintree{metadataFusionexportMetaJson, map[string]*bintree{}},
		"fusionexport-typings.json": &bintree{metadataFusionexportTypingsJson, map[string]*bintree{}},
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