# 阿里大鱼api接口

==================

*沙箱环境存在问题 因此此功能闲置*

调用方法

引用包

```go
import(
  "alidayu"
  )
```

初始化函数


```go
// 初始化阿里云

alidayu.Appkey = "xxxxxxx"

alidayu.AppSecret = "xxxxxxxxxxxxxxxxxxxxxxxxx"

alidayu.IsDebug = false
```

## 短信接口

* SendOnce 短信单条发送的接口

>// moblie-手机号码

>// signname-短信签名

>// templatecode-短信模板

>// param-传入参数

>// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错

```go
alidayu.SendOnce("136xxxxxxx8","测试",“SM_777777”,"{'code':'666666'}")
```

##  

* SendBatch 短信单条发送的接口

>// moblie-手机号码 ','隔开

>// signname-短信签名

>// templatecode-短信模板

>// param-传入参数

>// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错

```go
alidayu.SendBatch("136xxxxxxx8,136xxxxxxx3","测试","SM_777777","{'code':'666666'}")
```

## 文本转语音接口

* SendLecall 文本转语音的接口

>// moblie-手机号码

>// templatecode-短信模板

>// param-传入参数

>// 返回Result格式，请确保输出成功失败的结构体引用，如果Result.Success为true则获取Result.ResultError报错

```go
alidayu.SendLecall("136xxxxxxx8","TL_777777","{'code':‘666666’}")
```


## 返回格式

返回格式默认为model.go下的Result格式

在引用时不能引用为空的子结构

例如Result.Success==true时

引用到Result.resultError会报错

## 关于method里的方法

语音转文字的方式里，会默认将数字可读化

所以如果要发送验证码 是需要将数字用逗号隔开

此时可以调用SpritCode方法

它会将输入666666转换为6,6,6,6,6,6

方便语音播放

##  

技术来源于中遣互联，在开发时借鉴了数个开源项目。
