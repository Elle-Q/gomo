package main

import (
	"gomo/config"
)

func main() {
	config.Setup("config/config.yml")
	config.InitDB()

	// 上传本地图片
	UploadLocalFile("背景图片", "X:\\react_workspace\\gomo\\src\\assets\\bg")



	//localFile := "C:\\Users\\elle\\Desktop\\go_workspace\\Homolog\\public\\animation1.gif"
	//localFile := "X:\\react_workspace\\gomo\\public\\avatar\\avatar3.jpg"
	//qiniu.UploadLocal(localFile, "avatar3.jpg")
}
