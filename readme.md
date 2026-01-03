---

# ⚡ NetPulse

> A modern, high-performance CLI to check website availability and latency in real time.

![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat\&logo=go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

---

## 📖 About

NetPulse runs quick health checks against URLs, sorts results by latency, and renders a compact terminal table with status badges. It leans on Go’s concurrency (goroutines + channels) to process many targets at once while keeping a simple CLI surface.

---

## ✨ Features

* 🚀 Concurrency-first worker model for fast batches of URLs
* 🎨 Colored, table-based TUI powered by Lipgloss
* 📊 HTTP status, latency, and error surfacing
* 📂 CSV input for bulk checks
* 🛡️ Configurable request timeout per run

---

## 🛠️ Tech Stack

* Go 1.25+
* Cobra (CLI) + Viper (config-ready)
* Lipgloss (terminal styling)
* Testify (testing)

---

## 🚀 Quick Start

```bash
git clone https://github.com/dcastro0/netpulse.git
cd netpulse
go mod tidy
make build
```

The binary is written to `bin/netpulse`.

---

## 💻 Usage

Check a single URL:

```bash
./bin/netpulse check https://google.com
```

Check from a CSV (one URL per line):

```bash
./bin/netpulse check --file websites.csv
```

Tune timeout (seconds, default 5):

```bash
./bin/netpulse check https://example.com --timeout 8
```

### Flags

* `--file`, `-f` — Path to CSV with URLs
* `--timeout`, `-t` — Request timeout in seconds (default: 5)

### Sample Output

```
STATUS     URL                       LATÊNCIA
🟢 200     https://google.com        240ms
🔴 ERROR   https://broken-api.com    45ms
🟡 403     https://protected.com     110ms
```

---

## 🧪 Tests

```bash

# or
make test
```

---

## 🧩 Architecture (Fan-In)

1. Read input (arg or CSV) and fan out to goroutines
2. Workers execute HTTP GET with timeout
3. Results flow into a buffered channel (fan-in)
4. Aggregate, sort by latency, and render a styled table

---

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

Made with 💜 by **Caio Corrêa de Castro**

---

