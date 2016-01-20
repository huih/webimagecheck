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

type CheckState struct {
	CheckPage int
	CheckImage int
	CheckFail int
	ImageOk int
	ImageBad int
}

var checkState = &CheckState{0,0,0,0,0}

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
		
		//statistics
		checkState.CheckPage = checkState.CheckPage + 1
		
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
		
		checkState.CheckImage = checkState.CheckImage + 1
		
		canshow, err := ImageCanShow(imageUrl)
		if err != nil {
			//logs.Debug("pageUrl: %s, imageUrl: %s, err: %s", tipr.PageUrl, tipr.ImageUrl, err.Error())
			checkState.CheckFail = checkState.CheckFail + 1
			continue
		}
		
		if canshow == false {
			checkState.ImageBad = checkState.ImageBad + 1
			data := "page:" + tipr.PageUrl + " image:" + tipr.ImageUrl + "\n"
			WriteFile("bad.txt", []byte(data))
		} else {
			checkState.ImageOk = checkState.ImageOk + 1
		}
	}
	
	finished <- true
}

func PrintStatistics(interval time.Duration) {
	for {
		logs.Debug("checkPages: %d, checkImages: %d, checkFail: %d, imageOk: %d, imageBad: %d, pageListsize: %d, imageListSize: %d", 
			checkState.CheckPage, checkState.CheckImage, checkState.CheckFail, checkState.ImageOk, checkState.ImageBad, PageUrlList.Length(), ImageUrlList.Length())
		time.Sleep(interval * time.Second)
	}
}
func main(){
	
	//start parse pages
	go GetPages("http://www.zhiliaoyuan.com")
	
	//start get image url
	go GetPageImageUrl()
	
	//start check image
	go CheckImage()
	
	//start statistics
	go PrintStatistics(5)
	
	
	<- finished
}

