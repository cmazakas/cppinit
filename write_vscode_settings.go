package main

import "os"

func writeVSCodeSettings(proj string, file *os.File) (err error) {
	_, err = file.WriteString(`
		{
			"cmake.configureSettings": {
				"CMAKE_TOOLCHAIN_FILE": "${workspaceRoot}/` + proj + `.cmake"
			},
			"cmake.buildDirectory": "${workspaceRoot}/build_${buildType}",
			"cmake.generator": "Ninja",
			"cmake.environment": {
					"CC": "cl",
					"CXX": "cl"
			},
			"files.associations": {
				"*.cu": "cpp",
				"*.ipp": "cpp",
				"thread": "cpp",
				"memory": "cpp",
				"future": "cpp",
				"system_error": "cpp",
				"algorithm": "cpp",
				"array": "cpp",
				"atomic": "cpp",
				"chrono": "cpp",
				"cmath": "cpp",
				"codecvt": "cpp",
				"complex": "cpp",
				"condition_variable": "cpp",
				"cstddef": "cpp",
				"cstdint": "cpp",
				"cstdio": "cpp",
				"cstdlib": "cpp",
				"cstring": "cpp",
				"ctime": "cpp",
				"cwchar": "cpp",
				"deque": "cpp",
				"exception": "cpp",
				"fstream": "cpp",
				"functional": "cpp",
				"initializer_list": "cpp",
				"iomanip": "cpp",
				"ios": "cpp",
				"iosfwd": "cpp",
				"iostream": "cpp",
				"istream": "cpp",
				"iterator": "cpp",
				"limits": "cpp",
				"list": "cpp",
				"locale": "cpp",
				"map": "cpp",
				"mutex": "cpp",
				"new": "cpp",
				"numeric": "cpp",
				"ostream": "cpp",
				"queue": "cpp",
				"random": "cpp",
				"ratio": "cpp",
				"set": "cpp",
				"sstream": "cpp",
				"stack": "cpp",
				"stdexcept": "cpp",
				"streambuf": "cpp",
				"string": "cpp",
				"string_view": "cpp",
				"strstream": "cpp",
				"tuple": "cpp",
				"type_traits": "cpp",
				"typeindex": "cpp",
				"typeinfo": "cpp",
				"unordered_map": "cpp",
				"unordered_set": "cpp",
				"utility": "cpp",
				"valarray": "cpp",
				"vector": "cpp",
				"xfacet": "cpp",
				"xfunctional": "cpp",
				"xhash": "cpp",
				"xiosbase": "cpp",
				"xlocale": "cpp",
				"xlocbuf": "cpp",
				"xlocinfo": "cpp",
				"xlocmes": "cpp",
				"xlocmon": "cpp",
				"xlocnum": "cpp",
				"xloctime": "cpp",
				"xmemory": "cpp",
				"xmemory0": "cpp",
				"xstddef": "cpp",
				"xstring": "cpp",
				"xtr1common": "cpp",
				"xtree": "cpp",
				"xutility": "cpp",
				"cctype": "cpp",
				"csignal": "cpp",
				"cstdarg": "cpp",
				"cwctype": "cpp",
				"regex": "cpp"
			}
		}`)

	return
}
