package api

import (
	"matching-system/model"
	"matching-system/store"
	"matching-system/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddSinglePersonAndMatch 添加新用戶並尋找匹配。
func AddSinglePersonAndMatch(c *gin.Context, memoryStore *store.MemoryStore) {
	var newPerson model.Person
	if err := c.BindJSON(&newPerson); err != nil {
		utils.LogError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 檢查用戶是否有效
	if !newPerson.IsValid() {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid person")
		return
	}

	memoryStore.AddPerson(&newPerson)

	// 添加匹配邏輯
	matches := findMatches(&newPerson, memoryStore)

	c.JSON(http.StatusOK, gin.H{"person": newPerson, "matches": matches})
}

// findMatches 尋找與給定人員匹配的人員。
func findMatches(newPerson *model.Person, store *store.MemoryStore) []*model.Person {
	potentialMatches := store.GetAllPeople()
	var matches []*model.Person

	for _, person := range potentialMatches {
		if newPerson.CanMatch(person) {
			matches = append(matches, person)
			newPerson.MatchWith(person) // 更新可約會次數

			// 檢查並移除約會次數為零的用戶
			if newPerson.NumberOfDates <= 0 {
				store.RemoveIfNoDates(newPerson.Name)
				break
			}
			if person.NumberOfDates <= 0 {
				store.RemoveIfNoDates(person.Name)
			}
		}
	}

	return matches
}

// RemoveSinglePerson 從匹配系統中移除用戶。
func RemoveSinglePerson(c *gin.Context, memoryStore *store.MemoryStore) {
	name := c.Param("name")
	memoryStore.RemovePerson(name)

	c.JSON(http.StatusOK, gin.H{"message": "Person removed"})
}

// QuerySinglePeople 查詢最匹配的單身人士。
func QuerySinglePeople(c *gin.Context, memoryStore *store.MemoryStore) {
	nStr := c.DefaultQuery("n", "10")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		utils.LogError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}

	people := memoryStore.GetAllPeople()

	// TODO: 這裡可以添加排序和選擇最適合的匹配邏輯

	// 暫時簡化：僅返回前 n 個人
	if n > len(people) {
		n = len(people)
	}
	c.JSON(http.StatusOK, people[:n])
}
