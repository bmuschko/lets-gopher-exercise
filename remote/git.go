package remote

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmuschko/lets-gopher-exercise/utils"
	git "gopkg.in/src-d/go-git.v4"
)

type GitRepo struct {
	RepoUrl, TargetPath string
}

func (g *GitRepo) Install() {
	repoName := parseRepoName(g.RepoUrl)
	checkoutDir := filepath.Join(g.TargetPath, repoName)

	if _, err := os.Stat(checkoutDir); os.IsNotExist(err) {
		clone(g.RepoUrl, checkoutDir)
	} else {
		pull(checkoutDir)
	}
}

func clone(repoUrl string, targetPath string) *git.Repository {
	log.Printf("Cloning template from repository %s", repoUrl)
	repo, err := git.PlainClone(targetPath, false, &git.CloneOptions{
		URL:               repoUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	utils.CheckIfError(err)
	return repo
}

func pull(targetPath string) {
	log.Printf("Pulling latest changes for existing template in %s", targetPath)
	r, err := git.PlainOpen(targetPath)
	utils.CheckIfError(err)

	w, err := r.Worktree()
	utils.CheckIfError(err)

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})

	// Avoid handling the error if the worktree is up-to-date
	if err != git.NoErrAlreadyUpToDate {
		utils.CheckIfError(err)
	}
}

func parseRepoName(url string) string {
	lastDotSlash := strings.LastIndex(url, "/")
	lastDotIndex := strings.LastIndex(url, ".git")
	r := []rune(url)
	return string(r[lastDotSlash:lastDotIndex])
}
