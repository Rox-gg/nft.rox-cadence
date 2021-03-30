// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../templates/scripts/nft_collection_length.cdc (405B)
// ../../../templates/scripts/read_nft_id.cdc (631B)
// ../../../templates/scripts/rox_items_total_supply.cdc (150B)
// ../../../templates/transactions/AdminAssignRox.cdc (865B)
// ../../../templates/transactions/DestroyRox.cdc (386B)
// ../../../templates/transactions/SetupUser.cdc (921B)
// ../../../templates/transactions/TransferRox.cdc (788B)

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

var _scriptsNft_collection_lengthCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcd\x6e\xab\x30\x10\x85\xd7\x20\xf1\x0e\x73\xb3\xb8\x82\x0d\xe9\x3a\x6a\x1a\x21\x68\x24\x36\xb4\x22\xbc\x80\x31\x86\x5a\x35\x33\xc8\x8c\xd5\x54\x51\xde\xbd\x22\xe4\x87\x76\xe7\xc5\xf7\xf9\x9c\x33\xba\x1f\xc8\x32\x14\x84\x7b\x87\x9d\xae\x8d\xaa\xe8\x53\x21\xb4\x96\x7a\x78\x3a\x16\xfb\x2a\xc9\xb2\xf2\xf5\x70\x08\xfc\x2b\x5a\xd2\x31\x67\xd5\x8f\x0b\x24\x7d\x2b\xaa\x32\x49\x1f\x68\xe0\x0f\xae\x86\xd6\x21\xf4\x42\x63\x28\xa4\x4c\x9a\xc6\xaa\x71\xdc\xc0\xf5\x11\x6d\x20\x47\x86\xd3\xc4\x02\x00\x18\xc5\x20\xa4\x64\xd8\x42\xa7\x38\x91\x92\x1c\xf2\x42\x8c\x02\xdf\xf3\xbc\x89\x92\x64\x8c\x92\xac\x09\x4b\xd5\xc2\xf6\x62\xc5\x9d\xe2\x54\x0c\xa2\xd6\x46\xf3\x77\xb8\x1e\x5c\x6d\xb4\x5c\xdf\xaa\xa6\x77\x25\xfa\x17\xd7\x64\x2d\x7d\x3d\xff\x3f\xfd\xdd\x1c\x3f\xb0\xf7\x8b\x7f\x7e\x09\xe7\x58\x6f\xb7\x83\x41\xa0\x96\xe1\x2a\x25\x67\x1a\x40\x62\x98\xff\x01\x79\x8f\x9d\x0f\x32\x47\x2f\x5a\xae\xa2\xdb\x46\xab\xd8\x59\xfc\x3d\x60\x6a\x9e\x67\x63\x18\xc5\x46\x61\xc7\x1f\x81\x7f\x0e\xfc\x9f\x00\x00\x00\xff\xff\x87\x32\x68\xf6\x95\x01\x00\x00"

func scriptsNft_collection_lengthCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsNft_collection_lengthCdc,
		"scripts/nft_collection_length.cdc",
	)
}

func scriptsNft_collection_lengthCdc() (*asset, error) {
	bytes, err := scriptsNft_collection_lengthCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/nft_collection_length.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xea, 0xfd, 0xa, 0x3c, 0x64, 0x93, 0xdd, 0xcc, 0x58, 0x95, 0xbf, 0x57, 0x11, 0xa7, 0x2f, 0xae, 0x8d, 0x1c, 0x4c, 0xb, 0xf4, 0x58, 0x82, 0x4d, 0xd3, 0xf0, 0x41, 0x74, 0x48, 0xfb, 0x67, 0x3c}}
	return a, nil
}

