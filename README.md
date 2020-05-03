# distributed-uid-generator

> 使用go语言编写的分布式id生成器。项目使用twitter的snowflake算法生成唯一id.使用httpRouter和go原生网络模块向外界提供http访问.添加Dockerfile用于docker部署．

## 默认位数分配

- 时间位数:28bit,单位ms
- 机器位数:22bit,(可以自行添加DataCenter标识位)
- 序列位数:13bit
