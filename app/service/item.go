package service

import (
	"leetroll/app/service/vo"
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

func (e *ItemService) GetItemAndFilesByItemId(ID int, vo *vo.ItemWithFilesVO) *ItemService {
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

	if len(main) > 0 {
		vo.Main = main[0].QnLink
	}
	vo.Preview = prev
	vo.Attachment = attachment
	vo.ID = int64(ID)
	vo.Name = item.Name
	vo.Type = item.Type
	vo.Desc = item.Desc
	vo.Author = item.Author
	vo.Scores = item.Scores
	vo.Price = item.Price
	vo.DownCnt = item.DownCnt
	vo.Tags = item.Tags
	vo.CatID = item.Cat.ID
	vo.CatTitle = item.Cat.Title
	vo.Chapters = chapterVOs
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
		if len(chapterMain) > 0 {
			chapterVO.Main = chapterMain[0].QnLink
		}
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

// 获取章节信息
func (e *ItemService) GetChapters(itemId int, vos *[]vo.ChapterVO) *ItemService {
	fileHandler := e.FileHandler
	chapterHandler := e.ChapterHandler

	//查询item下的所有章节id
	chapters := make([]models.Chapter, 0)
	chapterHandler.QueryByItemId(itemId, &chapters)
	for _, chapter := range chapters {
		episodes := make([]models.File, 0)
		episodeIds := chapterHandler.QueryEpisodeIds(int(chapter.ID))
		fileHandler.QueryByIds(episodeIds, &episodes)
		e.setQnLink(episodes)

		chapterVo := vo.ChapterVO{}
		chapterVo.ID = chapter.ID
		chapterVo.Chapter = chapter.Chapter
		chapterVo.Episodes = episodes
		*vos = append(*vos, chapterVo)
	}
	return e
}
