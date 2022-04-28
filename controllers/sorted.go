package controllers

import (
	. "attempt_gin/pkg"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type SortedController struct{}

type SortedRequest struct {
	Arr []int `json:"arr"`
}

func (s SortedController) MergeSort(c *gin.Context) {

	var r SortedRequest

	if c.BindJSON(&r) == nil {
		log.Println(r)
	}

	Sort(r.Arr)
	//Mergs(&r.Arr, 0, len-1)
	c.JSON(http.StatusOK, r)
}
