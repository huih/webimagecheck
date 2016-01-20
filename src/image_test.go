package main
import (
	"github.com/gotools/logs"
	"testing"
)

func TestImageCanShow(t *testing.T) {
	//image ok
	//url := "http://zhiliaoyuan-zhiliao.stor.sinaapp.com/uploads/2016/01/20160119154259_54748.png"
	
	//image bad
	url := "http://zhiliaoyuan-zhiliao.stor.sinaapp.com/uploads/2016/01/20160117180444_86670.gif"
	
	ret, err := ImageCanShow(url)
	if err != nil {
		logs.Debug("err: %s", err.Error())
		return
	}
	if ret {
		logs.Debug("image is ok")
	} else {
		logs.Debug("image is not ok")
	}
}
