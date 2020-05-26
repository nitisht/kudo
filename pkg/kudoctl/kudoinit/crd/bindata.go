// Code generated for package crd by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
package crd

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

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5b\x6f\x1b\xb9\x15\x7e\xd7\xaf\x38\xf0\x3e\x38\xc6\x5a\xa3\xc6\x05\x8a\x42\x28\x0a\xa4\x4e\xb6\x50\xbb\xb5\x8d\xc8\x4e\xb1\xc8\xa6\x00\x45\x1e\x69\x58\x73\xc8\x09\x2f\x92\xd5\xcd\xfe\xf7\xe2\x90\x33\xa3\xcb\xcc\xe8\xb2\xdd\x6c\xf9\x92\x88\x22\xcf\xe5\xe3\x77\x2e\xa4\x35\x18\x0e\x87\x03\x56\xca\x0f\x68\x9d\x34\x7a\x0c\xac\x94\xf8\xe2\x51\xd3\x27\x97\x3d\xff\xd1\x65\xd2\x8c\x96\xaf\x67\xe8\xd9\xeb\xc1\xb3\xd4\x62\x0c\xb7\xc1\x79\x53\xbc\x47\x67\x82\xe5\xf8\x16\xe7\x52\x4b\x2f\x8d\x1e\x14\xe8\x99\x60\x9e\x8d\x07\x00\x4c\x6b\xe3\x19\x4d\x3b\xfa\x08\xc0\x8d\xf6\xd6\x28\x85\x76\xb8\x40\x9d\x3d\x87\x19\xce\x82\x54\x02\x6d\xd4\x50\xeb\x5f\xfe\x2e\xbb\xc9\xfe\x30\x00\xe0\x16\xe3\xf6\x47\x59\xa0\xf3\xac\x28\xc7\xa0\x83\x52\x03\x00\xcd\x0a\x1c\x83\xd4\xce\x33\xcd\xd1\x65\xcf\x41\x98\x4c\xe0\x72\xe0\x4a\xe4\xa4\x6c\x61\x4d\x28\xc7\xd0\xcc\xa7\x2d\x95\x1d\xc9\x87\x49\xb5\x3b\x4e\x29\xe9\xfc\xdf\x77\xa6\xbf\x97\xce\xc7\xaf\x4a\x15\x2c\x53\x5b\xda\xe2\xac\x93\x7a\x11\x14\xb3\x9b\xf9\x01\x80\xe3\xa6\xc4\x31\xdc\x91\xaa\x92\x71\x14\x34\x17\x66\xb6\xc2\xa9\x52\xef\x3c\xf3\xc1\x8d\xe1\xa7\x9f\x07\x00\x4b\xa6\xa4\x88\x5e\xa6\x2f\x4d\x89\xfa\xcd\xc3\xe4\xc3\xef\xa7\x3c\xc7\x82\xa5\x49\x00\x81\x8e\x5b\x59\xc6\x75\x8d\x89\x20\x1d\xf8\x1c\x21\x2d\x85\xb9\xb1\xf1\x63\x63\x28\xbc\x79\x98\x64\x95\x80\xd2\x9a\x12\xad\x97\xb5\x11\x34\xb6\x0e\xbd\x99\xdb\x53\x75\x49\xb6\xa4\x35\x20\xe8\x98\x31\xa9\xac\x0e\x0b\x05\xb8\xa4\xdc\xcc\xc1\xe7\xd2\x81\xc5\xd2\xa2\x43\x9d\x0e\x7e\x4b\x2c\xd0\x12\xa6\xc1\xcc\xfe\x8d\xdc\x67\x30\x45\x4b\x42\xc0\xe5\x26\x28\x41\xdc\x58\xa2\xf5\x60\x91\x9b\x85\x96\xff\x69\x24\x3b\xf0\x26\xaa\x54\xcc\x63\x75\x24\xf5\x90\xda\xa3\xd5\x4c\x11\x8a\x01\xaf\x81\x69\x01\x05\x5b\x83\x45\xd2\x01\x41\x6f\x49\x8b\x4b\x5c\x06\xff\x30\x96\x20\x9a\x9b\x31\xe4\xde\x97\x6e\x3c\x1a\x2d\xa4\xaf\x69\xce\x4d\x51\x04\x2d\xfd\x7a\x14\xc9\x2a\x67\xc1\x1b\xeb\x46\x02\x97\xa8\x46\x4e\x2e\x86\xcc\xf2\x5c\x7a\xe4\x3e\x58\x1c\xb1\x52\x0e\xa3\xe1\x3a\xb2\x3c\x2b\xc4\x37\xcd\x59\x5f\x6e\x59\xea\xd7\x44\x0b\xe7\xad\xd4\x8b\x66\x3a\xb2\xb0\x17\x77\x22\x23\x9d\x2f\xab\xb6\x25\xfb\x37\xf0\xd2\x14\xa1\xf2\xfe\xdd\xf4\x11\x6a\xa5\xf1\x08\x76\x31\x8f\x68\x6f\xb6\xb9\x0d\xf0\x04\x94\xd4\x73\xb4\xe9\xe0\xe6\xd6\x14\x51\x22\x6a\x51\x1a\xa9\x7d\xfc\xc0\x95\x44\xbd\x0b\xba\x0b\xb3\x42\x7a\x3a\xe9\xcf\x01\x9d\xa7\xf3\xc9\xe0\x36\x06\x3b\xcc\x10\x42\x29\x98\x47\x91\xc1\x44\xc3\x2d\x2b\x50\xdd\x32\x87\x5f\x1d\x76\x42\xd8\x0d\x09\xd2\xe3\xc0\x6f\xe7\xa8\xdd\x85\x09\xad\x66\xba\x4e\x26\x9d\x27\x54\x07\xe1\xb4\x44\xbe\x13\x1a\x02\x9d\xb4\x44\x5f\xcf\x3c\x12\xe9\xeb\x95\xd9\x96\xa8\xae\x70\xac\xc2\xdf\x32\x6f\x6c\x47\x5c\xb6\x2c\xb8\xdf\x5d\x1b\xcd\x95\x73\x89\x44\x1a\x8b\x73\xb4\x48\x39\xc2\x1b\xe2\x50\xfa\x8a\xef\xef\xd9\x13\x5f\xf3\x25\xdb\x9b\xef\xb3\x16\x7a\x93\x48\xa7\xc1\x6f\x1e\x26\x75\xe2\x48\xf9\x02\x6b\x3b\x5b\x1a\x7b\x0f\xaf\x1e\x73\x89\x4a\x3c\x30\x9f\x1f\xd5\x7a\x39\x99\x27\x35\x31\x8c\x22\x1c\xa5\x44\x8e\x3b\xf9\x28\x26\x4d\x64\x22\x4d\x76\x88\x04\x20\xb6\x59\xac\xd6\x5f\xa7\xa0\xa9\x62\x73\x93\xc3\x3c\x93\x1a\x58\xca\xea\xf0\xb7\xe9\xfd\xdd\xe8\xaf\x26\xd9\xda\x29\x93\x71\x8e\xce\x25\xaa\x14\xa8\xfd\x35\xb8\xc0\x73\x60\xae\x66\xd1\x94\xbe\xc9\x0a\xa6\xe5\x1c\x9d\xcf\x2a\x0d\x68\xdd\xc7\x9b\x4f\x5d\x98\x01\x7c\x67\x2c\xe0\x0b\x2b\x4a\x85\xd7\x20\x13\xca\x4d\x16\xa8\x49\x21\x5d\x02\xa2\x91\x07\x2b\xe9\x73\xd9\xed\x38\x83\xd2\x88\xca\xe1\x55\x74\xd4\xb3\x67\x04\x53\x39\x1a\x10\x94\x7c\xc6\x31\x5c\x10\xcb\xb6\x4c\xfc\x89\x4a\xee\xcf\x17\x9d\x32\x5f\xad\x72\xb4\x08\x17\xb4\xe4\x22\x19\xd6\x24\x7a\x9a\xab\xf9\xb1\x31\xd0\xe7\xcc\x83\xb7\x72\xb1\x40\x8b\xdd\x68\xc6\xec\x45\x59\xe1\x0a\x8c\x25\xdf\xb5\xd9\x12\x10\xc5\xd2\x99\x55\x61\x22\x5a\x06\x7f\xbc\xf9\xd4\x63\xed\x2e\x4e\x20\xb5\xc0\x17\xb8\x01\xa9\x13\x2a\xa5\x11\x57\x19\x3c\x46\x46\xac\xb5\x67\x2f\xa4\x87\xe7\xc6\xa1\x06\xa3\xd5\xba\xdb\x5a\x03\x39\x5b\x22\x38\x53\x20\xac\x50\xa9\x61\xca\x22\x02\x56\x6c\x4d\xfe\xd7\xc7\x45\x0c\x63\x50\x32\xeb\x77\x4b\x68\xa7\xd4\xc7\xfb\xb7\xf7\xe3\x64\x15\x51\x68\xa1\xc9\x14\x4a\xcd\x73\x49\x85\x92\x2a\x64\x4a\xf7\xc4\xc9\x08\x47\x48\xe4\xf0\x06\x78\xce\xf4\x02\x3b\xc5\x46\x4f\x11\xe6\x81\x12\x70\x76\x79\x6e\xb4\xee\xd7\xba\x7a\x74\xd4\xbc\xfd\xc4\xf0\x7f\xaa\x1c\x27\xb9\x15\xdb\xd0\xa3\x6e\xdd\x6d\xf1\xf9\xa0\x5b\xd4\x10\x5b\x8d\x1e\xa3\x67\xc2\x70\x47\x4e\x71\x2c\xbd\x1b\x99\x25\xda\xa5\xc4\xd5\x68\x65\xec\xb3\xd4\x8b\x21\x11\x71\x98\x98\xe0\x46\xb1\xb9\x1d\x7d\x13\xff\xf9\x45\x5e\xc4\x76\xf5\x34\x57\xe2\xd2\xdf\xc2\x1f\xd2\xe3\x46\x67\xbb\x53\x37\x43\xa7\x56\xa5\xcb\x69\x5d\x1c\xf7\x76\x52\x48\xac\x72\xc9\xf3\xba\xb3\xdd\x64\xcf\xce\x18\x29\x98\x48\x29\x97\xe9\xf5\x57\xa7\x2d\x01\x19\x2c\xd9\xb3\x1e\x56\xf7\xaa\x21\xd3\x82\xfe\xef\xa4\xf3\x34\x7f\x36\x72\x41\x9e\x10\xa4\x4f\x93\xb7\xbf\x0d\x99\x83\x3c\x3b\x22\x3b\xbb\x38\x1a\x25\xb3\xac\x40\x8f\xb6\xd5\xc0\x30\x21\xe2\xcd\x95\xa9\x87\x03\x4d\xce\x2f\xd2\xa9\x98\x7e\xf7\x82\x3c\xf8\x63\x8d\xdc\xe5\x63\x2c\x86\xcc\x22\xf8\x95\xa1\xf4\x4f\x2d\x1c\xed\x07\xac\x05\x00\x67\x9a\xda\xeb\xa6\x02\x8e\x01\x5e\x5f\xb5\x0c\x95\x5a\x48\x8b\xdc\xab\x35\xf8\xdc\x9a\xb0\xc8\xab\x86\x3c\x96\x0e\xe0\xc6\x5a\x74\xa5\xd1\x82\x8a\x4a\x83\x4a\x9d\xde\xb7\x7b\xda\xec\xa1\xc1\xac\xa5\xa5\x60\x25\xc0\xcd\x15\xb4\x74\x39\xf4\xf1\x66\x52\x11\x64\x57\xde\x36\x1e\xf1\x13\x65\x93\xee\xc6\x0e\xfe\x99\x4b\x85\x8d\x37\xf0\xea\xf5\x55\xed\xb9\x83\x9c\x95\x25\x6a\x47\xa5\xde\xae\xc1\xcb\x02\x81\x41\x70\x68\xab\x02\xd6\xb6\x97\x6d\x5c\xbd\x06\xb6\x31\xfb\xd5\xcd\xd5\x06\xd0\x04\x78\x0c\x74\x47\x57\x24\xd1\x5c\xa8\x9d\xf4\x21\xbd\x63\xb4\x24\xaf\x72\xd4\x5b\xec\x02\x61\xd0\xe9\xcb\x4b\x5f\x99\x02\x98\x2d\x32\x52\x8f\x56\x1a\x21\x39\xcc\x18\x7f\x0e\x65\xec\xbf\x7a\x5b\x19\x8a\x0e\x2b\x45\x7d\xc3\xc3\x17\xe9\x22\xa8\xd5\xde\xb9\x54\x98\xc1\x9b\x86\xb7\x6a\x5d\xf5\x66\x26\xa2\x62\x8d\x29\xda\xa0\x1a\x4b\x04\xe2\xa8\x62\x33\x41\x65\x76\xa3\x24\xe5\x11\xc2\xc3\x06\xad\x23\x31\x14\xd3\x6e\xaf\xe6\xb7\x64\xde\x19\x8f\x63\xd8\x39\xd5\xea\xf0\xea\xdb\x50\x04\x34\xb6\x5d\xa4\xb1\x87\x7b\x6d\x4c\x63\xa7\x37\x99\xc2\xed\xd3\xfb\xf7\xef\xee\x1e\xbf\xff\xa1\x8a\x02\xba\x54\xde\xc7\x2b\xcd\xd6\x23\xc7\xd6\xa3\x12\xbc\x9a\xdc\x5e\x11\xb4\xc2\xe8\x36\xaf\x62\xe3\x96\xf0\xac\xac\xbd\xde\xee\x84\x56\x52\x29\x8a\x2f\xae\x90\x59\xd2\xf4\x8e\xf1\x7c\x2f\x06\x5b\x32\x73\x46\x81\x1a\xb4\xfc\x1c\x10\x28\x31\x3a\x53\xdf\x05\x22\x6f\xc8\xf5\x28\x62\x46\xc9\x72\xb8\xa1\x9a\xf4\x49\x21\x35\x80\x1d\x6c\xd5\xb8\x22\x71\xfb\xd9\xef\xd0\x35\xac\xac\xe2\xa9\x2b\x81\x1f\x4c\xfa\xd5\x6b\xd4\xb1\xbc\xdf\x9c\xf1\x34\xae\x07\xce\x4a\xa2\x45\xba\xf2\x36\x57\xdd\x58\x15\x8c\x52\x26\x9c\x7f\xa3\x3b\xa5\xfa\x10\xc6\xf1\x51\x84\x24\x25\xa2\xe4\x46\x09\x57\x9f\xc1\xe4\x6d\xf5\xce\x73\x0d\x52\x73\x15\x44\x97\x22\x1a\x4f\x4f\x93\xb7\x2e\x03\xf8\x0b\x72\x16\x1c\xf5\xdf\xc4\x9a\x4b\x0f\xf7\x77\xdf\xff\x40\xc9\x24\xad\xa8\x28\x42\x2a\x35\x30\x25\xd3\x6b\x54\x72\x20\xee\xee\x93\x5f\x59\xd8\xa0\x24\xb5\x47\xed\x63\x1c\xe4\xa8\x4a\x07\x05\x5d\xa1\x5c\xb0\x95\x17\xa4\x2c\x7e\x1b\x8b\x5f\xa7\x48\x61\x62\x1f\xbf\x40\x4f\x9c\x9f\xab\xf8\xca\xf2\xab\xd4\xc7\xee\xc7\x8f\x16\x2f\xba\x9f\x3f\x12\x1d\xb6\x1f\x40\xcc\xac\xca\x9f\xad\x17\x90\x13\x1e\x40\x88\xc7\xd3\x4e\x4a\x9e\x56\xa9\x77\x8c\xbc\x78\x68\xa4\xc1\xf6\xdb\x64\xbc\xb3\xa7\xe9\x58\x16\x63\x8c\xfe\xa8\xe1\x31\x47\xd7\x75\xff\xa1\xca\x9c\xae\xf4\xd1\xa5\x74\x52\xde\x32\xed\xa2\x45\x8e\xf6\x76\x8f\x2f\x70\x47\x65\xaa\x43\x66\x9d\xd3\xe0\x4b\xcf\xd6\x8d\x88\xc3\x63\xd9\x21\x3c\xee\x79\x67\xad\xb1\xf1\xd3\x9f\x86\x71\xfc\x39\x4e\x3f\x60\xca\xc0\xdb\xb2\xff\xd5\xa7\xbb\x43\xf6\x8f\x87\xcd\x5a\xf6\x98\xfd\x6d\xb2\x61\x58\xff\x3b\xfc\xf6\x74\xd9\xed\xbd\x3d\x4a\x4e\xb0\x7b\xd9\x6b\xf7\x17\xf8\x8e\x79\xa6\x00\x23\x6e\x8d\xe8\xf8\x9f\x5b\x53\x94\x0a\x7d\x17\x39\x48\xee\x97\xf6\xbb\xc1\xa1\x64\x0d\xa0\x98\xf3\x4f\xe9\xa5\x74\xf3\xd7\x8d\xce\xb0\x9f\x1b\x5b\x30\x3f\x06\x5a\x3b\xa4\x66\xa7\x73\x95\x0e\x4a\xb1\x99\xc2\x31\x78\x1b\xba\x97\x1c\x4c\xbe\x00\x05\x3a\xc7\x16\x9d\xf5\xe3\xe8\xde\xbe\x7b\xf0\xd1\x8d\x65\xce\x5c\x37\x40\x00\xd2\x63\xd1\xf3\xd5\x5e\x98\x3f\x90\x94\x53\xc2\x9c\xd6\xf5\x08\x3c\x7c\x5c\x69\x1c\x84\xe8\x24\x7f\xd3\xe8\x87\xeb\x0c\x21\xfd\x75\xbb\x1e\xbf\x7e\xfd\x3e\xd3\x40\x2c\x0f\xda\x77\xf0\x80\x3b\x5c\x98\x7a\x2c\x4f\x38\x65\xd2\x7b\x50\xe8\x29\x47\x9d\xc6\x09\x07\x9e\xc6\x49\x80\xa4\x71\xec\xf0\xcf\x16\x78\x9c\x08\x69\x9c\x4f\x87\xa3\x22\xe1\x54\xc2\x9c\xe9\x54\xef\x9d\xba\x6b\x19\xb3\x96\x75\xbf\xb0\x9e\x20\xe8\xb0\x88\x43\xd0\x7e\x8d\xe8\x3a\x02\x50\x4f\x87\xfc\xb5\x7a\xe4\xff\xad\x4b\xee\x11\xb9\xdb\x3b\x9f\xdb\x27\xf7\xd9\xb9\xd3\x3d\x9f\xdc\x29\x1f\x01\xfc\x00\x79\x76\x00\x77\x4a\x72\xac\xfe\x32\x32\x43\x40\x1d\x9f\x42\xe2\x93\xce\x2c\x78\x02\x8d\xa7\xbf\x8e\x12\x60\x69\xf1\x2c\x01\xda\xbe\xfa\x5b\x81\x96\xa8\xe2\xf0\x73\x48\x6f\x8d\x1a\xd6\xac\x50\xf1\x0f\x0a\x46\x3b\x29\xe2\xfd\xd1\xc9\x85\x96\x73\xc9\x99\xf6\xb0\x8a\x4f\x25\x51\x9d\xf4\x97\xed\xbb\x83\x36\xfb\xd6\x9f\x7e\x0d\xd8\x9b\xda\xfc\x44\xa3\xfa\x35\x48\x33\x15\x83\x64\x58\xfd\x2e\x63\xf3\x2d\x40\xba\x0a\x6c\xb5\x25\xce\x1b\x4b\x39\x35\xcd\x6c\x22\x8c\x71\x8e\xa5\x47\x71\xb7\xff\x3b\x8d\x8b\xd4\x5b\xd5\x3f\xc3\x88\x1f\xb9\xd1\xe9\x32\xe0\xc6\xf0\xf1\xd3\x20\x49\x45\xf1\xa1\x36\x86\x26\xff\x1b\x00\x00\xff\xff\x45\x41\x85\x78\xd8\x22\x00\x00")

