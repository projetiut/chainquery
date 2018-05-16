// Code generated by go-bindata.
// sources:
// migration/000_init_schema.sql
// migration/001_supports.sql
// migration/002_decimals.sql
// DO NOT EDIT!

package migration

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

var _migration000_init_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xc0\xdc\xcb\x51\x53\x79\xc6\xf2\x25\xd7\x74\x3a\x79\xa0\x29\x3a\xe1\x45\xa2\x5c\x91\x4a\x93\xbe\x40\x10\x09\x49\xa8\x25\x90\x21\xc0\xd8\xee\xa7\xef\x10\xe0\x1f\xf0\x9f\x44\xc9\x4c\x2f\xf6\xf4\xd1\xc6\x12\x0b\x2c\x76\x7f\xfb\x5b\x2c\x74\x79\x09\xfe\xb2\x27\x9b\x08\x71\x0c\x16\xe1\xc5\x85\xfa\xb7\xc3\x11\xc7\x7b\x4c\xf9\x0d\xde\x10\x7a\x61\xcc\x4d\xdd\x35\x81\x63\x7c\x34\xa7\x3a\xb0\x6e\x81\x3d\x73\x81\xf9\xc5\x72\x5c\x07\x2c\xbd\x2d\x22\xf4\x5b\x8c\xa3\xa7\x25\x18\x9b\xb7\xfa\x62\xe2\x02\xe3\xa3\x3e\xd7\x0d\xd7\x9c\x03\xc7\x74\xc1\x0e\x71\x42\x47\xc0\x98\x4d\x26\xc9\x34\xf2\x4f\xb8\xc1\x14\x47\x68\x07\x3d\x02\xfe\xde\xac\xdc\xa4\x7e\x97\x65\xb9\xfa\xcd\xc4\xac\xae\x6a\xb5\x0b\xbc\xfb\xe5\x85\x76\x01\x00\x00\x4b\xe2\x2f\x81\x63\xce\x2d\x7d\x32\xbc\x90\xff\x59\x11\xce\x96\xe0\xb3\x3e\x4f\x96\xaa\x5d\x5f\x0d\xc4\xd7\xf6\x62\x32\x19\x4a\x01\xb1\xad\x87\x20\xba\x2f\xa4\xfe\x7a\x35\x38\x75\x67\xd5\x49\x03\xba\x26\xd1\x1e\x71\x12\x50\xb6\x04\x96\xed\x9a\x1f\xcc\x39\x58\xd8\x8e\xf5\xc1\x36\xc7\x55\x71\x9f\xac\xd7\xc4\x8b\x77\x3c\x31\xed\x6c\x71\x33\x31\xb5\xd1\xbb\xe1\xbb\xda\x5a\xb7\x88\x6d\x7b\x59\x26\xc8\x26\xc4\x64\xb3\xe5\x4b\x70\x63\x7d\xb0\x6c\xb7\x75\x7d\x7b\xec\x13\x44\x21\x27\x7b\xdc\x41\x36\xba\xdf\x61\x18\x05\x01\xef\xd3\xa2\x14\xed\x31\xf4\x76\x88\xec\xfb\x9f\x3a\xa0\xde\xf1\x7d\x85\x11\xfe\x4e\x82\x98\x41\xe1\x71\xf0\xb9\x47\x91\xe9\xc6\x8f\xbc\xdf\x19\xe5\x64\x8c\xfc\xe7\xf8\x96\x38\x8a\x36\xb8\x57\x53\x4a\xe5\x9d\xfc\xe4\x3b\x8e\x18\x09\x68\x57\x39\xb8\xc5\x8f\xc5\x4a\x47\xcf\x5e\x29\x8f\x10\x65\xc8\xe3\x62\x6a\xc4\xb6\x98\x2d\x81\x6b\x7e\x71\xeb\xc3\x0c\x86\x51\xe0\x61\xc6\xb0\xbf\x04\xae\x65\x7f\xb5\x6c\x57\x1b\x0d\x72\x0c\xbc\x52\xa6\x4e\xa3\x3f\xc2\x88\x27\xd2\x63\xdd\x35\x5d\x6b\x6a\x16\x71\x97\x03\xe7\x62\x3e\x37\x6d\x17\x26\xa3\x8e\xab\x4f\xef\xb2\xe8\x09\x7c\xb2\x26\x27\x7e\x0b\x66\x36\x58\xdc\x25\x1f\x34\xcd\x2b\x26\xbe\x9b\x5b\x53\x7d\xfe\x15\x7c\x32\xbf\x82\xe5\xdd\x27\x78\x23\x60\x13\x68\x09\x62\x0e\xa4\xee\x85\x6d\xfd\x63\x61\x4a\x09\xcb\x7f\x94\x22\x1f\x85\x53\x6a\x12\x79\x52\x41\x63\x66\x3b\xee\x5c\x4f\x0e\x6d\x69\x50\x0e\xdd\xc2\x56\x1f\x85\x25\x3f\xa3\x1d\xf1\xff\x60\xc9\xe1\x1a\x1f\x4d\xe3\x93\xd6\x68\x6d\xcb\x91\xdb\x9a\xcd\xc1\x1f\xce\xcc\x86\x9f\xf5\x89\x35\x6e\x14\x1d\xa4\x8a\x2d\x7b\x6c\x7e\x51\x17\x97\xa2\x97\x96\xe1\x58\x8b\x9c\x2b\xfc\x51\x53\xbd\xb3\x2e\x39\x15\x18\x97\x89\xaa\x88\x57\x97\xbd\x4b\xb1\xa0\x64\xa2\x26\x80\x68\x59\x90\x91\x39\x88\x96\xfb\x4a\x8b\xe4\x34\xf7\x07\xad\xf0\x8d\xc1\xc5\x00\x98\xf6\x07\xcb\x36\xdf\x5b\x94\x06\xe3\x9b\x52\x3e\x76\x4c\xf7\x7d\xcc\xd7\xef\xf6\xab\x37\x79\x38\xa4\x7f\xc3\x98\x12\x2f\xf0\x71\x12\x0f\xf3\xd9\x3f\xe1\xed\x6c\x3e\xd5\xdd\xf7\xc6\x6c\x7a\x37\x37\x1d\xc7\x1c\x27\xa7\x0f\x6f\x26\x33\xe3\x13\x74\xac\x7f\x99\xef\xdf\xfc\x80\x8c\xad\x1c\x70\x63\xde\x56\x80\x64\xf5\x24\xcc\x08\x93\xe1\x67\x03\x23\xa1\x61\xcc\xa1\x17\xc4\x94\x1f\xcf\xc8\x41\xcc\x4f\x90\xfe\x8e\x76\x31\x3e\x9c\xba\xd7\xb8\x2a\xd0\x84\x1f\x35\x68\x6a\x84\xd2\x06\xb9\x4e\x78\xdf\x1b\x7d\xa8\x42\x78\x66\x9f\xca\xb8\x92\x0c\x8e\x59\x30\x42\x0f\x25\xfc\xfd\xb9\x20\x54\x5d\x52\xba\xa1\x93\xd7\x55\x45\x60\x05\x35\xcb\x38\x7c\x3b\x9b\x9b\xd6\x07\x5b\x0a\xde\x96\x04\x4b\x78\x53\x0b\x91\x01\x98\x9b\xb7\xe6\xdc\xb4\x0d\x33\x27\xc6\x39\x74\x27\x1b\x1c\x9b\x13\x33\xd9\xa0\xee\x18\xfa\xd8\x54\xb6\x6c\xcf\x80\x6e\xb8\xd6\xcc\x06\xcd\xa9\xa0\x82\xf0\xd5\x84\xa0\x80\x96\x22\x99\x41\x69\xcd\x9d\x0f\x7e\x95\x42\x63\xf6\x71\xc9\xe8\x5d\x3e\x3c\x8c\xa9\x8a\xfc\x2b\x44\x56\xe4\xfb\x11\x66\xac\x1d\x55\x33\x81\x1c\x03\xde\x9c\x8e\x01\xa9\x6f\x54\xe1\x8d\x44\x8c\x43\x86\x31\x2d\x42\x23\x67\x99\x9b\x42\xe1\x6f\x57\x83\x74\x86\x62\x14\xc6\xd1\x4e\x2d\xd1\xae\x06\x3f\x2d\x08\x54\x83\x58\xcf\x0c\x7a\x88\x48\xa5\x42\x8a\x6c\x76\x0e\x99\x83\x52\x1f\x3f\x96\x64\xdd\xc4\x66\x9a\x30\x5d\xdd\x89\x53\x99\x4e\x0e\x9f\xca\xbe\x42\x67\x17\xe9\xbc\xdd\xd5\x55\xd4\x49\x06\x8f\x95\x42\x15\xd6\xd9\x67\x9a\x94\xc4\x23\x3d\xf3\xa6\xc5\x64\x72\x0c\x7a\x01\xa1\x2b\xc4\xf0\xd1\x42\x23\xbd\x65\xc8\xa4\x9f\x4d\x8e\x12\xf2\x1a\xc4\xbc\xaf\x22\x34\x9b\x8e\xd6\x33\x7f\x45\x82\x85\x98\xfa\x30\x0e\x7d\xe9\xcb\x5d\xb6\xcd\xf0\xb7\x18\x8b\x92\xfd\x0c\x5e\x96\xcd\xe1\x45\x24\xe4\x90\x91\x0d\x44\x6c\x2f\xa9\xc7\x79\x5b\x55\x66\x12\x25\xea\x59\x33\xbd\x14\xbc\xb3\x44\xd0\x1d\xa4\x2b\x42\x44\xc5\xc5\xaa\xf7\x97\x69\x0a\x2a\x23\x68\x37\x9a\x72\x40\x73\x85\x56\x55\x60\xa0\xac\x9b\xd7\x29\xd8\x29\xfa\x15\x98\x15\xaa\x3f\x4b\x6f\xd3\x52\xb7\x6b\x2e\x1b\x83\x98\xab\x05\x63\x1e\x73\x75\x69\x31\x67\x27\x90\x17\x92\xcd\x10\xdf\x2c\xdb\xc0\xe7\x6a\x00\xf8\x0a\xb2\x03\xac\x11\x22\xf1\xdf\x0e\xe9\xe0\x00\x58\x57\x44\x1b\x23\xa4\xe6\xfe\xc4\x5f\x0e\x4b\xb3\x36\x05\x4f\x7e\x3c\x19\x55\xc0\xac\x88\xb7\x6c\x9a\xb2\x07\x13\x35\x1e\xcf\x8e\x9d\x06\xc5\x75\xb2\xd2\x53\xe8\xbe\x68\xaf\x92\xb7\x02\x2f\x82\x74\xb4\xa7\xbe\x04\x73\x8e\xa7\x4e\xfe\x14\xe2\x72\xdf\xe4\x39\xd9\x31\x8c\x57\xf0\x1e\x3f\xf5\x94\x6b\xb3\xd9\xce\xcf\xb7\xf2\xce\x01\x7f\x8b\x49\x84\xfd\x24\x77\x53\xc4\xe3\x08\x37\xb4\x6a\x94\xab\x93\xd1\xef\x57\x85\x45\xde\x9e\x6b\x91\x2c\x9c\x76\x84\xf1\xe7\xac\x9e\x30\xc1\x9d\x78\x47\xce\x94\x88\xc2\xd5\x13\x6c\x85\xc0\x17\x53\x73\xcd\x64\x10\xb6\x96\x5c\x0b\x55\xa4\x16\x62\x43\xe9\xff\x2d\xe4\x45\x7e\xf8\x3f\xe4\x10\xcd\x0b\x70\x92\xd3\xba\x79\xca\xd1\xbf\x7e\x7a\x3d\xa5\x81\xea\x55\x7e\x8e\xff\xf5\x2b\xfc\xb2\xdf\xb6\x5c\xde\x97\x84\x1a\xae\xed\xe5\xf6\x8e\x12\xa5\x59\x1c\x06\x84\x8a\x1d\x89\xd3\x1a\x36\xd0\x13\x60\xcc\xa6\x53\xd3\x76\xc1\xaf\x14\x63\x1f\xfb\x60\x1d\x44\x20\xc2\x6b\x1c\x25\xb5\x01\x03\x84\x02\xbe\x25\x0c\x78\xc1\x2e\xde\x53\x10\x44\x3e\x8e\x7e\xad\xd7\xc7\xce\x74\x09\xb4\x3a\x4a\x69\xd7\x6f\xdf\x0e\xea\x3a\x08\x05\x1e\x62\x98\x81\x87\x2d\x8e\x30\xe0\x5b\x0c\x74\x67\x0a\x92\x51\x06\x78\x00\x56\x18\x84\x28\x62\xd8\x07\x0f\x84\x6f\x01\x0b\x31\xf6\xeb\x4a\x67\x71\xc8\x3b\x52\x3b\x69\xb0\x57\x58\xbe\xa7\x17\xec\x55\x86\x96\xfe\xfb\x87\x52\x34\x69\x53\x85\xe1\x14\x4a\xbb\x91\x34\x39\x41\x41\x96\x0a\xb8\x29\x66\x2a\x07\x68\x50\xc2\xac\x67\x11\xb5\xaa\xf2\xff\x33\xb5\x23\x4d\xa6\x9a\x8f\x9d\xc8\xd2\xba\x3b\xda\xd2\xc7\x2b\xc2\x21\xda\xcb\xae\xd1\x91\x6e\x4f\x01\x2e\x4e\xbc\x07\xc1\x5a\xa0\x89\x40\x72\x01\x25\x02\xbc\x52\xdd\x02\xdb\x92\x61\xfe\xf8\x6b\x91\xa7\xfd\x67\xea\x92\x4e\xd9\x45\xd9\x0e\x71\xcc\x38\xac\x37\xa6\xfa\xec\x82\x28\x8e\x5c\x39\xa1\x6e\x51\xa9\x4c\x55\x44\xc7\x9f\x95\xca\x5b\x17\xf4\xc3\xc2\xf5\x50\xaf\xa3\x50\x3f\x11\x27\xd9\xd0\xa5\x69\x3b\xe2\x83\x3d\x94\x62\xde\x71\xe2\xf9\xc9\x3c\xa5\x10\xe8\xf8\xb1\x21\x7c\x39\x4d\x84\x8a\x57\xbf\xec\x0c\x27\x1e\x2b\x75\xab\x15\x7b\xed\x73\x77\xab\xef\x28\xda\x63\xf5\x3d\xcd\xf5\x9b\xfa\xcb\x38\xf1\xda\x2a\x59\xd3\xb9\xbd\xa2\xc6\x09\x65\x65\x99\x15\x2d\xd7\x8a\x5a\x70\x79\x09\x46\xe0\x12\x18\x38\xe2\x64\x4d\x3c\xc4\xb1\xfb\x14\xe2\x21\xb8\x06\x97\xc0\xe1\x11\x46\xfb\xe4\xef\xf4\x02\x39\x5e\xed\x08\xdb\xe2\xe8\x59\x2b\xcc\x71\x51\x61\x8e\x08\x18\xc9\x42\xad\x94\xc4\x55\x56\x93\x61\x62\xa1\x9f\x91\x4d\xb9\x7d\x75\xde\xb9\x79\x85\x9e\x52\x2b\x9c\xf9\x95\xcb\x81\xd1\xd9\xd5\x78\xd7\x97\x05\x79\x63\xbf\xbf\xf7\x56\x82\xea\x43\xc4\x64\xd1\x3e\x35\xc7\xd6\x62\x2a\x8a\xdf\x36\xb9\x7f\x8b\xe2\xa3\x10\x2c\xc6\x89\x0f\x11\x87\xd9\xd3\xc5\x63\x8e\xde\x55\x0e\xaf\xd7\xd8\xe3\xe4\x3b\xee\x9a\x53\xd3\x5b\xfb\xcb\x4b\xa0\xfb\x3e\x49\x6c\x8a\x76\x60\x4d\xf0\xce\x97\x29\x14\x23\xf6\x04\x08\xf5\xf1\x23\xa1\x9b\x24\xe7\x32\xe1\xc1\x20\xf1\x7f\x96\x52\x8b\x98\x6f\x83\x48\xb9\x54\x18\x5d\x0f\x72\x2a\x21\x8b\x12\xd2\x68\x05\x2f\xa0\x3c\x29\x09\xcb\xb7\x34\xa3\xdf\xaf\xcf\x7d\x1b\xc3\x20\x85\x0c\xae\xe1\x43\xb7\xeb\x84\x1d\xa2\x9b\x18\x6d\xfa\xb8\x21\xe2\xdb\x78\xbf\xa2\x88\xec\x64\x1b\x58\x7d\x83\x47\xf8\xae\x1c\x0a\x27\x3c\xa2\x59\x63\x0c\xbd\x38\x4a\x42\xfa\x29\x85\x87\xdf\xf2\xc6\x72\x32\xd8\x47\x2f\xbc\xda\xe8\x63\x70\x4d\x76\x1c\x47\x5d\x3b\x59\x2b\xe2\x43\xc6\x45\xc0\xab\x76\xcc\xe4\x7f\xd1\x3d\x0f\x87\x1c\xfb\xbf\xfc\xe4\x4f\x0c\x9b\x98\x9d\xc0\xd0\xc3\xad\x22\x21\x72\x80\xa1\xb5\xbe\x6d\xa9\x30\xb5\x53\x5e\xb8\x1c\x58\xc9\x5d\x86\xe8\xa2\x39\xa3\xa6\x97\xb2\x76\x2f\xdb\x59\x9e\x20\x4f\xd2\x5d\x7d\x1e\x20\x74\x2f\x28\xf9\x16\xd7\x9e\xcb\x28\x06\x18\xa6\x89\x7d\xa8\xe4\xe5\x96\xf7\x98\x62\x42\x43\xcd\x26\xe9\x15\x4e\x29\xc3\xb4\xdc\xe0\xa8\x32\x83\x81\x48\xc9\xca\xbf\x04\x78\x55\x39\x9d\x51\xb7\x47\x8d\xf7\x55\x8f\xfa\x84\xc7\x41\x72\x3b\x5d\x2e\x4b\x84\x64\xc7\x3e\x98\x90\x15\xb7\x5c\x3a\x2f\xde\x91\x56\xb3\x4b\xcb\x77\x37\xc4\x77\x64\xd8\x6a\x4a\x0c\xb7\x08\xdb\x48\x6e\x55\x30\x2e\x79\xaf\x34\xbc\xb8\x68\x14\xd5\xd3\x7c\xa0\x65\x99\x41\x1b\xfd\x6d\xd4\x70\x8d\x26\x4d\x22\x73\x80\x2b\x52\x80\x56\x4e\x09\x2d\x9f\x4c\x72\xd8\xd6\x0a\x08\x6f\x3b\x2c\x89\xbe\x5a\x0a\xc3\x72\x25\x2f\x9a\x90\xb3\x38\x0c\x83\xe8\x40\xfb\x26\x15\xc0\x3e\xec\x9d\xfa\xa6\x53\x77\xe7\x16\x67\xe7\x87\xde\x3b\x4b\x5d\xcb\x8a\xc6\x2c\x90\xd9\xbc\x3d\x0f\xac\xef\x61\x6e\xf7\x1c\x59\x1b\x4e\xa2\x57\x0c\x56\x9c\x9d\x75\x09\xe4\xf2\x0a\x89\xdf\xb6\xc6\xda\x87\xbc\x3d\xbb\xb5\xb4\xfd\xa5\x9d\xb5\x52\x9b\x42\x19\x0e\x62\x9e\xdf\x8d\xb7\x36\x37\x5e\x74\x98\xc6\xf4\x9e\x06\x0f\x14\x1e\xa9\x9f\x4f\x2a\x60\x0b\x86\x5c\xff\x01\x18\x61\xe9\xf3\xa3\x8e\x9c\xad\xcf\x1f\xeb\xfc\x19\x71\x7a\xe2\x3d\x7b\x4f\x85\x5b\x15\x18\xd2\x53\xf6\xea\x2c\xb1\x4a\x67\x44\xb7\xa8\xde\x8f\xaa\xe8\x6b\xa1\x33\x65\xa9\x41\x0b\x11\x55\x17\x13\xfc\xb0\xcb\x7c\x25\x8a\x17\x42\xa1\xe4\x12\xf5\x37\xdd\x05\x32\x34\x7f\xd2\xd8\x6f\x68\x17\x77\x1f\x5f\xe9\x23\x1f\x14\x86\xbb\x84\x96\x8a\x1f\x40\x70\xc4\xe3\x43\x4f\x9f\xc3\x10\xaa\xbf\x57\xa8\x5d\xdd\x23\x8e\x72\x81\x86\x71\x14\x92\x96\xe1\x46\xf7\x96\x29\x82\x88\x2e\xdd\x4f\x60\xe0\xff\x06\x00\x00\xff\xff\xfa\xf5\x06\x12\x04\x3c\x00\x00")

