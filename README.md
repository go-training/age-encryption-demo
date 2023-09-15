# age-encryption-demo

[Age](https://github.com/FiloSottile/age) is a simple, modern and secure encryption tool. This repository contains a demo of how to use it.

## Before you start

Install age command line tool. See the [official installation instructions](https://github.com/FiloSottile/age#installation)

## Encrypting a file

```bash
$ age-keygen -o key.txt
Public key: your_age_public_key
$ echo "Hello World" > file.txt
$ age -r your_age_public_key -o file.txt.age file.txt
```

## Decrypting a file

```bash
age --decrypt -i key.txt file.txt.age > file.txt
```
