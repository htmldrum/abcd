package v0

import "net/http"

type feed_template string
type feed_id string
type Report struct {
	Title       string
	Link        string
	Description string
	DcCreator   string
	PubDate     string
	Guid        string
	Category    string
}

type Feedable interface {
	getFeed(c http.Client) ([]Report, error)
}

func ListReportsByState(s State) ([]Report, error) {
	c := http.Client{}
	return s.getFeed(c)
}

func ListReportsByCategory(c Category) ([]Report, error) {
	client := http.Client{}
	return c.getFeed(client)
}

func ListReports() ([]Report, error) {
	var r []Report

	return r, nil
}

func ListTopicsByCharacter(character string) ([]Topic, error) {
	var t []Topic

	return t, nil
}

func ListReportsForTopic(t Topic, page int) ([]Report, error) {
	var r []Report

	return r, nil
}
