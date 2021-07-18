package models

import (
	"ecom_backend/app/dto"
	"ecom_backend/app/modules"
)

//ModelItemsData is a class of model ms_games
type ModelItemsData struct{}

//Select get data
func (msg ModelItemsData) Select() ([]dto.ItemsData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT id, name, description, price, stock FROM ms_items")
	defer rows.Close()
	var listItem []dto.ItemsData
	var cron dto.ItemsData

	for rows.Next() {
		err = rows.Scan(&cron.ItemsId, &cron.ItemName, &cron.ItemDescription, &cron.ItemPrice, &cron.ItemStock)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, cron)
	}
		
	return listItem, nil
}

// Insert
func (msg ModelItemsData) Insert(data *dto.ItemsData) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO ms_items (name, description, price, stock) VALUES(?,?,?,?)", data.ItemName, data.ItemDescription, data.ItemPrice, data.ItemStock)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}

//update stock when purchased
func (msg ModelItemsData) UpdateStock(id int, qty int) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("UPDATE ms_items SET stock = (stock - ?) WHERE id = ?", qty, id);

	if err != nil {
		return -1, err
	}
	id_ret, err := res.RowsAffected()
	return id_ret, nil
}
