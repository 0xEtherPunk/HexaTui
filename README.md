# HexaTui 🖥️

## 🚀 About

HexaTui is a TUI (Text User Interface) client for interacting with AI language models through your terminal. Built with Go and the Bubble Tea framework, it provides a lightweight alternative to web-based chat interfaces.

Key technical aspects:

- Built on Go for efficient resource usage
- Uses Bubble Tea framework for TUI rendering
- Handles async I/O for responsive experience
- Implements VT100 terminal sequences
- Supports UTF-8 text encoding
- Maintains chat history in local storage


## 🎮 Controls

| Key         | Action           |
| ----------- | ---------------- |
| `Enter`     | Send message     |
| `Ctrl+C`    | Exit application |
| `Ctrl+U`    | Clear input line |
| `Backspace` | Delete character |


## 📥 Installation

### 1. Check Go Installation

```
go version
```

#### If Go is not installed: 

#### Ubuntu/Debian

```
sudo apt update && sudo apt install golang-go
```

#### macOS

```
brew install go
```

#### Windows:

download and install from https://go.dev/dl/


### 2. Get HexaTui

Clone the repository

```
git clone https://github.com/your-username/HexaTui.git
cd HexaTui
```

### 3. Install and Run

Install dependencies:

```
go mod tidy
```

Run the application:

```
go run cmd/main.go
```


### Troubleshooting

If you encounter issues:

- Clean and reinstall dependencies

```
rm go.sum
go mod tidy
```

> BTW, this is my first TUI project built with Go using the Bubble Tea framework for practice with terminal interfaces.


## Development

The project uses:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for TUI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for styling
- [Reflow](https://github.com/muesli/reflow) for text formatting


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.