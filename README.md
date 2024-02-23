# Project htmx

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```

### What I added:

Mac/Windows/Linux agnostic .air.toml file.
I was able to accomplish this by:

1. During `make watch`, I add `export WATCHING="true"` to check if I'm watching. So we can build the executable into `tmp/main` instead for `make build`.
2. Remove `bin` inside `.air.toml`. This is so that it checks `tmp/main.exe` for windows and checks `tmp/main` for other OS (default behavior).
3. Adding an if statement that checks `$$(go env GOHOSTOS) = "windows"`.
4. If "windows", I build it in tmp/main.exe.
5. If else, I build it in tmp/main.
