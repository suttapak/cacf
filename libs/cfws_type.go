package libs

import "github.com/suttapak/cacf/models"

//boardCase is a struct that holds the board and the client.
type BoardCase struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

//message form a facebook live chat.
type MsgFormFB struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	From    struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"from"`
}

//product form database
type Product struct {
	ID    uint    `json:"product_id"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

//when hub run it will has readPum for read a message form a facebook live chat,but if has a message is eqult product code,Run fucntion will send a message and product to processing later.
type MsgAndProduct struct {
	Message MsgFormFB      `json:"message"`
	Product models.Product `json:"product"`
}
type PageFacebook struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
