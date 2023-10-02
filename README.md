# Go Echo Server

A simple Go-based echo server that echoes back the URL path for incoming HTTP requests.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Building the Server](#building-the-server)
  - [Running the Server](#running-the-server)
- [Usage](#usage)
- [License](#license)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) v1.21.0 installed on your machine.

## Getting Started

To get started with this echo server, follow these steps:

### Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/Bloxico/golang-echo-server.git
   ```

2. Build the Go Echo Server

After cloning the repository, navigate to the project's root directory using your terminal. To build the Go echo server, run the following command:

  ```bash
  go build -o echo-server main.go
  ```

This command will compile the Go code and generate an executable binary named echo-server in the same directory.

3. Start the Echo Server

Now that you have the `echo-server` binary, you can start the echo server on your local machine. By default, it will listen on port `8080`. To start the server, run the following command:

  ```shell
  ./echo-server
  ```

You should see a message indicating that the server is running. By default, it listens on `http://localhost:8080`.  

To test the server, open your web browser or use a tool like curl to send HTTP GET requests to `http://localhost:8080/echo`. You should receive a JSON response like this:
```json
{
    "message": "Hello, World!"
}
```  

*Note: If you want to change the server's port, you can do so by setting the PORT environment variable before starting the server.*

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

