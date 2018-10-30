// Code generated by go-bindata.
// sources:
// storage/migrations/mysql/20170523194156-init.sql
// storage/migrations/mysql/20170529135133-custom_capabilities.sql
// storage/migrations/mysql/20170607184935-keys.sql
// storage/migrations/mysql/20170705192449-alter_node.sql
// storage/migrations/mysql/20180417152820-add_unique_key.sql
// storage/migrations/mysql/bindata.go
// DO NOT EDIT!

package mysql

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

var _storageMigrationsMysql20170523194156InitSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xd1\x4a\xc3\x30\x14\x86\xef\xf3\x14\xe7\x6e\x2b\x3a\x10\xa1\x20\x94\x5d\x74\x6b\xd4\x62\xc9\xb4\x6b\x05\xef\x7a\xba\x64\x35\xb0\x26\x21\x49\x1d\x7b\x7b\xe9\x90\xd1\x89\xba\x7a\x17\xc8\xf7\x27\x87\xef\xfc\x64\x36\x83\xab\x56\x36\x16\xbd\x80\xd2\x90\x65\x4e\xe3\x82\x42\x11\x2f\x32\x0a\x95\xd2\x5c\x54\x30\x25\x00\x95\x3f\x18\x51\xc1\x07\xda\xcd\x3b\xda\xe9\xed\x4d\x00\x6c\x55\x00\x2b\xb3\x0c\x12\x7a\x1f\x97\x59\x01\x93\xc9\x75\x4f\x3a\x8f\xbe\x73\xe3\x58\xe4\xdc\x0a\x37\x12\x76\xc2\x39\xa9\x55\xca\x07\x78\x18\xfe\xce\x77\x86\xa3\x17\xbc\x82\x5a\x36\x52\xf9\xb3\xb7\x8f\x80\x15\x8d\x74\xde\xfe\x81\x94\x2c\x7d\x29\x29\x3c\xd1\xb7\xc1\xac\xd3\xd3\x31\x20\x01\x50\xf6\x90\x32\x3a\x4f\x95\xd2\xc9\xe2\x34\xc2\xf2\x31\xce\xd7\xb4\x98\x77\x7e\x7b\x17\x91\x6f\x5a\x37\x68\xb0\x96\x3b\xe9\xa5\x70\x5f\x7a\x7b\xd3\xf1\x7f\x64\xd4\x56\xef\x9d\xb0\x0c\xdb\xc1\x5a\xc2\xcb\x81\x57\x61\x7b\x89\xe3\x32\x66\x87\x7e\xab\x6d\x3b\xfe\x17\x83\x8d\xc8\x34\xf2\xb5\xef\x1b\xd5\x1c\x2e\xa6\xc6\x1a\x1c\xf6\x34\xd1\x7b\x45\x92\x7c\xf5\xfc\xa3\xd0\xe8\xec\xea\x58\xe1\x88\x7c\x06\x00\x00\xff\xff\xd8\x81\xbe\x43\xe8\x02\x00\x00")

func storageMigrationsMysql20170523194156InitSqlBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysql20170523194156InitSql,
		"storage/migrations/mysql/20170523194156-init.sql",
	)
}

func storageMigrationsMysql20170523194156InitSql() (*asset, error) {
	bytes, err := storageMigrationsMysql20170523194156InitSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/20170523194156-init.sql", size: 744, mode: os.FileMode(420), modTime: time.Unix(1496588724, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageMigrationsMysql20170529135133Custom_capabilitiesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x51\x6b\xdb\x30\x14\x85\xdf\xf5\x2b\x2e\x79\x71\xcc\x5c\x08\x63\x85\x81\x97\x81\xea\xa8\xa9\xa8\x2a\xaf\x96\x1c\xd8\x53\xa5\xcc\x4a\x66\x88\x2d\x63\xa9\xcb\xf6\xef\x87\xe7\x78\x4b\xdd\x10\xb6\xbd\x94\xfa\xcd\xe2\x9c\x73\xf5\x5d\xec\x83\x2e\x2e\xe0\x4d\x55\x6e\x5b\xed\x0d\xe4\x0d\x42\x49\x46\xb0\x24\x20\xf1\x15\x23\xa0\xbe\xe8\x46\xaf\xcb\x5d\xe9\x4b\xe3\x1e\x5a\xb3\x51\x30\x45\x00\xaa\xb6\x85\xc1\x45\xd1\x1a\xe7\x14\xac\x70\x96\xdc\xe0\x6c\xfa\x76\x16\x02\x4f\x25\xf0\x9c\x31\x58\x90\x6b\x9c\x33\x09\x41\x10\x75\x06\x67\x3c\x2d\x14\xf4\xcf\x60\xb8\x3c\x63\xa8\x75\x65\x06\xfd\x5f\x19\xbe\xe9\xdd\xa3\xf9\x87\x09\x39\xa7\xf7\x39\x81\x5b\xf2\x19\x94\xee\x51\x84\xf1\xfc\xd7\xd8\xe9\x13\xc0\x68\xb8\x7e\x74\xb8\x56\x88\x42\x20\x7c\x49\x39\x99\xd3\xba\xb6\x8b\x2b\xe8\x86\x09\x22\xe7\x8f\x7e\xf3\x3e\x46\x88\x72\x41\x32\x09\x94\xcb\xf4\xe4\x0a\xcf\xc7\x47\x03\x4c\x88\x00\xa6\x82\x30\x92\x48\x18\x59\x92\x94\x27\x58\x3e\x0b\x5a\xb7\x76\xef\x4c\xcb\x0f\x31\x87\xd7\x95\x69\x5d\x69\xeb\xee\xa4\xd9\x69\xbf\xb1\x6d\x35\x28\x1a\xbd\x35\xcc\xea\x42\xf8\xee\x03\xd8\xfe\x50\x61\x04\x93\xa3\x94\xc9\x28\x14\xb4\xfb\xb3\xea\xeb\x2c\xbd\x83\x63\x3c\xb8\xc1\x2b\xca\x97\xbf\x15\x1f\x3e\x42\x10\x84\xfd\xb2\x53\xfe\xb2\x34\x07\xdb\xe4\x79\xd0\x2b\x64\x3a\x16\x4d\xc6\xa6\xd7\xc8\x33\x3a\x9c\x9c\x12\xfe\x0f\x57\x8c\xd0\x22\x4b\x3f\x9d\xaa\x33\x15\xa3\x8c\x70\x7c\x77\xa6\xeb\xc6\xbf\xaf\x8a\x11\xca\x39\x4b\x93\xdb\xde\x23\x62\xf4\xa4\x3f\x17\x76\x5f\x23\x41\x97\x1c\x33\x10\xf7\x4c\xc8\xae\x49\x83\x77\x97\xb3\xd9\x2c\x00\x41\x24\x54\xc6\x39\xbd\x35\x0f\xde\x7c\xf7\x30\x87\x80\x56\x8d\x75\xae\x5c\xef\x0c\x14\x76\x5f\x83\xff\x5a\x3a\xe8\xd3\x4a\x5b\x07\xf1\xcf\x00\x00\x00\xff\xff\x42\xc8\x60\xda\x9e\x05\x00\x00")

