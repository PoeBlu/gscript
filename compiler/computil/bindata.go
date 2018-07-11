// Code generated by go-bindata. DO NOT EDIT.
// sources:
// compiler/templates/debugger.gs
// compiler/templates/entrypoint.go.tmpl
// compiler/templates/hard_reserved
// compiler/templates/obfstring.go.tmpl
// compiler/templates/preload.gs
// compiler/templates/soft_reserved
// compiler/templates/vm_file.go.tmpl
package computil

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

var _debuggerGs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x53\x70\x49\x2d\xc8\xc9\xaf\xd4\xd0\x54\xa8\xe6\x52\x50\x70\x49\x4d\x2a\x4d\x77\xce\xcf\x2b\xce\xcf\x49\xd5\xd0\xb4\xe6\xaa\x05\x04\x00\x00\xff\xff\x61\x33\x1d\xaf\x27\x00\x00\x00")

func debuggerGsBytes() ([]byte, error) {
	return bindataRead(
		_debuggerGs,
		"debugger.gs",
	)
}

func debuggerGs() (*asset, error) {
	bytes, err := debuggerGsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "debugger.gs", size: 39, mode: os.FileMode(420), modTime: time.Unix(1531104485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _entrypointGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x4f\x6b\xdb\x40\x10\xc5\xef\xfa\x14\xaf\xc2\x01\x89\xd8\x72\x9b\x63\xc0\x87\x96\xa4\x7f\x68\x53\x0a\x6d\xd2\x43\x48\xcd\x5a\x1a\x8b\xc5\xf6\xae\x3c\xbb\x92\x1d\x84\xbe\x7b\x19\xad\x02\x71\x0a\xa5\x14\x27\x17\x69\x77\x34\xfb\x9b\xf7\x76\xc4\x54\x2a\x5f\xa9\x92\xb0\x51\xda\x44\xd1\xb2\x36\x79\xbf\x4c\x52\xb4\x11\xd0\x28\xc6\xae\x84\xbb\x37\x79\xf6\x53\x69\xff\x81\x6d\x5d\x45\x40\xdb\x4e\xc0\xca\x94\x84\xd1\x7c\x8c\x51\xc5\x1a\xe7\x33\x8c\xb2\x6b\xa3\xb7\xdf\x58\x5b\xd6\x5e\x93\x43\xd7\xfd\x99\xdb\x6c\x24\x35\xd1\xa6\xa0\x3d\x46\xd9\x77\xcb\x9e\x8a\x9b\x2b\xd7\x53\xd2\x70\x64\x57\x66\x6f\x8b\x22\x79\x93\x46\x40\x69\x21\xaa\x06\x41\x40\x41\x4b\x12\x51\xd9\x85\x35\x94\xa4\x7d\x4c\xb5\xad\x80\xb3\x4f\x17\xe8\x3a\xc1\x7f\xa5\xdd\xe3\xd0\x90\xa6\x97\x4f\x32\x67\x33\x18\xbd\x1e\xc0\x00\x93\xaf\xd9\xf4\x9b\xae\x7f\x12\xb3\xd0\x0e\x0e\x65\xc3\xe6\x7d\x6d\x72\xaf\xad\xf9\x4c\xf7\x88\xf5\xa6\xb2\xec\xe7\xca\x39\xf2\x2e\x3e\xa8\x28\x8c\x57\x87\x75\x0e\x79\x97\xd9\x17\x5b\x96\xc4\xd9\x25\xb3\xe5\x65\x12\xf7\x6f\x04\xa4\x36\x25\x02\x15\x4b\xcb\xb8\xb9\xc2\x89\x3b\xc7\x49\x13\x8f\x11\x3f\xa6\xc4\x63\x29\x94\xfe\xcd\xc8\xbf\xf8\xa8\x98\xd6\x56\x15\xc7\x71\x20\x82\x07\xa0\xd8\x78\x56\xe1\x43\x03\x8c\xf2\xba\xa1\x63\x37\x20\x50\x9f\xfd\xf2\x07\x0f\x2e\x67\x5d\xf9\x63\x7b\x08\xd4\x17\xf9\x89\x68\x4f\x79\xed\x8f\xd4\x85\x00\x13\x07\xff\x27\x3b\x48\x68\x5b\x90\x29\x30\x79\x98\x2e\x32\xcb\x9e\x7e\xe8\x86\xf1\x57\x26\x2b\x70\x6d\x68\x8c\x05\x6e\xef\x64\x95\xc2\x79\x16\x09\xa2\x5c\x6e\xb0\x9f\x76\x61\xa6\x2d\x06\x3b\x8b\x5b\x7d\x87\x5f\x33\xac\xfa\xdd\x4a\x96\x12\x8a\x82\x8a\xad\x1c\x08\x90\xe4\xf5\xfe\xec\x4c\x4a\xbb\x31\xe6\x43\x38\xb7\xa6\xc9\xae\xcd\xb6\xb6\x9e\x92\x2d\x4e\x1f\x52\x17\x29\x4e\xb1\x95\xe4\xe0\x0a\x4e\x54\x4e\xa7\x98\x4c\x26\xf8\xa1\x4c\xb9\xa6\x02\x1f\x95\x66\x87\x77\xb4\xb6\xbb\x68\x3a\x8d\xa2\xe8\x77\x00\x00\x00\xff\xff\x08\x2f\x9c\x10\xd3\x05\x00\x00")

func entrypointGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_entrypointGoTmpl,
		"entrypoint.go.tmpl",
	)
}

