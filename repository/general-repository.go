package repository

import (
	"api_log/entity"
	"gorm.io/gorm"
)

type GeneralRepository interface {
	Get() []entity.Log
	Post(g entity.Log) entity.Log
	Put(g entity.Log) entity.Log
	Delete(g entity.Log)
}

type generalConnection struct {
	connection *gorm.DB
}

func NewGeneralRepository(dbConn *gorm.DB) GeneralRepository  {
	return &generalConnection{
		connection: dbConn,
	}
}

func (db *generalConnection) Get() []entity.Log {
	var logs []entity.Log
	db.connection.Find(&logs)
	return logs
}

func (db *generalConnection) Post(g entity.Log) entity.Log {
	db.connection.Save(&g)
	return g
}

func (db *generalConnection) Put(g entity.Log) entity.Log {
	db.connection.Save(&g)
	return g
}

func (db *generalConnection) Delete(g entity.Log)  {
	db.connection.Delete(&g)
}