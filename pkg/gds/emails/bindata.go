// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/deliver_certs.html (1.408kB)
// templates/deliver_certs.txt (1.122kB)
// templates/reject_registration.html (532B)
// templates/reject_registration.txt (414B)
// templates/review_request.html (1.193kB)
// templates/review_request.txt (968B)
// templates/verify_contact.html (656B)
// templates/verify_contact.txt (432B)

package emails

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
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

var _deliver_certsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\x4d\x6f\xe3\x36\x10\x3d\x57\xbf\xe2\x35\xe8\xd1\x91\xb0\x7b\x0c\x54\xa1\xdb\xcd\x6e\x6b\x74\x91\x16\xb1\x5b\xa0\xc7\x31\x35\xb6\xd8\x90\x1c\x96\x1c\xd9\x75\x17\xf9\xef\x0b\xca\x1f\x89\x8d\xdc\xc4\x21\xdf\xcc\x9b\x37\x6f\xd4\xc6\xee\x57\x76\x4e\xf0\xf5\x2b\xea\x07\xf2\x8c\xe7\xe7\x59\xdb\xc4\xae\xaa\xda\xd8\xfd\x2d\x63\xc2\xf2\x71\xbe\xf8\x80\xc0\xba\x93\xf4\x84\xc4\x1b\x9b\x35\x91\x5a\x09\x18\x28\x63\xc5\x1c\x40\x31\x26\xd9\x72\xff\x3d\x26\x88\xa4\x0d\x05\xfb\xff\xd5\xa3\x4d\xa2\xa0\xdc\x57\xb6\xe7\xa0\x56\xf7\x30\x9c\xd4\xae\xad\x21\xe5\x8c\x2d\x39\xdb\x93\xda\xb0\xc1\xbe\xe4\xf0\xec\x57\x9c\xf2\x60\x23\x54\x20\x3a\xf0\x89\xca\xf1\x06\x5b\x4b\xd0\x81\x0f\xd1\xea\x17\x27\x2b\x72\xb8\xb7\x89\x8d\x4a\xda\xd7\xf8\xa0\x4a\x66\xe0\xbe\xe0\x75\xb0\x19\xec\xc9\x3a\x50\x62\xfc\xf1\xdb\xc7\xc5\xbb\xf7\xe0\x60\xd2\x3e\x2a\xf7\x97\x54\x72\x79\x4f\x5a\x78\x54\x86\x02\xac\x8f\x8e\x3d\x07\x7d\x29\x07\xc4\x24\x2a\x46\x1c\xc6\x5c\x38\xfb\xe5\x97\x05\x76\x56\x87\x23\xd5\x93\x5e\x27\xb2\x2a\xe0\xff\xcc\x40\x61\xc3\xd5\x32\xd1\x96\x1d\x1e\x47\xc7\x30\xe2\xa3\xb3\x14\x0c\xc3\x86\xb5\x24\x3f\x89\x56\x9f\x47\xb0\x1c\x18\x31\x59\x4f\x69\x8f\x9e\x95\xac\xcb\x90\xf5\x41\xa2\xfe\xd4\x2b\x38\x68\xda\x4f\x9d\x51\xc6\x5a\x9c\x93\x5d\xbe\x3b\xe6\x18\x5d\x57\x7d\xd7\x3a\xdb\xb5\x59\x93\x84\x4d\x37\xbf\xbf\x6b\x9b\xe3\xf7\x34\xf7\xbf\xe6\xf7\x78\x7e\x6e\x1b\x67\xaf\x5e\x7e\x14\xef\x25\xa0\xf8\xe2\x0a\x72\xb8\x39\x1a\xe6\x0d\xe4\x82\x93\x25\x87\x87\xb1\x74\x7f\x85\x3d\xdc\x1d\xae\xde\x46\x7f\x0a\x7d\x14\x1b\xf4\x0a\x78\x0a\x9f\x41\x6d\x53\x9a\x9b\x64\x12\xf4\x3c\x0d\xf3\x20\xcd\xeb\x79\xce\xd0\xb2\xef\xf6\x32\x62\x67\x9d\x43\xe0\x62\x89\xe1\x6c\x82\x48\x39\xef\x24\xf5\x6d\xc3\xbe\x7b\x19\x7c\x62\xc3\x76\xcb\x3d\x76\x03\x87\x12\xc1\xda\xa6\xac\xc8\xe3\xca\x5b\x2d\x9e\x99\x0a\xbd\xde\x87\x1a\xaf\x68\x94\x0a\x17\xae\x92\x50\x62\x95\x11\xef\x29\xf4\x70\x36\xf0\x6c\xca\x5b\x2c\x36\x66\x46\x6b\xa4\xe7\x4e\x22\x87\x9c\x5d\xdb\x4c\xa7\x37\xe6\x19\x13\x77\x3f\xe0\xf8\x0c\xf1\xc9\xe4\x77\xef\x71\x6b\x03\xe6\x0f\x9f\xe7\x5f\x3e\xd5\xb1\x1c\x65\x54\xfc\xfe\xe7\x72\x0a\x98\xa4\xb8\x0d\xd2\x73\x6e\x9b\x02\x9e\x04\xfb\x2c\x09\x5e\xd2\x85\xeb\x0a\x45\x1b\x94\x37\xe9\xb0\x87\x93\x9d\x5f\x2c\x7f\xb4\xf4\x0c\xd1\x31\x65\x46\x66\x86\x8c\xa9\xea\xc5\x8c\x65\x3b\x0e\x39\x48\xd1\x12\x86\xc4\xeb\x1f\x6f\x06\xd5\x98\xef\x9a\x46\x93\xcd\x45\x03\xad\x03\x6b\x73\xd3\x5d\x9c\xdb\x86\xba\x1a\xf3\xc9\xd3\x18\x68\xcb\xa0\xb0\xaf\xfe\x1d\x39\x97\x7c\xf9\xa0\x91\xa7\x3d\x8c\x04\x25\xa3\x18\xf3\x45\x91\xb2\xd4\x2a\x77\xd4\x7b\x1b\x7e\x9a\x32\xd7\x56\x6e\xba\xcb\x73\x29\x02\x49\xd5\x3f\x62\x8b\xd8\xa5\xd3\x32\xbd\xb7\x99\xde\x96\x36\x73\x24\xc3\x75\x76\x64\x9e\x6a\x23\xfe\xa6\x5b\x94\x4f\x94\x15\x0e\xec\x26\xd2\xe7\x25\xfd\x99\xb3\xe2\x91\x37\x94\xfa\x3c\x6b\x57\x09\x4d\x57\x2d\xcf\xba\x5d\xff\x99\xb0\xe0\xb4\xb5\x86\x0b\xfc\x5b\x00\x00\x00\xff\xff\xdb\x39\xfc\xc0\x80\x05\x00\x00")

