package cmd

import (
	"context"
	"fmt"
	"github.com/Bnei-Baruch/udb/api"
	"github.com/Bnei-Baruch/udb/utils"
	"github.com/Bnei-Baruch/udb/version"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitHTTP() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.Infof("Starting UDB API server version %s", version.Version)

	log.Info("Setting up connection to UDB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jerusalem",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
		viper.GetString("db.port"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gorm_logrus.New().LogMode(logger.Info),
	})
	if err != nil {
		log.Infof("UDB connection error: %s", err)
		return
	}
	db.AutoMigrate()

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = append(corsConfig.AllowMethods, http.MethodDelete)
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true

	// Authentication
	var oidcIDTokenVerifier *oidc.IDTokenVerifier
	if viper.GetBool("authentication.enable") {
		log.Info("Initializing Auth System")
		issuer := viper.GetString("authentication.issuer")
		oidcProvider, err := oidc.NewProvider(context.TODO(), issuer)
		if err != nil {
			log.Infof("KC init error: %s", err)
			return
		}
		oidcIDTokenVerifier = oidcProvider.Verifier(&oidc.Config{
			SkipClientIDCheck: true,
		})
	}

	// Setup gin
	gin.SetMode(viper.GetString("server.mode"))
	router := gin.New()
	router.Use(
		utils.MdbLoggerMiddleware(),
		utils.EnvMiddleware(db, oidcIDTokenVerifier),
		utils.ErrorHandlingMiddleware(),
		utils.AuthenticationMiddleware(),
		cors.New(corsConfig),
		utils.RecoveryMiddleware())

	api.SetupRoutes(router)

	srv := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: router,
	}

	// service connections
	log.Infoln("Running application")
	if err := srv.ListenAndServe(); err != nil {
		log.Infof("Server listen: %s", err)
	}
}
