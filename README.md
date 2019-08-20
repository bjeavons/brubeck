brubeck
=======

brubeck is a simple Unix time change & conversion command-line application
written in Go. Its purpose is to provide quick access to relative times and
convert to and from Unix time.

## Examples
```go
$ brubeck
1566265290
$ brubeck 1 week ago
1565660500
$ brubeck 1565660500 in edt
2019-08-12 21:41:40 -0400 EDT
```
