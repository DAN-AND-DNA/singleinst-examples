# examples

## KVStorage
非常简单的key value数据存储服务，用来说明单例模块的使用方式:

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

## 用法
1. git clone https://github.com/DAN-AND-DNA/singleinst 
2. git clone https://github.com/DAN-AND-DNA/singleinst-examples
3. 复制 ./singleinst-examples/examples 到 ./singleinst/ 中
4. 分别启动examples中的3个服务
5. postman 发送POST请求, 路径为 http://127.0.0.1:3737/nanogo/webbff/webbff/login 来获得token
6. 填充头部Token，发送POST请求 http://127.0.0.1:3737/nanogo/webbff/webbff/set 来设置值
7. 填充头部Token，发送POST请求 http://127.0.0.1:3737/nanogo/webbff/webbff/get 获得值

```json
// login request
{
    
    "name": "Dan",
    "password": "12345678"
}

// login response
{
    "token": "login_abcdef1234567",
    "base_userinfo": {
        "uid": "u10001",
        "username": "Dan",
        "age": 30
    }
}

// set request
{
    "new_value":{
        "key": "money",
        "value": "99991"
    }
}

// get request
{
    "key": "money"
}
```