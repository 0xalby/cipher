% cipher(1) v1.0 | User Commands
% Alberto Chiaravalli
% October 2023

# NAME

cipher - Encode and decode anything with a variety of popular ciphers with a mnemonic and constant command syntax

# SYNOPSIS

**cipher** [*OPTIONS*] *COMMAND* [*ARGS*]

# DESCRIPTION

**cipher** is a command-line tool for encoding and decoding text using a variety of popular ciphers. It supports Base64, Caesar, Rail Fence, Substitution, and Vigenère ciphers. The tool provides a consistent and mnemonic command syntax for ease of use.

# OPTIONS

-f, --file=*FILE*
: Input file to read from. If not provided, input is read from stdin.

-h, --help
: Display help information for the command.

# COMMANDS

## base64

Encode and decode using Base64.

-d, --decode
: Decode the input instead of encoding.

## caesar

Encode and decode using the Caesar cipher.

-s, --shift=*SHIFT*
: Shift value for the Caesar cipher (required).

-d, --decrypt
: Decrypt the input instead of encrypting.

## completion

Generate the autocompletion script for the specified shell.

**bash**
: Generate autocompletion script for Bash.

**zsh**
: Generate autocompletion script for Zsh.

**fish**
: Generate autocompletion script for Fish.

**powershell**
: Generate autocompletion script for PowerShell.

## railfence

Encode and decode using the Rail Fence cipher.

-r, --rails=*RAILS*
: Number of rails for the Rail Fence cipher (required).

-d, --decrypt
: Decrypt the input instead of encrypting.

## substitution

Encode and decode using the Substitution cipher.

-a, --alphabet=*ALPHABET*
: Substitution alphabet (26 unique letters, required).

-d, --decrypt
: Decrypt the input instead of encrypting.

## vigenere

Encode and decode using the Vigenère cipher.

-k, --key=*KEY*
: Key for the Vigenère cipher (required).

-d, --decrypt
: Decrypt the input instead of encrypting.

# EXAMPLES

Encode a file content using the Substitution cipher:
```bash
cipher substitution -a "zbcdvfghijkumnopqrstuvwxyz" -f input.txt
```

Encode stdin using the Caesar cipher:
```bash
echo "rotated" | cipher caesar --shift 3
```

Decode a Base64 string:
```bash
echo "ZW1waXJl" | cipher base64 -d