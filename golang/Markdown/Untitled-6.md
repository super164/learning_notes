# Markdown 图片
![替代文字](图片路径)
![替代文字](图片路径 "图片标题")
>开头一个感叹号 !
接着一个方括号，里面放上图片的替代文字
接着一个普通括号，里面放上图片的网址，最后还可以用引号包住并加上选择性的 'title' 属性的文字。

- 相对路径：  
![项目截图](./images/screenshot.png)
![用户界面](../assets/ui-demo.jpg "用户界面演示")
![图标](images/icon.svg "应用图标")
- 绝对路径：  
![本地图片](/Users/username/Documents/image.png)
![系统截图](C:\Users\username\Pictures\screenshot.png)

## 直接引用网络图片
![RUNOOB 图标](https://static.jyshare.com/images/runoob-logo.png)

![RUNOOB 图标](https://static.jyshare.com/images/runoob-logo.png "RUNOOB")
## 对图片地址使用变量
这个链接用 1 作为网址变量 [RUNOOB][1].
然后在文档的结尾为变量赋值（网址）

[1]: https://static.jyshare.com/images/runoob-logo.png

## 图片alt文本
![苹果公司总部大楼外观，现代玻璃幕墙建筑](./images/apple-headquarters.jpg)
![网站流量统计图表，显示过去六个月的访问量呈上升趋势](./charts/traffic-stats.png)
![用户登录界面，包含用户名和密码输入框](./screenshots/login-page.png)

## 链接和图片的高级用法
### 图片链接组合
[![图片alt文本](图片URL)](链接URL)