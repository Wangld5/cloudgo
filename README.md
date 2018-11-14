### cloudgo

#### 模板选择

我选用了三个模板，包括mux，negroni和render。其中而gorilla/mux是一个强大的路由，小巧但是稳定高效，不仅可以支持正则路由还可以按照Method，header，host等信息匹配，可以从我们设定的路由表达式中提取出参数方便上层应用，而且完全兼容http.ServerMux 。negroni是一个web中间件,方便使用net/http的库;它实现了http.Handle接口,兼容http.Handle。render主要用于模板的渲染，类似于vue。

#### curl测试

命令为：`curl -v http://127.0.0.1:8001/hello/wangld`

![](https://github.com/Wangld5/cloudgo/raw/master/pic/curl.png)

正常输入name，测试成功。

#### ab压力测试

命令为：`ab -n 1000 -c 100 http://127.0.0.1:8001/hello/wangld`

![](https://github.com/Wangld5/cloudgo/raw/master/pic/ab1.png)

![](https://github.com/Wangld5/cloudgo/raw/master/pic/ab2.png)

其中平均响应时间为58.615秒，测试成功。

#### 参数解释

-n 1000 表示总请求数位1000

-c 100表示并发用户数为100

http://127.0.0.1:8001/hello/wangld 表示这些请求的目标URL。

Server Hostname 表示请求的URL中的主机部分名称，它来自于http请求数据的头信息.。

Server Port 为服务器的端口

Document Path 表示请求的URL中根绝对路径，它同样来自于http请求数据的头信息。 

Document Length 表示http响应数据的正文长度。 

Concurrency Level 表示并发用户数，这是我们设置的参数。 

Time taken for tests 表示所有这些请求被处理完成花费的总时间。 

Complete requests 表示总请求数，这是我们设置的相应参数。 

Failed requests 表示失败的请求数，这里的失败是指请求的连接服务器、发送数据、接收数据等环节发生异常，以及无响应后超时的情况。 

Total transferred 表示所有请求的响应数据长度总和，包括每个http响应数据的头信息和正文数据的长度。 

HTML transferred 表示所有请求的响应数据中正文数据的总和 。

Time per request 这便是前面提到的用户平均请求等待时间 。

Transfer rate 表示这些请求在单位时间内从服务器获取的数据长度，它等于： Total transferred / Time taken for tests 这个统计项可以很好的说明服务器在处理能力达到限制时，其出口带宽的需求量。 