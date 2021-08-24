# /bin/sh

rm ./bin/main
go build -o ./bin/main ./cmd/main.go 
./bin/main
