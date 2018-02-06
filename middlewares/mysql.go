package middlewares

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinkdong/gox/log"
	"github.com/vinkdong/timing/types"
)

type MysqlMiddleware struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	db       *sql.DB
	Rule     types.Rule
}

const mysqlConnectScheme  = "%s:%s@tcp(%s:%d)/%s"

func (m *MysqlMiddleware) Init(rule types.Rule) {
	m.Rule = rule
}

func (m *MysqlMiddleware) Process() {
	m.init()
	defer m.db.Close()
	for _, sql := range m.Rule.Sql.Execute {
		m.Execute(sql)
	}
}

func (m *MysqlMiddleware) init() {
	dbConfig := m.Rule.Database
	db, err := sql.Open("mysql", fmt.Sprintf(mysqlConnectScheme,
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		panic(err.Error())
	}
	m.db = db
}

func (m *MysqlMiddleware) Execute(sql string) sql.Result {
	result, err := m.db.Exec(sql)
	if err != nil {
		log.Error(err.Error())
	}else {
		log.Info("executed")
	}
	return result
}
