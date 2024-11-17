# Password Generator

A simple password generator written in Go.

## Features
- Generate random passwords with customizable length.
- Option to include digits, special characters, and uppercase letters.

## How to Run
```bash
go build -o password-generator
./password-generator -length 16 -digits=true -special=true -upper=false

## Flags
- `-length`: Password length (default: 12).
- `-digits`: Use digits in the password (default: true).
- `-special`: Use special characters in the password (default: true).
- `-upper`: Use uppercase letters in the password (default: true).
