package services

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
	fb "github.com/huandu/facebook/v2"
	"github.com/suttapak/cacf/libs"
	"github.com/suttapak/cacf/models"
	"github.com/suttapak/cacf/repositories"
	"gorm.io/gorm"
)

type hubServices struct {
	//ws boardcase to client
	boardCase chan libs.BoardCase
	//ws conn
	conn *websocket.Conn
	//message form facebook will send to this channel
	messageFB chan []libs.MsgFormFB
	//message and product channel
	msgAndProduct chan libs.MsgAndProduct
	//this channel has reload products form database
	reloadProduct chan bool
	//product list from database
	productList []models.Product
	//token form facebook developer
	token string
	//repositories for database
	cartRepo     repositories.CartRepository
	customerRepo repositories.CustomerRepository
	liveRepo     repositories.LiveRespository
	messageRepo  repositories.MessageRepository
	productRepo  repositories.ProductRepository
}

func NewHubServices(
	cartRepo repositories.CartRepository,
	customerRepo repositories.CustomerRepository,
	liveRepo repositories.LiveRespository,
	messageRepo repositories.MessageRepository,
	productRepo repositories.ProductRepository) HubServices {

	return &hubServices{
		boardCase:     make(chan libs.BoardCase),
		messageFB:     make(chan []libs.MsgFormFB),
		msgAndProduct: make(chan libs.MsgAndProduct),
		reloadProduct: make(chan bool),
		productList:   nil,
		//websocket connection
		conn: nil,
		//token form facebook developer
		token: "",
		//repository form database
		cartRepo:     cartRepo,
		customerRepo: customerRepo,
		liveRepo:     liveRepo,
		messageRepo:  messageRepo,
		productRepo:  productRepo,
	}
}

func (s *hubServices) Run(conn *websocket.Conn, token string) {
	defer func() {
		go s.readMessageFormFacebookLive()
		go s.compareMessageAndCode()
		go s.readMessage()
		go s.writeMessage()
		go s.addCart()
		go s.reProduct()
	}()
	products, err := s.productRepo.GetAll()
	if err != nil {
		panic(err)
	}
	s.conn = conn
	s.token = token
	s.productList = products
}

func (s *hubServices) writeMessage() {
	for {
		select {
		case message := <-s.boardCase:
			if err := s.conn.WriteJSON(message); err != nil {
				return
			}
		}
	}
}

func (s *hubServices) reProduct() {
	for {
		select {
		case <-s.reloadProduct:
			product, err := s.productRepo.GetAll()
			if err != nil {
				//TODO : Handler error.
				return
			}
			s.productList = product
		}
	}
}

func (s *hubServices) readMessage() {
	for {
		_, message, err := s.conn.ReadMessage()
		if err != nil {
			//TODO : handler error.
			break
		}
		if string(message) == "repoduct" {
			s.reloadProduct <- true
		}
	}
}

func (s *hubServices) addCart() {
	for {
		select {
		case azip := <-s.msgAndProduct:
			product := azip.Product
			message := azip.Message
			s.boardCase <- libs.BoardCase{Code: 200, Msg: fmt.Sprintf("%v %v", message.From.Name, product.Name)}
			//Check user in database.
			customerID, err := strconv.Atoi(message.From.ID)
			if err != nil {
				s.boardCase <- libs.BoardCase{Code: 200, Msg: err.Error()}
				return
			}
			if _, err := s.customerRepo.GetByID(uint(customerID)); err != nil {
				//Check of create Customer.
				if err == gorm.ErrRecordNotFound {
					//Create Customer.
					if err := s.customerRepo.Create(models.Customer{
						Model: gorm.Model{ID: uint(customerID)},
						Name:  message.From.Name,
						Email: message.From.Email,
					}); err != nil {
						//TODO : handler error.
						s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
						return
					}
					//Create Cart of Customer.
					if err := s.cartRepo.Create(models.Cart{
						CustomerID: uint(customerID),
						Discount:   0,
						Total:      0,
					}); err != nil {
						//TODO : handler error.
						s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
						return
					}

				} else {
					//TODO : handler error.
					s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
					return
				}
			}
			//Add product to cart
			//Find customer's cart.
			cart, err := s.cartRepo.GetByCustomerID(uint(customerID))
			if err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
			newProduct := append(cart.Products, product)
			//Add product to cart.
			if err := s.cartRepo.Update(models.Cart{Model: gorm.Model{ID: cart.ID}, Products: newProduct}); err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
			//Remove procut count 1 item.
			if err := s.productRepo.Update(models.Product{Model: gorm.Model{ID: product.ID}, Count: product.Count - 1}); err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
			messagePatern, err := s.messageRepo.Get("customer")
			if err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
			cart, err = s.cartRepo.GetByID(cart.ID)
			if err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
			productPlanText := ""
			for index, productInCart := range cart.Products {
				productPlanText += fmt.Sprintf(`%v. %v %v บาท\n`, index+1, productInCart.Name, productInCart.Price)
			}
			newMessage := fmt.Sprintf(messagePatern.Message, productPlanText)
			if _, err := fb.Post("/me/messages", fb.Params{
				"access_token":   s.token,
				"recipient":      fmt.Sprintf(`{ 'id' : %v }`, message.From.ID),
				"message":        fmt.Sprintf(`{ 'text' : '%v' }`, newMessage),
				"messaging_type": "MESSAGE_TAG",
				"tag":            "POST_PURCHASE_UPDATE",
			}); err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: err.Error()}
				return
			}
		}
	}
}

