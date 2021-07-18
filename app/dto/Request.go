package dto

type LoginRequest struct {
	Username		string       `json:"username"`
	Password		string      `json:"password"`
}

type UpdateCartRequest struct {
	Cart []struct{
		IdItems		int      `json:"id_items"`
		Qty			int      `json:"qty"`
	}
}

type CheckoutRequest struct {
	Address		string       `json:"address"`
}

type GetTransactionResp struct {
	PurchaseKey		string      `json:"purchase_key" db:"purchase_key"`
	Address			string    	`json:"address" db:"address"`
	Total			string    	`json:"total" db:"total"`
	Status			int  		`json:"status" db:"status"`
	Detail			[]DetailData `json:"detail" db:"detail"`
}