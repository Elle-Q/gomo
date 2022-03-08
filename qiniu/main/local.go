package main

import (
	"fmt"
	"gomo/common/runtime"
	"gomo/config"
	"gomo/db/handlers"
	"gomo/db/models"
	"gomo/qiniu"
	"gomo/qiniu/regular"
	"gomo/tool"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

//上传本地文件夹, 并保存文件信息到数据库 (一些配置类数据)
func UploadLocalDir(name string, dir string)  {
	fmt.Printf(tool.Red("==================准备上传 '%s'===================="), name)
	tool.PrettyPrint(*config.QiniuConfig)

	service := &handlers.FileHandler{}
	service.DB = runtime.App.GetDb()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		wg.Add(1)
		file := file
		go func() {
			defer wg.Done()

			filePath := filepath.Join(dir, file.Name())
			link := regular.UploadLocal(filePath, file.Name())

			var model models.File
			makFileModel(file, link, &model)
			err := service.Save(&model).Error
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%s  --- %s", tool.Green(filePath), link)
			fmt.Println()
		}()
	}

	wg.Wait()
	fmt.Println(tool.Red("完成"))
}

//fixme:未完成的方法 (更新数据库七牛资源的有效时间) [暂时不用]
func UpdateLinkExpirrationInDB()  {
	//查找所有file, 对私有空间的file进行处理
	fmt.Printf(tool.Red("==================准备更新私有空间文件有效期 '%s'===================="))
	tool.PrettyPrint(*config.QiniuConfig)

	service := &handlers.FileHandler{}
	service.DB = runtime.App.GetDb()

	files := make([]models.File, 0)
	service.List(&files)

	for _, file := range files {
		file.QnLink = qiniu.GetPrivateUrl(file.Key)
	}

}
func makFileModel(file fs.FileInfo, link string, model *models.File) {
	model.Type = "avatar"
	model.Format = filepath.Ext(file.Name())
	model.Name = file.Name()[:strings.Index(file.Name(), ".")]
	model.Size = float32(file.Size())
	model.QnLink = link
}