func deliver_certsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_deliver_certsHtml,
		"deliver_certs.html",
	)
}

func deliver_certsHtml() (*asset, error) {
	bytes, err := deliver_certsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deliver_certs.html", size: 1408, mode: os.FileMode(0644), modTime: time.Unix(1622505401, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0xbd, 0xb8, 0x63, 0xa4, 0xd1, 0x9c, 0xad, 0x96, 0x2d, 0x7a, 0x18, 0x10, 0xbd, 0x74, 0xae, 0x1, 0x3b, 0xf0, 0x34, 0xcd, 0xbd, 0xac, 0x5c, 0x7c, 0x71, 0x65, 0xfc, 0xb0, 0xc0, 0x71, 0xe}}
	return a, nil
}

var _deliver_certsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x53\xcb\x72\x1b\x37\x10\xbc\xe3\x2b\x3a\x77\x6a\x55\xf6\x51\xa7\x38\x96\x9d\xb0\xe2\x52\x52\x22\x93\xaa\x1c\x47\xc0\x90\x3b\x11\x80\x41\x80\x59\x32\x8c\x4b\xff\x9e\xc2\x2e\x69\x8b\x3a\x6e\xcf\x03\x3d\xdd\xbd\xbf\x70\x8c\x8a\xaf\x5f\x31\x3c\x50\x62\xbc\xbc\xac\x9c\xfb\x4b\xa7\x8a\xed\xe3\x7a\xf3\x01\x99\xed\xa8\xf5\x19\x95\xf7\xd2\xac\x92\x89\x66\x8c\xd4\xf0\xc4\x9c\x41\xa5\x54\x3d\x70\xf8\x01\xf3\x88\xd6\x3d\x65\xf9\xef\x4d\xd3\xbe\x52\x36\x0e\x4e\x02\x67\x13\x3b\xc1\x73\x35\xd9\x89\x27\xe3\x86\x03\x45\x09\x64\x92\xf7\x38\xf5\x1d\x89\xd3\x13\xd7\x36\x4a\x81\x29\xd4\x46\xbe\x50\x39\x57\x70\x10\x82\x8d\xbc\xa0\xee\xe7\xa8\x4f\x14\x71\x2f\x95\xbd\x69\x3d\x0d\xf8\x60\x46\x7e\xe4\xd0\xe7\x6d\x94\x06\x4e\x24\x11\x54\x19\xbf\xff\xfa\x71\xf3\xee\x3d\x38\xfb\x7a\x2a\xc6\xe1\x9a\x4a\xeb\xfd\x64\x9d\x87\xf3\x94\x21\xa9\x44\x4e\x9c\xed\xfb\x73\x40\xa9\x6a\xea\x35\x62\x6a\x9d\x73\xda\x7e\xd9\xe0\x28\x36\x9e\xa9\x5e\xf4\xba\x90\x35\x05\xff\xeb\x47\xca\x7b\x76\xdb\x4a\x07\x8e\x78\x9c\x22\xc3\x6b\x2a\x51\x28\x7b\x86\xe4\x9d\xd6\x34\x8b\x36\x38\xb7\x1d\x19\xa5\x4a\xa2\x7a\x42\x60\x23\x89\x0d\xba\x5b\xb4\x09\x97\x23\xc1\xd9\xea\x69\x3e\x89\x1a\x76\x1a\xa3\x1e\xdb\x9d\x73\xeb\xfb\xbb\xd9\xca\x3f\xd7\xf7\x78\x79\x71\x1f\x35\x25\xcd\xe8\xc6\x2e\xf8\x02\x9c\x8d\x76\x1b\xae\x42\x11\x0f\x53\xa7\xba\x34\x2c\xd0\x82\xf4\x96\x4f\x39\x14\x95\x6c\x4b\xf5\xf2\xd5\x2b\x6e\xab\x08\x3c\xeb\xb8\x90\x7b\x2d\xe5\xaa\x43\x38\x4a\x8c\xc8\xdc\x9d\x18\xbf\x69\x5f\xa8\xb5\xa3\xd6\xf0\x4d\x6a\x54\xf6\x2c\x07\x0e\x38\x8e\x9c\x5d\x47\x76\x52\x9b\xa1\x4d\x4f\x49\xac\xbb\x34\xef\x7f\x9d\xc0\x01\xaf\x5e\xef\xcb\xaf\x7c\xd4\xbc\x60\x9a\x12\xe5\x80\x28\x99\x57\xf3\xde\x6e\xea\xd4\x18\x5a\x38\xb7\x16\xaf\xb5\xbb\x80\xe5\xd9\xb7\x77\xef\x71\x23\x19\xeb\x87\xcf\xeb\x2f\x9f\x86\xd2\x3f\x75\x32\xfc\xf6\xc7\x76\x06\x7c\x35\xdc\x64\x0d\xdc\x9c\xfb\xac\x15\x49\xeb\x95\x8d\x9d\x81\x64\xe3\x7d\x5d\x82\x3d\xe7\xe3\x7b\x86\xce\x19\x59\xa1\x44\xa6\xc6\x38\x48\x13\x83\x4e\xd5\x05\xf5\x53\x0f\xdc\xb2\x85\x0c\xa3\x59\x69\x77\xb7\xb7\x56\xa5\xf5\xdb\x6c\xc8\x6c\xb7\x03\xd6\x73\x22\x30\xd2\x81\x41\xf9\x84\x7f\x26\x6e\x7d\xe6\xac\x7c\xa2\x13\xbc\x66\x23\x6f\x98\x9a\x23\x03\x85\x24\xf9\xc7\x79\xcd\x20\x0a\xad\xf8\x5b\xa5\xab\xd1\xb9\x76\x79\x37\x91\xfc\x33\x7a\x52\x33\x47\xcc\x8d\x37\x9d\x65\x2b\xe4\x79\x68\xbd\x3a\x78\x4d\x83\x73\x3f\x71\x33\x3c\xf2\x9e\x6a\x68\xab\x39\xaf\xcb\x59\x6f\xff\x44\x6c\xb8\x1e\xc4\xf3\xff\x01\x00\x00\xff\xff\x36\x3b\x55\xda\x62\x04\x00\x00")