func entrypointGoTmpl() (*asset, error) {
	bytes, err := entrypointGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "entrypoint.go.tmpl", size: 1491, mode: os.FileMode(420), modTime: time.Unix(1531038670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hard_reserved = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x2f\x4e\x2e\xca\x2c\x28\xe1\x4a\x4f\xcd\x33\x48\xce\x4c\x49\xe5\x72\x0f\x76\x0e\xf2\x0c\x08\xe1\xf2\x0f\x09\xf1\xe7\xca\x4d\xcc\xcc\x8b\x8e\xb6\x2a\xcf\x2f\x4a\xb1\x8a\x8d\xd1\x8b\x89\xd1\x57\x88\xd5\xb2\xe7\x72\xcd\x4b\xcf\xcc\x4b\xc5\x22\xe1\x57\x9a\x93\xe3\x93\x9f\x9e\x9e\x5a\x84\x29\x09\x08\x00\x00\xff\xff\x6d\x3e\x4c\x6b\x6a\x00\x00\x00")

func hard_reservedBytes() ([]byte, error) {
	return bindataRead(
		_hard_reserved,
		"hard_reserved",
	)
}

func hard_reserved() (*asset, error) {
	bytes, err := hard_reservedBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hard_reserved", size: 106, mode: os.FileMode(420), modTime: time.Unix(1531177142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _obfstringGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\xbd\x4a\x04\x41\x10\x04\xe0\xf8\xe6\x29\x8a\x61\xcc\xee\xc7\x4c\x10\x2e\xbb\xc4\x17\x30\xf1\x0c\xda\x9b\xde\x75\x40\x7b\x96\xde\x5e\x59\x19\xfa\xdd\xc5\x09\x16\xc1\xb4\xfb\xab\xaa\xd3\x09\xad\x21\x1d\x9f\x2e\x70\xc7\xad\x8a\x51\x91\x19\x24\x60\xb9\xe9\xf7\x64\x9c\xa1\x3c\x29\xcf\x2c\x46\x56\xaa\xa0\x0e\xb0\x77\xc6\x6c\x5a\x64\x44\xec\xf1\x67\xfa\x58\x18\xee\x31\x7c\x91\xfe\x6d\x3c\xe3\xe5\x55\x17\xe1\x16\x76\xad\x1d\xa0\x24\x23\x23\x95\xbc\xee\x91\xde\xf0\x78\x46\x3a\x5e\xc8\x08\x07\xf7\xb0\xeb\xa4\x0c\xf8\xac\xb9\x1b\x3c\x6c\x77\x4c\x5a\xc4\x06\xc4\xab\x5c\x2d\x62\xd3\x2c\xf9\xbf\xb9\x5f\xef\xd6\x3d\xe2\xef\x42\xff\xb5\xd6\x9d\x7b\xf0\x9f\x00\x00\x00\xff\xff\xdf\x60\x4c\x83\xf0\x00\x00\x00")

func obfstringGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_obfstringGoTmpl,
		"obfstring.go.tmpl",
	)
}

