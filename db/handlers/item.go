package handlers

import (
	"gomo/admin/service/dto"
	"gomo/db"
	"gomo/db/models"
	"time"
)

type ItemHandler struct {
	db.Handler
}

func (h *ItemHandler) List(list *[]models.Item) *ItemHandler {

	rows, err := h.DB.Query("select i.id, i.name, i.desc, i.preview, i.type,  i.b_link, i.tags, i.price, i.author, i.down_cnt ,i.scores, i.update_time,i.create_time, c.id, c.title from public.item i inner join public.category c on i.cat_id = c.id")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()

	for rows.Next() {
		item := models.Item{}
		err := rows.Scan(&item.ID,
			&item.Name,
			&item.Desc,
			&item.Preview,
			&item.Type,
			&item.BLink,
			&item.Tags,
			&item.Price,
			&item.Author,
			&item.DownCnt,
			&item.Scores,
			&item.UpdateTime,
			&item.CreateTime,
			&item.Cat.ID,
			&item.Cat.Title)
		if err != nil {
			_ = h.AddError(err)
			return h
		}
		*list = append(*list, item)
	}

	return h
}

//新增或者更新item
func (h *ItemHandler) Update(item *dto.ItemUpdateReq) *ItemHandler {
	var sql string
	var err error
	if item.ID == 0 {
		returnID := 0
		sql = "insert into public.item(cat_id, name, tags, desc, preview, b_link, price, author, status, update_time, create_time) values($1,$2,$3,$4,$5,$6,$7) RETURNING id"
		err := h.DB.QueryRow(sql, &item.CatId, &item.Name, &item.Tags, &item.Desc, &item.Preview, &item.BLink, &item.Price, item.Author, item.Status, time.Now(), time.Now()).
			Scan(&returnID)
		if err == nil {
			item.ID = returnID
		}
	} else {
		sql = `update public.item set 
				cat_id=$1, 
				name=$2, 
				tags=$3, 
				desc=$4, 
				preview=$5, 
				b_link=$6, 
				price=$7, 
				author=$8,
				status=$9,
				update_time=$10 
				where id = $11`
		_, err = h.DB.Exec(sql,&item.CatId, &item.Name, &item.Tags, &item.Desc, &item.Preview, &item.BLink, &item.Price, item.Author, item.Status, time.Now(), item.ID)
	}

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}
