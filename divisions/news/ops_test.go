package news

import "testing"

func TestContentCategories(t *testing.T) {
	if len(ContentCategories) != 2 {
		t.Errorf("Expected 2 content categories")
	}
}

func TestListPodcasts(t *testing.T) {
	p, err := ListPodcasts()
	if err != nil {
		t.Errorf("Error when listing podcasts: %s", err)
	}
	if len(p) == 0 {
		t.Errorf("No Podcasts returned")
	}
}