var _scriptsRead_nft_idCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xc1\x6e\xd4\x40\x10\x44\xef\x2b\xed\x3f\x94\x72\x00\xaf\x84\xd6\x1c\x10\x87\x88\x10\x2d\x0e\x46\xb9\x18\xe4\x98\x0f\x18\xcf\xb4\x37\x2d\xec\x1e\x6b\xa6\x47\x09\x8a\xf2\xef\xc8\x71\xbc\x6b\xc0\x27\x5b\xae\xaa\xee\x7e\xc5\xc3\xe8\x83\xa2\xf2\x52\x26\x39\x72\xdb\x53\xe3\x7f\x91\xa0\x0b\x7e\xc0\xfb\xc7\xaa\x6c\x0e\x37\x37\xf5\xd7\xbb\xbb\xed\xe6\x55\x5a\xfb\xc7\x5b\xa5\x21\xae\x24\xc5\xf7\xaa\xa9\x0f\xc5\x59\xba\xdd\xe4\x39\x9a\x7b\x8e\x88\x36\xf0\xa8\x08\x64\x5c\xc4\x40\x6a\x9c\x51\x03\xd3\xfa\xa4\x30\x82\xaa\x6c\xc0\x02\x83\x14\x29\xbc\x8d\xb0\xbe\xef\xc9\x2a\x7b\xd9\x6e\xc6\xd4\xa2\x4b\x82\xc1\xb0\x64\xc6\x5a\x9f\x44\x2f\x71\x70\x2e\x50\x8c\xef\xc0\xee\x12\x3f\x6f\x45\x3f\x7e\xd8\x2d\x2f\x78\x9a\x46\x03\x40\x9e\xe3\x1b\x29\xf4\x9e\x30\xa6\xb6\x67\xbb\x4a\x86\xef\x5e\x7e\xf8\x07\xa1\xb0\x7c\xe8\x74\xf5\xec\xed\x49\x57\xea\x9a\x3a\x5c\xe1\x48\x7a\x98\x37\x58\x36\xd9\xcd\xe2\xe9\xd9\x1f\x49\x0b\x33\x9a\x96\x7b\xd6\xdf\x59\x3e\x4f\xcc\x17\x50\xc5\x29\x6b\xed\x69\x7d\x08\xfe\xe1\xd3\x9b\xa7\x7f\xd1\xef\xcf\xfa\x1f\x2f\x41\xcf\x9f\xb3\x95\xf1\xfa\x1a\xa3\x11\xb6\xd9\x45\xe1\x53\xef\x20\x5e\x31\x67\xc1\x9e\x76\x98\xbb\xf9\xef\xf2\x8b\xdd\x8a\xcf\x97\xd9\x64\x10\xa8\xa3\x40\x62\x27\x08\x30\x88\x23\x59\xee\xd8\x2e\xdd\x4c\x74\xd6\xb5\x2c\x88\xa4\x53\x5c\xfd\x0d\xea\xf5\xa8\xaa\x6c\xb2\xa9\x1d\x76\xa7\x79\x81\x34\x05\x99\x3c\x7b\x76\xdb\xcd\xf3\x9f\x00\x00\x00\xff\xff\xe9\x15\x71\x5a\x77\x02\x00\x00"

func scriptsRead_nft_idCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsRead_nft_idCdc,
		"scripts/read_nft_id.cdc",
	)
}

func scriptsRead_nft_idCdc() (*asset, error) {
	bytes, err := scriptsRead_nft_idCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/read_nft_id.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd8, 0xae, 0x6, 0xd2, 0xc4, 0x6, 0xb0, 0xc9, 0xd8, 0x34, 0x36, 0x3c, 0x8, 0xd1, 0xee, 0x26, 0xf4, 0x60, 0xd5, 0x47, 0x81, 0xde, 0x1, 0x12, 0x5e, 0x79, 0x45, 0x10, 0xf0, 0x77, 0x21, 0x4c}}
	return a, nil
}

var _scriptsRox_items_total_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcc\x2d\xc8\x2f\x2a\x51\xf0\xcb\xcf\x73\x2b\xcd\x4b\xcf\x4c\xca\x49\x0d\xc9\xcf\x4e\xcd\x53\x48\x2b\xca\xcf\x55\x30\xa8\xf0\x73\x0b\x71\x74\x71\x09\x72\x0d\x0e\xe6\xe5\x82\x2a\x0d\xca\xaf\xf0\x2c\x49\xcd\x2d\x46\x52\xe2\xec\xef\x17\x12\xe4\xe8\x8c\x50\xca\xcb\x55\x50\x9a\xa4\x90\x56\x9a\xa7\x90\x9b\x98\x99\xa7\xa1\x69\xa5\x10\xea\x99\x57\x62\x66\xa2\x50\xcd\xcb\xa5\xa0\xa0\xa0\x50\x94\x5a\x52\x5a\x94\x07\x37\x4a\xaf\x24\xbf\x24\x31\x27\xb8\xb4\xa0\x20\xa7\x92\x97\xab\x96\x97\x0b\x10\x00\x00\xff\xff\x18\x2a\x3b\x3e\x96\x00\x00\x00"

func scriptsRox_items_total_supplyCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsRox_items_total_supplyCdc,
		"scripts/rox_items_total_supply.cdc",
	)
}

