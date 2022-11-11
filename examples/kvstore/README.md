# KVStorage

key value数据存储服务:

1. WebBFF为web端提供Restful API，对后端接口进行封装，提供少量数据缓存
2. KVService，grpc服务，提供数据存取服务
3. UserService，grpc服务，提供用户权限校验服务

```c++
k8s中的架构                ______________________________ 
                          | k8s             KVService   |
                          |               /             |  
                          |        WebBFF - KVService   |  -- DB集群
                          |      /        \             |
Clients -- WAF -- SLB --  | Ingress         KVService   | 
                          |      \                      |
                          |        WebBFF - UserService |  -- Redis集群
                          |               \             |
                          |                 UserService |
                          |_____________________________|
```