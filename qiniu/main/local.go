package main

import (
	"fmt"
	"gomo/common/runtime"
	"gomo/config"
	"gomo/db/handlers"
	"gomo/db/models"
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

func makFileModel(file fs.FileInfo, link string, model *models.File) {
	model.Type = "avatar"
	model.Format = filepath.Ext(file.Name())
	model.Name = file.Name()[:strings.Index(file.Name(), ".")]
	model.Size = float32(file.Size())
	model.QnLink = link
}