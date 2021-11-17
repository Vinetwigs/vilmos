<div align="center">
    <h1 align="center">
        vilmos Language Specification
    </h1>
</div>

## Table of Contents

1. [Introduction](#introduction)
    1. [Idea](#idea)
    2. [Instructions](#instructions)
    3. [Types](#types)
    4. [Memory](#memory)
    5. [Errors](#errors)
2. [Instructions list](#instructions-list)
    1. [I/O](#io)
    2. [Arithmetic operators](#arithmetic-operators)
    3. [Logical operators](#logical-operators)
    4. [Bitwise operators](#bitwise-operators)
    5. [Stack operations](#stack-operations)
    6. [Control flow](#control-flow)
    7. [File management](#file-management)
    8. [Miscellaneous](#miscellaneous)
3. [Insert data in memory](#insert-data-in-memory)

## Introduction

### Idea

vilmos is an interpreted programming language that uses colors and pictures instead of code.
Its name is a tribute to the Hungarian painter Vilmos Huszár. Some of his most famous paintings    
rapresent blocky shapes, just like the pixels/squares needed to write vilmos programs.

<div>
    <div>
        <img src=".\assets\composition.jpg" height="300px">
    </div>
    <i>Composition</i> - Vilmos Huszár
</div>

[Back to top](#table-of-contents)

### Instructions

In a vilmos program, instructions are rapresented by squares of the same dimensions.
The minimum acceptable dimension is a single pixel. There is not a maximum dimension limit.    

The instructions are executed starting from the upper-left corner to the lower-right corner, but of course they can also   
be all on the same row.


<strong>NOTICE:</strong> An instruction, in vilmos, must match perfectly with the relative color code
so a program must be a <strong>.png</strong> image. In this way there will be no quality loss.

Instruction set is strongly inspired by [SuperStack!](https://esolangs.org/wiki/Super_Stack!#Instructions) one.

<i>Execution order example</i>


![exec-order-alt](.\assets\execution_order.jpg)    

[Back to top](#table-of-contents)

### Types

_vilmos_ supports following two data types:

 * **int**: a 32-bit signed integer [_-2147483648 to 2147483647_]
 * **string**: a sequence of ASCII characters with _**\0 delimiter at the beginning**_ of the string

### Memory

vilmos is a stack-based language, so the memory is rapresented by a stack (a little bit "stronger" than a classic    
one thanks to some powerful and useful operations provided by the language out of the box).

By default, memory has no maximum limit (it only depends to your device memory).
This can be changed setting a maximum stack size while using official interpreter.

[Back to top](#table-of-contents)

### Errors

If your vilmos painting encounters an error during runtime, the execution will be immediately stopped.

If you are using the official interpreter, when the execution is stopped, it will also be displayed an error   
message that describes what happened.
This can be avoided by using the debugger tool provided out of the box by the interpreter.

[Back to top](#table-of-contents)

## Instructions List

### I/O

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|INPUT_INT 	|Gets value from stdio as number and pushes it into the stack. If a file is opened,this instruction will read content from it and pushes all the characters in the file into the stack.   	|#ffffff   	| ![#ffffff](https://via.placeholder.com/25/ffffff/000000?text=+)  	|
|INPUT_ASCII   	|Gets values as ASCII char of a string and puts them into the stack. If a file is opened,this instruction will read content from it and pushes all the characters in the file into the stack.   	|#e3e3e3   	|![#e3e3e3](https://via.placeholder.com/25/e3e3e3/000000?text=+)|
|OUTPUT_INT   	|Pops the top of the stack and outputs it as number. If a file is opened,this instruction will write values into the file as integers and not in stdout.   	|#000001   	|![#000001](https://via.placeholder.com/25/000001/000000?text=+)   	|
|OUTPUT_ASCII   	|Pops the top of the stack and outputs it as ASCII char. If a file is opened,this instruction will write into the file as ASCII chars and not in stdout.   	|#4b4b4b   	|![#4b4b4b](https://via.placeholder.com/25/4b4b4b/000000?text=+)   	|

[Back to top](#table-of-contents)

### Arithmetic operators

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|SUM   	|Pops two numbers, adds them and pushes the result in the stack   	|#00ced1   	|![#00ced1](https://via.placeholder.com/25/00ced1/000000?text=+)   	|
|SUB   	|Pops two numbers, subtracts them and pushes the result in the stack   	|#ffa500   	|![#ffa500](https://via.placeholder.com/25/ffa500/000000?text=+)   	|
|DIV   	|Pops two numbers, divides them and pushes the result in the stack   	|#8a2be2   	|![#8a2be2](https://via.placeholder.com/25/8a2be2/000000?text=+)   	|
|MUL   	|Pops two numbers, multiplies them and pushes the result in the stack   	|#8b0000   	|![#8b0000](https://via.placeholder.com/25/8b0000/000000?text=+)   	|
|MOD   	|Pops two numbers, and pushes the result of the modulus in the stack   	|#ffdab9   	|![#ffdab9](https://via.placeholder.com/25/ffdab9/000000?text=+)   	|

[Back to top](#table-of-contents)

### Logical operators

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|AND   	|Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ecf3dc   	|![#ecf3dc](https://via.placeholder.com/25/ecf3dc/000000?text=+)   	|
|OR   	|Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#b7c6e6   	|![#b7c6e6](https://via.placeholder.com/25/b7c6e6/000000?text=+)   	|
|XOR   	|Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#f5e3d7   	|![#f5e3d7](https://via.placeholder.com/25/f5e3d7/000000?text=+)   	|
|NAND   	|Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#e1d3ef   	|![#e1d3ef](https://via.placeholder.com/25/e1d3ef/000000?text=+)   	|
|NOT   	|Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ff9aa2   	|![#ff9aa2](https://via.placeholder.com/25/ff9aa2/000000?text=+)   	|

[Back to top](#table-of-contents)

### Bitwise operators

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|BAND   	|Pops two numbers, and pushes the result of their bitwise AND |#8aa399   	|![#8aa399](https://via.placeholder.com/25/8aa399/000000?text=+)   	|
|BOR   	|Pops two numbers, and pushes the result of their bitwise OR  |#7d84b2   	|![#7d84b2](https://via.placeholder.com/25/7d84b2/000000?text=+)   	|
|BXOR   	|Pops two numbers, and pushes the result of their bitwise XOR |#8fa6cb   	|![#8fa6cb](https://via.placeholder.com/25/8fa6cb/000000?text=+)   	|
|BNOT   	|Pops one number, and pushes the result of its bitwise NOT |#dbf4a7   	|![#dbf4a7](https://via.placeholder.com/25/dbf4a7/000000?text=+)   	|
|RSHIFT   	|Pops two numbers, and pushes the result of the second >> the first |#439dba   	|![#439dba](https://via.placeholder.com/25/439dba/000000?text=+)   	|
|LSHIFT   	|Pops two numbers, and pushes the result of the second << the first |#2d6a7d   	|![#2d6a7d](https://via.placeholder.com/25/2d6a7d/000000?text=+)   	|

[Back to top](#table-of-contents)

### Stack operations

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|POP   	|Pops one element, and discardes it   	|#cc9e06   	|![#cc9e06](https://via.placeholder.com/25/cc9e06/000000?text=+)   	|
|SWAP   	|Swaps the top two elements in the stack   	|#ffbd4a   	|![#ffbd4a](https://via.placeholder.com/25/ffbd4a/000000?text=+)   	|
|CYCLE   	|Cycles clockwise the stack of one position   	|#e37f9d   	|![#e37f9d](https://via.placeholder.com/25/e37f9d/000000?text=+)   	|
|RCYCLE   	|Cycles counterclockwise the stack of one position   	|#e994ae   	|![#e994ae](https://via.placeholder.com/25/e994ae/000000?text=+)   	|
|DUP   	|Duplicates the top of the stack   	|#006994   	|![#006994](https://via.placeholder.com/25/006994/000000?text=+)   	|
|REVERSE   	|Reverses the content of the stack   	|#a5a58d   	|![#a5a58d](https://via.placeholder.com/25/a5a58d/000000?text=+)   	|

[Back to top](#table-of-contents)

### Control flow

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|WHILE   	|Enters in a while loop: if the top element is true loop, else exits while loop. It doesn't pop the element.   	|#2e1a47   	|![#2e1a47](https://via.placeholder.com/25/2e1a47/000000?text=+)   	|
|WHILE_END   	|Ends while loop   	|#68478d   	|![#68478d](https://via.placeholder.com/25/68478d/000000?text=+)   	|
|QUIT   	|Terminates program execution   	|#b7e4c7   	|![#b7e4c7](https://via.placeholder.com/25/b7e4c7/000000?text=+)   	|

[Back to top](#table-of-contents)

### File management

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|FILE_OPEN   	|Opens the file using the last string in the stack as path. The string is popped from the stack. While the file is open, INPUT_ASCII and INPUT_INT instructions will read all the file content and push each char into the stack. While the file is open, OUTPUT_INT and OUTPUT_ASCII instructions will write into the file. If the file doesn't exists, it will be created. An opened file is in read-write append mode. Only one file can be opened at a time.    	|#91f68b   	|![#91f68b](https://via.placeholder.com/25/91f68b/000000?text=+)   	|
|FILE_CLOSE   	|Closes the currently opened file. INPUT_ASCII, OUTPUT_ASCII, INPUT_INT and INPUT_ASCII will return to their standard behaviour.   	|#2fed23   	|![#2fed23](https://via.placeholder.com/25/2fed23/000000?text=+)   	|

[Back to top](#table-of-contents)

### Miscellaneous

|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|RND   	|Pops one number, and pushes in the stack a random number between [0, n[ where n is the number popped   	|#008000   	|![#008000](https://via.placeholder.com/25/008000/000000?text=+)   	|

[Back to top](#table-of-contents)

## Insert data in memory

vilmos supports the possibility to insert into the stack integers and strings directly from the given image.

When an operation square is encountered and its color code doesn't match with any of the default/custom color codes    
defined, the sum of red, green and yellow operation square's value is pushed into the memory.

**NOTICE:** When pushing a string into the stack, remember to insert at the beginning of the string the \0 delimiter (rapresented by #000000 RGB value)
Thanks to this delimiter it is possible to have in memory multiple strings and integers at the same time.
It is also possible to insert a single char, but remember to insert the delimeter!

_Example program that inserts 100 in memory and outputs it:_

![insert-int-alt](./assets/insert_int.png)

_Example program that inserts 'vilmos' string in memory and outputs it:_

![insert-string-alt](./assets/insert_string.png)

>Have fun ~~painting~~ coding

[Back to top](#table-of-contents)