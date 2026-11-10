# binary-clock

A terminal clock that shows the current time as **binary-coded decimal (BCD)**,
built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and
[Lip Gloss](https://github.com/charmbracelet/lipgloss).

Each decimal digit of `HH MM SS` gets its own column, and each lit cell is a
`1` bit, so you read the time the way a physical BCD desk clock works.

```
     H   h   M   m   S   s

 8 |     ▒▒      ▒▒      ▒▒
 4 |     ▒▒  ▒▒  ██  ██  ██
 2 | ▒▒  ██  ██  ▒▒  ▒▒  ██
 1 | ██  ▒▒  ██  ▒▒  ██  ▒▒
     1   2   3   4   5   6
```

The example above is `12:34:56` — the digit row at the bottom shows the answer.

## Features

- Live clock, refreshed every second
- Classic BCD layout: rows are bit values `8 4 2 1`, columns are decimal digits
- Digit summary row so you can cross-check your reading
- Adaptive colors for light and dark terminal themes
- Centered layout that follows terminal resizes

## How it works

Every digit of the time is encoded in binary on its own:

| Column | Meaning       | Digits | Bits |
|--------|---------------|--------|------|
| `H`    | hour tens     | 0–2    | 2    |
| `h`    | hour units    | 0–9    | 4    |
| `M`    | minute tens   | 0–5    | 3    |
| `m`    | minute units  | 0–9    | 4    |
| `S`    | second tens   | 0–5    | 3    |
| `s`    | second units  | 0–9    | 4    |

Rows are read top to bottom as bit values `8`, `4`, `2` and `1`. Add up the lit
rows in a column to get its digit — e.g. `██` on rows `2` and `1` is `2+1 = 3`.

## Requirements

- Go 1.24 or newer
- A terminal with Unicode block characters (`██`, `▒▒`)

## Install

```sh
go install github.com/Ah3ron/binary-clock/cmd/binary-clock@latest
```

Or build from source:

```sh
git clone https://github.com/Ah3ron/binary-clock.git
cd binary-clock
make build
./bin/binary-clock
```

## Usage

```sh
binary-clock
```

| Key       | Action |
|-----------|--------|
| `q`       | quit   |
| `ctrl+c`  | quit   |

## Development

```sh
make help    # list available targets
make run     # run the clock
make test    # run tests
make check   # formatting check + go vet + tests (used by CI)
make build   # build into ./bin
make fmt     # format sources in place
make clean   # remove build artifacts
```

## Project structure

```
.
├── cmd/binary-clock/   # program entry point
└── internal/
    ├── bcd/            # time -> BCD digits and bit layout
    └── ui/             # Bubble Tea model, styles and rendering
```

## License

MIT — see [LICENSE](LICENSE).
