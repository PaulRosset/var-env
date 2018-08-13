package actions

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Loader Load all the environment variables present in the specified yaml file, then write it on the the .bashrc to make it permanent.
func Loader(filename string) error {
	log.Println("Load " + filename)
	data, errOnYaml := openYml(filename)
	if errOnYaml != nil {
		return errOnYaml
	}
	varsEnv := make(map[interface{}]interface{})
	err := yaml.Unmarshal(data, &varsEnv)
	if err != nil {
		return err
	}
	errOnWriteEnv := writeEnv(varsEnv, os.Getenv("HOME")+"/.bashrc")
	if errOnWriteEnv != nil {
		return errOnWriteEnv
	}
	return nil
}

func openYml(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fileInfo, errStat := file.Stat()
	if errStat != nil {
		return nil, errStat
	}
	data := make([]byte, fileInfo.Size())
	_, errRead := file.Read(data)
	if errRead != nil {
		return nil, errRead
	}
	return data, nil
}

func writeEnv(varsEnv map[interface{}]interface{}, filename string) error {
	file, errOnOpenDestFile := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if errOnOpenDestFile != nil {
		return errOnOpenDestFile
	}
	defer file.Close()
	for name, val := range varsEnv {
		_, errOnWritting := file.WriteString("export VARENV_" + name.(string) + "=" + val.(string) + "\n")
		if errOnWritting != nil {
			log.Println(errOnWritting)
			return errOnWritting
		}
		log.Printf("%v added", name)
	}
	log.Println("Variable(s) set, you may need to refresh your bashrc to see the changes")
	return nil
}
