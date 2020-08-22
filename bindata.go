// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package main generated by go-bindata.// sources:
// README.md
// templates/index.html
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\x5f\x6f\xdb\x36\x14\xc5\xdf\xf9\x29\x0e\x9c\x3e\x6c\x40\x69\xd9\x8e\x17\xc7\x01\x0a\x2c\xed\xba\xad\x1b\xb6\x06\x6d\x3a\x6c\x6f\xa2\xc8\x2b\x8b\xb0\x44\x0a\x24\x15\xd7\x98\xbb\xcf\x3e\x90\x52\xec\x65\x73\xff\x25\x28\x30\x3d\x04\x97\xd6\xe5\xbd\x87\xbf\x73\x6d\xe6\x04\x57\xba\xa5\x5a\x1b\xba\xc0\x65\xa8\x85\x01\xc7\x53\x21\xd7\x64\x14\x9e\x9b\x95\x36\x44\x4e\x9b\x15\x5e\x98\x40\x2e\xbe\xbc\x16\x7e\xcd\xd8\xf3\xb7\x2d\x39\xdd\x90\x09\xf1\xe5\x46\x87\x0a\xad\xe8\x3c\x89\xa2\xa6\xc7\x70\xe4\xbb\x26\x86\x08\xc2\xaf\x3d\xb4\x81\xc0\x86\x0a\x78\x72\x37\x5a\xd2\x98\xb1\x93\x13\xbc\xf1\x62\x45\x31\x8a\x61\x2c\xf3\x9d\x95\x6b\x72\x8c\xbd\xea\x0c\x72\x95\x16\x70\x9d\x01\xd7\x01\xbc\xc5\xf9\xe4\x7c\x72\x11\xff\xa0\x75\x8d\x77\x7e\x13\xb2\x76\xd0\x9e\x8f\x71\x5d\x11\x2e\xaf\x5e\x60\xa3\xeb\x1a\x05\x41\x48\x49\xde\xeb\x28\xc2\x1a\xe4\x55\x08\xed\x45\x96\xd5\x56\x8a\xba\xb2\x3e\xa4\x42\x79\x12\x72\x72\x82\xa7\x9d\xae\x55\x94\xa0\x1b\xb1\x22\x6c\x6d\xe7\x3c\xd5\x25\x63\xbc\x7f\x85\x50\xd1\xf0\xae\x4b\x52\xe3\xba\x75\xf6\x46\x2b\x52\x83\xee\x52\xd7\xf1\x60\x40\x9e\xe7\x0c\x18\xf4\x17\x69\x3b\x0f\xb8\x95\x8a\xf1\x90\xc2\x38\x7e\xb5\x9b\xd8\x0b\x52\x98\x74\x50\x1d\x86\xf2\x7d\xc5\xff\x56\x3b\x4e\x63\xa8\x7c\xa8\xfb\xc7\x50\xd3\xd8\xcd\xc0\x21\x09\x8e\x78\x3e\xc6\xe2\x80\xa2\x74\xb6\x81\xb7\x9d\x93\x14\x6b\xfe\xd4\xf9\x90\xfa\xe7\x2b\x8b\x15\x05\xac\x74\xa8\xba\x62\x2c\x6d\x93\x1d\xf1\x23\x6e\x89\x96\x14\xda\x08\xb7\xed\x5d\x89\x72\xa2\x33\x37\x42\xd7\x69\x3a\xb4\xf1\x5a\xf5\xb8\x91\x3f\xfa\xe1\xe5\xd5\xe5\xf5\x8f\x59\xa1\x4d\x8e\xaf\xf2\xbf\xb2\x95\xed\x63\x6d\xd0\x58\x1f\x20\x85\x27\xff\xf5\x78\x7f\x3a\xaf\x9b\xb6\xde\xde\x05\xb7\xdf\x76\x47\x4a\x3c\xd7\xcf\x5d\x41\xce\x50\x20\xcf\xd8\x6d\x85\x52\x1b\x05\x7a\x2b\x9a\xb6\x26\x34\xc2\xe8\x92\x7c\x48\xe3\x1a\x71\xe5\xfb\x4f\xb2\x1c\x4a\x3b\x92\xc1\xba\xed\x18\xbf\x58\xa5\xcb\x6d\x4c\x69\x22\x5d\xeb\x12\xae\x60\xfb\x73\x18\x22\xe5\x21\x8c\x82\xa2\xb6\xb6\xdb\x5b\x61\xeb\xae\x20\x19\x6a\x48\x47\x22\x10\x78\x89\x71\xb6\x6f\xd0\x8b\x4c\x06\x39\x2a\xc9\x91\x91\x34\x8c\x66\x9e\x75\x6d\x6d\x85\xca\xc1\xf1\x26\x45\x78\xf6\xfa\x37\xc4\x69\x63\x6c\x07\x6d\xda\x2e\x00\x3b\x28\xf2\xd2\xe9\x36\x68\x6b\xf0\xbe\x67\xc7\x76\xe0\xe9\xc1\x6d\xf0\xc1\x27\x6d\xc8\x63\xab\x1c\x3b\x5c\xee\x1b\x63\x53\x69\x59\xf5\xa6\xc6\x51\x68\x9d\x8d\x43\x46\x0a\x3b\xc6\xf2\x3c\x2f\x84\xaf\xd8\x23\xc8\xce\xd5\xe0\xbf\xe3\xea\xe5\xeb\x6b\xf0\xef\x31\x8a\x7b\x9f\x7c\xdb\x8a\x50\x65\xc1\x66\x81\x7c\x18\x4b\x7f\x33\xc2\xd1\x81\x1c\xce\xcd\xd8\x9f\x0c\x18\xf9\x20\x42\xe7\x47\x17\x18\xf9\x2e\x4d\xf4\xe8\x71\xfc\x58\x89\x20\x46\x17\x88\x29\xc0\x48\xab\x98\x50\xd0\xf2\xf4\x6c\x21\x4f\xb9\x9c\x2f\x67\x7c\x2e\x69\xc1\xc5\xec\x9b\x33\x2e\xcb\x79\x39\x9b\x0a\xb1\x28\x4e\xe7\x23\x06\xbc\x63\xef\x58\xfa\xc2\x0c\x9c\xfb\x16\x91\xf3\xb3\x8a\xe4\x1a\xfd\x1a\xb6\x84\x48\xbf\x63\x07\xdc\x9f\x46\xfb\xdf\xd0\x3f\x8d\xf9\x1d\xf4\x5a\xe5\xd1\xdb\xf8\x45\x8a\x0a\xa0\x55\x94\x13\x6e\x97\xf1\xb7\x63\x23\x4c\xf8\x87\xd4\x0f\x1b\xa0\xd5\x93\x99\x5c\x9c\xd3\xe2\x6c\xc2\xa7\x72\xa2\xf8\x7c\x3a\x27\xbe\x5c\x8a\x39\x3f\x2d\xc4\x6c\x51\x2c\xce\xe4\xa4\x9c\xbc\xcf\x91\xbe\xcd\xe7\x38\x72\x48\x72\x9d\x31\xda\xac\x8e\x81\x4f\x37\x47\xe4\x7e\x15\x03\x08\x0c\xb9\x0f\xa0\x7e\x2f\xe8\x9f\xc7\x3c\xd8\xfe\xca\xfb\x38\x72\x52\x85\x98\x4e\xcf\x0b\x3e\x39\x55\x05\x9f\x17\x45\xc9\xc5\x72\x2e\xf9\x62\x52\x4e\x97\xcb\x59\x59\xce\xcb\xe9\xfb\x90\xa7\x16\x9f\x43\xbc\x21\x1f\xef\xd5\x98\x95\xf4\xa6\x02\xea\x18\xf6\x74\x4b\x27\xee\xaf\x52\x04\x31\x24\xff\xdf\xa7\x3d\xd8\xfe\x3f\x8c\x2f\x8c\xbe\xef\xf1\x10\xf6\x7d\x85\xa3\xf0\x03\xb9\x46\x1b\x11\x12\xff\xfd\xe2\x30\xfb\xd9\xc3\xad\xb8\xb7\x1b\xf7\x31\xe4\x70\x84\x2f\xea\xc9\xbe\xcd\x43\x6c\xd9\x17\xb9\xe3\xcc\xdf\x01\x00\x00\xff\xff\x78\x02\xbd\xf3\x0a\x0b\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 2826, mode: os.FileMode(420), modTime: time.Unix(1598118098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x91\x31\x73\xe3\x20\x10\x85\x6b\xe9\x57\xec\x51\x9f\x84\xdd\x79\x64\x50\x71\x97\xa4\x8d\x0b\xa7\x48\x89\xc4\x46\xda\x31\x02\x0d\x60\xd9\x8e\xc7\xff\x3d\x83\xc8\x24\x15\xb3\xdf\x5b\x1e\x6f\x59\xf1\xe7\xe9\xf5\xff\xf1\xfd\xf0\x0c\x63\x9c\x4c\x5b\x8a\x74\x80\x51\x76\x90\x0c\x2d\x4b\x00\x95\x6e\x4b\x00\x31\x61\x54\xd0\x8f\xca\x07\x8c\x92\xbd\x1d\x5f\xaa\x1d\xfb\x15\xac\x9a\x50\xb2\x85\xf0\x32\x3b\x1f\x19\xf4\xce\x46\xb4\x51\xb2\x0b\xe9\x38\x4a\x8d\x0b\xf5\x58\xad\xc5\x5f\x20\x4b\x91\x94\xa9\x42\xaf\x0c\xca\x6d\xbd\xc9\x46\x91\xa2\xc1\xf6\x40\x33\x1a\xb2\x28\x78\xae\x93\x62\xc8\x9e\xc0\xa3\x91\x2c\xc4\x9b\xc1\x30\x22\x46\x06\xa3\xc7\x0f\xc9\xc6\x18\xe7\xd0\x70\x7e\xb6\xf3\x69\xa8\x7b\x37\xf1\x69\x99\xeb\x3e\x04\xd6\x96\x85\x58\xfb\xdb\xb2\x28\xea\x94\x48\x91\x45\x0f\xf7\xb2\x28\x8a\x49\x5d\x73\x9c\x06\x76\x9b\xcd\x7c\xdd\x27\xd8\xb9\x6b\x15\xe8\x93\xec\xd0\x40\xe7\xbc\x46\x5f\x75\x2e\x4b\xb3\xd2\x7a\xe5\x5b\x9c\xf6\xd9\xc0\x0f\x64\x1b\x50\xe7\xe8\x12\x78\x94\x85\xe0\xdf\xcf\x09\x9e\xbf\x4d\x74\x4e\xdf\xd6\x09\x34\x2d\xd0\x1b\x15\x82\x64\x3f\x41\xd6\xa9\x01\xee\x77\xa8\xff\x39\x7d\x83\xc7\x23\x75\x72\x4d\x4b\x72\xc8\x57\x05\xcf\x8b\xf9\x0a\x00\x00\xff\xff\x0a\x8c\x46\xff\xa9\x01\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 425, mode: os.FileMode(420), modTime: time.Unix(1598123227, 0)}
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
	"README.md":            readmeMd,
	"templates/index.html": templatesIndexHtml,
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
	"README.md": &bintree{readmeMd, map[string]*bintree{}},
	"templates": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
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
