# Text-hider
[![Auto Compile and Release](https://github.com/Juliasmatius/Text-hider/actions/workflows/go.yml/badge.svg)](https://github.com/Juliasmatius/Text-hider/actions/workflows/go.yml)
![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/m/Juliamatius/Text-hider)
![GitHub contributors](https://img.shields.io/github/contributors/Juliasmatius/Text-hider)
![Depfu](https://img.shields.io/depfu/dependencies/github/Juliasmatius%2FText-hider)
![GitHub all releases](https://img.shields.io/github/downloads/Juliasmatius/Text-hider/total)
![GitHub issues](https://img.shields.io/github/issues/Juliasmatius/Text-hider)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Juliasmatius/Text-hider)
![GitHub forks](https://img.shields.io/github/forks/Juliasmatius/Text-hider)


Hide text inside all* files.\
*All binary files. No xml or json etc.\
This work is licensed under [Attribution-NonCommercial-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-nc-sa/4.0/)\
\
For commercial purposes send me an email [texthider@julimiro.eu.org](mailto:texthider@julimiro.eu.org)
# Demo
1. Download hardcoded.go and demo.png
2. Run hardcoded.go with ```go run hardcoded.go```
## Setup
0. Get a binary file
1. Open it in a hex editor. A good one is [hexed.it](https://hexed.it/)
2. Make your way into the end of the file. Add you hex pattern on a new line. Hex pattern can be any hex string, this is basically your key for the script.
3. Add text in hex. Turn your text into hex and add that to end of the file on a new line after the hex pattern. Good converter is [asciitohex.com](https://www.asciitohex.com/).
4. Save
5. Pick if you're using the hardcoded or argument based version. For hardcoded head over to [hardcoded](https://github.com/Juliasmatius/Text-hider#hardcoded). For argument based [argument based](https://github.com/Juliasmatius/Text-hider#argument-based)

## Hardcoded
6. Open the harcoded.go
7. On line 40 replace "0affffffafafafcd" with your hex pattern.
8. On line 41 replace "demo.png" with your filename.
9. (Optional) Compile the code with ```go build -ldflags "-H=windowsgui" hardcoded.go```. Make sure to set enviroment variables GOOS=windows GOARCH=386
10. To run uncompiled ```go run hardcoded.go``` and when compiled just launch the executable.

## Argument based
6. (Optional) Compile the code with ```go build -ldflags "-H=windowsgui" main.go```. Make sure to set enviroment variables GOOS=windows GOARCH=386
7. For uncompiled extraction command is ```go run main.go <hex pattern> <filename>``` and for compiled it's ```main.exe <hex pattern> <filename>```.
