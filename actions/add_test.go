package actions

import (
	"reflect"
	"testing"
)

func TestCheckArgsComptability(t *testing.T) {
	expected := "TESTDEV is not a valid argument to be written. Expected: [NAME VARIABLE]=[VALUE]"
	args := []string{"TEST=dev", "TESTDEV"}
	err := checkArgsComptability(args)
	if err.Error() != expected {
		t.Errorf("Error actual = %v, expected: %v", err, expected)
	}

	argsSuccess := []string{"test=dev", "prod=true", "env=prod", "token=1222aa"}
	errCheck := checkArgsComptability(argsSuccess)
	if errCheck != nil {
		t.Error(errCheck)
	}
}

func TestTransformArgs(t *testing.T) {
	argsChecked := []string{"test=dev", "prod=true", "env=prod", "token=abcde"}
	transformedArgs := transformArgs(argsChecked)
	eq1 := reflect.DeepEqual(transformedArgs["test"], "dev")
	eq2 := reflect.DeepEqual(transformedArgs["prod"], "true")
	eq3 := reflect.DeepEqual(transformedArgs["env"], "prod")
	eq4 := reflect.DeepEqual(transformedArgs["token"], "abcde")
	if !eq1 || !eq2 || !eq3 || !eq4 {
		t.Error("One of the value didnt match")
	}
	if transformedArgs == nil {
		t.Errorf("Expected non nil but got %v", transformedArgs)
	}
}
