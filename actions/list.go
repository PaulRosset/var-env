package actions

import (
	"fmt"
	"os"
)

// List All the env variables created by VARENV software...
func List() error {
	for _, varEnv := range os.Environ() {
		if varEnv[:7] == "VARENV_" {
			fmt.Printf("\t%s\n", varEnv)
		}
	}
	fmt.Printf("\n")
	return nil
}