func migration000_init_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration000_init_schemaSql,
		"migration/000_init_schema.sql",
	)
}

func migration000_init_schemaSql() (*asset, error) {
	bytes, err := migration000_init_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/000_init_schema.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration001_supportsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x4f\x6f\xda\x30\x18\xc6\xef\xf9\x14\xef\xad\xa0\x15\x29\x48\x1c\xa6\x4d\x1c\x8c\xf3\xd2\x5a\x4d\x1c\x64\x3b\x5b\xd9\xc5\x09\xc1\xa0\xb4\x24\x41\xb1\x73\xd8\xb7\x9f\x1c\x10\x74\xd5\x36\x69\x55\x1b\xf9\x90\xd7\x7f\x9e\xe7\xe7\xc7\xef\x64\x02\x9f\xea\x6a\xdf\x15\xce\x40\x76\x0c\x82\x97\xb5\x74\x85\x33\xb5\x69\xdc\xc2\xec\xab\x26\x20\xb1\x42\x01\x8a\x2c\x62\x84\xb6\x77\xc7\xde\x01\x89\x22\xa0\x69\x9c\x25\x1c\xf2\xf2\x50\x54\xb5\xae\xb6\x39\xd0\x7b\x22\x46\xb3\x70\x3c\xfc\x10\xea\x4f\x49\x54\x70\x28\x5c\xd5\x4c\xfd\xfe\x98\x28\x3c\x97\x7a\x6f\x1a\xd3\x15\x07\x5d\x56\x5f\xff\x6c\x8e\xcd\xf6\x2d\x58\xcb\x54\x20\xbb\xe3\xf0\x80\x6b\xc8\x77\xcf\x7a\xc0\xcb\x61\x74\xe5\x1c\x83\xc0\x25\x0a\xe4\x14\xe5\x19\xff\xd5\x7a\xca\x21\xc2\x18\x15\x02\x25\x92\x92\x08\xfd\x4c\xb6\x8a\x3c\x3e\x4f\x81\x50\xc5\x52\xfe\x0e\xd8\xb6\x3f\x1e\xdb\xee\x55\x9c\x9d\x29\x9c\xd9\xe6\xe0\xdd\x14\x4b\xbc\xa3\x02\x9e\xc5\x31\x44\xb8\x24\x59\xac\x80\x66\x42\x20\x57\xda\xaf\x4a\x45\x92\xd5\x07\xa1\xd4\xed\xb6\xda\x55\xff\xc9\xf2\x22\xaa\xf7\xe2\xa4\x02\xbd\xdc\x09\x94\x2d\x07\x08\x7c\x64\x52\x49\xc8\x9f\xda\x8d\xb6\xae\x70\xbd\xcd\x83\x51\x00\xa7\x89\xa6\xa8\x4d\x0e\xc3\xf7\x8d\x88\x37\xb6\xe5\xe5\xae\xb7\x5e\xf6\x50\x58\xa7\xed\xcf\xa6\x3c\xe9\xfe\x3d\x90\x9b\x30\x0c\xa7\x93\x61\x40\x18\x7e\x19\xc6\xcd\x20\x51\x59\x6d\xfb\xb2\x34\xd6\x0e\x1a\x8a\xf1\x35\xe3\x6a\x34\x1d\x5f\xce\x86\xbf\x7b\x9a\xae\x6b\x3b\x5d\x1b\x6b\x8b\xbd\xc9\x41\xe1\xa3\xba\x0d\x02\x80\x95\x60\x09\x11\xeb\x53\x8b\x1f\x9f\xf5\x53\xbb\x39\x47\x00\xa3\xeb\xfd\xc7\x41\x30\x46\x7e\xc7\x38\xce\x59\xd3\xb4\xd1\xe2\xfa\x66\xf7\x44\x48\x54\xf3\xde\xed\x3e\xd7\x9b\xd9\x25\x82\x73\xad\xfb\xa6\x2a\xdb\xad\xf1\x19\x88\xf4\xbb\x5e\xa6\x22\x21\x6a\x4e\xd3\x64\x25\x50\x4a\x8c\xbc\xb1\x5e\xc4\x29\x7d\xd0\x92\xfd\xc0\xf9\xec\x1f\xef\xfa\x2b\x00\x00\xff\xff\x84\xc3\x3d\xac\x68\x04\x00\x00")

func migration001_supportsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration001_supportsSql,
		"migration/001_supports.sql",
	)
}

