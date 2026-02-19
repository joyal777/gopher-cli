package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main initializes and runs the GX-Shell interactive environment
func main() {
	displayWelcome()

	scanner := bufio.NewScanner(os.Stdin)

	for {
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

		handleCommand(command, parts)
	}
}

// displayWelcome shows the welcome message and available commands
func displayWelcome() {
	fmt.Println("--- Gopher Shell (GX) V3.5 Activated ---")
	fmt.Println("=== File Operations ===")
	fmt.Println("gx  [name]        : Create File/Folder")
	fmt.Println("gxd [name]        : Delete")
	fmt.Println("gxc [path]        : Change Directory (cd)")
	fmt.Println("gxl               : List Files (ls)")
	fmt.Println("gxs [name]        : Check Storage Size")
	fmt.Println("gxmv [src] [dst]  : Move/Rename file")
	fmt.Println("gxcp [src] [dst]  : Copy file")
	fmt.Println("gxfind [name]     : Find files by name")
	fmt.Println("gxecho [text] [file] : Write text to file")
	fmt.Println("gxdup [file]      : Duplicate file")
	fmt.Println("\n=== File Viewing ===")
	fmt.Println("gxcat [file]      : View file contents")
	fmt.Println("gxhead [file]     : View first 10 lines")
	fmt.Println("gxtail [file]     : View last 10 lines")
	fmt.Println("gxgrep [text] [file] : Search text in file")
	fmt.Println("gxstat [file]     : Show file statistics")
	fmt.Println("\n=== System Info ===")
	fmt.Println("gxpwd             : Print working directory")
	fmt.Println("gxdate            : Show current date/time")
	fmt.Println("gxinfo            : Show system info")
	fmt.Println("gxwhich [cmd]     : Show command path")
	fmt.Println("gxtree [dir]      : Show directory tree")
	fmt.Println("\n=== Utilities ===")
	fmt.Println("gxcount [dir]     : Count files in directory")
	fmt.Println("gxempty [name]    : Create empty file")
	fmt.Println("gxmkdir [name]    : Create directory (mkdir)")
	fmt.Println("gxtouch [file]    : Update file timestamp")
	fmt.Println("gxhelp            : Show extended help")
	fmt.Println("\nType 'exit' or press Ctrl+X then Enter to quit")
	fmt.Println("--------------------------------------")
}

// handleCommand routes the command to the appropriate handler with security validation
func handleCommand(command string, parts []string) {
	// Security validation
	if !ValidateCommandInput(command, parts) {
		fmt.Println("❌ Error: Invalid command input")
		return
	}

	switch command {
	// File Operations
	case "gx":
		if len(parts) < 2 {
			fmt.Println("Error: Missing name")
			return
		}
		createItem(parts[1])

	case "gxd":
		if len(parts) < 2 {
			fmt.Println("Error: Missing name")
			return
		}
		deleteItem(parts[1])

	case "gxc":
		if len(parts) < 2 {
			fmt.Println("Error: Missing path")
			return
		}
		changeDir(parts[1])

	case "gxl":
		listItems()

	case "gxs":
		if len(parts) < 2 {
			fmt.Println("Error: Missing name")
			return
		}
		showSize(parts[1])

	case "gxmv":
		if len(parts) < 3 {
			fmt.Println("Error: Missing source or destination")
			fmt.Println("Usage: gxmv [source] [destination]")
			return
		}
		moveFile(parts[1], parts[2])

	case "gxcp":
		if len(parts) < 3 {
			fmt.Println("Error: Missing source or destination")
			fmt.Println("Usage: gxcp [source] [destination]")
			return
		}
		copyFile(parts[1], parts[2])

	case "gxfind":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename to search for")
			return
		}
		findFiles(parts[1])

	case "gxecho":
		if len(parts) < 3 {
			fmt.Println("Error: Missing text or filename")
			fmt.Println("Usage: gxecho [text] [filename]")
			return
		}
		echoToFile(parts[1], parts[2])

	case "gxdup":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		duplicateFile(parts[1])

	// File Viewing
	case "gxcat":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		viewFile(parts[1])

	case "gxhead":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		headFile(parts[1], 10)

	case "gxtail":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		tailFile(parts[1], 10)

	case "gxgrep":
		if len(parts) < 3 {
			fmt.Println("Error: Missing search text or filename")
			fmt.Println("Usage: gxgrep [search text] [filename]")
			return
		}
		grepFile(parts[1], parts[2])

	case "gxstat":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		showFileStats(parts[1])

	// System Info
	case "gxpwd":
		printWorkingDir()

	case "gxdate":
		showDateTime()

	case "gxinfo":
		showSystemInfo()

	case "gxwhich":
		if len(parts) < 2 {
			fmt.Println("Error: Missing command name")
			return
		}
		whichCommand(parts[1])

	case "gxtree":
		path := "."
		if len(parts) >= 2 {
			path = parts[1]
		}
		showTree(path, "")

	// Utilities
	case "gxcount":
		path := "."
		if len(parts) >= 2 {
			path = parts[1]
		}
		countFiles(path)

	case "gxempty":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		createEmptyFile(parts[1])

	case "gxmkdir":
		if len(parts) < 2 {
			fmt.Println("Error: Missing directory name")
			return
		}
		createDirectory(parts[1])

	case "gxtouch":
		if len(parts) < 2 {
			fmt.Println("Error: Missing filename")
			return
		}
		touchFile(parts[1])

	// Help
	case "gxhelp":
		showExtendedHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
