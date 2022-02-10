/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os/exec"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the CTcontainer binary file",
	Long:  `Install the CTcontainer binary file to the /usr/local/bin directory`,
	Run: func(cmd *cobra.Command, args []string) {
		installCtc()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func installCtc() {
	if err := install("./shell/install.sh"); err != nil {
		panic("Error occurred during installing ctcontainer-go")
	}
	color.Green("Installed successfully")
}

func install(shellPath string) error {
	install := exec.Command("/bin/bash", shellPath)
	err := install.Run()
	return err
}
