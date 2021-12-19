# tk-cache
Go 实现内存型分布式缓存，基于 lru 缓存淘汰算法和一致性 hash 算法

# 架构图
![image](https://raw.github.com/tomkiin/repositpry/master/tk-cache/doc/architecture.jpg)

架构设计为两层结构，node 为缓存节点，group 为控制节点；
### NODE
- 缓存类型为内存型
- 基于 lru 淘汰算法控制缓存容量；
### GROUP
- 负责用户交互和缓存节点的调度；
- 使用一致性 hash 算法实现对 node 节点负载均衡；
- 使用 singlefight 机制防止缓存击穿；
- router 组件负责对外提供 RESTful 操作接口，对内提供 node 注册接口；
- manager 组件负责对缓存节点的控制，如注册节点、注销节点、监控节点存活状态等；