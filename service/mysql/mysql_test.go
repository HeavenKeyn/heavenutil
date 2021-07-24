package mysql

import (
	"testing"
	"time"
)

func TestMDB_Insert(t *testing.T) {
	mdb, err := Connect(host, username, password, database)
	if err != nil {
		t.Error(err)
	}
	err = mdb.Insert("user", map[string]interface{}{"username": "321", "password": "111"})
	if err != nil {
		t.Error(err)
	}
}

type User struct {
	Id         int64
	Username   string `db:"username"`
	Email      string
	Phone      string
	Password   string
	CreateTime time.Time `db:"create_time"`
}

func TestMDB_FindByValue(t *testing.T) {
	mdb, err := Connect(host, username, password, database)
	if err != nil {
		t.Error(err)
	}
	var result User
	err = mdb.FindByValue("user", &result, map[string]interface{}{"email": "111"})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
