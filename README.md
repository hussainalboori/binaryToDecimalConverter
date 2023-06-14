# Binary to Decimal Converter

This is a simple command-line tool that converts a binary number to its decimal equivalent. It takes a binary number as input, converts it to decimal, and outputs the result.

## Usage

To use the Binary to Decimal Converter, follow these steps:

1. Ensure you have Go installed on your system. You can download and install it from the official Go website: https://golang.org/

2. Clone this repository to your local machine or download the source code.

3. Open a terminal or command prompt and navigate to the project directory.

4. Build the project by running the following command:
   ```
   go build
   ```

5. Run the program with the input and output filenames as command-line arguments:
   ```
   ./binary-to-decimal <input_file> <output_file>
   ```
   Replace `<input_file>` with the name of the file containing the binary number you want to convert, and `<output_file>` with the name of the file where you want to save the decimal result.

6. The program will convert the binary number to decimal and save the result in the specified output file.

7. You can open the output file to view the decimal result.

## Example

Let's say you have a file named `input.txt` containing the binary number `1010`. You want to convert this binary number to decimal and save the result in a file named `output.txt`. Here's how you can do it:

```
./binary-to-decimal input.txt output.txt
```

The program will read the binary number from `input.txt`, convert it to decimal (which is `10` in this case), and save the result in `output.txt`.

Feel free to customize the content and sections based on your specific project requirements.
