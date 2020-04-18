# sendkeys [![GoDoc](https://godoc.org/git.sr.ht/~mna/sendkeys?status.svg)](http://godoc.org/git.sr.ht/~mna/sendkeys) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/git.sr.ht/~mna/sendkeys)

Package sendkeys provides functions to simulate user input in the terminal by
sending keys to a TTY. Useful especially to test terminal programs that run in
raw mode. See the [package documentation][godoc] for details, API reference and
usage example (alternatively, on [pkg.go.dev][pgd]).

* Canonical repository: https://git.sr.ht/~mna/sendkeys
* Issues: https://todo.sr.ht/~mna/sendkeys

It only works on Unix-like systems. Note that programs using this will require
sudo-like privileges.

## License

The [BSD 3-Clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[godoc]: http://godoc.org/git.sr.ht/~mna/sendkeys
[pgd]: https://pkg.go.dev/git.sr.ht/~mna/sendkeys