func scriptsRox_items_total_supplyCdc() (*asset, error) {
	bytes, err := scriptsRox_items_total_supplyCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/rox_items_total_supply.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x29, 0x34, 0x35, 0xf7, 0x82, 0xbe, 0xce, 0x53, 0xc7, 0xca, 0x71, 0x84, 0x81, 0xd0, 0x86, 0xed, 0xfc, 0xbf, 0x20, 0xc5, 0x2b, 0x55, 0xa1, 0x5c, 0xbf, 0x38, 0xb3, 0xdc, 0x74, 0xd1, 0xa1, 0xaf}}
	return a, nil
}

var _transactionsAdminassignroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xcd\xca\xdb\x30\x10\xbc\x07\xf2\x0e\x4b\x0e\xc5\x86\x60\x7a\x28\x3d\x98\x26\xc1\x38\x35\xe4\x50\x37\xc4\xee\x03\xc8\xf2\xc6\x11\x95\x25\x23\xaf\xdb\x94\x90\x77\x2f\xf2\x7f\xbe\xe4\x83\x4f\x07\xd9\x12\xb3\xa3\x9d\xd9\x11\x65\xa5\x0d\x41\xac\x55\xd4\xa8\x42\x64\x12\x53\xfd\x1b\x15\x9c\x8d\x2e\xe1\xf3\x35\x8e\xd2\x60\xbf\x3f\x7d\x4f\x92\xe5\xa2\x87\x9e\xf4\xf5\x40\x58\xd6\x33\x48\xf8\x33\x4e\x4f\x41\x38\x41\x97\x0b\x32\x4c\xd5\x8c\x93\xd0\xca\x31\xc8\x45\x25\x50\x91\x0f\x41\x9e\x1b\xac\xeb\x35\x70\x2d\x25\x72\xb2\x0f\x1e\x72\x1f\x12\x32\x42\x15\x6b\x20\x81\x66\x3a\x95\x42\x51\xdc\x94\x19\x1a\xf0\xe1\xd7\x41\xd1\xd7\x2f\x2e\xdc\x96\x0b\x00\x80\x6e\x97\x48\x2d\xca\x56\x7d\x1a\x3a\xf3\xe2\x28\xfd\xd1\x5e\xda\x4e\x2c\xac\x32\x58\x31\x83\x0e\xcb\x4b\xa1\x02\xce\x7d\x08\x1a\xba\x04\x9c\xeb\x46\x51\x4b\xd9\xe1\xec\xaa\x51\x9e\xbd\x8e\x13\x36\x30\x54\x78\x99\x36\x46\xff\xfd\xf6\xe2\x91\xad\x63\x9d\xf0\x47\x63\xbc\xee\x3a\x21\x6d\x58\x81\x47\x46\x17\x77\x62\xb7\x6b\xb7\x83\x8a\x29\xc1\x9d\x55\xa8\x1b\x99\x83\xd2\x04\x1d\x3d\x30\x30\x78\x46\x83\x8a\x23\x90\x06\xba\x20\xc4\x51\xda\x4b\x5c\xf5\x3c\xf7\xa1\x5d\xbc\x22\x6f\x08\x07\x4b\x06\x43\x46\xbf\x61\x03\x05\x52\x2f\x73\x1a\x83\x3b\x97\xdb\x17\xa0\xf8\xd3\xea\x1d\x41\x8f\x2d\x7b\x05\x52\xc8\x2a\x96\x09\x29\xe8\x9f\x33\x4a\x0d\xfb\x31\x6a\x75\x6c\x32\x29\xf8\x0b\xb5\xa3\x73\xb7\xb7\x21\x7b\xaa\xbe\x6f\x9d\x0f\x58\x55\xcc\x3b\x7e\x69\xd7\xc4\xbb\x72\xdf\x19\x6d\xfb\x89\xa3\x74\x1e\xce\x81\xf4\x29\x9d\x0f\xc7\x21\xa4\x76\x9f\x47\xd4\x9f\xfd\x4f\x83\xba\xff\x0f\x00\x00\xff\xff\xd7\x6e\x8c\xcb\x61\x03\x00\x00"

func transactionsAdminassignroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsAdminassignroxCdc,
		"transactions/AdminAssignRox.cdc",
	)
}