func (s *hubServices) compareMessageAndCode() {
	for {
		select {
		//recive message from channel.
		case msg := <-s.messageFB:
			for _, message := range msg {
				//check message len, if meesage more than 10 ,it've continue always.
				if len(message.Message) > 10 {
					continue
				}
				for _, product := range s.productList {
					//check product count.
					if product.Count <= 0 {
						//TODO : handler error.
						s.boardCase <- libs.BoardCase{Code: 400, Msg: product.Name + "Product is out of stock."}
						break
					}
					//compare message and code with if condition.
					if message.Message == product.Code {
						//send message and product to channel
						s.msgAndProduct <- libs.MsgAndProduct{
							Message: message,
							Product: product,
						}
					}
				}
			}
		}
	}
}

func (s *hubServices) readMessageFormFacebookLive() {
	//get live videos for page with selected token form client side.
	pages, err := fb.Get("me/live_videos", fb.Params{
		"access_token": s.token, //token of a page.
		"fields":       "title,status,id",
	})
	if err != nil {
		//TODO : handler error
		s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
		return
	}
	pagesJson := []libs.PageFacebook{}
	if err := pages.DecodeField("data", &pagesJson); err != nil {
		//TODO : handler error
		s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
		return
	}
	pageSelected := libs.PageFacebook{}
	for _, page := range pagesJson {
		if page.Status == "LIVE" {
			pageSelected = page
			break
		}
	}
	if len(pageSelected.ID) <= 0 {
		//TODO : handler error
		s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + "no live video"}
		return
	}
	after := ""
	liveID, _ := strconv.Atoi(pageSelected.ID)
	//Find live in database,If don't have live in database,create new live,and if have live get the live,and then add live->after to after variavle.
	if live, err := s.liveRepo.GetByID(uint(liveID)); err == gorm.ErrRecordNotFound {
		//create live to database.
		if err := s.liveRepo.Create(models.Live{Title: pageSelected.Title, After: ""}); err != nil {
			//TODO : handler error
			s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
			return
		}
	} else if err != nil {
		//TODO : handler error
		s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
		return
	} else {
		after = live.After
	}
	//get message from facebook with for infinite loop.
	messagesJson := []libs.MsgFormFB{}
	for {
		comments, err := fb.Get(fmt.Sprintf("%s/comments", pageSelected.ID), fb.Params{
			"access_token": s.token,
			"fields":       "from{name,id,email},message",
			"limit":        "10",
			"after":        after,
		})
		if err != nil {
			//TODO : handler error
			s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
			return
		}
		if err := comments.DecodeField("data", &messagesJson); err != nil {
			//TODO : handler error
			s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
			return
		}
		if len(messagesJson) <= 10 {
			//send message to channel for process message and code message.
			s.messageFB <- messagesJson
			after = comments.Get("paging.cursors.after").(string)
			//update live after.
			if err := s.liveRepo.Update(
				models.Live{
					After: after,
					Model: gorm.Model{ID: uint(liveID)},
				},
			); err != nil {
				//TODO : handler error.
				s.boardCase <- libs.BoardCase{Code: 400, Msg: "error : " + err.Error()}
				return
			}
		}
	}
}
