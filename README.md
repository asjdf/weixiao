# weixiao
WeiXiao SDK for Go   （腾讯微校SDK）

[腾讯微卡官方文档](https://wiki.weixiao.qq.com/api/)

## 结构
包名遵照官方文档中的服务类型设定，包中相应的文件名也与文档保持一致

- third: 应用服务
- school: 主体信息服务
- guard: 门禁服务

## school

### 身份认证
 - [x] 账号+密码验证（接口对接）
 - [ ] 跳转校企页面认证
 - [ ] 姓名+证件号验证（接口对接）


###  校园卡 
 - [x] 用户校园卡面余额字段同步
 - [ ] 用户校园卡面借书字段同步
 - [ ] 用户校园卡面补助字段同步
 - [ ] 用户校园卡面餐次字段同步
 - [x] 用户卡面信息项主动更新
 - [ ] 用户卡面信息项自定义接口
 - [ ] 解绑学生校园卡
 ### 校园码
  - [x] 校园码解码
  - [ ] 校园码扫码事件回调

### 错误代码
error code
40001 bind error
40002 appkey is not equal
40003 CBCDecrypter error
40004 Unmarshal error
40005 marshal error