func configCrdsKudoDev_instancesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_instancesYaml,
		"config/crds/kudo.dev_instances.yaml",
	)
}

func configCrdsKudoDev_instancesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_instancesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\xdc\x36\x0c\xbd\xfb\x57\x10\xe9\x21\x97\x8e\xa7\x69\x81\xa2\xf0\x2d\xd8\xe6\x10\xb4\x9b\x06\xd9\x60\x2f\x45\x0f\xb4\xc5\x99\x61\xd7\x96\x54\x92\x1a\x74\xf3\xeb\x0b\x49\x9e\xcf\x9d\xd9\x4d\x0b\xd4\x37\x3d\x89\x7c\xd4\xe3\x87\xd5\x2c\x16\x8b\x06\x23\xdf\x93\x28\x07\xdf\x01\x46\xa6\xbf\x8d\x7c\x5e\x69\xfb\xf0\x93\xb6\x1c\x96\xdb\x37\x3d\x19\xbe\x69\x1e\xd8\xbb\x0e\x6e\x92\x5a\x98\x3e\x91\x86\x24\x03\xfd\x4c\x2b\xf6\x6c\x1c\x7c\x33\x91\xa1\x43\xc3\xae\x01\x40\xef\x83\x61\x86\x35\x2f\x01\x86\xe0\x4d\xc2\x38\x92\x2c\xd6\xe4\xdb\x87\xd4\x53\x9f\x78\x74\x24\x85\x61\xc7\xbf\xfd\xae\xfd\xbe\xfd\xb1\x01\x18\x84\x8a\xf9\x67\x9e\x48\x0d\xa7\xd8\x81\x4f\xe3\xd8\x00\x78\x9c\xa8\x83\x10\x49\xd0\x82\x68\xfb\x90\x5c\x68\x1d\x6d\x1b\x8d\x34\x64\xb2\xb5\x84\x14\x3b\xd8\xe3\xd5\x64\x8e\xa3\xde\xe1\xb7\xd9\xba\x40\x23\xab\xfd\x72\x02\xff\xca\x6a\x65\x2b\x8e\x49\x70\x3c\x62\x2b\xa8\xb2\x5f\xa7\x11\xe5\x80\x37\x00\x3a\x84\x48\x1d\x7c\xc8\x54\x11\x07\x72\x0d\xc0\x16\x47\x76\xe5\x1a\x95\x3c\x44\xf2\x6f\x3f\xbe\xbf\xff\xe1\x6e\xd8\xd0\x84\x15\x04\x70\xa4\x83\x70\x2c\xe7\xf6\x31\x00\x2b\xd8\x86\xa0\x1e\x85\x55\x90\xb2\xdc\x31\xc2\xdb\x8f\xef\x67\xf3\x28\x19\x34\xde\x5d\x31\x7f\x47\x39\xdd\x63\x67\x44\xaf\x73\x24\xf5\x0c\xb8\x9c\x45\xaa\x84\x73\x2e\xc8\x81\x56\xea\xb0\x02\xdb\xb0\x82\x50\x14\x52\xf2\x35\xaf\x47\x6e\x21\x1f\x41\x0f\xa1\xff\x93\x06\x6b\xe1\x8e\x24\x3b\x01\xdd\x84\x34\xba\x9c\xfa\x2d\x89\x81\xd0\x10\xd6\x9e\xbf\xec\x3d\x2b\x58\x28\x94\x23\x1a\xcd\x8a\xef\x3e\xf6\x46\xe2\x71\xcc\x1a\x26\xfa\x16\xd0\x3b\x98\xf0\x11\x84\x32\x07\x24\x7f\xe4\xad\x1c\xd1\x16\x6e\x83\x10\xb0\x5f\x85\x0e\x36\x66\x51\xbb\xe5\x72\xcd\xb6\xab\xe2\x21\x4c\x53\xf2\x6c\x8f\xcb\x52\x8b\xdc\xa7\x9c\xd0\xa5\xa3\x2d\x8d\x4b\xe5\xf5\x02\x65\xd8\xb0\xd1\x60\x49\x68\x89\x91\x17\x25\x70\x5f\x8a\xb8\x9d\xdc\x37\x32\x97\xbc\xbe\x3e\x8a\xd4\x1e\x73\xd6\xd5\x84\xfd\x7a\x0f\x97\x22\xbb\xaa\x7b\xae\xb5\x9c\x5d\x9c\xcd\x6a\xfc\x07\x79\x33\x94\x55\xf9\xf4\xee\xee\x33\xec\x48\x4b\x0a\x4e\x35\x2f\x6a\x1f\xcc\xf4\x20\x7c\x16\x8a\xfd\x8a\xa4\x26\x6e\x25\x61\x2a\x1e\xc9\xbb\x18\xd8\x5b\x59\x0c\x23\x93\x3f\x15\x5d\x53\x3f\xb1\xe5\x4c\xff\x95\x48\x2d\xe7\xa7\x85\x9b\xd2\xcb\xd0\x13\xa4\xe8\xd0\xc8\xb5\xf0\xde\xc3\x0d\x4e\x34\xde\xa0\xd2\xff\x2e\x7b\x56\x58\x17\x59\xd2\x97\x85\x3f\x1e\x41\xa7\x07\xab\x5a\x7b\x78\x37\x2b\x2e\x66\x68\xd7\x82\x77\x91\x86\x93\xd6\x70\xa4\x2c\xb9\x7c\x0d\x8d\x72\xd1\x9f\xcc\x91\xeb\xdd\x78\xce\x70\xb2\x71\xe5\x2a\xa5\x8e\x52\x4f\xe2\xc9\x48\x2f\x34\xf3\x0b\x96\x2e\xfc\x5b\x9b\x09\xd9\x1b\xb2\x27\xd1\x73\x1b\x36\x9a\x9e\x80\x67\xaa\xdd\xee\xcd\x67\xbc\x27\xcd\x53\x61\x3f\xd0\x0e\xfe\xdb\x27\x9e\xae\xa9\x56\x3f\x9a\x90\xc7\x4b\x1b\x67\x21\xbc\xcb\xe7\x4a\x6b\x79\x08\x05\xc3\xb1\x1a\x03\x3a\x27\xa4\x65\xe2\xe4\x3a\xc4\xa1\x34\xc1\x45\x97\xf5\x7f\xe1\x9e\x8d\xf7\x59\x21\x0f\x4e\xbe\x22\xe6\xfc\xc3\xa8\xd3\x20\x29\x49\xb1\x82\x20\x10\x64\x8d\x9e\xbf\x94\x51\x5b\xc0\xff\x10\xc3\xc5\xca\x3f\xde\x42\x11\x7c\x6c\xce\x83\x2e\xbf\xaf\x5b\xf4\xbc\x22\xb5\xaf\xae\x9e\x24\x4f\x32\x74\xe5\xec\xe5\x86\x34\xb4\xa4\x2f\xb7\x64\x39\x76\xd2\x94\xa1\xd7\x3c\xf6\x9e\xef\xca\x0b\x9c\x67\xd0\xe1\x09\x32\xbf\x76\xf6\x50\x89\x6a\x31\xbf\x3b\x0e\xbb\x00\x95\xb7\x03\x93\x54\x2b\x49\x2d\x08\xae\x69\x46\x0e\x57\xc2\x61\xa0\x68\xe4\x3e\x9c\xbf\x43\x5e\xbd\x3a\x79\x66\x94\xe5\x10\xbc\xe3\xfa\x72\x82\xdf\xff\x68\xaa\x57\x72\xf7\xbb\x60\x32\xf8\x4f\x00\x00\x00\xff\xff\x86\x23\x1d\x6c\xb8\x09\x00\x00")

