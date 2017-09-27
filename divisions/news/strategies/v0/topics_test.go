package v0

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestVCRListTopicsForLetterByLocation(t *testing.T) {
	r, err := recorder.New("fixtures/TestVCRListTopicsForLetterByLocation")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := http.Client{
		Transport: r,
	}

	resp, err := ListTopicsForLetterByLocation(client, "J")
	if err != nil {
		t.Fatalf("Failed to get url %s", err)
	}

	exp := 91

	if len(resp) != exp {
		t.Errorf("Expected %d podcasts!. Got: %d", exp, len(resp))
	}

}

func TestVCRListTopicsForLetterBySubject(t *testing.T) {
	r, err := recorder.New("fixtures/TestVCRListTopicsForLetterBySubject")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	client := http.Client{
		Transport: r,
	}

	resp, err := ListTopicsForLetterBySubject(client, "J")
	if err != nil {
		t.Fatalf("Failed to get url %s", err)
	}

	exp := 8

	if len(resp) != exp {
		t.Errorf("Expected %d podcasts!. Got: %d", exp, len(resp))
	}

}
