// Package snap Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// pkg/package-format/snap/desktop-scripts/desktop-common.sh
// pkg/package-format/snap/desktop-scripts/desktop-gnome-specific.sh
// pkg/package-format/snap/desktop-scripts/desktop-init.sh
package snap

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var _desktopScriptsDesktopCommonSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x7b\x7b\x73\xd3\xba\xb6\xf8\xff\xfe\x14\x0b\x27\xc3\x6e\x81\xc4\x6d\xd9\x70\xf8\x75\x37\xfc\x4e\x68\x43\xc9\x90\x3e\x6e\x93\xb2\xd9\x73\xd8\x37\xa3\xd8\x4a\xa2\x1b\x47\xf2\x91\xec\xb6\x39\xa5\xdf\xfd\xce\x92\xe4\x57\xec\x3e\x81\xdb\x19\x68\x63\x2f\xad\x97\xd6\x5b\x4a\xe3\x99\x37\x61\xdc\x9b\x10\x35\x87\x16\x75\x9c\xc6\xe3\x7e\x9c\x06\x0c\x48\xc2\xfd\x39\x95\xe0\x8b\xe5\x52\x70\xa0\x57\x91\x90\xb1\x82\xa9\x90\x40\xf8\x0a\x02\xaa\x16\xb1\x88\x80\x44\x11\x34\x1e\x8d\xdf\x99\x26\xdc\x8f\x99\xe0\x10\x49\x1a\x51\x1e\x8c\x03\x26\x37\x36\xe1\xda\x01\x08\x85\x4f\x42\xb8\x20\xb2\xe3\x36\xb7\xdd\xec\x41\xc0\xf0\xc1\x0e\x3e\x60\x53\xf8\x17\xb4\x02\x70\x9b\x01\x93\x2e\xfc\xfd\x07\xc4\x73\xca\x1d\x00\x00\x7a\x41\x42\x70\x0d\xb3\xd0\x44\x24\xdf\xdc\x6f\x08\xf6\xad\x79\x8d\x1f\x77\x5f\xee\x7e\x6b\xe2\x1f\x37\xdf\x5c\x44\x35\x65\xce\x4d\x81\x1b\x12\xfd\x7a\x66\x2c\x1f\x86\x8d\xdd\x1b\xc3\x5d\x1d\x33\x3e\xe1\x63\x11\x51\x3e\x9e\xb2\x90\x5a\x7e\xe6\x94\x04\xd0\xf2\xb7\x00\xd9\x81\xe7\xef\xc1\x0b\xe8\x85\xc7\x93\x30\x2c\x2d\x4d\xa2\x80\xc4\x74\x7c\x15\xcc\x50\x16\x35\xbe\x20\x61\x42\x55\x49\x26\x45\x2e\xe8\x98\x71\x16\x33\x12\xda\xf7\x9d\x29\x09\x15\xcd\x20\xbe\x1e\x1c\x8e\x0f\xfa\x67\xc3\x8e\x7b\x70\xb2\x7f\x7e\xd4\x3b\x1e\x0d\xe1\xa0\x37\xfc\x3c\x3a\x39\x85\x83\x93\x3f\x8f\x07\x27\xdd\x03\x38\x3a\x1f\xf6\xf7\xe1\xb4\xbf\x3f\x3a\x3f\xeb\x0d\xe1\x4b\xff\xa0\x77\x32\x84\xd3\xf3\x0f\x83\xfe\xfe\xf0\x53\xf7\xac\x07\xa3\xde\xd1\xe9\xa0\x3b\xea\x0d\x51\xc4\x84\x2b\x1a\x6b\xcc\xc3\xd3\xde\x7e\xbf\x3b\xd0\x14\xc6\xa7\xdd\xd1\xa7\xa1\x93\x69\x73\x0a\x6e\xf3\x1a\x81\xf6\x4f\x8e\x3f\xf6\x0f\xc7\x9f\x4e\x8e\x7a\xbb\xad\x26\xfe\xf2\xda\xbe\xe0\x53\x36\xbb\xf1\x12\x45\x65\x0b\xa5\x6b\xe3\x7f\x65\xcd\x37\x40\xcd\x69\x18\xfa\x73\xea\x2f\x40\x89\x44\xfa\xb4\x93\x6b\x0a\x21\xcc\xc3\x27\xd0\x31\xfb\x94\xb1\xfa\x1f\x30\x18\x8a\xe2\xbc\xbc\xba\x29\x71\x53\xa7\xea\x58\x26\x34\x43\x85\x4e\x15\x00\xe3\xd0\x4c\x75\xfe\x07\x04\x42\xaf\xd5\xa6\x87\x4f\x9b\xd7\xc1\x0d\xbe\x72\xed\xe3\x30\xa1\x1d\xb7\x79\xfd\x0c\x6d\xd9\x75\xf4\x43\xcd\x91\xdb\xac\x21\xe7\x42\x07\x90\x62\x89\x2d\xa8\xec\xc3\xcb\xce\x86\x8b\x56\xe9\x6e\x5a\x00\x23\x23\x07\x7c\x1a\x26\xd4\x5d\x5b\x5f\xc5\x30\xee\x1f\xf7\x47\xf8\x41\xef\xa8\xc5\x87\x2b\x53\x8c\x53\xe6\xd8\x5f\xce\x03\x08\xd4\x1b\x4a\x05\xad\x46\x1a\x08\x4e\x4b\x4e\xc0\xd4\x58\x25\x93\x88\xc4\x73\x6b\xf8\xc6\x6b\x37\x24\x25\x21\x3e\xd5\x3e\xb4\x89\xfa\x8c\x88\xa4\x3c\x5e\x7b\xb7\x63\xde\xa1\x46\xaf\x03\x26\x1b\x8d\xa6\x01\xf3\x6e\x5c\x78\xd6\xc9\xfc\x1d\x9e\x3f\x07\x49\xe3\x44\x72\xd8\x82\xef\xdf\xd3\xbf\xb7\x91\x95\x5c\xba\xe1\x71\xf7\x74\x6c\x9d\x67\x7c\x76\x7e\x3c\xea\x1f\xf5\x8a\xc2\x36\x80\x04\x01\xcc\x28\xa7\x92\x84\x80\x1c\x28\xe0\x22\xc6\xa7\x34\x80\xc9\x0a\x14\x27\x91\x2f\xc9\x34\x86\x20\xa1\x10\x0b\x90\x09\x8f\xd9\x92\xea\x17\x0e\x14\x22\x17\x0c\x0e\xc6\x83\xfe\x87\xb3\xee\xd9\x5f\x5a\x5d\xb7\x50\xf7\x42\x36\xf1\xca\x6f\xba\x67\xfb\x9f\xc6\xa3\xb3\xfe\xe9\xa0\x37\x72\x9f\x84\x33\x51\xf2\x71\x78\xef\x43\x36\x61\xdc\x75\xd0\x58\x1a\xf0\xf5\xf3\x07\x30\x3e\xe9\xd8\x90\xfa\xf5\xf3\x87\xd4\x73\xcf\x4e\x4e\x46\x9d\x3b\xf0\xa8\x39\x91\xd4\xfb\xba\xbd\xed\x5d\x2d\x26\x2e\xa2\x3b\x64\x17\x14\xbe\x9e\x44\x94\xf7\x8f\x80\x80\x3f\x27\xdc\xd7\x8a\xc5\xe0\x17\x53\x13\x03\x29\x04\x24\x26\x6d\xa7\x01\xa3\x39\x53\xc0\x14\x48\xfa\xef\x84\x49\x1a\x68\x9f\x8d\xe9\x55\x0c\x8c\x47\x49\x8c\x0b\x2f\x85\x5c\xa0\x13\x0f\x0f\x06\x3b\x30\x23\x4b\xaa\xda\x19\xa7\x83\x93\xfd\xee\xa0\x77\xd0\x3f\x7b\x18\x93\x86\xb8\xe6\x73\x88\xf1\x72\x3f\x91\x4a\x48\xa5\x0d\x23\xc3\xb9\x7f\x7e\x36\x3c\x39\xd3\xfb\x71\x3f\x56\xe6\x0b\xae\x5c\xa7\x90\x6e\x4b\x08\xec\x1e\x78\x28\x2e\x06\xbb\x14\xde\x69\xc0\x11\x55\x04\x06\x6c\x62\x72\x3f\x2a\xec\x70\x00\x2a\x89\x90\x09\xe7\x57\x98\x88\xb7\xa4\x8a\xb8\xbf\x0e\x75\x8b\xce\x42\x2d\xd9\x88\x86\x21\x84\x6c\x72\x38\x80\xcb\x39\x95\x7a\xf7\xa7\x8c\x07\xe8\x94\x10\x48\x76\x41\xa5\x4a\xb5\x3d\xe8\x7f\x38\x1c\x8c\x0f\xce\xfa\x5f\x7a\x36\x0c\xdd\xa5\xf3\x7b\xb8\x08\x24\xbb\x47\xbe\x2a\x39\xb7\xc0\xc9\x97\x6e\x99\x93\x7b\x18\x41\xda\x48\x13\x65\xfe\x53\xc8\x05\x91\x22\xe1\x3a\xdf\x60\xf0\x30\xa6\x1c\x49\x11\x49\x46\x63\x22\x57\xc0\xbf\xb0\x80\x91\x54\x03\xb0\x14\x09\x8f\x55\x51\x29\xc0\xb8\xd3\x00\xef\x82\x18\x02\x1a\x8b\xfe\x6b\x16\x42\x3c\x27\x31\x70\x4a\x03\x85\xea\x9c\x50\x24\xb3\x26\x1f\xb2\xc1\xe2\xb9\x40\xb7\x41\x68\x6b\x53\x89\x62\x7c\x86\xc1\x41\x41\x20\x74\xf4\xd3\x1e\x75\xc9\xe2\xb9\x26\x5e\x66\x0b\x9d\xf2\x8c\x4e\xdb\xbb\x30\x8f\xe3\x48\xed\x7a\xde\x24\x99\xa9\x76\xa8\xeb\xd6\x88\x04\x6d\x4e\x63\xcd\x59\xb4\xf2\x5e\x4e\x92\x99\xb7\xfd\xe6\xdd\xbb\xed\xff\xb7\x73\x97\xde\x6b\x25\x42\xad\x9d\x73\x16\xaf\xfe\x61\x8b\x60\xd8\xb8\xcc\x95\x88\xba\x4b\x19\x28\xd3\x46\x76\xbc\xed\xb7\xaf\xdf\xfd\xbe\xf5\x66\xf3\xd7\x58\x73\xc8\x26\x09\x72\xa6\xad\xf9\x34\x09\x15\x25\x49\xc0\x84\xe5\xf3\xd7\xd0\x8c\x32\x32\x9a\x6a\xef\x70\x00\x17\x94\x07\x42\x02\x96\xa9\x0a\x04\x87\x59\x78\xc1\x03\xa0\x9c\x4c\x42\x1a\x80\x5a\xa9\x98\x2e\x95\xa3\xab\xe4\x5a\x0d\x5f\xf0\xc0\xa3\xb3\x70\x6c\xf0\xb4\x03\x93\x54\xbf\xe9\xd4\x5e\x90\x61\x3c\xee\x1d\x0e\xc6\x5f\x7a\xc7\x07\x27\x67\x99\x34\x58\x15\x3c\x0c\x6b\xe6\xf1\x87\xc3\x58\x52\xb2\xa4\x72\xcd\xeb\x59\xac\x20\x0a\x93\x19\xe3\x99\xd7\x1f\x0e\x47\xe3\xd3\xc1\xf9\x61\xff\xb8\xe8\xf2\x0f\x51\xd3\x4c\x59\x22\xad\xed\xf6\x96\x5b\x83\x6f\xf8\xd7\x70\xd4\x3b\xfa\xd1\x48\xb2\x46\xa6\x01\x33\x15\x5b\x21\x40\xf9\x84\x73\x2a\x21\x10\x54\xf1\xdf\x30\x51\xa9\x98\x84\x21\x3a\x24\xba\x93\x2f\xa4\xa4\x7e\xac\x73\x4a\xee\x44\x33\x16\xcf\x93\x49\xdb\x17\x4b\x2f\x99\x24\x3c\x4e\xbc\xac\xf2\x68\xd9\x9e\xaf\x35\xa7\x61\x44\xa5\xf2\x98\x52\x09\x55\xde\xef\xaf\xeb\xc4\xdb\xef\x1e\x1f\xf7\xee\x4c\x77\x0f\x15\x6d\xbb\xbd\x55\x96\x13\x3f\xb5\x8c\x90\x2d\x2b\xa4\xb6\xc5\xaf\x07\x87\xb0\x6f\x2a\x84\xbb\xcb\x2e\x34\xaf\x52\x26\xcc\x5b\x00\x6d\x4f\xb7\x30\x4d\x63\xdf\xbb\x0a\x66\x6b\x59\xb4\x76\x6d\x0e\xeb\x34\xe0\x80\x4e\x19\x37\x95\x9a\xfa\x0d\xc4\x25\xd7\x55\x05\x96\xa3\x4f\xe0\xf3\xa0\x3b\xea\xde\xc9\x65\x96\xf3\xab\x7c\xae\xaf\x7d\x14\xec\x43\xe1\xd2\x12\xe2\x7e\xd0\xf1\xf9\xb0\x77\xa6\x9f\xe6\x95\x4e\x0a\x86\x3d\x58\x5a\x8d\x85\xe5\xba\xa7\x08\xd2\x59\xc7\xe4\xb5\xf5\x8a\x94\xdb\xe5\x02\x89\xb7\x22\x70\x9b\xa5\x75\xee\x5a\x2e\xc4\x30\x7e\x38\x60\x13\xd8\x83\x9d\xf6\x9b\xd7\xed\x1d\x9d\x7f\x14\x25\xd2\x9f\x63\x56\xc2\xf7\xca\x9f\xd3\x25\x51\x79\x9f\x96\x22\xdb\x75\x1a\x00\xc5\x34\xf4\x1f\x16\x86\xa4\x3d\xe3\x62\x49\xdb\x42\xce\x3c\x35\x17\x97\xe3\x49\x32\x6b\xfb\x33\xf6\xff\x59\xd0\xf9\xc7\xef\xdb\xaf\x5f\xbf\xb9\x53\x3f\x55\x6e\x51\x3d\x3e\xf1\xe7\x14\xa6\x22\x0c\xa8\xbc\x5d\x3b\xfb\xdd\xfd\x4f\xbd\xaa\x7a\xf6\x4f\x8e\x8e\x4e\x8e\xbd\xb6\x46\xe2\xea\xe6\xc4\x4e\x2e\xd6\x55\x68\x20\xd0\xfa\x9e\x41\x8b\x5a\x6e\x72\xac\x2e\xfc\x5d\xec\x5c\x30\x98\x98\x25\x68\xd2\xd4\x8f\x85\x5c\x41\xa2\x68\x60\x0b\x00\x15\x0b\x2c\x98\x13\x8e\x4c\xaf\xd1\x7a\x05\x4b\x36\x93\x58\x6f\xb3\xd8\x01\x58\x5e\xdc\xce\x4d\x8d\x28\xba\x35\x58\xdb\xe2\x02\x9b\x99\xd6\x74\x4c\xb8\x5f\x6d\xf9\x10\xa0\xc6\xac\x0c\x92\x8a\x45\x15\x16\x69\x7a\xfb\x92\xa2\x34\xfa\x9d\xf5\x48\xdc\x52\x6c\x74\xd1\xa2\xe8\x15\x53\xb1\x82\x0d\xa3\x19\x49\x97\xe2\x82\x06\x98\x87\x38\x0c\x4e\x77\xa1\xb1\xfd\xf6\xcd\xdb\xd7\xbf\x6f\x61\xa7\x31\x65\x57\x34\xd8\x2c\xf4\x90\x6b\x28\x8b\xed\xe3\x1a\x53\x45\x28\x07\xc0\x9f\x2f\x45\x00\xff\xd8\xda\xaa\x7b\x6d\x9a\xab\x1e\x57\x09\xa6\xc2\x39\xd5\x13\x3d\xcc\x87\x2a\x6b\x81\x30\x7c\x31\x6c\xa9\x15\x6c\xd8\xfe\x27\x7d\xa9\x5a\x98\x51\xd2\x42\x4f\xe7\x17\x64\xba\x58\x79\x9c\xec\x3f\xa4\xe2\x28\x74\x3c\xfd\x29\x10\xbe\x7a\x05\x0b\x4a\x23\x88\x25\xf1\x17\x20\xa6\x36\x57\x63\x88\x0f\x98\x54\x70\x89\x9f\x94\x80\x4b\x0a\x3e\xe1\x10\x89\x98\xf2\x98\x91\x30\x5c\x65\x16\x65\x92\x1c\xc7\x17\x10\x92\x98\x4a\xa7\x7e\x2c\x66\xcd\x24\x89\x20\x1f\xf5\xbc\x00\x21\xb1\xc5\x86\xab\x60\xd6\xca\x1e\xb7\x0c\x02\xbd\x99\x94\x06\x34\x70\x74\xa1\xab\x11\x9a\x57\x76\x84\x96\x3f\x96\x34\x14\x24\xa8\x3c\x0e\x19\x5f\xa4\xf3\x36\xc7\xce\x6d\x8c\x6b\xe9\xd1\xc2\x5a\x74\x34\x99\xe0\x59\x61\xa2\x51\x63\x7d\x29\x86\xcc\x28\x30\x66\x4d\x31\x5a\x95\x27\x58\x85\x8f\x46\xe9\xd9\xac\x29\x1f\xbe\x15\x27\x60\x5e\x73\xba\x3e\x96\xd1\x6e\x5a\x81\xa9\x71\x08\x03\x5e\xd1\x86\x1d\x7e\x15\x07\x37\x68\x86\x6c\x5a\x9e\x79\x82\xdb\x3c\xeb\x75\x07\x25\x42\xeb\x53\xbf\xe7\xcf\x1f\xb3\xc6\x5a\x59\x26\x8c\xd5\x7c\xc9\x32\xf7\x4f\x8e\x4e\x4f\x8e\x7b\xc7\xa3\xe1\xf8\xb8\xd7\x3b\x18\x9f\x9f\x1e\x74\x47\x3d\x17\x3a\xe0\x22\xdf\xb8\x19\xdf\xbf\xc3\xbf\xa0\xb9\x2e\x56\xdd\x8c\x0d\xa3\xa0\xeb\xfd\x77\xc3\x7b\xa6\x1a\xdf\xb4\xca\x1a\xcd\xeb\x94\xc1\x9b\xc6\xcc\x7d\x88\x8c\xef\xab\xaa\xad\x99\x4a\x02\xf8\x11\xb4\xc8\x83\x14\x70\xdb\x5e\x3d\xda\x64\x00\x96\xc1\x1b\x95\x2c\x61\xaf\x8e\x2e\x5a\x45\x1d\xf3\xcd\x69\xdb\x2c\x33\x54\xb5\x01\xdc\x65\x28\x53\xe6\x50\x33\x99\xae\xf8\x9b\x85\x58\xf7\x2c\xfd\xd8\xda\xd4\x83\xb6\xea\x96\xd0\x50\xc3\x93\x71\x5a\x13\x36\xf7\xf5\x80\x99\x4d\xd3\x30\x44\xfc\x38\xd1\x31\x48\x52\x12\x14\x03\x90\xc2\x10\x46\x89\x3f\xc7\x80\xa2\x2b\x3f\x54\xf5\xc6\x06\x83\x0e\x6c\xfd\x01\x0c\xf6\xa0\x79\xdd\xa8\x1f\x77\xfe\xeb\x9f\x7f\xdf\xfc\x01\xec\xe5\xcb\xcd\x4d\xab\x77\x36\x85\x67\x15\xb3\xaf\x8c\xa1\xed\xea\x26\xfb\xfb\xc6\x2d\x98\xe4\x6d\x2a\xbc\x45\x89\xf8\x62\x22\x29\x59\x98\x7d\xd0\x7b\x65\x82\xb4\x69\xf2\x31\x22\xdb\xb0\x48\x78\x50\x1f\x2f\x6d\xd6\xc3\xc2\x09\x79\x7a\xa5\x23\x2b\x8b\x2b\x9b\x63\xc1\xd3\xcd\xd1\xce\x2d\x96\x4b\xc4\xdb\xba\xa8\x47\xfd\x3e\x1b\xeb\x67\x22\xd6\xc1\xdd\xbe\xc3\x76\x23\x4d\xda\xd6\x72\x6b\x2f\x28\x3b\x40\x85\x53\x03\x58\xb1\xa2\x1f\xd9\x54\x80\xc9\xda\x04\xfa\xce\x1d\x85\x56\x4b\xd2\x90\xc4\xec\x82\xb6\x62\xd1\xb1\xc1\x7f\xd3\x2d\x44\xf2\x52\x30\x6c\x4e\xd6\xa3\x78\x7e\x72\x75\x0b\x00\x80\x5c\x62\x0a\x2f\x00\xec\x14\x4f\x9c\x4a\x13\x7d\x8b\xcf\x16\x8c\xb7\x62\x0c\x39\xb4\xd4\x3a\x5f\x39\x7c\xe5\x94\xc0\x46\x07\xeb\xfe\xda\xee\x2e\x29\x10\x49\xb1\xa3\xf5\x71\xd7\xb0\x3c\xe7\xf4\xd2\xec\xdd\x2b\xf0\x0b\x2e\x39\x27\x17\x79\x09\xa0\x08\x96\x5a\x8c\x83\x08\x03\x33\xe4\xd5\x15\x0d\x1a\x17\x56\x61\xa6\x00\x7d\xc8\x06\xd6\x6d\x9d\x08\x83\x4e\xdd\x86\x95\x8e\x43\xcc\xc6\x59\x57\xbb\xac\x85\x5f\x87\x33\x7b\x34\x00\xb7\x29\xc2\xc0\xd6\x01\x76\xd3\x38\xbd\xcc\x1e\x18\xbb\x09\x50\x05\x16\x74\xd3\x56\x12\x06\x6a\x3d\x7b\xb7\x2e\xb8\x85\xf3\x5e\x58\x20\x0f\x76\xde\x97\xb7\x96\x86\xb9\x85\xdc\x43\x5d\xbf\x2e\xd0\x83\xe7\xcf\xb3\x1d\x87\x8d\x52\xd9\xa2\x41\xab\x15\xce\xf7\xef\x70\x37\x98\x29\xf9\xdd\xcd\x27\x89\x52\xae\x33\xb4\x15\x05\x34\xa6\x7e\x0c\x97\x64\x15\xa2\x09\x28\x2a\x2f\xa8\x04\x25\xfc\x05\x8d\x5f\x69\x1a\xa0\x68\x0c\x94\x5f\x30\x29\xf8\x52\x5b\x90\xc0\xe2\x38\x64\xa9\xe9\x44\x92\x4e\xa9\x74\x1a\x29\x92\x57\x60\x30\x61\x29\xe9\x8b\x65\x44\x62\x50\xab\xa5\xde\x94\x8d\x84\xc7\x2c\x44\x9b\x4c\x94\xfe\x27\xed\x58\xb5\x0d\x67\x74\x49\x97\x13\x2a\x5f\x99\x09\x46\xb9\x6b\x50\xe0\xc9\x84\xeb\x04\xee\xed\x25\x2c\x78\xaf\xc7\x31\x6d\xad\x17\xe4\x27\x14\x62\x91\x0e\x74\xcc\xf1\x54\xde\x83\x39\x0d\x73\x48\x31\xa7\x56\xac\x36\x7c\x14\x12\x05\x27\x2c\x54\xd8\xb4\xa6\x2d\xeb\x54\xc8\x64\xd9\xce\x06\x3d\x6d\x26\xbc\xd8\xb3\x42\xb5\x02\x4c\xe6\x2d\xfc\x0b\x03\xab\x3d\x76\xd2\x67\x04\xdb\xef\xde\x7a\xdb\x5b\x4e\x03\xba\x45\xad\xe8\x91\x6e\x20\xd0\x3f\xed\x29\x41\xa6\x64\xae\xc3\x15\x96\xe8\x64\x65\xda\x0a\x0a\x07\xfd\x61\xf7\xc3\xa0\x37\xfe\xb3\xfb\xd7\xa0\x7b\x7c\xe0\x34\x74\x73\x44\xf8\x0a\xb8\xe0\x2d\xba\x8c\xe2\x95\x39\xed\xdc\xc4\x36\x43\x2d\x58\x64\x28\x84\x62\xc6\x7c\xc0\xa2\x5f\xd2\x70\xd5\x76\x4a\x35\x9c\x45\x36\xee\x7e\xe9\xf6\x07\x88\xde\x66\x6d\xdb\xf3\xd6\x76\x53\xcf\x9f\x43\xeb\x3f\xe0\x36\xd7\x18\x2a\x35\xbb\x00\x97\x01\x53\x51\x48\x56\x1d\x37\xd5\xcf\x96\xbb\x76\x92\x99\x12\x3f\xe8\x0f\x4f\x07\xdd\xbf\xaa\x31\x30\xc7\x51\x81\x75\x0a\xd1\xcf\x12\x18\xe3\xe6\xa1\x57\x74\x2a\x4c\x7b\xed\xb6\xd7\x4c\xb1\xb9\xe5\x45\x9c\x44\xb7\x2c\x5a\x5b\x61\x18\x1f\x82\xdb\x5c\x27\x58\xe5\xbc\x81\xd0\x32\xe1\x1c\x83\xae\x69\xe8\x33\xdb\x47\xc3\xd6\x31\xd4\x82\xda\xce\x3a\x93\xb0\xf7\xe1\xfc\xb0\xb3\x9d\xbd\xbf\x67\xbf\xb2\xda\xc3\x90\xf5\x4d\x7e\x36\x15\x55\xc9\xb3\xd0\xc4\xb9\xb8\xcc\x60\x8b\x59\x68\x5d\x19\x55\x81\x8a\x29\xa9\x2a\x7e\x0d\x86\x6c\x65\x9e\xa2\x4c\x48\x39\x22\x0b\x6a\x66\xf0\x5d\x3d\x83\x37\x2e\x07\xe4\x82\xb0\x90\x4c\x42\xdd\x1e\xb3\xc0\x88\x80\xd8\x5a\x2a\xa2\x3e\x9b\x32\xbf\x32\x2b\x78\x58\xcb\x0f\xa0\x27\xf1\xc6\xa5\x3a\xae\xfe\xe0\x99\x4f\x6e\xfe\x5a\x0f\xea\xef\xb3\xa0\x02\xa2\x75\x93\xa8\x41\x52\x55\xa2\xdd\xe9\xd3\xf3\xc1\xb0\x37\x1e\xf6\xce\xbe\xf4\xce\x3a\x6e\xc2\xd9\xd5\x6e\xf3\xba\x06\xc1\x8d\x5b\xd6\xdc\x61\x1f\x24\x8d\x84\x62\x3a\x68\x3d\x6e\x22\x7a\xd8\x1f\x8f\xfe\x3a\xed\x0d\xfa\x1f\x7e\xf4\x5c\x63\xc6\x72\x26\xcc\x30\xfd\x97\x72\x52\x25\x77\x3f\xb2\xa7\x89\xf1\x08\xbc\x3f\xb6\xf8\x7f\x54\x0d\x02\xa7\x01\x9f\x29\x8d\x00\x9b\x23\x29\xc9\x0a\xfb\xa0\x74\xfa\xad\x5e\x69\xef\x0d\x85\x88\x30\x9e\xc4\x73\x29\x92\x99\x3e\xeb\x5b\x3a\xfd\x8f\xc3\xce\x6f\xbb\xbf\x99\x46\xaa\x25\xb1\xb1\xc5\x65\xa6\x74\x37\x98\xf6\xf6\xf6\x8a\x83\x52\xac\x9f\x34\xbd\x8f\x82\xc7\xf6\x1c\x40\x27\x64\xc4\x47\xb3\x23\x9d\x8f\x27\xc7\x23\xdb\x8b\xde\x75\xf6\x42\x63\xdf\x9b\x0a\x1e\x2b\xb7\x66\xe1\xc7\xfe\xa0\x77\xef\x42\xf3\xbf\x6e\x85\xdd\xc2\xad\x94\x25\x59\xd0\x31\xe6\xf3\x31\xbe\xb7\xa3\xc9\x6b\x07\x80\xfa\x73\x01\xee\x5e\xfe\xf4\x7d\xf9\x7a\x59\xde\x5d\x17\x26\xdc\x96\xc7\xf2\xb5\x33\x8d\x08\x60\x2f\x60\xf2\xfd\x5d\xcb\xf6\x3c\x84\xb0\x37\x9a\xea\x49\x3d\x02\x7d\x1d\xca\xba\x72\x7a\x6d\x1b\xeb\xaa\xe9\x9c\x93\xeb\x75\x68\x2c\x8d\xeb\x84\x5e\xe7\xeb\x8e\x85\x39\x8f\xc5\xc2\xd0\x22\xf8\x0d\x60\x8f\x71\x3f\x4c\x02\x0a\x6c\xc6\x85\xa4\xe3\x25\x53\x8a\xf1\x59\xc7\x5d\x51\xe5\xbe\xc7\xbd\x69\x07\x7b\x9e\x05\x7a\xff\x9b\x6e\x4b\xfe\xa4\xba\x1d\xc6\x12\x25\x5d\x1d\xcf\x99\xc2\x22\x87\x24\xa1\x1d\xdc\xa3\x13\x4d\x99\x54\xba\x92\xd4\x45\x0c\x3e\x66\x7c\xa6\x51\x5c\x0a\xb9\x50\xbb\xfa\xb8\x5b\x24\x31\xb0\x18\xbd\x23\xb3\x90\x4b\x16\x86\x10\xcb\x95\xbe\x68\x22\x59\xac\x8f\x2d\x31\x97\x60\xbf\x68\x0a\xca\xb9\x58\x9a\x1e\x29\x23\x86\xd6\x3f\xa1\x30\x09\x31\x17\xe9\x6b\x44\xdd\x28\xea\xca\xa5\x90\xed\xa2\xbc\x19\x3c\x16\xb4\xec\xaa\xe3\x5e\x05\x33\xf7\x7d\x4e\x7c\xcf\x4b\x21\xb4\xb4\x75\x76\xa2\x01\xbc\x7c\xc9\x2d\x26\x93\xe1\xb9\x63\x69\x81\x5a\x66\x47\xd6\x37\xbc\x92\x73\xdc\x64\x93\xd3\x47\xce\xef\x32\xce\xe4\x12\x5a\x72\x5a\x39\x6d\xf1\xae\x73\x3a\xaf\xb4\xc9\x98\xff\x5b\x2f\x5e\x99\x58\xf2\xaa\x6d\x7e\xdf\x38\x5a\xdf\xfa\x6e\x50\x61\xaf\xa6\x92\xcc\x74\xc3\xc0\x54\x3e\x14\xc7\x2a\x9d\x64\x7d\xa7\xd9\x7e\xa6\xf4\x7a\x6b\x31\x7a\x83\x74\x79\xa0\x4f\xcc\x8b\x18\xcd\xaf\x44\xea\xb5\xbb\xc0\xc9\x12\xab\xe8\x78\x6e\xf6\x3b\x0f\x39\xc6\x38\xbd\x37\x5b\x7a\x00\xa2\x63\x8f\x3e\x96\x6f\x57\xcf\x08\x8a\xf3\xb8\xc2\xbe\x21\x60\x5d\x84\xaa\x1b\xe3\xe5\xaf\xcb\xb1\x2e\x3d\x14\x32\x4a\xca\x2a\x35\x22\xa9\x1d\x9b\x9b\xe3\xb6\xd1\x67\x78\xdd\xde\x7e\x67\x4e\x3f\x74\x03\xa3\xcd\x0f\x98\x3e\xad\xf6\xe7\x84\xcf\x68\xa0\x91\x19\xd0\x9d\x2d\xdd\xee\xa0\x4e\x63\x51\x3a\x30\x83\xcb\x39\xf3\xe7\x30\x27\x0a\x35\xc5\xa9\x1f\xd3\x34\xec\xa3\xe4\xba\xc0\x9b\x56\xae\x38\x19\x00\xb7\x7a\xd8\x96\xae\xe0\x77\x2c\x59\x3f\x1e\xb2\x6f\x6c\x49\xf3\x21\x61\x61\x00\x4b\xb6\xa4\xc6\xbc\x9d\x46\x51\xf2\x59\xbc\xd0\xae\xf9\xef\x18\x98\x2f\xf8\xaf\xb1\x62\x0f\xa9\x67\x19\xe4\x99\x99\xeb\xdf\x73\x49\x0c\x97\x78\x39\xd7\x65\x27\x66\xd3\xe2\x30\xce\x4c\xd3\x5a\x4b\xdd\x0b\x92\x98\x4c\x88\xaa\x1d\xc6\xa5\x73\xe9\x56\x24\xa9\xee\xb1\x3b\xd8\x3e\xaa\x98\x2c\x23\x05\xad\xe0\xec\x61\x4c\xd5\x6e\x12\xfe\xd4\xb2\x71\x9b\x26\xd2\xa4\x94\x15\x9e\x4c\xc0\x52\x04\x49\x48\xcd\x18\xc8\x9c\x61\x6e\x18\x77\xc4\x8a\x64\xa6\x68\x1c\x33\x3e\x53\x16\x6c\x33\xbb\xe8\xd0\x3f\x19\x1f\x9d\x1c\x9c\x0f\x74\x11\xdd\x59\x3f\x79\xf4\x66\x4c\xb4\x2c\x66\xf7\x47\xb7\x37\x3f\x95\x79\x72\x69\x1b\xb2\x49\x6b\xa7\xbd\xa5\xd9\xfa\x77\x42\xe5\x2a\xe5\xad\xb4\xc1\x99\x1d\x95\xa5\x33\x9a\x2b\x44\x8f\xba\xd7\x25\x17\x7b\x5a\xf9\x2d\x3c\xcb\x95\xe7\xbe\x68\x2b\xf1\x10\x42\x3f\x13\xf1\x2f\x51\x6f\x1d\xa9\xcc\xfc\x86\xd9\xa0\x88\x61\x7e\xc8\x8d\xcd\x5c\x34\x70\x0e\x87\xe3\xe1\xfe\xa7\xde\x51\xb7\x60\x63\xb9\x51\x67\x54\xed\xb5\x04\xb7\xf0\xc5\x01\x83\x72\x9c\x5e\x58\xb8\x2e\x59\xd1\xf6\x6d\xbb\x5e\x24\x57\xdd\xf4\xea\xdb\x27\x16\x77\x60\xe5\x1b\x9b\xeb\xd8\xf5\x85\x5a\x55\xba\xe2\xcc\x1a\xd9\xcd\x91\x78\x33\x0b\xd4\x4e\x75\x59\x37\xb8\xd0\x49\x3a\xbf\x9f\x40\x42\xec\x2c\x56\x36\x67\xd8\x2d\xb0\x78\xb2\x55\xbe\xe0\x31\xe3\xd9\x20\xa2\x3c\xe9\xd6\xcd\xe1\x46\xa8\xa0\xd5\x2d\xb1\xe3\x7a\x2f\xda\x57\xcb\xb0\x38\x83\xdc\xbc\x7d\x0e\x5e\x5d\x57\xab\xeb\xc7\x10\x17\x17\x54\x4a\x16\xd0\x27\x71\x90\x2d\xbe\x9b\x8d\xec\xc8\xae\x01\x27\x3c\x5c\xa5\x2a\xcc\xef\xc8\x98\x23\x31\x11\x31\x1a\x00\xe1\xab\xd8\x96\xb9\xb5\xec\x97\x09\xdd\xa2\x86\x87\x0a\xa6\x0d\xbc\xde\x5e\x59\xfa\x1d\x9b\x1f\x8c\xc8\xeb\xfe\xf5\x33\x42\x07\xfe\x61\xf1\xb6\x72\x8f\xb6\x77\x33\xf4\xfc\x28\x8f\x0f\xba\xda\x37\xf5\x91\x72\x1a\x59\x7d\xa5\x2f\xe8\x66\x25\x1f\x2a\x1a\x3b\x02\x7d\x1b\x10\x2b\xd1\xbc\x32\x42\xc7\x45\xeb\x77\x0e\xb0\x9a\x43\xce\x46\xb6\x94\xc9\xef\xa7\x57\x6e\xbe\x78\x81\xa9\xee\x8a\xe5\x44\xcd\x7a\x3d\xaa\xae\x4f\x5f\x95\xe3\x61\x8d\xb1\xb2\xa0\x14\x76\x6a\x28\x14\xf2\xc0\x7d\x48\x6f\x5d\x9f\xc5\xe0\x11\x96\x22\x13\x16\xb2\x78\xf5\xf0\x3b\xed\x5e\x9c\xaf\x7a\xc0\x3d\xf5\x22\xf8\x9d\xdf\x8c\xf8\x59\x98\xec\x05\x7a\xac\x70\x82\x45\x2b\x62\x57\x93\x64\x0a\xa1\x20\x41\xe1\x3e\xfb\xe1\xc1\xe7\xf1\x69\xff\xeb\x87\xf3\x8f\x69\x7a\xd2\x43\x8d\x4a\x21\x93\x21\x68\x59\x04\xb6\xa0\xbd\x0d\x0d\x9a\xd0\xd3\xdd\x21\xa7\x86\x4e\xb1\xd3\xde\xde\x6a\x6f\x79\x29\xe7\x3f\xa1\x4c\xd6\x79\xae\x56\x72\xf7\xe7\x54\x5a\x65\x01\x0a\x1f\x75\x5d\x90\xea\xb0\x6c\xee\xff\x07\xe4\xde\xdf\x29\x77\xe6\x0d\x7d\x5f\x77\xa6\xba\x6b\x33\xdb\xfc\x6b\x5a\x13\xfb\xe5\x92\x4a\x4f\x5a\x03\xf1\xc4\x22\xc3\x1e\x7c\x2d\xf5\xd7\x0f\x6e\x2b\x32\x34\x11\xcf\x7d\x51\x28\x4d\x72\x13\xd0\xab\x3d\xc6\x03\x7a\x65\x9a\xbb\xec\xb8\xf3\x59\x09\xc0\x17\xbc\xa5\xff\xac\xeb\x99\xf0\x47\xbf\xb4\xc5\x4e\x8d\x90\x5e\x73\x03\x9b\x16\x6c\xe9\x53\xa4\xe9\x29\x7e\xf1\x38\x23\x48\x5f\x8e\x2b\x5f\x33\x35\x3f\x05\x55\xe6\x70\x4e\xf5\x94\xc3\x50\xd0\x67\xa5\xb5\x70\x0f\x70\x02\x35\x61\xdc\xb3\x9d\x97\x96\x5f\x4b\x5e\x19\xc8\xa5\xf9\xf8\xf1\x78\x6e\x61\x2d\x3d\x90\x7e\x02\x73\xed\x59\xbc\xd8\xf9\x49\x0c\x5a\x5c\xb7\x30\x99\xd5\x69\x85\x3f\xcb\x45\x53\xe1\x1c\xfa\x70\xf4\xd9\x9a\xa9\x19\xd5\xcd\xc9\x05\x13\x12\xbb\x4d\x36\x65\xfa\x68\x79\x34\x17\xca\xdc\x2d\x62\xcb\x88\xf8\x71\x3e\x58\x01\xca\x67\x8c\x53\x73\xd3\x76\xb2\x82\xff\x8a\x81\x28\xb8\xa4\x61\xe8\xcc\xe2\xc5\xd8\xa4\x43\xd5\xd9\x98\xc5\x8b\xd6\x6b\xac\xa4\x6d\xf9\xd0\x66\x9c\x41\xfa\x70\x22\xc4\x62\x49\xe4\x42\xe9\x27\x3a\x88\xc4\x0b\x2c\x1f\xfc\xb9\x10\x8a\x4a\x04\xde\x74\xb2\x6b\x61\x6e\xf3\xba\x80\x1b\x1d\xcf\xb5\xce\x13\x50\x15\x77\xea\x6e\x7a\x15\x47\x10\x03\x70\x9b\x08\x78\x6b\xba\xdf\x08\x98\xb4\x9e\xa0\xe1\x36\xef\x49\xf6\xe6\xda\xa1\x06\x2d\x5f\x4f\xb2\xe7\x85\x6a\xb5\x9c\x88\x90\xf9\x60\x0a\x24\x01\x6c\x92\xa8\xf4\x68\x4e\xdf\x34\x40\xd1\xf4\xc3\xd8\x1e\xa6\x27\x91\xfe\x22\x88\x85\xd1\xdf\x69\x71\x1a\xb0\xa1\xa8\x3e\x79\x5c\xea\xbb\x5c\x8d\xd7\x7a\xbf\x1a\x6f\x41\xf0\x3b\xbf\xfd\xf3\xe6\xdd\xd6\xef\x6f\x5f\x6f\x3a\xfd\x0f\xe7\xc3\x71\xf9\xd4\x61\x5d\x53\xc8\x44\xe9\x76\xf1\xfa\x1a\xd7\xb1\x53\xd7\xf5\x17\x1e\xae\xb4\x5f\x44\x4d\xe3\x6d\x2d\x8c\x93\x8f\xb5\x2a\x9a\x44\xf2\x06\x51\x1d\xe5\x34\xdd\xdf\x7d\x68\x8b\x60\xd4\x07\xb7\xf9\x4f\xf7\x7f\x03\x00\x00\xff\xff\x18\x6c\xf5\xe6\x81\x40\x00\x00")

