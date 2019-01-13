package templ

import (
	"fmt"
	"github.com/bmuschko/lets-gopher/utils"
	"go/build"
	"os"
	"path/filepath"
)

func GenerateProject(templatePath string, goHomeBasePath string) {
	srcGoPath := filepath.Join(build.Default.GOPATH, "src")
	targetPath := filepath.Join(srcGoPath, goHomeBasePath)

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		fmt.Printf("The target directory %q already exists\n", targetPath)
		os.Exit(1)
	}

	fmt.Printf("Generating project in %q\n", targetPath)
	utils.CopyDir(templatePath, targetPath)
}