func transactionsAdminassignroxCdc() (*asset, error) {
	bytes, err := transactionsAdminassignroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/AdminAssignRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x19, 0x2e, 0x3f, 0x22, 0x48, 0xf5, 0x40, 0xfb, 0xa7, 0xd, 0xbe, 0xa0, 0x4, 0x3f, 0x36, 0xe5, 0x24, 0x89, 0xfb, 0x3d, 0x9, 0x2a, 0x7e, 0x88, 0xca, 0x93, 0x68, 0xd7, 0x4a, 0xaa, 0xde, 0xf3}}
	return a, nil
}

var _transactionsDestroyroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xc1\x6a\x83\x40\x10\x86\xcf\x2e\xec\x3b\x4c\x6f\x06\x1a\xe9\xa1\xf4\x20\x52\x2a\xda\x80\x17\x1b\xd4\x3e\xc0\x76\x9d\xc4\xa5\xba\x23\xeb\x98\xa4\x94\xbc\x7b\xd1\x34\xb5\x87\x42\x6f\xcb\xce\xf7\xff\xff\x67\xba\x9e\x1c\x43\x4e\x76\x33\xda\xbd\x79\x6b\xb1\xa2\x77\xb4\xb0\x73\xd4\xc1\xdd\x29\xdf\x54\x71\x9a\x16\xcf\x65\x29\xc5\x37\x5a\xd0\x29\x63\xec\x86\x5f\x48\xf2\x92\x57\x45\x9c\x2c\xa8\x14\xec\x94\x1d\x94\x66\x43\xd6\x37\x75\x08\xaf\x99\xe5\x87\xfb\x15\x7c\x4a\xe1\xf5\x0e\x7b\xe5\xd0\x57\x5a\x73\x08\xf1\xc8\x4d\xac\x35\x8d\x96\xe7\xb3\x14\x9e\xd7\x22\x83\xa6\xb6\xc5\xb9\x00\xa2\x35\x4c\x6c\xd0\x92\xaa\xa3\xa7\xab\x40\x90\xfc\x10\x5b\x67\x0e\x8a\xf1\xd1\x9f\x9c\x42\xf8\x83\x28\x99\x9c\xda\xe3\x56\x71\xb3\xba\x59\x46\xec\x8e\xa7\xf6\x65\x2b\x38\x1a\x6e\x6a\xa7\x8e\xfe\xf5\x91\xa5\x21\x98\x7a\x75\xc9\xd4\x38\xb0\xa3\x8f\x29\x77\xf9\x98\xbd\x06\x75\x40\x3f\x5a\x2f\x2d\xb7\xc0\xf4\xaf\x86\x14\xde\x59\x8a\xf3\x57\x00\x00\x00\xff\xff\xd3\x5f\x66\x27\x82\x01\x00\x00"

func transactionsDestroyroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsDestroyroxCdc,
		"transactions/DestroyRox.cdc",
	)
}

func transactionsDestroyroxCdc() (*asset, error) {
	bytes, err := transactionsDestroyroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/DestroyRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0x35, 0x9f, 0x53, 0xe0, 0x5f, 0xe2, 0x4e, 0xa1, 0x45, 0xe8, 0xa3, 0xa6, 0xe9, 0x63, 0xd4, 0x26, 0xe5, 0xc4, 0x3c, 0x22, 0x8f, 0xaf, 0xde, 0xa3, 0x17, 0x80, 0xde, 0x46, 0x63, 0x89, 0x3f}}
	return a, nil
}

