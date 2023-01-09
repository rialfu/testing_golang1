package auth

import (
	"fmt"
	"net/http"
	"rema/kredit/custom"

	"github.com/gin-gonic/gin"
)

// const (
// 	message = "Input data not suitable"
// )

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
func (h *Handler) CreateUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	_, status, err := h.Service.CreateAccount(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

}
func (h *Handler) UpdatePassword(c *gin.Context){
	var req UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	fmt.Println(req)
	code, err :=h.Service.UpdatePassword(req.OldPassword, req.NewPassword, req.Username)
	fmt.Println(err)
	if err != nil{
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message":"Berhasil",
	})
}
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	res, status, err := h.Service.Login(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := custom.GenerateJWT(res.Username, res.Name)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"token": token,
		"data": map[string]any{
			"name":      res.Name,
			"username":  res.Username,
		},
	})
	return
}


