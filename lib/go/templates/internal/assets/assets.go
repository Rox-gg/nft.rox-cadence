// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../templates/scripts/borrow_nft.cdc (640B)
// ../../../templates/scripts/nft_collection_length.cdc (414B)
// ../../../templates/scripts/rox_items_total_supply.cdc (156B)
// ../../../templates/transactions/adminMintRox.cdc (835B)
// ../../../templates/transactions/destroyRox.cdc (391B)
// ../../../templates/transactions/setupUser.cdc (929B)
// ../../../templates/transactions/transferRox.cdc (793B)

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

var _scriptsBorrow_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x51\x4d\x6f\xd3\x40\x10\xbd\x47\xca\x7f\x78\xea\x01\x1c\x09\x25\x1c\x10\x87\x8a\x52\x05\x97\x20\x2e\xa6\x4a\xcd\x0f\x18\xaf\xc7\xc9\x08\x67\xc6\x5a\xcf\xaa\x45\x55\xff\x3b\x72\x5d\x07\x43\xf7\xb4\xab\x7d\x6f\xe6\x7d\xc8\xa9\xb3\xe8\x28\x4c\x77\x49\x0f\x52\xb5\x5c\xda\x2f\x56\x34\xd1\x4e\x78\xff\x50\xec\xca\xed\xcd\xcd\xfe\xeb\xdd\xdd\x72\xf1\x02\xdd\xdb\x43\x6e\xea\x91\x82\xcf\x50\xf9\x8f\xa2\xdc\x6f\xf3\xbf\xe8\xe5\x62\xb3\x41\x79\x94\x1e\x7d\x88\xd2\x39\x22\x53\xdd\xe3\xc4\x4e\x35\x39\x81\x2a\x4b\x0e\x52\x14\xbb\x12\xa2\x20\xa4\x9e\xe3\xdb\x1e\xc1\xda\x96\x83\x8b\xe9\x72\xd1\xa5\x0a\x4d\x52\x9c\x48\x34\xa3\x10\x2c\xa9\x5f\x62\x5b\xd7\x91\xfb\xfe\x1d\xa4\xbe\xc4\xcf\xef\xea\x1f\x3f\xac\xa6\x0b\x1e\x87\xd5\x00\xb0\xd9\xe0\x1b\x3b\xfc\xc8\xe8\x52\xd5\x4a\x98\x4d\x86\x35\xcf\x1f\x76\xaf\x1c\xa7\x87\x0f\xc6\x47\x6e\xcb\x3e\x43\xef\xb9\xc1\x15\x0e\xec\xdb\x51\xc1\xa4\x64\x35\x82\x87\xb3\x3e\xb0\xe7\xd4\x51\x25\xad\xf8\xef\x6c\x96\xd1\x3a\x3f\xcf\xb9\x7d\x96\x71\x4b\x7e\x9c\x33\x2b\x8b\xd1\xee\x3f\xbd\x79\xfc\xbf\x83\x57\xcc\xa7\xcf\xd9\x8c\x78\x7d\x8d\x8e\x54\x42\x76\x91\x5b\x6a\x6b\xa8\x39\xc6\x59\x08\x67\x25\x63\x43\xaf\xfc\x5f\xac\x66\x29\x7d\x19\x49\x84\xc8\x0d\x47\xd6\x30\x44\x01\x42\xdf\x71\x90\x46\xc2\xd4\xd0\x90\xd1\xbc\x9c\x29\x28\x6d\x1c\x57\xff\xc6\xf5\x62\xaa\xd8\x95\xd9\xd0\x91\xd4\xe7\x7d\x91\x3d\x45\x1d\x38\x6b\xa9\x97\x8b\xa7\x3f\x01\x00\x00\xff\xff\xf8\x6a\xd8\x13\x80\x02\x00\x00"

func scriptsBorrow_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsBorrow_nftCdc,
		"scripts/borrow_nft.cdc",
	)
}

func scriptsBorrow_nftCdc() (*asset, error) {
	bytes, err := scriptsBorrow_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/borrow_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x27, 0x94, 0x90, 0x5a, 0x7d, 0xc1, 0x66, 0xf, 0x42, 0xb4, 0xf2, 0x2b, 0xe, 0xc1, 0xbc, 0x5e, 0x5c, 0x7b, 0x4f, 0xf3, 0x60, 0x76, 0x7a, 0x9b, 0xf9, 0xcf, 0x66, 0xe6, 0x56, 0x63, 0xb7}}
	return a, nil
}

