// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\x0a\x02\x31\x10\x45\xfb\x39\xc5\x4f\xa7\x57\x48\x67\x23\x58\x28\x88\xa5\x58\x0c\xeb\xec\x1a\xd8\x24\x4b\x66\x14\x17\xd9\xbb\x0b\x81\x88\x62\x37\xbc\x79\xbc\x2f\x4f\x93\x74\x85\xcd\x93\xe0\x78\x97\x32\xe3\x45\x00\x17\x0b\x3d\x77\xa6\xab\x76\x1d\x38\x8a\xc7\xc9\x4a\x48\x83\x5b\x7b\x6c\x9a\xb1\x4b\x7d\x76\xb4\x10\xd5\xc4\x0f\xae\xa9\x60\x12\xd5\xe3\xdc\x3e\xee\xf2\x6f\x57\xf1\x21\x45\x43\x4e\x9f\x11\x02\xba\x1b\xa7\x41\xc6\x3c\x7c\x43\x0b\x51\xd4\x38\x4e\x7b\xf5\xd8\x8e\x99\xcd\xd1\xf2\x0e\x00\x00\xff\xff\xc2\xab\x64\xad\xc7\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 199, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x5d\x6f\xe3\x36\xd6\xbe\xf7\xaf\x38\xd3\xb9\x68\x02\x04\x83\xe2\xc5\xdb\x62\xe1\xab\x55\x6d\xb5\xa3\x4d\xe2\x78\x63\x67\x66\x8b\xc5\x60\x40\x8b\xc7\x16\x11\x89\x54\x49\x2a\x89\x77\x31\xff\x7d\x71\xf8\x21\x89\xb6\xd2\xd9\x29\xf6\xce\xe2\xc7\x73\x9e\xf3\xc9\x43\x1a\x5f\x2c\x4a\x0e\xf6\xd8\x22\xfc\xbd\x43\x7d\x84\x7f\xcf\x00\x3a\x83\x7a\x0e\x0f\x06\x75\x21\xf7\xea\xcd\x0c\x40\xe9\xc3\x1c\xee\xf4\x21\x7e\xd3\x8a\x0d\x5a\x2b\xe4\xc1\xf8\x95\xf1\x2b\xce\x66\xd6\x6a\xb1\xeb\x2c\x86\xf9\xe1\x3b\xe0\xd1\xa0\x99\xc3\x3f\x7b\x31\x9f\x68\xa2\xac\x3b\x63\x51\x5f\x08\x3e\x87\x62\xf9\xe6\x72\x0e\x0b\x3f\x12\x25\x87\x05\x3f\x1f\x57\xac\xc1\x0b\xc9\x1a\x9c\xc3\xc6\x6a\x21\x0f\xaf\x2f\x26\x31\xe3\x99\xb1\xa4\x85\x92\x12\x4b\x2b\x94\x3c\x97\x39\xcc\x0d\x80\x22\xd3\x56\xec\x59\x69\x2f\x58\xf8\xb1\x3d\xb6\x38\x87\x6c\xf4\xe5\x20\x6e\x8a\x38\x44\x1b\x59\x67\x55\xa9\x9a\xb6\x46\x8b\x17\x42\xb6\x9d\x8d\xb4\xaf\xa0\xec\xb4\x51\x7a\xad\xcc\x1c\x0a\x69\xaf\x80\x39\x91\x73\xc8\x46\x7b\x32\x37\x46\xe0\x57\x91\xf9\x43\xb1\x8c\x18\x97\xe9\xe2\x7b\x34\x5d\x7d\x26\xf6\x17\x81\x35\x3f\x95\xbd\xa7\xc1\xa0\xc1\x68\x6d\x2e\xad\xb0\xc7\x6b\x21\xf9\xd5\x0c\x00\x40\xe3\xef\x9d\xd0\xc8\x33\x7d\xa0\xc5\x64\xd0\xe9\xe5\x9f\x5e\xa1\x97\x2c\xdf\x74\x87\x03\x1a\x52\xe8\xd3\x6c\x06\xf0\x16\x36\xa5\x16\xad\x6d\x0e\x1a\x50\xf2\x56\x09\x69\xcd\x15\x68\xdc\xa3\x06\xab\x80\xab\xd2\x80\x90\x50\xd6\xaa\xe3\xac\x15\xef\x5a\xad\xac\x9a\x01\xd4\xe2\x09\x3f\x08\x7c\x26\x3a\x37\xe1\xf7\x2d\x5a\xc6\x99\x65\xde\xc9\x71\xc5\x42\x49\x8b\xd2\x9a\x91\x8f\x6f\x4e\xa6\x68\xb9\x71\x3c\x08\xce\x33\x4a\xc1\xfc\xec\x04\xd4\x26\x99\x78\xe3\x75\x5a\x62\x5b\xab\x23\x3c\xe2\xd1\xcc\x00\xb8\xfb\x6a\x50\xda\x6b\x3c\x92\x80\xe5\x78\xc0\xe3\x27\x6b\x46\xf0\xe9\x52\x8f\x9e\xad\x8b\x08\xcd\x5a\x11\x30\xb3\x75\xd1\x83\xf9\xd1\x11\x4a\x98\x9c\x7d\x99\xcd\xc6\x59\x7f\xdb\x59\x46\x9e\x70\x89\xbf\xd0\xc8\x2c\x86\xe8\x4f\xb2\x09\xfe\xca\xb1\xd5\x58\x32\x8b\xfc\x42\x23\x33\x14\xa0\xdf\x85\x05\x06\x98\x46\x90\xea\x19\x4a\x07\xc0\xe1\x49\x30\x68\x5f\x82\x46\xdf\x5d\xce\x00\x1e\x5a\xce\x2c\x7e\x10\xff\x12\x2e\xaf\xf6\xe2\x70\x11\x02\x85\xe2\xa4\x58\xbe\xb9\x82\xa7\xd1\xe4\x1c\x72\x2e\x2c\xdb\xd5\xc9\x96\x89\x14\xf7\x94\x13\x13\x9d\x59\x0c\x60\x89\x14\x77\xcb\x57\x0c\xfc\xb3\x52\x35\x32\x39\xc0\x79\x5b\x0d\x36\x8b\x00\xfe\x7b\x7a\xa7\x57\x70\x5c\x0a\x2f\x4c\x5f\x21\xa3\x32\x49\xa5\xbc\x3c\xaf\x9c\x1b\xb4\x69\xb1\xbc\x60\xa3\x3a\x3a\x46\x19\xd5\xd3\xcb\xa9\x0a\x5b\xc8\x27\xe1\xe9\x5c\x60\xc3\x44\xdd\x57\x49\xca\x79\x6d\xec\x6a\x5c\x39\xaf\xa0\x66\x27\x43\x97\xf1\x00\x20\x98\x54\xbf\x35\xea\x46\x18\x23\x94\x34\x17\x54\xea\x7b\x07\x76\xe9\x64\x4a\x78\x34\x31\x80\x7b\x1f\x7a\xe8\x3b\x7d\xe8\x2d\xa7\xf4\xa1\x47\x55\xc3\xf8\x80\x38\x5a\x4c\x68\xfd\xd1\xf4\x65\x36\x73\x61\x1d\xe1\x5d\x58\x07\x7f\xcd\x00\x92\xf3\x62\x06\x90\x9a\x66\x06\xd0\x8a\xd2\x76\x3a\x59\xa3\xf4\x61\x75\xb2\x2d\xd0\x1b\x06\x84\xc9\xda\x56\xab\x27\xe4\xa3\x98\x88\x5c\x02\xb9\xaf\x51\x91\xa4\x97\x87\x61\xb5\x99\xc0\x19\x07\x8b\x03\x63\x92\xd5\x47\x2b\x4a\x73\xd7\x5a\x45\x55\x7d\x14\x8e\x51\xd0\x78\xf3\x10\x21\x6e\xbb\x55\x9d\xde\x20\xca\xd7\xf6\xb9\xa3\xe2\x95\xa0\x9b\x06\x98\xde\xf5\x5f\x71\xee\x89\xfa\x1c\x3b\x31\xd6\x23\x65\xe3\x60\xab\x50\x66\x32\x7b\x6b\xe6\xf0\x4b\xad\x98\xf5\xf5\xd3\x94\xc3\xaa\x88\x97\x24\xfd\xff\x04\x16\x65\xd7\x4c\x9c\x96\x1b\xcb\x2c\x3a\x01\x59\xbe\xf9\xfc\xb0\xba\x5e\xdd\x7d\x5c\x85\xaf\x75\xbe\x5a\x16\xab\x5f\xc3\xd7\xfd\xc3\x6a\x35\x7c\xfd\x92\x15\x37\xf9\x32\x7c\x6c\xf3\xfb\xdb\x62\x95\x6d\xf3\xe5\xa4\xa4\xa1\x0d\xf0\x82\xb2\xed\x48\xd0\x5b\xc8\x24\x20\x17\x36\x74\x10\xa0\x4a\x6a\x2d\x40\xec\x81\xb9\xdc\x84\x8a\x19\x68\x14\x17\x7b\x81\x1c\x6c\x85\xe0\x9d\x65\xf1\xc5\xc2\xee\x08\x42\x1a\xd4\xe4\x2a\x50\x1a\x38\x55\x3c\xfa\x5d\x56\x4c\xb3\x92\xca\xfc\x3b\x27\x64\x5b\x09\x3a\x8e\xcb\xba\xe3\x68\xe8\x10\x71\x1b\xa4\xc3\x7b\xc4\xe3\x4e\x31\xcd\x81\x49\x0e\x2d\x33\x1e\x40\x35\x0d\x93\xdc\x6d\x27\xc6\xf9\xb2\xd8\x7a\xba\x60\xb0\xc6\x72\xe0\x2b\xeb\xe3\x34\xe9\xb2\x52\x06\x25\x30\x99\x74\x34\x60\xfa\x46\xe2\x5d\xa4\xc5\x05\x9d\x51\x06\x5c\x83\xf0\xd6\x91\x4a\xb6\xd8\x8a\x59\x10\x16\x4c\xa5\xba\x9a\x43\xa3\x9e\xd0\x2d\x22\x51\xdf\x9b\xd0\x8b\x51\xd7\x41\x83\x92\x0c\xc3\x28\x25\x5b\x2d\xc8\xbb\x96\xed\xa2\x16\x9b\xfc\x26\x5f\x6c\xff\x20\x1e\xa8\x1d\x0a\xe1\x70\x9d\x84\xc3\xf5\xe7\xf5\xdd\x32\xfc\xda\x7c\x58\xc4\x5f\x8b\xfb\x62\xbd\x0d\x1f\xab\xec\x36\xdf\xac\xb3\x45\x3e\xa4\xc5\x64\xff\xe4\xf0\x1f\x85\xe4\xaf\xb5\x6f\x27\x85\x26\x84\x33\xb5\x2b\xae\xc5\xec\x47\x1b\x66\xcb\x0a\x79\x21\x39\xbe\xb8\xf6\xae\x90\xf6\x13\xf5\x3c\x14\xd4\x53\xe0\x2e\xda\x7b\x76\x5b\xb6\x3b\x21\x45\x71\x42\xf1\xc5\xf1\x05\xd4\xde\x59\xd3\xb2\x9d\x37\xbf\xad\xd0\x8c\x9d\xe7\xfb\x87\xbd\xd2\x64\x5b\xcb\x76\x8e\x85\x6b\x86\x1d\xd0\xc7\x0a\x6d\x85\x3a\x04\x0b\x45\x14\x1b\x6d\xa6\x7d\x60\xc9\xf9\x84\xef\x05\x3e\x8b\xba\x86\x86\x3d\x7a\xd7\x86\xf8\x03\x7c\xc1\xb2\x73\x55\x89\xe4\x0c\x5f\xd9\xde\x52\x91\x22\xf0\xa1\x1c\xc1\x98\xdf\x1f\xf4\xaf\x53\xfe\xf1\xfd\xf7\xc8\x0c\x7b\xa5\x1b\x66\xa9\x31\xf2\x09\x47\x64\xfb\xec\x33\xa1\x15\x7f\xae\x44\x59\xb9\x68\xdf\x21\x4a\x68\x99\x36\xc8\x29\x2d\xcf\x63\x58\xf5\x81\xee\x83\x9c\xed\x36\x56\xb5\xd0\x2a\x23\x1c\x5f\xd2\xaf\x97\x59\x8c\x3b\xfe\xc4\xa0\xa7\x1c\x88\x17\x83\x27\x56\x0b\x7e\x35\xb2\x4f\x34\xe0\x3b\x77\xd0\xe5\xfd\xf8\xd8\x58\x6f\x21\xab\xeb\xc4\xa5\xe4\x16\x64\x65\x35\xf2\x3e\x91\x34\xc1\xc7\x9b\xc4\xba\x49\xfc\x0c\x46\xa5\x96\x9a\x09\x89\x9a\xa2\xad\xf3\x07\xc8\xe9\xb9\x39\x5d\xb4\x43\xdc\x0e\xcb\x1a\x34\x86\x1d\x92\xa1\xd8\xc6\x9e\x9e\x18\xd7\x7f\x31\xf9\x13\x4a\xef\xc0\x89\x7d\xae\x7f\xda\x8a\x06\x13\x89\xd4\x41\x9d\x0c\x46\xc0\xb5\xe2\x7f\x4a\x81\xce\x7c\xa3\x06\x00\x65\xb4\x98\xbb\xf3\xa6\xe6\xf3\x17\x03\x24\xd5\x68\x36\xaa\x49\xc3\xb1\x8a\x85\xe6\x7a\xc4\xd6\x07\x30\xc7\x3d\xa3\x90\x76\x66\xa5\xda\x2c\x95\xad\x42\xc4\x3c\x4a\xf5\x2c\xc9\xab\x8b\x4d\x72\x18\xd1\xbe\xb0\xde\x40\x85\xac\xb6\xd5\x91\xb6\x56\xc8\xb4\xdd\x21\xb3\x3e\xed\x35\x96\x28\x9e\x90\xd3\x11\xa2\xf1\xd0\xd5\x4c\x83\x90\x16\x35\x75\x41\xee\x1c\xb1\x95\x0f\xf3\x70\x5f\x20\x38\x8d\xa6\x55\x92\x13\x03\xab\xdc\xed\x14\x8d\x35\x81\xc4\xfb\x3c\xbb\xd9\xbe\xff\xed\x9c\x44\x27\x47\x34\x5c\x65\x18\x10\x4b\x7f\xd7\xa7\x73\x51\xc1\x5a\xbc\x08\x84\x05\xdd\x37\x1d\x03\x61\x80\xda\x32\xc1\x63\x06\x0d\x3a\x5c\xc1\xce\x25\xb4\xfc\xde\xc2\xef\x1d\xea\xa3\xcb\x18\x0a\x7e\xa3\x1a\x0c\x1e\x0a\xa7\x93\x46\x83\xcd\xae\x46\x03\xef\xb7\xdb\xf5\xf7\x06\x7e\xfc\xe1\x87\xe0\xe8\xde\x7e\xd3\xe4\x5d\x41\x3b\x28\x77\x1b\x16\x66\xe0\x1a\xf4\xf8\xf5\x7e\xbd\x88\x1a\x50\x49\xdc\x69\x64\x8f\xe6\x9d\x03\xa8\x54\x8b\xbe\xe0\x30\xdb\x1f\x89\x51\x71\x87\x5b\x12\xd1\x1d\x2b\x1f\xe9\x00\x16\x12\x9d\xca\x1a\x4d\xd7\x50\x79\x80\xc0\xc8\x33\x09\x3c\x97\xc5\x66\x71\xb7\x5a\xe5\x8b\xad\xeb\x5c\x4e\xed\x4c\xfd\x3c\xf9\xe6\xb9\x42\x79\x6a\x68\xe1\x47\x5a\xad\x4a\x34\x86\xaa\x43\x5c\x1e\x6d\xb0\x5e\x66\x5b\xdf\x1e\x79\x5c\x7f\x2f\xf4\x7d\x40\xd4\xdc\x9b\x9d\x86\xa4\xb2\x60\x28\x5b\x99\x3c\x82\x72\x75\x6d\xdf\x69\x7f\x60\xf8\x30\x76\xf8\x68\x80\xed\x54\xe7\x4d\xf0\x1c\x0a\xa0\xb0\xe3\xd8\x54\xfa\x94\xca\xb9\x8e\x81\xcb\x33\x33\x60\xf5\x31\xc4\x9f\x17\xe0\x29\xed\x99\xa8\xb1\x8f\x1a\xba\x13\x0b\x09\x0c\x76\x8c\x27\x06\x74\x4a\xe6\xb1\xf7\x8b\x85\x62\x7c\xd7\x75\xd9\xd7\x32\x63\x6c\xa5\x55\x77\xa8\x72\x77\x3f\x98\xba\x5f\x8c\xaf\xe9\x69\x7f\x1b\x8b\x48\x92\xd6\xb1\x58\xbd\x8f\x31\x9c\xd4\x9d\xf4\x12\x9e\x5c\xbe\xfb\xd9\x0f\xa8\x8d\x38\xa9\x3b\x5e\xc2\xeb\x33\xa7\x57\xa8\x56\xa3\xb5\xc7\xc5\xf4\xe4\xf9\x13\x52\xac\x6d\x5a\xd5\xeb\x9a\x49\xec\x4b\xaa\x6b\x56\xfa\x2f\x5f\xe2\x64\xd7\xac\x14\x47\xff\x9e\x16\x06\x0a\x69\xac\xee\xe8\x1a\x80\x3c\x9d\xf4\x46\xba\x3d\xad\xae\x27\xd6\x4d\xdf\x02\xbd\x9d\xdb\x8c\x73\x8d\x26\xa9\xd2\x56\x3d\xe2\xc4\x99\x32\xdc\xa5\xdd\xd6\xb3\x7b\xa7\x70\x73\x37\x42\x3e\x26\x7b\xdf\xc2\xfd\x57\x5e\xc1\x1c\xfa\xe9\xe3\xd7\x57\xae\x9a\x67\x37\x9a\x6f\x14\x13\x5f\xba\xc2\x01\xe9\x65\xce\xcf\x58\x38\x37\xbf\xd4\x71\xf5\x98\xc1\x93\x30\x7f\xdb\xdc\xad\xfe\x0c\x89\xf4\x65\xee\x9b\x34\x05\x2a\x18\x91\x65\x9a\x48\xdf\x24\xfc\x15\xfd\x4f\xde\x0c\x43\xc4\xa6\xaa\xf7\xd7\x85\xd1\x73\xb1\x83\x01\x48\xee\x72\xee\xf3\xa6\x58\x3d\xfc\xe3\x73\x76\xbb\xfc\xe9\xff\xe3\xd0\x32\xbb\xff\x58\xac\xd2\xb1\xc5\xdd\x6a\x9b\x15\xab\xfc\xfe\xf3\x26\xdf\x7e\xfe\x2d\xbb\xbd\xd9\x4c\x4f\x4d\xe0\xa5\x0b\xb6\xf9\xed\xfa\x86\xea\x92\x07\xe9\x53\x60\x78\xcb\xf6\xff\x0f\xe8\x24\x76\x4d\xc5\xfe\xef\xc7\x9f\x12\x1d\xd3\x47\x80\x6f\x29\x6b\xd3\x4f\x08\xa3\xc7\x23\xef\xf1\xf3\xf7\x96\xf3\x8d\xa3\x37\x22\x9f\x74\xaf\xbc\xb0\xcc\xbe\xfc\x27\x00\x00\xff\xff\x23\xef\x38\x5c\x07\x19\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 6407, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
