package actions

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Remove one or multiple env variables written as permanent.
func Remove(args []string) error {
	fmt.Println("RM")
	errOnArgs := checkValidityArgs(args)
	if errOnArgs != nil {
		return errOnArgs
	}
	errOnPresentVars := checkPresent(args)
	if errOnPresentVars != nil {
		return errOnPresentVars
	}
	for _, arg := range args {
		cmd := exec.Command("sed", "-i", "'/^export"+arg+"/ d'", os.Getenv("HOME")+"/.bashrc")
		errorOnSed := cmd.Run()
		if errorOnSed != nil {
			return errorOnSed
		}
	}
	return nil
}

func checkValidityArgs(args []string) error {
	for _, variable := range args {
		isMatch, err := regexp.Match("(^VARENV_.+)", []byte(variable))
		if err != nil {
			return err
		}
		if !isMatch {
			return errors.New(variable + " is not at the correct format. Expected format: VARENV_[THE NAME OF YOUR VARIABLE]")
		}
	}
	return nil
}

func checkPresent(args []string) error {
	nbArgs := len(args)
	count := 0
	for _, varEnv := range os.Environ() {
		comb := strings.Split(varEnv, "=")
		for _, arg := range args {
			if comb[0] == arg {
				count++
			}
		}
	}
	if nbArgs != count {
		return errors.New("One of the variable you provided is not set, so we can't remove it")
	}
	return nil
}
