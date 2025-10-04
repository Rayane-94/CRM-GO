package cmd


import (
	"fmt"
	"github.com/spf13/cobra"
	"CRM-GO/internal/domain"
)


func init() {
	cmd := &cobra.Command{
	Use: "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
	c := domain.Contact{
		FirstName: mustGetString(cmd, "first"),
		LastName: mustGetString(cmd, "last"),
		Email: mustGetString(cmd, "email"),
		Phone: mustGetString(cmd, "phone"),
		Company: mustGetString(cmd, "company"),
		Notes: mustGetString(cmd, "notes"),	
	}
	res, err := Service.Add(c)
	if err != nil { return err }
		fmt.Printf("Créé: ID=%d %s %s\n", res.ID, res.FirstName, res.LastName)
		return nil
	},
}


cmd.Flags().String("first", "", "Prénom")
cmd.Flags().String("last", "", "Nom")
cmd.Flags().String("email", "", "Email")
cmd.Flags().String("phone", "", "Téléphone")
cmd.Flags().String("company", "", "Société")
cmd.Flags().String("notes", "", "Notes")


RootCmd().AddCommand(cmd)
}


func mustGetString(cmd *cobra.Command, name string) string {
v, _ := cmd.Flags().GetString(name)
return v
}