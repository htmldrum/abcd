package v0

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestVCRStateGetFeed(t *testing.T) {
	r, err := recorder.New("fixtures/TestVCRStateGetFeed")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := http.Client{
		Transport: r,
	}
	state := States[0]

	resp, err := state.getFeed(client)
	if err != nil {
		t.Fatalf("Failed to get url %s", err)
	}

	exp := 250
	if len(resp) != exp {
		t.Errorf("Expected %d reports!. Got: %d", exp, len(resp))
	}
}
