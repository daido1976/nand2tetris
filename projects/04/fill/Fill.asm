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
// 変数 start, end, current をセット（start, end はどこまで塗るかの判定、current は今どこを塗っているかの識別に使う）
    @SCREEN // 16384
    D=A
    @start // A = 16
    M=D
    @24575 // 16384（SCREEN）+ 8192 - 1
    D=A
    @end // A = 17
    M=D
    @SCREEN
    D=A
    @current // A = 18
    M=D
// キーボード操作待ちのループ処理
(LOOP)
    D=M[KBD] // キーボード入力されると KBD（24576）に該当キーの ASCII コードが入る（何も押されなければ 0 が入る）
    @WHITE // キーが押されていなければ WHITE へ
    D;JEQ
    @BLACK // // キーが押されていれば BLACK へ
(BLACK)
    // ガード節
    D=M[current]
    D=M[end]-D // end - current
    @LOOP
    D;JLT // end < current（最後まで黒で塗り潰し終わっている状態）ならば LOOP にジャンプする

    // 本処理
    A=M[current]
    M=-1 // 16 ビット分黒く塗り潰す
    @current // A = 17
    M=M+1 // current をインクリメント
    @LOOP
    0;JMP
(WHITE)
    // ガード節
    D=M[current]
    D=M[start]-D // start - current
    @LOOP
    D;JGT // start > current（最初まで白で塗り潰し終えた状態）ならば LOOP にジャンプする

    // 本処理
    A=M[current]
    M=0 // 16 ビット分白く塗り潰す
    @current // A = 17
    MD=M-1 // current をデクリメント
    @LOOP
    0;JMP
(END)
    @END
    0;JMP
