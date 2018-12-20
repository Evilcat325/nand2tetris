// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
//init screen_size
@screen_size
D=1
M=D+1	//screen_size = 2
D=M
M=M+D	//screen_size = 4
D=M
M=M+D	//screen_size = 8
D=M
M=M+D	//screen_size = 16
D=M
M=M+D	//screen_size = 32
D=M
M=M+D	//screen_size = 64
D=M
M=M+D	//screen_size = 128
D=M
M=M+D	//screen_size = 256
D=M
M=M+D	//screen_size = 512
D=M
M=M+D	//screen_size = 1024
D=M
M=M+D	//screen_size = 2048
D=M
M=M+D	//screen_size = 4096
D=M
M=M+D	//screen_size = 8192
M=M-1	//screen_size = 8191

//init stat to 0
@last_state
M=0

(START)
//read from kbd
@KBD
D=M
//if kbd != 0 then key_board_state = 1
@SKIP_IF_0
D;JEQ
	D=1
(SKIP_IF_0)
@key_board_state
M=D

//if last_state == key_board_state skip
@last_state
D=D-M
@START
D;JEQ

@i
M=0	 	//init i
@SCREEN
D=A
@screen_pointer
M=D

(LOOP)
	@i
	D=M
	M=M+1
	@screen_size
	D=D-M
	@LOOP_END
	D;JGT

	//loop start
	@key_board_state
	D=M
	@SKIP_IF_0_tmp
	D;JEQ
		D=-1
	(SKIP_IF_0_tmp)

	@screen_pointer
	A=M
	M=D
	@screen_pointer
	M=M+1

@LOOP
0;JMP
(LOOP_END)

//update last_state
@key_board_state
D=M
@last_state
M=D
@START
0;JMP
