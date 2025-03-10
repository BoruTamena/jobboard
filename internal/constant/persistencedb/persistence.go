package persistencedb

import (
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PersistenceDb struct {
	*db.Queries
}

func NewPersistenceDb(pool *pgxpool.Pool) PersistenceDb {

	return PersistenceDb{
		Queries: db.New(pool),
	}
}
