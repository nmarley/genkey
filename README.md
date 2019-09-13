# genkey

> ECDSA key generation utility

Generates an ECDSA key suitable for use in Dash or similar cryptocurrency. By
default, the private key is base58 encoded as a compressed WIF using Dash
prefixes.

## Table of Contents
- [Install](#install)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Install

```sh
go get -u github.com/nmarley/genkey
```

## Usage

Example:

```sh
genkey

6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000
```

## Contributing

Feel free to dive in! [Open an issue](https://github.com/nmarley/genkey/issues/new) or submit PRs.

## License

[ISC](LICENSE)
