# flipper

just a dumb cli tool cause im tire of things that watch only a subset of my files
im sure there are better tools out there but i dont care

## Usage

### General Usage

```bash
flipper -command "echo hello world"
```

### Exclude a directory

```bash
flipper -ex "vendor" -command "echo hello world"
```

#### Exclude multiple directories

```bash
flipper -ex "vendor" -ex "node_modules" -command "echo hello world"
```

### Exclude an extension

```bash
flipper -ex-ext "go" -command "echo hello world"
```

#### Exclude multiple extensions

```bash
flipper -ex-ext "go" -ex-ext "js" -command "echo hello world"
```

## Installation

```bash
go install github.com/karim-w/flipper@latest
```

## License

BSD 3-Clause License

## Contributing

Pull requests are welcome. For major changes, please open an issue first to
discuss what you would like to change.

## Author

Karim-W
