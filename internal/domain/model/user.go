package model

import "encoding/json"

type User struct {
	ID        int `gorm:"column:id"`
	Name      string
	Timestamp `gorm:"embedded"`
}

func (u *User) JSONTokenData() string {
	data := map[string]interface{}{
		"id":   u.ID,
		"name": u.Name,
	}

	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
