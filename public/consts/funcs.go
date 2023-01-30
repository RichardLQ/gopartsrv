package consts

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

//生成唯一uid
func Uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

//获取当天最后时间戳
func GetbeforeDawn() int64 {
	nows := time.Now().Format(FORMATDATESHORT)
	nows = nows + " 23:59:59"
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(FORMATDATELONG, nows, loc)
	timestamp := tmp.Unix()
	return timestamp
}

//get请求
func HttpGet(url string)string{
	res, err :=http.Get(url)
	if err != nil {
		return ""
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return ""
	}
	return string(robots)
}

//post请求
func HttpPost(url string, data interface{}, contentType string) (content string) {
	jsonStr, _ := json.Marshal(data)
	if contentType == "" {
		contentType = "application/json"
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return
}

//读文件 0:success,1:fail
func ReadContent(filePath string) (content string,err error) {
	data, err := ioutil.ReadFile(filePath)   // 读取文件
	if err != nil {
		return "",err
	}
	return string(data),nil
}

//写文件 0:success,1:fail
func WriteContent(filePath,content string) (status int,err error) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return 1,err
	}
	_,err=f.Write([]byte(content))
	if err != nil {
		return 1,err
	}
	return 0,nil
}

//运行cmd命令(cmds 为是什么命令，params为参数，第一个为地址)
func CmdScript(cmds string,params []string) (string,error) {
	cmd0 := exec.Command(cmds,params...)
	stdout0 , err := cmd0.StdoutPipe() // 获取命令输出内容
	if err != nil {
		return "",err
	}
	if err = cmd0.Start(); err != nil {  //开始执行命令
		return "",err
	}
	useBufferIO := false
	if !useBufferIO {
		var outputBuf0 bytes.Buffer
		for {
			tempoutput := make([]byte, 256)
			n, err := stdout0.Read(tempoutput)
			if err != nil {
				if err == io.EOF {  //读取到内容的最后位置
					break
				} else {
					return "",err
				}
			}
			if n > 0 {
				outputBuf0.Write(tempoutput[:n])
			}
		}
		return outputBuf0.String(),nil
	} else {
		outputbuf0 := bufio.NewReader(stdout0)
		touput0 , _ , err := outputbuf0.ReadLine()
		if err != nil {
			return "",err
		}
		return string(touput0),nil
	}
}