package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Db 数据库查询采用sqlx实现，参考 http://jmoiron.github.io/sqlx/
type Db struct {
	fd *sqlx.DB
}

// NewDb 新建一个数据库链接
func (d *Db) NewDb(driverName, dataSourceName string) error {
	var err error
	if d.fd, err = sqlx.Open(driverName, dataSourceName); err != nil {
		return err
	}
	return nil
}

// Ping 简单ping测试
func (d *Db) Ping() error {
	return d.fd.Ping()
}

// Exec 直接执行原生语句
func (d *Db) Exec(sql string) (sql.Result, error) {
	return d.fd.Exec(sql)
}
