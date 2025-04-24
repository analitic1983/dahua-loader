package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"koshmin/dahua-loader/config"
	"koshmin/dahua-loader/database"
	"koshmin/dahua-loader/services"
	"log"
)

// Show all commands
func printAllCommands(cmd *cobra.Command, prefix string) {
	// Collect all commands
	var commands []struct {
		Name        string
		Description string
	}
	collectCommands(cmd, &commands, "")

	// Calc max command length
	maxWidth := 0
	for _, c := range commands {
		if len(c.Name) > maxWidth {
			maxWidth = len(c.Name)
		}
	}

	// Print command with prefix
	for _, c := range commands {
		fmt.Printf("  %-*s  %s\n", maxWidth, c.Name, c.Description)
	}
}

// Collect commands list
func collectCommands(cmd *cobra.Command, commands *[]struct {
	Name        string
	Description string
}, prefix string) {
	commandColor := color.FgGreen
	if prefix != "" {
		commandColor = color.FgYellow
	}
	colored := color.New(commandColor).SprintFunc()

	for _, c := range cmd.Commands() {
		name := prefix + colored(c.Use)
		*commands = append(*commands, struct {
			Name        string
			Description string
		}{
			Name:        name,
			Description: c.Short,
		})

		if len(c.Commands()) > 0 {
			collectCommands(c, commands, "  ")
		}
	}
}

func main() {
	config.Init()

	rootCmd := &cobra.Command{
		Use:   "dahua_loader",
		Short: "Download videos from dahua cameras",
	}

	// Migrations
	migrationCmd := &cobra.Command{
		Use:   "migrations",
		Short: "Manage migrations",
	}

	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Run available migrations",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrationsUp()

		},
	}
	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Rollback one last migration",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrationsDown()
		},
	}
	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Migrations status",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrationsStatus()
		},
	}

	// Add subcommands
	migrationCmd.AddCommand(upCmd, downCmd, statusCmd)
	rootCmd.AddCommand(migrationCmd)

	// Check new videos/photos
	checkNewVideos := &cobra.Command{
		Use:   "download-new-video",
		Short: "1. Connect to cameras 2. Download new videos/photos 3. Convert to mp4 4. Save to download folder",
		Run: func(cmd *cobra.Command, args []string) {
			_ = services.CheckForNewVideos()
		},
	}
	rootCmd.AddCommand(checkNewVideos)

	// Create admin
	createAdmin := &cobra.Command{
		Use:   "create-admin",
		Short: "has two parameters: email, password",
		Run: func(cmd *cobra.Command, args []string) {
			email, _ := cmd.Flags().GetString("email")
			password, _ := cmd.Flags().GetString("password")
			if email == "" {
				log.Fatalln("Empty \"email\" parameter not allowed")
				return
			}
			if password == "" {
				log.Fatalln("Empty \"password\" parameter not allowed")
				return
			}

			if err := database.InitDB(); err != nil {
				log.Fatalln("Failed db connect: ", err)
			}

			_, err := services.AddAdmin(email, password)
			if err != nil {
				log.Fatalln("Failed create user:", err.Error())
			}
			log.Println("User created.")
		},
	}
	createAdmin.Flags().StringP("email", "", "", "Email address")
	createAdmin.Flags().StringP("password", "", "", "Password")
	rootCmd.AddCommand(createAdmin)

	// Override help
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Printf("  %s [command]\n\n", cmd.Use)
		fmt.Println("Available Commands:")
		printAllCommands(cmd, "")
	})

	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
	rootCmd.Execute()

}
