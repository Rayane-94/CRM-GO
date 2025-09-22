package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Nom   string
	Email string
}

var contacts []Contact

func addContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nom: ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	nom = strings.TrimSpace(nom)
	email = strings.TrimSpace(email)

	contacts = append(contacts, Contact{Nom: nom, Email: email})
	fmt.Printf("Contact ajouté: %s <%s>\n", nom, email)
}

func modifyContact() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact à modifier")
		return
	}

	// Afficher les contacts existants abser sur addContact
	for i, v := range contacts {
		fmt.Printf("%d: %s <%s>\n", i+1, v.Nom, v.Email)
	}

	
	var index int
	fmt.Print("Index du contact à modifier: ")
	fmt.Scanln(&index)
	index = index - 1 // index commencant par 1 

	if index < 0 || index >= len(contacts) {
		fmt.Println("Index invalide")
		return
	}

	fmt.Print("Nouveau nom: ")
	fmt.Scanln(&contacts[index].Nom)

	fmt.Print("Nouvel email: ")
	fmt.Scanln(&contacts[index].Email)

	fmt.Printf("Contact modifié: %s <%s>\n", contacts[index].Nom, contacts[index].Email)
}

func menuRef() {
	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1 - Ajouter un contact")
		fmt.Println("2 - Lister les contacts")
		fmt.Println("3 - Modifier un contact")
		fmt.Println("4 - Quitter")
		fmt.Print("Votre choix: ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			addContact()
		case 2:
			if len(contacts) == 0 {
				fmt.Print("aucun contact enregistré")
			} else {
				//afficher contact
				for i, v := range contacts {
					fmt.Printf("Contact %d: %s <%s>\n", i+1, v.Nom, v.Email)
				}
			}
		case 3:
			var index int
			fmt.Print("Index du contact à modifier: ")
			fmt.Scanln(&index)
			modifyContact()

		case 4:
			fmt.Print("stop script")
			//todo quitter le process
			os.Exit(0)
		default:
			fmt.Print("mauvais choix")
		}
	}
}

func main() {
	ajouter := flag.Bool("ajouter", false, "Ajouter un contact directement")
	flag.Parse()

	if *ajouter {
		addContact()
		return
	}

	menuRef()
}
