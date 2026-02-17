# GX-Shell 🐹 (Version 2.0)

A powerful, custom interactive shell built in Go. GX-Shell transforms your standard terminal into a specialized environment for rapid file system management with custom "GX" prefixed commands.

## ✨ Features
- **Interactive Environment**: Real-time path display in your prompt.
- **Smart File Management**: Fast creation, navigation, and deletion.
- **Storage Analytics**: Calculate the total size of files and directories (auto-converts to KB/MB).
- **Directory Navigation**: Full support for moving through your file system.

## 🚀 Command Reference

| Command | Action | Example |
| :--- | :--- | :--- |
| `gx`  | **Create** a File or Folder | `gx notes.txt` or `gx source_code` |
| `gxd` | **Delete** (Recursive) | `gxd old_folder` |
| `gxc` | **Change Directory** (cd) | `gxc ..` or `gxc projects` |
| `gxl` | **List** items in directory | `gxl` |
| `gxs` | **Check Storage Size** | `gxs my_videos` |
| `exit`| **Quit** GX-Shell | `exit` or `Ctrl+X` |

## 🛠️ Getting Started

### Prerequisites
- [Go](https://go.dev/doc/install) installed on your machine.

### Installation
1. **Clone & Enter:**
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