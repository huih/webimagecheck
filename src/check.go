package main
import (
	"time"
	"github.com/gotools/logs"
	"github.com/gotools/lists"
)

var ImageUrlList = fifolist.New()

var PageImageStop = false

var finished chan bool

type ImagePageRel struct {
	PageUrl string
	ImageUrl string
}
func GetPageImageUrl() {
	for {
		if PageStop && PageUrlList.Length() <= 0 {
			break
		}
		
		pageUrl := PageUrlList.GetAndRemove()
		if pageUrl == nil {
			time.Sleep(1 * time.Second)
			continue
		}
		
		imagePathArray := FindImagePathUseGoQuery(pageUrl.(string))
		imagePathArray = FilterImageUrl(imagePathArray, "zhiliaoyuan-zhiliao.stor.sinaapp.com")
		
		for _, imagePath := range imagePathArray {
			var ipr ImagePageRel
			ipr.PageUrl = pageUrl.(string)
			ipr.ImageUrl = imagePath
			
			ImageUrlList.Add(ipr)
		}
	}
	PageImageStop = true
}

func CheckImage (){
	for {
		if PageImageStop && ImageUrlList.Length() <= 0 {
			break
		}
		
		
		ipr := ImageUrlList.GetAndRemove()
		if ipr == nil {
			time.Sleep(time.Second)
			continue
		}
		
		tipr := ipr.(ImagePageRel)
		imageUrl := tipr.ImageUrl
		canshow, err := ImageCanShow(imageUrl)
		if err != nil {
			//logs.Debug("pageUrl: %s, imageUrl: %s, err: %s", tipr.PageUrl, tipr.ImageUrl, err.Error())
			continue
		}
		
		if canshow == false {
			logs.Debug("pageUrl: %s, imageUrl: %s is bad image", tipr.PageUrl, tipr.ImageUrl)
		}
	}
	
	finished <- true
}


func main(){
	
	//start parse pages
	go GetPages("http://www.zhiliaoyuan.com")
	
	//start get image url
	go GetPageImageUrl()
	
	//start check image
	go CheckImage()
	
	<- finished
}

