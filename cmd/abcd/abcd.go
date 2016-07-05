package main

import (
	log "github.com/Sirupsen/logrus"
)

func main(){
	log.Info("Starting abcd")
	ReadConfig()
	// config := ReadConfig)(
	RefreshFeeds()
	StartServer()
}

func BuildCache(){
	log.WithFields(log.Fields{
		"abcd": "#buildCache",
	}).Info("Executing abcd#buildCache.")
}

func RefreshFeeds(){
	log.Info("Refresing feeds")
}

func StartServer(){
	log.Info("Starting server that serves feeds")
}