func deliver_certsTxtBytes() ([]byte, error) {
	return bindataRead(
		_deliver_certsTxt,
		"deliver_certs.txt",
	)
}

func deliver_certsTxt() (*asset, error) {
	bytes, err := deliver_certsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deliver_certs.txt", size: 1122, mode: os.FileMode(0644), modTime: time.Unix(1622505270, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0x1, 0xfe, 0xff, 0xd5, 0xdc, 0x14, 0x89, 0x96, 0xd0, 0x8c, 0x57, 0xe2, 0xa9, 0x63, 0x5, 0x99, 0x90, 0x9f, 0xe1, 0x8b, 0x4c, 0xd2, 0x14, 0xc0, 0xcd, 0x8d, 0x8e, 0xd5, 0xb5, 0x56, 0x4a}}
	return a, nil
}

var _reject_registrationHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x4f\x8f\xd3\x40\x0c\xc5\xef\xf9\x14\x4f\x7b\xe1\x52\x25\xf7\x6a\x18\x01\xaa\x04\xbd\x20\xd4\x2d\xdc\x9d\xc6\x6d\x06\x39\xe3\xe0\x71\x5a\xaa\x55\xbf\x3b\x4a\xe8\xae\xe8\x81\xdb\x8c\xff\x3d\xff\xfc\xc2\x18\xbf\xb0\x88\xe2\xe5\x05\xf5\x57\x1a\x18\xb7\xdb\x2a\x34\x63\xac\xaa\x30\xc6\xef\xf9\xa8\xe6\x53\x26\x67\xb9\xae\xe0\x3d\x63\xbf\xdb\x3e\x7f\x84\xf1\x39\xf1\x05\xad\x92\x75\xe8\xa9\xc0\xf8\x27\x1f\x9c\x3b\x5c\x75\xb2\x7b\xd1\x67\xd1\x96\x04\x9b\x64\x7c\x70\xb5\x2b\x9e\xd9\xce\xe9\xc0\x95\xf1\x29\x15\x37\xf2\xa4\x19\xc6\xbf\x26\x2e\x5e\x63\xdf\x33\x8c\xa9\x68\x5e\x84\xee\x71\x5c\xfe\x9d\x9e\x0a\xa8\xe0\xa8\x22\x7a\x29\xeb\xb0\xac\x39\x49\xac\x80\x20\x29\x86\xe2\xa6\xf9\x14\xb7\x9b\x75\x68\xee\xef\x05\xec\xc7\x76\x83\xdb\x2d\x34\x92\x5e\x4b\xe7\xe8\xee\xaf\xd8\x6b\x22\x34\xf3\xa4\x19\xfb\x9b\x30\x15\x46\x56\x67\x78\x4f\x8e\x94\xa1\xd6\xb1\xc1\x15\xc7\xf4\x7b\xd9\x6f\x34\x6d\x85\x87\x82\x81\xf3\x0c\xc2\x1d\xa8\xd5\x33\xaf\xe6\x13\xbc\x13\x41\x4f\x67\x9e\x1b\xca\xd4\x0e\xc9\x2b\xc2\x03\x36\x8d\x23\x93\xd4\xb8\x6b\x15\xce\x1d\x28\x5f\xb1\x40\x27\xcd\x05\xae\x55\x20\xf4\xc6\xc7\xf7\x4f\x03\x25\x71\x5d\x53\x37\xa4\xfc\xc1\x2d\x15\xaa\x93\x3e\xc5\xc7\x7f\x68\x28\xd6\x6f\xde\x7d\x9a\x8f\xb7\xe3\x13\x59\x57\x56\xa1\x35\x34\xb1\xda\xbf\x19\xf8\x3f\x6f\xe6\xf6\x3f\x01\x00\x00\xff\xff\x80\x82\xbf\x94\x14\x02\x00\x00")

