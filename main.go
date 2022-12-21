package main

import (
	"fmt"
	"github.com/BtKent/swiftbarssh/ssh"
)

func main() {

	// This is the display name in the menu bar
	// This :xserve: is an icon from SF Symbole.
	// You can use the text you want.
	displayedText := ":xserve:"

	// This is the path to the config file
	configFile := "/Users/quentin/.ssh/config"

	fmt.Println(displayedText)
	fmt.Println("---")
	ssh.PrintHosts(configFile)
}
