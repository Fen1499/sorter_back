package controllers

import (
	. "sorter_gin/pkg"

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
	var sorter ISorter = new(Sorter)

	if c.BindJSON(&r) == nil {
		log.Println(r)
	}

	sorter.Sort(r.Arr)
	c.JSON(http.StatusOK, sorter.GetResult())
}
