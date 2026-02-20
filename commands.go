package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// ==================== FILE OPERATIONS ====================

// createItem creates a new file (if name contains ".") or directory
func createItem(name string) {
	// Security check
	if !validateFilename(name) {
		return
	}

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

// deleteItem removes a file or directory recursively
func deleteItem(name string) {
	// Security check
	if !validateFilename(name) {
		return
	}

	// Additional safety check - prevent deleting system files
	if isSuspiciousPath(name) {
		fmt.Println("❌ Error: Access denied - Cannot delete this path")
		return
	}

	err := os.RemoveAll(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("🗑️ '%s' deleted.\n", name)
}

// changeDir changes the current working directory
func changeDir(path string) {
	// Security check
	if !validatePath(path) {
		return
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory:", err)
	}
}

// moveFile moves or renames a file from source to destination
func moveFile(src, dst string) {
	// Security checks
	if !validatePath(src) || !validatePath(dst) {
		return
	}

	if !validateFilename(filepath.Base(dst)) {
		return
	}

	err := os.Rename(src, dst)
	if err != nil {
		fmt.Printf("Error moving '%s' to '%s': %v\n", src, dst, err)
		return
	}
	fmt.Printf("✅ Moved '%s' to '%s'\n", src, dst)
}

// copyFile copies a file from source to destination
func copyFile(src, dst string) {
	// Security checks
	if !validatePath(src) || !validatePath(dst) {
		return
	}

	if !validateFilename(filepath.Base(dst)) {
		return
	}

	// Check file size before copying
	info, err := os.Stat(src)
	if err != nil {
		fmt.Printf("Error accessing source file '%s': %v\n", src, err)
		return
	}

	if !checkFileSizeLimit(info.Size()) {
		return
	}

	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Printf("Error reading source file '%s': %v\n", src, err)
		return
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		fmt.Printf("Error writing to destination '%s': %v\n", dst, err)
		return
	}

	fmt.Printf("✅ Copied '%s' to '%s' (%d bytes)\n", src, dst, len(data))
}

// findFiles searches for files by name in the current directory
func findFiles(name string) {
	// Security check
	if !validateSearchTerm(name) {
		return
	}

	fmt.Printf("Searching for '%s' in current directory...\n", name)

	found := 0
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

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

// echoToFile appends text to a file
func echoToFile(text, filename string) {
	// Security checks
	if !validateFilename(filename) {
		return
	}

	if len(text) > 10000 {
		fmt.Println("❌ Error: Text too long (max 10000 chars)")
		return
	}

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

// duplicateFile creates a copy of a file with "_copy" suffix
func duplicateFile(filename string) {
	// Security check
	if !validateFilename(filename) {
		return
	}

	// Check file size before duplicating
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error accessing file '%s': %v\n", filename, err)
		return
	}

	if !checkFileSizeLimit(info.Size()) {
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	newFilename := base + "_copy" + ext

	if !validateFilename(newFilename) {
		return
	}

	err = os.WriteFile(newFilename, data, 0644)
	if err != nil {
		fmt.Printf("Error creating duplicate: %v\n", err)
		return
	}
	fmt.Printf("✅ File duplicated as '%s'\n", newFilename)
}

// ==================== FILE VIEWING ====================

// listItems displays all files and directories in the current directory
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

// viewFile displays the entire contents of a file
func viewFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	fmt.Printf("\n--- %s ---\n", filename)
	fmt.Println(string(data))
	if len(data) > 0 && data[len(data)-1] != '\n' {
		fmt.Println()
	}
	fmt.Printf("--- End of file (%d bytes) ---\n", len(data))
}

// headFile displays the first N lines of a file
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

// tailFile displays the last N lines of a file
func tailFile(filename string, lines int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

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

// grepFile searches for text within a file (case-insensitive)
func grepFile(searchText, filename string) {
	// Security checks
	if !validateSearchTerm(searchText) {
		return
	}

	if !validateFilename(filename) {
		return
	}

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

// ==================== SYSTEM INFORMATION ====================

// showSize calculates and displays the total size of a file or directory
func showSize(name string) {
	// Security check
	if !validatePath(name) {
		return
	}

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

// printWorkingDir displays the current working directory
func printWorkingDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}
	fmt.Printf("📂 Current directory: %s\n", dir)
}

// showDateTime displays the current date and time
func showDateTime() {
	now := time.Now()
	fmt.Printf("📅 Date: %s\n", now.Format("Monday, January 2, 2006"))
	fmt.Printf("⏰ Time: %s\n", now.Format("15:04:05 MST"))
	fmt.Printf("📆 Unix timestamp: %d\n", now.Unix())
}

// showSystemInfo displays system information (hostname, OS, CPU, etc.)
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

	tempDir := os.TempDir()
	fmt.Printf("📁 Temp Dir: %s\n", tempDir)

	if _, err := os.Stat(".git"); err == nil {
		fmt.Println("🔀 Git repo: Yes")
	}
}

// whichCommand locates a command in the system PATH
func whichCommand(cmd string) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("❌ Command '%s' not found in PATH\n", cmd)
		return
	}
	fmt.Printf("✅ '%s' found at: %s\n", cmd, path)
}

// showTree displays a tree structure of directories and files
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

// ==================== UTILITIES ====================

// countFiles counts the number of files and directories in a path
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

// createEmptyFile creates an empty file
func createEmptyFile(name string) {
	// Security check
	if !validateFilename(name) {
		return
	}

	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", name, err)
		return
	}
	file.Close()
	fmt.Printf("📄 Empty file '%s' created (0 bytes)\n", name)
}

// createDirectory creates a new directory
func createDirectory(name string) {
	// Security check
	if !validateFilename(name) {
		return
	}

	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Printf("Error creating directory '%s': %v\n", name, err)
		return
	}
	fmt.Printf("📁 Directory '%s' created\n", name)
}

// showFileStats displays detailed statistics about a file
func showFileStats(filename string) {
	// Security check
	if !validateFilename(filename) {
		return
	}

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

// touchFile creates or updates the timestamp of a file
func touchFile(filename string) {
	// Security check
	if !validateFilename(filename) {
		return
	}

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

	now := time.Now()
	err := os.Chtimes(filename, now, now)
	if err != nil {
		fmt.Printf("Error touching file '%s': %v\n", filename, err)
		return
	}
	fmt.Printf("✅ File '%s' timestamp updated\n", filename)
}

// gxmd5 computes and prints the MD5 checksum of a file
func gxmd5(filename string) {
	if !validateFilename(filename) {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	sum := hasher.Sum(nil)
	fmt.Printf("MD5(%s) = %x\n", filename, sum)
}

// gxlines counts and prints the number of lines in a file
func gxlines(filename string) {
	if !validateFilename(filename) {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	fmt.Printf("%s: %d lines\n", filename, lines)
}

// gxreplace replaces all occurrences of old with new in the provided file
func gxreplace(old, new, filename string) {
	if !validateFilename(filename) {
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", filename, err)
		return
	}

	content := strings.ReplaceAll(string(data), old, new)

	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file '%s': %v\n", filename, err)
		return
	}

	fmt.Printf("✅ Replaced '%s' with '%s' in '%s'\n", old, new, filename)
}

// ==================== HELP ====================

// showExtendedHelp displays the extended help menu
func showExtendedHelp() {
	fmt.Print(`
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
