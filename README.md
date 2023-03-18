# Recipe Keeper

## Getting Started
1. Make sure you have [Go](https://golang.org/) installed.
2. Clone the repo
```bash
git clone https://github.com/zuams/api-recipe-keeper.git
```
3. Remove unused dependencies
```bash
cd path/to/recipe-keeper && go mod tidy
```
4. Start
```bash
go run main.go
```
 
## Generate / Update Swagger Docs
1. Download [Swag](https://github.com/swaggo/swag) for Go by using:
```bash
go get github.com/swaggo/swag/cmd/swag
```
2. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
3. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```bash
$GOPATH/bin/swag init
```
4. Start
```bash
go run main.go
```
5. Open http://localhost:4000/swagger/index.html in your browser, you can see Swagger 2.0 Api documents.
