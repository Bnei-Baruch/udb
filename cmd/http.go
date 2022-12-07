package cmd

import (
	"context"
	"github.com/Bnei-Baruch/udb/api"
	"github.com/Bnei-Baruch/udb/models"
	"github.com/Bnei-Baruch/udb/utils"
	"github.com/Bnei-Baruch/udb/version"
	"github.com/spf13/viper"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitHTTP() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.Infof("Starting UDB API server version %s", version.Version)

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
		utils.EnvMiddleware(models.DB, oidcIDTokenVerifier),
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