func configCrdsKudoDev_operatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorsYaml,
		"config/crds/kudo.dev_operators.yaml",
	)
}

func configCrdsKudoDev_operatorsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x3b\x5d\x6f\xe3\x36\xb6\xef\xf9\x15\x07\xd3\x87\xb9\x05\x62\xfb\x76\x2e\x70\x71\xe1\xb7\xc1\x4c\x7b\x91\xed\x34\x09\x26\x99\x2e\x16\xdd\x02\xa1\xc5\x63\x8b\x0d\x45\x72\x49\xca\x1e\x6f\xd1\xff\xbe\x38\x24\x25\x59\x8e\x28\x2b\xc9\xa4\x3b\xab\x97\xc4\x12\x79\x78\xbe\xbf\x74\x74\x36\x9b\xcd\xce\x98\x11\x3f\xa3\x75\x42\xab\x25\x30\x23\xf0\xb3\x47\x45\xbf\xdc\xfc\xfe\xff\xdc\x5c\xe8\xc5\xf6\xbb\x15\x7a\xf6\xdd\xd9\xbd\x50\x7c\x09\xef\x6a\xe7\x75\xf5\x11\x9d\xae\x6d\x81\xef\x71\x2d\x94\xf0\x42\xab\xb3\x0a\x3d\xe3\xcc\xb3\xe5\x19\x00\x53\x4a\x7b\x46\xb7\x1d\xfd\x04\x28\xb4\xf2\x56\x4b\x89\x76\xb6\x41\x35\xbf\xaf\x57\xb8\xaa\x85\xe4\x68\xc3\x09\xcd\xf9\xdb\xff\x9e\xbf\x99\xff\xef\x19\x40\x61\x31\x6c\xbf\x15\x15\x3a\xcf\x2a\xb3\x04\x55\x4b\x79\x06\xa0\x58\x85\x4b\xd0\x06\x2d\xf3\xda\xa6\x9d\x6e\x7e\x5f\x73\x3d\xe7\xb8\x3d\x73\x06\x0b\x3a\x73\x63\x75\x6d\x96\xd0\xde\x8f\x3b\x13\x3a\x91\x94\xab\x04\x24\x91\x1f\x9e\x48\xe1\xfc\x8f\x43\x4f\x3f\x08\xe7\xc3\x0a\x23\x6b\xcb\xe4\x43\x14\xc2\x43\x27\xd4\xa6\x96\xcc\x3e\x78\x7c\x06\xe0\x0a\x6d\x70\x09\x97\x84\x86\x61\x05\xf2\x33\x80\x2d\x93\x82\x07\x4a\x23\x62\xda\xa0\x7a\x7b\x7d\xf1\xf3\xff\xdc\x14\x25\x56\x2c\xde\x04\xe0\xe8\x0a\x2b\x4c\x58\x77\x8c\x18\x08\x07\xbe\x44\x88\x3b\x60\xad\x6d\xf8\x79\x8c\x1e\xbc\xbd\xbe\x98\x27\x70\xc6\xd2\x53\x2f\x1a\x76\xd0\x75\xa0\x06\xed\xbd\xa3\x83\x5f\x13\x66\xe9\x50\x4e\x82\xc7\x78\x72\x3a\x02\x39\xb8\x88\x83\x5e\x83\x2f\x85\x03\x8b\xc6\xa2\x43\x15\x55\xe1\x00\x2c\xd0\x12\xa6\x40\xaf\x7e\xc3\xc2\xcf\xe1\x06\x03\x9e\xe0\x4a\x5d\x4b\x4e\xda\xb2\x45\xeb\xc1\x62\xa1\x37\x4a\xfc\xb3\x85\xec\xc0\xeb\x70\xa4\x64\x1e\x93\x3c\x9a\x4b\x28\x8f\x56\x31\x49\x3c\xad\xf1\x1c\x98\xe2\x50\xb1\x3d\x58\xa4\x33\xa0\x56\x07\xd0\xc2\x12\x37\x87\x9f\xb4\x45\x10\x6a\xad\x97\x50\x7a\x6f\xdc\x72\xb1\xd8\x08\xdf\x28\x7e\xa1\xab\xaa\x56\xc2\xef\x17\x41\x7d\xc5\xaa\xf6\xda\xba\x05\xc7\x2d\xca\x85\x13\x9b\x19\xb3\x45\x29\x3c\x16\xbe\xb6\xb8\x60\x46\xcc\x02\xe2\x2a\xe8\xfd\xbc\xe2\xdf\xd8\x64\x25\xee\xf5\x01\xa6\x7e\x4f\x5a\xe0\xbc\x15\x6a\xd3\xde\x0e\x0a\x99\xe5\x3b\x29\x24\x89\x99\xa5\x6d\x11\xff\x8e\xbd\x74\x8b\xb8\xf2\xf1\xfb\x9b\x5b\x68\x0e\x0d\x22\xe8\xf3\x3c\x70\xbb\xdb\xe6\x3a\xc6\x13\xa3\x84\x5a\xa3\x8d\x82\x5b\x5b\x5d\x05\x88\xa8\xb8\xd1\x42\xf9\xf0\xa3\x90\x02\x55\x9f\xe9\xae\x5e\x55\xc2\x93\xa4\xff\x51\xa3\xf3\x24\x9f\x39\xbc\x0b\xe6\x0f\x2b\x84\xda\x70\xe6\x91\xcf\xe1\x42\xc1\x3b\x56\xa1\x7c\xc7\x1c\xbe\x38\xdb\x89\xc3\x6e\x46\x2c\x3d\xcd\xf8\x43\xaf\xd5\x5f\x18\xb9\xd5\xde\x6e\xfc\xca\xa0\x84\x8e\x4c\xf2\xc6\x60\xd1\xb3\x10\x8e\x4e\x58\xd2\x62\xcf\x3c\x92\xee\x1f\x6d\x98\x1f\x00\x1e\x32\xce\x68\xa0\x66\xc0\x40\xb3\x84\x41\xf4\xba\x0a\x0b\x42\xf1\x26\x3c\x3c\xde\xd8\xa3\xe1\xdd\xd1\xe2\x96\x00\x06\x1e\x2b\x43\x16\xc7\x1b\xfd\xf3\x25\xf3\x50\x30\x05\x2b\x3c\x02\x09\x50\x3b\xe4\x64\xa6\xe9\x70\xfa\x97\x29\x10\xca\x79\xa6\x0a\x8c\xbe\x01\x5b\x06\xcc\xa7\xd2\xd2\xf8\xb3\x51\x1a\xae\x82\xcc\x3e\xe2\x1a\x2d\xd2\x61\xa4\x40\x4c\x28\x07\xa8\x74\xbd\x29\x83\xce\xd9\x2a\x78\x23\xc2\x4b\xa2\x87\xbd\xae\x1f\x90\x20\x14\x49\xdb\x83\xb6\x50\x69\x2e\xd6\xfb\x80\xb2\x25\xb0\x24\xc5\xe4\xb5\x8e\xb6\xe5\xe4\x06\x59\xe7\x3a\x48\xc2\xdb\xeb\x8b\xc6\xa1\x36\xbc\xb2\x91\x9e\x07\x27\x8e\xf2\x8b\xae\xb5\x40\xc9\xaf\x99\x2f\x4f\x9e\xfa\xfa\x62\x9d\xe8\x0b\xe2\xd5\xc0\xc0\x08\x8c\xe2\x6a\xfd\x74\x10\x22\x32\x1e\x6f\x0e\x80\x04\x20\x2b\xb4\x98\xd6\x9f\x47\x67\x92\x74\xa6\xf3\xed\x24\x12\x60\x31\xf6\xc1\x5f\x6e\xae\x2e\x17\xff\xaf\x23\xae\x83\x30\x59\x51\xa0\x73\xd1\x76\x2a\x54\xfe\x1c\x5c\x5d\x94\xc0\x5c\x63\x56\x37\xf4\x64\x5e\x31\x25\xd6\xe8\xfc\x3c\x9d\x80\xd6\xfd\xf2\xe6\xd7\x21\x9e\x01\xfc\xa0\x2d\xe0\x67\x56\x19\x89\xe7\x20\x22\x97\x5b\xef\xd8\x28\x8f\x70\x91\x11\x2d\x3c\xd8\x09\x5f\x8a\x61\xc2\x19\x18\xcd\x13\xc1\xbb\x40\xa8\x67\xf7\x08\x3a\x11\x5a\x23\x48\x71\x8f\x4b\x78\x45\x9a\x75\x80\xe2\xef\x94\x95\xfc\xf1\x6a\x10\xe6\x7f\xed\x4a\xb4\x08\xaf\x68\xc9\xab\x88\x58\x1b\x00\xe9\x5e\xa3\x1f\x1d\x82\xc1\x2e\xbd\x15\x9b\x0d\x5a\x1c\xe6\x66\xf0\xea\xe4\x2d\xbf\x25\xf5\x16\x6b\x50\xfa\x00\x40\x00\x4b\x32\x33\x58\x88\xb5\x40\xfe\x00\xe1\x5f\xde\xfc\x9a\xc1\xb6\xcf\x27\x10\x8a\xe3\x67\x78\x03\x42\x45\xae\x18\xcd\xbf\x9d\xc3\x6d\xd0\x88\xbd\xf2\xec\x33\x9d\x53\x94\xda\xa1\x02\xad\xe4\x7e\x18\x5b\x0d\x25\xdb\x22\x38\x5d\x21\xec\x50\xca\x59\xf4\x4a\x1c\x76\x6c\x4f\xf4\x37\xe2\x22\x0d\x63\x60\x98\xf5\xfd\xd4\x62\x10\xea\xed\xd5\xfb\xab\x65\xc4\x8a\x54\x68\x13\x72\x28\x0a\x59\x6b\x41\x09\x04\x65\x0e\x31\x0c\x92\x4e\x06\x76\xd4\x51\x39\xc8\xad\x95\x4c\x6d\x1e\x7a\x3d\x08\x6e\x23\x70\x77\x5d\x53\x60\x9a\xbf\x7e\xac\xb5\x1e\xe7\x00\xcd\x35\x90\x0b\x1c\x3b\x86\x7f\x53\x44\x9d\x44\x56\x48\xd8\x4f\x92\x75\x79\xa0\xcf\xa3\x64\x51\xe9\x60\x15\x7a\x0c\x94\x71\x5d\x38\x22\xaa\x40\xe3\xdd\x42\x6f\xd1\x6e\x05\xee\x16\x3b\x6d\xef\x85\xda\xcc\x48\x11\x67\x51\x13\xdc\x22\xe4\xff\x8b\x6f\xc2\x9f\x27\x51\x11\xb2\xf6\x69\xa4\x84\xa5\x7f\x06\x3d\x74\x8e\x5b\x3c\x9a\x9c\x26\x49\x9c\x1a\x95\x5e\xdf\x44\x87\x50\x1c\xef\x24\x93\xd8\x95\xa2\x28\x9b\x8c\xbf\xf3\x9e\x83\x36\x52\x31\x1e\x5d\x2e\x53\xfb\x17\x57\x5b\x62\x64\x6d\x09\x9f\xfd\x2c\x55\xa0\x33\xa6\x38\xfd\xef\x84\xf3\x74\xff\xd1\x9c\xab\xc5\x04\x23\xfd\x74\xf1\xfe\xcf\x51\xe6\x5a\x3c\xda\x22\x07\xb3\x5b\xba\x0c\xb3\xac\x42\x8f\xf6\x41\x02\x23\x3c\x56\x03\x59\x4d\x8f\xe6\xeb\x66\x37\x14\xcc\x90\x40\x52\x6d\xc8\xac\x60\x2b\x21\x85\xdf\x27\xc7\x3c\x54\x75\xf7\xaf\x15\x92\x37\x8f\x19\xa3\x17\x21\xef\xa4\x84\xa1\x4b\x22\x1f\x7a\xf5\xb1\xe4\x8b\x10\x5d\xb3\x5a\xfa\xa1\x47\x47\x54\xbc\x8f\x2b\x63\xb1\x95\xb6\xa5\xf8\x1d\x43\x65\xcb\x24\x5a\x62\xac\xde\x0a\x9e\x09\xb4\x00\xab\x98\x37\xe6\xb1\x86\x53\xca\xd6\xc7\x6e\x0a\xfa\xed\x8f\x4e\x0c\x0c\xa4\x56\x1b\xb4\x87\x4b\x49\x16\xa5\xde\x65\x10\x27\xac\x3b\x42\x77\x42\xca\x50\xcc\x39\xe4\x4f\xa3\x41\x38\x23\xd9\xfe\x32\x13\x08\x8e\x69\xe8\x56\xa7\x12\x23\x96\x14\xab\x3d\x7c\xba\x70\x4f\x42\x20\x17\x82\x8e\x4e\x7e\x75\x99\xb2\x1f\xa2\xff\xb0\xd2\x49\xa9\x6b\x83\x49\x8a\xf3\x4d\x55\x94\x61\xe2\x5a\x48\x0c\xdd\x98\xc3\x44\xf3\x2e\xb6\xaf\xde\x5d\x7d\xba\xbc\xbd\x23\x28\x0a\x6a\xd7\x94\xef\xd1\x56\x24\x69\x4c\x06\x26\x0b\x89\x59\x4a\x25\xff\xae\x62\x51\x1a\xdc\xb9\x91\xa2\x60\x6e\x09\xf0\xfb\xef\x30\x0f\xb6\xe8\xe6\xe1\x14\xf8\x23\x93\x5d\x9e\xe0\x19\x55\xf4\x94\x5c\x4f\xe0\xdb\xc7\xb4\xb4\xcd\x1a\x5d\x93\x53\xf7\xac\xa5\x81\x08\x5e\xe7\x0c\x06\x5b\x93\x22\x71\x33\x29\x5b\xe3\x71\xe7\x94\xae\xee\x4a\xf4\x25\xda\x03\xdb\x24\x0d\x71\xf5\x7a\x2d\xc6\xed\x6b\xa5\xb5\xc4\xc1\x9a\x25\x65\xcb\x13\xc8\xbc\x8d\x2b\x41\x70\x0a\x31\x81\xcc\x40\xa3\x64\x2a\xaa\xc9\x06\xbd\x03\xfc\x8c\x45\x4d\x2e\x6b\x57\x62\x4e\x8c\x31\x1f\xee\x1c\x66\x48\x29\x5d\xa3\x57\x17\x6d\xa9\x9c\xba\x63\x07\x4e\xe9\x2e\x76\x54\xee\x32\x80\x29\xae\x46\x84\x42\x0a\x1e\xb0\x0a\x29\x3d\x7e\x16\xce\x13\x0f\x89\x7d\x3b\xe1\x10\x84\x7f\xed\xe0\x8e\xa3\x91\x7a\x7f\xf7\x24\xab\x0a\x6e\x71\x16\x16\x4d\x60\xde\xde\xe0\x81\x7e\x44\x75\x27\xb7\x4a\xfb\x5b\x12\x43\x79\x73\x17\x4f\x7c\x0a\x52\xd9\xd8\xd6\x3c\x62\xd6\xb2\x7e\xa5\x41\xdc\x7a\x10\x34\x18\xe7\xa1\xa9\xcd\xe4\xf5\x48\x60\xe9\xc7\x3f\xe2\x7a\x47\x20\x03\x87\x96\xfe\xd1\x6b\xb8\x2e\x99\x0b\x34\x93\x34\x30\x76\x47\x56\x54\xb6\x91\x5b\xf0\x43\x4e\x75\x3c\x9c\x99\x00\x6f\x02\xd3\xd3\xc1\x15\x33\x84\x50\xd8\x16\xd5\x21\xd4\xb5\xe1\xe9\x68\x9d\x94\x89\xfb\xb9\x93\x7a\xe4\x4b\xe1\x42\x29\xe6\x3c\x9a\x44\x7b\x53\xfa\xff\xd8\x66\x3d\x19\xd0\x4d\xab\x32\xe3\xed\x4f\xf1\x27\x5e\x79\xa7\x1f\xaf\x13\xda\x1d\xaf\x80\xfd\x18\x94\x1e\x17\x6e\x02\xad\x89\xdd\xb4\xf5\x80\xdb\x0d\x3f\xda\x5e\xda\x08\x50\x38\x60\x51\xc3\x0a\x70\x5e\x93\xf3\x64\x5d\x1b\x38\xc7\x1d\x38\x25\xba\x0c\xea\x07\x1d\x3f\xd7\xa4\xfb\x0e\x03\xd6\xb1\xfb\xd6\xbe\xe6\xc8\x5f\x41\xd0\xba\x28\xea\x07\x9d\xbd\xfe\x35\x45\x82\xf1\x3a\x25\xc7\x74\xee\x14\x69\xa6\xa5\xcc\xdd\x9f\x3c\x75\x12\x07\x1f\x7d\x74\xde\x0d\x0d\xaf\x1b\xf4\x64\x8f\x05\xe7\xbc\x65\x1e\x37\xfb\xc9\x6a\x7c\x65\x39\xc6\x96\x5d\x6b\xcf\xa5\xde\xc5\xac\xa8\x5e\x05\xbe\x34\x5d\x9d\x71\x19\x4b\xa6\x16\xd1\xeb\x74\x19\x54\x78\x0b\xc8\x41\xd7\x19\x9f\x73\x48\xd7\x28\x4f\x4f\x72\x48\xd5\x52\x52\x3a\xb5\x04\x6f\xeb\xe1\x2c\x6d\x9c\x7d\xe3\x8c\x7b\x2a\xcb\x0e\xd8\x92\xa1\x6c\x3a\xb3\x9e\x1a\x0c\x1f\x44\xae\x2e\x48\x50\x18\xeb\xbc\x16\xfd\x3c\x3e\x7a\x94\xaf\xd9\x43\x07\xad\xae\x87\xc7\x87\xe4\x21\x29\xed\x0b\xab\x81\x6d\x99\x90\x29\x23\x8e\xbc\x1b\x79\x6f\x02\x13\x0b\xd5\x5b\xe6\xee\x63\x7d\xb7\x91\x7a\xc5\xe4\x39\x18\x2d\xf7\x95\xb6\xa6\x14\x05\x08\x8a\xc9\x55\xf3\xca\xb2\x41\xc7\xd4\x2b\x29\x8a\xc1\x1e\x65\x87\x63\xc0\xf9\x91\xa1\x3c\xd7\xf3\x7b\x46\x49\x73\x62\xe3\xf1\x7b\xac\x11\x2e\x85\xd7\x58\x58\xad\x90\xbb\xc8\x05\xed\x9c\x68\x28\x0d\x80\x5c\x6a\xe8\x32\x29\xf5\x2e\xe7\x0c\xea\xd8\x47\xdf\x6a\xc1\x61\x67\x45\x78\x5b\x59\x84\x29\x02\xa8\xd5\xa2\x62\xd6\x95\x4c\xca\xd0\xdb\xa6\xe0\x11\xbb\xe7\x5a\xc9\x3d\xa5\xc8\x59\x23\x29\xd0\x86\x64\x22\xf4\x68\x1d\x70\x34\xa8\x38\x81\xd6\xa9\x3a\x23\x1c\x7f\x14\x8a\x13\x8a\x08\x5c\xef\x94\x13\x1c\x9b\x77\xd3\xb9\x02\xcb\x18\xab\x59\x51\x82\x70\xe7\x11\x9d\x40\x3f\x15\x24\xa1\x07\x1a\xea\x0d\xa5\x7d\xec\x4a\xa7\xb3\x53\xae\x9d\x35\x67\xb2\xa6\xdf\x9c\x8e\x76\xe5\x28\x82\x8b\x86\xcc\x15\x16\xba\x42\x60\xd5\x4a\x6c\x6a\x5d\xbb\xf6\xf5\x7d\xaa\x6f\x72\xf9\x0f\x31\xc6\xce\xe1\xaf\x08\x95\xd8\x94\x1e\x2c\x6e\x85\x13\x3e\x1a\x49\x47\xc4\x61\x43\x3a\xb9\x95\xb1\x92\xa4\xc1\x46\x81\x70\xae\xce\x14\x54\xa7\x23\x77\xfe\x4d\x65\x77\xf5\x94\xed\x20\xdf\x60\xc6\xb4\xaf\xbf\x12\xfa\x9a\xaa\x3b\xc1\x24\x55\xb9\xfa\x7c\x34\x88\x75\x45\x04\x6d\xac\xb4\x0b\x33\x04\xc7\x2f\xad\x7b\x74\x9f\x0e\x34\x5c\xab\x91\x04\xe4\x54\x7d\x49\xd7\x9a\x79\x26\x9f\x07\xa2\xa9\x83\xf3\x4d\x94\x89\xd4\xe8\xbe\x07\x7d\xbc\x84\x1a\x00\x2f\x21\x26\xd0\x2a\x67\x45\x93\x88\x33\xac\xb8\x67\x9b\x11\x06\xf5\x88\x42\x11\x9a\x09\x84\x6f\xb3\x33\x98\xe8\x39\x48\x5d\x30\xd9\xde\x5b\x6b\xc9\xd1\x8e\x50\xa4\x2d\xd9\xcc\xa7\x8f\x1f\x88\xa4\x66\x97\x67\x76\xc5\xc2\xc8\xd2\x33\xc8\x49\x2d\x82\x67\x49\xdc\x88\x5c\x85\x0e\x53\xd2\xdc\x7e\xae\x20\x0c\xa6\xe9\x06\xba\xbb\x4a\x79\x0e\x8b\xdd\xae\x0d\x2a\xd2\x8d\xd0\xbd\x19\xcd\x0c\x59\x00\xd4\x44\xe2\xa6\xb2\x51\xbc\x7b\x7f\x30\x56\x1e\x4d\xad\x1d\x50\x6d\x7f\x10\xf2\x64\xf1\x30\x39\x7f\x5f\x7f\x49\x60\xf7\x38\x9a\x90\x3f\x0e\x56\x36\x8b\x78\x34\xb0\x49\x75\xc7\x84\xcc\x1a\x26\x15\x27\x46\x8f\xe0\x3d\x01\xe3\x76\xc0\xea\x19\xfa\x3d\x89\x31\x5f\x8c\xe2\x1d\x53\xfe\x7b\x7b\xd2\xa0\xc7\x82\xc1\xa8\x88\x9e\xd0\x01\x6b\x3a\x12\x4f\xec\x82\x8d\xf0\xaf\x9f\x51\x36\xc7\xc4\xe4\xbb\xe9\x87\xb4\x16\x1f\xe2\xc1\xdf\xde\xfe\xf4\xa1\x43\x28\xf8\xe1\xa1\x56\xc9\x51\x07\xde\x25\x17\x1d\x5c\x08\xdd\xb0\x07\x8e\x24\x0d\xb3\x51\x72\x3e\x3c\x6c\x34\xc0\xac\xda\x6c\x2c\xe3\x24\xf0\x1f\xac\xae\x46\xab\x96\x4f\xbd\xa5\x81\xac\x98\x2d\x1f\x95\x2a\xae\x1b\x9a\x8a\xd0\x1f\x6a\x51\x88\x87\x5f\xa6\xc8\xf9\x42\x83\x50\x4f\x1c\x85\x3a\xe5\xa2\xc7\xc7\xa1\x9e\x37\x10\x75\xd2\xa0\x47\x87\xa2\x9e\x39\x16\x95\x4b\x88\xd5\xd3\x07\xa3\x72\x29\x65\x18\x4d\x79\x81\xd1\xa8\x97\x18\x8e\x7a\x99\xf1\xa8\x17\x1a\x90\x7a\xc1\x11\xa9\x67\x0d\x49\xe5\xf2\x79\x2a\x94\x9f\x36\x26\x95\xd7\xd6\xe6\xcd\xd3\x63\x07\xa5\x72\x15\xf0\xf1\xf8\xd4\x84\x51\xa9\x93\x76\x9c\x4f\x7a\xbe\xfe\x81\xa9\x2f\xf5\xc6\xfa\x6b\x18\x9b\x9a\x44\x4b\x76\x74\xea\x2b\x1d\x9e\x9a\xf0\x7a\xfc\xe4\x00\xd5\x73\x47\xa8\x72\x9d\x00\xf7\x1f\x31\x44\x75\x92\x83\x99\x41\xaa\xaf\x6e\x94\xea\xcb\xbf\x72\xde\x3e\x6a\x8c\x7f\xf8\x6b\x04\xcf\x7c\xed\x26\x7f\x8f\x10\x56\xf7\xbe\x48\xd0\x2b\x87\x76\x3b\xf1\x93\x84\x01\x14\x8e\x6e\x75\x5f\x70\xa5\x8f\xc5\xda\x5b\x01\xc9\x59\xfa\x6c\xab\x7b\x0a\x10\xcf\x3f\x28\xa8\x9c\xd7\x96\x6d\x9a\x12\xab\xa3\x90\x92\x1c\xe3\x91\x5f\x1e\x7f\xbf\xf5\x2a\x46\xd9\xe6\x83\xac\xf0\xb3\xd0\x2a\x16\x2d\x6e\x09\xbf\xfc\x7a\x06\xa9\xb9\xd0\x24\xe1\xe1\xe6\xbf\x02\x00\x00\xff\xff\x10\x03\xb6\xf7\xf7\x36\x00\x00")

func configCrdsKudoDev_operatorversionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorversionsYaml,
		"config/crds/kudo.dev_operatorversions.yaml",
	)
}

func configCrdsKudoDev_operatorversionsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorversionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"config/crds/kudo.dev_instances.yaml":        configCrdsKudoDev_instancesYaml,
	"config/crds/kudo.dev_operators.yaml":        configCrdsKudoDev_operatorsYaml,
	"config/crds/kudo.dev_operatorversions.yaml": configCrdsKudoDev_operatorversionsYaml,
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
	"config": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"kudo.dev_instances.yaml":        &bintree{configCrdsKudoDev_instancesYaml, map[string]*bintree{}},
			"kudo.dev_operators.yaml":        &bintree{configCrdsKudoDev_operatorsYaml, map[string]*bintree{}},
			"kudo.dev_operatorversions.yaml": &bintree{configCrdsKudoDev_operatorversionsYaml, map[string]*bintree{}},
		}},
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
