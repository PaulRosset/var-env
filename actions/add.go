package actions

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Add arguments as CLI commands
func Add(args []string) error {
	errOnArgs := checkArgsComptability(args)
	if errOnArgs != nil {
		return errOnArgs
	}
	varsEnvs := transformArgs(args)
	errOnWriteEnv := writeEnv(varsEnvs, os.Getenv("HOME")+"/.bashrc")
	if errOnWriteEnv != nil {
		return errOnWriteEnv
	}
	fmt.Println("\tEnv variables Added.")
	return nil
}

func transformArgs(args []string) map[interface{}]interface{} {
	varsEnvs := make(map[interface{}]interface{})
	for _, vars := range args {
		combs := strings.Split(vars, "=")
		varsEnvs[combs[0]] = combs[1]
	}
	return varsEnvs
}

func checkArgsComptability(args []string) error {
	for _, arg := range args {
		isValidForProcess, err := regexp.Match("(.+={1}.+)", []byte(arg))
		if err != nil {
			return err
		}
		if !isValidForProcess {
			return errors.New(arg + " is not a valid argument to be written. Expected: [NAME VARIABLE]=[VALUE]")
		}
	}
	return nil
}
