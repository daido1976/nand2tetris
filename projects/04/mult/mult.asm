// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

// R2 = R0 * R1 の乗算プログラム（R0 と R1 にはあらかじめ 0 以上の値が入力されている前提）
// R1 の数だけ R2 += R0 している。（R1 が 0 の場合は何もせず 0 を返す）
// R1 をインクリメンタルな i （元は 1）と比較することで処理終了するかどうかを決めている。（D;JGT の部分）
    @R2
    M=0 // R2 = 0
    @i
    M=1 // i = 1
(LOOP)
    @i
    D=M // D = i
    @R1
    D=D-M // D = i - R1
    @END
    D;JGT // if D > 0 then go to END
    @R0
    D=M // D = R0
    @R2
    M=D+M // R2 += R0
    @i
    M=M+1 // i++
    @LOOP
    0;JMP // Goto LOOP
(END)
    @END
    0;JMP // Infinite loop
