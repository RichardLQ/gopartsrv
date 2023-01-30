package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopartsrv/public/consts"
	"io"
	"os"
	"time"
)

var DirAdress = "./logs/" + time.Now().Format(consts.FORMATDATESHORT) + "/" + time.Now().Format(consts.FORMATDATESHORT) + ".log"

func init() {
	judgeDate()
	f, _ := os.Create(DirAdress)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

//判断日期文件夹
func judgeDate() {
	dirs, _ := os.Getwd()
	_, err := os.Stat(dirs + "/logs")
	if err != nil {
		err = os.Mkdir(dirs+"/logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	_dir := dirs + "/logs/" + time.Now().Format(consts.FORMATDATESHORT)
	_, err = os.Stat(_dir)
	if err != nil {
		err = os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(DirAdress)
	if err != nil {
		_, err = os.Create(DirAdress)
		if err != nil {
			panic(err)
		}
	}
}

func Info(content string)  {
	file, err := os.OpenFile(DirAdress, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	cont,err:=file.WriteString(content)
	fmt.Println(cont)
	fmt.Println(err)
}
