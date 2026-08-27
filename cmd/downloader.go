package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var aap2_download_link string = "https://dl.google.com/dl/android/maven2/com/android/tools/build/aapt2/8.2.0-10154469/aapt2-8.2.0-10154469-linux.jar"
var p8jar_download_link string = "https://dl.google.com/dl/android/maven2/com/android/tools/r8/8.2.33/r8-8.2.33.jar"
var apksigner_jar_download_link string = "https://github.com/patrickfav/uber-apk-signer/releases/download/v1.3.0/uber-apk-signer-1.3.0.jar"

func android_jar_download_link(version string) string {
	return "https://github.com/Sable/android-platforms/raw/master/android-" + version + "/android.jar"
}

var download = &cobra.Command{
	Use:   "download",
	Short: "Installs the necessary tools needed to compile the apk",
	Run: func(cmd *cobra.Command, args []string) {
		tool_downloader("36")
	},
}

func init() {
	rootCmd.AddCommand(download)
}

func tool_downloader(version string) {
	exec.Command("sh", "-c", "mkdir -p ~/android-sdk/bin ~/android-sdk/platforms/android-"+version).Run()
	aap2_downloader()
	p8jar_downloader()
	android_jar_downloader(version)
	apksigner_jar_downloader()
}

func aap2_downloader() {
	cmd := exec.Command("sh", "-c", "wget "+aap2_download_link+" -O /tmp/aapt2.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("aap2 download Failed step 1: %v\n", err)
		return
	}
	exec.Command("sh", "-c", "unzip /tmp/aapt2.zip aapt2 -d ~/android-sdk/bin/").Run()
	exec.Command("sh", "-c", "chmod +x ~/android-sdk/bin/aapt2").Run()
	exec.Command("sh", "-c", "rm /tmp/aapt2.zip").Run()

}

func p8jar_downloader() {
	cmd := exec.Command("sh", "-c", "wget "+p8jar_download_link+" -O ~/android-sdk/bin/r8.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("r8 download Failed: %v\n", err)
		return
	}
}

func android_jar_downloader(version string) {
	cmd := exec.Command("sh", "-c", "wget "+android_jar_download_link(version)+" -O ~/android-sdk/platforms/android-"+version+"/android.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("android.jar download Failed: %v\n", err)
		return
	}
}

func apksigner_jar_downloader() {
	cmd := exec.Command("sh", "-c", "wget "+apksigner_jar_download_link+" -O ~/android-sdk/bin/uber-apk-signer.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("apksigner error: %v\n", err)
		return
	}
	fmt.Print("success")

}
