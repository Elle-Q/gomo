package main

import (
	"gomo/config"
	"gomo/qiniu"
)

func main() {
	config.Setup("config/config.yml")
	config.InitDB()

	//上传本地图片
	//UploadLocalDir("背景图片", "E:\\homolog资源文件\\默认背景")
	//img.UploadLocal("E:\\homolog资源文件\\默认背景\\bg6.jpg", "背景图片")

	//对qiniu上已经存在的视频进行分片操作
	//link := video.OpsVideoHLSForExistKey("1646038527976/C02L17_wrinkles.mp4", "C02L17_wrinkles.m3u8")
	//println("http://api.qiniu.com/status/get/prefop?id=",link)

	//fileBytes, _ :=os.ReadFile("Q:\\3d_parttime\\C02L17_wrinkles.mp4")
	//link := video.UploadVideoFileForHLS(fileBytes, "item01.mp4", "item01.m3u8")
	//println("---------------upload------------",link)


	qiniu.GetPrivateUrl("test.m3u8")
}
