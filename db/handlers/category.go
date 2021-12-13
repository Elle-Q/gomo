package handlers

import (
	"gomo/app/service/dto"
	"gomo/db"
	"gomo/db/models"
)

type CatHandler struct {
	db.Handler
}

func (h *CatHandler) List(list *[]models.Category) *CatHandler {

	rows , err := h.DB.Query("select id, title, sub_title, preview, desp, create_time ,update_time from public.category")

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
			&cat.Desc,
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

func (h *CatHandler) Get(req *dto.CatApiReq, cat *models.Category)  *CatHandler{

	row := h.DB.QueryRow("select id, title, sub_title, preview, desp, create_time ,update_time from public.category where id=$1", req.GetId())

	err := row.Scan(&cat.ID,
		&cat.Title,
		&cat.SubTitle,
		&cat.Preview,
		&cat.Desc,
		&cat.UpdateTime,
		&cat.CreateTime)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h

}
