package video

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"gomo/config"
	"io"
	"mime/multipart"
)

//文件分片上传(视频文件)
func OpsVideoHLSForExistKey(key string, m3u8Name string) (link string) {

	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	pipeline := "video-pipe" // 多媒体处理队列

	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		Zone:          &storage.ZoneXinjiapo, //对应机房
		UseHTTPS:      false,                 //是否使用https域名
		UseCdnDomains: false,                 //上传是否使用CDN加速
	}

	operationManager := storage.NewOperationManager(mac, &cfg)
	bucket := config.QiniuConfig.VideoBucket

	//处理指令
	mp4Fop := fmt.Sprintf("avthumb/m3u8/noDomain/1/segtime/15/vb/440k|saveas/%s", storage.EncodedEntry(bucket, m3u8Name))


	persistentId, err := operationManager.Pfop(bucket, key, mp4Fop, pipeline, "", true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("视频分片处理完毕>>> ", persistentId)
	return persistentId

	//deadline := time.Now().Add(time.Hour * 24).Unix() //24小时有效期
	//url := storage.MakePrivateURL(mac, config.QiniuConfig.PubDomain, key, deadline)
	//return fmt.Sprintf(url)

}

//文件分片上传(视频文件)
func UploadVideoForHLSFromFile(file multipart.File, len int64, key string, m3u8Name string) (string, error) {
	return UploadVideo(file,len,key,m3u8Name)
}

func UploadVideoForHLSFromBytes(file []byte, key string, m3u8Name string)(string, error) {
	length := len(file)
	return UploadVideo(bytes.NewReader(file), int64(length),key,m3u8Name)
}

//文件分片上传(视频文件)
func UploadVideo(reader io.Reader, len int64, key string, m3u8Name string) (string, error) {

	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	pipeline := "video-pipe" // 多媒体处理队列

	mac := qbox.NewMac(accessKey, secretKey)
	bucket := config.QiniuConfig.VideoBucket

	mp4Fop := fmt.Sprintf("avthumb/m3u8/noDomain/1/segtime/15/vb/440k|saveas/%s", storage.EncodedEntry(bucket, m3u8Name))

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		PersistentOps: mp4Fop,
		PersistentPipeline: pipeline,  // 多媒体处理队列
	}
	upToken := putPolicy.UploadToken(mac)


	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneXinjiapo //对应机房
	cfg.UseHTTPS = false             //是否使用https域名
	cfg.UseCdnDomains = false        //上传是否使用CDN加速

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, key, reader, len, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("视频分片任务提交完毕 >>> ", ret.PersistentID)


	//deadline := time.Now().Add(time.Hour * 24).Unix() //24小时有效期
	//url := storage.MakePrivateURL(mac,config.QiniuConfig.PubDomain, key, deadline)
	return ret.PersistentID, err

}