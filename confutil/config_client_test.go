package confutil

import (
	"testing"
	"time"
)

func TestType(t *testing.T) {
	t.Log()
}

func TestArray(t *testing.T) {
	t.Log(Array("aaa", []string{"aaa"}, "haha"))
}

func TestDuration(t *testing.T) {
	t.Log(Duration("aaa", time.Hour, "haha"))
}

func TestParse(t *testing.T) {
	Parse()
}