func reject_registrationHtmlBytes() ([]byte, error) {
	return bindataRead(
		_reject_registrationHtml,
		"reject_registration.html",
	)
}

func reject_registrationHtml() (*asset, error) {
	bytes, err := reject_registrationHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "reject_registration.html", size: 532, mode: os.FileMode(0644), modTime: time.Unix(1622505415, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4f, 0x82, 0x13, 0xfe, 0xbe, 0x77, 0x70, 0xf4, 0x98, 0x4b, 0x77, 0xf4, 0x9e, 0x47, 0x94, 0x15, 0x2c, 0xd1, 0x1a, 0x54, 0xa8, 0x60, 0x9d, 0x2b, 0x83, 0x7c, 0xc5, 0xa3, 0x37, 0x19, 0x4c, 0xde}}
	return a, nil
}

var _reject_registrationTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x41\x4e\xf3\x30\x14\x84\xf7\x3e\xc5\xec\xfe\x4d\x94\x03\x74\xf5\x83\x2a\x41\x37\x08\xb5\x85\xfd\x4b\x33\x6d\x8c\x1c\xbf\xf2\xec\xa4\x44\x55\xee\x8e\x1c\x2a\x04\x0b\x76\xb6\xe5\x99\xf9\x66\x1e\x19\x82\xe2\x7a\x45\xfd\x24\x3d\x31\xcf\x95\x73\x2f\xf1\xa8\x96\x87\x28\x99\x61\xaa\x90\x3b\x62\xbf\xdd\xec\xee\x60\x1c\x3d\x2f\x68\x54\xac\x45\x27\x09\xc6\x37\x1e\x32\x5b\x4c\x3a\xd8\xed\xd3\x43\xd0\x46\x02\xd6\xde\x78\xc8\x6a\x13\x76\xb4\xd1\x1f\xe8\x8c\x27\x9f\xb2\x49\xf6\x1a\x61\x7c\x1f\x98\x72\x8d\x7d\x47\x18\x25\x69\x5c\x82\x6e\xef\xb8\xfc\x74\xf7\x09\x92\x70\xd4\x10\xf4\x92\x56\xce\x6d\xd6\xab\x05\xf9\x75\xb3\xc6\x3c\x3b\x57\xce\xdb\x2f\x8f\x72\x7d\x0e\x94\x44\x44\xcd\x44\xee\x24\xc3\x47\xa8\xb5\x34\x64\xc5\xd1\x7f\x2c\x41\x67\xd3\x26\xb0\x4f\xe8\x19\x0b\x11\x5b\x48\xa3\x23\xab\xd2\xe5\x5f\x08\xe8\x64\x64\x11\xa4\xa1\xe9\x7d\x86\xfc\xe6\x97\xf3\x99\x12\x6a\xdc\xb2\x12\x63\x0b\x89\x13\x16\x7a\xaf\x31\x15\xa9\xb4\xbd\x8f\xff\xb3\xf9\x24\xb5\xd7\xda\xb9\xfb\x52\x6d\xcb\x93\x58\x9b\x2a\xb7\xff\x1e\xf6\xaf\xcd\x3e\x03\x00\x00\xff\xff\x09\xfe\x4c\xe1\x9e\x01\x00\x00")

func reject_registrationTxtBytes() ([]byte, error) {
	return bindataRead(
		_reject_registrationTxt,
		"reject_registration.txt",
	)
}

func reject_registrationTxt() (*asset, error) {
	bytes, err := reject_registrationTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "reject_registration.txt", size: 414, mode: os.FileMode(0644), modTime: time.Unix(1622505406, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf4, 0x3f, 0x84, 0x42, 0x3f, 0x96, 0x4c, 0xf, 0x2b, 0x40, 0x1e, 0x60, 0x8f, 0xe8, 0x3, 0xb4, 0x24, 0x6c, 0xdb, 0xba, 0x8b, 0x48, 0x73, 0xc7, 0x59, 0xa1, 0x49, 0xf1, 0xeb, 0x81, 0xa2, 0x7d}}
	return a, nil
}

