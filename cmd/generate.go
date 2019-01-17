package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/bmuschko/lets-gopher-exercise/templ"
	"github.com/bmuschko/lets-gopher-exercise/utils"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/core"
)

func doGenerateCmd(cmd *cobra.Command, args []string) {
	config, err := loadConfig()
	utils.CheckIfError(err)

	availableTemplates := templ.ListTemplates()
	availableTemplateNames := templateNames(availableTemplates)

	if len(availableTemplateNames) > 0 {
		core.SetFancyIcons()
		selectedTemplateName := promptTemplate(availableTemplateNames)
		defaultBasePath := buildDefaultBasePath(config, selectedTemplateName)
		enteredBasePath := promptBasePath(defaultBasePath)

		templ.GenerateProject(availableTemplates[selectedTemplateName], enteredBasePath)
	} else {
		log.Print("No templates found!")
	}
}

func promptTemplate(templateNames []string) string {
	selectedTemplate := ""
	prompt := &survey.Select{
		Message: "What template would you like to use to generate a project?",
		Options: templateNames,
	}
	err := survey.AskOne(prompt, &selectedTemplate, nil)
	utils.CheckIfError(err)

	fmt.Printf("You choose %q\n", selectedTemplate)
	return selectedTemplate
}

func promptBasePath(defaultBasePath string) string {
	enteredBasePath := ""
	prompt := &survey.Input{
		Message: "What base path would like to use?",
		Default: defaultBasePath,
	}
	err := survey.AskOne(prompt, &enteredBasePath, survey.MinLength(1))
	utils.CheckIfError(err)

	fmt.Printf("You choose %q\n", enteredBasePath)
	return enteredBasePath
}

func buildDefaultBasePath(genConfig templ.GenConfig, selectedTemplateName string) string {
	defaultBasePath := ""

	if genConfig.Domain != "" {
		defaultBasePath += genConfig.Domain

		if !strings.HasSuffix(defaultBasePath, "/") {
			defaultBasePath += "/"
		}

		defaultBasePath += selectedTemplateName
	}

	return defaultBasePath
}

func templateNames(availableTemplates map[string]string) []string {
	templateNames := make([]string, 0, len(availableTemplates))
	for k := range availableTemplates {
		templateNames = append(templateNames, k)
	}
	sort.Strings(templateNames)
	return templateNames
}

func loadConfig() (templ.GenConfig, error) {
	config, err := templ.Load()
	if err != nil && os.IsNotExist(err) {
		return templ.GenConfig{}, nil
	}

	return config, nil
}
