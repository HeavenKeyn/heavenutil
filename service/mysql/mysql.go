package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DBI interface {
	Drop(table string) error
	Insert(table string, m map[string]interface{}) error
	BatchInsert(table string, data []map[string]interface{}) error
	FindBy(table string, dest interface{}, value map[string]interface{}) (map[string]interface{}, error)
	Find(table string, desc interface{}) error
}

type MDB struct {
	*sqlx.DB
}

func (mdb *MDB) SetDB(db *sqlx.DB) error {
	if mdb.DB != nil {
		err := mdb.DB.Close()
		if err != nil {
			return err
		}
	} else {
		mdb.DB = db
	}
	return nil
}

func (mdb *MDB) Connect(host, username, password, database string) error {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		username, password, "tcp", host, 3306, database)+"?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		return err
	}
	err = mdb.SetDB(db)
	return err
}

func (mdb *MDB) Drop(table string) error {
	return nil
}

func (mdb *MDB) Insert(table string, m map[string]interface{}) error {
	_, err := mdb.DB.NamedExec(insertSQL(table, m), m)
	return err
}

func (mdb *MDB) BatchInsert(table string, data []map[string]interface{}) error {
	_, err := mdb.DB.NamedExec(insertSQL(table, data[0]), data)
	return err
}

func (mdb *MDB) FindBy(table string, dest interface{}, value map[string]interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (mdb *MDB) Find(table string, dest interface{}) error {
	return nil
}

func insertSQL(table string, m map[string]interface{}) string {
	var build strings.Builder
	build.WriteString("insert into ")
	build.WriteString(table)
	build.WriteString(" (")
	var values strings.Builder
	i := 0
	for k := range m {
		i++
		build.WriteString(k)
		values.WriteString(":")
		values.WriteString(k)
		if i != len(m) {
			build.WriteString(",")
			values.WriteString(",")
		}
	}
	build.WriteString(") values (")
	build.WriteString(values.String())
	build.WriteString(")")
	return build.String()
}

func structToMap(dest interface{}) (map[string]interface{}, error) {
	return nil, nil
}
