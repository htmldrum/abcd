package news

import (
	"net/http"

	"github.com/htmldrum/abcd/divisions/news/strategies/v0"
)

type ContentCategory struct {
	Name string
}

var (
	Podcasts          = ContentCategory{"Podcasts"}
	Reports           = ContentCategory{"Reports"}
	ContentCategories = []ContentCategory{Podcasts, Reports}
)

func ListPodcasts() ([]v0.Podcast, error) {
	c := http.Client{}
	p, err := v0.ListPodcasts(c)
	if err != nil {
		return p, err
	}
	return p, nil
}

func ListReports() ([]v0.Report, error) {
	r, err := v0.ListReports()
	if err != nil {
		return r, err
	}
	return r, nil
}
