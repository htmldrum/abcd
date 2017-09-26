package news

import "github.com/htmldrum/abcd/divisions/news/strategies/v0"

type ContentCategory struct {
	Name string
}

type Podcast struct {
	Name string
}

var (
	Podcasts          = ContentCategory{"Podcasts"}
	Reports           = ContentCategory{"Reports"}
	ContentCategories = []ContentCategory{Podcasts, Reports}
)

// 2 kinds of content:
// podcasts
// reports
// Browsable by topic

func ListPodcasts() ([]Podcast, error) {
	var p []Podcast

	resp, err := v0.ListPodcasts
	if err != nil {
		return p, err
	}
	for n, pod_json := range resp {
		p = append(p, Podcast{pod_json.Name})
	}

	return p, nil
}
