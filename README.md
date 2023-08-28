# Ariadne: CLI tool for exploring directories
This project is amed after Ariadne who was the daughter of King Minos and gave Theseus a golden ball of thread to help him escape the Labyrinth.


Ariadne is a command line utility designed to help you visualize the directory structure of a given path in a tree format. It offers features that enhance the basic tree visualization by providing options to skip specific directories and color-coding for an improved experience.

## Features

- Display a tree structure of a directory.
- Skip specific directories using flags.
- Visual distinction between files and directories.

## Usage

- Navigate to the directory where Pathos is located and run:
`go run main.go`

This will display the tree structure of the current directory.

To visualize a different directory, use the -dir flag:
`go run main.go -dir /path/to/directory`

## Flags

- dir: Specify the directory to walk through. Defaults to the current directory.
    - Example: `go run main.go -dir /path/to/directory`
- ng: Skip .git directories.
    - Example: `go run main.go -ng`
- skip: Skip specific directories. Provide a comma-separated list of directory names.
   - Example: `go run main.go -skip node_modules,build,logs`

## Sample Output

`go run main.go -ng`

```
.
|-- README.md
|-- go.mod
|-- main.go
|-- traverse
  |-- directory.go
  |-- traverse.go
  |-- visualize.go
```