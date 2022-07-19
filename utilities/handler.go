package utilities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/cacf/errs"
)

type Reply struct {
	Code       int
	MessageOBJ interface{}
}

type ErrorMessage struct {
	Message string   `json:"message"`
	Error   []string `json:"error"`
}

type FundationMessage struct {
	Message string `json:"message"`
}

type GetRequest struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type any interface {
}

func HandlerReply(statusCode int, obj any, c *gin.Context) (int, any) {
	switch e := obj.(type) {
	case errs.AppError:
		return e.HttpCode, ErrorMessage{
			Message: e.Msg,
			Error:   []string{e.Error()},
		}
	case error:
		return 500, ErrorMessage{
			Message: "internal server error.",
			Error:   []string{"internal server error"},
		}
	default:
		switch statusCode {
		case http.StatusOK:
			switch c.Request.Method {
			case "GET":
				return http.StatusOK, GetRequest{
					Message: "success",
					Data:    obj,
				}
			case "POST":
				return http.StatusCreated, FundationMessage{Message: "created successfully."}
			case "PUT":
				return http.StatusOK, FundationMessage{Message: "updated successfully."}
			case "DELETE":
				return http.StatusOK, FundationMessage{Message: "deleted successfully."}
			default:
				return http.StatusOK, GetRequest{
					Message: "success",
					Data:    obj,
				}
			}
		case http.StatusCreated:
			return http.StatusCreated, FundationMessage{Message: "created successfully."}
		case http.StatusNoContent:
			return http.StatusNoContent, FundationMessage{Message: "no content."}
		case http.StatusBadRequest:
			return http.StatusBadRequest, ErrorMessage{
				Message: "bad request.",
				Error:   []string{"bad request"},
			}
		case http.StatusUnauthorized:
			return http.StatusUnauthorized, ErrorMessage{
				Message: "unauthorized.",
				Error:   []string{"unauthorized"},
			}
		case http.StatusForbidden:
			return http.StatusForbidden, ErrorMessage{
				Message: "forbidden.",
				Error:   []string{"forbidden"},
			}
		case http.StatusNotFound:
			return http.StatusNotFound, ErrorMessage{
				Message: "not found.",
				Error:   []string{"not found"},
			}
		case http.StatusInternalServerError:
			return http.StatusInternalServerError, ErrorMessage{
				Message: "internal server error.",
				Error:   []string{"internal server error"},
			}

		default:
			return http.StatusInternalServerError, ErrorMessage{
				Message: "internal server error.",
				Error:   []string{"internal server error"},
			}
		}
	}
}