func storageMigrationsMysql20170529135133Custom_capabilitiesSqlBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysql20170529135133Custom_capabilitiesSql,
		"storage/migrations/mysql/20170529135133-custom_capabilities.sql",
	)
}

func storageMigrationsMysql20170529135133Custom_capabilitiesSql() (*asset, error) {
	bytes, err := storageMigrationsMysql20170529135133Custom_capabilitiesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/20170529135133-custom_capabilities.sql", size: 1438, mode: os.FileMode(420), modTime: time.Unix(1496588724, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageMigrationsMysql20170607184935KeysSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x6a\xe3\x30\x18\x45\xf7\x7e\x8a\x0f\x66\xe1\x84\x49\xc0\x33\xb4\x2b\xd3\x85\x12\xab\x41\x54\x91\x13\x4b\x86\x66\x65\x29\x95\x9a\x0a\x62\xcb\x44\x0a\xe9\xe3\x17\x37\x3f\x60\xb7\xdd\x75\xab\xab\x7b\xcf\x11\x28\x9a\x4e\xe1\x6f\x6d\x77\x07\x15\x0c\x94\x6d\x14\x21\x2a\x70\x01\x02\xcd\x28\x06\xd9\x38\x6d\x24\xa0\x2c\x83\x79\x4e\xcb\x25\x03\x69\xb5\x84\x19\x59\x10\x26\x46\xff\x93\x31\xb0\x5c\x00\x2b\x29\x05\x54\x8a\xbc\x22\x6c\x5e\xe0\x25\x66\x02\x56\x05\x59\xa2\x62\x03\x4f\x78\x03\x8f\xa4\xe0\x22\xed\x0f\xbf\xa8\x56\x6d\xed\xde\x06\x6b\xfc\x6f\x01\xe6\x05\x46\x02\x03\x61\x19\x7e\x06\xe9\x83\x0a\x47\x5f\xb9\xd7\xea\xd8\x6a\x15\x8c\x96\x90\xb3\xeb\x8b\x46\x97\x58\x4e\xe4\x35\x1d\x0f\x07\x94\xd6\x07\xe3\xfd\xb9\xd6\xf7\x1d\x7d\xce\xa0\xcb\x85\x2f\x4d\x6f\xbc\xb7\xae\x21\x43\xe4\xed\x78\x9c\x46\x59\x91\xaf\x06\x24\x6e\x02\x53\xb5\xf9\x06\x78\x03\x94\x8c\xac\xcb\xa1\xe1\xcf\xbd\x4e\xb4\x4b\x26\xd0\x13\x9e\x74\x8a\xe1\xec\xf1\x07\x38\xc5\x78\x05\xff\xd2\xa8\xf7\x15\x32\x77\x6a\x22\x4e\x16\x0c\x51\xe0\x6b\xca\x45\x87\x8f\xef\xee\x93\x24\x89\x81\x63\x01\xb5\xf1\x5e\xed\x4c\x15\xcc\x7b\x80\x07\x88\x49\xdd\x3a\xef\xed\x76\x6f\x40\xbb\x53\x03\xe1\xcd\x7a\x38\xaf\x59\xd7\xc4\xe9\x47\x00\x00\x00\xff\xff\x23\x9f\x25\xdf\x69\x02\x00\x00")

func storageMigrationsMysql20170607184935KeysSqlBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysql20170607184935KeysSql,
		"storage/migrations/mysql/20170607184935-keys.sql",
	)
}

