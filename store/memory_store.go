package store

import (
	"matching-system/model"
	"sync"
)

// MemoryStore 為儲存單身人士資料的結構。
type MemoryStore struct {
	mu     sync.RWMutex
	people map[string]*model.Person
}

// NewMemoryStore 創建一個新的 MemoryStore 實例。
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		people: make(map[string]*model.Person),
	}
}

// AddPerson 新增一個單身人士到儲存系統。
func (store *MemoryStore) AddPerson(person *model.Person) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.people[person.Name] = person
}

// RemovePerson 從儲存系統中移除一個單身人士。
func (store *MemoryStore) RemovePerson(name string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.people, name)
}

// GetPerson 根據名字獲取一個單身人士的資料。
func (store *MemoryStore) GetPerson(name string) (*model.Person, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	person, exists := store.people[name]
	return person, exists
}

// GetAllPeople 返回儲存系統中所有單身人士的列表。
func (store *MemoryStore) GetAllPeople() []*model.Person {
	store.mu.RLock()
	defer store.mu.RUnlock()

	peopleList := make([]*model.Person, 0, len(store.people))
	for _, person := range store.people {
		peopleList = append(peopleList, person)
	}
	return peopleList
}

// RemoveIfNoDates 檢查並移除約會次數為零的用戶。
func (store *MemoryStore) RemoveIfNoDates(name string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	person, exists := store.people[name]
	if exists && person.NumberOfDates <= 0 {
		delete(store.people, name)
	}
}
