package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"CRM-GO/internal/store"
	"CRM-GO/internal/store/app"
	jsonstore "CRM-GO/internal/store/json"
	memstore "CRM-GO/internal/store/memory"
	gormstore "CRM-GO/internal/store/gorm"
	
)



var (
	rootCmd = &cobra.Command{
		Use: "crm-final",
		Short: "Mini CRM en ligne de commande",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return initConfig() },
	}

	Service *app.ContactService
)


func Execute() {
	if err := rootCmd.Execute(); err != nil {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
}



type Config struct {
	Type string `mapstructure:"type"` 
	DBPath string `mapstructure:"db_path"`
	JSONPath string `mapstructure:"json_path"`
}


func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")


	viper.SetDefault("type", "memory")
	viper.SetDefault("db_path", "mini-crm.db")
	viper.SetDefault("json_path", "contacts.json")


	_ = viper.ReadInConfig()


	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil { return err }


	var s store.Storer
	var err error
	switch cfg.Type {
	case "gorm":
		s, err = gormstore.New(cfg.DBPath)
	case "json":
		s, err = jsonstore.New(cfg.JSONPath)
	case "memory":
		s = memstore.New()
	default:
		return fmt.Errorf("type de stockage inconnu: %s", cfg.Type)
	}	
	if err != nil { return err }


	Service = app.NewContactService(s)
	return nil
}


func RootCmd() *cobra.Command { return rootCmd }