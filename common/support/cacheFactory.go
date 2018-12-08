package support

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/astaxie/beego/cache"
	"time"
)

var Cc cache.Cache

func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if Cc == nil {
		return errors.New("cache instance is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("set cache error caught: %v\n", r)
			Cc = nil
		}
	}()
	timeouts := time.Duration(timeout) * time.Second
	err = Cc.Put(key, data, timeouts)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	if Cc == nil {
		return errors.New("cache instance")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			Cc = nil
		}
	}()

	data := Cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}
	// log.Pinkln(data)
	err := Decode(data.([]byte), to)
	if err != nil {
		//fmt.Println("获取Cache失败", key, err)
	} else {
		//fmt.Println("获取Cache成功", key)
	}

	return err
}

func DelCache(key string) error {
	if Cc == nil {
		return errors.New("cache instance")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			Cc = nil
		}
	}()

	err := Cc.Delete(key)
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		//fmt.Println("删除Cache成功 " + key)
		return nil
	}
}

// --------------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
