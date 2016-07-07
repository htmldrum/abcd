package main

var c Config

func main(){
	c = ReadConfig()
	feeds := RefreshFeeds()
	SaveFeeds(*feeds)
	StartServer()
}
