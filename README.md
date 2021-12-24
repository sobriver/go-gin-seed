go project seed, include 
- gin
- logrus
- mysql

# 编码规范
- 所有错误码统一定义在code文件
- 所有函数返回信息必须包含错误码， 且错误码在返回信息的最后一个
- model层返回实体信息和错误码
- service层只返回错误码