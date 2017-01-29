// Code generated by go-bindata.
// sources:
// assets/Dockerfile
// assets/Dockerfile-go
// assets/Dockerfile-python
// assets/entrypoint
// assets/entrypoint-go
// assets/entrypoint-python
// DO NOT EDIT!

package command

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

var _assetsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x51\x6f\xd3\x3e\x14\xc5\xdf\xfd\x29\xae\xd2\xfd\xa7\x3f\x13\x8d\x3b\x1e\x78\xe8\xe8\x44\x47\x3b\x11\xa0\x49\x94\x05\xc4\xb4\x4e\x93\x97\xdc\x36\x66\x89\x6d\x6c\x67\x5b\x15\x85\xcf\x8e\xe2\x94\xb5\x0c\xc4\x43\x14\x5f\xfb\x77\xee\x3d\xc7\xf2\x80\x0c\x60\x26\xb3\x3b\xd4\x2b\x5e\x22\xe9\xca\x77\x52\x6d\x34\x5f\x17\x16\xfe\xcf\x5e\xc0\xab\xd1\xf1\x6b\xf8\x50\x0b\x85\x1c\x3e\xb2\x07\x56\x49\x2b\x1d\x96\x16\xdc\x80\x91\x2b\xfb\xc0\x34\x02\x37\xa0\xb1\x44\x66\x30\x87\x5a\xe4\xa8\xc1\x16\x08\x8b\x20\x85\x4f\x3c\x43\x61\xd0\x77\xa2\xc2\x5a\x35\xa6\x54\x2a\x14\x46\xd6\x3a\x43\x5f\xea\x35\x2d\x7b\xc4\xd0\x8a\xdb\xe1\xb6\xf0\x55\xa1\xc8\x80\x34\x4d\x8e\x2b\x2e\x10\xbc\x5b\x66\xd0\x6b\x5b\x72\x9e\x44\x0b\x68\x1a\xff\x8c\x19\x0c\x2a\xb6\xc6\xb6\x25\x8b\x69\x10\xa6\xd3\x20\x9c\x27\xcf\xad\xc2\x9b\xbb\xed\xca\xff\xe6\x4e\xde\xae\x2b\xc6\x4b\x3f\x93\xd5\x29\x21\xf3\xf0\x0b\xa4\xf3\x64\x01\xf7\xf6\x78\x34\x72\xe5\x6c\x7e\x16\x4c\xc3\x9b\xf3\x24\x0a\xd3\x79\x38\x03\x21\x05\x17\x16\x35\xcb\x2c\xbf\x47\xd2\x34\x0f\xdc\x16\xe0\xbf\x4f\xd3\x38\xd6\xf2\x71\xd3\xb6\x4e\xd6\xd5\x37\x71\x12\x7d\xbd\xec\xbc\xb5\x2d\x69\x1a\x14\xb9\xfb\xef\x04\x17\xcf\x15\x17\xff\x94\x84\x72\x9f\x0f\xa3\xbf\xc2\x4f\xf4\x54\xd9\x5f\x78\xf2\x39\x04\xcc\x0a\x09\xde\x34\xfb\x5e\x73\x8d\xe3\x71\x77\xf1\xd0\x80\x23\x60\xe9\xb9\x1e\x4b\xef\x04\xda\x13\x0f\x4e\x4f\x81\xa2\xcd\x28\x53\xb6\xfb\xfc\x4c\x8a\x95\x9f\xd3\xd1\xb1\xea\xe8\x3f\x27\xc5\x1b\xc5\xf7\x47\xc5\x97\x71\xd0\x7b\x9b\xb8\xbe\x70\x78\x08\x4b\x02\x00\x10\xc4\x71\x94\xa4\x93\x83\x66\x87\x0c\x8e\x28\xdd\x23\xaa\xbb\x9c\x6b\x18\x2a\xf8\x41\x7d\xc5\x15\xdd\x9d\xf4\x01\xae\xd6\xa5\xbc\x65\xe5\xf5\x52\x70\x91\xe3\xe3\xb0\xd6\xe5\xe4\x60\xd7\x8d\x6a\x29\x2d\x55\x1b\xc5\x97\xc2\xea\xda\x58\xcc\x87\x85\x34\x76\x72\xd0\xf4\xb3\xff\x1b\x1f\xb5\x2e\xe1\xb6\xbf\xe2\xca\xe5\xdb\x4f\x65\xb1\x52\x25\xb3\x08\x1e\x53\xd6\x83\xfe\x76\x77\x9b\xfd\x43\xed\xf7\xc9\x74\x36\x03\x14\x56\x6f\x94\xe4\xc2\xfa\xa6\x00\x67\x81\xcc\xc3\x34\xb9\x8c\xa3\x20\x4c\xe1\xaa\x7b\xaa\x85\xf7\x12\xbc\xde\xdd\x6f\xb8\x77\xfd\x34\xf9\x67\x00\x00\x00\xff\xff\xd9\xcc\x02\xe5\x7b\x03\x00\x00")

func assetsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfile,
		"assets/Dockerfile",
	)
}

func assetsDockerfile() (*asset, error) {
	bytes, err := assetsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDockerfileGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x51\x6b\xdb\x3e\x14\xc5\xdf\xf5\x29\x2e\xcd\x9f\xf0\xdf\x86\xad\xb6\x0f\x1b\xf4\x2d\x2c\xa3\xcb\xda\xb5\x25\x64\xdb\xcb\x5e\x6e\xa5\x1b\x59\x44\x96\x84\x74\x9d\x2c\x08\x7f\xf7\x11\x3b\xa4\x81\x32\xe8\x8b\xc1\xc7\xe7\xe7\x23\x9d\x33\x11\x13\x98\x07\xb5\xa1\xb4\xb6\x8e\x2a\x13\xc4\x41\xf9\x1c\xe2\x3e\x59\xd3\x30\xfc\xaf\xde\xc1\xf5\xe5\xd5\xc7\xea\xfa\xf2\xea\x13\x7c\xeb\x7c\x24\x0b\x77\xb8\xc3\x36\xf0\xe8\x5d\x35\x36\x43\x0e\x6b\xde\x61\x22\xb0\x19\x12\x39\xc2\x4c\x1a\x3a\xaf\x29\x01\x37\x04\xdf\x17\x2b\xb8\xb7\x8a\x7c\xa6\x7a\x80\x1a\xe6\x78\x23\x65\x88\xe4\x73\xe8\x92\xa2\x3a\x24\x23\xdd\x68\xc9\xb2\xb5\x5c\x1d\x5f\xea\xd8\x44\x31\x11\xa5\x68\x5a\x5b\x4f\x70\x81\x91\x2f\xfa\x5e\x2c\x7f\x3c\x00\x46\xae\x0c\x31\x74\x51\x23\x13\x4c\xa7\xf0\x5b\x00\xc0\x49\xb7\x3e\x33\x3a\x07\xd5\x7e\x90\x3a\xb6\x2e\x83\x09\x0e\xbd\x01\x63\x19\x54\x97\x1c\xec\x0e\xce\x52\x12\x7a\x43\x50\xaf\x12\x6e\x6d\xae\x67\x5a\x07\x9f\xeb\x59\xe4\xfa\x09\xd5\x06\x0d\xe5\xbe\x87\x52\xea\xe1\x49\x5e\xf7\xfd\xeb\xb8\x2e\x9a\x84\x9a\x0e\x71\xd3\xe9\x49\x55\x8e\xd0\xbf\x98\x53\x0b\x55\x5a\x83\xdc\x62\x92\x0a\x55\x43\x12\x23\x4b\x4c\xaa\xb1\x5b\xca\xf2\xfd\xf8\xc5\xd9\xe7\x41\x77\x36\x73\x96\x42\x7c\x79\xf8\x09\xb7\x8f\x4f\xb3\xd5\x57\x90\x1a\x19\x85\x98\xc0\xe2\x78\x3b\x63\xdb\x96\x86\x3e\xda\x8d\xb6\x09\xaa\x08\xff\x8d\x5e\xf9\x6c\xcf\x92\x87\xdb\x56\xf9\x1e\xaa\x70\x6e\x90\x03\x3f\x2c\x92\x6f\xa4\x4c\xb8\xab\x8d\xe5\xa6\x7b\xee\x32\x25\x15\x3c\x93\xe7\x5a\x85\x56\xf2\x50\x4d\xa5\xec\x48\xc8\x16\x33\x53\x3a\xe2\x2f\x29\x4d\x1b\x34\x7c\xf8\xf3\x3a\x42\x1c\x8b\x13\x67\x63\x8e\xe3\x1f\xf6\x2c\xc5\xae\x4f\xf5\xdf\x86\x45\x1b\x43\xe2\x27\xe4\xa6\xef\xc5\x6c\x3e\x3f\x74\x3f\x1b\x4b\xea\xfb\xd3\xbf\x73\x52\xb2\x94\x7f\x50\xa5\x90\xcb\xf4\x06\x7c\x49\x31\x64\xcb\x21\xed\x47\x6a\x38\xe4\xaf\xc7\xe5\xdd\x7c\xb1\x7c\x93\xf7\x6f\x00\x00\x00\xff\xff\x6e\x1a\x7c\x76\x45\x03\x00\x00")

func assetsDockerfileGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfileGo,
		"assets/Dockerfile-go",
	)
}

