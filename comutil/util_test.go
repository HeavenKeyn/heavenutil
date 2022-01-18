package comutil

import "testing"

func TestLoadProperties(t *testing.T) {
	err := LoadProperties("", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestValueToFloat64(t *testing.T) {
	t.Log(ValueToFloat64(11.0))
	t.Log(ValueToFloat64(11))
	t.Log(ValueToFloat64("11"))
	t.Log(ValueToFloat64("11.0"))
	t.Log(ValueToFloat64(""))
	t.Log(ValueToFloat64(19143692e0))
}

func TestValueToInt64(t *testing.T) {
	t.Log(ValueToInt64("11"))
	t.Log(ValueToInt64(11.0))
	t.Log(ValueToInt64("11.0"))
	t.Log(ValueToInt64(""))
	t.Log(ValueToInt64(1.636714431914e+12))
	t.Log(ValueToInt64(19143692e0))
}

func TestHumpToUnderline(t *testing.T) {
	t.Log(HumpToUnderline("MonMON"))
}

func TestGetMapKeys(t *testing.T) {
	t.Log(GetMapKeys(map[string]interface{}{"aaa": "aaa", "bbb": 1}).([]string))
}

func TestMsUnixToTime(t *testing.T) {
	t.Log(MsUnixToTime(1636967988881))
}
