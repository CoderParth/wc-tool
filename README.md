# wc-tool
My own version of the Unix command line tool wc!

## To try it out

## Clone the repo:

```
git clone https://github.com/CoderParth/wc-tool.git
```

```
cd wc-tool
```

## Build 

```
go build -o wc
```

## Usage

The tool supports several command-line options:

- `-c`: Count the number of bytes in a file.
- `-l`: Count the number of lines in a file.
- `-w`: Count the number of words in a file.
- `-m`: Count the number of characters in a file.

If no option is provided, the tool will count bytes, lines, and words.

```
./wc [option] filename
```

## Examples:
```
./wc -c test.txt

Bytes:  342190
```
