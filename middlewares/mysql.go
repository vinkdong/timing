package middlewares

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinkdong/gox/log"
	"github.com/vinkdong/timing/types"
	"sync"
	"time"
)

type MysqlMiddleware struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	db       *sql.DB
	Rule     *types.Rule
}

const mysqlConnectScheme  = "%s:%s@tcp(%s:%d)/%s"

func (m *MysqlMiddleware) Init(rule *types.Rule) {
	m.Rule = rule
}

func (m *MysqlMiddleware) Process() {
	if m.Rule.Count != 0 && m.Rule.Executed >= m.Rule.Count {
		return
	}
	m.init()
	defer m.db.Close()
	for _, sql := range m.Rule.Sql.Execute {
		if m.Rule.Thread > 0 {
			m.MultiThreadExecute(sql)
		}else {
			m.Execute(sql)
		}
	}
}

func (m *MysqlMiddleware) MultiThreadExecute(sql string) {
	var k int16
	var wg sync.WaitGroup
	wg.Add(int(m.Rule.Thread))
	for k = 0; k < m.Rule.Thread; k++ {
		go func() {
			defer wg.Done()
			m.Execute(sql)
		}()
	}
	wg.Wait()
}

func (m *MysqlMiddleware) checkProcessed() bool {
	fmt.Println(m.Rule.Executed)
	if m.Rule.Count != 0 && m.Rule.Executed > m.Rule.Count {
		spent := time.Now().UnixNano() - m.Rule.Started
		thtime := time.Duration(spent) * time.Nanosecond
		log.Info("request count is over limit stopped")
		log.Infof("spent %s", thtime.String())
		return true
	}
	return false
}

func (m *MysqlMiddleware) init() {
	if m.db != nil{
		return
	}
	dbConfig := m.Rule.Database
	db, err := sql.Open("mysql", fmt.Sprintf(mysqlConnectScheme,
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database))
	if err != nil {
		panic(err.Error())
	}
	m.db = db
}

func (m *MysqlMiddleware) Execute(sql string) sql.Result {
	if m.checkProcessed(){
		log.Info("request count is over limit stopped")
		m.Rule.Skip = true
		return nil
	}
	if m.db == nil{
		m.init()
	}
	result, err := m.db.Exec(sql)
	if err != nil {
		log.Error(err.Error())
	} else {
		m.Rule.Executed += 1
		//log.Info("executed")
	}
	return result
}
