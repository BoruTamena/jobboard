package initiator

import (
	"context"
	"log"
	"time"

	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitPgDb(cfg dto.Config) *pgxpool.Pool {

	url := cfg.Db.PgUrl

	con_config, err := pgxpool.ParseConfig(url)

	if err != nil {
		log.Fatalf("pg connection creation error:%v", err.Error())
	}

	con_config.ConnConfig.ConnectTimeout = time.Second

	// con_config.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {

	// 	return true
	// }
	// create connection pool

	con_pol, err := pgxpool.ConnectConfig(context.Background(), con_config)

	if err != nil {
		log.Fatalf("connection error::%v", err)
	}

	return con_pol

}