func assetsDockerfileGo() (*asset, error) {
	bytes, err := assetsDockerfileGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile-go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDockerfilePython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4d\x6f\xdb\x30\x0c\xbd\xeb\x57\x10\x0d\x50\x6c\x07\x49\x4d\x06\x74\x40\x6f\xc1\x72\xc9\xda\x7d\x20\x48\xb1\xcb\x2e\xb2\xc5\xd8\x44\x15\x49\x93\xe8\x04\xae\xe7\xff\x3e\xd8\x71\xb2\x6c\xa8\x2f\x26\x1f\xf9\x1e\xfc\x1e\x3d\x13\x33\x58\x85\xf2\x05\xd3\x8e\x1c\xca\xd8\x72\x1d\xbc\x18\xd0\x4f\x21\xb6\x89\xaa\x9a\xe1\x5d\xf9\x1e\x16\x77\xf3\x7b\xb9\xb8\x9b\x7f\x84\xcf\x8d\x8f\x48\xf0\x68\x8e\x66\x1f\x38\x8c\xbb\xdb\x9a\x32\xe4\xb0\xe3\xa3\x49\x08\x94\x21\xa1\x43\x93\xd1\x42\xe3\x2d\x26\xe0\x1a\xe1\xcb\x7a\x0b\x4f\x54\xa2\xcf\xa8\x46\x52\xcd\x1c\x1f\xb4\x0e\x11\x7d\x0e\x4d\x2a\x51\x85\x54\x69\x77\x5a\xc9\x7a\x4f\x2c\xa7\x46\xc5\x3a\x8a\x99\xe8\x3a\x8b\x3b\xf2\x08\x37\x26\xf2\x4d\xdf\x8b\x19\xac\x7d\x66\xe3\x1c\x98\xc8\x10\x4d\xf9\x62\x2a\xcc\x4a\x6c\x9e\xbf\x0e\x88\xac\x90\xa1\x89\xd6\x30\xc2\xed\xed\x15\x52\x25\x63\x11\x64\x3b\xa0\x3f\x05\x00\x5c\x66\x34\xe9\xc9\x76\x84\x1a\x26\x97\xe1\x14\xca\xf4\x92\x91\x22\x98\x44\x66\x01\x65\x93\x1c\x1c\x07\x5a\x45\x3c\x09\xbd\xf1\xbc\x3a\x2a\xe6\x95\xb4\x78\x00\x47\x45\x42\x63\x1d\x79\xbc\x3f\x03\xc5\xeb\xe2\x5c\xe6\x5f\x8e\x18\x3f\x5c\xda\xec\x86\x52\x74\x5d\x32\xbe\x42\x50\xdb\x64\x0e\x94\xd5\xd2\xda\xe0\xb3\x5a\x46\x56\xdf\x27\xcf\x7d\xff\x8f\xe9\x2b\x1b\x5d\xa7\xfa\x5e\x74\x1d\x7a\xdb\xf7\x42\xcc\xe0\xf9\x14\x48\xa4\x38\x52\x06\x3b\x97\x75\x79\xce\x66\x98\x5e\xc5\x1b\x5b\xf4\x87\x71\x7d\xf4\x2c\x9f\xc6\xe3\xe5\x07\xad\x93\x39\xaa\x8a\xb8\x6e\x8a\x26\x63\x2a\x83\x67\xf4\xac\xca\xb0\xd7\x6d\xdb\x34\x7a\x24\xca\x49\x1f\x93\xde\x9b\xcc\x98\x74\x41\xfe\xff\x11\xfc\x86\xc2\xe4\xfa\xef\x97\x5e\x9d\xfb\xf4\x7b\x0c\x17\x5f\xae\x56\x83\xa3\x65\x2a\x6b\x3a\x60\xdf\x83\xb6\x86\x8d\xf8\xf1\x6d\xf3\xb8\x5a\x6f\xa6\xee\x2c\xf1\x27\x00\x00\xff\xff\xe4\x5c\xa5\x11\xdc\x02\x00\x00")

func assetsDockerfilePythonBytes() ([]byte, error) {
	return bindataRead(
		_assetsDockerfilePython,
		"assets/Dockerfile-python",
	)
}

func assetsDockerfilePython() (*asset, error) {
	bytes, err := assetsDockerfilePythonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Dockerfile-python", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypoint = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x8f\xbd\x4e\x2b\x31\x14\x84\x7b\x3f\xc5\xdc\xec\x2d\xa0\x88\x1d\x28\x28\x82\x68\xa0\xe2\xaf\x4a\xba\x28\x85\xe3\x3d\xbb\xb6\x94\xd8\x96\x7d\x56\x21\xb2\xfc\xee\x68\x77\x43\x89\x44\x39\x9a\x6f\x3e\x69\x9a\x7f\xea\xe0\xbc\x3a\xe8\x6c\xb1\xa4\x41\x34\xa2\x01\x79\x4e\x97\x18\x9c\xe7\x29\xbe\x84\x78\x49\xae\xb7\x8c\x1b\x73\x8b\xfb\xd5\xdd\x03\xde\x06\x1f\xc9\xe1\x5d\x9f\xf5\x29\x70\x98\xb0\xad\x75\x19\x39\x74\x7c\xd6\x89\xe0\x32\x12\x1d\x49\x67\x6a\x31\xf8\x96\x12\xd8\x12\x3e\x5f\xb7\xf8\x70\x86\x7c\x26\x39\x8d\x2c\x73\x5c\x2b\x15\x22\xf9\x1c\x86\x64\x48\x86\xd4\xab\xe3\x8c\x64\x75\x72\xbc\xbc\x06\x19\x6d\x14\x8d\x70\x1d\x76\x3b\xfc\x6f\xf0\x84\x15\xf6\xfb\xc7\x51\xeb\x05\x40\x5f\x64\x30\xbe\x10\x9d\x13\x82\x8c\x0d\x58\x3c\x53\x17\x12\x61\x63\x92\x8b\x8c\x0d\x53\xcc\xeb\x85\x28\x25\x69\xdf\x13\xe4\x5c\xcf\x6d\xad\xd7\x4d\x29\xb2\xd6\x11\x92\xb5\x8a\x52\xc8\xb7\xb5\xfe\xf8\x7e\x13\xfd\x45\xf1\x1d\x00\x00\xff\xff\xae\x78\xf8\x79\x68\x01\x00\x00")

func assetsEntrypointBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypoint,
		"assets/entrypoint",
	)
}

