package actions

import (
	"fmt"
	"os"
)

// List All the env variables created by VARENV software...
func List() ([]string, error) {
	getVarsEnv := []string{}
	for _, varEnv := range os.Environ() {
		if varEnv[:7] == "VARENV_" {
			getVarsEnv = append(getVarsEnv, varEnv)
		}
	}
	fmt.Printf("\n")
	return getVarsEnv, nil
}
