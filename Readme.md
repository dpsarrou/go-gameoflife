# Game Of Life (in Go)

Implementation of the GameOfLife game in Golang (1.19)

## Installation

Cloning this repository is enough to run the program if you have a recent
version of Golang already installed on your system. For this implementation
Go **1.19.5** was used.

After you have cloned the repository, you can run the tests to verify that
everything is working correctly:

```shell
    make test
```

or alternatively use golang directly: `go test -cover ./gameoflife/...`

You should expect a very similar output:

```text
    ok      go-gameoflife/gameoflife        0.240s  coverage: 100.0% of statements
```

## Usage

To run the program using a randomly generated universe you can use:

```shell
    make run
```

or alternatively use golang directly: `go run .`

To run the program using a premade universe with a Glider pattern you can use:

```shell
    make run-glider
```

or alternatively use golang directly: `go run . -glider`

Provided a correct execution, the first output would look something like this:

```text




                 *
                  *
                ***






```

To stop the program simply stop the process pressing `ctrl+c` on your keyboard.

## Documentation

### Code structure

- The [gameoflife](./gameoflife/) folder is a single go package that contains
  the Entities of the game along with their tests. It captures the modeling of
  the GameOfLife Domain and the business rules.
- The [main.go](main.go) file is the entrypoint for the CLI program that
  instantiates the game Engine using a seed provided from a flag as user
  input and runs it.

### Architecture & Design

The GameOfLife is in principle a very simple application. The whole
functionality can be implemented in a single file using a minimal amount of
functions, a lot of nested loops and direct access to the internal state.
However such implementation is not very flexible. It is not easy to test for
the correctness of individual pieces of code, and simple modifications can
result into changes to multiple places. And that is fine when a single person
is working in the codebase, but that is rarely true in modern software
development. This implementation chooses to adhere most closely to the modern
software development practises and makes a concious decision to increase the
code size by separating concerns, optimizing function APIs for testability and
maintainability.

The guiding principles that were used for designing the solution:

- Get feedback fast

  Using this principle, to ensure the program works and behaves as expected,
  every code file is acompanied by automated tests, prioritizing testing wide
  areas of code rather than creating tests for every single function. This is
  a deliberate choice to balance effort vs value.

- Single Responsibility, Separation of concerns

  The functionality has been explicitly grouped under specific models to ensure
  that each is simple to argue, has clear responsibilities and cohesive
  behavior. As a result, the following entities have been identified:

  - Cell

    A single Cell in the GameOfLife which controls its state of being Alive/Dead
    and can Interact with its neighbooring cells to change it

  - Universe

    A 2D matrix of Cells, that can evolve by triggering interaction betwen Cells
    at a given position in the matrix.

  - Engine

    The game engine handling the continuous generation of Universe evolutions

- InformationExpert Principle

  Each model controls its own information as as a result must be responsible for
  enforcing related business rules. This allows for the rules to be closer to
  the state they control, and contain them in a single place without allowing
  for information "leaking". In this case Cell controls the Alive/Dead
  state and it is the one that enforces the GameOfLife rules who alter it.
  Applying this principle results in Telling the Cell to update its state rather
  than Asking the Cell of its state, checking the rules outside of it, then
  updating its state (a common antipattern also called the AnemicDomainModel).
