package main

import (
	log "github.com/Sirupsen/logrus"
)

func main(){
	log.WithFields(log.Fields{
		"abcd": "main",
	}).Info("Starting abcd")
	BuildCache()
	return
}

func BuildCache(){
	log.WithFields(log.Fields{
		"abcd": "#buildCache",
	}).Info("Executing abcd#buildCache.")
}
