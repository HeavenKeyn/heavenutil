package comutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
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

func HumpToUnderline(title string) string {
	if title == "" {
		return title
	}
	var build strings.Builder
	st := []rune(title)
	build.WriteRune(st[0])
	for i := 1; i < len(st); i++ {
		if unicode.IsUpper(st[i]) {
			if !unicode.IsUpper(st[i-1]) {
				build.WriteString("_")
			} else if i+1 < len(st) && unicode.IsLower(st[i+1]) {
				build.WriteString("_")
			}
		}
		build.WriteRune(st[i])
	}
	return strings.ToLower(build.String())
}

func ValueToFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("空值")
	}
	switch value.(type) {
	case float64:
		return value.(float64), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	case json.Number:
		return value.(json.Number).Float64()
	case float32:
		return float64(value.(float32)), nil
	case int64:
		return float64(value.(int64)), nil
	case int32:
		return float64(value.(int32)), nil
	case int:
		return float64(value.(int)), nil
	default:
		return 0, errors.New(fmt.Sprint(value, "不是float64类型"))
	}
}

func ValueToInt64(value interface{}) (int64, error) {
	if value == nil {
		return 0, errors.New("空值")
	}
	switch value.(type) {
	case int64:
		return value.(int64), nil
	case string:
		return strconv.ParseInt(value.(string), 10, 64)
	case json.Number:
		return value.(json.Number).Int64()
	case int32:
		return int64(value.(int32)), nil
	case int:
		return int64(value.(int)), nil
	default:
		return 0, errors.New(fmt.Sprint(value, "不是int64类型"))
	}
}
