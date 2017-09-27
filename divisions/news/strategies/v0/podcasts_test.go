package v0

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestVCRListPodcasts(t *testing.T) {
	r, err := recorder.New("fixtures/TestVCRListPodcasts")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	client := http.Client{
		Transport: r,
	}

	resp, err := ListPodcasts(client)
	if err != nil {
		t.Fatalf("Failed to get url %s", err)
	}

	exp := 18

	if len(resp) != exp {
		t.Errorf("Expected %d podcasts!. Got: %d", exp, len(resp))
	}

	for _, p := range resp {
		s := reflect.ValueOf(&p).Elem()
		for i := 0; i < s.NumField(); i++ {
			typeOf := s.Type()
			f := s.Field(i)
			if f.Type().String() == "string" {
				if f.Interface() == "" {
					t.Errorf("Did not expect field %s to equal \"\"", typeOf.Field(i).Name)
				}
			}
		}
	}
}