func desktopScriptsDesktopCommonShBytes() ([]byte, error) {
	return bindataRead(
		_desktopScriptsDesktopCommonSh,
		"desktop-scripts/desktop-common.sh",
	)
}

func desktopScriptsDesktopCommonSh() (*asset, error) {
	bytes, err := desktopScriptsDesktopCommonShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "desktop-scripts/desktop-common.sh", size: 16513, mode: os.FileMode(420), modTime: time.Unix(1560800605, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _desktopScriptsDesktopGnomeSpecificSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\xdd\x8f\xd2\x40\x14\xc5\xdf\xfb\x57\x5c\xa7\x3c\x19\xfb\x21\xfb\x68\x6a\x9c\x6d\x67\xa1\xa1\x5f\x5b\x06\x3f\x62\xcc\xa4\x94\x29\x1d\xb7\xb4\x38\x33\xcd\xc2\x7f\x6f\x20\x28\xeb\x02\x6e\x8c\xae\x7d\xbd\xbd\xf7\x9c\xf3\x3b\x69\xcd\x17\xce\x5c\xb4\xce\xbc\x50\x35\x58\xdc\x30\xcc\xdf\x3e\x86\x09\x23\x3a\x81\xa6\xe8\xdb\xb2\xe6\x12\xd4\x9a\x97\xa2\x12\x25\xac\x0b\xa9\xc1\x7c\x6a\xdb\x10\x15\x7c\x06\x34\x98\x26\x38\x63\x01\x99\x4e\x68\x9a\xb1\x0f\xf8\x53\x84\x93\x80\xe1\xf7\x38\x8c\xf0\x75\x44\x10\x78\x80\xb4\xec\x39\x82\x2f\x6f\x40\xd7\xbc\x35\x00\xf8\x66\xdd\x49\x0d\xa3\x60\xc2\xae\xb1\x3f\x21\x49\xe0\xa1\xfb\x62\xdb\x14\xed\x02\x1d\xc7\x7e\x34\xa3\x94\xe4\x67\x5f\x31\x21\xe8\xb8\x82\xb6\xd3\x50\xf7\x52\x83\xee\x0e\xf6\xb7\xa0\x6b\xa1\xa0\x50\x70\xcf\x9b\xe6\x15\x7c\xed\x95\x06\xd1\x42\x59\x28\x7e\x3c\x7d\x4b\xd9\x6d\x86\x59\x16\x61\x7a\x93\xe6\xb1\x77\xb8\x6c\xf1\x65\x63\x54\xc2\x30\x7e\x18\xa4\x13\x96\x61\x3a\xf6\x1e\xa5\xcc\x67\x09\x0d\x63\xe2\xf4\x4a\x3a\x8d\x98\x3b\xbf\x4e\x71\xee\x8f\x19\xcd\xc3\x2c\x22\xd4\x59\xea\x3b\xeb\xca\x76\x91\x61\x98\x20\xe6\xbd\x82\xa2\x5d\x40\x55\x0a\xbd\x01\xd1\x6a\xbe\x94\x85\x16\x5d\x6b\xec\x94\xc2\x98\xc5\x69\x30\x8b\x08\x0b\xc2\xdc\x1b\x7c\x0c\x46\xcc\xc7\xfe\x98\xb0\x71\x1a\x13\x47\xac\x56\xdd\xa2\x6f\xb8\x7a\x68\xee\xb8\x72\x13\x46\xc4\x1b\x9c\x9c\x39\xae\xd9\x65\x51\xd6\xfc\x6c\x69\x7e\x1a\x67\x69\x42\x12\x3a\x65\x09\x21\x01\x9b\x65\x01\xa6\xe7\x9b\x93\x2b\xb0\x64\x05\xe8\x54\x69\xd7\xca\xea\x6e\x21\x24\x58\xeb\x4b\xf3\xbd\xb8\xb5\x79\xac\xff\x07\x38\x1b\x31\xdf\x13\xb5\xdc\x3d\xd9\x6f\x3d\x97\x5b\xeb\x67\xc6\x3d\xe9\x07\x6e\x01\x9a\x16\x2c\x55\xfd\x85\xe0\xa1\x3f\xe7\xca\x76\x6d\xf7\x48\x13\x39\x2f\x6d\xd5\x5d\xca\x09\xcf\x9c\xf0\xed\x89\xf0\xae\xfe\xfd\xa7\xd3\xfc\x2b\xc6\x43\xdb\xbd\x60\x61\xf8\x2c\x90\x87\xb6\xeb\x0c\xed\xd7\xff\x95\xf2\x13\x19\x2f\x63\xae\xc4\xe1\x37\xc1\x4b\x40\x83\x77\xe8\x7b\x00\x00\x00\xff\xff\xd7\xce\x8a\x81\x79\x05\x00\x00")

func desktopScriptsDesktopGnomeSpecificShBytes() ([]byte, error) {
	return bindataRead(
		_desktopScriptsDesktopGnomeSpecificSh,
		"desktop-scripts/desktop-gnome-specific.sh",
	)
}

func desktopScriptsDesktopGnomeSpecificSh() (*asset, error) {
	bytes, err := desktopScriptsDesktopGnomeSpecificShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "desktop-scripts/desktop-gnome-specific.sh", size: 1401, mode: os.FileMode(420), modTime: time.Unix(1560800605, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _desktopScriptsDesktopInitSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x7f\x6b\xdb\x30\x10\xfd\x5f\x9f\xe2\xaa\x98\xb6\x81\x25\x86\xd1\x85\xd1\xcd\x65\xa6\xf6\xda\xb0\x34\x09\x49\x3a\x06\xa5\x18\x45\x3e\xc7\x62\xb2\x15\x24\xb9\xcd\x20\x1f\x7e\xc8\x69\x7e\xf5\x67\xf4\x87\x41\xba\x77\xef\xde\x9d\x9e\xdc\x38\xf2\xa7\xa2\xf4\xa7\xcc\xe4\xd0\x42\x42\x1a\xcf\x17\x69\x40\x8f\x55\x25\xcf\x51\x83\x28\x85\x85\xc6\x2b\x18\x32\xee\x87\xc3\x24\x8a\xc7\xbf\x26\x83\x61\x72\x39\xb8\x19\x0e\xfa\x71\x7f\x32\x4e\xfa\x71\x1c\x25\xb7\xc3\x28\x9c\xc4\x01\xb5\xba\x42\x4a\x48\x03\x4c\x8e\x52\xf2\x1c\xf9\x5f\x30\xaa\xd2\x1c\x03\x3f\xc5\x07\xbf\xac\xa4\x24\x6d\xa0\x5e\xcd\x76\x3b\x8e\x47\x49\x14\x4e\x42\xbf\x2d\x99\xb1\x89\xc6\x07\x61\x84\x2a\x29\x7c\xbe\xd8\xc0\x61\xb9\x04\x47\x4b\x44\x06\x77\xeb\xcc\xb5\x8e\x5e\x38\x9e\x24\xa3\xf8\x77\x77\xdc\x1d\xf4\x29\x04\xeb\xf8\xf6\xe8\xfe\x1b\xd8\x1c\x4b\x02\x70\x50\x03\x19\x93\x06\x29\x41\x69\x90\x00\x20\xcf\x15\xd0\xb7\x2b\x06\xcf\xab\x5d\x7c\xd8\x1a\xc9\x84\x9b\xcf\x18\x2d\x78\xa3\x38\xec\x5d\x0f\x6e\x62\xb0\xca\x89\x84\xca\xa0\x36\xa0\x91\x49\xc8\x55\x81\x90\x0a\x8d\xdc\x2a\xfd\x8f\xac\x91\x01\xf5\x4e\x67\x68\xb1\xb4\x30\x67\xc6\x3c\xa6\xe0\xdd\x76\x23\x58\x02\xaf\x2c\xb4\x52\x38\x39\x3f\x81\x56\x06\x9d\x66\x7d\x09\xdd\x6c\x43\x0b\x39\x33\x50\xa8\x54\x64\x02\x53\x77\x2a\x74\x7d\xde\x4a\x85\x36\x60\xd0\x5a\x51\xce\xcc\x27\xc8\x94\xe6\x08\xac\x84\x6a\x9e\x32\xbb\x1a\xfa\x9d\xa3\xa4\xde\x9f\xe8\x2a\xb9\x1c\xf4\x7f\x76\xaf\x12\x27\xc5\xdf\xa4\xb7\xeb\x4f\x91\x7e\x31\x55\x41\xe1\xf8\xf8\x23\xb8\x54\x9c\x49\xdc\x24\xdc\x6f\xaf\x68\x55\x8d\x7a\xa7\xab\x18\x7c\x07\xba\x19\x92\xdf\xe6\xaa\xcc\xc4\xec\x59\x5d\xda\xa4\x70\xe4\xee\xfd\x94\x33\x7b\xb0\xca\x26\x85\xe5\x92\xc0\xd3\x3a\xb8\xe0\x4a\xf9\xe1\x25\xf7\x3b\x6d\xee\xf5\x7a\xa0\x21\x57\x2f\x0a\x20\x13\xb5\x73\x76\x1f\x41\x38\xba\xbc\xa6\x10\x04\x40\x59\x91\x76\xce\x76\xbd\xee\x42\x01\x5d\x7c\xed\x24\x9d\xb3\x96\x14\x65\xb5\x68\xcd\xca\xca\xf9\xfa\x0d\x02\x5d\xe4\xd9\x4b\x02\xa6\x8b\x6d\x36\xb2\xa9\xc8\xb3\x77\x39\x5e\x13\xc1\x98\xe6\xf9\x61\x2a\xe6\x73\xde\x39\x43\xf9\x92\x63\xae\x1e\x51\xbb\xa0\xc4\x7d\x9e\xfa\x95\xae\x30\x5b\xb2\x5d\x88\x9b\xd9\xde\x9c\x1d\x20\x99\x8c\xba\xc3\x5e\x3c\x09\xa8\x57\x17\x7f\x1a\x6b\x6d\x5b\x07\xf6\xa5\x98\xba\x7f\x66\x6a\x71\x61\x53\x55\x30\x51\xb6\x8d\xda\x95\x85\x8b\xb9\xd2\x16\x7a\x51\x32\x1c\xc5\xbd\x41\x18\x05\xd4\xdb\x6e\xce\xdf\x63\xa9\x25\x3d\xe5\xaf\xad\xb6\xde\x1f\xe0\x88\x57\xa1\xbb\x4d\x39\x72\xe4\x40\xbd\x1f\xf4\x7f\x00\x00\x00\xff\xff\xc0\xd6\xa2\xf0\xfa\x05\x00\x00")

func desktopScriptsDesktopInitShBytes() ([]byte, error) {
	return bindataRead(
		_desktopScriptsDesktopInitSh,
		"desktop-scripts/desktop-init.sh",
	)
}

func desktopScriptsDesktopInitSh() (*asset, error) {
	bytes, err := desktopScriptsDesktopInitShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "desktop-scripts/desktop-init.sh", size: 1530, mode: os.FileMode(420), modTime: time.Unix(1560800605, 0)}
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
	"desktop-scripts/desktop-common.sh":         desktopScriptsDesktopCommonSh,
	"desktop-scripts/desktop-gnome-specific.sh": desktopScriptsDesktopGnomeSpecificSh,
	"desktop-scripts/desktop-init.sh":           desktopScriptsDesktopInitSh,
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
	"desktop-scripts": &bintree{nil, map[string]*bintree{
		"desktop-common.sh":         &bintree{desktopScriptsDesktopCommonSh, map[string]*bintree{}},
		"desktop-gnome-specific.sh": &bintree{desktopScriptsDesktopGnomeSpecificSh, map[string]*bintree{}},
		"desktop-init.sh":           &bintree{desktopScriptsDesktopInitSh, map[string]*bintree{}},
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