var _transactionsSetupuserCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcf\x6e\x9b\x40\x10\xc6\xef\x48\xbc\xc3\xd7\x1c\x5a\x2c\x39\x49\xcf\x96\x1b\xc9\x72\x12\xa9\x17\x37\xb2\xfd\x02\x0b\x0c\xb0\xca\xb2\x83\x86\x21\x89\x15\xf9\xdd\x2b\x30\x04\x1a\x5b\x6d\x6f\x96\xe7\x37\x7c\x7f\x76\x6c\x59\xb1\x28\x36\xec\x1f\x1b\x9f\xdb\xd8\xd1\x9e\x9f\xc9\x23\x13\x2e\xf1\xfd\x6d\xf3\xb8\x5f\xdd\xdf\x6f\x1f\x76\xbb\x30\xe8\xd1\x2d\xbf\xfd\x54\x2a\xeb\x09\xb2\xfe\xb5\xd9\x6f\x57\xeb\x11\x0d\x03\x15\xe3\x6b\x93\xa8\x65\x8f\xf7\x30\x00\x80\x4a\xa8\x32\x42\x51\x6d\x73\x4f\xb2\xc0\xaa\xd1\x62\x95\x24\xdc\x78\x9d\x0d\x0c\x70\x7b\x0b\x9b\x41\x0b\x82\x39\xcd\x90\x32\xd5\xfe\x9b\xc2\x38\x21\x93\x1e\x50\x98\x17\x82\x41\xc2\xce\x51\x27\x30\xac\xda\x0c\xa7\x6f\xdf\xc4\x2c\xc2\xaf\xcb\xaf\x83\xd7\x9b\xf5\x07\xfc\x24\xf6\xc5\x28\xdd\x45\xad\xfd\x05\x2e\x10\x3b\x65\x31\x39\x3d\x19\x2d\x66\xf8\xf2\x03\xde\xba\xd1\x1e\xe0\x38\x8f\xae\x76\x9d\x0e\x0a\x53\x23\xe1\xb2\x72\xa4\x94\xa2\x26\x45\x53\x21\xa6\x8c\x85\xae\x66\xe3\x8a\x90\x36\xf2\x61\xf3\x38\x89\x9a\x08\x19\x6d\xd3\x78\x7a\x05\x95\x95\x1e\x2e\xe4\x72\xa4\x93\x7f\xb1\xbc\x1e\x5d\x9f\xf6\x1f\xda\xc5\x31\x40\x34\x91\x06\x26\x6a\x75\xdb\x9c\x55\x28\x4f\x0b\x1e\x80\xbe\xbb\x16\x8a\x96\xd7\xa3\xe0\x1c\xca\xff\x2c\xaa\x7d\xf3\xf3\x54\x55\x13\x3b\x9b\x20\x31\x95\x89\xad\xb3\x7a\x40\xc6\xd2\x69\x9f\xa7\xec\xe5\x9d\xf5\xcf\x7f\x7b\xb8\xf7\xcf\xa7\x3a\x45\x3a\xb9\xf9\xe8\x75\xf8\xf1\x19\x39\xde\x45\x97\x14\xba\x59\x1b\x67\x0e\x35\x92\x93\xfe\x47\xec\x3f\x4b\xee\x8e\x63\x3d\xbd\x88\xa6\xea\x32\xf7\x5d\x2f\xc6\xb3\x68\xd1\x3e\xb3\x49\x53\xa1\xba\xee\x47\xc7\x30\x38\x86\xc1\xef\x00\x00\x00\xff\xff\x16\xd8\x8f\x7c\x99\x03\x00\x00"

func transactionsSetupuserCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsSetupuserCdc,
		"transactions/SetupUser.cdc",
	)
}

func transactionsSetupuserCdc() (*asset, error) {
	bytes, err := transactionsSetupuserCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/SetupUser.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5c, 0xfd, 0xd4, 0x7b, 0xdb, 0xb1, 0x31, 0x69, 0x4d, 0x31, 0x98, 0x41, 0x89, 0xca, 0xfb, 0xd6, 0x31, 0x77, 0x92, 0xf7, 0xb5, 0xf2, 0x11, 0xa7, 0x8e, 0xb6, 0x68, 0x5c, 0x46, 0x99, 0x84, 0xd7}}
	return a, nil
}

