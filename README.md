# kanbn2md

[Kanbn](https://github.com/basementuniverse/kanbn) is a CLI Kanban board.  It stores 
the task information in markdown format and has a [vscode extension](https://github.com/basementuniverse/vscode-kanbn)
for viewing and editing the board from vscode.  What is missing is a way to render the board as a markdown table.  
This simple program takes JSON output of `kanbn board -j` on stdin and renders a markdown table on stdout.

## Installation

Download a binary from the releases page or install via go with

```bash
go install github.com/tjdavis3/kanbn2md
```

## Usage

```bash
kanbn board -j | kanbn2md > board.md
```

## Example

| Backlog | Todo | In Progress | Done |
| --- | --- | --- | --- |
| Comment the code | Create github repo | Create a README.md | Build app |
| Add tests | Add error handling |  |  |
