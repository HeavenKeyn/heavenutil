package comutil

import "reflect"

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

type CountFloat struct {
	total float64
	count int
}

func (c *CountFloat) Add(values ...float64) {
	for _, value := range values {
		c.total += value
	}
	c.count += len(values)
}

func (c CountFloat) Total() float64 {
	return c.total
}

func (c CountFloat) Count() int {
	return c.count
}

func (c CountFloat) Average() float64 {
	return c.total / float64(c.count)
}
