package img

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gomo/config"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UploadLocal(filePath string, fileName string) (link string){
	//get upToken
	upToken := GetPubImgToken()

	//make key (timestamp)
	key := strconv.FormatInt(time.Now().UnixMilli(), 10) + "/" +fileName

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneXinjiapo
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "avatar", //这是测试的好玩的
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取公开访问外链
	return fmt.Sprintf(storage.MakePublicURL(config.QiniuConfig.PubDomain, key))
}


func UploadFile(file multipart.File, fileName string) (link string){
	//get file size
	fileSize := getFileSize(file)

	//get upToken
	upToken := GetPubImgToken()

	//make key (timestamp)
	key := strconv.FormatInt(time.Now().UnixMilli(), 10) + "/" + fileName

	//cfg
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneXinjiapo
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "avatar",
		},
	}
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("上传成功:", ret.Key, ret.Hash)

	return config.QiniuConfig.PubDomain + "/" + key
}

func DeleteFile(key string)  {

}

func getFileSize(file multipart.File) int64 {
	var fileSize int64
	switch t := file.(type) {
	case *os.File:
		fi, _ := t.Stat()
		fileSize = fi.Size()
	default:
		fileSize, _ = file.Seek(0, 0)
	}
	return fileSize
}


func GetPubImgToken() string {
	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	bucket := config.QiniuConfig.PubBucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}



