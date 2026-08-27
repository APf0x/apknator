package cmd

import (
	"embed"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

//go:embed template/*
var templateFS embed.FS

var template = &cobra.Command{
	Use:   "template",
	Short: "Builds a simple template for a simple helloworld Activity",
	Run: func(cmd *cobra.Command, args []string) {
		template_init()
	},
}

func init() {
	rootCmd.AddCommand(template)
}

func template_init() {
	subFS, err := fs.Sub(templateFS, "template")
	if err != nil {
		panic(err)
	}
	e := os.CopyFS("my-app", subFS)
	if e != nil {
		panic(e)
	}
}
