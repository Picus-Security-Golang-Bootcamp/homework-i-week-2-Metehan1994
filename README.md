# Homework | Week 2 | Booklist App

## Overview

The program interprets a number of books formatted in json file, list them and searches a specific book among them.

## How to Use the App ?

The program works with two command line arguments which are "list" and "search". "list" command does not require any additional keywords whereas "search" command is used with the book name which is sought.
***

### list command

```
go run main.go list
```

"list" command shows the booklist.
***

### search command

```
go run main.go search <bookName>
```

"search" command with book name gives the book name if it is available in the list. Otherwise, the program informs you about absence of book which is asked.
***

### Some Notes for Usage

1. Program produces an error message when it is executed without any arguments.

2. Using "search" argument without book name or "list" argument with multiple words triggers another error.

3. The arguments other than "list" and "search" are not allowed and same error message in the second note pops up.

4. The program is insensitive to uppercase & lowercase letters for booknames which makes easier to search the book.

## Package Used

- The program is created with **GO main package**.

- Its import modules like **"fmt", "strings", "os", "io/ioutil"** and **"encoding/json"** are used.
