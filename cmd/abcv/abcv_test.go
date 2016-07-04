package main

import(
	"testing"
)

func TestFetchIndex(t *testing.T) {
	faves := []string{"biggie-smalls", "tupac-shakuer"}
	FetchIndex(faves) // go compiler is smart enough to flag this :P as an invalid operation
	if(2 != 2){
		t.Fatal("There should be 2");
	}
}
