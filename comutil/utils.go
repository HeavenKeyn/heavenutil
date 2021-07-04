package comutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strings"
	"unicode"
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

func HumpToUnderline(title string) string {
	var build strings.Builder
	for _, u := range title {
		if unicode.IsUpper(u) {
			build.WriteString("_")
		}
		build.WriteRune(u)
	}
	return strings.ToLower(build.String()[1:])
}

func ValueToFloat(value interface{}) (float64, error) {
	switch value.(type) {
	case float64:
		return value.(float64), nil
	case json.Number:
		return value.(json.Number).Float64()
	default:
		return 0, errors.New(fmt.Sprint(value, "不是float64类型"))
	}
}
