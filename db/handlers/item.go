package handlers

import (
	"database/sql"
	"github.com/lib/pq"
	"leetroll/admin/service/dto"
	"leetroll/db"
	"leetroll/db/models"
	"strings"
	"time"
)

type ItemHandler struct {
	db.Handler
}

func (h *ItemHandler) List(list *[]models.Item) *ItemHandler {

	rows, err := h.DB.Query("select i.id, i.name, i.desp, i.type,  i.b_link, i.tags, i.price::decimal, i.author, i.down_cnt ,i.scores, i.update_time,i.create_time, c.id, c.title from public.item i inner join public.category c on i.cat_id = c.id")

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	defer rows.Close()
	parseErr := parseItemRows(rows, list)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}

	return h
}

// 新增或者更新item
func (h *ItemHandler) Update(item *dto.ItemUpdateReq) *ItemHandler {
	var sql string
	var err error
	tags := strings.Split(item.Tags, ",")
	if item.ID == 0 {
		returnID := 0
		sql = `insert into public.item(cat_id, name, type, tags, desp, b_link, price, author, status, update_time, create_time)
				values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
		err := h.DB.QueryRow(sql, &item.CatId, &item.Name, "", pq.Array(tags), &item.Desc, &item.BLink, &item.Price, item.Author, item.Status, time.Now(), time.Now()).
			Scan(&returnID)
		if err == nil {
			item.ID = returnID
		} else {
			_ = h.AddError(err)
			return h
		}
	} else {
		sql = `update public.item set  cat_id=$1,  name=$2,  tags=$3,  desp=$4,  b_link=$5,  price=$6, 
  				author=$7, status=$8, update_time=$9  where id = $10`
		_, err = h.DB.Exec(sql, &item.CatId, &item.Name, pq.Array(tags), &item.Desc, &item.BLink, &item.Price, item.Author, item.Status, time.Now(), item.ID)

		if err != nil {
			_ = h.AddError(err)
			return h
		}
	}

	return h
}

func (h *ItemHandler) Delete(id int64) *ItemHandler {
	sql := "delete from public.item where id=$1"

	_, err := h.DB.Exec(sql, id)

	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 根据id查询item
func (h *ItemHandler) Get(id int, item *models.Item) *ItemHandler {

	row := h.DB.QueryRow("select i.id, i.name, i.desp, i.type,  i.b_link, i.tags, i.price::decimal, i.author, i.down_cnt ,i.scores, i.update_time,i.create_time, i.cat_id from public.item i  where i.id=$1",
		id)

	if row.Err() != nil {
		_ = h.AddError(row.Err())
		return h
	}

	err := row.Scan(
		&item.ID,
		&item.Name,
		&item.Desc,
		&item.Type,
		&item.BLink,
		pq.Array(&item.Tags),
		&item.Price,
		&item.Author,
		&item.DownCnt,
		&item.Scores,
		&item.UpdateTime,
		&item.CreateTime,
		&item.Cat.ID)
	if err != nil {
		_ = h.AddError(err)
		return h
	}
	return h
}

// 查询分类下热门的item (4条 用来首页做展示)
func (h *ItemHandler) ListPopularByCat(catID int64, maxSize int, items *[]models.Item) *ItemHandler {
	rows, err := h.DB.Query(`select i.id, i.name, i.desp, CASE WHEN f.key IS NULL THEN '' ELSE f.key END, i.type, i.b_link, i.tags, i.price::decimal,
			i.author, i.down_cnt ,i.scores, i.update_time,i.create_time, 0, '' from public.item i left join 
			    (select * from public.file where type = 'main') f on f.item_id = i.id where i.cat_id =$1 limit $2`,
		catID, maxSize)

	if err != nil {
		_ = h.AddError(err)
		return h
	}

	defer rows.Close()
	parseErr := parseItemRowsWithMain(rows, items)

	if parseErr != nil {
		_ = h.AddError(parseErr)
		return h
	}
	return h
}

// 解析rows为list
func parseItemRows(rows *sql.Rows, list *[]models.Item) (err error) {
	for rows.Next() {
		item := models.Item{}
		item.Cat = &models.Category{}
		err := rows.Scan(&item.ID,
			&item.Name,
			&item.Desc,
			&item.Type,
			&item.BLink,
			pq.Array(&item.Tags),
			&item.Price,
			&item.Author,
			&item.DownCnt,
			&item.Scores,
			&item.UpdateTime,
			&item.CreateTime,
			&item.Cat.ID,
			&item.Cat.Title,
		)
		if err != nil {
			return err
		}
		*list = append(*list, item)
	}
	return nil
}

func parseItemRowsWithMain(rows *sql.Rows, list *[]models.Item) (err error) {
	for rows.Next() {
		item := models.Item{}
		item.Cat = &models.Category{}
		err := rows.Scan(&item.ID,
			&item.Name,
			&item.Desc,
			&item.Main,
			&item.Type,
			&item.BLink,
			pq.Array(&item.Tags),
			&item.Price,
			&item.Author,
			&item.DownCnt,
			&item.Scores,
			&item.UpdateTime,
			&item.CreateTime,
			&item.Cat.ID,
			&item.Cat.Title,
		)
		if err != nil {
			return err
		}
		*list = append(*list, item)
	}
	return nil
}
