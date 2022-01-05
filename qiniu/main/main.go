package main

import (
	"gomo/config"
	"gomo/qiniu"
)

func main() {
	config.Setup("config/config.yml")

	localFile := "C:\\Users\\elle\\Desktop\\go_workspace\\Homolog\\public\\animation1.gif"

	qiniu.UploadLocal(localFile, "animation1.gif")
}