var _scriptsNft_collection_lengthCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x41\x6e\xc2\x30\x10\x45\xd7\x89\x94\x3b\x4c\x59\x54\xc9\x26\xea\x1a\x95\xa2\xc8\x14\x89\x0d\x45\x81\x0b\x38\x83\x09\x56\x9d\x99\xc8\x19\xab\x54\x88\xbb\x57\x21\x94\xa6\xed\xce\x8b\xf7\xfc\xff\x1f\xdb\xb4\xec\x05\xd6\x4c\xcb\x40\xb5\xad\x9c\xd9\xf1\xbb\x21\x38\x78\x6e\xe0\xe9\xb4\x5e\xee\x8a\xc5\xa2\x7c\xdd\x6e\x93\xf8\x86\x96\x7c\x52\x4c\xe2\x35\xca\x88\x52\x6f\xeb\x5d\x59\xa8\x1f\x3a\x89\xdb\x50\xc1\x21\x10\x34\xda\x52\xaa\x11\x8b\xfd\xde\x9b\xae\x9b\xc2\xed\x91\x4d\x61\x45\x02\xe7\x9e\x05\x00\x70\x46\x40\x23\x0a\xcc\xa0\x36\x52\x20\x72\x20\x19\x89\x59\x12\x47\x51\xd4\x53\xc8\xce\x19\x14\xcb\x54\x9a\x03\xcc\xae\x56\x5e\x1b\x51\xba\xd5\x95\x75\x56\x3e\xd3\x51\xcb\x5c\xdd\xf1\x4d\xa8\x9c\xc5\x8d\x96\x63\xf6\x90\x57\xec\x3d\x7f\x3c\x3f\x9e\xff\x8e\xff\x27\x5c\x5e\xd2\x21\x3c\x9a\xcf\xa1\xd5\x64\x31\x9d\x28\x0e\x6e\x0f\xc4\x02\xc3\x3f\x80\xf7\xf0\xe1\x2c\xed\x55\x1d\x75\x9d\x64\xdf\x4b\xbd\x91\xe0\xe9\xf7\x8c\xbe\xff\x6a\xd1\xa5\x59\xee\x0c\xd5\x72\x4c\xe2\x4b\x12\x7f\x05\x00\x00\xff\xff\x25\xcf\x4f\x62\x9e\x01\x00\x00"

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5f, 0x3a, 0x97, 0x13, 0xf2, 0x40, 0xce, 0x72, 0x72, 0x5b, 0x40, 0x7c, 0x9c, 0xb1, 0xea, 0xc0, 0x17, 0x58, 0xe3, 0x70, 0x4d, 0x86, 0x47, 0x32, 0x78, 0x1c, 0x7e, 0xbc, 0xba, 0xf7, 0x5c, 0x42}}
	return a, nil
}

var _scriptsRox_items_total_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x31\x0a\xc2\x30\x14\x06\xe0\x3d\x90\x3b\xfc\xa3\x2e\xe2\x20\x0e\x6e\xa5\xb5\xe0\x52\xa1\x89\x07\x48\x25\x95\x60\xf2\x5e\x08\x2f\x50\x11\xef\xee\x22\xd8\xfd\xe3\x0b\x29\x73\x11\x0c\x4c\x7d\xa5\x47\x98\xa2\xb7\xfc\xf4\x84\xb9\x70\xc2\x7e\x19\x7a\xdb\x74\xdd\x78\x36\x46\xab\x1f\x1d\x79\x69\x99\xa4\xb8\xbb\xac\x54\x7b\x1d\xec\xd8\xb4\x7f\xad\x55\xae\x13\xe6\x4a\x48\x2e\xd0\x66\x7b\xc2\xed\x42\x72\x3c\xe0\xad\x15\x00\x14\x2f\xb5\xd0\x7a\xdb\x09\x8b\x8b\xa6\xe6\x1c\x5f\x5a\x7d\xb4\xfa\x06\x00\x00\xff\xff\x17\xf8\x57\x1f\x9c\x00\x00\x00"

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x95, 0xb9, 0xfd, 0x2f, 0x77, 0x74, 0xd7, 0x2d, 0x11, 0xbc, 0x42, 0x6d, 0x61, 0x6, 0x6e, 0x2a, 0x5d, 0xc, 0x57, 0xa8, 0xe8, 0x15, 0xde, 0xf5, 0x35, 0xca, 0x50, 0xb8, 0x9d, 0x63, 0x8b, 0xc1}}
	return a, nil
}

