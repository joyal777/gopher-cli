# GX-Shell 🐹

A lightweight, interactive command-line interface (CLI) built in Go for fast file and directory management. Instead of complex terminal commands, GX-Shell provides a dedicated environment for quick file operations.

## ✨ Features
- **Interactive Mode**: Runs in a dedicated shell loop.
- **Smart Creation**: Automatically detects if you want a file or a folder.
- **Recursive Delete**: Easily remove files or entire directories.

## 🚀 Commands
| Command | Action | Example |
| :--- | :--- | :--- |
| `gx` | Create a File or Folder | `gx my_folder` or `gx note.txt` |
| `gxd` | Delete a File or Folder | `gxd old_file.txt` |
| `exit` | Exit the shell | `exit` or `Ctrl+X` |

## 🛠️ Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/joyal777/gopher-cli.git](https://github.com/joyal777/gopher-cli.git)
   cd gopher-cli

2. **Run the shell directly:**
    ```bas
    go run main.go

3. **(Optional) Build the executable:**
    ```bash
    go build -o gx.exe
    ./gx.exe

**📂 Project Structure**

main.go: Contains the core logic for the interactive loop and file system operations.

go.mod: Manages project dependencies.

**🤝 Contributing**
This is an open-source project! Feel free to fork, add new commands (like file renaming or listing), and submit a pull request.

Built with ❤️ using Golang.