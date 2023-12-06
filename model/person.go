package model

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Person struct {
	Name          string
	Height        int // cm
	Gender        Gender
	NumberOfDates int
}

// NewPerson 創建一個新的 Person 實例。
func NewPerson(name string, height int, gender Gender, numberOfDates int) *Person {
	return &Person{
		Name:          name,
		Height:        height,
		Gender:        gender,
		NumberOfDates: numberOfDates,
	}
}

// CanMatch 檢查此人是否可以與另一個人匹配。
func (p *Person) CanMatch(other *Person) bool {
	if p.Gender == other.Gender {
		return false
	}

	if p.NumberOfDates == 0 || other.NumberOfDates == 0 {
		return false
	}

	if p.Gender == Male && p.Height <= other.Height {
		return false
	}

	if p.Gender == Female && p.Height >= other.Height {
		return false
	}

	return true
}

// MatchWith 用於更新可約會次數。
func (p *Person) MatchWith(other *Person) {
	if p.CanMatch(other) {
		p.NumberOfDates--
		other.NumberOfDates--
	}
}

// IsValid 檢查 Person 是否有效。
func (p *Person) IsValid() bool {
	result := p.Name != "" &&
		p.Height > 0 &&
		(p.Gender == "male" || p.Gender == "female") &&
		p.NumberOfDates > 0

	return result
}
