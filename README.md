# examples

[中文文档](README_cn.md)

## KVStorage
A very simple key value data storage service to illustrate the use of single instance modules:

1. WebBFF provides Restful API for the web side, encapsulates the back-end interface and provides a small amount of data caching
2. KVService, grpc service, provides data access service
3. UserService, grpc service, provides user permission verification service

```c++
Architecture in k8s        ______________________________ 
                          | k8s             KVService   |
                          |               /             |  
                          |        WebBFF - KVService   |  -- DB clusters
                          |      /        \             |
Clients -- WAF -- SLB --  | Ingress         KVService   | 
                          |      \                      |
                          |        WebBFF - UserService |  -- Redis clusters
                          |               \             |
                          |                 UserService |
                          |_____________________________|
```

## 用法
1. git clone https://github.com/DAN-AND-DNA/singleinst 
2. git clone https://github.com/DAN-AND-DNA/singleinst-examples
3. cp -r ./singleinst-examples/examples ./singleinst/
4. Start each of the 3 services in examples
5. Send a POST request using postman to get the token:  
    http://127.0.0.1:3737/nanogo/webbff/webbff/login 
6. Fill the header Token and send a POST request to set the value:  
    http://127.0.0.1:3737/nanogo/webbff/webbff/set 
7. Fill the header Token and send a POST request to set the value:   
    http://127.0.0.1:3737/nanogo/webbff/webbff/get

```golang
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
