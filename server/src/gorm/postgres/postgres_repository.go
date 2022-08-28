package gorm

import (
	gpaasgorm "github.com/jackyuan2010/todoapp/server/src/gorm"
	model "github.com/jackyuan2010/todoapp/server/src/model"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	dsn       *gpaasgorm.DbDsn
	dbcontext *gpaasgorm.DbContext
	DB        *gorm.DB
}

func NewPostgresRepository(dsn *gpaasgorm.DbDsn) *PostgresRepository {
	pgrepository := PostgresRepository{}
	pgrepository.dsn = dsn
	var dbcontext gpaasgorm.DbContext = PostgresDbContext{}
	pgrepository.dbcontext = &dbcontext
	pgrepository.DB = (*pgrepository.dbcontext).GetDb(pgrepository.dsn)
	return &pgrepository
}

func (pgrepository PostgresRepository) QueryById(id string) *model.ToDoTask {
	todTask := model.ToDoTask{}

	pgrepository.DB.First(&ToDoTask, "id=?", id)

	return &todTask
}