# Cipher
> Encode and decode anything with a variety of ciphers

## Features
* Mnemonic and consistent flags(e.g. -d/--decode for every cipher that supports that)
* Encrypts using the selected cipher by default
* Supports Substitution, Caesar, Vigenère, Rail Fence, Base64 with more coming in the future

## Examples
```
cipher caesar -s 1 -f input.txt
echo "Tom Tom - Holy Fuck" | cipher base64
echo "Rocket" | cipher substitution -a "zbcdvfghijkumnopqrstuvwxyz" -f input.txt
echo "Q2FnZSBUaGUgRWxlcGhhbnQ=" | cipher base64 -d
```

## Usage
```zsh
cipher -h
```
```zsh
Encode and decode anything with a variety of popular ciphers with a mnemonic and constant command syntax

Usage:
  cipher [command]

Available Commands:
  base64       Encode and decode using Base64
  caesar       Encode and decode using the Caesar cipher
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  railfence    Encode and decode using the Rail Fence cipher
  substitution Encode and decode using the Substitution cipher
  vigenere     Encode and decode using the Vigenère cipher

Flags:
  -f, --file string   Input file
  -h, --help          Help for cipher

Use "cipher [command] --help" for more information about a command.
```

## Installing
1. You can get an executable from the [release page](https://github.com/0xalby/cipher/releases)
2. Using Go's package manager
```zsh
go install github.com/0xalby/cipher@latest
```
3. From source
```zsh
git clone https://github.com/0xalby/cipher
cd cipher
go mod tidy
make install
man cipher
```

## Utilities
```zsh
go install github.com/go-delve/delve/cmd/dlv@latest
sudo pacman -Syu pandoc || sudo apt install pandoc
```