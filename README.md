# Build Your Own wc Tool

## Overview

This challenge invites you to create your own version of the Unix command-line tool `wc`. The Unix philosophy emphasizes writing simple tools that do one thing well and can be combined to create complex data processing pipelines. 

The `wc` command stands for **word count** and provides various statistics about text files, including the number of lines, words, characters, and bytes. This project will guide you through building a simplified version, named `ccwc` (Coding Challenges Word Count).

## Unix Philosophy

The principles of Unix philosophy that you will be following are:

1. **Write simple parts connected by clean interfaces**: Each tool does just one thing and provides a simple command-line interface (CLI) that handles text input from either files or streams.
2. **Design programs to be connected to other programs**: Each tool can easily connect to others, creating powerful compositions.

You can read more about the Unix Philosophy in the book *The Art of Unix Programming*.

## The Challenge - Building `wc`

The functional requirements for `wc` are concisely described in its man page. You can view this by running `man wc` in your terminal.

### Requirements
- Count the number of lines, words, characters, and bytes in a file.
- Support command-line options:
  - `-c` for byte count
  - `-l` for line count
  - `-w` for word count
  - `-m` for character count (considering multibyte characters)
