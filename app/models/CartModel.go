package models

import (
	"ecom_backend/app/dto"
	"ecom_backend/app/modules"
)

type ModelCartData struct{}

//Select get data
func (msg ModelCartData) Select(username string) ([]dto.CartData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT tuc.id, username, mi.name, COUNT(*) as count_item, id_items, price FROM tr_user_cart tuc JOIN ms_items mi on tuc.id_items = mi.id JOIN ms_user mu on tuc.id_user = mu.id WHERE username = ? GROUP BY id_items", username)
	defer rows.Close()
	var listItem []dto.CartData
	var cron dto.CartData

	for rows.Next() {
		err = rows.Scan(&cron.CartId, &cron.Username, &cron.ItemName, &cron.Qty, &cron.IdItems, &cron.Price)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, cron)
	}
		
	return listItem, nil
}

// Insert
func (msg ModelCartData) Insert(data *dto.CartTable) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO tr_user_cart (id_items, id_user) VALUES(?,?)", data.IdItems, data.IdUser)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}

//update 
func (msg ModelCartData) PurgeCart(id int) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("DELETE FROM tr_user_cart WHERE id_user = ?", id);

	if err != nil {
		return -1, err
	}
	id_ret, err := res.RowsAffected()
	return id_ret, nil
}
