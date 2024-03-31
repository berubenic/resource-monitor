# Resource Monitor

This is a simple Go application that prints out system resource usage, including CPU and memory statistics.

## Structure

The application is structured into separate packages for handling CPU and memory statistics:

- `cmd/cpu`: Contains functions for fetching and printing CPU information.
- `cmd/memory`: Contains functions for fetching and printing memory information.

## Usage

To run the application, navigate to the root directory of the project and use the `go run` command:

```sh
go run cmd/main.go
```
