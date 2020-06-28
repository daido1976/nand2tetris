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
    @20480 // 16384（SCREEN） + 4096（半分）
    D=A
    @end // A = 16
    M=D
    @SCREEN
    D=A
    @current // A = 17
    M=D
(LOOP)
    D=M[KBD]
    @LOOP
    D;JEQ
    @BLACK
(BLACK)
    A=M[current]
    M=-1 // 16 ビット分黒く塗り潰す
    @current // A = 17
    MD=M+1 // current をインクリメント
    D=M[end]-D // end - current
    @BLACK
    D;JGT // end - current が 0 より大きければ LOOP する（つまり end > current ならば LOOP する）
(END)
    @END
    0;JMP
