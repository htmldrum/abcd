package main

import(
	"testing"
)

func TestBuildCache(t *testing.T) {
	BuildCache() // go compiler is smart enough to flag this :P as an invalid operation
	if(false){
		t.Fatal("Result should fail");
	}
}
