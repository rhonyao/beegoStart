# Beego Start
开始学习beego，你可以用这份代码快速进行分析和开发.
这不是一份完全restful的Api接口实例，但是通过修改路由中的post,get你可以将这份代码在10分钟内快速迭代为restful api。
使用这份代码，里面包含了一些控制器的例子，他们分别可以完整的用于：
    </br>支付宝支付
    </br>github登录
    </br>阿里云oss存储
    </br>支付定订单
    </br>253.com短信验证码
    </br>用户(推荐重构这个控制器)
    </br> 关于用户登录的这个控制器，只是简易的用户登录，我会新建一个git做一个完整强大独立的golang版本Oauth2的server。后续开源在github并将此框架融入。目前已经在开发中。将让你获得一个类似QQ登录的授权系统。

### 使用代码，你需要知道的技术栈
> Web framework Beego
> Vue, Pug, ECMA Script 2017
> Nuxt.js

### Nginx的反向代理配置

>for api endpoint

您需要在您的电脑安装nginx，并使用nginx进行反向代理来使用您的业务。

```
upstream dev{
    server localhost:1993 max_fails=3 fail_timeout=3s;
}

server {  
  listen 80;  
  server_name  dev.endopint.com;  

    location / {
        proxy_redirect off ;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://dev;
    }
}

```

### 启动

你需要在你的操作系统安装Golang和beego。
```
    cd /beegoStart
    bee run
```

为Linux打包
``` sh
    bee pack -be GOOS=linux -ba "-tags prod" -exr="^(?:front|vendor|tests|oauth|uploads|utils)$"
```

推荐使用supervisor做进程守护。

### 数据库
你需要在你的机器安装MySQL, Redis
MySQL数据表将在orm启动时自动写入您的数据库，无需提前导入.sql文件。

### 软件配置
现在还没有发布完整的可用于生产环境的版本。在此之前发送邮件到virskor@gmail.com获取所有的配置文件。我们会告诉你配置文件的作用。

### 前端启动

你的电脑需要安装Node或者你可以使用cnpm替代npm。
```sh
    cd /front
```

1. install
```sh
$ npm install
```
2. development
```sh
$ npm run dev
```
3. build
```sh
$ npm run prod
```