# Cipher
> Encrypt and decrypt anything with a variety of ciphers

## Features
* Mnemonic and consistent flags(e.g. -d/--decrypt for every cipher that supports that)
* Encrypts using the selected cipher by default
* Supports Substitution, Caesar, Vigenère and Rail Fence ciphers with more coming in the future

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
```
Usage:
  cipher [command]

Available Commands:
  aes          Encrypt and decrypt using the AES cipher
  affine       Encrypt and decrypt using the Affine cipher
  caesar       Encrypt and decrypt using the Caesar cipher
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  railfence    Encrypt and decrypt using the Rail Fence cipher
  substitution Encrypt and decrypt using the Substitution cipher
  vigenere     Encrypt and decrypt using the Vigenère cipher

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
sudo make install
man cipher
```

## Utilities
```zsh
go install github.com/go-delve/delve/cmd/dlv@latest
sudo pacman -Syu pandoc || sudo apt install pandoc
```

## Contributing
Check out [TODO.md](./TODO.md) and send a PR for me to review