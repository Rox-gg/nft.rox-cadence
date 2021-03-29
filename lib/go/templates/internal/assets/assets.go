// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/AdminAssignRox.cdc (773B)
// ../../../transactions/SetupUser.cdc (921B)
// ../../../transactions/TransferRox.cdc (792B)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _adminassignroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x6a\xe3\x30\x10\xc6\xef\x81\xbc\xc3\x90\xc3\x62\xc3\x62\xf6\xb0\xec\xc1\x6c\x12\x8c\x53\x43\x0e\x75\x43\xe2\x3e\x80\x2c\x4f\x1c\x51\x59\x12\xf2\xb8\x4d\x08\x79\xf7\xe2\xff\x6e\x9b\x42\x75\x90\x40\x7c\xf3\x69\xbe\x9f\x46\x14\x46\x5b\x82\x58\xab\xa8\x52\xb9\x48\x25\x26\xfa\x05\x15\x1c\xad\x2e\xe0\xcf\x39\x8e\x92\x60\xb3\xd9\x3f\x1c\x0e\xf3\x59\x27\xdd\xeb\xf3\x96\xb0\x28\x27\x92\xf0\x29\x4e\xf6\x41\x38\x4a\xe7\x33\xb2\x4c\x95\x8c\x93\xd0\xca\xb1\xc8\x85\x11\xa8\xc8\x87\x20\xcb\x2c\x96\xe5\x6f\xa0\x8b\xc1\xed\xc6\x87\xe7\xad\xa2\x7f\x7f\x5d\xb8\xce\x67\x00\x00\xed\x2e\x91\xa0\x10\x8a\xd0\xfa\xf0\xab\x7f\xcf\x8b\xa3\xe4\xb1\xb9\xac\xfd\x6b\x99\xb1\x68\x98\x45\x87\x65\x85\x50\x01\xe7\x3e\x04\x15\x9d\x02\xce\x75\xa5\xa8\xb1\x6c\x75\xf5\x2a\x51\x1e\xbd\xd6\x13\x96\xd0\x57\x78\xa9\xb6\x56\xbf\xfd\xbf\xf3\xc8\xca\xa9\xf3\xf9\x43\x5c\xaf\xbd\x3e\x90\xb6\x2c\xc7\x1d\xa3\x93\x3b\xba\xd7\x6b\xbd\x06\xc3\x94\xe0\xce\x22\xd4\x95\xcc\x40\x69\x82\xd6\x1e\x18\x58\x3c\xa2\x45\xc5\x11\x48\x03\x9d\x10\xe2\x28\xe9\x22\x2e\x3a\x9f\x5b\xdf\x2e\x9e\x91\x57\x84\x3d\x92\x1e\xc8\x40\x11\x96\x90\x23\x75\x31\x47\xb8\xee\x34\x6e\x57\x80\xe2\xb5\xc9\x3b\x88\x3e\xb6\xec\xe5\x48\x21\x33\x2c\x15\x52\xd0\xc5\x19\xa2\x86\x5a\x4a\x6c\xfe\x6e\x57\xa5\x52\xf0\x3b\x69\x07\x72\xd7\xcf\xa3\xf3\xa5\xfa\xb6\x72\x7e\x80\x2a\x9f\x76\x7c\x17\xd7\xe8\xbb\x70\xbf\xf9\xda\xe6\x88\xa3\x64\x3a\x72\xbd\xe9\x38\x73\xed\x39\x62\xbf\xbd\x07\x00\x00\xff\xff\x04\xc9\x05\xfa\x05\x03\x00\x00"

func adminassignroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_adminassignroxCdc,
		"AdminAssignRox.cdc",
	)
}

func adminassignroxCdc() (*asset, error) {
	bytes, err := adminassignroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "AdminAssignRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0x8d, 0x8c, 0xe0, 0x2d, 0x30, 0x8f, 0x8d, 0x61, 0x8b, 0x61, 0xc, 0x78, 0x91, 0x94, 0x80, 0xc3, 0xef, 0x78, 0xba, 0xe8, 0x14, 0x53, 0x2c, 0xd7, 0xd0, 0x79, 0xd1, 0x55, 0x36, 0x95, 0xf8}}
	return a, nil
}

