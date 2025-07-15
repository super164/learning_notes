# Markdown 高级技巧
## 支持的HTML元素
>不在 Markdown 涵盖范围之内的标签，都可以直接在文档里面用 HTML 撰写。
目前支持的 HTML 元素有：`<kbd> <b> <i> <em> <sup> <sub> <br> `

使用 <kbd>Ctrl</kbd>+<kbd>Alt</kbd>+<kbd>Del</kbd> 重启电脑

## 转义
>需要使用`\`来实现转义

**文本加粗** 
\*\* 正常显示星号 \*\*

## 数学公式
#### 基本语法结构
>命令：以反斜杠 \ 开头，如 \alpha、\sum
参数：用花括号 {} 包围，如 \frac{a}{b}
下标：使用 _，如 x_1
上标：使用 ^，如 x^2
分组：用花括号将多个字符组合，如 x_{i+1}

#### 常用命令
>\alpha, \beta, \gamma  % 希腊字母
\sum, \prod, \int      % 求和、乘积、积分
\frac{分子}{分母}      % 分数
\sqrt{表达式}          % 平方根
\sqrt[n]{表达式}       % n次根
