package handlers

import (
	"database/sql"
	"fmt"
	"gomo/db"
	"gomo/db/models"
	"time"
)

type FileHandler struct {
	db.Handler
}

//根据item_id查找所有文件
func (h *FileHandler) QueryByItemId(itemId int, list *[]models.File) *FileHandler {

	rows, err := h.DB.Query("select id, type, item_id, name, size, format, bucket, key, mark1, update_time, create_time from public.file where item_id=$1", itemId)

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	defer rows.Close()
	parseErr := parseRows(rows, list)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}
	return h

}

//根据ids查找所有文件
func (h *FileHandler) QueryByIds(ids string, list *[]models.File) *FileHandler {

	sql := fmt.Sprintf("select id, type, item_id, name,size,format,bucket,key,mark1, update_time,create_time from public.file where id IN(%s)", ids)
	rows, err := h.DB.Query(sql)

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	defer rows.Close()
	parseErr := parseRows(rows, list)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}
	return h

}

//所有文件
func (h *FileHandler) List(list *[]models.File) *FileHandler {

	rows, err := h.DB.Query("select id, type, item_id, name,qn_link,size,format,bucket,key,update_time, create_time from public.file")

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	defer rows.Close()
	parseErr := parseRows(rows, list)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}
	return h

}

//所有文件
func (h *FileHandler) Save(model *models.File) *FileHandler {
	returnID := 0
	sqlS := "insert into public.file(item_id, type, name, size, format, bucket, key, mark1, update_time, create_time) " +
		"values($1, $2 ,$3 ,$4 ,$5 ,$6 ,$7,$8,$9) RETURNING id"
	err := h.DB.QueryRow(sqlS,
		&model.ItemId,
		&model.Type,
		&model.Name,
		&model.Size,
		&model.Format,
		&model.Bucket,
		&model.Key,
		time.Now(),
		time.Now(),
		) .Scan(&returnID)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	model.ID = int64(returnID)
	return h

}

func (h *FileHandler) Delete(id int) *FileHandler{
	sql := "delete from public.file where id=$1"

	_, err := h.DB.Exec(sql, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}


//解析rows为list
func parseRows(rows *sql.Rows, list *[]models.File) (err error){
	for rows.Next() {
		file := models.File{}
		err := rows.Scan(
			&file.ID,
			&file.Type,
			&file.ItemId,
			&file.Name,
			&file.Size,
			&file.Format,
			&file.Bucket,
			&file.Key,
			&file.Mark,
			&file.UpdateTime,
			&file.CreateTime,
		)
		if err != nil {
			return err
		}
		*list = append(*list, file)
	}
	return nil
}