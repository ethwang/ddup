## 问题
1. golang内存分配机制；
2. golang垃圾回收机制；
3. golang调度器&GPM模型；
4. golang csp并发模型及channel原理;
5. golang new和make的区别；
6. go context原理；
7. go mutex原理
8.  切片传参
9.  map原理
10. slice原理
11. 基于信号的抢占式调度原理
12. go defer原理
13. go http原理
14. go 实现map并发
15. goroutine栈扩缩容
16. go 定时器原理和坑





## 解答

4. golang channel原理;
    - 作用
      - channel是goroutine之间的通信的媒介
    - 设计原理  
        - goroutine可以使用共享内存加互斥锁进行通信，同时也提供了另一种并发模型-csp模型。goroutine和channel分别对应csp中的实体和传递信息的媒介，goroutine之间会通过channel传递数据
        - channel运行时内部数据结构是hchan(循环队列+双向链表(阻塞的goroutine链表)),该结构中包含了用于保护成员变量的互斥锁，channel是一个用于同步和通信的有锁队列
    - 数据结构(循环队列+双向链表+互斥锁)
      - channel运行时使用hchan结构体表示
      - hchan中持有互斥锁，一个循环队列和两个双向链表
        - 循环队列:存放的缓冲区的数据指针
        - 双向链表:当前channel缓冲区不足而阻塞的goroutine列表
            - sendq:发送数据时并没有找到接收方并且缓冲区已满，此时goroutine会将自己加入到channel的sendq队列中
            - recvq:
        - 互斥锁
    - 创建channel(make->makechan)
      - 如果当前channel不存在缓冲区，只会为hchan结构体分配(mallocgc)一块内存空间
      - 如果当前channel存在缓冲区且缓存区存放的数据类型不是指针,会为hchan结构体和底层数组分配一块连续的内存空间
      - 如果当前channel存在缓冲区且缓冲区存放的数据类型是指针,会为hchan结构体和底层数组分别分配内存
      - 问题：为什们有指针的内存分配需要单独配置呢?
  
    - 发送数据(发送之前先为当前channel加锁，防止多个goroutine并发的修改数据)
        - 如果channel已经关闭，向该channel发送数据会报错中止程序
        - 当存在等待的接受者时，将数据直接发送给阻塞的接受者
          - 将发送数据直接拷贝到接收变量所在的内存地址上
          - 唤醒等待接受数据的goroutine并标记成可运行状态，把该goroutine放到发送方处理器P下一个可执行的位置，该p下一次调度时会执行被唤醒的goroutine
        - 当缓存区存在空余空间时，将发送的数据写入channel缓存区
            - sendx(指向循环数组下一个可被存放的发送位置)索引所在位置
        - 当不存在缓冲区或者缓冲区已满时，将该goroutine放入发送等待队列并挂起(触发调度)等待唤醒
    - 接受数据
        - 当存在等待的发送者时，从阻塞的发送者或者缓存区中获取数据
            - 分两种情况
              - 如果channel不存在缓存区，将channel发送队列中的goroutine数据复制到目标内存地址中，然后将当前处理器P的下一个可执行位置设置成发送数据的goroutine
              - 如果channel存在缓存区，将队列中的数据复制到接收方的内存地址，将发送等待队列头的数据复制到缓冲区中，释放一个阻塞的发送方
        - 当缓存区存在数据时，从缓存区接受数据
        - 当缓存区不存在数据时，将该goroutine放入接受等待队列并挂起等待唤醒
    - 关闭channel
        - 当channel时一个空指针或者已经被关闭时，直接抛出异常
        - 唤醒rendq和recvq中所有被阻塞的goroutine

5. golang new和make的区别；
    - make的作用是初始化内置数据结构(切片，哈希表和channel):申请内存+初始化
      - make([]int,0,100):make slice是返回一个包含data,cap和len的结构体
      - make(map[int]bool):make map返回一个hmap结构体指针
      - make(chan int,5):make channel返回一个hchan结构体的指针
    - new的作用是根据传入的类型分配一块内存空间并返回指向这块内存空间的指针:申请内存+置为零值

