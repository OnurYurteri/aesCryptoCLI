# aesCryptoCLI

A command line interface written in Go for encrypting and decrypting files.

~~~
NAME:
   Aes Crypto CLI - Encrypt files with Advanced Encryption Standard (AES)

USAGE:
   aesCryptoCLI [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Onur Yurteri

COMMANDS:
     createKey, c  Create AES Key
     run, r        Run Encrypt/Decrpyt operation for given file
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -k value       Path for key file
   -i value       Path for input file
   -o value       Path for output file
   --help, -h     show help
   --version, -v  print the version
~~~

## How to build & run

```bash
go mod init aes-crypto-cli
go get github.com/urfave/cli

# Build executable named 'aes-crypto-cli'
go build -o aes-crypto-cli
```

## How to use

```bash
# Show help
aes-crypto-cli --help
```

```bash
# Create a key file named aes.key
aes-crypto-cli createKey
```

```bash
# Encrypt README.md file with aes.key and save it as README.encrypt
aes-crypto-cli -i README.md -o README.encrypt -k aes.key run
```

```bash
# Decrypt README.encrypt file with aes.key and save it as README.decrypt
aes-crypto-cli -i README.md -o README.encrypt -k aes.key run
```