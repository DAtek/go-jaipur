[![codecov](https://codecov.io/gh/DAtek/go-jaipur/graph/badge.svg?token=7QLVUTCXMZ)](https://codecov.io/gh/DAtek/go-jaipur)

# Japipur board game in Go

## Project structure

### `app`

The package implements interactions between an `io.Reader` , `io.Writer` and the core game logic.

### `cmd`

A console interface to the game.

### `core`

Pure domain logic.

### `fsm`
A very minimalistic implementation of a Finite State Machine, used by the `app` package.
