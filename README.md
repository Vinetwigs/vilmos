
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Vinetwigs/vilmos)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Vinetwigs/vilmos)
![GitHub last commit](https://img.shields.io/github/last-commit/Vinetwigs/vilmos)
[![stars - vilmos](https://img.shields.io/github/stars/Vinetwigs/vilmos?style=social)](https://github.com/Vinetwigs/vilmos)
[![forks - vilmos](https://img.shields.io/github/forks/Vinetwigs/vilmos?style=social)](https://github.com/Vinetwigs/vilmos)
[![License](https://img.shields.io/badge/License-Apache_License_2.0-orange)](#license)
[![issues - vilmos](https://img.shields.io/github/issues/Vinetwigs/vilmos)](https://github.com/Vinetwigs/vilmos/issues)

<div>
   <h1 align="center">
      vilmos interpreter - let's put software in museums🎨
   </h1>
</div>

<div align="center">
   Uninstall all your IDE's, close the terminal, install your favourite drawing software and start programming today using only your favourite colors.
   Create your custom palettes, take your brushes and paint your softwares using <a href="./LANGUAGE.md">vilmos language</a>.
   
   <strong>:exclamation: Brushes not included in the kit :exclamation:</strong>
</div>

## Prerequisites
Make sure you have at least installed Go v1.17 or your paintings may be blue :sob:

## Installing
```
go install github.com/Vinetwigs/vilmos/v2
```
Now you can start using the interpreter via terminal using 'vilmos' command

## How to use
###### Getting help
`vilmos help` prints the complete usage of vilmos interpreter. You can optionally use `vilmos -h` as short form or `vilmos --help`.

```
NAME:
   vilmos - Official vilmos language interpreter

USAGE:
   vilmos.exe [global options] command [command options] [arguments...]

VERSION:
   2.0.1

AUTHOR:
   Vinetwigs <github.com/Vinetwigs>

COMMANDS:
   version, v  show installed version
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                                                   enable debug mode (default: false)
   --config FILE_PATH, --conf FILE_PATH, -c FILE_PATH            load configuration from FILE_PATH for custom color codes
   --max_size value, -m value                                    set max memory size (default: -1)
   --instruction_size value, --is value, --size value, -s value  set instruction size (default: 1)
   --help, -h                                                    show help (default: false)
   --version, -V, -v                                             Shows installed version (default: false)                        Shows installed version (default: false)
```

###### Interpret program
`vilmos <FILE_PATH>` is the easiest way to see your colors come to life. Make sure your image is in .png format.

###### Debug program
vilmos debugger lets you execute your programs step by step printing in each step the content of the stack.
To enable debugger mode all you need to do is putting -d flag in interpreter terminal execution.
`vilmos -d <FILE_PATH>`.
There is also a longer version for the flag: `vilmos --debug <FILE_PATH>`.

###### Use custom color codes
The true power of vilmos visual language is the capability of setting custom color codes for your programs.
To achieve that, you have to simply specify in a config file (example config file provided in the repository) a color code for the instruction.
The color codes must be in the HEX format (without the #) and there are two supported formats:

1. Full Hex code (a044d1)
2. Short Hex code (fff)

Once you have chosen your favourite colors, your program execution must be in the format:
`vilmos -c <CONFIG_FILE_PATH> <FILE_PATH>`.
You can optionally use longer versions for config flag:
`vilmos --conf <CONFIG_FILE_PATH> <FILE_PATH>` or `vilmos --config <CONFIG_FILE_PATH> <FILE_PATH>`.

###### Set instructions size
To specify the pixel size of each instruction in your colorful vilmos program, there is -instruction_size flag to help you.
`vilmos -s <size> <FILE_PATH>`. You can optionally use the alternative forms:
1. `vilmos -instruction_size <size> <FILE_PATH>`
2. `vilmos -size <size> <FILE_PATH>`

When you specify the instructions size, make sure that each instruction in your vilmos program is a square having the specified size as sides length. This permits to use bigger images to make your vilmos program more appealing.
Read [language specifications](https://github.com/Vinetwigs/vilmos/blob/main/LANGUAGE.md) for more informations about 

###### Set maximum memory size
To specify a maximum size for the memory usable for your painting execution you have to use -m flag.
`vilmos -m <size> <FILE_PATH>`.
You can optionally use the longer version: `vilmos --max_size <size> <FILE_PATH>`.
If you try to put in memory another element when the stack is full, an error will be launched and the execution will be stopped.

###### Version
To print actual vilmos interpreter version you have different choices:
1. `vilmos version`
2. `vilmos --version`
3. `vilmos -V`
4. `vilmos -v`

## Author
- [Vinetwigs](https://github.com/Vinetwigs)

## Contributors
<a href="https://github.com/Vinetwigs/vilmos/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Vinetwigs/vilmos" />
</a>

## License
```text
   Copyright 2021 Vinetwigs

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```
