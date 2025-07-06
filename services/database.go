package services

import (
	"api-core/libs"
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	DB *gorm.DB
}

func (s *DatabaseService) Open() {
	sqlDB, err := sql.Open("pgx", libs.DB_STRING)
	if err != nil {
		log.Fatalln("Can't connect to database")
	}

	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{},
	)
	if err != nil {
		log.Fatalln("Can't connect to database")
	}

	s.DB = gormDB
}
