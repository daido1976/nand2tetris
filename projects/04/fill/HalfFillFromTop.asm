// スクリーンの上半分を上から順番に塗り潰す
    @20480 // 16384（SCREEN） + 4096（半分）
    D=A
    @end // A = 16
    M=D // end（アドレスは16） = 20480
    @SCREEN
    D=A
    @current // A = 17
    M=D // current（アドレスは17） = 16384（スタート地点）
(LOOP)
    A=M[current] // A = 16384（2回目以降は 16385, 16386... とインクリメントしていく）
    M=-1 // 16 ビット分黒く塗り潰す
    @current // A = 17
    MD=M+1 // current をインクリメント
    D=M[end]-D // end - current
    @LOOP
    D;JGT // end - current が 0 より大きければ LOOP する（つまり end > current ならば LOOP する）
(END)
    @END
    0;JMP
