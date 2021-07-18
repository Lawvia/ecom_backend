package models

import (
	"ecom_backend/app/dto"
	"ecom_backend/app/modules"
)


type ModelHeaderData struct{}

//Select get data
func (msg ModelHeaderData) SelectHeader(key int) ([]dto.HeaderData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT purchase_key, address, total, status FROM tr_header_purchase WHERE id_user = ?", key)
	defer rows.Close()
	var listItem []dto.HeaderData
	var cron dto.HeaderData

	for rows.Next() {
		err = rows.Scan(&cron.PurchaseKey, &cron.Address, &cron.Total, &cron.Status)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, cron)
	}
		
	return listItem, nil
}

func (msg ModelHeaderData) SelectDetail(key string) ([]dto.DetailData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT id_items, name, qty FROM tr_detail_purchase tdp JOIN ms_items mi on mi.id = tdp.id_items WHERE header_key = ?", key)
	defer rows.Close()
	var listItem []dto.DetailData
	var cron dto.DetailData

	for rows.Next() {
		err = rows.Scan(&cron.IdItems, &cron.ItemName, &cron.Qty)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, cron)
	}
		
	return listItem, nil
}

// Insert
func (msg ModelHeaderData) InsertHeader(data *dto.HeaderData) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO tr_header_purchase (purchase_key, id_user, address, total, status) VALUES(?,?,?,?,?)", data.PurchaseKey, data.IdUser, data.Address, data.Total, 3)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}

func (msg ModelHeaderData) InsertDetail(data *dto.DetailData) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO tr_detail_purchase (id_items, qty, header_key) VALUES(?,?,?)", data.IdItems, data.Qty, data.HeaderId)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}
