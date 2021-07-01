package comutil

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

func LoadProperties(path string, out interface{}) error {
	doc, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(doc, out)
	if err != nil {
		return err
	}
	return nil
}

func StructToMapByTag(obj interface{}, tagName string) map[string]interface{} {
	var m = make(map[string]interface{})
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	for i := 0; i < obj1.NumField(); i++ {
		tag := obj1.Field(i).Tag.Get(tagName)
		if tag != "" {
			m[tag] = obj2.Field(i).Interface()
		}
	}
	return m
}
