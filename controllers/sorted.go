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
	var sorter Sorter

	if c.BindJSON(&r) == nil {
		log.Println(r)
	}

	tupl := sorter.Sort(r.Arr)
	c.JSON(http.StatusOK, tupl)
}
