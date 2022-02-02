package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func main() {

	command := os.Args

	if len(command) > 1 {
		switch command[1] {
		case "help":
			helperFunc()
		case "h":
			helperFunc()
		case "deploy":
			if len(command) > 2 {
				fmt.Println("[-] sammy deploy command... ")
				deploy(command[2])
			} else {
				throwError("No Directory provided !", 1)
				helperFunc()
			}
		case "restart-nginx":
			fmt.Println("[-] sammy restart-nginx command...")
		case "stop":
			fmt.Println("[-] sammy stop command...")
		case "update-field":
			fmt.Println("[-] sammy update-field command...")
		default:
			throwError("No valid command provided...", 0)
			helperFunc()
		}
	} else {
		throwError("No command provided", 0)
		helperFunc()
	}
}

func throwError(err string, statusCode int) {
	fmt.Printf("[x] Error: %s", err)
	fmt.Println()

	if statusCode == 1 {
		log.Fatal(err)
	} else if statusCode == 0 {
		log.Println(err)
	}
	os.Exit(statusCode)
}

func execCommand(program string, args string) {
	cmd := exec.Command(program, args)

	err := cmd.Run()

	if err != nil {
		log.Println(err)
	}
}

func deploy(dir string) {
	fmt.Printf("Deploying %s\n", dir)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			throwError("Directory '"+dir+"' not found : "+err.Error(), 1)
		} else {
			throwError(dir+" : "+err.Error(), 1)
		}
	}

	listDeployFolders(dir)

	gitPull(dir)

	// We run the docker-compose here
	for _, p := range getDockerComposeFiles(dir) {
		execCommand("docker-compose", "-f "+p+" up -d")
	}

	// we run the docker stack deploy here
	for _, p := range getDockerStackFiles(dir) {
		execCommand("docker stack deploy", "-c "+p)
	}
}

func listDeployFolders(dir string) {
	servicesDir := dir + "/services"
	confDir := dir + "/conf"

	fmt.Printf("Config folder %s\n", confDir)
	listContent(confDir)

	fmt.Printf("Services folder %s\n", servicesDir)
	listContent(servicesDir)
}

func getDockerComposeFiles(dir string) []string {
	return findFiles(dir, []string{"*-compose.yml", "*-compose.yaml"})
}

func getDockerStackFiles(dir string) []string {
	return findFiles(dir, []string{"*-stack.yml", "*-stack.yaml"})
}

func findFiles(targetDir string, pattern []string) []string {
	result := []string{}

	for _, v := range pattern {
		matches, err := filepath.Glob(targetDir + v)
		if err != nil {
			throwError("failed to find files: "+err.Error(), 1)
		}

		if len(matches) != 0 {
			fmt.Println("Found: ", matches)
			result = append(result, matches...)
		}
	}
	return result
}

func gitPull(targetDir string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(targetDir)
	if err != nil {
		fmt.Println(err)
	}
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		fmt.Println(err)
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	fmt.Println("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Println(err)
	}
}

func listContent(stringPath string) {
	err := filepath.Walk(stringPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != stringPath {
				fmt.Println("      " + path)
			}
			return nil
		})
	if err != nil {
		throwError("failed to list content: "+err.Error(), 1)
	}
}

func helperFunc() {
	fmt.Println("\nEx: use sammy...")
}
