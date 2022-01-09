package main

import (
	"gomo/config"
)

func main() {
	config.Setup("config/config.yml")
	config.InitDB()

	// 上传本地图片
	UploadLocalFile("默认头像", "E:\\homolog资源文件\\默认头像")



	//localFile := "C:\\Users\\elle\\Desktop\\go_workspace\\Homolog\\public\\animation1.gif"
	//localFile := "X:\\react_workspace\\gomo\\public\\avatar\\avatar3.jpg"
	//qiniu.UploadLocal(localFile, "avatar3.jpg")
}
