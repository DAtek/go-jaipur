coverfile := ".coverage"
pkgs := "./core ./app ./fsm"


test *opts:
    go test {{ pkgs }} {{ opts }}


test-cover *opts:
    go test -coverprofile {{ coverfile }} {{ pkgs }} {{ opts }}


show-coverage:
    go tool cover -html={{ coverfile }}


test-and-show-coverage: test-cover show-coverage

run:
    go run ./cmd