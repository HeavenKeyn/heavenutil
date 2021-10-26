package comutil

import "testing"

func TestValueToFloat64(t *testing.T) {
	t.Log(ValueToFloat64(11.0))
	t.Log(ValueToFloat64(11))
	t.Log(ValueToFloat64("11"))
	t.Log(ValueToFloat64("11.0"))
	t.Log(ValueToFloat64(""))
}

func TestValueToInt64(t *testing.T) {
	t.Log(ValueToInt64("11"))
	t.Log(ValueToInt64(11.0))
	t.Log(ValueToInt64("11.0"))
	t.Log(ValueToInt64(""))
}

func TestHumpToUnderline(t *testing.T) {
	t.Log(HumpToUnderline("MonMON"))
}

func TestGetMapKeys(t *testing.T) {
	t.Log(GetMapKeys(map[string]interface{}{"aaa": "aaa", "bbb": 1}).([]string))
}
