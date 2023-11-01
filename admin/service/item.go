package service

import (
	"leetroll/admin/service/vo"
	"leetroll/db/handlers"
	"leetroll/db/models"
	"leetroll/qiniu"
	"leetroll/tool"
)

type ItemService struct {
	ItemHandler    *handlers.ItemHandler
	FileHandler    *handlers.FileHandler
	ChapterHandler *handlers.ChapterHandler
	Error          error
}

func (e *ItemService) GetFilesByItemId(ID int, filesVO *vo.ItemFilesVO) *ItemService {
	itemHandler := e.ItemHandler
	fileHandler := e.FileHandler
	chapterHandler := e.ChapterHandler

	//获取item的相关信息(名称)
	item := models.MakeItem()
	err := itemHandler.Get(ID, item).Error
	if err != nil {
		e.Error = err
		return e
	}

	files := make([]models.File, 0)
	//获取文件信息
	fileHandler.QueryByItemId(ID, &files)

	main := make([]models.File, 0)
	prev := make([]models.File, 0)
	attachment := make([]models.File, 0)

	for _, f := range files {
		p := &f
		if tool.IsVideo(f.Format) {
			p.QnLink = qiniu.GetPrivateUrlForM3U8(f.Key)
		} else {
			p.QnLink = qiniu.GetPrivateUrl(f.Key)
		}
		switch f.Type {
		case "main":
			main = append(main, f)
			break
		case "preview":
			prev = append(prev, f)
			break
		case "attachment":
			attachment = append(attachment, f)
			break
		}
	}

	//获取章节信息
	chapterVOs := e.getChapters(ID, chapterHandler, fileHandler)

	filesVO.Main = main
	filesVO.Preview = prev
	filesVO.Attachment = attachment
	filesVO.ID = int64(ID)
	filesVO.ItemName = item.Name
	filesVO.RescType = item.Type
	filesVO.Chapters = chapterVOs

	return e
}

func (e *ItemService) getChapters(ID int, chapterHandler *handlers.ChapterHandler, fileHandler *handlers.FileHandler) []vo.ChapterVO {
	chapterVOs := make([]vo.ChapterVO, 0)
	chapters := make([]models.Chapter, 0)
	chapterHandler.QueryByItemId(ID, &chapters)

	for _, ch := range chapters {
		chapterVO := vo.ChapterVO{}
		chapterMain := make([]models.File, 0)
		episodes := make([]models.File, 0)

		episodeIds := chapterHandler.QueryEpisodeIds(int(ch.ID))
		fileHandler.QueryByIds(episodeIds, &episodes)
		fileHandler.QueryByIds([]int64{ch.Main}, &chapterMain)

		e.setQnLink(chapterMain)
		e.setQnLink(episodes)

		chapterVO.Chapter = ch.Chapter
		chapterVO.ID = ch.ID
		chapterVO.Main = chapterMain
		chapterVO.Episodes = episodes

		chapterVOs = append(chapterVOs, chapterVO)
	}
	return chapterVOs
}

func (e *ItemService) setQnLink(files []models.File) {
	for i, _ := range files {
		file := files[i]
		if tool.IsVideo(file.Format) {
			files[i].QnLink = qiniu.GetPrivateUrlForM3U8(file.Key)
		} else {
			files[i].QnLink = qiniu.GetPrivateUrl(file.Key)
		}
	}
}
