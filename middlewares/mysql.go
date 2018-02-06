package middlewares

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinkdong/gox/log"
)

type MysqlMiddleware struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	db       *sql.DB
}

const mysqlConnectScheme  = "%s:%s@tcp(%s:%d)/%s"

func (m *MysqlMiddleware) init() {
	db, err := sql.Open("mysql", fmt.Sprintf(mysqlConnectScheme,
		m.Username, m.Password, m.Host, m.Port, m.Database))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	m.db = db
}

func (m *MysqlMiddleware) Execute(sql string) sql.Result {
	result, err := m.db.Exec(sql)
	if err != nil {
		log.Error(err.Error())
	}
	return result
}
