package gorm

import (
	todogorm "github.com/jackyuan2010/todoapp/server/src/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDbContext struct {
}

func (pgdc PostgresDbContext) GetDb(dsn *gpaasgorm.DbDsn) *gorm.DB {
	pgsqlconfig := postgres.Config{
		DSN:                  (*dsn).Dsn(),
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlconfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		pgsqlconfig := (*dsn).(PostgresDbConfig)
		sqlDB.SetMaxIdleConns(pgsqlconfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(pgsqlconfig.MaxOpenConns)
		return db
	}
}