% cipher(1) v1.0 | User Commands
% Alberto Chiaravalli
% March 2025

# NAME
cipher - Encrypt and decrypt anything with a variety of ciphers

# SYNOPSIS

**cipher** [*OPTIONS*] *COMMAND* [*ARGS*]

# DESCRIPTION

**cipher** is a command-line tool for encoding and decoding text using a variety of popular ciphers. It supports Caesar, Vigenère, Substitution and Rail Fence with more coming in the future!

# OPTIONS

-f, --file=*FILE*
: Input file to read from. If not provided, input can be read from stdin.

-h, --help
: Display help information for the command.

# COMMANDS

## caesar

Encrypt and Decrypt using the Caesar cipher.

-s, --shift=*SHIFT*
: Shift value for the Caesar cipher(required).

-d, --decrypt
: Decrypt the input.

## vigenere

Encrypt and Decrypt using the Vigenère cipher.

-k, --key=*KEY*
: Key for the Vigenère cipher (required).

-d, --decrypt
: Decrypt the input.

## substitution

Encrypt and decrypt using the Substitution cipher.

-a, --alphabet=*ALPHABET*
: Substitution alphabet(26 unique letters, required).

-d, --decrypt
: Decrypt the input.

## railfence

Encrypt and decrypt using the Rail Fence cipher.

-r, --rails=*RAILS*
: Number of rails for the Rail Fence cipher(required).

-d, --decrypt
: Decrypt the input.

## aes

Encrypt and decrypt using the AES cipher.

-k, --key=*KEY*
: AES encryption/decryption key(32 bytes, required).

-d, --decrypt
: Decrypt the input.

## affine

-a, --a=*KEY*
: Key 'a' for the Affine cipher(must be coprime with 26 and defaults to 1)

-b, --b=*KEY*
: Key 'b' for the Affine cipher

-d, --decrypt
: Decrypt the input.

# EXAMPLES

Encrypt a file content using the Substitution cipher:
```bash
cipher substitution -a "zbcdvfghijkumnopqrstuvwxyz" -f input.txt
```

Encrypt stdin using the Caesar cipher:
```bash
echo "rotated" | cipher caesar --shift 3
```

Decrypt a Base64 string:
```bash
echo "ZW1waXJl" | cipher base64 -d