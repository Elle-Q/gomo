package handlers

import (
	"leetroll/admin/service/vo"
	"leetroll/db"
	"leetroll/db/models"
)

type CatHandler struct {
	db.Handler
}

func (h *CatHandler) List(list *[]models.Category) *CatHandler {

	rows, err := h.DB.Query("select id, title, sub_title, preview, page_img, desp, status, create_time ,update_time from public.category")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()

	for rows.Next() {
		cat := models.Category{}
		err := rows.Scan(&cat.ID,
			&cat.Title,
			&cat.SubTitle,
			&cat.Preview,
			&cat.PageImg,
			&cat.Desc,
			&cat.Status,
			&cat.UpdateTime,
			&cat.CreateTime)
		if err != nil {
			_ = h.AddError(err)
			return h
		}
		*list = append(*list, cat)
	}

	return h
}

func (h *CatHandler) Get(id int, cat *models.Category) *CatHandler {

	row := h.DB.QueryRow("select id, title, sub_title, preview, page_img, desp, create_time ,update_time from public.category where id=$1",
		id)

	err := row.Scan(&cat.ID,
		&cat.Title,
		&cat.SubTitle,
		&cat.Preview,
		&cat.PageImg,
		&cat.Desc,
		&cat.UpdateTime,
		&cat.CreateTime)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *CatHandler) Save(cat *models.Category) *CatHandler {
	var sql string
	var err error
	if cat.ID == 0 {
		returnID := 0
		sql = "insert into public.category(title, sub_title, preview,page_img, desp, status, update_time, create_time) values($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id"
		err := h.DB.QueryRow(sql, &cat.Title, &cat.SubTitle, &cat.Preview, &cat.PageImg, &cat.Desc, &cat.Status, cat.UpdateTime, cat.CreateTime).
			Scan(&returnID)
		if err == nil {
			cat.ID = int64(returnID)
		}
	} else {
		sql = "update public.category set title=$1, sub_title=$2, preview=$3, desp=$4, status=$5, page_img=$6, update_time=$7 where id = $8"
		_, err = h.DB.Exec(sql, &cat.Title, &cat.SubTitle, &cat.Preview, &cat.Desc, &cat.Status, &cat.PageImg, cat.UpdateTime, cat.ID)
	}

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *CatHandler) Delete(catIds int) *CatHandler {
	sql := "delete from public.category where id=$1"

	_, err := h.DB.Exec(sql, catIds)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *CatHandler) ListName(list *[]vo.CatNameVO) *CatHandler {
	rows, err := h.DB.Query("select id, title from public.category")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()

	for rows.Next() {
		cat := vo.CatNameVO{}
		err := rows.Scan(&cat.ID, &cat.Title)
		if err != nil {
			_ = h.AddError(err)
			return h
		}
		*list = append(*list, cat)
	}
	return h
}
