package repository

import (
	"programming/golang/rest-api/entity"

	"gorm.io/gorm"
)

//BookRepository is a ...
type PirateRepository interface {
	InsertPirate(p entity.Pirate) entity.Pirate
	UpdatePirate(p entity.Pirate) entity.Pirate
	DeletePirate(p entity.Pirate)
	AllPirate() []entity.Pirate
	FindPirateByID(pirateID uint64) entity.Pirate
}

type pirateConnection struct {
	connection *gorm.DB
}

//NewPirateRepository creates an instance PirateRepository
func NewPirateRepository(dbConn *gorm.DB) PirateRepository {
	return &pirateConnection{
		connection: dbConn,
	}
}

func (db *pirateConnection) InsertPirate(p entity.Pirate) entity.Pirate {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *pirateConnection) UpdatePirate(p entity.Pirate) entity.Pirate {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *pirateConnection) DeletePirate(p entity.Pirate) {
	db.connection.Delete(&p)
}

func (db *pirateConnection) FindPirateByID(pirateID uint64) entity.Pirate {
	var pirate entity.Pirate
	db.connection.Preload("User").Find(&pirate, pirateID)
	return pirate
}

func (db *pirateConnection) AllPirate() []entity.Pirate {
	var pirates []entity.Pirate
	db.connection.Preload("User").Find(&pirates)
	return pirates
}
