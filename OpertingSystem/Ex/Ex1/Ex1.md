- 实验步骤
    1. 实验步骤
        - 从软盘加载驱动程序到`0x7C00`
        ```
        org 0x7C00
        jmp short Start
        nop
        ``` 
        - 