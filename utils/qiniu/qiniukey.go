package qiniu

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

const (
	ACCESS_KEY = "MQUnVFi5W6YPDHkasxm5btuo3hW3QHeZJIqUX7oL"
	SECRET_KEY = "62Z_JQ9PvuVUoK5wIApSOR4ls3CGOmpk1RSVaJH2"
	QINIU_BUCKET = "sourcandy"
	NETWORK_PREFIX = "https://cdn.sourcandy.cn/"
)

//七牛云上传文件-表单上传（localFile 文件地址，prefix 文件夹前缀，names 文件名称）
//返回值（address 文件地址，fileName 文件名称，certificate 证明，err 错误信息）
func QiNiu_UpLoadFile(localFile,prefix,names string) (address,fileName, certificate string,err error) {
	putPolicy := storage.PutPolicy{
		Scope: QINIU_BUCKET,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "sourcandy's bucket",
		},
	}
	if prefix != "" {
		prefix = prefix + "/"
	}
	suffixs := strings.Split(localFile,".")
	if names == ""{
		names,_= UniqueId()
	}
	key := prefix + names + "." + suffixs[len(suffixs)-1]
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return  "","","",err
	}
	return NETWORK_PREFIX + prefix + ret.Key,ret.Key,ret.Hash,nil
}

//七牛云上传文件-字节流上传（localFile 文件地址，prefix 文件夹前缀）
func QiNiu_ByteUploadFile(localFile,prefix,names string) (address,fileName, certificate string,err error)  {
	putPolicy := storage.PutPolicy{
		Scope: QINIU_BUCKET,
	}
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "sourcandy's bucket",
		},
	}
	res, err := http.Get(localFile)
	if err != nil {
		return "","","", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "","","", err
	}
	dataLen := int64(len(data))
	if prefix != "" {
		prefix = prefix + "/"
	}
	suffixs := strings.Split(localFile,".")
	suffix,_ := Prefix(suffixs[len(suffixs)-1])
	if suffix == "" {
		suffix = "jpg"
	}
	if names == ""{
		names,_= UniqueId()
	}
	key := prefix + names + "." + suffix
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "","","", err
	}
	return NETWORK_PREFIX + ret.Key,ret.Key,ret.Hash,nil
}

//文件流上传
func QiNiu_SourceUploadFile(source *multipart.FileHeader,prefix,names string) (address,fileName, certificate string,err error) {
	src, err := source.Open()
	if err != nil {
		return "","","",nil
	}
	defer src.Close()
	putPolicy := storage.PutPolicy{
		Scope: QINIU_BUCKET,
	}
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "sourcandy's bucket",
		},
	}

	if prefix != "" {
		prefix = prefix + "/"
	}
	suffixs := strings.Split(source.Filename,".")
	suffix,_ := Prefix(suffixs[len(suffixs)-1])
	if suffix == "" {
		suffix = "jpg"
	}
	if names == ""{
		names,_= UniqueId()
	}
	key := prefix + names + "." + suffix
	err = formUploader.Put(context.Background(), &ret, upToken, key,src , source.Size, &putExtra)
	if err != nil {
		return "","","", err
	}
	return NETWORK_PREFIX + ret.Key,ret.Key,ret.Hash,nil
}
//获取后缀
func Prefix(prefix string) (string,error) {
	prefixList := map[string]string{
		"jpg":"jpg",
		"png":"png",
		"jpeg":"jpeg",
		"mp4":"mp4",
		"xls":"xls",
	}
	return prefixList[prefix],nil
}

func QiNiu_DeleteFile(adder string) error {
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(QINIU_BUCKET, adder)
	if err != nil {
		return err
	}
	return nil
}


//生成uid字串
func UniqueId() (string,error) {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "",err
	}
	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))
	return hex.EncodeToString(h.Sum(nil)),nil
}
