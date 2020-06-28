// 1+...+5の和を求める
    @i
    M=1
    @sum
    M=0
(LOOP)
    // D=i
    @i
    D=M

    // D=i-5
    @5
    D=D-A

    // if (i-5)>0, Goto END
    @END
    D;JGT

    // D=i
    @i
    D=M

    // sum=sum+i
    @sum
    M=D+M

    // i=i+1
    @i
    M=M+1

    // Goto LOOP
    @LOOP
    0;JMP
(END)
    // Infinite loop
    @END
    0;JMP
