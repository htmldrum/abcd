package ex_lib  // import "camlistore.org/pkg/ex_lib"

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T){
	fmt.Sprintf("Beginning TestHello...");
	result := Hello(4)
	if(result == 3){
		fmt.Sprintf("Passed! %d", result)
	} else {
		t.Errorf("Failed! %d", result)
	}
}