6. go context原理(context为同一个任务的多个goroutine之间提供退出信号通知和数据传递的功能)
    - 一个接口(四个待实现的方法)
        - Deadline()
        - Done()
        - Err()
        - Value(key interface{})
    - 四种实现
        - emptyctx:空context
        - cancelctx:可取消的context
        - timerctx:在cancelctx基础上封装了一个定时器和一个截止时间,达到定时时间或者截止时间自动取消
        - valuectx:给context附加键值对信息，可以向下传递信息
    - 六个函数
        - background:返回一个空context
        - TODO:返回一个空context
        - WithCancel:把一个context包装成cancelctx并返回一个取消函数
        - WithDeadline:可以创建timerctx,传入截止时间点
        - WithTimeout:可以创建timerctx,传入截止时间段
        - WithValue:附加键值对信息向下传递

7. go mutex原理(读写互斥锁,同一时间只能有一个goroutine抢占锁，其他goroutine等待锁释放)
    - 结构体
      - state:存储互斥锁的状态。加锁解锁都是通过atomic包提供的函数原子性操作该字段。
      - sema:用作信号量,主要用作等待队列。
    - 模式 
      - 正常模式:一个尝试加锁的goroutine会先自旋几次尝试通过原子操作获取锁,若几次自旋之后不能获取到锁，则通过信号量排队等待，所有的等待者都会按照先入先出的顺序排队，但是当锁释放，第一个等待者被唤醒后并不会直接拥有锁而是需要和后来者竞争（也就是那些处于自旋阶段尚未排队等待的goroutine）,这种情况下后来者更有优势(一方面他们正在cpu上运行另一方面处于自旋的goroutine可以很多被唤醒的gouroutine只有一个),所以被唤醒的goroutine大概率拿不到锁。这种情况下被唤醒的goroutine重新被插入到队列头部，而一个goroutine在本次加锁等待时间超过1ms后，它会把当前mutex模式从正常模式切换到饥饿模式。
      - 饥饿模式：锁的所有权会从执行unlock的goroutine直接传递给等待队列头部的goroutine。后来者不会自旋获取锁，它们会直接从队列尾部排队等待。当一个等待者获取锁后，它会有两种情况把mutex从饥饿模式切换到正常模式：1.获取锁的时间小于1ms;2.排队队列中没有其他goroutine
      - 总结：
        - 正常模式：排队+自旋。先自旋原子抢锁失败后再排队等待锁。优势：有更高的吞吐量，因为频繁的挂起唤醒goroutine会带来额外的开销。但不能无限制自旋（自旋4次）。缺点：队列尾端的goroutine长时间不能获取到锁
        - 饥饿模式：排队（先来后到获取锁）,优势：解决尾端goroutine长时间无法获取锁的问题
    - 实现细节
      - lock
        - fast path
            - mutex处于unlock状态且没有goroutine排队，通过cas获取锁。如果没有获取到锁进入slow path继续获取锁。(该goroutine没有获取到锁说明有其他goroutine获取到了)
        - slow path
      - unlock
        - fast path
            - 通过原子操作从state中mutexLocked减一(释放锁),如果失败通过slow path继续释放锁(失败说明有排队的goroutine获取锁)
        - slow path(除lock位为0外其他位不为0)
            - 若等待队列为空或goroutine被唤醒或获得了锁或锁进入了饥饿模式，不需要唤醒某个goroutine直接返回即可
    
    - 其他
      - 正常模式下是否一定会自旋呢？不一定
        - 单核场景不自旋
        - 如果只有一个p时不自旋
        - 如果有多个p每个p本地队列都不为空不自旋
      - 自旋的goroutine会先争抢state的唤醒标志位,来告知持有锁的goroutine在unlock时不用再唤醒其他goroutine了
      - 自旋结束
        - 自旋已经>4次
        - 锁进入饥饿模式
        - 锁被释放
      - 自旋结束的goroutine会尝试原子操作(cas)修改mutex的状态(修改state),有如下几种修改状态:
        - 如果state标识是饥饿模式则goroutine去排队
        - 如果是正常模式，就尝试修改lock位
        - 如果当前goroutine等待锁时间超过1ms,就将mutex状态切换为饥饿模式
      - 原子修改
        - 成功
          - 获取到锁
          - 进入等待队列
        - 失败
          - 继续自旋检查，释放之前唤醒的标识位

