# cppinit

`cppinit` is a CLI utility written in Go that generates empty C++ projects for
my particular configuration.

Namely, VSCode + CMake + Catch + Boost.

To acquire/install:

1. Download and install the Go toolkit (1.10 was used at time of writing)
2. Clone the repository
3. `go build`
4. Take the compiled binary and add it to a dedicated folder on your system's PATH


To use:

Build an empty project with:
```
./cppinit -name=sample-proj
```

This will fill the current directory with all the files you need to get up and running
with a C++ project (writing tests and slight VSCode support).

To purge the current directory,
```
./cppinit -clean
```

*Note*: Inovking `cppinit -clean` will delete anything in the `include/` and `test/` directories which means one can lose quite a bit of work. Always make sure `cppinit -clean` is invoked with caution. And always have backups of anything you care about.