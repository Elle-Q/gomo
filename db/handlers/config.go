package handlers

import (
	"gomo/db"
	"gomo/db/models"
)

type ConfigHandler struct {
	db.Handler
}

func (h *ConfigHandler) FindByName(name string, model *models.Config) *ConfigHandler {

	row := h.DB.QueryRow("select id, name, val, create_time ,update_time from public.config where name=$1",
		name)

	err := row.Scan(&model.ID,
		&model.Name,
		&model.Val,
		&model.UpdateTime,
		&model.CreateTime)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}
