package db

import (
	"fmt"
	"sync"

	"github.com/coderi421/gframework/app/pkg/code"
	"github.com/coderi421/gframework/app/pkg/options"
	"github.com/coderi421/gframework/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbFactory *gorm.DB
	once      sync.Once
)

func GetDBFactoryOr(mysqlOpts *options.MySQLOptions) (*gorm.DB, error) {
	if mysqlOpts == nil && dbFactory == nil {
		return nil, errors.WithCode(code.ErrConnectDB, "failed to get mysql store factory")

	}
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlOpts.Username,
			mysqlOpts.Password,
			mysqlOpts.Host,
			mysqlOpts.Port,
			mysqlOpts.Database)
		dbFactory, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
		sqlDB, _ := dbFactory.DB()
		//允许连接多少个mysql
		sqlDB.SetMaxOpenConns(mysqlOpts.MaxOpenConnections)
		//允许最大的空闲的连接数
		sqlDB.SetMaxIdleConns(mysqlOpts.MaxIdleConnections)
		//重用连接的最大时长
		sqlDB.SetConnMaxLifetime(mysqlOpts.MaxConnectionLifetime)

	})
	if dbFactory == nil || err != nil {
		return nil, errors.WithCode(code.ErrConnectDB, "failed to get mysql store factory")
	}
	return dbFactory, nil
}
