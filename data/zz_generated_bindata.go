// Code generated by go-bindata.
// sources:
// data/ca_bundle.pem
// data/public_key.json
// data/aws_identity_cert.pem
// DO NOT EDIT!

package data

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

var _dataCa_bundlePem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x97\x49\x93\xa3\x30\x9e\xf6\xef\x7c\x8a\xf7\xee\x78\xc3\x80\x01\xe3\x43\x1f\x24\xf6\x45\x60\x76\xc3\x8d\xc5\x06\x63\x83\xc1\x06\xb3\x7c\xfa\x91\x33\xab\xaa\xab\xaa\x7b\xa6\x27\x62\x32\x22\x23\xd3\xb2\x24\x84\xa4\xff\xf3\x7b\x9e\xff\xff\xf9\x81\x92\xa2\x59\xff\x4f\x90\x5c\x5f\x93\x35\x01\xf8\xd2\x57\x2b\x81\x34\x4d\xf2\x7c\x41\x00\xab\x52\x82\x49\x83\xa0\xd4\x7c\xd8\x6e\x82\x90\x89\x54\xce\x98\x77\xbb\xa7\xee\xa3\xfb\xc8\x37\xf1\x10\x89\x3e\xb0\x60\x79\xeb\xab\xdb\x55\x39\x4c\x24\x04\xce\x4b\x26\x80\x08\x58\xe4\xbc\x26\xc1\x89\xc5\xd0\x71\x14\x69\xd2\xc3\x60\x95\x8e\x08\x90\x0a\xa0\x02\x49\xa8\x90\xe2\x44\x54\x75\x6e\x0e\x23\x72\x6f\x93\x3c\x7d\xf5\x13\xa5\x19\xc2\x2c\x92\x39\x22\xa3\x99\x32\x68\x0e\xef\x42\x80\xa2\xe3\x81\x19\x41\xe6\x24\xfa\x12\x85\x7c\xb0\xa0\x1a\x4c\x48\xc4\xbf\xf2\x03\xb7\x69\x9f\xb6\xd9\xf6\x7f\xb5\x4d\x6e\x2d\x99\x04\x02\xb7\xaf\x27\xc1\x0a\x09\x61\x88\x66\x71\xfd\x2c\xd3\x0a\x21\x78\xf8\xf0\x26\x0f\xf1\xa9\x7b\x67\xb5\x14\x22\x88\xbe\x57\x34\x23\x14\xd0\xe1\x52\x34\xe1\xa2\x49\x16\x24\x34\x51\x16\x90\xc3\x4f\xa2\xf3\x63\x69\x53\xf2\xb5\x34\xbc\xb2\xa9\x2c\xa5\xeb\xe7\x55\x04\xaf\x57\x3c\x2d\xdb\x89\x8e\x84\xdf\x3b\x00\x80\xd1\xa0\x38\x81\xcf\xf7\x06\x01\x1e\x78\xe3\x1c\x51\xf0\xab\xa4\xdd\xe5\x1c\x9f\xbe\xca\x5d\x34\x4e\x1c\x32\x41\xc1\x0e\x92\xc7\x41\xed\xe1\xa1\xf5\x61\xe4\x0a\x9b\xdd\x8f\xe1\x63\xf3\x12\x6d\xf7\x89\x0a\xe6\x42\x83\xac\x4d\x88\x1c\xad\xc7\x94\xd9\xd5\xcc\x64\xcd\x55\x77\x1f\x16\x2e\x0d\x8c\x1b\xe3\x53\xfd\xf9\x00\xed\xc9\x90\xeb\xc9\xe0\xd6\x66\x9e\xcd\x30\x7e\xec\x33\x35\xbc\x9e\x5e\xc7\xbb\xce\xf5\x36\xea\xe4\xf2\xcc\x12\xd9\x5d\x3c\x6e\x28\x7e\xde\x40\x9a\x03\x64\x77\xbd\x3a\xf6\x78\xbc\x5d\x16\x71\x39\xbb\xcc\xec\x94\xd5\xa5\xe6\xb8\xf8\x41\x1d\xc2\xcd\xb9\x09\x76\x6d\xba\x5e\xde\x9d\x0c\x36\xae\xbd\x72\xd1\x23\x6c\x08\xc8\xce\x1b\x99\xee\x42\x7e\x3e\x1b\x96\xbb\x1f\xb9\x74\x15\x8b\x80\x8d\xc3\x13\xe5\xa7\x53\xf7\x6c\x66\x57\xa0\x36\xd1\x0b\xc4\xcd\xca\xf5\xc7\xcd\xca\x83\xa7\xa8\xf9\x02\x2d\xa3\x70\xa1\x2f\x13\x41\x6a\xb5\x61\x0f\xd2\x29\xdf\x86\x97\x66\xf0\xe5\xbc\x62\x37\xe0\xa2\xc4\x48\x41\x7d\xff\xd6\x39\x33\x3f\x5d\x41\xd5\x2b\xac\xaf\x6d\xc4\x27\xe9\x0e\x88\xe7\x6f\x1b\xfe\x14\x40\xe1\xec\xf0\x9a\x42\x18\x23\xb0\x52\x73\xef\x5f\xb5\xc1\x48\xe2\xd9\xa0\x10\x1a\x7d\x7d\x08\x0f\x5a\x76\x07\x25\x82\x00\x28\x35\xde\xed\x3d\xbe\xb1\xd0\x5a\x81\xf7\x39\x62\xd5\xc5\xed\x17\x5e\x12\x44\xa0\x10\xc0\x51\xb7\xa0\x94\x00\x02\xcc\xe7\x9c\x0b\x71\x92\xe0\x76\x72\xa4\xcf\xb5\xae\x6a\x50\x7c\x0d\x70\x18\x49\x2e\x9d\x20\x4a\xdd\x06\xde\x0d\xbc\x75\x91\x77\x34\xec\x00\x2e\xe7\x88\x28\xe4\x37\x7b\x2c\xc0\xa4\x7e\xae\xa8\x4b\xd6\x10\x96\x93\xfc\x00\x41\x05\x6b\x54\xf9\xfe\xeb\x0d\x96\xe0\x2e\x30\x5a\x94\xac\xea\xab\x82\xb6\x50\x96\xd3\x79\x8a\x35\x63\x8a\x21\x74\x02\x15\xaf\x40\x92\xb2\x15\x0e\x48\xe0\x15\x41\x78\x29\xf8\xa2\xc8\x70\x05\xb0\xba\x5a\x8f\x42\x75\x27\xfb\xca\xbf\x33\xda\x5a\x73\x81\x5d\x32\xfa\x40\xc6\xb4\x34\x9b\xcd\x8f\x1b\xda\xba\x0b\x51\x9c\x2c\xd2\x6c\xac\x77\xe6\x03\x0e\x96\xe5\x13\x96\x92\x0c\x9d\x7c\x02\x8f\x78\x4c\x55\x97\xcc\xc5\xc7\xdb\xa4\xad\xa5\xf8\x8f\x13\x78\x87\xdf\x3a\x58\xf7\xbc\x06\xdb\xef\x0d\xe3\x25\x5b\x04\x34\x12\xbd\x12\xf5\x60\xaa\xae\xec\xaf\x95\xc5\x3b\xfd\x45\x98\xad\xfe\xce\x76\x6e\x1d\xfb\xd2\x18\xff\x28\xdb\x42\xd5\xa9\x7c\xe7\x8c\x31\x7d\x18\xcc\xdd\x6f\xdf\xe3\x01\x3f\x8a\xaa\xd0\x80\x83\x10\x78\x4c\x84\x00\x62\x25\x61\x24\x04\x1c\x0d\xfe\x5d\x38\xc2\xf4\x5d\x38\xc0\x11\xe4\xdb\x93\xa1\xc6\x5d\x7b\x7c\x30\xb2\xa0\xda\x7e\x1d\xef\x2c\xdf\x0e\x35\x8a\x60\x0f\xca\xb0\x4d\xb9\xe4\xda\x2f\xba\x74\xdd\xec\x59\x7a\x93\x52\x01\xbb\x70\x57\x30\xc5\x97\xe6\xf4\x7a\xd1\x77\x7d\xd2\xe5\x1e\x75\xf4\xb1\xab\x8c\x92\xe5\x68\xf6\x76\x2a\xf9\xdb\x91\x16\x2c\x76\x20\x38\x65\x9f\x21\x27\xf7\x79\x81\x9f\xc5\xc4\x1a\x62\xbf\xd8\x47\x47\x91\x0f\x92\xab\x6b\x00\xfd\x08\x4f\xe9\x8e\xdc\x82\x6c\x1a\x93\x33\xa9\xa4\xb2\xe4\xf0\x63\x99\xc7\x4e\xe9\xb5\x1b\x4d\x81\x1a\xc1\x6f\xcd\xa9\x82\x96\x9f\xf8\x41\x28\x45\xa3\x10\x04\x30\x0c\x29\x3e\x1e\x32\x70\x3d\xf6\xbb\xe5\xd4\x63\x8d\xe2\xed\x75\x93\x0f\x30\x1a\x13\xef\x96\xbd\x6f\xf8\x80\x8f\x66\xda\xdc\x20\x5d\x52\xc4\xd8\xb9\xcb\xea\xec\x7b\xb1\xa5\x4e\x7c\xdb\xf2\x16\x1f\xee\x63\x9d\x5b\x38\x1e\x0c\x37\x35\xb7\x3c\x17\xb4\x9d\x7f\x1d\x66\x68\xd4\x83\x71\xd4\x3c\x13\x85\xc2\xbc\xbf\x32\x55\x9b\xcf\x6a\xe2\x11\x8b\xb9\x18\xce\xa9\x9a\xe8\x88\x3e\xbd\xc8\xde\x3c\x0b\xd4\x79\x00\x9b\xda\x57\x44\x83\x09\x2e\xf8\x33\xe9\xc9\x7b\xd9\xbb\xf2\x0f\xd6\x34\x69\xca\xe4\xb5\x15\x74\xe9\x93\xee\xdc\x2d\xf1\x25\xec\x92\x25\xfe\xab\xd8\xff\x27\x10\xdc\x6a\x0c\x82\x5d\xff\x07\x08\xbc\xeb\xa3\x5d\x2f\x47\x6e\xea\x99\x27\xb8\xdc\xb4\xfd\xb3\x95\xe6\xfa\xdf\x82\x40\x80\x8d\x28\x99\x7f\xcb\xb1\x24\x02\xfb\x5b\x8e\x4b\x1f\x92\xf2\x92\x7e\xee\x6a\x23\xcd\xd2\x0a\xdc\xef\xf6\xdc\x17\xee\x56\x4d\xe0\x6b\x45\xe6\xb4\x5b\x65\x4a\x30\xeb\x3e\xa8\x7f\x6a\xb8\x2a\x5b\x64\x7c\xd2\x9b\x34\x0a\x5f\x89\x00\x83\x24\xb2\x1e\x59\x73\x78\x65\x74\xd1\x25\x27\xf4\xd2\xa4\xfb\x18\x2f\xcc\x4c\xd8\x2b\x60\xbf\x07\x21\xcc\xac\x3f\x06\xf9\xc9\x49\xa7\xd3\x08\x17\xc2\x02\xbd\x4f\x61\x60\xe1\xbf\xe7\xad\xdb\x25\xcd\xbd\x8e\x4f\xee\x5d\x93\x64\x8a\x28\x94\xea\x9d\x37\x77\xf2\xec\x81\x41\x93\xf2\xe5\x17\x8a\xc4\x60\xb1\x7c\x69\xf9\x27\x8a\xd0\x0e\xf9\xda\xfa\xc1\xd1\x4f\x14\xd9\xfe\xff\x1e\x45\x09\x82\xf9\x57\x3f\x3c\xdc\xf9\x09\x47\x4d\xd6\xdf\x78\x0f\x9c\xd2\x21\xa5\x12\x7d\x90\xec\x69\x1f\x24\xe9\xc6\x23\xd1\xf0\xba\x2c\xac\x2e\x58\x66\x80\x88\x45\xf0\x88\x37\x5b\x00\xce\x03\x2b\x10\xc6\x91\xa9\x33\xa5\xaa\x1a\x67\xeb\x54\x13\x79\x7a\x50\x4b\x19\x92\x97\x68\x1f\x53\x4c\x45\x1f\xf4\xfb\xe3\x40\x95\x55\x7c\xbc\x93\x15\x90\xde\x4f\xa0\x0d\xd5\x60\x97\xce\xae\xb3\x5f\xbd\xef\x58\xcf\x07\x7c\x3f\x76\x99\x87\x4a\x55\x5e\x13\x44\x1c\x6c\x4e\xd3\xf8\x7c\xc3\xad\x17\x6a\x70\x5b\xc6\x8b\xae\xd3\x6e\x38\xb3\x45\x5d\x16\x71\xc2\xdd\xb6\x0f\x8d\xee\xce\xa1\x11\x8e\xae\xcc\x5c\xda\xc3\x00\x33\xae\xb0\xfa\xbc\x59\x03\xd6\xdc\xf6\x13\xa1\xc9\x40\xc9\xd4\xa7\x53\x9a\x46\xb3\x49\xb7\x2f\x77\x6e\x8e\x81\x58\xaa\x3b\xc3\x50\xed\xb0\x66\xc6\x21\xea\x36\x41\xd5\x22\x3d\x1b\xef\x6a\x75\xce\x98\xa6\x0e\x72\x2c\x45\x4d\x5a\xb9\x51\xca\x11\xa1\x3d\xd6\x13\xab\xb2\x9e\xb5\x6e\xc9\x73\x39\x99\x27\x72\x28\x54\x40\x51\x4c\x79\x3b\xb0\x7b\x29\x8a\xb8\x7d\xce\xe4\x27\xbe\xd6\x15\xc3\xac\xc4\xcd\x33\x2f\xfa\x57\x4f\xf2\x1d\x7f\x13\xaf\x94\x49\x1c\x76\x72\x7e\x6a\xda\x2d\xd7\x05\xc2\xb2\x5e\x8d\xe7\x1d\x30\xd9\xe1\xbd\x37\x23\x2d\x9b\xf3\xfc\x1c\xda\xf2\x8e\x51\x2e\x9a\xc8\x2e\xaa\x76\x88\xb7\x8e\x00\xb7\x9a\x26\x4a\xa5\x34\x6d\xec\xc5\x69\x88\xba\xf4\xc6\x4c\x7f\x6a\x7d\x49\x0a\x60\x92\x00\x48\xed\x4f\xa1\x7c\x60\x3f\x20\xc0\x7f\x09\x9e\xf4\x4d\x17\x19\x81\x2f\x04\x4d\x62\xf9\x05\x8d\xe3\x87\x42\xc4\xc7\x0d\x80\x52\x51\x10\x24\xbf\x51\x54\x3a\x11\x84\x9e\xa4\xa0\x45\xb6\xec\x85\x17\x75\x2f\x30\x4b\xec\x16\xd0\x79\x91\x24\xc6\xc0\xfc\xba\x7c\xc9\xb3\x87\x24\x45\x04\x11\x51\x42\x2f\x3f\x4d\xe2\xa5\x2f\xd5\x13\x12\x5e\xcc\xd5\x30\x98\x2c\xe8\x73\xbe\x52\xdc\x72\x85\xcc\x6f\x70\xc0\xcf\x72\x5f\x48\x79\x4c\x66\xf9\x1b\x8a\x10\x50\x14\xad\xa9\x48\xbc\x71\x9c\xb9\x1c\xb0\xb8\x5b\xd3\x0f\x6d\x6f\xd1\x95\xad\xbe\xfd\x0d\x4b\xe6\x6d\xb8\x62\x9a\xd4\x19\x4d\x4e\x36\xf8\x63\x02\x43\x31\x95\x5f\x13\xd4\x79\xeb\x8c\xf9\xc7\xa6\x29\xf9\xf2\x3b\x72\xfe\x26\x4e\xb2\x6a\x18\x14\xe1\x42\x20\xf1\xfb\xdd\xd5\xc9\xc1\xb8\x71\xa6\x0f\x6e\x1e\x02\xf7\xe7\xa4\xcd\xf4\xdf\x4e\x4a\xfc\xcb\xac\x5f\xb8\x91\x7e\xe2\xc6\x40\xa0\x9c\x60\x19\x4b\x78\xd7\x4b\xec\x42\xff\x8d\x22\x01\x7c\x6c\x98\xff\x71\x5d\x08\x27\x73\x72\x06\x9f\x33\x4d\xfb\x86\x1a\x7a\x96\x99\x32\x07\xe7\x77\x2b\x47\x60\x64\x05\x6d\xda\x60\x64\xdc\x8f\x66\xf8\x0e\x6c\xdf\xb2\xa2\xbe\xbd\xad\x5e\x44\xa0\xab\xd2\x79\x92\xf7\x6c\x6d\xf2\x30\x18\xdd\x9a\x9d\xdd\xad\xfc\x10\xf4\x0c\xf1\x0f\x30\x5f\x45\x77\xd7\xd4\x92\xca\x4c\x11\x37\xed\x5f\x8a\x58\x16\x07\x47\xeb\x46\xa9\xb8\xc8\x7b\x30\x6e\x9b\x94\x38\x2f\x46\xd1\x4d\x40\xbf\xf4\xb3\x22\x33\xc7\xbc\x15\x92\x53\xe3\x03\x36\xee\xd2\xe3\xbe\x78\x9e\xfb\xd7\x09\x29\xeb\xde\x70\xe8\xea\x15\xce\x29\x4f\x39\x4c\x69\xbe\xf7\xdb\xa9\xe9\x0a\xb3\x87\x06\x91\xb9\x6e\x5c\xb1\x7e\x63\xfb\xf2\xe5\xa2\x1e\xcd\x9b\x56\xf5\x15\x54\x22\x9d\xcb\x06\x3a\x96\x95\xae\xe5\xea\xbc\x04\x46\x50\x73\xe2\x15\x14\x75\xc1\xdc\x3b\x79\xe2\xd9\xaa\x30\x9e\x82\x14\x5a\x04\x29\x4b\xdc\x36\xa4\x0a\x8b\x76\xd1\xa5\x16\x96\xd0\x73\x85\x16\x3b\xba\x53\x32\x9d\xca\x48\x9d\x97\xf7\xcd\x01\x57\xef\xc9\x4d\x14\x79\x8b\xa9\xbd\xeb\xdd\x1d\x3b\xbe\x2e\xdd\xe3\x46\xe9\x2e\x13\x10\xe9\x2d\xaf\x91\x77\xc8\x9b\x77\x3f\x34\x25\x7b\x0d\x52\xa7\xef\x73\x9f\xb5\x74\xb2\x52\xc0\x3f\xfe\xf1\x7f\xc0\x4d\xf1\x11\x39\x8a\xfe\x89\x1b\x1d\x18\xb9\xed\xdd\x26\xb2\x7c\x16\xdb\x7f\x67\x1d\x90\x52\xce\xc2\x0a\xf4\x4f\xa5\xe0\x12\x8b\x7d\x70\x0f\x7d\xe4\x05\x93\xf6\x9d\x29\x0c\x69\x9e\xfd\x42\x91\x17\x2c\xea\xf7\x4c\x71\xca\x50\x09\xeb\x54\x61\xdf\x99\x72\x68\x31\x08\x56\x53\x80\x7a\xd6\xa0\x11\x6b\xf7\x84\x2f\xf8\xd7\x20\x53\x5a\xee\x7f\x0c\x72\xe8\xb9\xca\x77\xa8\x44\x57\x28\x62\x6a\x90\x69\x94\x74\x31\x2d\xe3\xbf\x58\xa3\x31\x25\x7e\x42\x82\x38\xfb\xe0\x2c\x4f\xb8\x6c\xf0\x81\x22\x51\xfb\x0e\x25\x22\x4c\x71\xdb\x6a\x7d\x3c\x56\x5d\xce\xd6\x8a\x58\xe4\x27\x29\xd2\xb0\x13\xfe\x2b\x25\x11\x38\x26\x39\x3f\xcc\xaa\x24\x40\xa4\x3a\x27\xbd\xfb\x90\x00\x7b\x2b\xff\xc7\x45\x97\xe0\x8c\x8c\x80\xb6\xde\xd8\xa4\xad\x89\x22\xbf\x12\x5f\xba\x23\xe1\x3b\xc8\x10\x38\x5b\xe5\xc1\xce\xad\xf2\x26\xe9\x92\x68\xbe\x69\xb2\x7b\x8f\xe9\x6a\xcc\xe8\xf9\x9d\xd0\x77\x0c\xbb\xa9\xf4\x22\xb6\x36\x6b\x6c\x9b\xc5\xdb\x4f\xe4\x2c\x3f\xc7\x10\xdf\x83\x3e\x24\xfc\xbc\x63\xb8\x7e\x18\xf4\x85\xa0\x4f\x3c\x52\xee\x98\x9e\x56\x55\x28\x41\xe9\x9c\x42\x32\x55\x0e\x98\xdd\x2e\xab\x09\x64\xe9\xae\xda\x57\x32\x22\xfe\x53\x34\xfa\x91\x8c\x42\x51\x7c\x4a\xc6\xfb\x6e\x33\xef\x68\xa3\x24\xc5\xa5\x7e\x54\xfe\xcb\xe5\xb7\xc4\xc2\x6f\x2e\xd8\xcf\x0f\x86\xff\xd4\xc4\x1d\xc9\x1f\xe8\x81\xb6\x95\x63\x62\x35\xc2\xa3\xa1\xd8\x1c\x68\xc2\x62\x52\xf7\xed\xe1\x71\x61\xf5\xc0\x56\x58\xfa\x96\x05\x5d\xef\x30\x27\xb5\xa6\x05\xd2\x22\xfc\x66\x4b\x2f\x52\x9b\x0c\x6f\x94\x86\xd8\xa3\x0c\xad\x13\x70\xfc\x76\xaf\x8f\x28\x1d\x2b\x3a\x32\x9b\xc7\x3e\xd2\x3d\xdd\xa5\xb2\xad\x9e\x0a\x7e\x2e\xdb\x22\xfd\x70\x49\x19\x59\x6d\x5b\xba\x0f\xc2\x1e\x36\xb6\x23\x3f\x0a\xef\xb6\x3f\x3a\xac\xb4\x67\xa9\x2c\x02\xaa\x68\x06\x23\xbb\xbf\xa4\x0c\xc7\xee\xa7\x79\x13\x9c\xe8\xa9\x11\x8f\x12\x75\x13\x0c\x46\x44\x96\x74\xb9\x8c\x05\xe7\x24\x11\x29\x10\xeb\xb2\xb8\x5d\x9f\xb5\xbb\x47\x10\x7b\xa7\xb9\xf1\x7b\xc4\x65\x69\x43\xed\x6b\x67\x1c\x4b\x52\x1c\xc5\xe3\xc5\xdd\x8c\x73\xca\x90\x77\x3a\x79\xdb\x65\x21\xcb\x6e\x6d\x44\xb9\x76\x39\x83\x92\xd5\x09\x87\x89\xe8\x4c\xb5\xf7\x89\xdd\x55\x4e\xba\xea\x94\xec\x5f\xaa\x65\xab\x6a\x4f\xad\x59\xf5\x43\x12\x2a\xd7\xcb\xd6\x64\x7a\x93\x77\x43\x55\x0d\x41\x0c\xcf\x32\xb8\x07\xec\x75\xc7\xcb\x9f\x64\x44\x7c\x45\xa3\x0b\xde\xf1\x33\x39\x89\xdf\xe1\xc5\xff\x70\x08\x82\x00\x9f\x03\xe6\xd4\xfa\x6d\xc1\x54\x87\xff\x8a\x4b\xd8\x20\x61\x9b\x1e\x4f\xaa\xf3\xd5\xd7\x26\x20\x8c\x25\x59\x9f\x2f\xc0\xda\xf4\xa0\xc8\x27\x63\xbd\x6a\x8f\x67\x35\x78\xdd\xba\x48\x89\x22\x22\xf8\x0d\x3d\x6d\x72\x62\x04\x53\x20\x9b\x87\xcb\x70\x50\xed\x9d\xbb\x51\x0e\xf2\x10\x12\x4f\x6b\x3d\x49\x48\xb3\xfb\xb8\x6e\x91\xf4\x5b\xfa\xc1\x4f\x87\x12\x9a\x1c\x1f\xe4\x7f\x04\x9a\x38\x76\xfe\x19\x68\xf8\x91\xc0\x15\xb7\x98\x98\x45\xe6\x0a\xaa\xbf\x92\x4f\xf8\xab\xe3\xae\x1c\x7f\xf5\xc3\xff\x7f\x30\x83\x84\xf8\x6b\x65\x04\xc6\xcc\x05\xe3\x76\x52\x38\x90\x3c\xe0\x59\x91\xa3\x5f\x98\x59\xcd\x16\xd7\xe1\x95\xa5\xf2\xe5\x80\x31\x63\x2d\x99\xf8\x6d\x3d\x55\x0f\x48\x42\x0d\x34\x2c\x05\x0a\x01\xe5\x0f\x56\xc0\xdf\x96\xcb\xfc\x61\xb9\x20\x10\x70\x5e\x67\xba\xd3\xd2\xee\x3a\x79\x77\x47\xbc\xcb\x86\xdb\x7c\xf6\xb3\x9a\x55\xc5\xc3\x56\x21\xc2\x8b\xb1\x40\x31\x1b\x4a\x78\xf0\x67\x99\x24\x0d\x65\xdc\xcc\x38\x00\x24\x9b\xa7\x79\xdc\x6d\x9c\x57\x7d\xb4\x14\xaa\x74\xca\xd2\x64\x36\xc2\x96\x92\x68\x31\x80\xf9\x7e\x2e\x9d\x1a\xee\xd2\x82\x22\xee\x24\x1f\x8f\xd1\xee\x7c\x60\x6d\x57\x30\xbb\x8d\x23\xac\xc3\x74\xee\xf7\x45\xc7\xac\xb0\xcd\x0b\xd1\xa9\xb6\xc1\x81\xcc\x12\x63\x14\xf4\xad\xdc\x51\x01\x75\x7e\xbe\xbd\x6a\xda\x45\x6d\x24\x39\x03\xc1\xd7\xf3\xd4\x18\x0b\x97\x66\x69\x58\xec\xf8\x23\x0a\x99\xd7\xd6\xc0\xf1\xea\x56\x74\xbc\x7a\xbf\x1c\xa0\x1b\x74\x7a\x78\x96\x4e\xa5\x17\x0b\x17\xbb\xe5\xf5\xdd\x16\x3b\xdb\x82\xa2\xb9\xed\xa6\x4b\x70\xf6\x7a\x1f\x9f\xac\x11\xed\x17\x2f\xb5\x5c\xc8\xb5\xba\xaa\x88\x2d\x9d\x1c\x6a\x3e\xd9\x6d\xc3\x25\xb4\xa5\xb0\x77\x8a\xe4\xcc\xd8\xdb\xe0\xca\x2a\xb5\xa9\x81\x44\x8d\x73\xcf\x3a\xc6\xe7\x73\x35\x12\xe1\x6b\x19\x4d\x60\x3b\xd4\x7c\x63\x9a\xb3\x6f\x08\xee\x3d\xdb\x4e\xe7\xe8\x65\x54\x5b\x4b\x6a\x2f\x61\xdf\xee\x5e\xf2\x76\x40\x1b\x1a\xb9\x7b\x7c\xbc\xd4\x8e\x04\xcc\xf4\x3f\x51\xe4\xbf\x02\x00\x00\xff\xff\xb9\x5d\xc3\x84\xbe\x12\x00\x00")

