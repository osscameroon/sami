package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func main() {
	if len(os.Args) == 0 {
		println("[x] No command provided")
		return
	}

	switch command := os.Args[1]; command {
	case "deploy":
		println("[-] sammy deploy command... ")
		dir := os.Args[2]
		deploy(dir)
	case "restart-nginx":
		println("[-] sammy restart-nginx command...")
	case "stop":
		println("[-] sammy stop command...")
	case "update-field":
		println("[-] sammy update-field command...")
	default:
		println("[x] No valid command provided...")
		helperFunc()
	}

}

func deploy(dir string) {
	fmt.Printf("Deploying %s\n", dir)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: %s: %s\n", dir, err)
			return
		} else {
			fmt.Printf("Error: unexpected error occurred for %s: %s\n", dir, err)
			return
		}
	}

	listDeployFolders(dir)

	gitPull(dir)

	// We run the docker-compose here
	for _, p := range getDockerComposeFiles(dir) {
		exec.Command("docker-compose", "-f "+p+" up -d")
	}

	// we run the docker stack deploy here
	for _, p := range getDockerStackFiles(dir) {
		exec.Command("docker stack deploy", "-c "+p)
	}
}

func listDeployFolders(dir string) {
	servicesDir := dir + "/services"
	confDir := dir + "/conf"

	fmt.Printf("Config folder %s\n", confDir)
	listContent(confDir)

	fmt.Println("Services folder %s\n", servicesDir)
	listContent(servicesDir)
}

func getDockerComposeFiles(dir string) []string {
	return findFile(dir, []string{"*-compose.yml", "*-compose.yaml"})
}

func getDockerStackFiles(dir string) []string {
	return findFile(dir, []string{"*-stack.yml", "*-stack.yaml"})
}

func findFile(targetDir string, pattern []string) []string {
	result := []string{}

	for _, v := range pattern {
		matches, err := filepath.Glob(targetDir + v)

		if err != nil {
			fmt.Println(err)
		}

		if len(matches) != 0 {
			fmt.Println("Found : ", matches)
			result = append(result, matches...)
		}
	}
	return result
}

func gitPull(targetDir string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(targetDir)
	if err != nil {
		println(err)
	}
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		println(err)
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	println("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		println(err)
	}
}

func listContent(stringPath string) {
	err := filepath.Walk(stringPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != stringPath {
				println("      " + path)
			}
			return nil
		})
	if err2 != nil {
		println(err2)
	}
}

func helperFunc() {
	println("\nEx: use sammy...")
}