var _review_requestHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xd1\x8f\xe3\x34\x10\xc6\x9f\xc9\x5f\xf1\xa9\x42\x02\xa4\xb6\x11\x3c\xae\x42\xa4\x85\x95\x60\x01\x1d\xa7\xb6\xdc\x3d\x4f\xed\x49\x62\xce\xf1\x84\xb1\x9b\xa8\x3a\xdd\xff\x8e\x9c\xa6\xd7\x5d\x96\x7d\xba\xd7\xf1\x67\xfb\xf7\x7d\x9e\x71\x35\xd4\xbf\xb2\xf7\x82\xc3\xee\x71\x7f\x8f\x7b\xdb\xbb\x10\xd7\x55\x39\xd4\x45\x51\x0d\xf5\x7b\x46\x47\x23\x43\xd9\xb0\x1b\xd9\x82\x10\x78\x5a\xc4\xbf\x78\x39\x92\xc7\x83\x53\x36\x49\xf4\x0c\xe5\xd6\xc5\xa4\x94\x9c\x04\x28\xff\x73\xe2\x98\xd0\xa8\xf4\x20\xbc\xbb\xdf\xbf\x45\xea\x28\x15\x81\xd9\x46\x24\xc1\x31\x9f\x3b\x3a\x9e\xd8\x6e\x71\xe8\xf8\xba\x45\x14\x1d\x45\x8c\xac\xae\x71\x6c\x91\x3a\x76\x0a\xee\xc9\x79\x90\xb5\xca\x31\x7e\xcb\xf1\x3b\x50\xb0\x4f\xc0\x8a\xb7\xbf\xff\xbc\xff\xfe\x07\x0c\x14\xe3\x24\x6a\xf3\x0d\x96\x8d\x9e\x87\x04\x82\x61\x4d\xae\x71\x86\x12\x63\xea\x9c\xe9\x30\x39\xef\x33\x42\xcb\x81\x95\x12\x5b\x48\xf0\x67\xb8\x06\x67\x39\x81\x86\x41\x65\x64\xa4\xce\xc5\x62\xe1\x7a\x06\x89\xdf\xf6\x7f\xbe\x81\x8b\x77\xd7\xac\x94\xeb\x8f\x1f\xb1\xdd\x2d\xcb\x9f\x3e\x55\x65\xae\xcd\x31\x1e\xe4\xe2\xe6\x0c\x51\x28\xff\xcd\x26\x65\x57\xff\x1b\xd8\x1a\xa7\xc8\xf3\x6a\x23\xde\xcb\xe4\x42\x8b\x9e\x13\x59\x4a\x84\xc9\xa5\x2e\xaf\x15\x95\x11\xcb\x75\x6b\xe3\x12\x61\x55\xce\x05\x18\xe9\x7b\x0a\xf6\x4a\x75\xf2\x75\xf1\x55\xe5\x5d\xfd\xf8\x70\x87\x2a\x26\x95\xd0\xce\x98\xef\x1e\x1f\x66\xc4\xa5\x54\x95\xde\x2d\xca\x83\x7c\xe0\xf0\x5c\x3c\x97\x5e\xca\xab\x32\x1f\xbf\xf8\x23\x63\x78\x78\xe9\xea\xe6\x66\x21\x83\x77\x81\x91\x44\x3c\x28\x2e\x16\x6f\x21\xd6\x17\x5f\x5f\xe3\xe6\x0c\x1b\x87\x1b\x31\x36\x09\x4f\x91\xb0\xa1\xc5\xfa\xe7\x23\x0e\xf2\x5a\xc4\x6b\xe4\x47\x75\x96\x41\x8b\x24\xe7\xde\x73\x8c\xd4\x72\xc6\xc9\x1b\x3c\xc5\x04\xd2\xf6\xd4\x73\x48\x5f\x02\xb6\xc3\xa6\xc7\x6a\xf5\x5f\xbc\x37\x92\x78\x1e\x04\x4c\xfc\x8d\x32\x26\xd1\x0f\xf9\x8d\x93\xc0\x28\xe7\xfe\x24\xfc\xb5\xfb\x03\x1c\xec\x20\x2e\x24\x34\xa2\x97\x27\x27\x74\xca\xcd\x8f\xab\x2e\xa5\x21\xde\x95\xe5\x48\x71\xb0\xd7\xe1\xdb\x06\x4e\xe5\xaa\x7e\x51\xab\x4a\xaa\x11\x5d\xbe\x53\x10\x5d\x3f\xf8\xdc\x86\x97\x64\xb2\x8b\x62\x50\x31\x1c\xe3\x16\xef\x67\x9c\x28\xaa\x67\xd0\x51\x4e\x97\xfc\x5c\x30\x12\x46\x0e\x8e\x83\x61\x48\x33\x4f\x44\x6e\xc8\x4e\x2c\xe8\xa2\x61\x52\x7f\x46\x4c\x39\xc5\x59\xc1\xc5\xe5\x87\x08\x9c\xb2\xbd\xed\x67\xf3\x3f\xe5\xd1\xd8\x71\x4b\x6a\xe3\xba\x3a\x2a\xca\xba\xc8\x43\xf5\xca\x87\xb2\x67\x1d\x9d\xe1\xbc\xfd\xdf\x00\x00\x00\xff\xff\x37\x7d\x8c\xf0\xa9\x04\x00\x00")

