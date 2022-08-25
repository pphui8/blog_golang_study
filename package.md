## package management in golang

#### how to use
1. run `go init [project name]`
2. edit the go mod file, add the requirement
```mod
module firstApp

go 1.17

require github.com/gin-gonic/gin v1.8.1

require (
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
```
3. add a `main.go` file and add a entry point
4. run `go run starter.go `