func obfstringGoTmpl() (*asset, error) {
	bytes, err := obfstringGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "obfstring.go.tmpl", size: 240, mode: os.FileMode(420), modTime: time.Unix(1531098832, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _preloadGs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x3f\x6f\xea\x30\x14\xc5\xf7\x7c\x8a\x23\x26\xfb\xc1\x8b\xc2\x5b\xde\x90\x52\x89\x3f\x5d\x3a\x94\x01\xb6\xb6\x83\x81\xeb\xe0\xd6\xd8\xd1\xb5\xd3\x2a\xaa\xf8\xee\x15\x09\xa0\x06\x51\xb5\x63\x6e\xce\xf9\xdd\xa3\x7b\xac\x2b\xb7\x8e\xc6\x3b\x2c\x22\x1b\x57\x2c\xfd\xa4\x8e\x34\x66\x56\xb5\x08\x12\x1f\x09\xf0\xa6\x18\x1b\x15\x15\x46\x78\x7c\xce\x13\x40\x7b\x86\x38\x4c\x0d\x46\xc8\x72\x18\xdc\x20\xa4\x96\x5c\x11\xb7\x39\x4c\xbf\xdf\xfa\xd0\xb8\xd2\xb2\x0a\x5b\x11\xd2\xf5\x56\xf1\xd4\x6f\x68\x1c\x85\x91\xf2\x80\xd9\x27\x00\x53\xac\xd8\x35\xc2\x3c\xd9\x27\xc9\x39\xcd\x39\xc5\xd2\xb7\xc1\x84\x6a\xa9\x47\x47\x3b\x4c\x35\xfb\xdd\xf4\x48\x4e\x55\x59\xda\x5a\xb4\x7f\x06\x50\xb2\x4b\x9c\x55\xbb\x52\xf8\xd5\x4b\x07\xd3\x7b\x72\x3d\xf4\x71\xbf\x98\x3f\xa4\xa1\x31\x1a\x5d\x1f\x54\x03\xb8\xca\xda\x01\xfe\x5d\x50\x26\xa4\x3d\xd3\x8c\x4a\xeb\x6b\xd1\x41\x45\xae\xe8\x62\xe3\x15\x95\x56\x36\x5c\xc8\xc6\x3a\x12\xff\x8e\x38\x77\x77\xcc\x9e\x7f\x44\x2e\x2c\x51\x29\x02\xad\xbd\xdb\x7c\x69\x31\x44\xc5\x11\x23\x38\x7a\xc7\x4c\x45\x12\x32\x2d\x28\x2e\xcd\x8e\x84\xfc\xae\xd8\x21\xfd\xef\x74\x6a\x34\x84\xb8\x06\xc0\xdf\x96\x2f\x71\x8b\xd3\x6a\xfc\xc1\x30\xcb\x32\x79\x32\x03\x2b\x26\xf5\x9a\x37\x1f\xfb\xe6\x0d\x74\x0f\xb6\xaa\x8a\xa9\x77\xc1\x5b\xba\x7a\x8a\xcf\x00\x00\x00\xff\xff\xf8\xdd\x34\xe2\xad\x02\x00\x00")

func preloadGsBytes() ([]byte, error) {
	return bindataRead(
		_preloadGs,
		"preload.gs",
	)
}

func preloadGs() (*asset, error) {
	bytes, err := preloadGsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "preload.gs", size: 685, mode: os.FileMode(420), modTime: time.Unix(1530143358, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _soft_reserved = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0a\x80\x50\x0c\x03\xd0\x3d\x97\xc9\x51\x74\xf9\xe0\xea\x50\xaa\xa8\xe4\xd3\x06\xcf\xff\x1f\x2f\x7d\x41\x70\x74\x54\x13\x2c\xc9\x04\x67\xdd\xff\xe9\x20\x52\xca\x37\x90\x9a\x4f\x62\xec\x07\x64\x0b\x9b\xad\x15\x00\x00\xff\xff\xe0\xe3\x66\x7a\x3a\x00\x00\x00")

func soft_reservedBytes() ([]byte, error) {
	return bindataRead(
		_soft_reserved,
		"soft_reserved",
	)
}

func soft_reserved() (*asset, error) {
	bytes, err := soft_reservedBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "soft_reserved", size: 58, mode: os.FileMode(420), modTime: time.Unix(1531122400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _vm_fileGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x6d\x6f\x1b\xb9\xf1\x7f\x6d\x7d\x8a\xf9\xef\xdf\x2e\x56\x81\xbc\x4e\xae\x45\x0b\xb8\x4d\x01\xc7\xd2\xe5\xdc\x38\x76\x60\xf9\x72\x28\x7c\x81\x8f\xda\x1d\xad\x19\xaf\xc8\x2d\x49\xd9\x56\x0c\x7d\xf7\x62\x86\xdc\x27\x59\xf2\xc3\x5d\x5e\xf4\x45\xe0\xd5\x92\x9c\xf9\xcd\x6f\x1e\x97\x29\x45\x7a\x2d\x72\x84\x99\x90\xaa\xd7\x93\xb3\x52\x1b\x07\x71\x0f\x20\x9a\x2c\x1c\xda\xa8\xb7\x15\xa5\x7a\x56\x1a\xb4\x76\x2f\xff\x26\xcb\x88\x96\x52\xb3\x28\x9d\xde\x13\xb4\xde\xfc\x4c\x65\x79\x85\x86\x4e\xa0\x4a\x75\x26\x55\xbe\x37\x11\x16\xff\xfa\x17\xde\x34\x9d\x39\x5a\x92\x3a\xea\xd1\xcf\x5c\xba\xab\xf9\x24\x49\xf5\x6c\x2f\x47\xf5\x3a\x95\x19\xee\xe5\x36\x35\xb2\x74\x7b\xa8\x72\xa9\x30\x5a\xd9\x66\xf4\x04\x8d\xbb\x36\x72\x86\x6a\x4f\x3b\xa7\x69\xc3\xfd\x3d\xc8\x29\x6c\x27\x3f\x09\x3b\xc4\xc9\x3c\xcf\xa5\xca\x47\x4a\x4c\x0a\xcc\x60\x77\xb9\x7c\x4a\x53\xc6\x67\x08\x34\x89\xda\x05\x54\x19\xf0\xa9\x96\xdc\x63\xfd\x52\xa9\x85\x26\x99\x7b\xd6\x09\x95\x09\x93\xad\x08\xf7\xd2\x8d\x50\x39\xc2\xb6\xb2\x03\xd8\x2e\xaf\x73\xd8\x7f\x0b\xdb\xc9\x7b\xfd\xc9\xbb\xe3\xdd\xe2\x44\xcc\xd0\x96\x22\x45\x8f\x67\x6f\x0f\x8e\xd8\x37\x52\xe5\x90\x61\x89\x2a\x43\x95\x2e\xa0\x72\xdf\xfd\x3d\x8b\x49\xe8\x18\x2c\x97\x30\xd5\x06\x94\x70\xf2\x06\x5b\xbb\x09\x77\xb5\xd1\x4b\xfb\x24\xdc\x15\x2c\x97\x81\xc9\x00\xb0\xdf\xeb\xdd\x08\x03\x71\x8f\x40\x07\x9c\x97\x03\xd8\xc6\xd9\x04\x33\x0f\x74\x44\x8f\xb6\x86\x46\x42\x79\x35\x39\x1a\x92\xf6\x2b\x5d\x64\x16\xdc\x15\x42\x26\x9c\x60\x30\xbc\x9c\x61\x06\x53\x59\x60\xeb\xc0\xa9\x91\x79\x00\xed\x41\x74\x04\xbd\xf5\x80\xfd\xab\x21\xc9\x22\xb0\xbd\x16\x9b\xfd\x5e\x2f\x00\x08\x47\x6e\x8d\x28\xbd\xee\x1c\x15\x5a\x69\xe1\xf3\x47\x46\xc0\x7b\x2a\x55\x6e\x51\x62\xfb\x94\x75\x66\x9e\x3a\xb8\xef\x01\x8c\xe0\x95\x0f\xc0\x64\xc4\x7f\x7a\x00\x1f\xe0\xe2\x0b\xa5\xc3\xb3\x02\x6e\x08\xaf\xaa\xb0\x4a\x86\xe1\xa1\x1b\x01\x4b\xc6\x7c\x82\xb7\x2d\x00\xa9\x41\xe1\xb0\x0b\xdc\xcc\x95\x93\x33\x64\xf4\x3e\xb4\xba\x46\x4c\xe7\x2a\xed\x8a\x89\xfb\xf0\xaa\x25\x94\xcc\x71\x48\x2e\x0b\x16\x9d\xe0\x6d\x1c\xb5\x65\x44\x03\xcf\xb0\x3f\x10\x0d\xe0\xcf\xaf\x07\x10\x8d\xee\x30\x9d\x3b\x8c\xfa\xcf\xc8\x04\x51\x90\xfc\x2a\xd6\x49\xc3\x38\x3c\x1f\x73\x1a\xc4\x4a\x16\x95\x92\x46\xe9\x54\x14\x16\xc3\x9f\x3e\xa3\x4c\xc6\xe8\xc2\x09\x51\xf4\x1f\xc9\xc7\x07\xb4\xf3\x86\x8c\xcd\xac\x89\x27\x43\x1d\xf6\x3b\x71\x0d\xa0\x69\xcf\x9f\x1a\x7b\x89\x1f\x80\xd1\x3e\x38\x1c\xf0\xe3\xd3\xde\x05\x18\xee\x43\x56\x6f\x6f\x41\xa4\x7f\x06\xdd\xdc\x28\xd0\xc1\xc5\xac\xe9\xc7\xb9\x4a\x9d\xd4\xea\x03\x2e\x20\xf2\x05\xf6\x52\x58\x8b\xce\x46\xe4\x23\xff\xc6\x82\x7f\x05\x52\x39\xfd\xf2\x20\x88\x75\xdb\xef\xfd\xe7\x28\x8e\xfb\x80\xc6\x50\x5e\x74\x0a\x12\xca\x6c\x6d\xaa\xef\x76\x72\x5d\x66\x3e\x3b\x37\xe5\xb1\x4e\x46\xa1\xc2\xd8\x8b\xa8\x39\x11\x7d\x81\xb7\x40\x80\xe3\x7e\xc8\x28\xf0\x2e\xd8\xdb\x83\x9f\x15\xd5\x32\x34\x9c\xc0\x25\xfd\x95\xee\x0a\x52\xad\xac\x83\x0c\xd3\x42\x18\xcc\x40\x4c\xf4\x0d\xf2\x89\xc0\x74\x63\xf4\x30\x5e\x29\x1e\xfd\xe0\x93\x10\x00\xde\x80\x70\x4c\xc9\x62\xa3\x8b\x4a\x83\x85\x16\x59\xc7\x39\xe4\x90\xca\x11\x61\x1d\x0a\x39\x31\xc2\x2c\xfe\xa8\x63\x5a\xea\x3a\x2e\xa9\x42\x29\x19\x25\xc7\x5a\x64\x63\xd6\x10\x57\xdb\x93\xdc\x46\x83\x2e\xcd\x97\x97\x9f\xce\x46\xc7\xa7\x07\xc3\xe8\x4b\xdc\xef\x3f\x15\x81\x1e\x31\x1b\x49\xf2\xbc\x89\x6b\xcc\x78\x18\x90\x9f\x3f\xbe\x30\xe6\x1a\x55\xcf\x31\x70\xa5\x50\xac\x98\x38\x3a\x39\x3f\xfb\xf7\xa7\xd3\xa3\x93\xf3\xc7\xad\xc4\x50\xc1\xc8\x82\xf0\xfc\xa0\x29\x08\xc7\x6f\x50\x39\xb3\x80\x52\x4b\xe5\x20\xce\x70\x2a\x15\x66\x30\x59\xc0\x0d\x1a\x2b\xb5\x1a\xc0\xdc\xce\x45\x51\x2c\x60\x88\x65\xa1\x17\xa4\xf5\x99\xf6\xb7\x40\xac\x66\xdb\x13\x55\x46\x27\xc3\xe4\x48\x7d\xc5\xd4\xf1\x86\x43\xad\xac\x2e\x30\x7e\x50\x16\x2f\x07\x24\x96\x12\x95\x78\x3a\x14\x45\x51\x01\xf8\x45\xba\xab\x73\x39\x43\x3d\x77\x71\xe4\x91\x73\x39\x0f\xac\xa3\x31\x4f\x45\x88\x1f\x1f\x7c\x1a\x30\x94\xd0\xd1\x17\x4a\xcc\x64\xca\x8c\x14\x52\x5d\x63\x56\x0d\x1a\xd3\x20\x64\x4d\x11\x0b\x2d\xf8\x8f\xd7\xaf\x06\x54\x87\x51\x9a\x57\x88\x08\x7e\xd3\x03\x7e\x7e\xcb\x59\xee\x8b\xd6\xee\x2e\xbc\x1b\xbd\x3f\x3a\x81\x93\x83\xf3\xa3\xcf\x23\xf8\x74\x70\xf8\xe1\xe0\xfd\x08\x8e\x3e\x7e\x3a\x3d\x3b\x1f\xbf\x74\x24\xdb\x7d\x38\x93\xad\x9d\xab\x60\xae\x32\xa4\x41\xac\x3a\x48\xbb\x54\x98\x9b\x2e\x55\x59\xff\xe4\xbe\x54\xf5\x68\x36\x31\x68\xf5\xe5\xb1\x11\xba\xbf\x69\x82\xf3\x0d\x89\x20\xb6\xb6\xd4\x69\xc4\x8b\xe3\xc5\x6c\xa2\x8b\x73\x8a\xb3\x7d\x98\x89\xf2\xc2\x3a\x23\x55\xfe\xe5\x55\x47\x31\xb1\x7e\xbf\x1c\x34\xc5\x33\xd0\x32\x95\xd9\xdd\x00\xb6\xa7\x4c\x0b\x09\x3f\x66\xe7\xd3\xf6\x56\x77\x38\xc8\x68\xee\xaf\x43\xc1\xa7\x15\x9a\xf6\x3c\x5a\xaf\xad\xa0\x4c\xe8\xf7\xb4\xf6\x3a\xb1\xe2\x34\x58\x06\x0d\x8e\x50\x33\x69\xd3\xcb\x9a\x35\x3e\x20\xb3\xbb\xb5\x04\xb2\x1d\xab\xa4\x74\xc4\x07\x5a\xe8\xcd\x3e\xe8\xa4\x35\x16\xf0\x4e\x7e\xac\x78\x68\x3b\x2b\x69\x11\x79\xb1\x46\x2c\x35\xb9\x06\x67\x07\xe6\x6a\x3b\xa2\x10\xe2\xdc\x22\xd2\x02\x3f\xd5\x48\xbf\xae\xf0\x56\x91\xdd\x54\xc5\x4e\xb4\xf8\xea\xc9\x20\xa3\x41\x07\x33\xa5\xbe\x9c\xf2\xe9\xff\xe3\xc4\x08\x8d\xb7\x55\x0e\xd6\x74\x4b\x9f\x39\xa3\x93\xe1\xe6\xbc\x59\x5f\x4f\x98\xbb\xd3\xe1\x87\xb0\x1c\x2a\x07\xf2\x87\x22\xb1\x74\x8d\xbe\x6f\xae\x56\x64\x9e\x4e\x7c\x49\xe8\xc8\xe9\xce\x0b\x0f\x7a\xff\xe9\xf0\x6c\xb8\x46\xfd\xd9\x90\xe6\x07\x27\xe4\xe3\x00\x50\xf1\x6b\x1a\x2f\xbc\x7e\xaa\x26\x2b\x72\xde\x06\xfd\xf7\xbd\xad\xd6\x57\x91\x4f\x89\x89\xaf\x14\xc3\x5a\xfa\x07\x5c\x1c\x18\x23\x16\xcc\xe2\x16\x1f\x90\x53\x98\xe9\x8c\x4f\xc0\xdf\xea\xf7\x50\x1a\xa9\xdc\x14\xa2\x5f\xd5\xaf\xdc\x20\xc3\xee\xca\x03\x9d\x3d\xaf\xef\x76\xee\x06\x10\x91\x3e\x5e\x6b\x06\xdb\x55\xcb\x87\x20\x6b\x83\x75\x37\x1f\x3b\x9f\x62\x61\xe4\x7c\xb4\x32\xb7\x07\x2c\x0b\xbe\x62\x74\x9c\x31\xa9\x5b\x90\x40\x4b\x73\xf7\x21\x5f\x03\xc4\x2b\xfe\x7b\x22\x00\x03\xbb\xd5\x1c\x9d\x4d\xde\x90\x44\x85\xb7\x31\x5f\x43\x24\xef\xe6\xd3\x29\x9a\x3e\x2f\xfd\xb0\x7e\x69\xcb\x9a\x94\x56\xfc\xdb\x13\xbc\x3d\x43\x91\xa1\x89\xbd\xe8\xd8\x32\x04\x72\xad\xbc\x81\x0b\xc2\xfa\xae\xd0\xe9\xf5\x58\x7e\xc3\xea\xdb\xce\x3a\x83\x62\x46\x32\xfc\x55\x06\x09\x39\xfd\xf1\x5d\x3c\x19\x80\xbc\xb9\xd8\xff\xd2\xef\x6d\x31\xa3\xc8\xe6\xfa\xbb\x0d\xda\x33\xf4\x2f\xe3\xf0\x66\xec\xb2\x51\xb8\x00\x19\x80\x35\x29\xe9\x45\x95\x7a\x38\x5c\xa8\x82\xf8\x31\xeb\xf3\xef\xef\xc7\xfb\x41\xff\x00\xce\xf6\x21\xe8\x59\x7a\xd6\x9a\x36\x2f\x75\x72\xa8\xcb\x45\x9c\x4d\xde\x0c\x1a\xa1\xfd\xbf\x3f\x8f\xd8\xad\xfc\x9b\xa9\x45\xe5\xdf\x64\xd9\xa2\x29\x9b\xbc\xe9\xf7\xb6\x56\x5d\xb4\xb5\xb5\x2a\x67\x6b\xd9\xdb\x0a\x78\xda\x70\x7e\x18\x40\xfe\xcd\x3c\x5b\x42\xfe\xcd\x24\x87\x85\xb6\x34\xd3\x54\x1b\xb2\xc9\x0f\xc9\x3b\x72\x5e\x5c\xcd\x75\x43\x9e\xfb\x29\x80\x87\xad\xa9\xc3\x37\x1e\xa8\x4b\xee\x2f\xfe\x83\xc1\xf6\x9a\x66\x75\x59\x77\x2a\xdf\xa6\x4c\xe2\xfb\xd4\x72\xd9\xcd\x94\x56\xa5\xaf\x72\x86\x67\x1a\xd3\xcd\x18\xbf\xad\x1e\x05\x36\x77\xac\x47\xa6\x99\x55\x7d\x31\x19\x03\xda\x39\x5d\x0b\xa0\xf9\xad\xef\x5f\x7d\x16\xc5\xdc\xe7\x17\x35\x55\x93\xcf\x67\xa8\x1c\x8c\x85\x92\x6e\x01\x87\x57\x98\x5e\x5b\x1f\x1a\x05\x2a\x96\x94\x54\x9b\x8e\xa5\x75\x7d\xf8\x27\x29\x2d\x50\x79\xdc\x07\x26\xb7\xd5\x7d\x00\x84\x89\x9b\x3f\x93\x47\x34\x2f\x4d\xe3\xc8\x69\x0d\x33\xa1\x16\x20\x82\x14\x0b\x25\x55\x87\x8c\x5a\x70\x4d\xc5\x8e\xa5\x99\x79\xc7\x56\x97\x06\x2b\xdd\x14\x18\x07\x19\x81\xe6\x58\xa7\x82\x16\x7c\xe2\xd7\xd1\xc8\x3b\x4e\xc9\xc0\x8f\xe2\x1a\x0f\xe7\xd6\xe9\x19\x63\x88\xa3\x5a\x0b\xcf\x70\xa4\x62\x33\x28\x6e\x8d\xd5\x81\xa8\xfa\xd0\xdb\xcc\xc6\x3f\x5e\xce\xc6\x14\x6f\xff\xe7\xc8\x58\x8b\x69\x1d\x17\x9d\xc1\x4d\xf8\x2e\x25\x38\x1d\x6a\xfb\xeb\x06\xef\x07\x88\x26\x9b\xea\x60\xfb\x7f\xb2\x4a\xd4\x73\x0b\x55\x4e\xd1\x7a\xc5\x51\x2d\x92\xd1\x9d\x1b\xcb\xdc\x6f\x31\xed\xf5\xba\xce\x74\xfc\x11\xb7\x76\xf4\x93\xd1\x1d\x8d\x30\xf1\xa6\xde\xb0\xc6\x35\xa9\x9e\x17\x19\x28\xed\x00\xf9\x6c\xcd\x07\xec\x64\xa0\xa7\x6b\xdd\xd3\x01\xf5\x52\x5f\xad\x62\x88\xd1\x98\xdf\xed\xc5\xe9\xcc\x25\x63\xdf\xcf\xdb\xa6\x38\x23\x94\x2d\x84\xc3\x8e\x35\xec\xd8\x5c\xc3\x0d\x95\x82\xae\x19\xfd\x2a\xe4\xed\xad\x74\xe9\x15\xdc\x10\xcf\x1d\xf6\x93\xd8\x2d\x4a\xec\x33\x91\xa9\xb0\xb8\xea\xad\x7d\x36\xa1\xe3\xcf\x07\x12\x56\x8e\x70\xf3\xc5\xa9\x98\x17\xce\x9f\x46\x63\x3e\x5a\xfe\x4c\xea\xd8\x55\x07\x10\xdf\xb1\xce\xa4\x9d\x09\x97\x5e\xed\x93\xc3\x30\xa5\x29\x6b\xc7\x0e\x20\xd7\x0e\x76\xce\xab\xdc\x69\x69\x89\x06\x70\xb3\x9e\xf9\x69\x1c\x89\x8e\xe8\x54\xab\xf0\x8d\xee\x09\xde\xf7\xee\xf6\xb0\x7e\xb7\x8f\x9a\xe3\xcb\x95\xdb\xc3\x26\xa5\x0c\x51\x34\x80\x6d\x53\xe7\xd4\x59\x98\x73\x37\xa6\x95\xdf\xe0\x93\xca\x74\x92\xca\xb4\x5e\xb1\x9b\x4c\x27\xa9\x3a\x03\x39\x2f\xa3\x3b\x46\x45\x9a\xeb\x8a\x56\x69\x6f\x5d\x95\xe6\xae\xde\xf9\xba\xbe\xa7\xa7\x10\xe7\x26\xb7\xf2\x29\x16\x24\x4f\x93\xf7\xa8\xd0\x08\x87\x5e\xe0\x98\xa7\x3d\x88\x4c\xd4\xdc\xf6\x3d\xab\x0d\xc6\x5d\x69\x07\x26\xaf\x44\x89\x28\x44\x92\x47\x89\xff\xa9\x51\xbe\xa9\x51\x8e\xa5\xca\x0b\xac\x7d\xa7\x95\xc3\x3b\x07\xbb\x20\x9c\xc3\x59\xe9\xa8\x0a\x87\x35\x77\x25\xf8\x27\x7f\x19\x19\x74\x9f\x45\xd1\x2d\x3b\xec\xf3\x73\xcd\xdd\x34\x36\xaf\x37\x15\x9a\x0d\x91\x1c\xb4\x70\x02\xb6\xa3\x6d\x2a\x64\x31\x37\xd8\xc4\x5b\x28\x0d\x1b\x4a\xc6\x34\x7e\x59\x48\xae\xc6\xf5\x83\xa0\x0c\x22\xbc\xc1\x21\x44\x0a\xdb\xfc\x4f\xd1\xf9\x95\xb4\x4d\x25\xbc\x12\x16\x66\xf3\xc2\xc9\xb2\x26\xd5\xc2\x6e\xb8\xd3\x21\xa7\x70\xa1\x11\xf0\xaf\x31\x08\xfe\x6a\xe1\xef\x81\xb5\x3e\x48\xf5\xac\x14\x4e\x4e\x64\x21\xdd\xa2\x07\xf0\xd5\x9e\x4e\xbe\xae\x61\xfc\x74\x42\xb2\xe3\xdf\x2e\xbe\xfc\xb6\x66\x22\x7c\x84\xf1\xa6\x26\xce\xc4\x35\x82\x50\x01\x92\x66\x81\x8c\xac\x36\x45\x58\x2b\x73\xc5\xe5\xc0\xc3\x8c\xbe\x0f\xfd\xd5\x3d\xef\x3a\xee\x1f\x54\x81\x01\x6c\x5f\x6e\x2e\x02\x21\xe9\xfd\x30\xb7\x92\xfa\x7e\x80\x66\x06\x93\x31\x86\x8b\xcf\xb0\x1c\x0d\x3a\x55\xe1\x85\x61\xdb\x90\x28\xb2\x0c\xb0\xc0\x56\x7f\xe4\xc2\xb9\x73\x4e\x49\xf3\xd5\x56\x64\x10\x41\x14\xcc\xfd\xd0\x65\x4c\xd5\x2c\x4d\xe7\xd7\x77\x0f\xf4\xc7\x98\xee\x96\xde\x86\xcb\xea\xc6\x80\xea\x4a\x16\xe2\xa2\x49\x09\x4f\xa7\x4f\xf7\xfe\x9a\xe2\xf9\x44\x9a\xa8\xaa\xb2\x58\x3f\x13\xb0\xb6\x7a\x87\x50\x59\xa5\x67\xae\xc2\x45\x71\xd2\x94\xce\xef\x52\x16\xab\x4b\x71\xe2\xeb\xe7\x4a\xc9\x5a\x7b\x96\xbd\x86\x20\xfa\xa8\x39\x17\x94\xb1\x19\xfc\x24\xa4\xb1\x30\xc1\x42\xdf\xf6\xf8\xfe\x66\xf7\xbf\x01\x00\x00\xff\xff\xb2\x95\xaa\x62\xe3\x1f\x00\x00")

func vm_fileGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_vm_fileGoTmpl,
		"vm_file.go.tmpl",
	)
}

func vm_fileGoTmpl() (*asset, error) {
	bytes, err := vm_fileGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vm_file.go.tmpl", size: 8163, mode: os.FileMode(420), modTime: time.Unix(1531269801, 0)}
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
	"debugger.gs": debuggerGs,
	"entrypoint.go.tmpl": entrypointGoTmpl,
	"hard_reserved": hard_reserved,
	"obfstring.go.tmpl": obfstringGoTmpl,
	"preload.gs": preloadGs,
	"soft_reserved": soft_reserved,
	"vm_file.go.tmpl": vm_fileGoTmpl,
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
	"debugger.gs": &bintree{debuggerGs, map[string]*bintree{}},
	"entrypoint.go.tmpl": &bintree{entrypointGoTmpl, map[string]*bintree{}},
	"hard_reserved": &bintree{hard_reserved, map[string]*bintree{}},
	"obfstring.go.tmpl": &bintree{obfstringGoTmpl, map[string]*bintree{}},
	"preload.gs": &bintree{preloadGs, map[string]*bintree{}},
	"soft_reserved": &bintree{soft_reserved, map[string]*bintree{}},
	"vm_file.go.tmpl": &bintree{vm_fileGoTmpl, map[string]*bintree{}},
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

