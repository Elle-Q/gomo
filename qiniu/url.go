package qiniu

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"gomo/config"
	"time"
)

// 公开空间访问
func GetPubUrl(key string) string {
	domain := config.QiniuConfig.PubDomain
	publicAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println("公开空间访问链接为: ", publicAccessURL)
	return publicAccessURL
}

// 私有空间访问
func GetPrivateUrl(key string) string{
	accessKey := config.QiniuConfig.AK
	secretKey := config.QiniuConfig.SK
	mac := auth.New(accessKey, secretKey)

	domain := config.QiniuConfig.VideoDomain
	deadline := time.Now().Add(time.Hour * 360).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURLv2WithQueryString(mac, domain, key, "pm3u8/0", deadline)

	//fmt.Println("私有空间访问链接为: ", privateAccessURL)
	return privateAccessURL
}
