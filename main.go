package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- Gopher Shell (GX) Activated ---")
	fmt.Println("Commands: gx [name] (Create), gxd [name] (Delete)")
	fmt.Println("Press Ctrl+X and Enter to Exit")
	fmt.Println("-----------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("gx-shell> ") // The separate command prompt
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		// Handle Exit Logic (Ctrl+X is \x18 in ASCII)
		if command == "\x18" || command == "exit" {
			fmt.Println("Exiting Gopher Shell. Bye!")
			break
		}

		switch command {
		case "gx":
			if len(parts) < 2 {
				fmt.Println("Error: Please provide a name (e.g., gx foldername)")
				continue
			}
			target := parts[1]
			createItem(target)

		case "gxd":
			if len(parts) < 2 {
				fmt.Println("Error: Please provide a name to delete")
				continue
			}
			target := parts[1]
			deleteItem(target)

		default:
			fmt.Printf("Unknown command: %s. Use gx or gxd.\n", command)
		}
	}
}

// Logic for creating files or folders
func createItem(name string) {
	if strings.Contains(name, ".") {
		// It's a file
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Printf("File '%s' created successfully.\n", name)
	} else {
		// It's a folder
		err := os.Mkdir(name, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Printf("Folder '%s' created successfully.\n", name)
	}
}

// Logic for deleting
func deleteItem(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Error deleting:", err)
		return
	}
	fmt.Printf("'%s' deleted.\n", name)
}
