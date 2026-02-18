# GX-Shell 🐹 (Version 3.0)

A powerful, custom interactive shell built in Go. GX-Shell transforms your standard terminal into a specialized environment for rapid file system management with custom "GX" prefixed commands. Now with 15+ commands for complete file system control!

## ✨ Features

- **Interactive Environment**: Real-time path display in your prompt.
- **Complete File Operations**: Create, delete, move, copy, and find files/folders.
- **File Viewing**: View entire files or just the beginning/end.
- **System Information**: Check dates, paths, and system details without leaving the shell.
- **Storage Analytics**: Calculate the total size of files and directories (auto-converts to KB/MB/GB).
- **Directory Navigation**: Full support for moving through your file system.
- **Utility Tools**: Count files, create empty files, and more.

## 🚀 Command Reference

### File Operations

| Command | Action | Example |
| :--- | :--- | :--- |
| `gx` | **Create** a File or Folder | `gx notes.txt` or `gx source_code` |
| `gxd` | **Delete** (Recursive) | `gxd old_folder` |
| `gxmv` | **Move/Rename** file or folder | `gxmv old.txt new.txt` |
| `gxcp` | **Copy** file | `gxcp source.txt backup.txt` |
| `gxfind` | **Find** files by name | `gxfind .go` |
| `gxempty` | **Create** empty file | `gxempty temp.txt` |
| `gxmkdir` | **Create** directory | `gxmkdir newfolder` |

### Navigation & Listing

| Command | Action | Example |
| :--- | :--- | :--- |
| `gxc` | **Change Directory** (cd) | `gxc ..` or `gxc projects` |
| `gxl` | **List** items in directory | `gxl` |
| `gxpwd` | **Print** working directory | `gxpwd` |
| `gxcount` | **Count** files in directory | `gxcount` or `gxcount ./folder` |

### File Viewing

| Command | Action | Example |
| :--- | :--- | :--- |
| `gxcat` | **View** entire file contents | `gxcat notes.txt` |
| `gxhead` | **View** first 10 lines of file | `gxhead log.txt` |
| `gxtail` | **View** last 10 lines of file | `gxtail log.txt` |

### System Information

| Command | Action | Example |
| :--- | :--- | :--- |
| `gxs` | **Check Storage Size** | `gxs my_videos` |
| `gxdate` | **Show** current date and time | `gxdate` |
| `gxinfo` | **Show** system information | `gxinfo` |
| `gxwhich` | **Locate** a command in PATH | `gxwhich go` |

### Shell Control

| Command | Action | Example |
| :--- | :--- | :--- |
| `exit` | **Quit** GX-Shell | `exit` or `Ctrl+X` |

## 🛠️ Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) 1.21+ installed on your machine.

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/joyal777/gopher-cli.git
   cd gopher-cli

2. **Run the shell directly:**
    ```bash
    go run main.go

3.  **(Optional) Build the executable:**
    ```bash
    go build -o gx.exe
    ./gx.exe

**📖 Usage Examples**

1.  **File Management**
    ```bash
    gx-shell> gxl                    # List files
    gx-shell> gx notes.txt            # Create a file
    gx-shell> gx projects             # Create a folder
    gx-shell> gxmv notes.txt old.txt  # Rename file
    gx-shell> gxcp old.txt backup.txt # Copy file
    gx-shell> gxd old.txt             # Delete file

2.  **Navigation**
    ```bash
    gx-shell> gxpwd                   # Show current directory
    gx-shell> gxc projects            # Enter projects folder
    gx-shell> gxc ..                   # Go back

3.  **Viewing Files**
    ```bash
    gx-shell> gxhead log.txt          # See first 10 lines
    gx-shell> gxtail log.txt          # See last 10 lines
    gx-shell> gxcat config.json       # View entire file

4.  **System Info**
    ```bash
    gx-shell> gxdate                   # Check current time
    gx-shell> gxinfo                   # Show system details
    gx-shell> gxs downloads             # Check folder size
    gx-shell> gxwhich python            # Find python location

5.  **Finding & Counting**
    ```bash
    gx-shell> gxfind .md               # Find all markdown files
    gx-shell> gxcount                   # Count items in current dir
    gx-shell> gxcount ./src             # Count items in src folder

**🤝 Contributing**

This is an open-source project! Feel free to:

Fork the repository

Add new commands or features

Improve existing functionality

Submit a pull request

Whether it's bug fixes, new commands, or documentation improvements, all contributions are welcome!

**📝 License**
MIT License - feel free to use this project for learning or building your own tools.

Built with ❤️ using Golang | Version 3.0

