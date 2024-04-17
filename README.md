# Exploring Go

This repository contains some simple projects built for learning and practice purposes.

## Usage

- Either create a binary using `go build` as directed and run the binary file.
- Or navigate into a directory and run `go run .` to run the code without building it explicitly.

1. [**CLI Todo App**](https://github.com/iamrishupatel/exploring-go/tree/main/cli-todo)

   - Navigate to the `cli-todo` directory.
   - Compile the Go code using `go build` command.
   - Run the compiled executable to start the CLI todo app.
     ```bash
     ./cli-todo
     ```
   - Follow the on-screen instructions to add, remove, update, or list todo items.

2. [**Simple API Client**](https://github.com/iamrishupatel/exploring-go/tree/main/crypto-masters)

   Crypto Masters is a simple client to fetch data from the web. It fetches the price of the some currencies
   and displays them. It use goroutines, watchgroups and http.

   - Navigate to `crypto-masters` directory
   - Compile the code using `go build` comamnd
   - Run the compiled executable.
     ```bash
     ./crypto-masters
     ```
   - It should print rates of some currencies from the cex api.

3. [**Simple Web Server**](https://github.com/iamrishupatel/exploring-go/tree/main/simple-web-server)

   - Navigate to `simple-web-server` directory.
   - Compile the code using `go build` command.
   - Run the compiled executable.
     ```bash
     ./simple-web-server
     ```
   - Visit `http://localhost:8080` in your browser to see the server in action.
   - You can add new data by sending a POST request to `http://localhost:8080/add-museum` with a JSON payload.
