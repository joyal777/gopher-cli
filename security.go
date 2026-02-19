package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Security constants
const (
	MAX_PATH_LENGTH     = 260               // Windows MAX_PATH
	MAX_FILE_SIZE       = 512 * 1024 * 1024 // 512MB limit
	MAX_FILENAME_LENGTH = 255
	ALLOWED_NAME_CHARS  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789._-@+= ()"
)

// isPathTraversal checks if a path tries to escape the current working directory
func isPathTraversal(path string) bool {
	// Reject absolute paths
	if filepath.IsAbs(path) {
		return true
	}

	// Normalize the path
	cleanPath := filepath.Clean(path)

	// Check for directory traversal patterns
	if strings.Contains(cleanPath, "..") {
		return true
	}

	// Check for suspicious patterns
	if strings.Contains(path, "\x00") { // Null byte injection
		return true
	}

	if strings.Contains(path, ";") { // Command separator
		return true
	}

	return false
}

// validatePath checks if a path is safe to use
func validatePath(path string) bool {
	// Check for empty path
	if len(strings.TrimSpace(path)) == 0 {
		return false
	}

	// Check length
	if len(path) > MAX_PATH_LENGTH {
		fmt.Println("❌ Error: Path exceeds maximum length")
		return false
	}

	// Check for path traversal
	if isPathTraversal(path) {
		fmt.Println("❌ Error: Access denied - Invalid path")
		return false
	}

	return true
}

// validateFilename checks if a filename is safe
func validateFilename(filename string) bool {
	if len(strings.TrimSpace(filename)) == 0 {
		fmt.Println("❌ Error: Filename cannot be empty")
		return false
	}

	if len(filename) > MAX_FILENAME_LENGTH {
		fmt.Println("❌ Error: Filename exceeds maximum length (255 chars)")
		return false
	}

	// Check for path traversal attempts
	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		fmt.Println("❌ Error: Filename cannot contain path separators")
		return false
	}

	// Reject suspicious characters
	if strings.Contains(filename, "..") {
		fmt.Println("❌ Error: Invalid filename")
		return false
	}

	if strings.Contains(filename, "\x00") {
		fmt.Println("❌ Error: Invalid filename")
		return false
	}

	// Check for command injection characters
	if strings.ContainsAny(filename, ";|&$()><`\\") {
		fmt.Println("❌ Error: Filename contains invalid characters")
		return false
	}

	return true
}

// validateSearchTerm checks if a search term is safe
func validateSearchTerm(term string) bool {
	if len(strings.TrimSpace(term)) == 0 {
		fmt.Println("❌ Error: Search term cannot be empty")
		return false
	}

	if len(term) > 1000 {
		fmt.Println("❌ Error: Search term too long")
		return false
	}

	// Check for regex injection or command injection patterns
	if strings.Contains(term, "\x00") {
		fmt.Println("❌ Error: Invalid search term")
		return false
	}

	return true
}

// validateInputArgs checks if input arguments are safe
func validateInputArgs(args []string) bool {
	for _, arg := range args {
		if len(arg) > MAX_PATH_LENGTH {
			fmt.Println("❌ Error: Argument too long")
			return false
		}

		if strings.Contains(arg, "\x00") {
			fmt.Println("❌ Error: Invalid argument")
			return false
		}
	}
	return true
}

// isSymlink checks if a path is a symbolic link
func isSymlink(path string) (bool, error) {
	info, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false, err
	}

	// If EvalSymlinks changes the path, it was a symlink
	absPath, _ := filepath.Abs(path)
	absInfo, _ := filepath.Abs(info)

	return absPath != absInfo, nil
}

// sanitizePath cleans and validates a path
func sanitizePath(path string) (string, error) {
	if !validatePath(path) {
		return "", fmt.Errorf("invalid path: %s", path)
	}

	cleanPath := filepath.Clean(path)

	if isPathTraversal(cleanPath) {
		return "", fmt.Errorf("path traversal detected: %s", path)
	}

	return cleanPath, nil
}

// sanitizeFilename cleans and validates a filename
func sanitizeFilename(filename string) (string, error) {
	if !validateFilename(filename) {
		return "", fmt.Errorf("invalid filename: %s", filename)
	}

	cleanName := strings.TrimSpace(filename)
	cleanName = strings.ReplaceAll(cleanName, "\n", "")
	cleanName = strings.ReplaceAll(cleanName, "\r", "")

	return cleanName, nil
}

// checkFileSizeLimit validates that a file doesn't exceed size limit
func checkFileSizeLimit(size int64) bool {
	if size > MAX_FILE_SIZE {
		fmt.Printf("❌ Error: File size exceeds maximum allowed (%d MB)\n", MAX_FILE_SIZE/(1024*1024))
		return false
	}
	return true
}

// ValidateCommandInput validates command input for safety
func ValidateCommandInput(command string, args []string) bool {
	// Check command length
	if len(command) > 50 {
		fmt.Println("❌ Error: Command name too long")
		return false
	}

	// Validate arguments count
	if len(args) > 10 {
		fmt.Println("❌ Error: Too many arguments")
		return false
	}

	// Check for null bytes in command
	if strings.Contains(command, "\x00") {
		fmt.Println("❌ Error: Invalid command")
		return false
	}

	// Validate all arguments
	if !validateInputArgs(args) {
		return false
	}

	return true
}

// RateLimitCheck prevents abuse through rapid repeated operations
var operationCount = 0
var lastOperationTime int64 = 0

// CheckRateLimit implements basic rate limiting
func CheckRateLimit() bool {
	// Placeholder for rate limiting logic
	// Can be enhanced with actual time-based rate limiting
	return true
}

// isSuspiciousPath checks if a path looks suspicious or dangerous
func isSuspiciousPath(path string) bool {
	// List of dangerous patterns
	dangerousPaths := []string{
		"system", "system32", "windows", "winnt",
		"boot.ini", "autoexec.bat", "config.sys",
		"/etc/", "/bin/", "/sbin/", "/usr/", "/var/",
		".bashrc", ".bash_profile", ".ssh",
	}

	lowerPath := strings.ToLower(path)

	for _, dangerous := range dangerousPaths {
		if strings.Contains(lowerPath, dangerous) {
			return true
		}
	}

	return false
}
