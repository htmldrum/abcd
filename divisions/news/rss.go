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

func ListReportsByState(s v0.State) ([]v0.Report, error) {
	r, err := v0.ListReportsByState(s)
	if err != nil {
		return r, err
	}
	return r, nil
}

func ListReportsByCategory(c v0.Category) ([]v0.Report, error) {
	r, err := v0.ListReportsByCategory(c)
	if err != nil {
		return r, err
	}
	return r, nil
}

func ListTopicsForLetter(l string) ([]V0.Topic, error) {
	r, err := v0.ListTopicsForLetter(l)
	if err != nil {
		return r, err
	}
	return r, nil
}

func ListReportsForTopic(t v0.Topic, page int) ([]V0.Report, error) {
	r, err := v0.ListReportsForTopic(t, page)
	if err != nil {
		return r, err
	}
	return r, nil
}
