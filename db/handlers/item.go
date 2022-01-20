package handlers

import (
	"gomo/db"
	"gomo/db/models"
)

type ItemHandler struct {
	db.Handler
}

func (h *ItemHandler) List(list *[]models.File) *ItemHandler {

	rows , err := h.DB.Query("select id, title, sub_title, preview, desp, status, create_time ,update_time from public.category")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()

	for rows.Next() {
		cat := models.Config{}
		err := rows.Scan(&cat.ID,
			&cat.Title,
			&cat.SubTitle,
			&cat.Preview,
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
