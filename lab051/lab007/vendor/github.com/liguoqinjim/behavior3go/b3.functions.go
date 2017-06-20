package behavior3go

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
)

func CreateUUID() string { // This function is used to create unique IDs for trees and nodes.
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	s := base64.URLEncoding.EncodeToString(b)

	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type RegisterStructMaps struct { //节点map
	maps map[string]reflect.Type
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{make(map[string]reflect.Type)}
}

func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	fmt.Println("New ", name)
	var c interface{}
	var err error
	if v, ok := rsm.maps[name]; ok {
		c = reflect.New(v).Interface()
		fmt.Println("found ", name, "  ", reflect.TypeOf(c))
		return c, nil
	} else {
		err = fmt.Errorf("not found %s struct", name)
		fmt.Println("New no found", name, "  ", len(rsm.maps))
	}
	return nil, err
}

func (rsm *RegisterStructMaps) CheckElem(name string) bool { //检查节点是否存在
	if _, ok := rsm.maps[name]; ok {
		return true
	}
	return false
}

func (rsm *RegisterStructMaps) Register(name string, c interface{}) { //注册节点
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}
