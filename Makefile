gen:
	protoc --proto_path=proto --go_out=./out --go_opt=paths=source_relative .\proto\*.proto

clean:
	rm out/*.go

run:
	go run main.go
