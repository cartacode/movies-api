///////////////////////////////
//// Proprietary Code
//// Company: Vuli
//// Creator: Richard Hoop
//// Contact: richard@spry.is
///////////////////////////////

package main

import (
	"fmt"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/VuliTv/go-movie-api/router"
	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

var log = logging.GetProdLog()

func init() {
	// ENV Prefix
	defaultENV := "local"
	defaultPATH := "."
	envarCONFIGPATH := "CONFIG_PATH"
	envarENV := "ENV"
	viper.SetEnvPrefix("VULI_API")

	// Not sure what this error is
	if err := viper.BindEnv(envarENV); err != nil {
		panic(err)
	}

	// Set the env to local if not found
	if viper.Get(envarENV) == nil {
		viper.Set(envarENV, defaultENV)
		log.Infow(envarENV+" envar not found. defaulting", "env", defaultENV)
	}

	if viper.Get(envarCONFIGPATH) == nil {
		viper.Set(envarCONFIGPATH, defaultPATH)
		log.Infow(envarCONFIGPATH+" envar not found", "path", defaultPATH)
	}
	viper.AddConfigPath(viper.Get(envarCONFIGPATH).(string))

	// Set our config defaults
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings." + defaultENV)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	log.Info("config imported", viper.AllKeys())

}

func main() {
	logging.SetLevel(logging.DEBUG)
	log.Info("Server started")

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	viper.SetConfigName("api") // name of config file (without extension)
	viper.AddConfigPath(".")
	// viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv() // read in environment variables that match

	// Default config values
}
