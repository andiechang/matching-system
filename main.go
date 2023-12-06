package main

import (
	"matching-system/api"
	"matching-system/store"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 初始化記憶體存儲
	memoryStore := store.NewMemoryStore()

	// 設置路由
	r.POST("/add_single_person_and_match", func(c *gin.Context) {
		api.AddSinglePersonAndMatch(c, memoryStore)
	})
	r.DELETE("/remove_single_person/:name", func(c *gin.Context) {
		api.RemoveSinglePerson(c, memoryStore)
	})
	r.GET("/query_single_people", func(c *gin.Context) {
		api.QuerySinglePeople(c, memoryStore)
	})

	// 啟動服務器
	r.Run() // 預設監聽並在 0.0.0.0:8080 上啟動服務
}
