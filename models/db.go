package models

import pg "gopkg.in/pg.v5"

/**
Constanta untuk koneksi database
*/
const (
	Host     = "localhost:5432"
	User     = "prototype"
	Password = "your_password"
	DBName   = "skripsi_call"
)

type DB struct {
	*pg.DB
}

func SetupDB() *DB {
	//inisialisai database ke postgreSQL
	db := pg.Connect(&pg.Options{
		User:     User,
		Password: Password,
		Addr:     Host,
		Database: DBName,
	})
	//return instansi struct
	return &DB{db}
}
