package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)


func init() {
	cmd := &cobra.Command{
		Use: "list",
		Short: "Lister les contacts",
		RunE: func(cmd *cobra.Command, args []string) error {
			items, err := Service.List()
			if err != nil { return err }
			if len(items) == 0 { fmt.Println("Aucun contact"); return nil }
			for _, c := range items {
				fmt.Printf("%d\t%s\t%s\t%s\t%s\n", c.ID, c.FirstName, c.LastName, c.Email, c.Phone)
			}
			return nil
			},
		}
	RootCmd().AddCommand(cmd)
}