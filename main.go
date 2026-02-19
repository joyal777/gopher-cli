package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
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

		switch command {
		// Original commands
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

		case "gxc":
			if len(parts) < 2 {
				fmt.Println("Error: Missing path")
				continue
			}
			changeDir(parts[1])

		case "gxl":
			listItems()

		case "gxs":
			if len(parts) < 2 {
				fmt.Println("Error: Missing name")
				continue
			}
			showSize(parts[1])

		// NEW: File operations
		case "gxmv":
			if len(parts) < 3 {
				fmt.Println("Error: Missing source or destination")
				fmt.Println("Usage: gxmv [source] [destination]")
				continue
			}
			moveFile(parts[1], parts[2])

		case "gxcp":
			if len(parts) < 3 {
				fmt.Println("Error: Missing source or destination")
				fmt.Println("Usage: gxcp [source] [destination]")
				continue
			}
			copyFile(parts[1], parts[2])

		case "gxfind":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename to search for")
				continue
			}
			findFiles(parts[1])

		// NEW: File viewing
		case "gxcat":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			viewFile(parts[1])

		case "gxhead":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			headFile(parts[1], 10)

		case "gxtail":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			tailFile(parts[1], 10)

		// NEW: System info
		case "gxpwd":
			printWorkingDir()

		case "gxdate":
			showDateTime()

		case "gxinfo":
			showSystemInfo()

		case "gxwhich":
			if len(parts) < 2 {
				fmt.Println("Error: Missing command name")
				continue
			}
			whichCommand(parts[1])

		// NEW: Utilities
		case "gxcount":
			path := "."
			if len(parts) >= 2 {
				path = parts[1]
			}
			countFiles(path)

		case "gxempty":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			createEmptyFile(parts[1])

		case "gxmkdir":
			if len(parts) < 2 {
				fmt.Println("Error: Missing directory name")
				continue
			}
			createDirectory(parts[1])

		case "gxecho":
			if len(parts) < 3 {
				fmt.Println("Error: Missing text or filename")
				fmt.Println("Usage: gxecho [text] [filename]")
				continue
			}
			echoToFile(parts[1], parts[2])

		case "gxdup":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			duplicateFile(parts[1])

		case "gxgrep":
			if len(parts) < 3 {
				fmt.Println("Error: Missing search text or filename")
				fmt.Println("Usage: gxgrep [search text] [filename]")
				continue
			}
			grepFile(parts[1], parts[2])

		case "gxstat":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			showFileStats(parts[1])

		case "gxtouch":
			if len(parts) < 2 {
				fmt.Println("Error: Missing filename")
				continue
			}
			touchFile(parts[1])

		case "gxtree":
			path := "."
			if len(parts) >= 2 {
				path = parts[1]
			}
			showTree(path, "")

		case "gxhelp":
			showExtendedHelp()

		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

// ==================== ORIGINAL FUNCTIONS ====================

func changeDir(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
	}
}

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

	const unit = 1024
	if totalSize < unit {
		fmt.Printf("Size of '%s': %d B\n", name, totalSize)
	} else if totalSize < unit*unit {
		fmt.Printf("Size of '%s': %.2f KB\n", name, float64(totalSize)/float64(unit))
	} else {
		fmt.Printf("Size of '%s': %.2f MB\n", name, float64(totalSize)/float64(unit*unit))
	}
}

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

func deleteItem(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("🗑️ '%s' deleted.\n", name)
}

// ==================== NEW FUNCTIONS ====================

// File Operations

func moveFile(src, dst string) {
	err := os.Rename(src, dst)
	if err != nil {
		fmt.Printf("Error moving '%s' to '%s': %v\n", src, dst, err)
		return
	}
	fmt.Printf("✅ Moved '%s' to '%s'\n", src, dst)
}

func copyFile(src, dst string) {
	// Read source file
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Printf("Error reading source file '%s': %v\n", src, err)
		return
	}

	// Write to destination
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Printf("Error writing to destination '%s': %v\n", dst, err)
		return
	}

	fmt.Printf("✅ Copied '%s' to '%s' (%d bytes)\n", src, dst, len(data))
}