var _transactionsAdminmintroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xdd\xca\x9b\x40\x10\xbd\x0f\xe4\x1d\x86\x5c\x14\x85\x0f\xe9\xb5\xf4\x4b\x10\x53\xa1\x17\xb5\x21\xfa\x02\xeb\x3a\x31\x4b\xd7\x1d\x59\xc7\x36\x25\xe4\xdd\xcb\xfa\x13\x4d\x93\xc2\xb7\x17\x1b\x36\x9e\x39\x67\xce\x99\x51\x75\x43\x96\x21\x25\x93\x74\xa6\x52\x85\xc6\x9c\x7e\xa2\x81\x93\xa5\x1a\x3e\x5f\xd2\x24\x8f\xf6\xfb\xe3\xd7\x2c\x5b\xaf\x46\xe8\x91\x2e\x31\x19\xb6\x42\xf2\x02\x15\xff\x48\xf3\x63\x14\xcf\xe8\xf5\x8a\xad\x30\xad\x90\xac\xc8\x78\x16\xa5\x6a\x14\x1a\x0e\x21\x2a\x4b\x8b\x6d\xfb\x06\x92\xb4\x46\xc9\x4e\xf3\x5b\x19\x42\xc6\x56\x99\xea\x0d\x58\xa1\x9d\x5e\x3e\x5c\xd7\x2b\x00\x80\xe1\xd6\xc8\x50\x2b\xc3\x0e\xf0\x69\xd1\x47\x90\x26\xf9\xf7\xfe\x7f\xa7\xeb\x90\x8d\xc5\x46\x58\xf4\x44\x59\x2b\x13\x49\x19\x42\xd4\xf1\x39\x92\x92\x3a\xc3\x3d\xeb\x80\x73\xa7\x45\x7d\x0a\x06\x5a\x78\x87\xa9\x22\x28\xc8\x5a\xfa\xfd\xe5\xb5\xce\xd6\x73\xd6\xc3\x65\x18\xc1\xf0\x25\x63\xb2\xa2\xc2\x83\xe0\xb3\x3f\x6b\xb8\xb3\xdb\x41\x23\x8c\x92\xde\x26\xa6\x4e\x97\x60\x88\x61\x10\x01\x01\x16\x4f\x68\xd1\x48\x04\x26\xe0\x33\x42\x9a\xe4\xa3\xd7\xcd\xc8\x73\x9b\x9a\xc6\x0b\xca\x8e\x71\xca\x66\x4a\xe6\x9e\x31\xbc\x43\x85\x3c\x9a\x9d\xa3\xf7\x97\xa6\xc7\x02\x54\xbf\x7a\xd7\x77\xd0\x63\xcb\x41\x85\x1c\x8b\x46\x14\x4a\x2b\xfe\xe3\x2d\xdd\xc6\xe3\xf4\xc8\x1c\xba\x42\x2b\xf9\xc2\xf0\x3d\xc2\xeb\xbf\xeb\xf5\x54\x7d\xdb\x7a\x1f\x48\xab\x5a\x36\xfd\x32\xb1\x99\x77\xe3\xff\x67\xc6\xfd\x4f\x9a\xe4\xcb\x9d\x9c\x48\x9f\x96\xf2\xe1\x39\xed\xa6\xbb\xe7\x99\xdc\xfe\x06\x00\x00\xff\xff\xfd\x4e\x82\xbf\x43\x03\x00\x00"

func transactionsAdminmintroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsAdminmintroxCdc,
		"transactions/adminMintRox.cdc",
	)
}

func transactionsAdminmintroxCdc() (*asset, error) {
	bytes, err := transactionsAdminmintroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/adminMintRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5d, 0x47, 0x74, 0x98, 0x92, 0xf, 0x37, 0x2, 0xc5, 0x84, 0xd, 0x65, 0x8f, 0x80, 0xc, 0xb1, 0x79, 0x11, 0x54, 0x3e, 0xda, 0x9e, 0xa8, 0x8f, 0xc4, 0xc9, 0x50, 0xbe, 0x67, 0x40, 0x95, 0xa7}}
	return a, nil
}

