package persistencedb

import (
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersistenceDb struct {
	*gorm.DB
}

func NewPersistenceDb(config dto.Config) PersistenceDb {

	g_db, err := gorm.Open(postgres.Open(config.Db.PgUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return PersistenceDb{
		g_db,
	}
}
