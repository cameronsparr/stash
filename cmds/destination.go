package cmds

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cameronsparr/stash/config"
)

// Destination specifies the 'stash destination' command.
var Destination = &Command{
	Usage: "destination [add | list | remove] [arguments]",
	Short: "Add, list, or remove backup destinations.",
	Long: `
Add, list, or remove backup destinations.
`,
	Run: runDestination,
}

func runDestination(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.UsageExit()
	}
	switch args[0] {
	case "add":
		runAdd(args)
	case "list":
		runList(args)
	case "remove":
		runRemove(args)
	}
}

func runAdd(args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which type of backup destination would you like to add?")
	fmt.Println("	1. Amazon (S3 or Glacier)")
	fmt.Println("	2. Google Cloud")
	fmt.Println("")
	fmt.Print("Choose an option [1-2]: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	switch text {
	case "1":
		addAmazon()
	case "2":
		return
	}
}

func addAmazon() {
	// TODO: implement
	confMngr := config.NewMngr()
	reader := bufio.NewReader(os.Stdin)
	confEntry := config.Entry{
		Name:        getName(reader, confMngr),
		Folders:     getFolders(reader),
		Type:        "Amazon",
		Credentials: getCredentials(reader),
	}
	confMngr.AddDestination(confEntry)
}

func getName(reader *bufio.Reader, confMngr *config.Mngr) string {
	fmt.Print("Specify a name for this destination: ")
	text, _ := reader.ReadString('\n')
	name := strings.TrimSpace(text)
	if confMngr.IsDuplicateEntry(config.Entry{Name: name}) {
		log.Fatalf("Attempted to add duplicate entry [%s], if you were "+
			"trying to add folders to an existing backup destination, use "+
			"'stash folder add'",
			name)
	}
	return name
}

func getFolders(reader *bufio.Reader) []string {
	fmt.Print("Specify directories to backup (space-separated): ")
	text, _ := reader.ReadString('\n')
	dirs := strings.Split(strings.TrimSpace(text), " ")
	// Check that folders are valid and accessible
	for _, dir := range dirs {
		if _, err := isValidDirectory(dir); err != nil {
			log.Fatalln(err)
		}
	}
	return dirs
}

func isValidDirectory(dir string) (bool, error) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			e := fmt.Sprintf("Directory %s does not exist", dir)
			return false, errors.New(e)
		} else {
			e := fmt.Sprintf("Error accessing %s, do you have read permission?", dir)
			return false, errors.New(e)
		}
	}
	return true, nil
}

func getCredentials(reader *bufio.Reader) map[string]string {
	return map[string]string{"key": "supersecret"}
}

func runList(args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("This is not implemented yet, but do you love marutaro? [Y/y]")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func runRemove(args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("This is not implemented yet, but do you love marutaro? [Y/y]")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
