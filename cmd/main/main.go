package main

import (
	"fmt"

	"github.com/apavanello/game-of-life/versioninfo"
)

func main() {
	fmt.Println("Project: "+ versioninfo.ProjectName)
	fmt.Println("Description: "+ versioninfo.ProjectDescription)
	fmt.Println("Copyright: "+ versioninfo.ProjectCopyright)
	fmt.Println("Version: "+ versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)
}
