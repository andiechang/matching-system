package test

import (
	"matching-system/model"
	"matching-system/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStore(t *testing.T) {
	// 創建一個新的 MemoryStore 實例
	store := store.NewMemoryStore()

	// 測試添加用戶
	person := model.NewPerson("Alice", 170, model.Female, 3)
	store.AddPerson(person)
	retrievedPerson, exists := store.GetPerson("Alice")
	assert.True(t, exists)
	assert.Equal(t, person, retrievedPerson)

	// 測試獲取所有用戶
	people := store.GetAllPeople()
	assert.Equal(t, 1, len(people))
	assert.Equal(t, "Alice", people[0].Name)

	// 測試移除用戶
	store.RemovePerson("Alice")
	_, exists = store.GetPerson("Alice")
	assert.False(t, exists)
}

func TestRemoveIfNoDates(t *testing.T) {
	store := store.NewMemoryStore()

	person := model.NewPerson("Bob", 180, model.Male, 1)
	store.AddPerson(person)

	// 減少約會次數直到達到0，然後檢查是否被移除
	person.NumberOfDates = 0
	store.RemoveIfNoDates("Bob")
	_, exists := store.GetPerson("Bob")
	assert.False(t, exists)
}
