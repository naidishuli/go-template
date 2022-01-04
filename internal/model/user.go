package model

import "encoding/json"

type User struct {
	Id        int `gorm:"column:id"`
	Name      string
	Timestamp `gorm:"embedded"`
}

func (u *User) JSONTokenData() string {
	data := map[string]interface{}{
		"id":   u.Id,
		"name": u.Name,
	}

	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