func assetsEntrypoint() (*asset, error) {
	bytes, err := assetsEntrypointBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypointGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\x5d\x6b\xe3\x48\x10\x7c\x9f\x5f\x51\x27\x0b\xce\x0e\x48\x8a\x73\x70\x07\x0e\x86\x4b\xee\x82\xcf\xb7\xeb\x64\x89\xbd\x79\x09\x26\x4c\xa4\xd6\x78\x88\x34\x23\x66\x46\x4e\x8c\xac\xff\xbe\xc8\x72\x94\x0f\x67\xd9\xbc\x49\xd5\xd5\xd5\x55\xd3\x74\xef\xb7\xe8\x5e\xaa\xe8\x9e\xdb\x15\x02\x2a\x59\x8f\xf5\x40\xca\x99\x4d\xa1\xa5\x72\x81\xd0\x3b\xe4\x1f\x5d\x6c\x8c\x14\x2b\x87\x7e\x3c\xc0\xc9\xf1\xf0\xcf\xe0\xe4\x78\xf8\x17\xfe\x2f\x55\x41\x12\x5f\xf8\x23\xcf\xb5\x6b\xb9\x8b\x95\xb4\xb0\x3a\x75\x8f\xdc\x10\xa4\x85\xa1\x8c\xb8\xa5\x04\xa5\x4a\xc8\xc0\xad\x08\xb3\xe9\x02\x5f\x65\x4c\xca\x52\xb8\x6b\x5a\x39\x57\x8c\xa2\x48\x17\xa4\xac\x2e\x4d\x4c\xa1\x36\x22\xca\x5a\x8a\x8d\x72\xe9\x82\xfd\x4f\x58\xac\x0a\xd6\x63\xdf\xce\x16\xff\x8d\x3d\x7f\x72\xd5\x7c\x34\x11\x46\x7e\xf3\xe5\x31\x26\x53\xdc\xde\xc2\xef\x61\x8c\x63\x2c\x97\xa7\xcd\x40\xc5\x00\x7a\xa2\x18\x4d\x4e\x96\xca\x8e\x35\xc4\x18\x1e\x57\x1b\xef\x35\xd3\x10\x4f\xb4\xca\x36\x58\x5c\x9f\xdd\x4c\xe7\x77\x93\xab\xbb\x9b\x8b\xeb\xf9\xf4\xea\x72\xec\xf7\x85\xc6\x9a\x8c\x95\x5a\x61\x8b\xb8\x74\x08\x52\xfc\x81\x20\x81\x07\x0f\x5b\x34\x39\x6d\x24\x74\x14\x89\x01\xa3\xcc\xd2\x2f\xf4\x86\x8d\xb3\x78\xa5\xe1\xcd\xc9\x61\xd2\x89\x8f\xe0\x1f\xb0\xbd\x86\xbb\xe6\x19\x3c\xbf\x3f\x99\xce\x66\x17\x6f\x94\x0e\xe8\x10\x32\xcf\x69\xe0\xed\xf2\xa6\xda\x80\x20\x15\xfc\xea\xef\xd1\xc9\xa8\xef\xf7\x82\xe1\xa0\x3e\x45\xa2\xdf\x18\x20\xb5\x96\x46\xab\x9c\x94\x1b\xc1\xa7\x66\x62\x42\x71\xd6\x6c\x32\x78\x82\x4f\x2c\xd1\x8a\x18\x6b\x1b\xce\x29\xd5\x86\x30\x55\xd6\xf1\x2c\xc3\xdc\x51\x61\x47\x1e\xab\x2a\xc3\x95\x20\x84\x6d\x7d\x5f\xae\xeb\x7d\x57\x55\x85\x75\xdd\xb0\xc2\xba\x66\x55\x45\x2a\xa9\xeb\x67\xc5\x43\x29\x99\x22\x7c\x51\xe8\xa4\x3f\x25\x5a\x55\xcd\x02\x3a\x8e\xd0\x10\xe4\x10\x38\x84\x51\x18\x86\x1e\x7b\x07\xbc\x37\xb3\x8f\x37\x8f\x8d\x2c\xdc\xcf\xd2\xb5\xd5\x4f\x85\x3b\x10\x8a\x8e\x20\x53\xd8\x16\x96\x56\xfd\xee\x20\xe4\x9a\x54\x8e\xd2\x12\x72\xfe\x40\xd0\xed\xbd\x24\x94\xf2\x32\x73\x10\x3a\xe3\x4a\xc0\x91\x75\x88\x75\x9e\x73\x95\x84\x38\x8a\x76\xb3\x9a\x97\xea\xdc\x74\x2e\x3f\xe3\xaf\x7b\xa7\xf6\x28\x02\xc2\xe4\xf2\xfb\x8c\x3f\x50\x2a\x33\xc2\x72\x89\xed\x76\x8f\x7f\x08\x9e\xcf\xff\xcd\x3f\xc2\x5f\x81\xdd\x69\x35\xd8\xf3\x59\xb4\x96\x27\xfa\xbc\x94\x59\x72\x66\x84\xad\x6b\x06\x08\xdd\xc6\xab\xaa\x83\x52\xe7\xf3\x85\x15\xac\xf7\xbb\xdb\x55\x77\x69\x52\xd9\x05\xfb\x11\x00\x00\xff\xff\x88\x1d\xa1\x48\xdd\x04\x00\x00")

func assetsEntrypointGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypointGo,
		"assets/entrypoint-go",
	)
}

