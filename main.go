package main

import (
	"os"
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
		println("[-] sammy deploy... ")
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

	println("[>] DEPLOY_DIR : " + dir)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			println("[x] ERR: " + dir + " doesn't exist.")
			println(err)
			return
		} else {
			println("[x] ERR: Unexpected error occured for the dir: " + dir)
			println(err)
			return
		}
	}
	servicesDir := dir + "/services"
	confDir := dir + "/conf"

	println("[>] CONF_DIR " + confDir)
	listContent(confDir)

	println("[>] SERVICE_DIR " + servicesDir)
	listContent(servicesDir)

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(dir)
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
	err2 := filepath.Walk(stringPath,
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
