package main

import(
	log "github.com/Sirupsen/logrus"
	"strings"
)

type Doc struct {
	url, body string
}

type DocNdx struct {
	name string
	docs []Doc
}

func main(){
	faves := []string{"prime-possum", "big-dog"}
	ndxs := []string{"index", "value"}
	FetchIndex(faves);
	log.WithFields(log.Fields{
		"ndxs": ndxs,
	}).Info("Ndx")
}

func FetchIndex(faves []string) {
	for _, name := range faves {
		url := strings.Join([]string{"http://", name, ".com"}, "")
		body := strings.Join([]string{"In news today, ", name, " has been killed. Music is dead."},"")
		log.WithFields(log.Fields{
			"url": url,
			"body": body,
		}).Info("Println sucks!")
	}
}
