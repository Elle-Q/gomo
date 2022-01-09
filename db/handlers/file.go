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

	rows, err := h.DB.Query("select id, type, item_id, name,qn_link,size,format,update_time, create_time from public.file where item_id=$1", itemId)

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

	sql := fmt.Sprintf("select id, type, item_id, name,qn_link,size,format,update_time, create_time from public.file where id IN(%s)", ids)
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

	rows, err := h.DB.Query("select id, type, item_id, name,qn_link,size,format,update_time, create_time from public.file")

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
func (h *FileHandler) Save(model models.File) *FileHandler {

	sql := "insert into public.file(type, name, qn_link, size, format, update_time, create_time) " +
		"values($1, $2 ,$3 ,$4 ,$5 ,$6 ,$7) "
	_, err := h.DB.Exec(sql,
		&model.Type,
		&model.Name,
		&model.QnLink,
		&model.Size,
		&model.Format,
		time.Now(),
		time.Now(),
		)

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
			&file.QnLink,
			&file.Size,
			&file.Format,
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