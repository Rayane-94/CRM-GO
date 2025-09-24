package main

// import (
// 	"fmt"
// 	"testing"
// )

// // Fonctions ajouter manuellement pour tester
// func addContactTest(nom, email string) {
// 	contacts = append(contacts, Contact{Nom: nom, Email: email})
// }

// func modifyContactTest(index int, nom, email string) error {
// 	if index < 0 || index >= len(contacts) {
// 		return fmt.Errorf("index invalide")
// 	}
// 	contacts[index].Nom = nom
// 	contacts[index].Email = email
// 	return nil
// }

// func TestAddContact(t *testing.T) {
// 	contacts = []Contact{} // reset contacts

// 	addContactTest("toto", "toto@a.com")

// 	if len(contacts) != 1 {
// 		t.Errorf("Expected 1 contact, got %d", len(contacts))
// 	}
// 	if contacts[0].Nom != "toto" || contacts[0].Email != "toto@a.com" {
// 		t.Errorf("Contact mal ajouté: %+v", contacts[0])
// 	}
// }

// func TestModifyContact(t *testing.T) {
// 	contacts = []Contact{{Nom: "toto", Email: "toto@a.com"}}

// 	err := modifyContactTest(0, "test", "test2@a.com")
// 	if err != nil {
// 		t.Errorf("Erreur inattendue: %v", err)
// 	}

// 	if contacts[0].Nom != "test" || contacts[0].Email != "test2@a.com" {
// 		t.Errorf("Contact non modifié correctement: %+v", contacts[0])
// 	}
// }
