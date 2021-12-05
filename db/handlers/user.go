package handlers

import (
	"gomo/app/service/dto"
	"gomo/db"
	"gomo/db/models"
)

type UserHandler struct {
	db.Handler
}

func (h *UserHandler) Find( req *dto.UserApiReq, model *models.User) *UserHandler {

	row := h.DB.QueryRow("select id, name, phone, qr_code,address,gender,vip,bg_imag,admin,update_time, create_time from public.user where id=$1", req.GetId())

	err := row.Scan(&model.ID,
		&model.Name,
		&model.Phone,
		&model.QRCode,
		&model.Address,
		&model.Gender,
		&model.Vip,
		&model.BgImag,
		&model.Admin,
		&model.UpdateTime,
		&model.CreateTime,
	)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}