func review_requestHtmlBytes() ([]byte, error) {
	return bindataRead(
		_review_requestHtml,
		"review_request.html",
	)
}

func review_requestHtml() (*asset, error) {
	bytes, err := review_requestHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "review_request.html", size: 1193, mode: os.FileMode(0644), modTime: time.Unix(1622505767, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9b, 0x5, 0xcb, 0x2b, 0x14, 0xea, 0xb6, 0x58, 0x7f, 0x6c, 0xcb, 0xd3, 0xd4, 0xca, 0x7e, 0xe4, 0xa4, 0x16, 0x1b, 0xda, 0x1c, 0xbc, 0x85, 0xfb, 0x4b, 0xca, 0x24, 0xa5, 0xbc, 0x5d, 0x27, 0xc0}}
	return a, nil
}

var _review_requestTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x41\x6f\x13\x3f\x10\xc5\xef\xfe\x14\x4f\xd5\x5f\xfa\x83\x94\x44\x82\x63\x6f\x85\x4a\x50\x40\xa5\x4a\x42\x7b\x9e\xda\x93\xac\xa9\xd7\xb3\x8c\x27\xbb\x5a\x55\xfd\xee\xc8\x9b\x2d\xb4\x2a\x1c\xb8\x59\xf6\x78\x7e\x6f\x9e\x9f\x3f\x72\x4a\x82\xed\xfa\x62\x73\x86\xb3\xd0\xc6\x5c\x16\xce\xdd\x30\x1a\xea\x19\xca\x9e\x63\xcf\x01\x84\xcc\xc3\x5c\xf5\x21\xc9\x2d\x25\x9c\x47\x65\x6f\xa2\x23\x94\xf7\xb1\x98\x92\x45\xc9\x50\xfe\x71\xe0\x62\xd8\xa9\xb4\x20\x5c\x9f\x6d\xae\x60\x0d\x99\xcb\xcc\xa1\xc0\x04\xb7\xb5\x6f\x1f\x79\xe0\xb0\xc2\xb6\xe1\xc7\x2b\xa2\x68\xa8\xa0\x67\x8d\xbb\xc8\x01\xd6\x70\x54\x70\x4b\x31\x81\x42\x50\x2e\xe5\x15\x97\xd7\xa0\x1c\x9e\x08\x73\x57\x9f\xdf\x6f\xde\xbc\x45\x47\xa5\x0c\xa2\xa1\x12\x02\x7b\x1d\x3b\x03\xc1\xb3\x5a\xdc\x45\x4f\xc6\x18\x9a\xe8\x1b\x0c\x31\xa5\x2a\x61\xcf\x99\x95\x8c\x03\x24\xa7\x11\x71\x87\x51\x0e\xa0\xae\x53\xe9\x19\xd6\xc4\xe2\x66\x5d\xcf\x44\xe2\xd3\xe6\xeb\x25\x62\x39\x75\xee\xfe\x1e\xab\xf5\xbc\xfb\xf0\xe0\xdc\x56\x8e\xda\x47\x88\x42\xf9\x3b\x7b\xab\x33\xfc\xd1\x9e\x05\x0e\x85\xa7\xd3\x9d\xa4\x24\x43\xcc\x7b\xb4\x6c\x14\xc8\x08\x43\xb4\xa6\x9e\xb9\x7d\x28\xb3\x55\xf0\xd2\xb6\x94\xc3\xa9\x73\x17\xe7\xa7\xa8\xe8\xeb\x8b\xf3\x8a\xdd\xca\x1d\xe7\xe3\xce\xb4\x7c\x94\x42\xde\x73\xf7\x52\xc0\x6f\xf0\xdc\x11\x29\x66\x86\x89\x24\x50\x99\xd5\xd4\xe9\xfe\xc3\x13\xfa\x32\x3e\x41\x62\x69\xcf\x70\x58\xd2\x44\xfc\xcb\xc8\x0b\x54\x4b\x63\x60\xd0\x5c\x52\x7d\x68\xb9\x14\xda\x73\x65\xd6\x0b\x89\x8a\x81\x74\x7f\x68\x39\xdb\x3f\xd2\xd7\x58\xb6\x38\x39\x71\xee\x52\x8c\xa7\xac\x61\xe0\xff\x95\x31\x88\xde\x55\x63\x4d\xe0\x95\x6b\x04\x08\xdf\xd6\x5f\xc0\x39\x74\x12\xb3\x61\x27\x3a\xd1\x7b\x2a\x5d\x78\x8c\xf3\x2a\xb3\xa1\xc4\xda\x4a\x5c\x89\x6d\x97\xea\x93\x1e\xa7\x9a\xe4\x74\x2a\x9e\x4b\x59\xe1\x66\xa2\x14\x51\x1d\x41\xb7\x72\x38\xce\x1e\xb3\x97\xdc\x73\x8e\x9c\x3d\x43\x76\x53\x96\xea\xe3\x36\x12\x40\xe6\x6a\x0d\x93\xa6\x11\xc5\xaa\x03\x53\x05\xcf\x7f\x2b\xb3\x55\xd5\x2b\xe7\xde\xd5\x58\xad\x79\x4f\x1a\xca\xc2\x6d\x7f\x55\xbc\xf8\x7d\x1b\xd6\x3e\x7a\xfe\x19\x00\x00\xff\xff\x57\x36\xd6\xa1\xc8\x03\x00\x00")

