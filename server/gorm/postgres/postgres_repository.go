package gorm

import (
	todogorm "github.com/jackyuan2010/todoapp/server/gorm"
	model "github.com/jackyuan2010/todoapp/server/model"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	dsn       *todogorm.DbDsn
	dbcontext *todogorm.DbContext
	DB        *gorm.DB
}

func NewPostgresRepository(dsn *todogorm.DbDsn) *PostgresRepository {
	pgrepository := PostgresRepository{}
	pgrepository.dsn = dsn
	var dbcontext todogorm.DbContext = PostgresDbContext{}
	pgrepository.dbcontext = &dbcontext
	pgrepository.DB = (*pgrepository.dbcontext).GetDb(pgrepository.dsn)
	return &pgrepository
}

func (pgrepository PostgresRepository) QueryById(id string) *model.ToDoTask {
	todoTask := model.ToDoTask{}

	pgrepository.DB.First(&todoTask, "id=?", id)

	return &todoTask
}