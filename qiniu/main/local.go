package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"leetroll/common/runtime"
	"leetroll/config"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
	"leetroll/qiniu/regular"
	"leetroll/tool"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

// 上传本地文件夹, 并保存文件信息到数据库 (一些配置类数据)
func UploadLocalDir(typeConfig string, dir string) {
	fmt.Printf(tool.Red("==================准备上传 '%s'===================="), typeConfig)
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
			key := regular.UploadLocal(filePath, file.Name(), typeConfig)

			var model models.File
			model.Key = key
			model.Bucket = config.QiniuConfig.PubBucket
			model.Type = typeConfig
			makFileModel(file, &model)

			err := service.Save(&model).Error
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%s  --- %s", tool.Green(filePath), key)
			fmt.Println()
		}()
	}

	wg.Wait()
	fmt.Println(tool.Red("完成"))
}

// fixme:未完成的方法 (更新数据库七牛资源的有效时间) [暂时不用]
func UpdateLinkExpirrationInDB() {
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
func makFileModel(file fs.FileInfo, model *models.File) {
	model.ItemId = 0
	model.Format = filepath.Ext(file.Name())
	model.Name = file.Name()[:strings.Index(file.Name(), ".")]
	model.Size = float32(file.Size())
}
