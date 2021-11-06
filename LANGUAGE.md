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
###### Memory
vilmos is a stack-based language, so the memory is rapresented by a stack (a little bit "stronger" than a classic stack thanks to some powerful operations provided by the language out of the box) that has a non specified max storage capability so, while there is enough space on your device, you can continue putting things into the memory.
Using official interpreter, you can specify a maximum memory size for your program execution.
###### Errors handling
If your vilmos programs encounters an error during runtime, its interpreter will display the error, and the execution will be stopped. Hopefully you can run your software in debug mode to keep an eye on every single step of your painted code.

## Instructions
|  Instruction 	| Description  	| Color code   	| Color preview   	|
|:-:	|:-:	|:-:	|:-:	|
|INPUT_INT 	|Gets value from stdio as number and pushes it into the stack   	|#ffffff   	| <div style="width:20px; height:20px; background: #fff; border: 1px solid black"></div>  	|
|INPUT_ASCII   	|Gets values as ASCII char of a string and puts them into the stack   	|#e3e3e3   	|<div style="width:20px; height:20px; background: #e3e3e3;"></div>|
|OUTPUT_INT   	|Pops the top of the stack and outputs it as number   	|#000000   	|<div style="width:20px; height:20px; background: #000000;"></div>   	|
|OUTPUT_ASCII   	|Pops the top of the stack and outputs it as ASCII char   	|#4b4b4b   	|<div style="width:20px; height:20px; background: #4b4b4b;"></div>   	|
|SUM   	|Pops two numbers, adds them and pushes the result in the stack   	|#00ced1   	|<div style="width:20px; height:20px; background: #00ced1;"></div>   	|
|SUB   	|Pops two numbers, subtracts them and pushes the result in the stack   	|#ffa500   	|<div style="width:20px; height:20px; background: #ffa500;"></div>   	|
|DIV   	|Pops two numbers, divides them and pushes the result in the stack   	|#8a2be2   	|<div style="width:20px; height:20px; background: #8a2be2;"></div>   	|
|MUL   	|Pops two numbers, multiplies them and pushes the result in the stack   	|#8b0000   	|<div style="width:20px; height:20px; background: #8b0000;"></div>   	|
|MOD   	|Pops two numbers, and pushes the result of the modulus in the stack   	|#ffdab9   	|<div style="width:20px; height:20px; background: #ffdab9;"></div>   	|
|RND   	|Pops one number, and pushes in the stack a random number between [0, n[ where n is the number popped   	|#008000   	|<div style="width:20px; height:20px; background: #008000;"></div>   	|
|AND   	|Pops two numbers, and pushes the result of AND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ecf3dc   	|<div style="width:20px; height:20px; background: #ecf3dc;"></div>   	|
|OR   	|Pops two numbers, and pushes the result of OR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#b7c6e6   	|<div style="width:20px; height:20px; background: #b7c6e6;"></div>   	|
|XOR   	|Pops two numbers, and pushes the result of XOR [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#f5e3d7   	|<div style="width:20px; height:20px; background: #f5e3d7;"></div>   	|
|NAND   	|Pops two numbers, and pushes the result of NAND [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#e1d3ef   	|<div style="width:20px; height:20px; background: #e1d3ef;"></div>   	|
|NOT   	|Pops one number, and pushes the result of NOT [0 is false, anything else is true] [pushes 1 if true or 0 is false]   	|#ff9aa2   	|<div style="width:20px; height:20px; background: #ff9aa2;"></div>   	|
|POP   	|Pops one element, and discardes it   	|#cc9e06   	|<div style="width:20px; height:20px; background: #cc9e06;"></div>   	|
|SWAP   	|Swaps the top two elements in the stack   	|#ffbd4a   	|<div style="width:20px; height:20px; background: #ffbd4a;"></div>   	|
|CYCLE   	|Cycles clockwise the stack of one position   	|#e37f9d   	|<div style="width:20px; height:20px; background: #e37f9d;"></div>   	|
|RCYCLE   	|Cycles counterclockwise the stack of one position   	|#e994ae   	|<div style="width:20px; height:20px; background: #e994ae;"></div>   	|
|DUP   	|Duplicates the top of the stack   	|#006994   	|<div style="width:20px; height:20px; background: #006994;"></div>   	|
|REVERSE   	|Reverses the content of the stack   	|#a5a58d   	|<div style="width:20px; height:20px; background: #a5a58d;"></div>   	|
|QUIT   	|Terminates program execution   	|#b7e4c7   	|<div style="width:20px; height:20px; background: #b7e4c7;"></div>   	|
|WHILE   	|Enters in a while loop: if the top element is true loop, else exits while loop. It doesn't pop the element.   	|#2e1a47   	|<div style="width:20px; height:20px; background: #2e1a47;"></div>   	|
|WHILE_END   	|Ends while loop   	|#68478d   	|<div style="width:20px; height:20px; background: #68478d;"></div>   	|

## Insert data from the image
In vilmos language it is possible to store integers and chars (depending on the output instruction used) into the stack directly from the image.
<strong>For each pixel into the image that has a color code different from all the operations' color codes</strong> (or custom ones if defined), <strong>the interpreter pushes into the stack a value equals to the sum of the pixel's red, green and blue values</strong>

## Examples
You can find into the main repository an 'examples' folder containing some basics programs written in vilmos language. Of course it is possible to write programs a lot more complex than those one.
The only limit is your palette!

>Have fun ~~painting~~ coding