1. init函数:初始化函数,可以用来进行一些初始化的操作 每一个源文件都可以包含一个init函数,该函数会在main函数执行前,被Go运行框架调用。 
2. 全局变量定义, init函数, main函数的执行流程?
    - 全局,init,main
3. 多个源文件都有init函数的时候,如何执行:
    - 导入包中的先执行，接着main包中的全局变量，init,main函数
>![alt text](image-3.png)