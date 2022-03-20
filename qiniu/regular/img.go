package regular

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gomo/config"
	"mime/multipart"
	"os"
)

func UploadLocal(filePath string, fileName string, name string) (link string){
	//get upToken
	upToken := GetToken(config.QiniuConfig.PubBucket)

	//make key (timestamp)
	key := fmt.Sprintf("%s/%s/%s", "config",name, fileName)

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
	return key
}

func UploadFilePub(file multipart.File,len int64, key string) (string, error){
	//get upToken
	upToken := GetToken(config.QiniuConfig.PubBucket)
	return uploadFile(file,len,key, upToken)
}

func UploadFilePrivate(file multipart.File,len int64, key string) (string, error){
	//get upToken
	upToken := GetToken(config.QiniuConfig.VideoBucket)
	return uploadFile(file,len,key, upToken)
}

func uploadFile(file multipart.File,len int64, key string, token string) (string, error){

	//cfg
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneXinjiapo
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "avatar",  //啊啊啊啊啊啊~
		},
	}
	err := formUploader.Put(context.Background(), &ret, token, key, file, len, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("上传任务提交成功:", ret.PersistentID)

	return ret.PersistentID, err
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


func GetToken(bucket string) string {
	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}