func storageMigrationsMysql20170607184935KeysSql() (*asset, error) {
	bytes, err := storageMigrationsMysql20170607184935KeysSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/20170607184935-keys.sql", size: 617, mode: os.FileMode(420), modTime: time.Unix(1497343574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageMigrationsMysql20170705192449Alter_nodeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xc8\xcb\x4f\x49\x4d\x50\xf0\xf5\x77\xf1\x74\x8b\x54\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x48\x4c\x49\x29\x4a\x2d\x2e\x4e\x50\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x32\x35\xd5\x54\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x57\xb7\xe6\x42\x31\xdc\x25\xbf\x3c\x8f\x2c\xe3\x0d\xb0\x9b\x0e\x08\x00\x00\xff\xff\x05\x76\xd0\xcb\xba\x00\x00\x00")

func storageMigrationsMysql20170705192449Alter_nodeSqlBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysql20170705192449Alter_nodeSql,
		"storage/migrations/mysql/20170705192449-alter_node.sql",
	)
}

func storageMigrationsMysql20170705192449Alter_nodeSql() (*asset, error) {
	bytes, err := storageMigrationsMysql20170705192449Alter_nodeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/20170705192449-alter_node.sql", size: 186, mode: os.FileMode(420), modTime: time.Unix(1499289882, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageMigrationsMysql20180417152820Add_unique_keySql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xc1\x4b\xc3\x30\x14\x87\xef\xfd\x2b\x7e\xb7\x6e\xc8\x60\x88\x3d\x15\x0f\xcf\x36\x6e\xc5\x98\xb2\x36\x15\x6f\x26\x5b\xc3\x16\x5c\x9b\xd2\x04\xa7\xff\xbd\xd4\xa9\xd0\xe3\x7b\xf0\xbe\xef\x7b\xd1\x6a\x85\x9b\xce\x1e\x47\x1d\x0c\x9a\x21\x22\x2e\x59\x05\x49\x0f\x9c\x41\xf5\xae\x35\x0a\x94\xe7\xc8\x4a\xde\x3c\x0b\xa8\x77\xf3\xa5\xf0\xa1\xc7\xc3\x49\x8f\x8b\xdb\x24\x59\x42\x94\x12\xa2\xe1\x1c\xf4\x38\x5d\x2a\xdb\xaa\x74\x4e\x39\xe8\x41\xef\xed\xd9\x06\x6b\xbc\x42\xb6\x25\xb1\x61\xff\xc0\x49\x41\x6d\x3b\x1a\xef\xd5\x75\x7a\x9a\x14\x2f\x54\x65\x5b\xaa\xe6\x8a\x34\xca\x2a\x46\x92\xa1\x11\xc5\xae\x61\x28\x44\xce\x5e\x7f\x9b\x4a\xf1\x97\xbb\xf8\x59\x2c\xd3\x68\xf6\x5a\xee\x2e\x7d\x54\x17\x1b\x41\x1c\xf5\x8e\xd7\x72\xe2\xc4\x77\xc9\x7a\xbd\x8e\x51\x33\x89\xce\x78\xaf\x8f\xe6\x2d\x98\xcf\x80\x7b\xc4\x45\x37\x38\xef\xed\xfe\x6c\xd0\xba\x4b\x8f\x70\xb2\x1e\x57\x9a\x75\x7d\x9c\x7e\x07\x00\x00\xff\xff\xf2\x9d\x9b\xd9\x39\x01\x00\x00")

func storageMigrationsMysql20180417152820Add_unique_keySqlBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysql20180417152820Add_unique_keySql,
		"storage/migrations/mysql/20180417152820-add_unique_key.sql",
	)
}

