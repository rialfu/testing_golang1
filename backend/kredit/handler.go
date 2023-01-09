package kredit

import (
	"fmt"
	"net/http"
	"rema/kredit/custom"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
func (h *Handler) HandleGetChecklistPencairan(c *gin.Context) {
	page := c.Query("page")
	if page == ""{
		page = "1"
	}
	limit := c.Query("limit")
	if limit == ""{
		limit= "10"
	}
	// data, err:=h.Service.ServiceChecklistPencairan(page, limit)
	data, err:=h.Service.ServiceChecklistPencairan(page, limit)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		return
	}
	fmt.Println(data)
	c.JSON(200, gin.H{"message":"sukses", "data":map[string]any{
		"data":data.data,
		"total":data.total,
		"total_page":data.total_halaman,
	}})
	return 
}

func (h *Handler) HandleApprovePencairan(c *gin.Context) {
	var req UpdateChecklistPencairan
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	err := h.Service.UpdateChecklistPencairan(req)
	if err != nil{
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message":"sukses",
	})
}

func (h *Handler) HandleGetDataReport(c *gin.Context) {
	page := c.Query("page")
	if page == ""{
		page = "1"
	}
	limit := c.Query("limit")
	if limit == ""{
		limit= "10"
	}
	start := c.Query("start")
	_, err := time.Parse("2006-01-02", start)
	if err != nil{
		start = ""
	}
	end := c.Query("end")
	_, err = time.Parse("2006-01-02", end)
	if err != nil{
		end = ""
	}
	company := c.Query("company")
	approval := c.Query("approval")
	if approval !="1" && approval != "0"{
		approval = ""
	}
	branch := c.Query("branch")
	
	data, err:=h.Service.ServiceDataReport(page, limit, start, end, company, branch, approval)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		return
	}
	
	c.JSON(200, gin.H{"message":"sukses", "data":map[string]any{
		"data":data.data,
		"total":data.total,
		"total_page":data.total_halaman,
	}})
	return 
}
func (h *Handler) HandleGetSearch(c *gin.Context){
	data, err := h.Service.GetCompanyService()
	
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		return
	}
	dataBranch, err := h.Service.GetBranchService()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"dataCompany":data, "dataBranch":dataBranch})
}