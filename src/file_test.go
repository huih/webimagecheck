package main
import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	data := "hello\n"
	data1 := "world\n"
	
	WriteFile("test.txt", []byte(data))
	WriteFile("test.txt", []byte(data1))
	
}
