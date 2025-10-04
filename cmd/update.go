package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"CRM-GO/internal/domain"
)

func init() {
	cmd := &cobra.Command{
	Use: "update",
	Short: "Mettre à jour un contact par ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, _ := cmd.Flags().GetUint("id")
		c := domain.Contact{
			ID: id,
			FirstName: mustGetString(cmd, "first"),
			LastName: mustGetString(cmd, "last"),
			Email: mustGetString(cmd, "email"),
			Phone: mustGetString(cmd, "phone"),
			Company: mustGetString(cmd, "company"),
			Notes: mustGetString(cmd, "notes"),
		}
		if err := Service.Update(c); err != nil { return err }
		fmt.Println("Mise à jour ok")
		return nil
	},
}
	cmd.Flags().Uint("id", 0, "ID du contact (requis)")
	_ = cmd.MarkFlagRequired("id")
	cmd.Flags().String("first", "", "Prénom")
	cmd.Flags().String("last", "", "Nom")
	cmd.Flags().String("email", "", "Email")
	cmd.Flags().String("phone", "", "Téléphone")
	cmd.Flags().String("company", "", "Société")
	cmd.Flags().String("notes", "", "Notes")
	RootCmd().AddCommand(cmd)
}