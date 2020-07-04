# Cryptography program

Simple encryption and decryption program inplemented in Golang.

## Table of contents

1. [Simple run](#simple-run)
2. [Create an executable program](#executable-binary-program)
3. [Usage](#usage)
4. [Command line arguments mode](#command-line-arguments-mode)

**N.B.:** To use all the command, you must have installed the [Golang environment](https://golang.org/).

## Simple run

You can run the program directly with :

```bash
go run <filename>.go
```

## Executable binary program

You can run the followig command for build one specific binary :

```bash
go build <filename>.go
```

Or you can also run this command to build all binaries in one time :

```bash
make build
```

**N.B.:** `make` application must be installed in your machine

## Usage

You can use all the program from 2 ways :

- you can just open the binary file which open an interactive program;
- you can just pass the arguments to the command when you call it.

## Command line arguments mode

For the `encryptmessage.go` file, you must run :

```bash
encryptmessage.exe <msg:string> <key:string>
```

For the `decryptmessage.go` file, you must run :

```bash
decryptmessage.exe <encrypted_byte_data:byte_array> <key:string>
```

For the `encryptmessagetofile.go` file, you must run :

```bash
encryptmessagetofile.exe <filename:string> <msg:string> <key:string>
```

For the `decryptmessagefromfile.go` file, you must run :

```bash
decryptmessagefromfile.exe <filename:string> <key:string>
```
