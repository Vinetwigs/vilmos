<div align="center">
    <h1>
        vilmos - Visual Interpreted Language to Make Obedient Serial art 🎨
    </h1>
</div>

## Main concepts
##### Idea
vilmos is an interpreted programming language that uses colors and pictures instead of code.
Its name is a tribute to the Hungarian painter Vilmos Huszár. Some of his most famous paintings rapresent some blocky shapes, just like the pixels needed to write vilmos programs.

##### Instructions
Each pixel's color, in a vilmos program, represents a single instruction (IMPORTANT: pixel's color must be perfectly equal to the istructions ones or custom ones, so the official interpreter accepts in input .png files with lossless compression).
The vilmos program image is read starting from the upper-left corner to the lower-right corner, but all the instructions can of course also be all on the same row.
Instruction set is strongly inspired by SuperStack! one.

From v2.0.0 it is possible to use bigger dimensions for the images and instructions, thanks to instruction_size (-s) interpreter's flag which specifies how many pixels rapresents a single instruction in the given image.
In this case the interpreter will execute only the upper-left pixel of each <strong>square</strong>. It is useful to create more appealing ~~images~~ programs.

###### Memory
vilmos is a stack-based language, so the memory is rapresented by a stack (a little bit "stronger" than a classic stack thanks to some powerful operations provided by the language out of the box) that has a non specified max storage capability so, while there is enough space on your device, you can continue putting things into the memory.
Using official interpreter, you can specify a maximum memory size for your program execution.

###### Errors handling
If your vilmos programs encounters an error during runtime, its interpreter will display the error, and the execution will be stopped. Hopefully you can run your software in debug mode to keep an eye on every single step of your painted code.