var _setupuserCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x6e\x9b\x40\x10\xc6\xef\x48\xbc\xc3\xa7\x1c\x5a\x2c\x39\x49\xcf\x96\x13\xc9\x72\x12\xa9\x17\x37\xb2\xfd\x02\x0b\x0c\xb0\xca\xb2\x83\x86\x21\x89\x15\xf9\xdd\x2b\x30\x04\x1a\x5b\x6d\x6f\x96\xe7\x37\x7c\x7f\x76\x6c\x59\xb1\x28\x36\xec\x9f\x1a\x9f\xdb\xd8\xd1\x9e\x5f\xc8\x23\x13\x2e\xf1\xe3\x7d\xf3\xb4\x5f\x3d\x3c\x6c\x1f\x77\xbb\x30\xe8\xd1\x2d\xbf\xff\x54\x2a\xeb\x09\xb2\xfe\xb5\xd9\x6f\x57\xeb\x11\x0d\x03\x15\xe3\x6b\x93\xa8\x65\x8f\x8f\x30\x00\x80\x4a\xa8\x32\x42\x51\x6d\x73\x4f\xb2\xc0\xaa\xd1\x62\x95\x24\xdc\x78\x9d\x0d\x0c\x70\x7b\x0b\x9b\x41\x0b\x82\x39\xcd\x90\x32\xd5\xfe\xbb\xc2\x38\x21\x93\x1e\x50\x98\x57\x82\x41\xc2\xce\x51\x27\x30\xac\xda\x0c\xa7\x6f\xdf\xc4\x2c\xc2\x6f\xcb\x6f\x83\xd7\x9b\xf5\x27\xfc\x2c\xf6\xd5\x28\xdd\x47\xad\xfd\x05\x2e\x10\x3b\x65\x31\x39\x3d\x1b\x2d\x66\xb8\xbb\x83\xb7\x6e\xb4\x07\x38\xce\xa3\xab\x5d\xa7\x83\xc2\xd4\x48\xb8\xac\x1c\x29\xa5\xa8\x49\xd1\x54\x88\x29\x63\xa1\xab\xd9\xb8\x22\xa4\x8d\x7c\xda\x3c\x4e\xa2\x26\x42\x46\xdb\x34\x9e\xde\x40\x65\xa5\x87\x0b\xb9\x1c\xe9\xe4\x5f\x2c\xaf\x47\xd7\xa7\xfd\xc7\x76\x71\x0c\x10\x4d\xa4\x81\x89\x5a\xdd\x36\x67\x15\xca\xd3\x82\x07\xa0\xef\xae\x85\xa2\xe5\xf5\x28\x38\x87\xf2\x3f\x8b\x6a\xdf\xfc\x3c\x55\xd5\xc4\xce\x26\x48\x4c\x65\x62\xeb\xac\x1e\x90\xb1\x74\xda\xe7\x29\x7b\x79\x67\xfd\xcb\xdf\x1e\xee\xe3\xeb\xa9\x4e\x91\x4e\x6e\x3e\x7a\x1d\x7e\x7c\x45\x8e\xf7\xd1\x25\x85\x6e\xd6\xc6\x99\x43\x8d\xe4\xa4\xff\x11\xfb\xcf\x92\xbb\xe3\x58\x4f\x2f\xa2\xa9\xba\xcc\x7d\xd7\x8b\xf1\x2c\x5a\xb4\xcf\x6c\xd2\x54\xa8\xae\xfb\xd1\x31\x0c\x8e\x61\xf0\x3b\x00\x00\xff\xff\x7c\xc2\xc4\x36\x99\x03\x00\x00"

func setupuserCdcBytes() ([]byte, error) {
	return bindataRead(
		_setupuserCdc,
		"SetupUser.cdc",
	)
}

func setupuserCdc() (*asset, error) {
	bytes, err := setupuserCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "SetupUser.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe7, 0xfe, 0x7c, 0xd5, 0xab, 0xd2, 0x39, 0x34, 0xaf, 0xcf, 0x9a, 0xf2, 0x74, 0x68, 0x1b, 0xfc, 0xdd, 0x23, 0xce, 0x8b, 0x30, 0x8d, 0xe3, 0xb5, 0x6c, 0x8a, 0xf9, 0xa2, 0x83, 0x3b, 0xb8, 0x8d}}
	return a, nil
}

