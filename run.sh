export PATH=$(go env GOPATH)/bin:$PATH 
swag init -g routers/endpoints.go # init swag
go run . # runs app from main.go