8. 切片传参问题(append会引发异常):切片传参是值传递,会复制一个相同的slice结构体，复制原切片结构体中的底层数组指针，长度，容量;所以复制出来的切片结构体和原切片结构体中的底层数组指针都是同一个, 会引发一些异常。
  [参考](https://blog.csdn.net/bestzy6/article/details/119981699)


9. map原理
    1. 结构体:hmap
       1. 元素个数
       2. 桶数量
       3. 溢出桶数量
       4. 桶:数组指针。数组值是bmap结构
       5. 扩容桶:发生扩容时，被赋值为扩容前桶的数组指针
       6. 溢出桶:一个桶有八个位置，也就是可以允许八个哈希冲突的键值对存入。超过八个就需要链接一个溢出桶来存储。
    2. 新增元素
       1. 通过key的hash值后"B"位确定哪个桶
       2. 遍历桶，通过key的tophash和hash值从数组0到7位置查找，找到第一个可以插入的位置
       3. 如果当前桶已满，会通过overflow链接一个新桶存储
    3. 查找元素
       1. 计算key的hash值
       2. 通过hash值后"B"位来确定桶
       3. 通过hash值前8位(tophash)确定在桶中的位置 
       4. 比较完整hash是否一致，一致则找到
       5. 如果没找到则到下一个溢出桶继续找
    4. 扩容
       - 扩容方式
         - 等量扩容
         - 2倍扩容
       - 扩容条件
         - 平均每个桶中数据过多
         - 溢出桶数量过多 
       - 扩容过程
         - 扩容时，将旧数据放入oldbucket中
         - 每次对map操作时，会触发从旧桶到新桶的数据迁移
         - 扩容没有完成之前，查找操作都会先遍历旧桶再遍历新桶
  [参考1](https://zhuanlan.zhihu.com/p/495998623)
  [参考2](https://golang.design/go-questions/map/assign/)

10. slice原理
    1. 结构体:slice
       1. 数组指针
       2. 长度
       3. 容量
    2. 扩容
    3. 值传递:将结构体中的长度容量和指针进行传递;[参考](https://mp.weixin.qq.com/s?__biz=MzUxMDI4MDc1NA==&mid=2247489302&idx=1&sn=c787d1fa4546e12c7e55e880da73c91f&scene=21#wechat_redirect
  [参考](https://halfrost.com/go_slice/)

11. 基于信号的抢占式调度原理
    1. 线程注册信号处理函数；
    2. sysmon线程检测到执行时间过长的goroutine或者gc stw时，会向相应的线程发送信号；
    3. 线程收到信号后，内核执行信号注册函数。注册函数的主要作用是：将抢占函数插入到当前goroutine运行位置； 
    4. 信号注册函数执行完成后，回到当前用户goroutine继续执行，此时用户goroutine会执行被插入的抢占函数。抢占函数有四个作用：
       1. 将当前goroutine从running状态变为runable状态；
       2. 将当前goroutine与执行它的线程分离
       3. 将当前goroutine放入全局可运行队列
       4. 调用schedule函数进入循环调度 
  [*参考](https://cloud.tencent.com/developer/article/1836265)
12. go defer原理
    1. defer用作延迟函数调用,return执行顺序是: 设置返回值->执行defer->ret
    2. _defer结构体会保存执行延迟函数的信息:
       1. sp函数栈指针;
       2. pc程序计数器;
       3. fn函数地址;
       4. link指针,链接多个defer
    3. 每个goroutine结构中有一个defer指针指向defer的单链表，每次申明一个defer时就将defer插入到单链表的表头，每次执行defer时就从单链表表头取出一个defer执行。(后进先出)

16. go 定时器原理和坑
    [参考](https://pengrl.com/p/1785/)



17. make slice,map,channel对比
    1. make slice 返回的是slice实体
    2. make map 返回的是map指针
    3. make channel 返回的是channel指针 