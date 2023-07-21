# Text-hider
Hide text inside all* files.
*All binary files. No xml or json etc.
## Setup
0. Get a binary file
1. Open it in a hex editor. A good one is [hexed.it](https://hexed.it/)
2. Make your way into the end of the file. Add you hex pattern on a new line. Hex pattern can be any hex string, this is basically your key for the script.
3. Add text in hex. Turn your text into hex and add that to end of the file on a new line after the hex pattern. Good converter is [asciitohex.com](https://www.asciitohex.com/).
4. Save
5. Pick if you're using the hardcoded or argument based version. For hardcoded head over to [hardcoded](https://github.com/Juliasmatius/Text-hider#hardcoded). For argument based [argument based](https://github.com/Juliasmatius/Text-hider#argument-based)

## Hardcoded
6. Open the code
7. On line 40 replace "0affffffafafafcd" with your hex pattern.
8. On line 41 replace "demo.png" with your filename.
9. (Optional) Compile the code with ```go build -ldflags "-H=windowsgui" main.go```.
10. To run uncompiled ```go run main.go``` and when compiled just launch the executable.

## Argument based
6. (Optional) Compile the code with ```go build -ldflags "-H=windowsgui" main.go```.
7. For uncompiled extraction command is ```go run main.go <hex pattern> <filename>``` and for compiled it's ```main.exe <hex pattern> <filename>```.