func findFiles(name string) {
	fmt.Printf("Searching for '%s' in current directory...\n", name)

	found := 0
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip files we can't access
		}

		// Check if filename contains the search term
		if strings.Contains(info.Name(), name) {
			fmt.Printf("  📍 %s\n", path)
			found++
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error during search: %v\n", err)
		return
	}

	if found == 0 {
		fmt.Printf("No files found matching '%s'\n", name)
	} else {
		fmt.Printf("Found %d matching file(s)\n", found)
	}
}

// File Viewing

func viewFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	fmt.Printf("\n--- %s ---\n", filename)
	fmt.Println(string(data))
	if len(data) > 0 && data[len(data)-1] != '\n' {
		fmt.Println() // Add newline if file doesn't end with one
	}
	fmt.Printf("--- End of file (%d bytes) ---\n", len(data))
}

func headFile(filename string, lines int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	fmt.Printf("\n--- First %d lines of %s ---\n", lines, filename)
	for scanner.Scan() && count < lines {
		fmt.Println(scanner.Text())
		count++
	}

	if count == 0 {
		fmt.Println("(file is empty)")
	} else if count < lines {
		fmt.Printf("--- End of file (only %d lines) ---\n", count)
	} else {
		fmt.Printf("--- End of head (showed %d lines) ---\n", lines)
	}
}

func tailFile(filename string, lines int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	// Read all lines
	var allLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("\n--- Last %d lines of %s ---\n", lines, filename)

	// Calculate start index
	start := 0
	if len(allLines) > lines {
		start = len(allLines) - lines
	}

	for i := start; i < len(allLines); i++ {
		fmt.Println(allLines[i])
	}

	if len(allLines) == 0 {
		fmt.Println("(file is empty)")
	} else {
		fmt.Printf("--- End of tail (showed %d of %d lines) ---\n",
			len(allLines)-start, len(allLines))
	}
}

// System Info

func printWorkingDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}
	fmt.Printf("📂 Current directory: %s\n", dir)
}

func showDateTime() {
	now := time.Now()
	fmt.Printf("📅 Date: %s\n", now.Format("Monday, January 2, 2006"))
	fmt.Printf("⏰ Time: %s\n", now.Format("15:04:05 MST"))
	fmt.Printf("📆 Unix timestamp: %d\n", now.Unix())
}

func showSystemInfo() {
	hostname, _ := os.Hostname()
	cwd, _ := os.Getwd()

	fmt.Println("=== System Information ===")
	fmt.Printf("💻 Hostname: %s\n", hostname)
	fmt.Printf("📂 Current Dir: %s\n", cwd)
	fmt.Printf("🔧 OS: %s\n", runtime.GOOS)
	fmt.Printf("🖥️  Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("⚙️  Go Version: %s\n", runtime.Version())
	fmt.Printf("🧵 CPUs: %d\n", runtime.NumCPU())

	// Get temp directory
	tempDir := os.TempDir()
	fmt.Printf("📁 Temp Dir: %s\n", tempDir)

	// Check if we're in a git repo (simple check)
	if _, err := os.Stat(".git"); err == nil {
		fmt.Println("🔀 Git repo: Yes")
	}
}

func whichCommand(cmd string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("❌ Command '%s' not found in PATH\n", cmd)
		return
	}
	fmt.Printf("✅ '%s' found at: %s\n", cmd, path)
}

// Utilities

func countFiles(path string) {
	fileCount := 0
	dirCount := 0

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory '%s': %v\n", path, err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
	}

	fmt.Printf("📊 Directory '%s' contains:\n", path)
	fmt.Printf("  📁 %d directories\n", dirCount)
	fmt.Printf("  📄 %d files\n", fileCount)
	fmt.Printf("  📦 Total: %d items\n", dirCount+fileCount)
}

func createEmptyFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", name, err)
		return
	}
	file.Close()
	fmt.Printf("📄 Empty file '%s' created (0 bytes)\n", name)
}

func createDirectory(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Printf("Error creating directory '%s': %v\n", name, err)
		return
	}
	fmt.Printf("📁 Directory '%s' created\n", name)
}

// ==================== NEW FUNCTIONS (V3.5) ====================

func echoToFile(text, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Printf("Error writing to file '%s': %v\n", filename, err)
		return
	}
	fmt.Printf("✅ Text written to '%s'\n", filename)
}

func duplicateFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	// Generate new filename with _copy suffix
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	newFilename := base + "_copy" + ext

	err = os.WriteFile(newFilename, data, 0644)
	if err != nil {
		fmt.Printf("Error creating duplicate: %v\n", err)
		return
	}
	fmt.Printf("✅ File duplicated as '%s'\n", newFilename)
}

