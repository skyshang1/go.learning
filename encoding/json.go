package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Mail struct {
	Value string
}

func (m *Mail) UnmarshalJSON(data []byte) error {
	if !bytes.Contains(data, []byte("@")) {
		return fmt.Errorf("mail format error")
	}
	m.Value = string(data)
	return nil
}

func (m *Mail) MarshalJSON() (data []byte, err error) {
	if m != nil {
		data = []byte(m.Value)
	}
	return data, nil
}

type Phone struct {
	Value string
}

func (p *Phone) UnmarshalJSON(data []byte) error {
	if len(data) != 11 {
		return fmt.Errorf("phone format error")
	}
	p.Value = string(data)
	return nil
}

func (p *Phone) MarshalJSON() (data []byte, err error) {
	if p != nil {
		data = []byte(p.Value)
	}
	return
}

type UserRequest struct {
	Name  string
	Mail  Mail
	Phone Phone
}

func main() {
	user := UserRequest{
		Name:  "sky",
		Mail:  Mail{Value: "skyshang@gmail.com"},
		Phone: Phone{Value: "18813137113"},
	}

	info, _ := json.Marshal(user)
	fmt.Println(string(info))

	userStr := `{"Name":"ysy","Mail":{"Value":"skyshanggmail.com"},"Phone":{"Value":"18813137113"}}`
	var user1 UserRequest
	err := json.Unmarshal([]byte(userStr), &user1)
	fmt.Println(err) // Warn: mail format error
}
