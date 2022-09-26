coverfile := ".coverage"
apps := "./core ./app ./fsm"


test:
    go test {{ apps }}


test-cover:
    go test -coverprofile {{ coverfile }} {{ apps }}


show-coverage:
    go tool cover -html={{ coverfile }}


test-and-show: test-cover show-coverage

run:
    go run ./cmd