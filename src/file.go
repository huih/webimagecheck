package main
import (
	"os"
	"github.com/gotools/logs"
)

func WriteFile(fileName string, data []byte) error {
	err := os.Remove(fileName)
	if err != nil {
		logs.Debug(err.Error())
	}
	
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR |os.O_TRUNC, os.ModePerm|os.ModeTemporary)
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