package handlers

import (
	"leetroll/app/service/dto"
	"leetroll/db"
	"leetroll/db/models"
)

type UserHandler struct {
	db.Handler
}

func (h *UserHandler) FindById(id int, model *models.User) *UserHandler {

	row := h.DB.QueryRow("select id, name, phone, qr_code,address,gender,vip,bg_imag,admin,status,avatar,moto,update_time,create_time from public.user where id=$1", id)

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
		&model.Avatar,
		&model.Moto,
		&model.UpdateTime,
		&model.CreateTime,
	)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *UserHandler) FindByUserName(req *dto.UserApiReq, model *models.User) *UserHandler {

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

// 更新用户信息
func (h *UserHandler) Update(model *models.User) *UserHandler {
	sql := "update public.user set Name=$1, Address=$2, Gender=$3, Status=$4, Moto=$5, update_time=$6 where id = $7"
	_, err := h.DB.Exec(sql,
		&model.Name,
		&model.Address,
		&model.Gender,
		&model.Status,
		&model.Moto,
		&model.UpdateTime,
		&model.ID)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 用户注册
func (h *UserHandler) Register(model *models.User) *UserHandler {
	sql := "insert into public.user(Phone, Name) values($1, $2)"
	_, err := h.DB.Exec(sql,
		&model.Phone,
		&model.Name)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 修改绑定手机
func (h *UserHandler) updatePhone(phone string, id int) *UserHandler {
	sql := "update public.user set Phone= $1 where id=$2"
	_, err := h.DB.Exec(sql, phone, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 修改绑定微信
func (h *UserHandler) updateQRCode(QRCode string, id int) *UserHandler {
	sql := "update public.user set QRCode= $1 where id=$2"
	_, err := h.DB.Exec(sql, QRCode, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 设置vip
func (h *UserHandler) updateVip(vip bool, id int) *UserHandler {
	sql := "update public.user set vip= $1 where id=$2"
	_, err := h.DB.Exec(sql, vip, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

func (h *UserHandler) FindUserByPhone(u *dto.UserLoginApiReq, model *models.User) *UserHandler {

	row := h.DB.QueryRow("select id, name, phone,avatar, qr_code,address,gender,vip,bg_imag,admin from public.user where phone=$1", u.UserName)

	err := row.Scan(&model.ID,
		&model.Name,
		&model.Phone,
		&model.Avatar,
		&model.QRCode,
		&model.Address,
		&model.Gender,
		&model.Vip,
		&model.BgImag,
		&model.Admin,
	)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 更新用户头像
func (h *UserHandler) UpdateAvatar(id int, link string) *UserHandler {
	sql := "update public.user set avatar= $1 where id=$2"
	_, err := h.DB.Exec(sql, link, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 更新用户背景
func (h *UserHandler) UpdateBG(id int, link string) *UserHandler {
	sql := "update public.user set bg_imag= $1 where id=$2"
	_, err := h.DB.Exec(sql, link, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}
