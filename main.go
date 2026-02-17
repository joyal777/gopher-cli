package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("--- Gopher Shell (GX) V2 Activated ---")
	fmt.Println("gx  [name]  : Create File/Folder")
	fmt.Println("gxd [name]  : Delete")
	fmt.Println("gxc [path]  : Change Directory (cd)")
	fmt.Println("gxl         : List Files (ls)")
	fmt.Println("gxs [name]  : Check Storage Size")
	fmt.Println("Type 'exit' or press Ctrl+X then Enter to quit")
	fmt.Println("--------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Show current working directory in the prompt
		cwd, _ := os.Getwd()
		fmt.Printf("\n%s\ngx-shell> ", cwd)

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		if command == "\x18" || command == "exit" {
			fmt.Println("Exiting Gopher Shell. Bye!")
			break
		}

		switch command {
		case "gx":
			if len(parts) < 2 {
				fmt.Println("Error: Missing name")
				continue
			}
			createItem(parts[1])

		case "gxd":
			if len(parts) < 2 {
				fmt.Println("Error: Missing name")
				continue
			}
			deleteItem(parts[1])

		case "gxc": // Change Directory
			if len(parts) < 2 {
				fmt.Println("Error: Missing path")
				continue
			}
			changeDir(parts[1])

		case "gxl": // List Files
			listItems()

		case "gxs": // Storage Size
			if len(parts) < 2 {
				fmt.Println("Error: Missing name")
				continue
			}
			showSize(parts[1])

		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

// gxc: Change Directory
func changeDir(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
	}
}

// gxl: List Items
func listItems() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Mode        Size         Name")
	fmt.Println("----        ----         ----")
	for _, file := range files {
		info, _ := file.Info()
		indicator := "📄"
		if file.IsDir() {
			indicator = "📁"
		}
		fmt.Printf("%-10s  %-10d   %s %s\n", info.Mode(), info.Size(), indicator, file.Name())
	}
}

// gxs: Show Storage Size
func showSize(name string) {
	var totalSize int64

	err := filepath.Walk(name, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error calculating size:", err)
		return
	}

	// Convert bytes to KB/MB for readability
	const unit = 1024
	if totalSize < unit {
		fmt.Printf("Size of '%s': %d B\n", name, totalSize)
	} else if totalSize < unit*unit {
		fmt.Printf("Size of '%s': %.2f KB\n", name, float64(totalSize)/float64(unit))
	} else {
		fmt.Printf("Size of '%s': %.2f MB\n", name, float64(totalSize)/float64(unit*unit))
	}
}

// gx: Create
func createItem(name string) {
	if strings.Contains(name, ".") {
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		file.Close()
		fmt.Printf("📄 File '%s' created.\n", name)
	} else {
		err := os.Mkdir(name, 0755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("📁 Folder '%s' created.\n", name)
	}
}

// gxd: Delete
func deleteItem(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("🗑️ '%s' deleted.\n", name)
}
