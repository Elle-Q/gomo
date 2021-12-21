package handlers

import (
	"gomo/app/service/dto"
	"gomo/db"
	"gomo/db/models"
)

type UserHandler struct {
	db.Handler
}

func (h *UserHandler) Find(req *dto.UserApiReq, model *models.User) *UserHandler {

	row := h.DB.QueryRow("select id, name, phone, qr_code,address,gender,vip,bg_imag,admin,status,update_time, create_time from public.user where id=$1", req.GetId())

	err := row.Scan(&model.ID,
		&model.Name,
		&model.Phone,
		&model.QRCode,
		&model.Address,
		&model.Gender,
		&model.Vip,
		&model.BgImag,
		&model.Admin,
		&model.Status,
		&model.UpdateTime,
		&model.CreateTime,
	)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *UserHandler) List(list *[]models.User) *UserHandler {
	rows, err := h.DB.Query("select id, name, phone, qr_code,address,gender,vip,bg_imag,admin,status, update_time, create_time from public.user")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()

	for rows.Next() {
		model := models.User{}
		err := rows.Scan(&model.ID,
			&model.Name,
			&model.Phone,
			&model.QRCode,
			&model.Address,
			&model.Gender,
			&model.Vip,
			&model.BgImag,
			&model.Admin,
			&model.Status,
			&model.UpdateTime,
			&model.CreateTime,
		)
		if err != nil {
			_ = h.AddError(err)
			return h
		}
		*list = append(*list, model)
	}

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

//更新用户信息
func (h *UserHandler) Update(model *models.User) *UserHandler{
	sql := "update public.user set Name=$1, Phone=$2, QRCode=$3, Address=$4, Gender=$5, BgImag=$6, Avatar=$7, update_time=$8 where id = $9"
	_, err := h.DB.Exec(sql,
		&model.Name,
		&model.Phone,
		&model.QRCode,
		&model.Address,
		&model.Gender,
		&model.BgImag,
		&model.Avatar,
		&model.UpdateTime,
		&model.ID)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

//用户注册
func (h *UserHandler) Register(model *models.User) *UserHandler{
	sql := "insert into public.user(Phone, Name) values($1, $1)"
	_, err := h.DB.Exec(sql,
		&model.Phone,
		&model.Name)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}


//修改绑定手机
func (h *UserHandler) updatePhone(phone string, id int) *UserHandler{
	sql := "update public.user set Phone= $1 where id=$2"
	_, err := h.DB.Exec(sql,phone, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}


//修改绑定微信
func (h *UserHandler) updateQRCode(QRCode string, id int) *UserHandler{
	sql := "update public.user set QRCode= $1 where id=$2"
	_, err := h.DB.Exec(sql,QRCode, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

//设置vip
func (h *UserHandler) updateVip(vip bool, id int) *UserHandler{
	sql := "update public.user set vip= $1 where id=$2"
	_, err := h.DB.Exec(sql,vip, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}
