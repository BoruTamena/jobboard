package initiator

import (
	"log"
	"net/http"

	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/gin-gonic/gin"
)

// Initiate
// @title JobBoard
// @version 1.0.0
// @description This is a  Swagger API documentation for JobBoard Open source Project.
// @contact.name Boru Tamene Yadeta
// @contact.url  https://github.com/BoruTamena

// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your_token}" to authenticate

// @securityDefinitions.basic BasicAuth
// @description Basic authentication using username and password

// @schemes http
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

	// initalizing platform

	log.Println("initalizing redis cache")
	platform := InitPlatform(*config)

	log.Println("initalizing module")
	// Initalizing Module
	md := InitModule(persistence, platform)
	// Initalizing Handler
	handler := InitHandler(md)

	InitRouting(rg, handler, md)

	server := http.Server{

		Addr:    "localhost:" + config.Server.Port,
		Handler: ServerEngine,
	}

	log.Fatal(server.ListenAndServe())

}
