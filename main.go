package main

import "bufio"
import "os"
import "strings"
import "fmt"
import "flag"

type Contact struct {
	Nom string
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
	fmt.Printf("✅ Contact ajouté: %s <%s>\n", nom, email)
}

func menuRef() {
	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1 - Ajouter un contact")
		fmt.Println("2 - Lister les contacts")
		fmt.Println("3 - Quitter")
		fmt.Print("Votre choix: ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			addContact()
		case 2:
			if len(contacts) == 0 {
				fmt.Print("pas de contact PD")
			} else {
				//afficher contact
				for i, v := range contacts {
					fmt.Printf("✅ afficher les contacts: %s <%s>\n", i+1, v.Nom, v.Email)
				}
			}
		case 3:
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
