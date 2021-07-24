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
	FindByValue(table string, dest interface{}, value map[string]interface{}) error
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

func Connect(host, username, password, database string) (*MDB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		username, password, "tcp", host, 3306, database)+"?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		return nil, err
	}
	var mdb = &MDB{}
	err = mdb.SetDB(db)
	return mdb, err
}

func (mdb *MDB) Drop(table string) error {
	_, err := mdb.DB.Exec("drop table " + table)
	return err
}

func (mdb *MDB) Insert(table string, m map[string]interface{}) error {
	_, err := mdb.DB.NamedExec(insertSQL(table, m), m)
	return err
}

func (mdb *MDB) BatchInsert(table string, data []map[string]interface{}) error {
	_, err := mdb.DB.NamedExec(insertSQL(table, data[0]), data)
	return err
}

func (mdb *MDB) FindByValue(table string, dest interface{}, value map[string]interface{}) error {
	query, args := selectSQL(table, value)
	return mdb.DB.Get(dest, query, args...)
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

func selectSQL(table string, value map[string]interface{}) (string, []interface{}) {
	var builder strings.Builder
	args := make([]interface{}, 0)
	builder.WriteString("select * from ")
	builder.WriteString(table)
	builder.WriteString(" where ")
	c := 0
	for s, i := range value {
		c++
		builder.WriteString(s)
		builder.WriteString("=?")
		if c != len(value) {
			builder.WriteString(" and ")
		}
		args = append(args, i)
	}
	return builder.String(), args
}

func structToMap(dest interface{}) (map[string]interface{}, error) {
	return nil, nil
}
