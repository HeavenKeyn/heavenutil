package errutil

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("test error")
	err = Error("错误为：", err)
	t.Log(err)
}
