run: 
	go run main.go

test:
	go test -v -cover . ./state_machine

install:
	go get -u -v