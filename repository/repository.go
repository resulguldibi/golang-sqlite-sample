package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	GetById(id int64) (interface{}, error)
	GetAll(instanceType interface{}) (interface{}, error)
	Update(query string, args ...interface{}) (result sql.Result, err error)
	Delete(query string, args ...interface{}) (result sql.Result, err error)
	Insert(query string, args ...interface{}) (result sql.Result, err error)
}

type BaseRepository struct {
	dbClient *DBClient
}

type UserRepository struct {
	BaseRepository
}

type UserBalanceRepository struct {
	BaseRepository
}

type DBClient struct{
	pool *sqlx.DB
}

func NewUserRepository(dbClient *DBClient) UserRepository{
	return UserRepository{BaseRepository : BaseRepository{dbClient : dbClient}}
}

func NewUserBalanceRepository(dbClient *DBClient) UserBalanceRepository{
	return UserBalanceRepository{BaseRepository : BaseRepository{dbClient : dbClient}}
}

type DBClientFactory struct{
    driverName string
 	dataSourceName string
}

func NewDbClientFactory(driverName string, dataSourceName string) DBClientFactory{
	return DBClientFactory{driverName : driverName, dataSourceName : dataSourceName}
}

func (dbCLientFactory DBClientFactory) NewDBClient() *DBClient{
	client := &DBClient{}

	pool, err := Connect(dbCLientFactory.driverName, dbCLientFactory.dataSourceName)

	if err !=nil{
		panic(err)
	}

	client.pool = pool

	return client
}







