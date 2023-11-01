package handlers

import (
	"database/sql"
	"leetroll/db"
	"leetroll/db/models"
	"time"
)

type ChapterHandler struct {
	db.Handler
}

// 保存章节信息
func (h *ChapterHandler) Save(model *models.Chapter) *ChapterHandler {
	//新增资源
	if model.ID == -1 {
		return h.Add(model)
	} else {
		return h.Update(model)
	}
}

func (h *ChapterHandler) Update(model *models.Chapter) *ChapterHandler {
	var err error
	if 0 != model.Main {
		sql := `update public.chapter set main=$1 where id = $2`
		_, err = h.DB.Exec(sql, model.Main, model.ID)
	}
	if err != nil {
		_ = h.AddError(err)
	}
	return h
}

func (h *ChapterHandler) GetById(id int64, model *models.Chapter) *ChapterHandler {

	row := h.DB.QueryRow("select id, item_id, chapter, main, update_time, create_time from public.chapter where id=$1", id)

	if row.Err() != nil {
		_ = h.AddError(row.Err())
		return h
	}
	err := row.Scan(
		&model.ID,
		&model.ItemId,
		&model.Chapter,
		&model.Main,
		time.Now(),
		time.Now())
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *ChapterHandler) Add(model *models.Chapter) *ChapterHandler {
	returnID := 0
	sqlS := "insert into public.chapter(item_id, chapter, main, update_time, create_time) " +
		"values($1, $2 ,$3 ,$4 ,$5) RETURNING id"
	err := h.DB.QueryRow(sqlS,
		&model.ItemId,
		&model.Chapter,
		&model.Main,
		time.Now(),
		time.Now(),
	).Scan(&returnID)

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	model.ID = int64(returnID)
	return h
}

func (h *ChapterHandler) QueryByItemId(itemId int, list *[]models.Chapter) *ChapterHandler {
	rows, err := h.DB.Query("select id, item_id,chapter, main, update_time, create_time from public.chapter where item_id=$1", itemId)

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	defer rows.Close()
	parseErr := parseChapter(rows, list)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}
	return h
}

func (h *ChapterHandler) DelFileByType(chapterId int, type_ string, fileId int) {
	var err error
	if type_ == "Main" {
		sql := `update public.chapter set main=$1 where id = $2`
		_, err = h.DB.Exec(sql, -1, chapterId)
	} else {
		// 保存章节关联关系
		sql := `delete from public.chapter_file where chapter_id = $1 and file_id = $2`
		_, err = h.DB.Exec(sql, chapterId, fileId)
	}
	if err != nil {
		_ = h.AddError(err)
	}
}

// 保存章节关联关系
func (h *ChapterHandler) SaveChapterEpisode(chapterId int, episodeFileIds []int64) *ChapterHandler {
	for _, episodeFileId := range episodeFileIds {
		sqlS := "insert into public.chapter_file(chapter_id, file_id) " +
			"values($1, $2)"
		_, err := h.DB.Query(sqlS,
			chapterId,
			episodeFileId)

		if err != nil {
			_ = h.AddError(err)
			return h
		}
	}
	return h
}

// 获取章节关联的所有文件信息
func (h *ChapterHandler) QueryEpisodeIds(chapterId int) []int64 {
	rows, err := h.DB.Query("select file_id from public.chapter_file where chapter_id=$1", chapterId)
	if err != nil {
		_ = h.AddError(err)
	}
	episodeFileIds := make([]int64, 0)
	defer rows.Close()
	for rows.Next() {
		var fileId int64
		err = rows.Scan(&fileId)
		episodeFileIds = append(episodeFileIds, fileId)
	}
	if err != nil {
		_ = h.AddError(err)
	}
	return episodeFileIds
}

// 解析rows为list
func parseChapter(rows *sql.Rows, list *[]models.Chapter) (err error) {
	for rows.Next() {
		chapter := models.Chapter{}
		err := rows.Scan(
			&chapter.ID,
			&chapter.ItemId,
			&chapter.Chapter,
			&chapter.Main,
			&chapter.UpdateTime,
			&chapter.CreateTime,
		)
		if err != nil {
			return err
		}
		*list = append(*list, chapter)
	}
	return nil
}
