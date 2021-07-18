package models

import (
	"ecom_backend/app/dto"
	"ecom_backend/app/modules"
)


type ModelWhitelistData struct{}

//Select get data
func (msg ModelWhitelistData) CheckList(ip string) (*dto.WhitelistData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT id, name, ip FROM ms_whitelist WHERE ip = ?", ip)
	defer rows.Close()
	rows.Next()
	cron := new(dto.WhitelistData)
	err = rows.Scan(&cron.WhitelistId, &cron.IpName, &cron.Ip)
	if err != nil {
		return nil, err
	}
	return cron, nil
}

// Insert
func (msg ModelWhitelistData) Insert(data *dto.WhitelistData) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO ms_whitelist (name, ip) VALUES(?,?)", data.IpName, data.Ip)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}
