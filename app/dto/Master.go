package dto

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type ItemsData struct {
	ItemsId			int       	`json:"id" db:"id"`
	ItemName		string      `json:"name" db:"name"`
	ItemDescription	string    	`json:"description" db:"description"`
	ItemPrice		string    	`json:"price" db:"price"`
	ItemStock		int  		`json:"stock" db:"stock"`
}

type UserData struct {
	UserId			int       	`json:"id" db:"id"`
	Username		string      `json:"username" db:"username"`
	Password		string    	`json:"password" db:"password"`
	Active			int    		`json:"active" db:"active"`
	LastLogin		time.Time  	`json:"last_login" db:"last_login"`
}

type WhitelistData struct {
	WhitelistId		int       	`json:"id" db:"id"`
	IpName			string      `json:"name" db:"name"`
	Ip				string    	`json:"ip" db:"ip"`
}

type CartData struct {
	CartId			int       	`json:"id" db:"id"`
	Username		string      	`json:"username" db:"username"`
	ItemName		string    		`json:"item_name" db:"name"`
	IdItems			int      	`json:"id_items" db:"id_items"`
	Qty 			int 		`json:"qty" db:"count_item"`
	Price 			string		`json:"price" db:"price"`
}

type CartTable struct {
	IdItems			int      	`json:"id_items" db:"id_items"`
	IdUser			int      	`json:"id_user" db:"id_user"`
}

type HeaderData struct {
	HeaderId		int       	`json:"id" db:"id"`
	PurchaseKey		string      `json:"purchase_key" db:"purchase_key"`
	IdUser			int 	   	`json:"id_user" db:"id_user"`
	Address			string    	`json:"address" db:"address"`
	Total			string    	`json:"total" db:"total"`
	Status			int  		`json:"status" db:"status"`
}

type DetailData struct {
	IdItems			int      	`json:"item_id" db:"id_items"`
	ItemName		string      `json:"item_name" db:"name"`
	Qty				int 	   	`json:"qty" db:"qty"`
	HeaderId		string  	`json:"header_key" db:"header_key"`
}

type Token struct {
	UserId string
	jwt.StandardClaims
}


