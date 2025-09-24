package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContactJSON struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"` 
}

func SaveToJSON(filename string, contacts []*Contact) error {
	var data []ContactJSON
	for _, c := range contacts {
		data = append(data, ContactJSON{ID: c.ID, Name: c.Name, Email: c.Email})
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonBytes, 0644)
}


func LoadFromJSON(filename string) ([]*Contact, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []ContactJSON
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	var contacts []*Contact
	for _, d := range data {
		contacts = append(contacts, &Contact{ID: d.ID, Name: d.Name, Email: d.Email})
	}

	return contacts, nil
}

func PrintJSON(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lecture:", err)
		return
	}
	fmt.Println(string(data))
}
