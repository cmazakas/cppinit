# cppinit

`cppinit` is a CLI utility written in Go that generates empty C++ projects for
my particular configuration and tastes.

It requires the following:
* [VSCode](https://code.visualstudio.com/) (and the plugin CMake Tools by vector-of-bool)
* [CMake](https://cmake.org/) (v3.11+)
* [Boost](https://www.boost.org/) (v1.67)
* [Catch](https://github.com/catchorg/Catch2) (v2)
* [Ninja](https://ninja-build.org/)

Download Boost via `brew`, `pacman` or `vcpkg`. Because Ubuntu is usually so behind, simply download and build the source.

1. Download and install the Go toolkit (1.10 was used at time of writing)
2. `go get github.com/LeonineKing1199/cppinit`
3. `cd $GOPATH/src/github.com/LeonineKing1199/cppinit`
3. `go build`
4. Take the compiled binary and add it to a dedicated folder on your system's PATH

To use:

Build an empty project with:
```
mkdir my-new-project
cd my-new-project
./cppinit -name=my-new-project

=> creates
my-new-project/
  CMakeLists.txt
  my-new-project.cmake (empty toolchain file)
  .vscode/
    settings.json (for VSCode)
  include/
    my-new-project
  test/
    main.cpp
```

To purge the current directory,
```
./cppinit -clean
```

**Note**: Inovking `cppinit -clean` will delete anything in the `include/` and `test/` directories which means one can lose quite a bit of work. Always make sure `cppinit -clean` is invoked with caution. And always have backups of anything you care about.

## For C++ Beginners

It's a goal of this project to make using C++ for newcomers simple and
straight-forward.

### Boost

Boost installation steps vary wildly from system to system but in most cases,
one can simply download the compressed source, decompress it and then use the
documentation to build Boost with `b2`, making sure that one uses the `prefix`
option to create an easy-to-use Boost root folder. Otherwise, the include and
library directories will both need to be specified.

### Creating Your CMake Toolchain File

One may notice in the above example that a `my-new-project.cmake` file is generated
by `cppinit`. This is intended to be an empty CMake toolchain file. Toolchain files
are easy ways to feed system-specific variables into CMake without polluting the root
`CMakeLists.txt` file which should only be used for system-agnostic requirements of
the project.

Sample toolchain file:
```
# Sample file for Windows users

set(BOOST_ROOT "/Users/cmaza/source/boosts/boost_1_67_0_b1")
set(Boost_USE_STATIC_LIBS ON)

# directory that contains our catch.hpp file
include_directories(
  "/vcpkg/installed/x64-windows/include/catch"
)
```

Here one can see that we define the `BOOST_ROOT` to be the directory where we told
`b2` to install Boost. We also wish to use only the static libraries and then
we help CMake find our installation of Catch which was done via `vcpkg`.

Building and testing your project:
```
cd my-new-project
mkdir build_Debug
cd build_Debug
cmake \
-DCMAKE_TOOLCHAIN_FILE=../my-new-project.cmake \
-DCMAKE_BUILD_TYPE=Debug \
-G Ninja ..
cmake --build .
./my-new-project_tests
```

`cppinit` ultimately gives the programmer a CMakeLists.txt file that they can
now edit and maintain themselves.

The core project is pulled into a "core" lib that is then linked to by the testing binary.
It defaults to an INTERFACE library by default. This needs to be changed if the
core library itself becomes non-header-only.