func review_requestTxtBytes() ([]byte, error) {
	return bindataRead(
		_review_requestTxt,
		"review_request.txt",
	)
}

func review_requestTxt() (*asset, error) {
	bytes, err := review_requestTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "review_request.txt", size: 968, mode: os.FileMode(0644), modTime: time.Unix(1622505656, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x15, 0x45, 0x4e, 0xcd, 0x96, 0x7b, 0x3e, 0x6c, 0xd9, 0xce, 0xd9, 0x6a, 0xa6, 0x7, 0xe8, 0x32, 0xad, 0x9e, 0x69, 0xab, 0xcf, 0xbb, 0xd4, 0x70, 0x76, 0xf3, 0x3e, 0xc0, 0x11, 0x54, 0xcc, 0xfd}}
	return a, nil
}

var _verify_contactHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xbd\xae\xdb\x3a\x10\x84\x7b\x3d\xc5\xc0\xcd\x6d\x7c\xa5\xde\x50\x84\xfc\x01\xc9\x01\x82\x14\xb6\x93\x7e\x45\xad\x2d\xc2\x14\x57\x59\xae\x6d\x08\x07\x7e\xf7\x80\x22\xe2\x24\x45\x3a\x69\xb9\x9c\xf9\x66\xd8\xce\xdd\x67\x0e\x41\xf0\xfa\x8a\xfa\x2b\x4d\x8c\xc7\x63\xdb\x36\x73\x57\x55\xed\xdc\x1d\x47\x8a\x17\x2c\x72\xc5\x49\x14\xe9\xda\x4f\xde\xcc\xc7\x33\x08\xc7\xfd\xcb\xe1\x1d\x22\xdb\x5d\xf4\x02\xe5\xb3\x4f\xa6\x64\x5e\x22\x94\x7f\x5c\x39\x59\x8d\xa3\xa0\xe7\xb3\x8f\xb0\x91\xa1\x7c\xf3\x7c\xaf\x66\x15\xc7\x29\x6d\x31\x07\xa6\xc4\x68\x09\xa3\xf2\xe9\xcd\x26\x03\x7c\x67\xf5\xa7\xe5\x83\x44\x23\x67\xdf\xf6\x5f\xf0\x78\x6c\xba\xdb\x3a\xcc\x14\x0a\x9e\xc8\x07\xf4\x0b\x5c\xf0\xee\x92\x49\x6c\xf4\xa9\x0a\x3e\x5e\xea\xb6\xa1\xae\x90\xbf\x9c\x56\x66\x47\x31\x8a\x95\x55\x48\xa1\xc8\x9b\x4f\x6f\xd3\x05\x4e\xe6\x05\x14\x07\xcc\x94\xac\x08\x96\x2d\xf4\x1c\xe4\x0e\x1f\x4d\x56\xef\xaa\x57\xb9\x27\x56\xd0\x30\x28\xa7\x84\x9e\x74\xf7\x6c\xaa\x75\x32\x70\xf7\x8f\x0c\x6d\xb3\x9e\xfe\x09\xf7\x9f\x32\x46\xba\xad\x86\x2a\xd7\x3e\x30\x4a\xce\x3c\x59\xa3\xba\xa2\x50\x22\x3f\x91\x7f\x4d\x33\x65\x79\x02\x1a\x26\x1f\x53\x45\xf6\xbb\xca\x7c\xc3\x64\xb7\x9e\xbc\x35\xf5\x89\x6a\x2f\x9b\xee\xef\xff\x5c\x17\x44\x73\x31\xc5\x6e\x9a\xae\xd1\xdb\x52\x3d\x65\x46\xb3\x39\xed\x9a\x66\xbd\xf1\x7f\x7e\xe7\x34\x93\xe3\x3a\x05\x72\x97\xda\xc9\xb4\xe9\x0e\xf9\x13\x6e\xa4\x18\x39\x64\xc5\x7a\x6d\xe3\x3d\x27\xc3\x9e\xcf\xa4\x43\xda\xb6\xbd\xa2\xe9\xaa\xe3\x93\xf8\x53\x90\x9e\x02\x3e\x7a\x65\x67\xa2\x0b\x0e\xac\x37\xef\x38\xf7\xf3\x33\x00\x00\xff\xff\x23\x44\xec\x26\x90\x02\x00\x00")

func verify_contactHtmlBytes() ([]byte, error) {
	return bindataRead(
		_verify_contactHtml,
		"verify_contact.html",
	)
}