var _transactionsDestroyroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xc1\x4e\x83\x40\x10\x86\xcf\x6c\xb2\xef\x30\xde\x20\xb1\xc4\x83\xf1\x40\x88\x91\x80\x4d\x7a\x41\x03\xf8\x00\xe3\xb2\x2d\x1b\xb7\x3b\x64\x19\x6c\x8d\xe9\xbb\x1b\xa8\x4a\x0f\x1e\xbc\x6d\x76\xbe\xff\xff\x3f\xb3\xef\xc9\x33\x94\xe4\xd6\xa3\xdb\x99\x57\xab\x1b\x7a\xd3\x0e\xb6\x9e\xf6\x70\x73\x2c\xd7\x4d\x56\x14\xd5\x63\x5d\x4b\xf1\x8d\x56\x74\xcc\xc9\xb1\x47\xc5\x17\x54\xfe\x54\x36\x55\x96\x2f\xb4\x14\xec\xd1\x0d\xa8\xd8\x90\x0b\x4d\x9b\xc0\xcb\xc6\xf1\xdd\x6d\x04\x9f\x52\x04\xbd\xd7\x3d\x7a\x1d\xa2\x52\x9c\x40\x36\x72\x97\x29\x45\xa3\xe3\xf9\x2c\x45\x10\x58\xcd\xa0\xc8\x5a\x3d\x17\x40\xba\x82\x89\x8d\x2d\x61\x9b\x3e\x5c\x38\xc4\xf9\x2f\x74\x1f\x4e\x3e\x09\xfc\x7d\xad\x99\x3c\xee\xf4\x33\x72\x17\x5d\x2d\x1b\x6e\xcb\x53\xf9\x32\x15\x1f\x0c\x77\xad\xc7\x43\xf8\xf3\xd8\x14\x09\x98\x36\x3a\x67\x5a\x3d\xb0\xa7\x8f\x29\x77\xfe\x98\xb5\x06\x7c\xd7\x61\xba\x5a\x5a\xae\x81\xe9\x3f\x26\x52\x04\x27\x29\x4e\x5f\x01\x00\x00\xff\xff\x9f\x48\x8b\x2e\x87\x01\x00\x00"

func transactionsDestroyroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsDestroyroxCdc,
		"transactions/destroyRox.cdc",
	)
}

func transactionsDestroyroxCdc() (*asset, error) {
	bytes, err := transactionsDestroyroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/destroyRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x5, 0x81, 0x86, 0xb5, 0x19, 0x13, 0x7d, 0x39, 0x5c, 0x81, 0x31, 0x60, 0x57, 0x39, 0x27, 0xd3, 0x2b, 0x9a, 0x49, 0x31, 0x4a, 0x1d, 0xac, 0xa, 0x8a, 0xb, 0x53, 0x78, 0xe9, 0xd6, 0xa2}}
	return a, nil
}

