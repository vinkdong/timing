package middlewares

import "testing"

func TestMysqlConnect(t *testing.T)  {
  m := MysqlMiddleware{Host:"192.168.99.100"}
  m.init()
}