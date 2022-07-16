package dto

type SignInDTO struct {
	ID        uint   `json:"id" binding:"required"`
	UserToken string `json:"token" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type SignInReply struct {
	Data []struct {
		AccessToken string `json:"access_token"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Picture     struct {
			Data struct {
				Url string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
	} `json:"data"`
}