var _transactionsSetupuserCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcd\x4e\xe3\x40\x10\x84\xef\x96\xfc\x0e\xb5\x1c\x76\x1d\x29\xc0\x9e\xa3\x2c\x52\x64\xe0\x98\x45\x49\x5e\x60\x6c\xb7\xed\x11\x93\x69\xab\xdd\x5e\x12\xa1\xbc\xfb\xca\x3f\xc1\x06\x82\xc4\xd5\x5d\xed\xaa\xaf\xa6\xed\xbe\x62\x51\xac\xd9\x3f\x36\xbe\xb0\x89\xa3\x1d\x3f\x93\x47\x2e\xbc\xc7\xef\xc3\xfa\x71\xb7\xba\xbf\xdf\x3c\x6c\xb7\x61\x30\x48\x37\x7c\x88\xd9\xab\x98\x54\x27\xaa\xf8\xef\x7a\xb7\x59\xc5\xa3\x3a\x0c\x54\x8c\xaf\x4d\xaa\x96\x3d\x5e\xc3\x00\x00\x2a\xa1\xca\x08\x45\xb5\x2d\x3c\xc9\x02\xab\x46\xcb\x55\x9a\x72\xe3\x75\x76\xd6\x00\xb7\xb7\xb0\x39\xb4\x24\x98\x7e\x86\x8c\xa9\xf6\xbf\x14\xc6\x09\x99\xec\x88\xd2\xfc\x23\x18\xa4\xec\x1c\x75\x06\xe7\x55\x9b\xa3\xff\xf7\x4d\xc2\x22\xfc\xb2\xfc\x39\x89\x7b\x13\xbf\xe9\xef\xa2\x36\xfa\x02\x97\xa7\x5b\x65\x31\x05\x3d\x19\x2d\x67\xf8\xf1\x07\xde\xba\x31\x1d\xe0\xb8\x88\xae\xb6\x9d\x0d\x4a\x53\x23\xe5\x7d\xe5\x48\x29\x43\x4d\x8a\xa6\x42\x42\x39\x0b\x5d\xcd\xc6\x15\x21\x6d\xe4\x2d\xe5\x69\x42\x9a\x0a\x19\x6d\x61\x3c\xbd\x80\xf6\x95\x1e\x2f\x60\x39\xd2\xc9\x57\x2c\xaf\xdf\x05\xef\x7f\xf1\xd0\xee\x8e\x0c\xd1\xc4\x1d\x98\x18\xd6\x6d\x77\x56\xa1\x3c\xad\xf8\x2c\x18\xda\x6b\x45\xd1\xf2\x7a\xf4\x9c\x43\xf9\x3b\x75\xb5\x0f\xff\x99\xad\x6a\x12\x67\x53\xa4\xa6\x32\x89\x75\x56\x8f\xc8\x59\x3a\xfb\xcf\xac\x43\x02\x67\xfd\xf3\x57\xaf\xf7\xfa\xf1\x5c\x27\xb3\xa7\xce\x6a\xfe\x45\xd4\x0d\x1f\x7a\xc1\xe9\x2e\xba\xac\xe8\xc7\x2d\xcb\x1c\x6a\xa4\x20\xfd\x1e\xf6\xfb\x9e\xbb\x13\x89\xa7\x77\xd1\x54\x1d\xf3\x50\xf7\x62\x3c\x8e\x56\x3a\x30\x9b\x2c\x13\xaa\xeb\x61\x74\x0a\x83\x53\x18\xfc\x0f\x00\x00\xff\xff\x87\x43\x37\x6f\xa1\x03\x00\x00"

func transactionsSetupuserCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsSetupuserCdc,
		"transactions/setupUser.cdc",
	)
}

func transactionsSetupuserCdc() (*asset, error) {
	bytes, err := transactionsSetupuserCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/setupUser.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0xb1, 0x3e, 0xa5, 0xf5, 0x8f, 0x42, 0xfe, 0xeb, 0xfc, 0x4a, 0x75, 0xe4, 0xfb, 0x56, 0x9f, 0x3e, 0x23, 0xcb, 0xad, 0x0, 0x3c, 0x62, 0x1, 0x98, 0x8f, 0x78, 0x6f, 0xc4, 0xb5, 0x84, 0x2c}}
	return a, nil
}

