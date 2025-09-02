# 🌐 Website Health Checker

[![Go Report Card](https://goreportcard.com/badge/github.com/<your-username>/website-health-check)](https://goreportcard.com/report/github.com/<your-username>/website-health-check)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A simple **Go CLI tool** to check if websites are up or down.  
It tries to open a TCP connection to the given domain and port, and reports the status.

---

##  Features
- Check if one or multiple websites are reachable  
- Custom ports (defaults to `80`)  
- Connection details (local → remote)  
- Lightweight and fast (pure Go)  

---

## Installation

### From source
Make sure you have [Go installed](https://go.dev/doc/install).

### 1. Clone the repository
```sh
git clone https://github.com/dharwee/website-health-check.git
cd website-health-check
go build -o healthcheck

```
### 2. Build the Binary
```sh
go build -o healthcheck
```
### 3. Run the tool
```sh
./healthcheck --help
```

### Examples
Check a single website (default port 80):
./healthcheck -d google.com

Check multiple websites:
./healthcheck -d google.com -d github.com