var _transactionsTransferroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\xcf\x8e\xda\x40\x0c\xc6\xef\x48\xbc\x83\x97\x43\x3b\x91\xda\xa8\x87\xaa\x87\x68\xff\x34\x0a\x45\xe2\x92\xae\x20\x7d\x80\x61\xe2\x84\x51\xc3\x78\xe4\x71\x0a\xd5\x8a\x77\x5f\x05\x02\xc9\xb2\x9b\x4b\x7c\xf8\x3e\xfb\xf3\xcf\x63\x77\x9e\x58\x20\x27\xb7\x68\x5d\x6d\x37\x0d\x16\xf4\x17\x1d\x54\x4c\x3b\xf8\x76\xc8\x17\x45\x3a\x9f\xaf\x7e\xad\xd7\xd3\x49\x2f\x5d\xd1\x61\x29\xb8\x0b\x23\x49\xf6\x3b\x2f\x56\x69\x36\x48\xa7\x13\x61\xed\x82\x36\x62\xc9\x29\x46\x63\xbd\x45\x27\x09\xa4\x65\xc9\x18\xc2\x17\xd8\x5b\xd9\x96\xac\xf7\xcb\x79\x02\x7f\x96\x4e\x7e\x7c\x8f\xe0\xa5\x33\x02\x00\x34\x28\x20\xb8\xf3\xf9\xa2\x48\xe0\xe7\x6d\xb6\x38\x5f\x14\x17\xa5\x67\xf4\x9a\x51\x05\x5b\x3b\xe4\x04\xd2\x56\xb6\xa9\x31\xd4\x3a\x19\xf5\xbb\xf4\x34\xd4\x34\x78\xca\xb4\xc2\x0a\x1e\xe0\xec\x8a\x37\xc4\x4c\xfb\xfb\x4f\x97\xcd\xe2\xec\xaa\x7b\x66\xfb\x4f\x0b\x3e\xaa\x6e\xd9\x04\x3e\x50\xac\x85\x58\xd7\xf8\xac\x65\x1b\x0d\xd3\xba\xef\xe9\x09\xbc\x76\xd6\xa8\x59\x46\x6d\x53\x82\x23\x81\xf3\x28\xd0\xc0\x58\x21\xa3\x33\x08\x42\x20\x5b\x04\xda\x3b\xe4\xcf\x61\x94\x71\x16\x8d\xf3\x07\x6c\xaa\xb8\x87\x02\xf7\x5f\xdf\xee\x12\x5f\x70\xaa\x31\xd7\xa1\xee\x93\x1d\xcf\x3f\x3c\xa0\x69\x05\xdf\xf3\xb9\x1e\x0a\x1e\xa0\x46\xe9\x41\x0e\xf7\x8b\x6e\x0d\x25\x7a\x0a\x56\xce\x34\xaf\xb2\xb8\x46\xc9\xb4\xd7\x1b\xdb\x58\xf9\xaf\x3e\xc2\xda\x6e\x1a\x6b\x4e\xcc\xee\xae\xf8\x5f\xde\x1d\xfa\xd6\x70\x7c\x54\xd1\xdd\x38\xc3\x30\x3f\xee\x4b\x25\x9d\x33\xe9\x08\x8d\x89\xbd\x4d\x4e\xb5\x9a\x75\x18\x8b\xee\x91\x56\xc8\x8c\xe5\xec\x84\xe8\x38\x9d\x1c\xa7\x93\xd7\x00\x00\x00\xff\xff\x91\xe7\xc2\x25\x14\x03\x00\x00"

func transactionsTransferroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsTransferroxCdc,
		"transactions/TransferRox.cdc",
	)
}

func transactionsTransferroxCdc() (*asset, error) {
	bytes, err := transactionsTransferroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/TransferRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf8, 0xa3, 0x18, 0x68, 0x4e, 0x36, 0x1a, 0xb0, 0x18, 0x3c, 0x77, 0x68, 0x56, 0xf3, 0x8f, 0xde, 0x1b, 0x5d, 0x10, 0x65, 0x3, 0x68, 0x9c, 0xc5, 0x78, 0x9a, 0x6d, 0x25, 0xfd, 0x50, 0x38, 0x9}}
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
	"scripts/nft_collection_length.cdc":  scriptsNft_collection_lengthCdc,
	"scripts/read_nft_id.cdc":            scriptsRead_nft_idCdc,
	"scripts/rox_items_total_supply.cdc": scriptsRox_items_total_supplyCdc,
	"transactions/AdminAssignRox.cdc":    transactionsAdminassignroxCdc,
	"transactions/DestroyRox.cdc":        transactionsDestroyroxCdc,
	"transactions/SetupUser.cdc":         transactionsSetupuserCdc,
	"transactions/TransferRox.cdc":       transactionsTransferroxCdc,
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
	"scripts": {nil, map[string]*bintree{
		"nft_collection_length.cdc": {scriptsNft_collection_lengthCdc, map[string]*bintree{}},
		"read_nft_id.cdc": {scriptsRead_nft_idCdc, map[string]*bintree{}},
		"rox_items_total_supply.cdc": {scriptsRox_items_total_supplyCdc, map[string]*bintree{}},
	}},
	"transactions": {nil, map[string]*bintree{
		"AdminAssignRox.cdc": {transactionsAdminassignroxCdc, map[string]*bintree{}},
		"DestroyRox.cdc": {transactionsDestroyroxCdc, map[string]*bintree{}},
		"SetupUser.cdc": {transactionsSetupuserCdc, map[string]*bintree{}},
		"TransferRox.cdc": {transactionsTransferroxCdc, map[string]*bintree{}},
	}},
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