var _transferroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xcf\x8e\xda\x40\x0c\xc6\xef\x48\xbc\x83\x97\x43\x3b\x91\xda\x08\x55\x2d\x5d\x45\xfb\xa7\x88\x0a\x89\x0b\x5a\x6d\xe9\x03\x4c\x26\x4e\x18\x75\x18\x8f\x3c\x4e\xa1\x5a\xf1\xee\x55\x20\x90\x2c\xcb\xe6\x12\x1f\xbe\xcf\xfe\xfc\xf3\xd8\x4d\x20\x16\x58\x92\x9f\xd7\xbe\xb2\xb9\xc3\x15\xfd\x41\x0f\x25\xd3\x06\xc6\xbb\xf2\xb6\x98\xe0\xf8\xdb\xed\x24\x1f\xeb\x2f\x63\xf3\x7d\x38\x68\x0d\xcf\xb4\x5b\x08\x6e\xe2\xfb\xc2\xe1\x40\x58\xfb\xa8\x8d\x58\xf2\x8a\xd1\xd8\x60\xd1\x4b\x06\xd3\xa2\x60\x8c\xf1\x13\x6c\xad\xac\x0b\xd6\xdb\xc5\xcf\x0c\x7e\x2f\xbc\x4c\xbe\x26\xf0\xd2\x18\x01\x00\x1c\x0a\x08\x6e\xc2\x72\xbe\xca\xe0\xc7\x65\xbe\x74\x39\x5f\x9d\x94\x81\x31\x68\x46\x15\x6d\xe5\x91\x33\x98\xd6\xb2\x9e\x1a\x43\xb5\x97\x5e\xbf\x53\x4f\x43\xce\xe1\x21\xd3\x33\x96\x70\x0f\x47\x57\x9a\x13\x33\x6d\xef\x3e\x9c\xf6\x4a\x67\x67\xdd\x13\xdb\xbf\x5a\xf0\x41\x35\xab\x66\x70\x45\xf1\x4b\x88\x75\x85\x4f\x5a\xd6\x49\x37\xad\xf9\x1e\x1f\x21\x68\x6f\x8d\x1a\xcd\xa8\x76\x05\x78\x12\x38\x8e\x02\x0d\x8c\x25\x32\x7a\x83\x20\x04\xb2\x46\xa0\xad\x47\xfe\x18\x7b\x19\x47\x49\x3f\x7f\x44\x57\xa6\x2d\x14\xb8\xfb\xfc\x7a\x97\xf4\x84\x53\xf5\xb9\x76\x75\x9b\x6c\x7f\xfc\xe1\x0e\x4d\x2d\xf8\x96\xcf\xf9\x50\x70\x0f\x15\x4a\x0b\xb2\xbb\x5f\x72\x69\x28\x30\x50\xb4\x72\xa4\x79\x96\xa5\x15\xca\x4c\x07\x9d\x5b\x67\xe5\x9f\xba\x86\xb5\xce\x9d\x35\x07\x66\x37\x67\xfc\x2f\x6f\x0e\x7d\x69\xd8\x3f\xa8\xe4\xa6\x9f\xa1\x9b\x9f\xb6\xa5\x92\xc6\x99\x35\x84\xfa\xc4\x5e\x27\xa7\x4a\x8d\x1a\x8c\xab\xe6\x91\x96\xc8\x8c\xc5\xe8\x80\x68\x3f\x1c\xec\x87\x83\xff\x01\x00\x00\xff\xff\xd6\x48\x6b\x1b\x18\x03\x00\x00"

func transferroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transferroxCdc,
		"TransferRox.cdc",
	)
}

func transferroxCdc() (*asset, error) {
	bytes, err := transferroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "TransferRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8e, 0xce, 0x5b, 0x96, 0x8a, 0xde, 0x31, 0x31, 0x61, 0xa0, 0x21, 0xeb, 0x1f, 0x75, 0xc, 0x78, 0x62, 0xbb, 0x3d, 0x13, 0x48, 0x6e, 0xf1, 0x2d, 0x4e, 0x9a, 0x17, 0x4c, 0x1c, 0xf2, 0xb4, 0xba}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"AdminAssignRox.cdc": adminassignroxCdc,
	"SetupUser.cdc":      setupuserCdc,
	"TransferRox.cdc":    transferroxCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"AdminAssignRox.cdc": {adminassignroxCdc, map[string]*bintree{}},
	"SetupUser.cdc": {setupuserCdc, map[string]*bintree{}},
	"TransferRox.cdc": {transferroxCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
