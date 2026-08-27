package cmd

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

const currentVersion = "0.0.1"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update ap CLI to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		v := semver.MustParse(currentVersion)
		selfupdate.EnableLog()

		latest, err := selfupdate.UpdateSelf(v, "apf0x/apknator")
		if err != nil {
			fmt.Printf("Binary update failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(latest.Version)
		if latest.Version.Equals(v) {
			fmt.Printf("Current version (%s) is up to date.\n", currentVersion)
		} else {
			fmt.Printf("Successfully updated to version %s!\n", latest.Version)
		}
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "gives what version ap you are curretly using",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(version)
}
