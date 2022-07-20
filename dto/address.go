package dto

type AddressReply struct {
	ID          uint   `json:"id"`
	Address     string `json:"address"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	PostalCode  string `json:"postal_code"`
}
type CreateAddress struct {
	Address     string `json:"address" binding:"required"`
	SubDistrict string `json:"sub_district" binding:"required"`
	District    string `json:"district" binding:"required"`
	Province    string `json:"province" binding:"required"`
	PostalCode  string `json:"postal_code" binding:"required"`
	CustomerID  uint   `json:"customer_id" `
	ShopID      uint   `json:"shop_id"`
}