func assetsEntrypointGo() (*asset, error) {
	bytes, err := assetsEntrypointGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint-go", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsEntrypointPython = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x5d\x6f\xda\x30\x14\x7d\xf7\xaf\x38\x33\x51\x05\x0f\x4e\x08\x0f\x9b\x94\x0a\x69\xed\x56\xa9\x6c\xfd\x40\x85\x4d\xaa\xba\xb6\x4a\x93\x0b\xb1\x16\x6c\xcf\x36\x29\x88\xe6\xbf\x4f\x01\xda\x22\xb6\x4a\x7d\x89\x94\x7b\xae\xcf\xc7\x3d\xad\x0f\xd1\x83\x54\xd1\x43\xea\x0a\x08\x9a\xb3\x16\x6b\x81\x94\xb7\x4b\xa3\xa5\xf2\xc2\x2c\x7d\xa1\xd5\x7a\xfa\x45\x9b\xa5\x95\xd3\xc2\xa3\x9d\x75\xd0\xeb\xc6\x1f\x45\xaf\x1b\x7f\xc2\xb7\xb9\x32\x24\xf1\x3d\x7d\x4c\x67\xda\xeb\xf5\xee\xb8\x90\x0e\x4e\x4f\xfc\x63\x6a\x09\xd2\xc1\x52\x49\xa9\xa3\x1c\x73\x95\x93\x85\x2f\x08\xe7\x83\x31\xce\x64\x46\xca\x51\xb8\x7e\x54\x78\x6f\x92\x28\xd2\x86\x94\xd3\x73\x9b\x51\xa8\xed\x34\x2a\x37\x2b\x2e\x9a\x49\x2f\xb6\x3f\xa1\x29\x0c\x6b\xb1\xf1\xc9\xd5\x79\x7f\xe1\xc9\xce\xd8\xf0\x68\x7c\xda\xe7\x91\xd5\xda\x47\xa1\x59\x92\xaa\x9a\x58\x49\xd0\xcc\x39\xcb\x29\x2b\x1b\x23\x62\x81\xe1\xf5\xf8\xf4\xf2\xe2\xfe\xf8\xc7\xe0\xec\xeb\xfd\xd1\xd5\xe0\xa8\x77\x7f\x39\x1c\x8f\xfa\x5c\x2c\x10\x77\x21\x7e\x23\x3e\xe7\x8c\xaa\xb4\x04\x0f\xda\x6b\x26\x48\x25\x3d\x44\x87\xe3\xe0\x00\xbf\xf6\xb0\x4a\x5a\x3f\x4f\x4b\x52\x95\x78\x5e\x63\x4c\x4e\x70\x73\x83\xa0\x85\x3e\xba\xb8\xbd\x3d\x6c\xf2\x2a\x06\xd0\x82\x32\x34\xa7\x66\x13\xc9\x98\xa5\x34\xd7\xaa\x5c\x62\x78\xfd\xf3\xe4\x6a\x34\xb8\xbc\xe8\xbf\x2a\x3a\x9f\x96\x25\x44\x89\x27\x34\x67\xe3\x2e\x42\x14\x4d\x39\x9e\x30\xb5\x64\x20\x4e\xc0\xef\x82\xb8\x1d\x3c\xdd\xdc\x89\xdb\x4e\x33\xf7\xa9\x2c\x21\xe2\x0e\xa3\xac\xd0\xe0\x23\xf2\x18\xae\xdb\x43\x45\xd6\x49\xad\x12\x04\x2f\x4a\x9c\xed\x09\x39\x88\x6a\x07\xdf\xc2\xa5\xce\xd2\x72\x77\xcc\x26\xda\x82\x20\x15\x82\xd5\xe7\xa4\x97\xb4\x83\x96\x88\x3b\xf5\x21\x72\xdd\xe4\x7b\x51\x26\x55\x49\xab\xd5\x8c\x94\x4f\x10\x10\x67\xc0\x4e\x09\x01\xb1\x5c\x2b\x62\x5b\xab\xc7\x34\xd1\x96\x30\xd8\x7a\x19\x79\x32\x2e\xe1\x6c\xb5\xb2\xa9\x9a\x12\xc2\x0d\xbe\x85\xeb\x7a\xf3\x6a\xb5\x0a\xeb\x9a\x3d\x7f\x49\xe5\x75\xfd\xcc\xb7\x4f\xb4\xa9\x43\x10\x2c\xfd\x99\x4b\x4b\x8d\x2b\x17\xfa\x85\xdf\xed\xc6\x48\xf3\x7a\x0d\xfb\xcf\x6a\xd3\xd8\x8b\xa1\x3d\x2b\x7c\xed\x82\xff\xdf\xcc\x36\xdc\x28\xb3\xd2\xf8\xb7\xb2\x6d\xd0\x77\xf1\xbd\x45\xf4\x1e\x8a\xbf\x01\x00\x00\xff\xff\x98\x38\xfc\x89\xed\x03\x00\x00")

func assetsEntrypointPythonBytes() ([]byte, error) {
	return bindataRead(
		_assetsEntrypointPython,
		"assets/entrypoint-python",
	)
}

func assetsEntrypointPython() (*asset, error) {
	bytes, err := assetsEntrypointPythonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/entrypoint-python", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/Dockerfile": assetsDockerfile,
	"assets/Dockerfile-go": assetsDockerfileGo,
	"assets/Dockerfile-python": assetsDockerfilePython,
	"assets/entrypoint": assetsEntrypoint,
	"assets/entrypoint-go": assetsEntrypointGo,
	"assets/entrypoint-python": assetsEntrypointPython,
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
	"assets": &bintree{nil, map[string]*bintree{
		"Dockerfile": &bintree{assetsDockerfile, map[string]*bintree{}},
		"Dockerfile-go": &bintree{assetsDockerfileGo, map[string]*bintree{}},
		"Dockerfile-python": &bintree{assetsDockerfilePython, map[string]*bintree{}},
		"entrypoint": &bintree{assetsEntrypoint, map[string]*bintree{}},
		"entrypoint-go": &bintree{assetsEntrypointGo, map[string]*bintree{}},
		"entrypoint-python": &bintree{assetsEntrypointPython, map[string]*bintree{}},
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

