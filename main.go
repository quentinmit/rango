package main

import (
	"log"
	"net/http/fcgi"
	"os"

	"github.com/spf13/viper"
	"github.com/quentinmit/rango/rangolib"
)

func main() {

	// setup config file
	viper.SetConfigName("config")
	viper.ReadInConfig()

	// set config defaults
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("AdminDir", "admin")
	viper.SetDefault("AssetsDir", "static/assets")

	// make sure content dir exists
	contentDir := viper.GetString("ContentDir")
	_, err := os.Stat(contentDir)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(contentDir, 0755)
	}

	// make sure assets dir exists
	assetsDir := viper.GetString("AssetsDir")
	_, err = os.Stat(assetsDir)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(assetsDir, 0755)
	}

	// create router
	router := NewRouter(&RouterConfig{
		Handlers: &Handlers{
			Config:     rangolib.NewConfig("config.toml"),
			Dir:        rangolib.NewDir(),
			Page:       rangolib.NewPage(),
			ContentDir: contentDir,
			AssetsDir:  assetsDir,
		},
		AdminDir: viper.GetString("AdminDir"),
	})

	if err := fcgi.Serve(nil, router); err != nil {
		log.Fatal(err)
	}
}
