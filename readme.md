---

# ⚡ NetPulse

> A modern, high-performance CLI tool for website health monitoring built with Go.

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat\&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

---

## 📖 About

**NetPulse** is a command-line interface (CLI) tool designed to check the availability and latency of websites in real time.
Built with performance in mind, it leverages Go’s native concurrency model to process large volumes of URLs simultaneously, rendering results in a clean and structured terminal UI.

This project demonstrates practical usage of:

* **Goroutines**
* **Channels**
* **Fan-In concurrency patterns**
* **Clean Architecture principles**

---

## ✨ Features

* 🚀 **High Concurrency**
  Checks hundreds of websites in parallel using worker pools.

* 🎨 **Modern UI**
  Beautiful terminal output with tables and status badges (powered by **Lipgloss**).

* 📊 **Detailed Metrics**
  Displays HTTP status codes, latency, and error details.

* 📂 **Bulk Processing**
  Supports loading target URLs from CSV files.

* 🛡️ **Reliable**
  Configurable timeouts and robust error handling.

---

## 🛠️ Tech Stack

* **Language:** Golang
* **CLI Framework:** [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper)
* **UI / TUI:** [Lipgloss](https://github.com/charmbracelet/lipgloss)
* **Testing:** [Testify](https://github.com/stretchr/testify)

---

## 🚀 Getting Started

### Prerequisites

* Go **1.22+**

### Installation

```
git clone https://github.com/seu-usuario/netpulse.git
cd netpulse
go mod tidy
```

### Build

```
make build
```

The binary will be generated in the `bin/` directory.

---

## 💻 Usage

### Check a Single URL

```
./bin/netpulse check https://google.com
```

### Check Multiple URLs (via CSV)

Create a `websites.csv` file containing one URL per line:

```
./bin/netpulse check --file websites.csv
```

### Output Example

```
STATUS  URL                      LATÊNCIA
🟢 200  https://google.com        240ms
🔴 500  https://broken-api.com     45ms
🟡 403  https://protected.com     110ms
```

---

## 🧪 Running Tests

```
make test
```

---

## 🧩 Architecture Highlight

The application follows a **Fan-In concurrency pattern**:

1. **Dispatcher**
   Reads input (CLI args or CSV) and dispatches URLs to workers.

2. **Workers**
   Each worker runs in its own Goroutine and performs an HTTP check independently.

3. **Aggregator (Fan-In)**
   All results are sent to a buffered channel.

4. **Renderer**
   The main routine collects results, sorts them by latency, and renders the final UI table.

This approach ensures **high throughput**, **low latency**, and **clean separation of concerns**.

---

## 📄 License

This project is licensed under the **MIT License**.

---

Made with 💜 by **Caio Corrêa de Castro**

---
