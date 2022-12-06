package controller

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type ImgCommand struct {
	ImgName   string `form:"ImgName" json:"ImgName"`
	URL       string `form:"URL" json:"URL"`
	Condition string `form:"Condition" json:"Condition"`
}

type userbox struct {
	Items []string
}

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{"error": 1, "msg": "上傳失敗"})
		return
	}
	fileName := file.Filename
	myfolder := `./public/img/` + fileName
	if err := c.SaveUploadedFile(file, myfolder); err != nil {
		c.JSON(200, gin.H{"error": 2, "msg": "上傳失敗"})
		return
	}
	c.JSON(200, gin.H{"error": 0, "msg": "上傳成功"})
}

func AddImg(c *gin.Context) {
	var ImgCmd ImgCommand
	c.ShouldBind(&ImgCmd)
	if ImgCmd.URL == "" || ImgCmd.Condition == "" {
		c.JSON(200, gin.H{"error": 1, "msg": "參數遺失"})
		return
	}
	count, err := UrlGetImg(ImgCmd.URL, ImgCmd.Condition)
	if err != nil {
		c.JSON(200, gin.H{"error": 2, "msg": "操作失敗"})
	}

	c.JSON(200, gin.H{"error": 0, "msg": "成功", "data": count})
}

func GetImg(c *gin.Context) {
	myfolder := `./public/img/`
	var data = userbox{}
	files, _ := ioutil.ReadDir(myfolder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			data.Items = append(data.Items, file.Name())
		}
	}
	c.JSON(200, gin.H{"error": 0, "msg": "成功", "data": data.Items})
}

func DelImg(c *gin.Context) {
	var ImgCmd ImgCommand
	c.ShouldBind(&ImgCmd)
	if ImgCmd.ImgName == "" {
		c.JSON(200, gin.H{"error": 1, "msg": "參數遺失"})
		return
	}
	myfolder := `./public/img/` + ImgCmd.ImgName
	err := os.Remove(myfolder)
	if err != nil {
		c.JSON(200, gin.H{"error": 2, "msg": "刪除失敗"})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{"error": 0, "msg": "刪除成功"})
	return
}

func UrlGetImg(url, Condition string) (count int, err error) {
	c := colly.NewCollector()
	count = 0
	c.OnHTML(Condition, func(e *colly.HTMLElement) {
		getImg(e.Attr("src"))
		count++
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	//收到返回資訊
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("收到响应后调用:", response.Request.URL)
	})

	// 此為 爬蟲目標網站 網站
	err = c.Visit(url)
	if err != nil {
		fmt.Println("具体错误:", err)
	}
	fmt.Println("查詢到項目 :", count)
	return count, err
}

func getImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = "./public/img/" + path[len(path)-1]
	}
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	// defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}
