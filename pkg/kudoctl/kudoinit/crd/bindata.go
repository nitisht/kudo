// Code generated for package crd by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
// config/crds/kudo.dev_teststeps.yaml
// config/crds/kudo.dev_testsuites.yaml
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

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xcf\x8e\xd4\x30\x0c\xc6\xef\x7d\x0a\x6b\x1f\x60\xd0\x8a\x0b\xca\x0d\xc1\x65\x2f\x0b\x62\xd1\x5e\x10\x07\x4f\xe2\x29\x66\xda\x24\xb2\x9d\x11\x08\xf1\xee\x28\x69\x3b\x3b\x53\xca\x9f\xcd\xad\xbf\xe4\x73\xbe\xcf\xa9\x31\xf3\x23\x89\x72\x8a\x0e\x30\x33\x7d\x33\x8a\xf5\x4b\x77\xc7\x57\xba\xe3\xf4\xe2\x74\xbb\x27\xc3\xdb\xee\xc8\x31\x38\x78\x53\xd4\xd2\xf8\x81\x34\x15\xf1\xf4\x96\x0e\x1c\xd9\x38\xc5\x6e\x24\xc3\x80\x86\xae\x03\xf0\x42\x58\xe1\x47\x1e\x49\x0d\xc7\xec\x20\x96\x61\xe8\x00\x22\x8e\xe4\x80\xa3\x1a\x46\x4f\xba\x3b\x96\x90\x76\x81\x4e\x9d\x66\xf2\x55\xda\x4b\x2a\xd9\xc1\x99\x4f\x12\xad\x5b\x00\x93\x85\xbb\x59\xdd\x50\x1e\x8a\xe0\x70\x51\xb2\x51\xe5\xd8\x97\x01\xe5\x89\x77\x00\xea\x53\x26\x07\xf7\xb5\x5e\x46\x4f\xa1\xb2\xb2\x97\x39\xcb\x7c\x87\x1a\x5a\x51\x07\x3f\x7e\x76\x00\x27\x1c\x38\xb4\x28\xd3\x66\xca\x14\x5f\xbf\xbf\x7b\x7c\xf9\xe0\xbf\xd0\x88\x13\x04\xc8\x92\x32\x89\xf1\x52\xa3\xae\x8b\xbe\x9e\x19\x80\x7d\xaf\x16\xd4\x84\x63\x7f\xc6\x2d\xd6\xbf\x0e\x5d\xf6\xf7\xfa\x60\xda\x7f\x25\x6f\x67\xbc\x74\x72\x59\x5b\xe6\xe6\x2c\x82\x96\x64\xc3\x65\x5d\x81\xd4\x0b\xe7\x96\x1d\xde\x5d\x9f\x6d\x77\xf0\x81\x49\x01\x41\xe8\x40\x42\xd1\x13\x58\x02\x5c\xb6\xfc\x5a\xb3\x2a\x0f\xb3\xed\xdd\x8a\x6f\x46\x6a\x31\x50\x70\x24\x23\xd1\xb5\xd3\x3f\x48\xb6\x9b\x33\xbd\xee\x7f\xb4\x07\xfb\x5e\xa8\x47\xa3\xf0\xf0\x9b\xe6\xef\x46\x07\x8c\xcf\x92\x6c\xe0\x15\x3a\x2d\x03\xba\xcc\xe2\x53\x0c\xf4\x9e\xb2\x51\xb8\x5f\x4f\xc9\xcd\xcd\xd5\x7c\xb4\x4f\x9f\x62\x68\xd3\xaa\x0e\x3e\x7d\xae\xbf\xbf\x25\xa1\x30\xbf\xd0\x04\x7f\x05\x00\x00\xff\xff\x1e\x0c\x26\x86\x10\x04\x00\x00")

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

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 1040, mode: os.FileMode(420), modTime: time.Unix(1576836957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xcf\x8e\x13\x31\x0c\xc6\xef\xf3\x14\xd6\x3e\x40\xd1\x8a\x0b\xca\x0d\xc1\x85\xcb\x82\x00\xed\x05\x71\x70\x13\x53\x4c\x27\x7f\x64\x3b\x15\xfb\xf6\x28\x99\x9d\x6e\x3b\x3b\xa5\xaa\x6f\xf9\x39\x9f\xf3\x25\x8e\xb1\xf0\x23\x89\x72\x4e\x0e\xb0\x30\xfd\x35\x4a\x6d\xa5\x9b\xfd\x3b\xdd\x70\x7e\x73\xb8\xdf\x92\xe1\xfd\xb0\xe7\x14\x1c\x7c\xa8\x6a\x39\x7e\x25\xcd\x55\x3c\x7d\xa4\x5f\x9c\xd8\x38\xa7\x21\x92\x61\x40\x43\x37\x00\x78\x21\x6c\xf0\x3b\x47\x52\xc3\x58\x1c\xa4\x3a\x8e\x03\x40\xc2\x48\x0e\x72\x21\x41\xcb\xa2\x9b\x7d\x0d\x79\x13\xe8\x30\x68\x21\xdf\xa4\x3b\xc9\xb5\x38\x38\xf2\x49\xa2\x2d\x05\x30\x59\xf8\xfc\xac\xee\xa8\x8c\x55\x70\x3c\x29\xd9\xa9\x72\xda\xd5\x11\xe5\x85\x0f\x00\xea\x73\x21\x07\x0f\xad\x5e\x41\x4f\x61\x00\x38\xe0\xc8\xa1\x7b\x9d\x4e\xc8\x85\xd2\xfb\x2f\x9f\x1e\xdf\x7e\xf3\xbf\x29\xe2\x04\x01\x8a\xb4\x3a\xc6\xb3\x91\x16\x27\x0f\x77\x64\x00\xf6\xd4\xce\x50\x13\x4e\xbb\x23\xee\xbe\xaf\x6d\x3a\x7d\xc0\xf3\x8d\x79\xfb\x87\xbc\x1d\xf1\xfc\x54\x73\xac\x99\x6b\x11\x48\xbd\x70\xb1\x85\xc3\x8b\x06\xba\xd3\xba\x25\x49\x64\xa4\x2b\x77\xbb\xa2\x0c\xf9\x56\x4d\x44\x4e\x86\x9c\x48\x74\xa9\x61\xa3\xf8\x0a\x5e\xbe\xeb\x14\x14\x91\xc7\xb5\xc4\x7f\x3c\x4c\xd1\xff\xe5\xed\xc2\xd5\xfe\x9c\xa6\x50\x04\x9f\xce\x32\x55\x5e\x39\xbc\x70\xc4\x7a\xf3\x0d\xad\xea\x95\x3f\xb2\x40\x87\x79\xbe\xe7\x51\x7e\x29\x82\xde\x53\x31\x0a\x0f\xcb\x21\xbb\xbb\x3b\x1b\xaf\xbe\xf4\x39\x85\x3e\xec\xea\xe0\xc7\xcf\x36\x51\x96\x85\xc2\x73\xd7\x27\xf8\x2f\x00\x00\xff\xff\xc6\xc7\xac\x12\x4f\x04\x00\x00")

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

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 1103, mode: os.FileMode(420), modTime: time.Unix(1576836957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\x4d\x6f\x1b\x37\x13\xbe\xeb\x57\x0c\x7c\x79\x2f\x81\x5e\x04\xbd\x14\x7b\x0b\x1c\x14\x30\x90\x3a\x46\x62\x07\x28\xda\x02\x1a\x91\xb3\xda\xa9\xb9\x24\x4b\x72\x25\x0b\x41\xfe\x7b\x31\xa4\x76\x25\xad\x56\xb6\xaa\xea\xa4\x1d\x0e\xe7\xe3\x99\x67\x86\x24\x7a\xfe\x46\x21\xb2\xb3\x15\xa0\x67\x7a\x49\x64\xe5\x2b\xce\x9f\x7f\x8e\x73\x76\xff\x5f\xbf\x5f\x52\xc2\xf7\xb3\x67\xb6\xba\x82\xdb\x2e\x26\xd7\x7e\xa1\xe8\xba\xa0\xe8\x23\xd5\x6c\x39\xb1\xb3\xb3\x96\x12\x6a\x4c\x58\xcd\x00\x54\x20\x14\xe1\x23\xb7\x14\x13\xb6\xbe\x02\xdb\x19\x33\x03\xb0\xd8\x52\x05\xce\x53\xc0\xe4\xc2\xba\x38\x8e\xf3\xe7\x4e\xbb\xb9\xa6\xf5\x2c\x7a\x52\x62\x61\x15\x5c\xe7\x2b\x18\xe4\x65\x67\x94\x25\x80\x12\xc9\xe7\x9d\x91\x5d\xf4\x79\xc5\x9b\x2e\xa0\x39\x75\x90\x17\x23\xdb\x55\x67\x30\x9c\x2c\xcf\x00\xa2\x72\x9e\x2a\xb8\x17\x27\x1e\x15\xe9\x19\xc0\x1a\x0d\xeb\x9c\x47\x71\xeb\x3c\xd9\x0f\x0f\x77\xdf\x7e\xfa\xaa\x1a\x6a\xb1\x08\x01\x7c\x10\x73\x89\xfb\xe8\xe4\x77\x00\xea\x20\x03\x48\x5b\xf1\x11\x53\x60\xbb\x1a\xc4\x39\x99\xb7\x94\x0e\xc1\x3d\x56\x74\xcb\xbf\x48\xa5\x41\xdc\xe3\xd7\xff\xa6\x82\x2b\x01\xfa\x89\x00\xcf\xfa\x97\x9f\x72\xd6\x92\x12\x38\xbe\xe6\xc5\xf1\x46\x4d\x51\x05\xf6\x19\x2f\xb8\x1d\x29\x83\x16\xa6\x50\x04\x84\x44\xad\x37\x98\x48\xef\x9c\x40\x6a\x30\x81\x42\x0b\x4b\x1a\x99\x04\xe8\x22\x69\x48\xae\x77\x2e\x7f\xd1\x02\xdb\x98\xd0\x2a\x02\x57\x43\x6a\x68\xa0\xc2\xfc\xd2\x5c\x7a\x02\x4c\x27\x3f\xc2\x34\xe3\x88\x01\x5b\x4a\x14\xe2\x78\x0b\x27\x6a\x4f\x84\xe7\x81\xef\xb1\xaa\xb1\x33\x69\x6a\x69\x04\xe4\xc7\xa2\x09\x2c\xd0\xed\xb6\x09\x33\x3b\x02\xae\xc1\xba\x7d\x64\xa2\xe2\x83\x5b\xb3\xce\xec\x9d\xfa\x2d\xb7\x19\xae\x1e\xbe\x31\x5c\x6f\x80\x76\x1a\xdd\x25\xe1\x0f\x1f\xa0\xd0\xa7\x2e\x64\x0e\x18\x67\x57\x14\x0e\x55\xa5\x94\x8d\xdb\x9c\x09\x5c\xa2\xde\x27\xba\x61\x63\x60\x49\x99\x1c\xd7\xe5\xc0\xd1\x1b\xdc\x4a\xb7\x5f\x92\xc3\x5e\x7b\x47\xd3\x42\xcb\xe5\x16\x9e\xee\xe2\x55\x01\xd8\xcb\x3c\xdf\x64\x9f\x1c\x73\xfe\x87\xdd\x12\x1b\xd7\x19\x3d\x44\xc2\x36\x6b\xf4\x9d\x75\x06\xc4\x9a\x0d\x41\xed\x02\xd0\x0b\xb6\xde\xd0\x3b\x61\xd0\xa2\x4c\xe4\xdb\xcf\x4f\xf7\x8f\x0b\xb1\x62\xa1\x8b\xc5\x0f\xc1\x1a\x03\xe3\xd2\x08\x63\xce\xd8\xc4\x3c\x71\xc0\xf0\x33\x55\xf0\x87\x2d\xf3\x07\x00\x02\x79\xc3\x0a\x63\x05\xf0\xfd\x3b\xcc\x1f\xa4\x76\x71\x9e\xbd\xc0\x8f\x1f\x37\xd7\x60\x16\xe8\xef\x8e\x03\xe9\x0b\x70\xfb\xb2\x53\xcd\xe1\x70\xcd\x14\x25\xd5\x63\x12\x71\x1c\x2c\x42\x72\xe7\x1a\x86\x86\x96\x92\x72\xa3\x31\x43\xf3\xc4\x77\xe0\x02\x6c\x1a\x4a\x0d\x85\x83\xde\x14\x86\xc4\xae\xae\xf9\xf5\xfe\x5a\x3a\x67\x08\xa7\x60\x4d\x81\x57\x2b\x3a\x19\x4d\x13\x69\x3e\x16\x4d\x60\x4d\x36\x95\x34\x73\x8e\x06\x6d\xa1\xc9\x8a\x52\x04\x7a\x21\xd5\xc9\xb8\xdd\x34\x74\xae\x8c\xa9\x91\xe1\x31\x60\xa3\x1a\xb4\x2b\x01\xad\xf0\xea\x6e\x18\xb7\x79\x30\xce\x0f\x87\xd2\xa2\xf3\x1a\x13\x2d\xce\x18\xe6\x1a\xb0\x04\xb4\xe1\xd4\x94\xa8\x84\x71\x40\x2f\x1c\x93\x60\x28\xf0\x6d\x38\x12\x70\xfa\x5f\x84\x85\x26\x6f\xdc\x76\x71\x45\x57\x9d\x9d\xdd\xfd\x12\x86\x80\xdb\xe3\xa9\x6e\xd0\x9e\xcc\xe7\x23\x88\x1f\x44\x03\x5a\xf4\xb1\xcf\x23\x47\x2f\xc7\x50\xfe\x9c\x3e\x71\x26\x82\x48\x18\x9f\x5f\x77\xf5\x89\x63\x92\x29\x28\x24\xcb\xda\x80\x6b\x64\xb3\xeb\xbf\x52\xa2\xd1\x9d\x67\xec\xfd\xaa\xb3\x68\x7c\x01\x19\x27\x73\xc5\x10\x7b\x63\xe3\xf8\x92\x32\xde\x38\x59\xc3\xab\x0a\xdc\x4f\xc4\xd7\x91\x7f\xec\xb5\xca\x29\x6b\x76\x85\x08\x54\x53\x20\x69\x75\xa9\xf7\x6f\x1f\x7e\xfd\xb4\xb7\x07\xc6\x29\xb9\xc3\x9c\x04\x39\x1a\xc5\x11\x6a\x67\xb4\x0c\x08\xab\x41\x04\x61\x6f\x56\x43\x1d\x5c\x5b\x6a\x7d\x31\x8f\x3a\xbf\x0a\xa8\x85\x14\xbf\x04\xd7\xbe\x9a\xd6\xd3\x91\x6a\x4e\x2b\x66\x76\x8d\x58\x14\xf7\x37\xb0\x62\xfd\xf4\x00\x49\xee\xbf\xf0\xef\x8a\xc2\xad\xff\xd5\xe5\x74\xfa\x2a\x9c\x30\x75\xf1\x8d\x1b\xf3\x48\xd4\xbb\x85\xfe\xd1\xb3\x37\x82\x4a\x91\x4f\xa4\xef\xc7\xef\x90\x9b\x9b\xa3\xa7\x47\xfe\x54\xce\xea\xfc\x2c\x8a\x15\xfc\xfe\xa7\xbc\x2f\x92\x0b\xa4\x7b\xc0\xb3\xf0\x9f\x00\x00\x00\xff\xff\x97\x79\x8b\x24\x79\x0d\x00\x00")

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

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 3449, mode: os.FileMode(420), modTime: time.Unix(1576836957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_teststepsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x41\x6f\x1b\xb9\x0e\xbe\xfb\x57\x10\xee\x21\x2d\x10\x8f\x5f\xfb\x1e\x1e\x1e\xe6\x56\xa4\x6f\x17\xd9\x62\x93\xa2\xc9\xf6\x52\xf4\x40\x4b\xb4\x47\x6b\x8d\x34\x2b\x51\x76\xbc\x8b\xfd\xef\x0b\x4a\x33\x8e\x1d\x4f\x9a\x06\x28\x9a\x4b\x6b\x8d\x44\x7e\x24\x3f\x51\x1f\xb1\x33\x9f\x28\x44\xe3\x5d\x0d\xd8\x19\xba\x63\x72\xf2\x2b\x56\xeb\xff\xc5\xca\xf8\xf9\xe6\xf5\x82\x18\x5f\x4f\xd6\xc6\xe9\x1a\x2e\x52\x64\xdf\x7e\xa4\xe8\x53\x50\xf4\x8e\x96\xc6\x19\x36\xde\x4d\x5a\x62\xd4\xc8\x58\x4f\x00\x54\x20\x94\xc5\x5b\xd3\x52\x64\x6c\xbb\x1a\x5c\xb2\x76\x02\x60\x71\x41\x36\xca\x1e\x00\xe5\x1d\x07\x6f\x2d\x85\x19\x7b\x6f\x07\x87\x35\x4c\x5f\x57\xff\x9a\x4e\x00\x1c\xb6\x54\x03\x53\xe4\xc8\xd4\xc5\x6a\x9d\xb4\xaf\x34\x6d\x26\xb1\x23\x25\x36\x56\xc1\xa7\xae\x86\xfd\x7a\x39\xd2\x9b\x2f\x78\x6f\x29\xf2\x0d\x53\x97\x97\x3a\x9b\x02\xda\x03\x93\x13\x80\xa8\x7c\x47\x35\x5c\xc9\xc1\x0e\x15\xe9\x09\xc0\x06\xad\xd1\x39\x82\x62\xca\x77\xe4\xde\x7e\xb8\xfc\xf4\xef\x1b\xd5\x50\x8b\x65\x11\xa0\x0b\xbe\xa3\xc0\x66\xf0\x28\x7f\x07\xe9\xdc\xaf\x01\x68\x8a\x2a\x98\x2e\x5b\x84\x33\x31\x55\xf6\x80\x96\x04\x52\x04\x6e\x08\x36\x65\x8d\x34\xc4\xec\x06\xfc\x12\xb8\x31\x11\x02\x75\x81\x22\x39\xce\x90\x0e\xcc\x82\x6c\x41\x07\x7e\xf1\x3b\x29\xae\xe0\x86\x82\x18\x81\xd8\xf8\x64\xb5\x24\x78\x43\x81\x21\x90\xf2\x2b\x67\xfe\xdc\x5b\x8e\xc0\x3e\xbb\xb4\x28\xa9\x38\xb2\x68\x1c\x53\x70\x68\x25\x09\x89\xce\x01\x9d\x86\x16\x77\x10\x48\x7c\x40\x72\x07\xd6\xf2\x96\x58\xc1\xaf\x3e\x10\x18\xb7\xf4\x35\x34\xcc\x5d\xac\xe7\xf3\x95\xe1\x81\x40\xca\xb7\x6d\x72\x86\x77\xf3\x5c\x71\xb3\x48\xec\x43\x9c\x6b\xda\x90\x9d\x63\x67\x66\x19\xa7\xe3\x4c\xba\x56\xbf\x08\x3d\xb9\xe2\xd9\x01\x30\xde\x49\x95\x22\x07\xe3\x56\x93\xfb\xb4\x5a\x62\x7a\x34\xd1\xd7\x39\x2d\x39\xd8\xb2\x13\x90\x73\xd8\x0b\x5a\x19\xe7\x8c\x5b\x95\x14\x53\x26\x04\x08\x23\xaa\x03\x5b\x86\xa9\x8d\xf5\x51\x72\xc6\x2a\xfe\xb5\xba\x8f\x80\x7a\xfb\xe1\x72\xa8\xf4\xe0\x3d\xd0\x92\x02\x39\xae\x4e\x4e\x8e\x46\x5d\xfe\x96\x86\xac\xfe\x80\xdc\x3c\xe1\xef\xec\x72\x59\x1c\x88\x0d\x49\x05\x42\x67\x48\xd1\x11\x75\xc0\xb8\xc8\x84\xba\x5f\x94\x62\x04\x3a\xb1\x0b\xfd\xee\xf3\x42\xcb\x02\xeb\x80\x6c\x8c\xc6\x01\x96\xdb\x03\xbf\xdc\x5c\x5f\xcd\x7f\xf6\x05\x27\xa0\x52\x14\xe3\x88\xc5\xc8\xc8\xd4\x92\xe3\x73\x88\x49\x35\x80\x51\xe0\x9b\x40\xfa\x46\xbe\x54\x2d\x3a\xb3\xa4\xc8\x55\x6f\x9f\x42\xfc\xfc\xe6\x4b\x05\x3f\xf9\x00\x74\x87\x6d\x67\xe9\x7c\xc4\xac\x29\x99\xed\xa3\xeb\x13\xac\x08\x4c\x2c\x29\xd8\x5b\x83\xad\xe1\x26\xc3\xee\xbc\xee\x03\xdb\x4a\x40\x23\x46\x19\xd7\x04\xbe\x0f\x31\x11\x58\xb3\xa6\x1a\xa6\xd2\x8d\x0e\xe0\xfd\x25\x3d\xe8\xef\x29\xbc\xdc\x36\x14\x08\xa6\xf2\x73\x5a\x20\x8c\x25\xa0\xbf\x89\xb2\x6b\xe0\xc3\x3d\x38\x6e\x84\xb1\xc1\xac\x56\x14\x48\xe7\x8f\x24\x57\xe5\x15\xf8\x20\x31\x3a\x3f\x62\xf1\xfe\x78\x36\x2a\x95\xea\x48\x99\xa5\x21\x7d\x02\xf6\xf3\x9b\x2f\x53\x78\x79\x9c\x8d\xb1\x74\x3a\x4d\x77\xf0\x06\x8c\x2b\x19\xea\xbc\x7e\x55\xc1\x6d\x66\xc1\xce\x31\xde\x89\x17\xd5\xf8\x48\x0e\xbc\xb3\x3b\x89\xaa\xc1\x0d\x41\xf4\x2d\xc1\x96\xac\x9d\x95\x4e\x37\x96\xd6\x2d\xee\x24\xf2\xa1\x48\xc2\x29\x84\x0e\x03\x3f\xe8\x6e\xb7\xd7\xef\xae\xeb\xe2\x5f\x28\xb2\x72\xe2\xd4\x79\x1e\x31\xb9\x34\xd2\xc1\xa4\x75\xe5\xed\x85\x83\x02\x36\x15\x42\xb0\x07\xd5\xa0\x5b\x51\x89\x88\x60\x99\x38\x05\xaa\xce\x9e\x73\x03\xf3\x03\xf3\xc4\xe5\x7b\x6f\x9c\x3e\xb9\xe6\xdf\xb9\x5f\x0a\xc4\x38\x13\x34\xf1\x59\xf8\xf3\xf3\xfa\x04\xfe\xab\x03\x5a\x7e\x15\xff\x3a\x2d\x28\x38\x62\xca\x21\x68\xaf\xa2\xa0\x57\xd4\x71\x9c\xfb\x0d\x85\x8d\xa1\xed\x7c\xeb\xc3\xda\xb8\xd5\x4c\x58\x36\x2b\x65\x8d\xf3\xfc\x64\xcf\x5f\xe4\x7f\x9e\x8d\x3f\x3f\xd9\xdf\x12\x44\xde\xf8\x23\x22\x11\x3f\x71\xfe\xac\x40\x86\x07\xef\xdb\x1e\x90\xb3\x9b\x72\x9b\xd5\xc3\x73\xc2\xea\x6d\x63\x54\x33\xa8\x86\xfb\xa6\x37\x72\x45\x5a\xd4\x74\x2e\x3d\x04\xdd\xee\x7b\x33\x52\xf2\x95\x82\x38\xdf\xcd\x7a\x95\x37\x43\xa7\xe5\xff\xd1\x44\x96\xf5\x67\x25\x28\x99\x27\x2f\xda\x6f\x97\xef\x7e\x0c\x4f\x93\x79\xd6\x35\x2b\x1f\x8a\x8d\x13\x39\x83\x21\xe0\x6e\xbf\x9a\x9b\xec\x61\x9c\x4b\x1f\x5a\xe4\x5a\x24\xd9\x7f\xff\x73\x72\x58\x84\xda\x8a\xc2\x7e\xfd\x61\x3b\x1a\x69\x44\x26\x02\x0e\xcf\x76\x79\xc2\xf6\xd2\x32\x8b\x83\x86\xe0\xe3\xff\x6f\x6e\xf7\xc4\xca\x44\x3a\xd6\x9b\xc3\x83\xda\x1f\x8b\xf7\xa2\x53\x44\xa2\x71\xcb\xfc\x6a\x49\xcf\x0d\xbe\x2d\x0f\x96\xd3\x9d\x37\xae\x28\x2f\x65\x0d\xb9\xe3\x8e\x1d\xd3\xa2\x35\x2c\x7c\xfd\x23\x89\x2e\x07\xf6\x15\x5c\xa0\x73\x9e\x61\x41\x90\x3a\x8d\x4c\xba\x82\x4b\x07\x17\xd8\x92\xbd\xc0\x48\x3f\xa8\x85\x8e\x56\x55\x18\xa4\xd8\x3e\x9a\xea\xf7\xe5\x3b\x08\x06\x74\x3a\x0b\x8e\x90\xdc\xa0\x3d\x23\xf7\x2f\xdb\xa0\x3b\xbf\xae\x38\x1f\x21\xd6\x18\x7f\x0e\x27\xb0\x51\x64\xd7\x1b\x51\x81\x9a\xb2\x6b\x4d\x4b\x4c\x96\xf7\xa7\xa4\x90\xdc\xcf\x67\xa2\x2e\xfc\xe1\xe6\x93\x29\x41\xb4\x72\x6e\xc0\xd5\x09\xa8\x07\x5c\x97\x32\xc8\x10\xf6\x28\xaa\x4b\xa7\x8d\x92\x39\xa4\x57\x3b\x42\x9e\x4c\x54\x39\x59\x94\xf9\x0c\x22\x2e\x69\x48\xa4\xdc\x4a\x9f\x18\xf0\x08\x51\x20\xb4\x39\xf7\xe5\x76\x83\xb2\x29\x32\x85\x53\x7c\x0b\xef\x2d\xe1\x30\x46\x09\xeb\x44\x6d\x0e\xf0\x66\x43\x7d\xf7\xbf\x87\x00\x64\x2e\x1c\x06\xe5\x61\x26\x16\xf5\x9a\x72\xc5\x44\xe0\x76\x4c\xfa\xea\xe1\x00\x3a\x9d\x1e\x8d\x9e\xf9\xa7\xf2\x4e\xe7\xa9\x39\xd6\xf0\xf9\x8b\x0c\xa1\xec\x03\xe9\xbe\x8f\x97\xc5\x7f\x02\x00\x00\xff\xff\x87\xb0\xbe\xfb\x98\x0f\x00\x00")

func configCrdsKudoDev_teststepsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_teststepsYaml,
		"config/crds/kudo.dev_teststeps.yaml",
	)
}

func configCrdsKudoDev_teststepsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_teststepsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_teststeps.yaml", size: 3992, mode: os.FileMode(420), modTime: time.Unix(1576836957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_testsuitesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x4d\x8f\x1b\x37\x0c\xbd\xfb\x57\x10\xdb\x43\x36\x40\xd6\x9b\x20\x45\x11\xcc\x2d\xb0\x7b\x58\x04\xcd\x2e\xb2\x9b\xf4\x50\xf4\x20\x4b\x1c\x9b\xb5\x46\x9a\x52\x94\x9b\xed\xaf\x2f\xa8\xf1\x8c\x3d\xfe\xc8\x26\x69\x7c\xf2\x70\x24\xf2\xe9\xf1\x91\x1c\x99\x96\x3e\x21\x27\x8a\xa1\x02\xd3\x12\x7e\x16\x0c\xfa\x94\xa6\xeb\x37\x69\x4a\xf1\x7a\xf3\x6a\x81\x62\x5e\x4d\xd6\x14\x5c\x05\xb3\x9c\x24\x36\x1f\x30\xc5\xcc\x16\xe7\x58\x53\x20\xa1\x18\x26\x0d\x8a\x71\x46\x4c\x35\x01\xb0\x8c\x46\x8d\x0f\xd4\x60\x12\xd3\xb4\x15\x84\xec\xfd\x04\xc0\x9b\x05\xfa\xa4\x6b\x00\x6c\x0c\xc2\xd1\x7b\xe4\x2b\x89\xd1\xf7\x01\x2b\xb8\x78\x35\x7d\x79\x31\x01\x08\xa6\xc1\x0a\x04\x93\xa4\x4c\x82\x69\xba\xce\x2e\x4e\x1d\x6e\x26\xa9\x45\xab\x4e\x96\x1c\x73\x5b\xc1\x60\xef\xf6\x6c\xfd\x77\x80\x1f\x30\xc9\xbd\x6e\x2f\xb6\xd6\x67\x36\x7e\xdf\xe9\x04\x20\xd9\xd8\x62\x05\xef\x75\x6b\x6b\x2c\xba\x09\xc0\xc6\x78\x72\xe5\x10\x9d\xb3\xd8\x62\x78\x7b\x77\xf3\xe9\xf5\xbd\x5d\x61\x63\x3a\x23\x40\xcb\xb1\x45\x16\xea\x63\xea\x6f\x8f\xd1\xc1\x06\xe0\x30\x59\xa6\xb6\x78\x84\x67\xea\xaa\x5b\x03\x4e\x39\xc4\x04\xb2\x42\xd8\x74\x36\x74\x90\x4a\x18\x88\x35\xc8\x8a\x12\x30\xb6\x8c\x09\x83\x14\x48\x7b\x6e\x41\x97\x98\x00\x71\xf1\x17\x5a\x99\xc2\x3d\xb2\x3a\x81\xb4\x8a\xd9\x3b\xe5\x78\x83\x2c\xc0\x68\xe3\x32\xd0\xbf\x83\xe7\x04\x12\x4b\x48\x6f\x94\x8b\x91\x47\x0a\x82\x1c\x8c\x57\x12\x32\xbe\x00\x13\x1c\x34\xe6\x11\x18\x35\x06\xe4\xb0\xe7\xad\x2c\x49\x53\xf8\x2d\x32\x02\x85\x3a\x56\xb0\x12\x69\x53\x75\x7d\xbd\x24\xe9\x35\x64\x63\xd3\xe4\x40\xf2\x78\x5d\x92\x4e\x8b\x2c\x91\xd3\xb5\xc3\x0d\xfa\x6b\xd3\xd2\x55\xc1\x19\xa4\xe8\xae\x71\x3f\xf1\x56\x5f\xe9\xd9\x1e\x30\x79\xd4\x2c\x25\x61\x0a\xcb\x1d\xd9\x2c\x54\x1b\x2b\x69\x4e\x7c\x96\xee\x87\x15\x82\x23\x46\x2b\x91\x1f\xf5\xe0\x31\x4b\x9b\x65\xb7\x59\x6d\x97\x36\x33\x63\x10\xf8\x27\xf2\x9a\xc2\x72\xb7\x63\x4c\x4e\x0d\x21\x0a\xa8\x02\xa9\x26\x74\xcf\xa7\x4f\x41\xb4\xec\xbe\x04\xee\xce\xc8\x4a\xe3\xcf\x3e\xcc\x0b\x0e\x0a\x49\x8c\xf7\xb0\xc0\x5a\x39\xe5\x1c\x82\xa2\x29\x8a\x7d\x32\x56\xd1\xfc\x59\xd5\xbd\xa3\xe0\x80\x12\x98\xed\xb6\x2e\x7b\x3b\x71\x95\x38\x2b\x84\x0f\xbf\xde\x3f\x40\x9f\x83\x22\xc0\xb1\xe2\x8a\xd6\x76\xdb\xd2\x4e\x76\x2a\x13\x0a\x35\x72\x27\xdb\x9a\x63\x53\x3c\x62\x70\x6d\xa4\x20\xe5\xc1\x7a\xc2\x30\x96\x5c\xca\x8b\x86\x44\x75\xfe\x77\xd6\x83\x82\xc4\x29\xcc\x4c\x50\xaa\x17\x08\xb9\x75\x46\xd0\x4d\xe1\x26\xc0\xcc\x34\xe8\x67\x26\xe1\x8f\x16\x9d\x12\x9a\xae\x94\xc1\xa7\x65\xa7\xab\x66\x31\xd4\xb4\x7c\x32\xaf\x7a\xe2\x77\x37\xef\xe7\x5a\x8b\x35\x2d\x33\x97\x12\x86\x9a\x3c\xea\xeb\x9c\xf0\xab\xb2\x3a\x8b\x41\xf0\xb3\x9c\x0d\xd7\x87\xd0\x45\x5f\xed\x37\x2f\xd0\x8a\x3f\xef\xb3\x7b\x0f\xca\xa4\x09\xae\xc8\x93\x73\x38\x94\xa6\x09\x8f\xc7\xf2\x24\xc1\x26\x55\xa3\x24\x9f\xc4\xd0\x9b\x0d\xb3\xd9\x55\x5a\x63\x02\xd5\x98\x64\x4e\x9c\xbe\xc8\x70\xc1\xd4\x57\x2a\x61\x2a\x14\x18\x2a\xb8\x7a\x27\x27\xca\x6a\x84\xeb\x6c\x89\xfd\xcf\x33\xec\x4d\xc4\x93\xf8\xef\x51\xb6\x03\x11\x22\x17\x9d\x28\x02\x28\x73\xa9\x0c\xb2\xe3\x04\x76\xa5\x37\x98\x5b\xc3\xc6\x7b\x3c\x9f\xc1\x67\xda\xfa\x1a\xf3\x99\x9a\xdc\x40\xc8\xcd\x02\xb9\xcc\x94\x9e\x15\xcd\xa6\x11\x88\xc1\x22\x5c\x3a\xac\x4d\xf6\x52\xc1\x9b\xe7\xd3\xfd\x0a\xa8\x23\x37\x46\x2a\x9d\x0b\xbf\xfc\x7c\x04\x49\xa7\xc5\x12\x79\xb0\xa7\x35\xb5\x33\x9f\x93\x20\xcf\xd1\xa3\xe0\x59\x70\x37\x35\x24\x94\x17\xe0\x62\x69\xa9\xae\xac\x2e\x3c\x34\xd1\xae\xd1\xf5\x5f\x08\xd0\x7a\x13\x50\x39\xd2\x42\x00\xdb\x39\x3f\x26\x67\x11\xa3\x47\x13\x46\x48\xbe\x1b\xc2\x30\x82\xc0\xd4\x82\xbc\x13\xc9\x36\x4d\xe3\x96\x78\x49\x4d\xeb\x55\x7e\xf7\x87\xa7\x3f\x31\x1f\x8e\x70\x8a\x61\x99\x75\x67\xbd\xd3\xa3\x9e\x85\xfb\xfb\x0a\x65\xa5\x29\xe4\x02\x57\x62\xb7\x17\x0c\xf8\x68\x8d\x07\x14\xeb\xca\xb8\xd6\xca\xe6\x80\x82\x09\xde\xde\xdd\x40\x2a\xed\x79\x84\xb8\xde\x93\xdc\x89\xb9\x72\x12\xa3\xf6\x98\xef\xc6\xb6\x9f\xbb\xef\x8b\xfe\x71\x7e\xfb\xcd\xd1\x4b\xf7\xfd\x38\xbf\xdd\xfb\xda\xfc\xd6\xe0\xf2\x54\x1f\x9a\x9f\x6e\x3f\xa5\x98\xad\x49\xd8\x57\xda\x8f\x6b\x2d\x42\x0d\xc6\x7c\x7e\x18\xdc\x6e\x90\x99\x5c\xa7\xe4\x6d\x59\xf7\x9b\xb4\xfe\x5f\xbf\x84\x84\x36\x6a\x4b\xbf\xa4\xd0\xff\x1f\x49\xf5\x5b\x8a\x5e\xa7\x36\x31\x0e\x5f\x1e\x57\xa3\xcf\xb2\xc1\xd8\x7d\x08\x0d\x8f\xbb\x19\x7a\x68\xd2\x19\xb6\xb3\x75\x43\x68\x78\xde\x1f\x0d\x83\xb1\x6f\x84\x83\xe1\xa8\x0b\x8d\xde\x1c\x9a\x0e\x0b\x70\xfc\x46\x65\x7f\x60\xf9\x38\xbf\x1d\x2c\x72\x08\x66\x4b\xb4\xde\x22\xfa\x9b\x55\x7f\x89\x4a\x62\x24\x97\xbc\x1b\x6b\xb1\x15\x74\xef\x0f\x2f\x2c\x17\x17\xa3\x9b\x4a\x79\xd4\xf4\x94\x6b\x56\xaa\xe0\x8f\x3f\xf5\xca\x22\x91\xd1\x6d\xaf\x10\x9d\xf1\xbf\x00\x00\x00\xff\xff\x95\x40\x67\x7a\xc9\x0d\x00\x00")

func configCrdsKudoDev_testsuitesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_testsuitesYaml,
		"config/crds/kudo.dev_testsuites.yaml",
	)
}

func configCrdsKudoDev_testsuitesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_testsuitesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_testsuites.yaml", size: 3529, mode: os.FileMode(420), modTime: time.Unix(1576836957, 0)}
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
	"config/crds/kudo.dev_teststeps.yaml":        configCrdsKudoDev_teststepsYaml,
	"config/crds/kudo.dev_testsuites.yaml":       configCrdsKudoDev_testsuitesYaml,
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
			"kudo.dev_teststeps.yaml":        &bintree{configCrdsKudoDev_teststepsYaml, map[string]*bintree{}},
			"kudo.dev_testsuites.yaml":       &bintree{configCrdsKudoDev_testsuitesYaml, map[string]*bintree{}},
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