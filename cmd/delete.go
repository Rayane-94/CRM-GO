package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


func init() {
	cmd := &cobra.Command{
		Use: "delete",
		Short: "Supprimer un contact par ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmd.Flags().GetUint("id")
			if err := Service.Delete(id); err != nil { return err }
			fmt.Println("Supprimé")
			return nil
		},
	}
	cmd.Flags().Uint("id", 0, "ID du contact (requis)")
	_ = cmd.MarkFlagRequired("id")
	RootCmd().AddCommand(cmd)
}