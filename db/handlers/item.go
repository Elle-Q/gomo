package handlers

import (
	"gomo/db"
	"gomo/db/models"
)

type ItemHandler struct {
	db.Handler
}

func (h *ItemHandler) List(list *[]models.Item) *ItemHandler {

	rows , err := h.DB.Query("select i.id, i.name, i.desc, i.preview, i.type,  i.b_link, i.tags, i.price, i.author, i.down_cnt ,i.scores, i.update_time,i.create_time, c.id, c.title from public.item i inner join public.category c on i.cat_id = c.id")

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
