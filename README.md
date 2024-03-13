# Price Calculator with Tax Rates

Hello there! I've developed a Go project that's all about calculating the price of items including different tax rates, and thoughtfully outputting the results in JSON format. I wanted to leverage Go's powerful concurrency features to make the processing efficient and fast. Let me walk you through what I've built, the concepts I've employed, and how you can get it running on your machine.

## Project Overview

This project is designed to read a list of item prices from a text file, calculate the added tax for multiple tax rates concurrently, and output the results into separate JSON files for each tax rate. It showcases Go's concurrency model, interface abstraction, and error handling, among other features.

### Key Features

- **Concurrent Tax Calculation**: Uses goroutines and channels to perform tax calculations for different rates simultaneously.
- **Flexible File Handling**: Implements file reading and writing to manage input prices and output the calculated prices with tax.
- **Interface Driven Design**: The `IOManager` interface abstracts away the details of input and output operations, allowing for easy extension.
- **Comprehensive Error Handling**: Carefully handles potential errors that might occur during file operations and data processing.

### Concepts and Technologies Used

- **Goroutines and Channels**: For managing concurrent tasks and communication between them.
- **Interfaces**: To abstract the input and output operations, allowing for different implementations (e.g., files or command line).
- **Structs and Receiver Functions**: Organizes data and associated methods for handling tax calculations.
- **Error Handling and Propagation**: Ensures robustness by checking and handling errors appropriately.
- **JSON Encoding**: For outputting the calculation results in a structured format.

### Project Structure

- `main`: Sets up the concurrency logic and orchestrates the flow of the program.
- `prices`: Contains the core logic for tax calculations and managing the calculation jobs.
- `filemanager`: Provides a concrete implementation of the `IOManager` interface for handling file input and output.
- `iomanager`: Defines the `IOManager` interface for reading and writing operations.
- `conversion`: Offers utility functions for data type conversions.
- `cmdmanager`: An example (commented out) of another possible `IOManager` implementation using command line inputs and outputs.

## Getting Started

To get this project up and running, you'll need Go installed on your system. Follow these steps:

1. Clone the project to your Go workspace.
2. Navigate to the project directory and run `go build`.
3. Start the application by executing the built binary, e.g., `./price-calculator`.

Make sure you have a `prices.txt` file in the project directory with one price per line before running the application. The output will be generated in JSON files named according to the tax rates, e.g., `result_7.json` for a 7% tax rate.