func dataCa_bundlePemBytes() ([]byte, error) {
	return bindataRead(
		_dataCa_bundlePem,
		"data/ca_bundle.pem",
	)
}

func dataCa_bundlePem() (*asset, error) {
	bytes, err := dataCa_bundlePemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ca_bundle.pem", size: 4798, mode: os.FileMode(420), modTime: time.Unix(1494540991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPublic_keyJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\x56\x2a\x28\x4d\xca\xc9\x4c\x8e\xcf\x4e\xad\x54\xb2\x52\x0a\x28\xf1\x0c\xf6\xab\x2a\x0c\xcc\x0d\x0c\x70\xaa\xf0\xcb\x09\x2d\x37\x76\x4e\xa9\x28\x0e\x4f\xae\x8a\x48\xf2\x2c\xaf\x74\xad\xa8\xc8\xc9\x0e\x2a\x8c\x32\x77\x35\xb3\x34\x50\xaa\xe5\x02\x04\x00\x00\xff\xff\x13\xd9\xef\xab\x3d\x00\x00\x00")

func dataPublic_keyJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataPublic_keyJson,
		"data/public_key.json",
	)
}

func dataPublic_keyJson() (*asset, error) {
	bytes, err := dataPublic_keyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/public_key.json", size: 61, mode: os.FileMode(420), modTime: time.Unix(1494540991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataAws_identity_certPem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\x4b\xcf\xba\x38\x1f\x86\xf7\x7c\x8a\x77\x4f\xde\x28\x08\x28\xcb\x52\x4e\x05\x0b\x82\x80\xc0\x8e\x93\xa8\x9c\x11\x2c\xf2\xe9\x87\x79\x56\xf3\xcf\xe4\x49\xa6\x9b\x26\x57\x0f\xb9\x72\xff\x92\xfb\xff\x7f\x2f\x49\xd1\x90\xf5\x3f\xa8\xb8\x1e\x52\x11\x04\x9e\xf2\x43\x29\x8c\x10\x3c\x7a\x10\x82\x61\x0f\xa1\x03\x6f\x73\xf5\x8a\xf9\x80\x4b\xe2\x15\x18\x52\x99\x0d\x8f\xea\x65\xdb\xc0\x91\xb1\x4a\x16\xf8\xc3\xac\x40\x02\x91\x07\xea\xc0\xc3\x6e\x45\x28\x95\x44\x72\xe0\x38\x48\x59\xa4\x30\x0a\xad\x2e\xb9\xf1\x6d\xae\x89\x33\x52\xad\x7d\x14\xba\x35\x76\x01\x91\xcb\x9f\x3b\xba\x42\x72\x2f\xbe\xa9\xfb\x5c\x5b\x6a\x7c\x05\x44\xdf\x38\xb5\x1d\x98\xca\x92\x4b\xe9\x4d\x15\x52\x96\x2b\x03\x36\x78\x6e\x8f\xeb\xac\x8d\xfb\x88\x0d\xd6\xed\x67\xec\xac\xa0\x50\xc9\x7e\xc1\x2f\xb0\x60\xd9\xdf\x76\x9f\xc5\x9e\x91\x6c\x6c\xa5\x6c\xf9\x4f\xf8\x9b\xea\x6f\xa6\xd4\x7f\x55\xfd\xcd\x94\xfa\x53\x75\x0b\x33\xcd\x48\x59\x2a\xef\x7f\x04\x28\x6d\x41\x4b\x3a\x31\xa5\xd2\x81\xaf\xea\x93\x5d\xd9\x34\x65\x02\x87\xfb\x4e\x3b\xbe\xa0\x9e\x0f\xde\xb6\x85\xca\xdc\xb5\xcc\x79\xad\xeb\xf1\x28\x9f\x62\x32\x39\x97\xd3\xdd\x56\xfa\x9e\x57\xd8\xb6\xa4\x65\xc1\xcf\x99\x98\x29\xa3\x67\x3f\xf2\x27\xf3\x75\x68\xdf\xef\xab\xd5\x23\x21\x0d\x0f\x54\xf0\x45\xce\x6a\x1e\xc9\x39\xab\xdb\x7c\x17\x75\xeb\x60\x59\x4d\x89\xbe\x71\x91\x59\x47\xa5\xac\x4d\x11\x79\xba\x71\xbe\xd0\xcb\x49\x9d\xfc\x7e\x3a\x38\xe9\x37\x0a\x8d\xbc\x09\x70\x51\x5a\xc2\x85\x7a\x7c\x9e\xd1\xc4\x1b\xfa\xae\x8d\x6a\xee\xf1\x38\x5c\x12\x46\x37\xf2\x77\x55\x3a\x28\x00\xe7\xc0\x38\x28\x2e\xc3\xd0\x66\xc7\x4d\x17\xa1\x25\x1f\x9d\x3c\x04\x5a\x71\x23\x17\x74\x9a\x04\x10\xf3\xa2\x2a\x7a\xaa\x06\x1c\xe8\x33\x50\xb3\x0f\xd0\xcc\x2e\xf3\xca\xf7\xcb\x28\x1a\x61\xc7\x0f\x5d\xb3\x1a\xb3\x29\x37\xa9\x61\xcd\xa2\xb3\x90\xc3\x08\xba\x29\x4c\xe6\x93\x53\xd0\x58\xca\x8c\x7a\xe7\x53\x8f\xc7\x97\x31\xf5\xa0\x87\x5a\x2d\xde\xe7\xc2\x61\xdf\x02\x3a\xef\x61\x62\xef\xd2\xf9\x9b\xf9\x0c\x7c\x46\x4e\xc5\xed\x4d\x4b\x87\x99\x7e\xb7\x9e\x71\x9a\xd7\x0b\xa3\x88\x63\xef\xf7\xc7\xb4\x55\xa9\xda\x4d\xd8\x0f\xd3\x4e\x38\x3c\x64\x89\x1b\xc8\x79\x3a\x5d\x94\x5b\x93\x2f\x57\x18\xbd\x23\x55\xae\xb8\x26\x1e\xed\xb3\x04\x38\x4d\x01\x60\x9b\x84\x92\x36\xc5\xa7\xe0\xef\xa7\x33\x52\x76\xda\x9d\xc2\x56\x73\x11\x21\xe6\x8b\xee\xe3\xd8\xda\xc2\x3f\xba\xd3\x6d\x90\xe9\xc4\x2b\xd2\x37\x5d\xb1\x53\x2b\xb2\x92\x74\x19\x8a\x78\xe8\x6f\x6e\xc2\x5f\x76\xf4\x6b\xcc\xcd\xa6\x66\x86\x85\xab\x6b\xfd\x46\xe1\x70\x7c\x1f\x50\x89\x52\x81\x7e\xf8\x48\xa2\xaf\xa7\x7c\x3d\xed\x9a\xc6\xde\xa7\xfd\x78\x14\xdc\x2e\x0e\x60\x18\x25\x29\x0b\xe3\x22\x57\xe7\xe9\x38\x64\x87\x9b\xaf\x8b\xb4\xe2\x03\x9d\x6f\x08\xf5\xb9\x16\x32\xb4\x7d\x1c\x39\xee\xd1\x15\xcf\xc8\x8a\x48\x37\xeb\x68\x7d\x0e\x4e\x84\x41\xa5\x49\xdf\x41\xbb\x62\x8e\x93\x00\x96\xcf\x04\x90\x33\x40\xfe\x2d\x94\xea\x2d\x97\xc5\x23\x57\x42\x1d\xf5\xf0\xc0\xe2\x25\x8c\xc6\xf9\x5d\x88\x00\xaa\x92\xa5\x35\x79\xc8\xc6\xd2\x18\x58\xda\x68\x89\x16\x7b\x17\x5c\xbb\xda\x57\xa2\x49\xfd\x34\x80\x62\xc9\xff\x6e\x85\xbf\x02\x00\x00\xff\xff\xa8\x78\xa1\x05\x32\x04\x00\x00")

func dataAws_identity_certPemBytes() ([]byte, error) {
	return bindataRead(
		_dataAws_identity_certPem,
		"data/aws_identity_cert.pem",
	)
}

func dataAws_identity_certPem() (*asset, error) {
	bytes, err := dataAws_identity_certPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws_identity_cert.pem", size: 1074, mode: os.FileMode(420), modTime: time.Unix(1495415324, 0)}
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
	"data/ca_bundle.pem": dataCa_bundlePem,
	"data/public_key.json": dataPublic_keyJson,
	"data/aws_identity_cert.pem": dataAws_identity_certPem,
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
	"data": &bintree{nil, map[string]*bintree{
		"aws_identity_cert.pem": &bintree{dataAws_identity_certPem, map[string]*bintree{}},
		"ca_bundle.pem": &bintree{dataCa_bundlePem, map[string]*bintree{}},
		"public_key.json": &bintree{dataPublic_keyJson, map[string]*bintree{}},
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

