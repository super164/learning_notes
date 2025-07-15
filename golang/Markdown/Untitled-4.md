Markdown代码
===
行内代码
---
>如果是段落上的一个函数或片段的代码可以用反引号把它包起来（`）
- 函数名：使用`console.log()`输出信息
- 变量名：将值赋值给`userName`变量
- 命令行：运行`npm install`安全依赖
- 键盘按键：按`Ctrl+C`复制内容
- 文件名：编辑`index.html`文件

特殊字符转义
---
使用双反引号包围单反引号：

``使用 `反引号` 包围代码``

使用多个反引导号包围：
```包含 `` 双反引号的代码```

代码区块
---
正常文本段落

    这是缩进式代码块
    每行前面有四个空格
    保持代码的原始格式
    
继续正常文本

三反引号代码块
---
>可以用 ``` 包裹一段代码，并指定一种语言（也可以不指定）
```javascript
$(document).ready(function () {
    alert('RUNOOB');
});
```

代码的高级特性
===
行号显示
---
```javascript {.line-numbers}
function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

console.log(fibonacci(10));
```

代码差异对比
---
Diff 语法：

```diff
function calculateTotal(items) {
-   let total = 0;
+   let total = 0.0;
    
    for (let item of items) {
-       total += item.price;
+       total += parseFloat(item.price);
    }
    
+   // 保留两位小数
+   total = Math.round(total * 100) / 100;
    return total;
}
```

语言特定的差异对比
---

```javascript
// 之前的代码
const oldFunction = () => {
    var x = 10;  // &#x274c; 使用 var
    console.log("Value: " + x);  // &#x274c; 字符串拼接
}

// 改进后的代码  
const newFunction = () => {
    const x = 10;  // &#x2705; 使用 const
    console.log(`Value: ${x}`);  // &#x2705; 模板字符串
}
```