func grepFile(searchText, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	found := 0

	fmt.Printf("\n--- Searching for '%s' in %s ---\n", searchText, filename)
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(searchText)) {
			fmt.Printf("  Line %d: %s\n", lineNum, line)
			found++
		}
	}

	if found == 0 {
		fmt.Printf("No matches found for '%s'\n", searchText)
	} else {
		fmt.Printf("--- Found %d match(es) ---\n", found)
	}
}

func showFileStats(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error accessing file '%s': %v\n", filename, err)
		return
	}

	fmt.Printf("\n=== File Statistics: %s ===\n", filename)
	fmt.Printf("📄 Name: %s\n", info.Name())
	fmt.Printf("📊 Size: %d bytes\n", info.Size())
	fmt.Printf("🔒 Mode: %v\n", info.Mode())
	fmt.Printf("⏰ Modified: %v\n", info.ModTime())
	fmt.Printf("📁 Is Dir: %v\n", info.IsDir())

	// Calculate human-readable size
	size := float64(info.Size())
	if size < 1024 {
		fmt.Printf("💾 Size (readable): %.2f B\n", size)
	} else if size < 1024*1024 {
		fmt.Printf("💾 Size (readable): %.2f KB\n", size/1024)
	} else if size < 1024*1024*1024 {
		fmt.Printf("💾 Size (readable): %.2f MB\n", size/(1024*1024))
	} else {
		fmt.Printf("💾 Size (readable): %.2f GB\n", size/(1024*1024*1024))
	}
}

func touchFile(filename string) {
	// Create file if it doesn't exist
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file '%s': %v\n", filename, err)
			return
		}
		file.Close()
		fmt.Printf("✅ File '%s' created (touched)\n", filename)
		return
	}

	// Update timestamp
	now := time.Now()
	err := os.Chtimes(filename, now, now)
	if err != nil {
		fmt.Printf("Error touching file '%s': %v\n", filename, err)
		return
	}
	fmt.Printf("✅ File '%s' timestamp updated\n", filename)
}

func showTree(dirPath string, prefix string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		isLast := i == len(entries)-1
		currentPrefix := "├── "
		nextPrefix := prefix + "│   "

		if isLast {
			currentPrefix = "└── "
			nextPrefix = prefix + "    "
		}

		icon := "📄"
		if entry.IsDir() {
			icon = "📁"
		}

		fmt.Printf("%s%s%s %s\n", prefix, currentPrefix, icon, entry.Name())

		if entry.IsDir() {
			fullPath := filepath.Join(dirPath, entry.Name())
			showTree(fullPath, nextPrefix)
		}
	}

	return nil
}

func showExtendedHelp() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════════════╗
║              GX-Shell Extended Help (Version 3.5)               ║
╚══════════════════════════════════════════════════════════════════╝

📁 FILE OPERATIONS:
  gx [name]         - Create file (with .) or folder without extension
  gxd [name]        - Delete file or folder recursively
  gxc [path]        - Change directory
  gxl               - List files in current directory
  gxs [name]        - Show total size of file/folder
  gxmv [src] [dst]  - Move or rename a file/folder
  gxcp [src] [dst]  - Copy a file
  gxfind [name]     - Search for files containing name
  gxecho [text] [file] - Append text to file
  gxdup [file]      - Create a duplicate copy of file

📖 FILE VIEWING:
  gxcat [file]      - Display entire file contents
  gxhead [file]     - Show first 10 lines
  gxtail [file]     - Show last 10 lines
  gxgrep [text] [file] - Find lines containing text
  gxstat [file]     - Show detailed file statistics

🖥️  SYSTEM INFO:
  gxpwd             - Print current working directory
  gxdate            - Show current date and time
  gxinfo            - Display system information
  gxwhich [cmd]     - Find command location in PATH
  gxtree [dir]      - Display directory tree structure

🛠️  UTILITIES:
  gxcount [dir]     - Count files in directory
  gxempty [file]    - Create empty file
  gxmkdir [dir]     - Create directory
  gxtouch [file]    - Create/update file timestamp
  gxhelp            - Show this help message

⏹️  CONTROL:
  exit or Ctrl+X    - Exit the shell

`)
}
