# EvHelper
EvHelper 旨在帮助在上海的人们迅速抢定到自己想要定的evcard车。本项目主要使用beego作为后台的框架，缓存使用redis。

### 整体框架
整个项目使用c/s架构，client端可以为android，ios，web端，整个的项目架构如下图：

1. client和server交互，server端返回给client一个token，client再后面的访问携带该token与server交互
2. client将位置信息发送给server，server访问redis，从redis获得用户所在位置的位置编号（evcard为每个借车/还车点都设置了编号）
3. client根据获得位置编号，和evcard用户名，密码模拟登陆订车，client一直保持订车状态直到订到车
4. client将订车信息发送给server，供server后续操作

###TODO
1. beego后台的搭建
2. redis的搭建
3. client的框架

