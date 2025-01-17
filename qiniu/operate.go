package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"leetroll/config"
	"leetroll/qiniu/regular"
	"leetroll/qiniu/video"
	"leetroll/tool"
	"mime/multipart"
	"strings"
)

// 删除文件
func DeleteFile(bucket string, key string) error {
	bucketManager := getBucketManager()
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 上传item资源文件
func UploadItemResc(fileHeader *multipart.FileHeader, rescType string, itemId int) (string, string, error) {

	//文件格式
	format := fileHeader.Header.Get("Content-Type")
	//文件名称
	fileName := fileHeader.Filename
	prefix := fmt.Sprintf("item/%d/%s", itemId, rescType)
	//七牛云存储的key('/'表示分割文件夹, eg:/item/99/preview)
	key := fmt.Sprintf("%s/%s", prefix, fileName)

	var m3u8Key string
	var persistentID string
	var err error
	//判断是否为视频(视频需要分片处理.其他文件流程一样)
	file, _ := fileHeader.Open()
	if tool.IsVideo(format) {
		//分片文件的m3u8文件名称
		palinName := strings.Split(fileName, ".")[0]
		m3u8Key = fmt.Sprintf("%s/%s.m3u8", prefix, palinName)
		persistentID, err = video.UploadVideoForHLSFromFile(file, fileHeader.Size, key, m3u8Key) //上传视频
	} else {
		persistentID, err = regular.UploadFilePrivate(file, fileHeader.Size, key) //上传普通文件
	}
	if err != nil {
		fmt.Println("上传出错: ", err)
		return "", "", err
	}
	//保存文件上传处理单号(PersistID) => db
	fmt.Println("提交PersistentID到数据库 >>> ", persistentID)

	//返回七牛云链接 (都是私有访问链接)
	return key, m3u8Key, err

}

func getBucketManager() *storage.BucketManager {
	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager
}
