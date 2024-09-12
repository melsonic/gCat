### gcat: A cat implementation in Go
**gcat** is a custom implementation of the `cat` tool in unix-based systems. This tool covers the basic functionality of the `cat` command. Basically, it concatenates the input file contents and prints to stdout. It also has the ability to read from `stdin`. 

### Running the Application
- first run `make build`
    - this will produce an executable binary file named `gcat`.
- `gcat` can take upto two flags, these are
    - `-n` to print line number in front of each line
    - `-b` to print line number only in front of non blank lines
    - `-b` has more priority than `-n`
- `gcat` takes filenames separated by commas
    - if a file path is not found, it throws a **file not found** error.
    - a special filename is `-`, which means take input from stdin.

### Example commands
- `./gcat test.txt`
    - prints `test.txt` file content.
- `./gcat test.txt -n`
    - prints `test.txt` file content with line number.
- `head -n1 test.txt | ./gcat -`
    - prints input passed through the pipe from the previous command.
- `./gcat test.txt test2.txt`
    - concatenates `test.txt` & `test2.txt` file contents and prints it.
- `sed G test.txt | ./gcat -n | head -n4`
    - `sed G` inserts a blank line after each line in `test.txt`.
    - processes input passed through the pipe from the previous command and passes the output through a pipe.
    - `head -n4` prints first 4 lines from the pipe.
- `sed G test.txt | ./gcat -b | head -n5`
    - `sed G` inserts a blank line after each line in `test.txt`.
    - processes input passed through the pipe from the previous command and passes the output through a pipe.
    - `head -n5` prints first 5 lines from the pipe.

### Installing the Application 
- just run './install.sh' and it will automatically install the application.
- make sure, you have built the application `make build` before.