func migration001_supportsSql() (*asset, error) {
	bytes, err := migration001_supportsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/001_supports.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration002_decimalsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8d\x41\xcb\x82\x40\x14\x45\xf7\xfe\x8a\xbb\x53\xf9\x3e\x63\x08\x02\xc1\xd5\xa8\x93\x09\xaf\x27\xd4\xbc\x75\x8a\x3c\x43\xc8\x8a\xb0\x7e\x7f\x1b\x77\x05\xdd\xc5\x5d\x9c\xc5\x39\x49\x82\xbf\x69\x3c\x3f\xba\x59\x21\xf7\xc0\x92\x77\x07\x78\x9b\x93\x43\x7f\xe9\xc6\x09\xc5\xce\x72\xe5\x50\x34\x24\x7b\x46\xab\xc3\xa0\xfd\x3c\xbe\xf4\xd4\x4d\xb7\xe7\x75\x6e\xbf\xa1\xbc\xae\x6a\xf6\xd1\xda\xc4\x10\x3e\xd6\x15\xbb\x12\xdc\x78\xb0\x10\xa1\x74\x5b\x2b\xe4\x61\x90\xfd\xce\x0d\xaa\xed\xf2\x65\x23\x39\xb9\x68\x93\xfe\xa7\xf1\xa7\x2d\x34\x2b\xb3\x2c\x44\x16\x04\xef\x00\x00\x00\xff\xff\xde\x43\xaf\xbd\xd8\x00\x00\x00")

func migration002_decimalsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration002_decimalsSql,
		"migration/002_decimals.sql",
	)
}

func migration002_decimalsSql() (*asset, error) {
	bytes, err := migration002_decimalsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/002_decimals.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"migration/000_init_schema.sql": migration000_init_schemaSql,
	"migration/001_supports.sql":    migration001_supportsSql,
	"migration/002_decimals.sql":    migration002_decimalsSql,
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
	"migration": &bintree{nil, map[string]*bintree{
		"000_init_schema.sql": &bintree{migration000_init_schemaSql, map[string]*bintree{}},
		"001_supports.sql":    &bintree{migration001_supportsSql, map[string]*bintree{}},
		"002_decimals.sql":    &bintree{migration002_decimalsSql, map[string]*bintree{}},
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