func storageMigrationsMysql20180417152820Add_unique_keySql() (*asset, error) {
	bytes, err := storageMigrationsMysql20180417152820Add_unique_keySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/20180417152820-add_unique_key.sql", size: 313, mode: os.FileMode(420), modTime: time.Unix(1528440940, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageMigrationsMysqlBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x99\x5d\x6f\x1b\xc7\xf5\xc6\xaf\xb9\x9f\x62\x23\x20\x01\xf9\x87\x2c\xed\xec\xfb\x0a\xf0\xc5\x3f\x4e\x0a\x18\x45\x12\xa0\x49\xaf\x3a\x85\xb1\x2f\x33\xea\x22\x14\x29\x93\x54\x3a\xb2\xe1\xef\x5e\xfc\xe6\x1c\x2a\xb2\xe3\xd8\x92\xe3\x06\xbd\xa0\xb8\xbb\xdc\x39\x73\xde\xe6\x79\x9e\x19\x9d\x9f\xa7\xcf\xb6\x93\x4b\x2f\xdd\xc6\xed\xfa\x83\x9b\xd2\xe1\x36\xbd\xdc\x3e\x19\xe6\xcd\xd4\x1f\xfa\xb3\xe4\xfc\x3c\xdd\x6f\x6f\x76\xa3\xdb\x5f\xc4\xeb\xc3\x76\xd7\x5f\xba\xf3\xab\xf9\x72\xd7\x1f\xe6\xed\x66\x7f\x7e\x75\xbb\x7f\xb9\x3e\xcf\x33\xd3\x64\x55\x5e\x98\xae\x34\x55\xfd\x64\xde\xcc\x87\xb3\xfd\xcb\xf5\x43\xc6\x74\xa6\xa8\x4c\x51\x3c\x19\x6f\xf6\x87\xed\xd5\x8b\xb1\xbf\xee\x87\x79\x3d\x1f\x66\xb7\x7f\x98\x89\x3a\x6b\x4c\x5b\x76\x45\xf5\xe4\x67\x77\xfb\xc0\x31\x4d\x56\x99\x2e\x2f\xcb\xee\x49\xbf\x3e\xb8\xdd\x8b\xcd\x76\x72\x0f\x19\xd9\x66\xa5\x69\x4c\x95\xb7\x79\xf6\xa4\x9f\xa6\x17\x37\x9b\xf9\xe5\x8d\x7b\xf1\xb3\xbb\xfd\xe8\xe8\x63\x4e\x2f\xb7\xbc\xf6\xcd\x0f\xe9\xf7\x3f\xfc\x94\x7e\xfb\xcd\xf3\x9f\xbe\x48\x92\xeb\x7e\xfc\xb9\xbf\x74\x69\x7c\x33\x49\xe6\xab\xeb\xed\xee\x90\x2e\x93\xc5\xc9\x70\x7b\x70\xfb\x93\x64\x71\x32\x6e\xaf\xae\x77\x6e\xbf\x3f\xbf\x7c\x35\x5f\xf3\xc0\x5f\x1d\xf8\x9a\xb7\xf2\xf7\x7c\xde\xde\x1c\xe6\x35\x37\xdb\x38\xe0\xba\x3f\xfc\xeb\xdc\xcf\x6b\xc7\x05\x0f\xf6\x87\xdd\xbc\xb9\x8c\xbf\x1d\xe6\x2b\x77\x92\xac\x92\xc4\xdf\x6c\xc6\x54\x3d\xfb\x9b\xeb\xa7\x25\x17\xe9\x3f\xfe\xc9\xb4\xa7\xe9\xa6\xbf\x72\xa9\x0c\x5b\xa5\xcb\xe3\x53\xb7\xdb\x6d\x77\xab\xf4\x75\xb2\xb8\x7c\x15\xef\xd2\x8b\xa7\x29\x5e\x9d\x7d\xef\xfe\x8d\x11\xb7\x5b\x46\xb7\xb9\xff\xfa\xc6\x7b\xb7\x8b\x66\x57\xab\x64\x31\xfb\x38\xe0\x8b\xa7\xe9\x66\x5e\x63\x62\xb1\x73\x87\x9b\xdd\x86\xdb\xd3\xd4\x5f\x1d\xce\xbe\xc5\xba\x5f\x9e\x60\x28\xfd\xf2\xe5\x45\xfa\xe5\x2f\x27\xe2\x49\x9c\x6b\x95\x2c\xde\x24\xc9\xe2\x97\x7e\x97\x0e\x37\x3e\x95\x79\x64\x92\x64\xf1\x42\xdc\x79\x9a\xce\xdb\xb3\x67\xdb\xeb\xdb\xe5\x57\xc3\x8d\x3f\x4d\x2f\x5f\xad\x92\xc5\xb8\xfe\xf6\xe8\xe9\xd9\xb3\xf5\x76\xef\x96\xab\xe4\x73\xf9\x83\x19\xb1\xff\x3b\x86\xdc\x6e\x27\x7e\xeb\xc3\xe1\xc6\x9f\x7d\x8d\xeb\xcb\xd5\x29\x6f\x24\x6f\x92\xe4\x70\x7b\xed\xd2\x7e\xbf\x77\x07\x52\x7e\x33\x1e\xb0\x12\xe3\xd3\x7a\x24\x8b\x79\xe3\xb7\x69\xba\xdd\x9f\xfd\x65\x5e\xbb\xe7\x1b\xbf\xbd\x1b\xa7\x25\x3c\x3e\xbf\x67\x21\xd6\x30\x4d\xb5\x8c\xc9\x62\x3f\xbf\x8a\xf7\xf3\xe6\x50\x97\xc9\xe2\x8a\xe5\x9f\xde\x19\xfd\x6e\x3b\xb9\xf8\xf0\xa7\xf9\xca\xa5\xb4\xc9\x19\x57\xcc\x13\x5b\x65\xe9\xe7\x77\xe7\x5a\xa5\xdf\xf7\x57\x6e\xb9\xd2\x19\x98\x53\xa3\xf4\xf3\x19\xb3\x27\x6f\x3e\x30\xf6\xc7\xf9\x15\x63\xa3\x37\x6f\x0f\xc5\xd1\x0f\x0e\xc5\xd7\xe5\xea\xbe\xe7\x6f\x1b\x20\xb4\x8f\x19\x20\xb8\xe5\xea\xd7\x40\x7f\x63\x41\xa3\xff\x7d\x23\xcf\xf7\xdf\xcc\xbb\xe5\x2a\x1d\xb6\xdb\xf5\xfd\xd1\xfd\x7a\xff\x91\xc8\x6f\xf7\x12\xb8\xdb\xf9\x7e\x74\xaf\xdf\xdc\x1b\xad\x2d\x41\x97\xbf\x50\x40\xf9\xee\x0e\x4f\xbe\x03\x24\xde\x06\xdc\xe7\x9b\xf9\xf0\xe3\xcb\x75\xfa\x54\x5b\x65\x79\x62\x83\xf1\x36\xb4\x83\x0d\x59\x6b\x43\x96\xbd\xff\xe3\xbd\x0d\x5d\x69\x43\x97\xdb\x30\x19\x1b\xca\xde\x86\xb1\xb0\xa1\xc8\x6c\x30\xa5\x0d\x6d\x6d\x83\xf3\x36\xf8\x42\xee\x5d\x63\x43\xed\x6c\xc8\x07\x1b\x8a\xde\x06\x93\xd9\xd0\x1b\x1b\xf2\x4c\xec\x54\x93\x0d\x4d\x69\x43\x3d\xd8\x30\xf1\x9d\xdb\x30\x76\x36\x0c\xfa\x2c\xab\xc4\x5e\xd3\xdb\x30\xf4\x36\xd4\xa5\x0d\x45\x65\xc3\x90\xd9\x90\xd7\x36\xe4\xf8\xd0\xd9\x60\xb0\x33\xc8\xc7\x75\x36\x74\x99\xf8\xd7\x76\x32\x8e\xf1\xa6\xb1\x61\x6c\x6d\xf0\x8d\x0d\x79\x63\x43\xdb\xa8\xaf\xa3\xda\xad\x6d\x68\x0b\x1b\xfa\xc1\x86\xaa\x96\x7b\x53\xdb\x30\x4c\x36\xb4\xd8\xcb\xc5\x6e\x5d\xd9\x50\x3a\x1b\x5c\x61\x43\x9b\xdb\x50\xe6\x36\x18\x62\xf2\x36\x14\xb9\x0d\x59\x6f\x43\x57\xc9\xfb\xd5\x68\x43\x55\x4a\x7e\xf2\x4a\x72\xc8\x6f\x85\xb7\xc1\xb4\x36\x54\xc6\x86\xd1\xd8\x90\x35\x36\x4c\xe4\x72\xb2\xa1\x18\xe4\x9a\x38\xdc\x64\x43\x39\xc9\xb8\x1a\x5b\x6a\x83\x7c\x0e\x85\x0d\xd9\x68\x83\xc9\x35\x3e\x6a\xd3\xd8\x50\x75\x36\x64\xc6\x86\xae\x90\x5c\x36\xf8\xeb\x25\xff\xad\xb7\x61\x70\x36\x34\x85\xf8\x5f\xb5\x36\xb8\xd2\x86\x69\x14\xbf\x8b\x46\xed\xd5\x36\x8c\xb9\x0d\x45\x67\x43\xdf\xc9\xbc\x63\x2f\x7e\x36\xad\xf8\xee\x9d\x0d\x23\xb6\x1a\xa9\x7b\x5f\x48\x8e\x87\x51\xf2\x52\xf5\x92\xc3\x2a\xb7\xc1\x77\xe2\xef\xd0\x88\x0f\xe4\xd3\x54\x36\xb4\x5a\xff\xc9\x89\xbd\xd6\x48\x5f\xe4\xa3\x0d\xcd\x64\x43\xde\xc9\xa7\x18\xa5\x9e\x8c\x27\x5f\x3d\xf7\x85\x7c\x0a\xed\x27\x62\xae\xe8\xd1\x5a\xea\x44\x2d\x88\x99\xd8\x8f\xf5\x20\x1f\x8e\xe7\xa5\xc4\xe7\x73\x19\x4f\x1c\x53\x23\x7d\xd7\xb5\x12\x53\xe3\xa4\x9f\x88\xa9\x33\x36\xd4\x5e\x63\xa2\xa7\x5b\xe9\xc1\xae\x96\x7a\x51\xab\xbe\xb2\xa1\xc9\x6c\xa8\xb4\x5f\x19\x8b\x7f\xde\xd8\xd0\x78\xe9\x31\x7a\xbc\xd2\x35\xd2\x4d\x62\x83\x1a\x4e\x83\xc4\x85\x7d\xfc\x1a\x07\xc9\x45\xc5\x5c\x46\x6c\xd1\xcb\xd4\x8c\xb8\xea\x5a\xfa\x17\x1f\xe9\xd7\x5a\xfb\x86\x1c\xe2\x2f\x7d\x4c\x6e\xe9\x77\x62\x25\xce\x41\xd7\x92\xc1\xaf\xca\x06\x33\xda\x90\x33\x9e\x5a\xd3\xeb\xbd\x3c\x23\x7f\x8c\x99\x74\xde\xb2\x92\xf5\xde\xf0\x5b\x25\xeb\x85\x5a\x4f\x99\x0d\x8e\xfe\x19\x6d\x70\xbd\xf6\x12\x6b\xae\x95\x77\xb3\xfa\x6d\xdc\xe0\x33\xb5\x12\x17\x3d\x58\x16\x32\x3e\xcb\x8f\xef\x9d\x1c\x05\xc6\x23\x20\x4c\x59\xf1\x7d\x6a\xe3\xc8\x9d\xf7\xd4\x4a\xb2\x58\x3c\x06\x1f\x4f\x93\x05\x32\xe8\x71\x1a\xf6\xe4\x34\x59\xac\xee\x08\xf0\x11\xb3\x11\xc5\xff\x45\x42\xbf\x1f\x45\x64\xf4\x3b\xd9\xf4\xf8\xcc\x7c\x4c\xb3\xdc\x49\x8d\x28\x16\x2e\x9e\xbe\x4b\x3c\xaf\xa1\xe4\x8b\xf4\x13\xd2\x90\xc2\xc8\x17\x69\x53\x96\xa7\x29\xdc\x7a\x71\x9f\x7a\x97\x65\x9e\xad\xe2\x73\x18\xf3\x42\x18\xf5\xef\x9b\x39\x2c\x4d\xd9\xd5\x55\xdb\x36\x79\x79\x9a\x66\xab\x37\xc9\xa2\xc7\xab\xaf\x62\x62\x5e\xc7\x6c\x5c\xa4\x9a\x14\x5c\xbe\x88\x7f\xdf\xdc\x95\xbb\x3f\x7d\x38\x1b\xea\x56\xe2\xd9\x6f\x77\x12\x9f\x4a\x8e\x2c\xf0\x48\x6c\x46\x49\x6d\xb8\x47\x8e\x80\x90\x97\x45\x04\x80\xb3\xf8\x1a\x00\x1a\x00\x1a\x85\x2c\xb0\x5f\x17\xf2\x6e\x04\xc3\x46\xbe\x59\x60\x7d\x2b\x60\xcc\x77\xce\xbd\x17\x00\x62\xd1\xb2\xb0\xaa\x42\x40\x08\x5b\x10\x73\x04\x09\xde\x9d\xc4\x26\x63\x01\x16\xaf\x00\x14\x09\x50\x01\xbd\xc4\xd7\x49\x08\x7a\x50\xd2\x23\x0e\xaf\xa4\x04\x68\x76\xa3\x10\x07\xfe\x43\xdc\x2c\x7c\x80\x86\x38\xf8\xb8\x4c\xc8\x0a\xb2\x80\xf4\xab\x41\x08\x2c\x9b\x84\x64\xf8\x86\x2c\x21\xec\xb2\x56\x12\x2f\x05\xbc\x01\x49\x08\x21\xc7\xcf\x4c\x00\x02\x70\xe0\x3d\xe2\xc4\xef\x38\x67\x27\xbe\x02\x84\xc6\x09\x60\x42\x2c\xe4\x9b\x3c\x97\x4a\x90\x7d\x2f\x71\x90\x47\x80\x95\xe7\x00\x1a\x20\x47\xcc\xd4\x03\x12\x69\x14\xc4\x21\x40\xa7\x24\x4b\xcc\x8d\x92\x3f\xe0\x04\x81\x40\xde\x3e\x93\x1c\x40\x18\x80\x5d\x14\x03\x83\x3e\x83\x8c\xf1\xcd\x48\x0e\x21\x5d\x00\xb0\x6e\x84\xbc\xa8\x01\x73\x7a\x88\x07\x32\x20\xc6\x56\x7e\x8b\xb5\x69\x65\x4c\xad\xe3\xfc\x24\x44\x62\x3a\xcd\x45\x27\xf5\x01\x78\x21\x53\x6a\xc7\x7c\x91\xa0\x1b\x11\x35\x5c\x47\xd2\x18\x04\xf0\x19\x4b\x0d\x9d\x93\xfc\xb4\xa5\xe4\x19\x00\x87\xec\xe8\x19\xec\x1a\x84\x42\xa6\xa4\xd6\x2a\xf1\x95\x72\x5d\x69\x0f\x51\x37\xea\x04\xb0\x97\x3a\x6f\xd7\x09\xc1\x0d\xc7\x5c\x0f\xd2\xd7\xd4\x0d\x41\x80\x7f\x79\x2e\xf9\x86\xf4\xb1\x8b\x30\x2c\x9c\xd4\x35\x12\x45\x2e\x39\x8b\xc2\x49\xfd\x8d\x3d\x5a\x6a\xcf\xf4\x92\xaf\xb1\xb1\xa1\x24\xe6\xc2\x86\x72\x94\xb1\xb1\xce\xb5\x88\x0d\x6a\x0f\x49\x95\x2a\x4e\x10\x41\xdc\x47\x31\xd1\x48\xbd\x98\x37\x1b\xb4\x67\x1a\xa9\x31\xeb\x00\x7b\xcc\x99\x79\xa9\x2d\xb9\x25\x3f\x88\x86\xba\x93\xbe\xe3\xdb\x0d\x92\xcb\x1e\x72\xec\xe4\xd9\x80\xa0\x32\x4a\xc0\xf4\x6f\x2b\x3d\x36\xe8\xfd\xa8\xe4\x48\xfe\x7c\x2b\xe3\x89\x61\x52\x01\x85\x60\x81\xe0\xb3\x52\x04\x1b\xf5\x8d\x39\xe8\xd4\x56\x29\xa2\xc4\xab\xe0\xc3\x16\x7e\x20\x90\xa2\x30\x2d\xa4\x7f\xe8\x23\xea\x45\xaf\x53\x03\xc4\x1a\xd8\x81\x7f\xac\x2b\xf0\x2b\xe6\x3d\x97\xfe\x8c\xbd\x80\x00\xc9\x05\x37\xf0\x65\x50\x61\x14\xc5\xe8\x20\xf9\x07\x9f\x26\x15\x9c\x51\x24\xf7\xb2\x2e\x58\xe7\x88\x06\xf2\x3f\xa9\xb0\x66\x0d\xf0\x3b\xeb\xa1\x53\x61\x8a\x5d\xf2\x8a\x90\x69\x47\xb5\x95\xcb\x1a\x2e\x54\x04\xb1\x56\x19\x8b\x9f\xc4\xdf\xea\xba\xa4\xd7\x10\x62\xcc\x41\xcc\xcc\x47\x3c\xe0\x41\xdb\x8b\xe8\x26\xce\x42\x31\x81\xfc\xe1\x37\x62\x12\x1c\x61\x33\x41\x1f\x51\xaf\x28\xb8\xb5\xe6\xac\x2b\x7e\x2f\x15\x5f\x8d\xae\x5f\xf2\xc2\x3a\x8a\xfd\x00\xd6\x3a\xe9\x75\x72\x8c\x1f\x31\x9f\x85\xd4\x9e\xfc\xd3\x7f\xd8\x00\xc3\x10\xf9\xf8\x87\x2f\xc7\xda\x13\x37\xbe\x44\xe1\xd4\x48\x8f\xd2\x5b\x88\xdf\xb8\xae\x26\x59\xf3\xcc\x03\xf6\xd3\x8b\x51\xd0\x97\xbf\xfa\xc9\xdc\xfe\x28\x3c\x7b\x59\x5b\xac\x39\xf0\x9d\x75\x4e\xad\x58\xd3\xac\x91\x77\xb9\x88\x0f\xf5\x26\x16\xf0\x86\x4d\x45\xe7\x64\x53\xf5\x18\xc1\xf5\x41\x96\xfc\xcc\xfa\xeb\x83\x73\x3d\x48\x8e\x7d\xe4\x78\xf0\x11\xea\xec\x83\xbe\xfc\x61\xb1\xf6\x90\xac\xfe\x59\xda\xed\xa3\x39\x53\x29\x67\xca\xa2\xfd\xdf\xd4\x72\x77\x67\xba\x7f\x75\xb7\x9f\x2c\xde\x10\x04\x90\x06\x82\xa7\xee\x75\xb7\x95\x09\xa1\x20\x20\xbc\xee\xb8\x00\x1f\x16\x37\x02\xcb\x29\xb1\x02\x14\x90\x28\x40\x02\x60\xe7\xba\xfb\x43\x80\x00\x84\xec\xd2\x00\x0b\x80\x82\xdd\x24\xa4\x0b\x68\x41\x94\xd8\x81\xf8\xd9\xe9\x42\x3a\x5d\x2f\xc4\x57\xeb\x4e\xb0\x54\x22\x74\xba\x03\x04\xc8\xd8\x85\x16\x47\x21\xd1\x88\x38\x88\x60\x32\xc8\x87\x9d\x1b\xa0\x10\x4f\x26\x5a\xb1\x19\x4f\x2d\x74\x27\x0b\x61\x01\x6a\x00\x08\xc0\x0b\xf0\x40\x3c\xb5\x0a\x95\xdc\x88\xc8\x6c\x74\x67\x9d\x69\x5e\xb0\x85\x40\x01\x04\x8b\x56\xde\x07\x00\x11\x6d\xb9\x0a\x42\x48\x83\xb9\xf0\x3d\x9e\x7a\x14\x4a\x96\x95\xe4\x0a\x61\x02\x41\x1b\x3d\xc5\x41\x0c\xc4\x93\x25\xdd\xcd\x23\x0c\xab\xf1\xd7\x93\x0e\xf2\x02\x68\x91\x3b\x72\x3f\x8c\x42\x22\xf1\x44\x44\x85\x27\xf3\xd4\x2a\xd8\x00\xd7\xf8\x3e\xe2\x28\x97\x3c\x46\xe1\xe0\x45\xec\x72\x8d\x20\x81\x18\x9c\x12\x52\x14\xaa\x5e\x09\x46\x85\x0f\xb1\xf1\x1c\x10\x8f\x62\x6e\xd0\xd3\x22\x2f\x39\x71\x3a\x4f\xa9\xf3\xc6\x38\x8d\xc4\xd7\xe8\x18\xa7\x04\x42\xfd\x10\x30\x80\xfa\xd0\x09\x69\x42\xe6\x08\x01\x7a\x2d\x9e\x86\x8c\x22\x46\x11\x96\x91\xec\x07\xe9\x5b\xec\x77\x2a\x62\x62\x0d\xf5\x14\xcc\xa8\xcf\xd4\x90\xda\x4d\xb5\x5c\xd3\x23\x91\xdc\x3b\x79\x46\xdf\xc6\x13\xb2\x49\x4e\x3a\xa8\x55\x3c\x61\xa8\xe4\xd4\x0a\x32\x27\x26\xf2\x4a\x2f\x41\x48\xf1\x64\xad\x90\xb9\xc8\x01\x79\x83\xd4\xf1\x25\x8a\x2b\x23\x44\x4c\x8c\xd4\x1f\xc1\x9f\xa9\x98\xa0\xce\x88\x54\x7e\x3b\xe6\x1b\xff\x88\xaf\x57\xb1\xd5\x1b\xe9\x45\x7a\x14\xc2\x24\x2e\xd6\x0d\x6b\x82\x9e\x40\x28\xb0\x46\x3a\x15\xa3\xd4\x9d\xf7\xd9\x74\x78\x3d\x31\xa3\xff\xc6\x4a\xe7\xa8\x74\x03\x96\x4b\xfd\xbc\xf6\x35\xc2\x92\x3e\x27\xc7\x51\x78\x38\x15\xf6\xa3\x88\x62\xa7\x22\x1e\x21\xc1\x1a\x6f\x55\xac\x42\xc6\x88\x35\x7a\x92\xf8\x5a\x15\xd7\x88\x2d\x6a\x1f\x63\x34\xda\x7f\x93\x88\x06\xa3\x82\x2f\x9e\xd6\x64\xe2\x63\x14\xcf\x2a\xde\x23\x81\x7b\xcd\x67\x2d\x39\x2f\xa9\xf5\x20\x79\x23\x4f\x31\x27\x93\x9c\x20\x11\x1f\x39\xae\xb4\x5f\xc6\x52\x37\x40\xcd\xfb\xc9\x1e\x51\xd3\xe9\x66\x85\x8d\x26\x6b\xee\x71\xa7\x2b\xef\xc2\xe8\x67\x65\xf7\x77\x8d\x3f\x80\xce\x7f\xfb\xaf\xba\x07\xf3\xf7\xbb\xb3\xfd\x41\xc2\xfe\x9d\xcc\xfc\x39\x0c\xfd\xbe\x34\x28\x25\xd7\xa6\x79\x2c\x23\x37\x45\x59\x54\xcd\x7f\x9b\x91\xef\xfe\x63\xfa\xff\x77\xff\x30\xfd\x54\x5e\x76\xb9\xac\xeb\xa9\x92\xc3\x05\xb0\x01\x1c\x1b\x75\xdd\x45\x0c\xeb\x85\x7f\x4a\x3d\xad\xce\xf4\xe0\x83\xf5\xdd\xe8\x46\x20\xd3\xc3\x96\x6c\x92\x53\x69\xde\xe1\xbe\xd1\x13\x52\xae\x19\xcb\x26\x12\xd1\x1c\xb9\xd7\xeb\x81\xc5\xa4\xa7\xcb\x99\x1c\x80\x34\x2a\xbc\x23\x36\x0d\xc2\x4d\x8d\x9e\x3e\x7b\xdd\x9c\xc7\xd3\xd2\x4a\xd6\x36\x36\xe3\x47\x7d\x84\xd3\xa2\xcf\x93\x1c\x9e\x94\xba\x19\xcc\xf5\x30\x26\xd7\x43\x1c\xe2\x64\x1c\xf7\x93\x6e\x74\xe3\xc9\x76\xa5\xf9\x28\xd5\xa7\x42\x37\x5a\x46\xee\xd9\xe8\xfa\x23\x46\x19\xdd\x50\x19\xd1\x08\x8d\xfe\xe7\x80\xdf\xf1\xb3\xd2\x93\x73\x36\x69\xe0\x3f\xbc\x56\xea\xc9\xf8\x34\x0a\x9e\xc0\x8b\xc5\xf8\x9f\x00\x00\x00\xff\xff\x2b\xa9\x76\x2b\x00\x20\x00\x00")

func storageMigrationsMysqlBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_storageMigrationsMysqlBindataGo,
		"storage/migrations/mysql/bindata.go",
	)
}

func storageMigrationsMysqlBindataGo() (*asset, error) {
	bytes, err := storageMigrationsMysqlBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storage/migrations/mysql/bindata.go", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1528455037, 0)}
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
	"storage/migrations/mysql/20170523194156-init.sql": storageMigrationsMysql20170523194156InitSql,
	"storage/migrations/mysql/20170529135133-custom_capabilities.sql": storageMigrationsMysql20170529135133Custom_capabilitiesSql,
	"storage/migrations/mysql/20170607184935-keys.sql": storageMigrationsMysql20170607184935KeysSql,
	"storage/migrations/mysql/20170705192449-alter_node.sql": storageMigrationsMysql20170705192449Alter_nodeSql,
	"storage/migrations/mysql/20180417152820-add_unique_key.sql": storageMigrationsMysql20180417152820Add_unique_keySql,
	"storage/migrations/mysql/bindata.go": storageMigrationsMysqlBindataGo,
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
	"storage": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"mysql": &bintree{nil, map[string]*bintree{
				"20170523194156-init.sql": &bintree{storageMigrationsMysql20170523194156InitSql, map[string]*bintree{}},
				"20170529135133-custom_capabilities.sql": &bintree{storageMigrationsMysql20170529135133Custom_capabilitiesSql, map[string]*bintree{}},
				"20170607184935-keys.sql": &bintree{storageMigrationsMysql20170607184935KeysSql, map[string]*bintree{}},
				"20170705192449-alter_node.sql": &bintree{storageMigrationsMysql20170705192449Alter_nodeSql, map[string]*bintree{}},
				"20180417152820-add_unique_key.sql": &bintree{storageMigrationsMysql20180417152820Add_unique_keySql, map[string]*bintree{}},
				"bindata.go": &bintree{storageMigrationsMysqlBindataGo, map[string]*bintree{}},
			}},
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

