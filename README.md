# bingo-game (Simple Go HTTP Server with API)

This is a simple Go program that implements an HTTP server with a basic API. The API includes three endpoints:

1. **Generate a Number (/generate)**: Generates a distinct random number between 1 and 99, ensuring it is different from the last generated number.

2. **List All Numbers (/list)**: Retrieves a list of all numbers generated by the server.

3. **Reset Numbers (/reset)**: Resets the list of generated numbers.

## How to Run

1. Ensure you have Go installed on your machine.

2. Save the provided code to a file, for example, `main.go`.

3. Open a terminal and navigate to the directory where the file is saved.

4. Run the following command to start the server:

   ```bash
   go run main.go
