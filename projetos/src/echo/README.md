# echo

repositório: https://github.com/labstack/echo

## Installation

```bash
// go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4
```


```bash
✗ go run echo.go 
echo.go:4:3: no required module provides package github.com/labstack/echo/v4: go.mod file not found in current directory or any parent directory; see 'go help modules'
echo.go:5:3: no required module provides package github.com/labstack/echo/v4/middleware: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

```bash
✗ go mod init github.com/labstack/echo/v4
go: creating new go.mod: module github.com/labstack/echo/v4
go: to add module requirements and sums:
        go mod tidy
```

```bash
✗ go mod init github.com/labstack/echo/v4/middleware
go: /home/fahham/projetos/projetos-go/src/echo/go.mod already exists
```

```bash
✗ go mod tidy
go: finding module for package github.com/labstack/echo/v4/middleware
go: downloading github.com/labstack/echo v3.3.10+incompatible
go: github.com/labstack/echo/v4 imports
        github.com/labstack/echo/v4/middleware: module github.com/labstack/echo@latest found (v3.3.10+incompatible), but does not contain package github.com/labstack/echo/v4/middleware
```