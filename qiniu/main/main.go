package main

import (
	"leetroll/config"
)

func main() {
	config.Setup("config/config.yml")
	config.InitDB()

	//上传本地图片
	//UploadLocalDir("default_avatar", "F:\\dev\\homolog\\默认头像")
	//UploadLocalDir("default_bg", "F:\\dev\\homolog\\默认背景")
	//img.UploadLocal("F:\\dev\\homolog\\默认背景\\bg6.jpg", "背景图片")

	//对qiniu上已经存在的视频进行分片操作
	//link := video.OpsVideoHLSForExistKey("item/3/main/video2.mp4", "item/3/main/video2.m3u8")
	//println("http://api.qiniu.com/status/get/prefop?id=",link)

	//fileBytes, _ :=os.ReadFile("Q:\\3d_parttime\\C02L17_wrinkles.mp4")
	//link := video.UploadVideoFileForHLS(fileBytes, "item01.mp4", "item01.m3u8")
	//println("---------------upload------------",link)

	//qiniu.GetPrivateUrl("item/5/main/video3.m3u8")
	//qiniu.GetPubUrl("item01.m3u8")
}
