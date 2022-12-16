# 对象存储简介

### 1、传统网络存储

*  NAS：Network Attached Storage	==文件级别== 

  NAS是一个提供了存储功能和文件系统的网络服务器，客户端可以访问NAS上的文件系统，进行上传和下载。

* SAN：Storage Area Network      ==块级别== 

  SAN是只提供了块存储，对于文件系统的管理直接交给客户端来实现。对于客户端而言，SAN是一块磁盘，可以对其格式化，创建文件系统并挂载。

  

### 2、对象存储的不同

* 管理方式：以对象形式来管理数据，包含三部分：

  * 对象的数据
  * 对象的元数据
  * 对象的全局唯一标识符（ID）

  对象存储提供了编程接口，使应用程序能够操作数据。在基础层面，这包括用于基础读、写和删除操作的增删改查（[CRUD](https://zh.wikipedia.org/wiki/增刪查改)）功能。一些对象存储的实现更进一步，支持[对象版本控制](https://zh.wikipedia.org/w/index.php?title=对象版本控制&action=edit&redlink=1)、对象复制、生命周期管理以及对象在不同层级和类型的存储之间的移动等附加功能。大多数API实现是基于[REST](https://zh.wikipedia.org/wiki/表现层状态转换)的，允许使用许多标准的[HTTP](https://zh.wikipedia.org/wiki/HTTP)调用。

* 访问方式：对象存储通过REST访问数据/对象。也就是GET,PUT,POST,DELETE几种方式。

* 优势：

  * 提高了存储系统的可拓展性（全局唯一标识符）：

    对象存储架构的拓展只需要增加新的存储节点就可，其管理开销比起传统的网络存储小。

  * 以更低的代提供数据冗余：

    传统网络存储通常通过保持数据的多个副本来实现数据的冗余，分布式对象存储在存储系统中一个或多个节点失效时，对象仍然可用。

    

### 3、对象存储原型

* Put /objetcs/<object_name> 在对象存储中，通常通过一个Put来将对象上传到服务器。

* Get /objetcs/<object_name> 在对象存储中，通常通过一个Get将数据下载到本地。如果对象不存在，则返回404。

* Go语言实现：

  * main函数启动。

    ```go
    package main
    
    import (
    	"ObjectStorage/chapter1/objects"
    	"log"
    	"net/http"
    )
    
    func main() {
    	// http.HandleFunc来绑定对应pattern的处理函数，无法区分Method，必须自己编码区分
    	http.HandleFunc("/objects/", objects.Handler)
    	// 监听默认端口并启动服务
    	log.Fatal(http.ListenAndServe(objects.Port, nil))
    }
    ```

    

  * Get请求

    ```sh
    curl localhost:8080/objects/1.txt
    ```

  * Put请求

    ```sh
    curl localhost:8080/objects/2.txt -XPUT -d"2.txt"
    ```

    

