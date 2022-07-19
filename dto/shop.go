package dto

type ShopReply struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Summary    string `json:"summary"`
	FacebookID uint   `json:"facebook_id"`
}

type UpdateShopDTO struct {
	ID      uint   `json:"id"`
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Summary string `json:"summary" binding:"required"`
}
