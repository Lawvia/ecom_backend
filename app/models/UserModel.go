package models

import (
	"ecom_backend/app/dto"
	"ecom_backend/app/modules"
)

//ModelUserData is a class of model ms_games
type ModelUserData struct{}

//Select get data
func (msg ModelUserData) CheckUsername(username string) (*dto.UserData, error) {
	db := new(modules.DbConnect).GetDB()
	rows, err := db.Query("SELECT id, username, password, active FROM ms_user WHERE username = ?", username)
	defer rows.Close()
	rows.Next()
	cron := new(dto.UserData)
	err = rows.Scan(&cron.UserId, &cron.Username, &cron.Password, &cron.Active)
	if err != nil {
		return nil, err
	}
	return cron, nil
}

// Insert
func (msg ModelUserData) Insert(data *dto.UserData) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("INSERT INTO ms_user (username, password, active) VALUES(?,?,1)", data.Username, data.Password)

	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	return id, nil
}

func (msg ModelUserData) UpdateLogin(id int) (int64, error) {
	db := new(modules.DbConnect).GetDB()
	res, err := db.Exec("UPDATE ms_user SET last_login = NOW() WHERE id = ?", id);

	if err != nil {
		return -1, err
	}
	id_ret, err := res.RowsAffected()
	return id_ret, nil
}
