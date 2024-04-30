package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	config1 "novachat_engine/pkg/config"
	"sync"
)

var db *gorm.DB
var once sync.Once

func NewORM(dbConfig *config1.MySQLConfig) (*gorm.DB, error) {
	once.Do(func() {
		gormConfig := &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
		if dbConfig.Debug {
			gormConfig.Logger = logger.Default.LogMode(logger.Info)
		}

		_db, err := gorm.Open(mysql.Open(dbConfig.DSN), gormConfig)
		if err != nil {
			panic(err)
		}
		sqlDB, err := _db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(int(dbConfig.Idle))
		sqlDB.SetMaxOpenConns(int(dbConfig.Active))
		db = _db
	})
	return db, nil
}
