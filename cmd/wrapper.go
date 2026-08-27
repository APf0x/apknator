package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "build function of the ap tool, builds the Activity generating a signed .apk executable",
	Long:  `Builds Android APK executable directly from raw sources using SDK command-line tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		wrapper_build()
	},
}

func init() {
	rootCmd.AddCommand(build)
}

func runCmd(cmdStr string) {
	fmt.Printf("\n--> Executing: %s\n", cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("\n Build failed during: %s\nError: %v\n", cmdStr, err)
		os.Exit(1)
	}
}

func wrapper_build() {
	runCmd(`rm -rf build dist && mkdir -p build/gen build/obj build/apk dist`)

	runCmd(`~/android-sdk/bin/aapt2 compile --dir res -o build/compiled_res.zip`)

	runCmd(`~/android-sdk/bin/aapt2 link -I ~/android-sdk/platforms/android-36/android.jar --manifest AndroidManifest.xml -R build/compiled_res.zip --auto-add-overlay -o build/unsigned_base.apk --java build/gen/`)

	runCmd(`javac -source 8 -target 8 -classpath ~/android-sdk/platforms/android-36/android.jar -d build/obj/ $(find src build/gen -type f -name "*.java")`)

	runCmd(`java -cp ~/android-sdk/bin/r8.jar com.android.tools.r8.D8 --lib ~/android-sdk/platforms/android-36/android.jar --output build/apk $(find build/obj -type f -name "*.class")`)

	runCmd(`zip -j build/unsigned_base.apk build/apk/classes.dex`)

	runCmd(`java -jar ~/android-sdk/bin/uber-apk-signer.jar --apks build/unsigned_base.apk --out dist/`)

	fmt.Println("\n Signed APK available in dist/")
}
