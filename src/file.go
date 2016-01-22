package main
import (
	"strings"
	"io/ioutil"
	"os"
)

func WriteFile(fileName string, data []byte) error {	
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR |os.O_APPEND, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return err
	}
	defer f.Close()
	
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(fileName string) ([]string, error) {
	var dataList []string
	f, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm|os.ModeTemporary)
	if err != nil {
		return dataList, err
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return dataList, err
	}
	datastr := string(buf)
	dataList = strings.Split(datastr, "\n")
	return dataList, nil
}