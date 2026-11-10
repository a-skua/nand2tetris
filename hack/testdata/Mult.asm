// RAM[2] = RAM[0] * RAM[1] (RAM[0] >= 0)
// RAM[0] is used as the loop counter and becomes 0.
    @R2
    M=0
(LOOP)
    @R0
    D=M
    @END
    D;JEQ
    @R1
    D=M
    @R2
    M=D+M
    @R0
    M=M-1
    @LOOP
    0;JMP
(END)
    @END
    0;JMP
