package main
import (
	"errors"
	"path"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
//	"github.com/gotools/logs"
)

//filter some image url
func FilterImageUrl(imageUrlArray []string, subUrl string) []string{
	var newImageUrlArray []string
	for _, imagePath := range imageUrlArray {
		if true == strings.Contains(imagePath, subUrl) {
			newImageUrlArray = append(newImageUrlArray, imagePath)	
		}
	}
	return newImageUrlArray
}

func FindImagePathUseGoQuery(url string) []string {
	query,_ := goquery.NewDocument(url)
	imgs := query.Find("img")
	
	var imageUrlArray []string
	for index:= 0; index < imgs.Length(); index++ {
		imagePath, exist := imgs.Eq(index).Attr("src")
		if exist {
			imageUrlArray = append(imageUrlArray, imagePath)
		}
	}
	return imageUrlArray
}

//http://zhiliaoyuan-zhiliao.stor.sinaapp.com/uploads/2016/01/20160117180444_86670.gif
func ImageCanShow(imageUrl string) (bool, error) {
	
	var UserAgent = "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36"
	req, _ := http.NewRequest("GET", imageUrl, nil)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "zhiliaoyuan-zhiliao.stor.sinaapp.com")
	req.Header.Set("Accept", "image/webp,*/*;q=0.8")
	req.Header.Set("User-Agent", UserAgent)
	resp, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	
	//get image error
	contentType := resp.Header.Get("Content-Type")
	if false == strings.Contains(contentType, "image") {
		return false, errors.New("get image error")
	}
	
	buf, _ := ioutil.ReadAll(resp.Body)
	if len(buf) < 4 {
		return false, nil
	}
	fileExt := path.Ext(imageUrl)
	
	var ret string
	if fileExt == ".jpg" {
		ret = fmt.Sprintf("%02x%02x%02x",(buf[0]&0xff), (buf[1]&0xff), (buf[2]&0xff))
	}  else {
		ret = fmt.Sprintf("%02x%02x%02x%02x",(buf[0]&0xff), (buf[1]&0xff), (buf[2]&0xff), (buf[3]&0xff))
	}

	switch fileExt {
		case ".gif":
			if ret == "47494638" {
				return true, nil
			} else {
				return false, nil
			}
		case ".tif":
			if ret == "49492a00" {
				return true, nil
			} else {
				return false, nil
			}
		case ".png":
			if ret == "89504e47" {
				return true, nil
			} else {
				return false, nil
			}
		case ".jpg":
			if ret == "ffd8ff" {
				return true, nil
			} else {
				return false, nil
			}
		default:
			errStr := "unkown image format: " + fmt.Sprintf("ext: %s, ret: %s", fileExt, ret)
			return false, errors.New(errStr)
	}
	
	return true, nil	
}