func verify_contactHtml() (*asset, error) {
	bytes, err := verify_contactHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "verify_contact.html", size: 656, mode: os.FileMode(0644), modTime: time.Unix(1622506303, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0x8c, 0x42, 0xf0, 0x8b, 0xda, 0x63, 0x83, 0x76, 0x25, 0xcb, 0x6b, 0xde, 0x7c, 0x19, 0x66, 0xd9, 0x15, 0x9f, 0xc5, 0xec, 0x42, 0x6b, 0x28, 0xdc, 0x69, 0xfe, 0xae, 0x2a, 0x8, 0x47, 0x71}}
	return a, nil
}

var _verify_contactTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x90\xcd\xae\x9b\x30\x10\x85\xf7\x7e\x8a\xb3\xeb\x26\xe5\x01\xba\xea\x9f\xd4\x46\xaa\xba\x48\xb8\x77\x3f\x38\x13\x18\x61\x3c\xdc\xb1\x01\xa1\x88\x77\xbf\x32\x28\xd9\x59\xe3\x39\xdf\xf9\xec\xbf\x1c\x82\xe2\xf1\x40\xf5\x9f\x06\xc6\xb6\x9d\x9c\xab\x3b\x8a\x3d\x56\x9d\x70\x57\x43\x9a\x9a\x41\x72\x96\xd8\x82\x50\x5f\xce\xd7\x1f\x88\x9c\x17\xb5\x1e\xc6\xad\xa4\x6c\x94\x45\x23\x8c\x3f\x26\x4e\xb9\x42\xad\x68\xb8\x95\x88\xdc\x31\x8c\x67\xe1\xc5\x8d\xa6\x9e\x53\x3a\x61\x0c\x4c\x89\x31\xb3\xc9\x7d\x2d\x1d\x06\x1e\x48\x02\x9a\x15\x5e\xc7\x15\x14\x6f\x18\x29\xed\x7d\x05\x10\x24\xf6\x68\x38\xe8\x02\x89\x59\x8f\x48\x63\xba\x24\xb6\x6f\xce\x15\xf3\xf7\x1d\xf6\x4b\x63\x26\x9f\xdf\x2e\xff\xb0\x6d\xce\x9d\xef\x65\xf5\x8b\x31\x3a\x9a\x77\x98\xe9\xd4\x84\x67\x75\x99\xec\x28\x7f\xc4\x0e\x8b\x97\xdf\x73\x5a\x0c\x8e\x37\xd3\x6d\x90\x98\x1c\xe5\xe3\xf4\x3d\x9b\x24\xaa\x44\xa1\x06\x8d\x38\x50\xc3\x30\x45\xc9\x2b\xae\x81\x7c\x0f\xdf\x51\x8c\x1c\xb0\xaf\x7e\x2d\x3f\x96\x46\xf2\x5c\xa5\x72\x5b\x79\x1d\x2a\xe7\x7e\x72\xca\xb8\x70\x4b\x76\x4b\x27\x57\xbf\xea\xfe\x04\x6d\x28\xe0\xb7\x18\xfb\xac\xb6\xe2\xca\x36\x8b\xe7\xcf\x00\x00\x00\xff\xff\xc1\xf9\xcb\xc4\xb0\x01\x00\x00")

func verify_contactTxtBytes() ([]byte, error) {
	return bindataRead(
		_verify_contactTxt,
		"verify_contact.txt",
	)
}

func verify_contactTxt() (*asset, error) {
	bytes, err := verify_contactTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "verify_contact.txt", size: 432, mode: os.FileMode(0644), modTime: time.Unix(1622506302, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3a, 0xee, 0x30, 0x2, 0xb7, 0xf9, 0xf4, 0xe1, 0x32, 0x9d, 0x19, 0x22, 0xc8, 0xb9, 0xfc, 0x31, 0x1e, 0x67, 0xe3, 0xc4, 0x5b, 0x41, 0xb8, 0xf6, 0x5c, 0xf7, 0x2c, 0x8d, 0xc5, 0x8c, 0x4f, 0xdc}}
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
	"deliver_certs.html":       deliver_certsHtml,
	"deliver_certs.txt":        deliver_certsTxt,
	"reject_registration.html": reject_registrationHtml,
	"reject_registration.txt":  reject_registrationTxt,
	"review_request.html":      review_requestHtml,
	"review_request.txt":       review_requestTxt,
	"verify_contact.html":      verify_contactHtml,
	"verify_contact.txt":       verify_contactTxt,
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
	"deliver_certs.html": {deliver_certsHtml, map[string]*bintree{}},
	"deliver_certs.txt": {deliver_certsTxt, map[string]*bintree{}},
	"reject_registration.html": {reject_registrationHtml, map[string]*bintree{}},
	"reject_registration.txt": {reject_registrationTxt, map[string]*bintree{}},
	"review_request.html": {review_requestHtml, map[string]*bintree{}},
	"review_request.txt": {review_requestTxt, map[string]*bintree{}},
	"verify_contact.html": {verify_contactHtml, map[string]*bintree{}},
	"verify_contact.txt": {verify_contactTxt, map[string]*bintree{}},
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