# Goroutines and load shedding

The sub directory is a module that will start up on its own if needed. If run directly (ie with `go run main.go` in the `sub` directory), you can ctrl+c out of it as normal, because handling and capturing the sigint / sigterm happens outside of the actual app logic.

The `main` directory is nothing but an orchestrator. It spins up n numbers of the subs, and it takes care of capturing and handling the sigint/sigterms, and then communicating them to the spun up apps.

That way when the orchestrator is terminated, it first tears down all the other sub packages, and only then cleans up after itself.

## Demo
### Sub package on its own
```shell
$ cd sub
$ go run main.go
```
This will start some logging. Pressing ctrlc will show you it cleanly exits.

### Main package
```shell
$ cd main
$ go run main.go
```
This will start up the main package, start up 3 sub packages, keep track of them, and loop until needed. When you ctrlc, you will see all the other logs as the subs are being torn down.
