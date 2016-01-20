package main
import (
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