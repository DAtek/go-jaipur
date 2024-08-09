coverfile := ".coverage"
pkgs := "./core ./app ./fsm"


test *opts:
    gotestsum -f dots-v2 -- {{ opts }} {{ pkgs }}

test-cover *opts:
    gotestsum -f dots-v2 -- {{ opts }} -coverprofile {{ coverfile }} {{ pkgs }}

show-coverage:
    go tool cover -html {{ coverfile }}

test-and-show-coverage: test-cover show-coverage

run:
    go run ./cmd