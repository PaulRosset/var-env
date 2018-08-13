package actions

import (
	"fmt"
	"testing"
)

// Create pre-function to create VARENV, the create after-function to clear the env
func TestList(t *testing.T) {
	//errOnAdd := Add([]string{"TEST_LIST=list", "TEST_LIST2=list2"})
	// if errOnAdd != nil {
	// 	t.Error("Error on pre-function", errOnAdd)
	// }
	varsEnvs, err := List()
	fmt.Println(varsEnvs)
	if err != nil {
		t.Error(err)
	}
}
