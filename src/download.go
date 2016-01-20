package main

import (
	"errors"
	"net/http"
	"io/ioutil"
	"github.com/gotools/logs"
	
)

func DownLoad(url string) (pageContent string, err error){
	resp, err := http.Get(url)
	if err != nil {
		logs.Error("http get error")
		return "", errors.New("http get error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("read response body error")
		return "", errors.New("read response body error")
	}	
	return string(body), nil
}