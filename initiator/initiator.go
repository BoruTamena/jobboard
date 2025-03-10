package initiator

import (
	"log"
	"net/http"

	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/gin-gonic/gin"
)

// @title JobBoard
// @version 1.0.0
// @description This is a  Swagger API documentation for JobBoard Open source Project.
// @contact.name Boru Tamene Yadeta
// @contact.url  https://github.com/BoruTamena
func Init() {

	log.Println("Initalizing config file")

	err, config := InitViper("./")

	if err != nil {

		log.Println("file configuration initalizing  fail")
	}
	// initalizing server
	ServerEngine := gin.Default()
	rg := ServerEngine.Group("v1")

	log.Println("making migiration")
	// initalizing migiration
	mg := InitMigiration(config.Migration.Path, config.Db.PgUrl)

	UpMigiration(mg)

	// Init Persistence layer
	con_pol := InitPgDb(*config)

	pdb := persistencedb.NewPersistenceDb(con_pol)

	persistence := InitPersistence(pdb)

	log.Println("initalizing module")
	// Initalizing Module
	md := InitModule(persistence)
	// Initalizing Handler
	handler := InitHandler(md)

	InitRouting(rg, handler)

	server := http.Server{

		Addr:    "localhost:" + config.Server.Port,
		Handler: ServerEngine,
	}

	log.Fatal(server.ListenAndServe())

}
