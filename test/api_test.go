package test

import (
	"bytes"
	"matching-system/api"
	"matching-system/model"
	"matching-system/store"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	memoryStore := store.GetInstance() // 取得 MemoryStore 實例

	r.POST("/add_single_person_and_match", func(c *gin.Context) {
		api.AddSinglePersonAndMatch(c, memoryStore)
	})
	r.DELETE("/remove_single_person/:name", func(c *gin.Context) {
		api.RemoveSinglePerson(c, memoryStore)
	})
	r.GET("/query_single_people", func(c *gin.Context) {
		api.QuerySinglePeople(c, memoryStore)
	})
	return r
}

func TestAddSinglePersonAndMatch(t *testing.T) {
	r := setupRouter()

	// 創建一個新的請求
	requestBody := bytes.NewBufferString(`{"name":"Alice", "height":170, "gender":"female", "numberOfDates":3}`)
	req, _ := http.NewRequest("POST", "/add_single_person_and_match", requestBody)

	// 發送請求並獲取響應
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestRemoveSinglePerson(t *testing.T) {
	r := setupRouter()

	// 先添加一個人
	memoryStore := store.GetInstance()
	memoryStore.AddPerson(model.NewPerson("Bob", 180, model.Male, 2))

	// 創建一個新的請求來移除這個人
	req, _ := http.NewRequest("DELETE", "/remove_single_person/Bob", nil)

	// 發送請求並獲取響應
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestQuerySinglePeople(t *testing.T) {
	r := setupRouter()

	// 先添加幾個人
	memoryStore := store.GetInstance()
	memoryStore.AddPerson(model.NewPerson("Bob", 180, model.Male, 2))
	memoryStore.AddPerson(model.NewPerson("Alice", 168, model.Female, 3))

	// 創建一個新的請求
	req, _ := http.NewRequest("GET", "/query_single_people?n=2", nil)

	// 發送請求並獲取響應
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
