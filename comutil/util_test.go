package comutil

import "testing"

func TestValue(t *testing.T) {
	t.Log(ValueToFloat64(11))
	t.Log(ValueToInt64(11.0))
}

func TestHumpToUnderline(t *testing.T) {
	t.Log(HumpToUnderline("MonMON"))
}
