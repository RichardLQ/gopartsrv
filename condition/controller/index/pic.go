package index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

//GetRandomPic 获取图片
func GetRandomPic(c *gin.Context)  {
	urlList := getData()
	url := urlList[getRandm(len(urlList))]
	c.JSON(http.StatusOK, gin.H{"errs": 1, "msg": "请求成功", "data": url})
	return
}

//获取开始地址
func getdounimeiUrl() string {
	url := "https://www.dounimei.co/page/%d?orderby=hot"
	return fmt.Sprintf(url,getRandm(70))
}

func getRandm(ints int) int {
	rand.Seed(time.Now().UnixNano())
	// 生成1-100之间的随机数
	number:=rand.Intn(ints)+1
	return number
}
//抓取数据
func getData() []string {
	c := colly.NewCollector()
	urlList := []string{}
	c.OnHTML(".n-s", func(e *colly.HTMLElement) {
		urls := strings.Split(e.ChildAttr("img","data-src"),".php?src=")
		url := strings.Split(urls[1],"&h=260&w=260&zc=1")
		urlList = append(urlList,url[0])
	})
	c.Visit(getdounimeiUrl())
	return urlList
}