# Linux命令基础格式

无论是什么命令,用于什么用途,在Linux中,命令有其通用的格式: 

command [-options] [parameter] 

- command: 命令本身 
- options: [可选，非必填]命令的一些选项，可以通过选项控制命令的行为细节
- parameter: [可选,非必填]命令的参数,多数用于命令的指向目标等
- 语法中的[]，表示可选的意思 





示例：
ls-l/home/itheima, ls是命令本身，-l是选项， /home/itheima是参数 

意思是以列表的形式，显示/home/itheima目录内的内容 

cp -r test1 test2，cp是命令本身，-r 是选项，test1 和test2是参数

意思是复制文件夹test1成为test2