## Instructions
|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|INPUT_INT 	|Gets value from stdio as number and pushes it into the stack. If a file is opened,this instruction will read content from it and pushes all the characters in the file into the stack.   	|#ffffff   	| ![#ffffff](https://via.placeholder.com/25/ffffff/000000?text=+)  	|
|INPUT_ASCII   	|Gets values as ASCII char of a string and puts them into the stack. If a file is opened,this instruction will read content from it and pushes all the characters in the file into the stack.   	|#e3e3e3   	|![#e3e3e3](https://via.placeholder.com/25/e3e3e3/000000?text=+)|
|OUTPUT_INT   	|Pops the top of the stack and outputs it as number. If a file is opened,this instruction will write values into the file as integers and not in stdout.   	|#000001   	|![#000001](https://via.placeholder.com/25/000001/000000?text=+)   	|
|OUTPUT_ASCII   	|Pops the top of the stack and outputs it as ASCII char. If a file is opened,this instruction will write into the file as ASCII chars and not in stdout.   	|#4b4b4b   	|![#4b4b4b](https://via.placeholder.com/25/4b4b4b/000000?text=+)   	|
|SUM   	|Pops two numbers, adds them and pushes the result in the stack   	|#00ced1   	|![#00ced1](https://via.placeholder.com/25/00ced1/000000?text=+)   	|
|SUB   	|Pops two numbers, subtracts them and pushes the result in the stack   	|#ffa500   	|![#ffa500](https://via.placeholder.com/25/ffa500/000000?text=+)   	|
|DIV   	|Pops two numbers, divides them and pushes the result in the stack   	|#8a2be2   	|![#8a2be2](https://via.placeholder.com/25/8a2be2/000000?text=+)   	|
|MUL   	|Pops two numbers, multiplies them and pushes the result in the stack   	|#8b0000   	|![#8b0000](https://via.placeholder.com/25/8b0000/000000?text=+)   	|
|MOD   	|Pops two numbers, and pushes the result of the modulus in the stack   	|#ffdab9   	|![#ffdab9](https://via.placeholder.com/25/ffdab9/000000?text=+)   	|
|RND   	|Pops one number, and pushes in the stack a random number between [0, n[ where n is the number popped   	|#008000   	|![#008000](https://via.placeholder.com/25/008000/000000?text=+)   	|
|AND   	|Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ecf3dc   	|![#ecf3dc](https://via.placeholder.com/25/ecf3dc/000000?text=+)   	|
|OR   	|Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#b7c6e6   	|![#b7c6e6](https://via.placeholder.com/25/b7c6e6/000000?text=+)   	|
|XOR   	|Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#f5e3d7   	|![#f5e3d7](https://via.placeholder.com/25/f5e3d7/000000?text=+)   	|
|NAND   	|Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#e1d3ef   	|![#e1d3ef](https://via.placeholder.com/25/e1d3ef/000000?text=+)   	|
|NOT   	|Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ff9aa2   	|![#ff9aa2](https://via.placeholder.com/25/ff9aa2/000000?text=+)   	|
|POP   	|Pops one element, and discardes it   	|#cc9e06   	|![#cc9e06](https://via.placeholder.com/25/cc9e06/000000?text=+)   	|
|SWAP   	|Swaps the top two elements in the stack   	|#ffbd4a   	|![#ffbd4a](https://via.placeholder.com/25/ffbd4a/000000?text=+)   	|
|CYCLE   	|Cycles clockwise the stack of one position   	|#e37f9d   	|![#e37f9d](https://via.placeholder.com/25/e37f9d/000000?text=+)   	|
|RCYCLE   	|Cycles counterclockwise the stack of one position   	|#e994ae   	|![#e994ae](https://via.placeholder.com/25/e994ae/000000?text=+)   	|
|DUP   	|Duplicates the top of the stack   	|#006994   	|![#006994](https://via.placeholder.com/25/006994/000000?text=+)   	|
|REVERSE   	|Reverses the content of the stack   	|#a5a58d   	|![#a5a58d](https://via.placeholder.com/25/a5a58d/000000?text=+)   	|
|QUIT   	|Terminates program execution   	|#b7e4c7   	|![#b7e4c7](https://via.placeholder.com/25/b7e4c7/000000?text=+)   	|
|WHILE   	|Enters in a while loop: if the top element is true loop, else exits while loop. It doesn't pop the element.   	|#2e1a47   	|![#2e1a47](https://via.placeholder.com/25/2e1a47/000000?text=+)   	|
|WHILE_END   	|Ends while loop   	|#68478d   	|![#68478d](https://via.placeholder.com/25/68478d/000000?text=+)   	|
|FILE_OPEN   	|Opens the file using the last string in the stack as path. The string is popped from the stack. While the file is open, INPUT_ASCII and INPUT_INT instructions will read all the file content and push each char into the stack. While the file is open, OUTPUT_INT and OUTPUT_ASCII instructions will write into the file. If the file doesn't exists, it will be created. An opened file is in read-write append mode. Only one file can be opened at a time.    	|#91f68b   	|![#91f68b](https://via.placeholder.com/25/91f68b/000000?text=+)   	|
|FILE_CLOSE   	|Closes the currently opened file. INPUT_ASCII, OUTPUT_ASCII, INPUT_INT and INPUT_ASCII will return to their standard behaviour.   	|#2fed23   	|![#2fed23](https://via.placeholder.com/25/2fed23/000000?text=+)   	|

## Insert data from the image
In vilmos language it is possible to store integers and strings (depending on the output instruction used) into the stack directly from the image.
<strong>For each pixel into the image that has a color code different from all the operations' color codes</strong> (or custom ones if defined), <strong>the interpreter pushes into the stack a value equals to the sum of the pixel's red, green and blue values</strong>
#### Strings
To insert into the stack a string, you have to push at the beginning of the string a delimeter rapresented by RGB value 000000 (\000). When reading a string, if the delimeter is not found, an error will be launched and the execution will be stopped. Thanks to this delimiter you can have into the stack multiple strings and integers at the same time. You can also insert a single character, but <strong>remember to insert the delimiter.</strong>

:exclamation: <strong>For this feature there is no compatibility for programs written before v2.0.0</strong> :exclamation:

## Examples
You can find into the main repository an 'examples' folder containing some basics programs written in vilmos language. Of course it is possible to write programs a lot more complex than those one.
The only limit is your palette!

>Have fun ~~painting~~ coding