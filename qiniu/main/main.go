package main

import (
	"gomo/config"
	"gomo/qiniu"
)

func main() {
	config.Setup("config/config.yml")

	//localFile := "C:\\Users\\elle\\Desktop\\go_workspace\\Homolog\\public\\animation1.gif"
	localFile := "X:\\react_workspace\\gomo\\public\\avatar\\avatar3.jpg"
	qiniu.UploadLocal(localFile, "avatar3.jpg")
}