var _transactionsTransferroxCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4f\x8f\x9b\x30\x10\xc5\xef\x91\xf2\x1d\x66\x73\x68\x8d\xd4\xa2\x1e\xaa\x1e\xd0\xfe\x29\x22\x8d\xb4\x17\xba\x4a\xe8\x07\x70\xcc\x40\xac\x12\x8f\x35\x0c\x4a\xaa\x55\xbe\x7b\x45\x42\xc0\x9b\x5d\x2e\x58\xe2\x3d\xcf\x9b\xdf\xc3\xee\x3d\xb1\x40\x4e\x6e\xd5\xb9\xda\x6e\x1b\x2c\xe8\x2f\x3a\xa8\x98\xf6\xf0\xed\x98\xaf\x8a\x74\xb9\x5c\xff\xda\x6c\xe6\xb3\x41\xba\xa6\x63\x46\x4e\x58\x1b\x09\x54\xd9\xef\xbc\x58\xa7\xd9\xa4\x9e\xcf\x84\xb5\x6b\xb5\x11\x4b\x4e\x31\x1a\xeb\x2d\x3a\x49\x20\x2d\x4b\xc6\xb6\xfd\x02\x07\x2b\xbb\x92\xf5\xe1\x79\x99\xc0\x9f\x67\x27\x3f\xbe\x47\xf0\xda\x1b\x01\x00\x1a\x14\x10\xdc\xfb\x7c\x55\x24\xf0\xf3\x36\x5e\x9c\xaf\x8a\xab\xd2\x33\x7a\xcd\xa8\x5a\x5b\x3b\xe4\x04\xd2\x4e\x76\xa9\x31\xd4\x39\x09\xee\xbb\xde\x69\xa8\x69\xf0\x9c\x69\x8d\x15\x3c\xc0\xc5\x15\x6f\x89\x99\x0e\xf7\x9f\x82\xe5\xe2\x6c\x94\x3e\xaa\x7e\xd1\x04\x3e\xfe\xba\x11\x62\x5d\xe3\x8b\x96\x5d\x34\x0d\xeb\x9f\xa7\x27\xf0\xda\x59\xa3\x16\x19\x75\x4d\x09\x8e\x04\x2e\x93\x40\x03\x63\x85\x8c\xce\x20\x08\x81\xec\x10\xe8\xe0\x90\x3f\xb7\x41\xc4\x45\x14\xc6\x6f\xb1\xa9\xe2\x81\x09\xdc\x7f\x7d\xbb\x4a\x7c\xa5\xa9\x42\xac\xd3\x79\x48\x76\xba\xbc\xf0\x88\xa6\x13\x7c\x8f\x67\xec\x09\x1e\xa0\x46\x19\x38\x4e\xf5\x45\xb7\x86\x12\x3d\xb5\x56\x2e\x30\x47\x59\x5c\xa3\x64\xda\xeb\xad\x6d\xac\xfc\x53\x1f\x73\x7b\xe9\xb6\x8d\x35\x67\x6c\x77\x63\x01\xaf\xef\xaa\xbe\x35\x9c\x1e\x55\x74\x17\xc6\x98\x22\xc4\xc3\x51\x49\xef\x4c\x7a\x48\x21\xb4\xb7\xe1\xa9\x56\x8b\x9e\x64\xd1\xff\xa6\x15\x32\x63\xb9\x38\x53\x3a\xcd\x67\xa7\xf9\xec\x7f\x00\x00\x00\xff\xff\x69\xa0\xb0\x0d\x19\x03\x00\x00"

func transactionsTransferroxCdcBytes() ([]byte, error) {
	return bindataRead(
		_transactionsTransferroxCdc,
		"transactions/transferRox.cdc",
	)
}

func transactionsTransferroxCdc() (*asset, error) {
	bytes, err := transactionsTransferroxCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transactions/transferRox.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0xe3, 0xc6, 0x41, 0x5, 0x65, 0xe1, 0x22, 0x39, 0xf8, 0x47, 0xe, 0xc2, 0x4, 0xe8, 0x77, 0xe7, 0xee, 0x3e, 0x31, 0x98, 0xf3, 0x8f, 0xb8, 0x27, 0x73, 0xa6, 0x42, 0x13, 0x9d, 0xea, 0xaa}}
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
	"scripts/borrow_nft.cdc":             scriptsBorrow_nftCdc,
	"scripts/nft_collection_length.cdc":  scriptsNft_collection_lengthCdc,
	"scripts/rox_items_total_supply.cdc": scriptsRox_items_total_supplyCdc,
	"transactions/adminMintRox.cdc":      transactionsAdminmintroxCdc,
	"transactions/destroyRox.cdc":        transactionsDestroyroxCdc,
	"transactions/setupUser.cdc":         transactionsSetupuserCdc,
	"transactions/transferRox.cdc":       transactionsTransferroxCdc,
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
		"borrow_nft.cdc": {scriptsBorrow_nftCdc, map[string]*bintree{}},
		"nft_collection_length.cdc": {scriptsNft_collection_lengthCdc, map[string]*bintree{}},
		"rox_items_total_supply.cdc": {scriptsRox_items_total_supplyCdc, map[string]*bintree{}},
	}},
	"transactions": {nil, map[string]*bintree{
		"adminMintRox.cdc": {transactionsAdminmintroxCdc, map[string]*bintree{}},
		"destroyRox.cdc": {transactionsDestroyroxCdc, map[string]*bintree{}},
		"setupUser.cdc": {transactionsSetupuserCdc, map[string]*bintree{}},
		"transferRox.cdc": {transactionsTransferroxCdc, map[string]*bintree{}},
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
