## 内存管理
- 堆内存管理
  - 分配器（从**堆**中初始化相应的内存区域）  
    - 设计原理
         - 分配方法  
             - 线性分配器
             - 空闲链表分配器  
                 * **策略**   
                     * 首次适应:从链表头开始遍历,选择第一个大小大于申请内存的内存块;
                     * 循环首次适应:从上次遍历的结束位置开始遍历，选择第一个大小大雨申请内存的内存块;
                     * 最优适应:从链表头遍历整个链表，选择最合适的内存块;
                     * **隔离适应:将内存分割成多个链表，每个链表中内存块大小相同，申请内存时先找到满足条件的链表，再从链表中选择合适的内存块；（go使用的内存分配策略与这个相似）**
         - 分级分配:***核心理念是使用多级缓存将对象根据大小分类，并按照类别实施不同的分配策略***  
           - 对象大小
             - 微对象
             - 小对象
             - 大对象
           - 多级缓存(线程缓存，中心缓存，页堆)
               - 线程缓存属于每一个独立的线程，它能够满足线程上绝大数的内存分配需求，**因为不涉及多线程，所以也不需要使用互斥锁来保护内存，能够减少锁竞争带来的性能损耗。** 当线程缓存不能满足需求时，运行时会使用中心缓存作为补充解决小对象的内存分配，再遇到32kb以上的对象时，内存分配器会选择页堆直接分配大内存
         - 虚拟内存布局
    - 内存管理组件:内存管理单元(mspan),线程缓存(mcache),中心缓存(mcentral),页堆(mheap)
        - 内存管理单元:mspan
        - 线程缓存：mcache
        - 中心缓存：mcentral
        - 页堆：mheap
    - 内存分配
      - 微对象
        - 线程缓存(mcache)中tiny字段指向了maxTinySize大小的块,如果当前块中还包含大小合适的空闲内存，运行时会通过**基地址**和**偏移量**获取并返回这块内存；
      - 小对象
          1. 确定分配对象的大小以及跨度类；
          2. 从线程缓存、中心缓存或者堆中获取内存管理单元并从内存管理单元找到空闲的内存空间；
      - 大对象 
        - 运行时对于大于32kb的大对象会单独处理,不会从线程缓存和中心缓存中获取内存管理单元，直接调用runtime.mcache.allocLarge分配大片内存;
    - 总结
      1. 分级分配：核心理念是使用多级缓存将对象按照大小分类，并按照类别实施不同的分配策略 
      2. mspan是基本的内存管理单元,每一个mspan承载着一种对象大小的内存管理;
      3. mcache是线程缓存,与线程上的处理器一一绑定，主要用来缓存用户程序申请的微小对象。每一个线程缓存都持有68x2个mspan（noscan:非指针类型,scan:指针类型）
      4. mcentral是中心缓存,与线程缓存不同,访问中心缓存中的内存管理单元需要使用互斥锁；
      5. mheap:堆上初始化的所有对象都由该结构体管理;该结构体包含两组重要的结构：
         (1). 全局的中心缓存列表mcentral
         (2). 管理堆区内存区域的arenas(稀疏矩阵) 
      6. 不同大小的对象分配内存的策略
        (1). 微对象(0,16B)
              将多个较小的内存分配请求合入同一个内存块中，只有当内存块中的所有对象都需要被回收时，整片内存才可能被回收
        (2). 小对象[16B,32KB]
            a. 确定分配对象的大小以及跨度类;
            b. 从线程缓存，中心缓存或者堆中获取内存管理单元并从管理单元中找到空闲的内存空间
        (3). 大对象(32KB,+00)  
           -  不会从线程缓存或者中心缓存中获取内存管理单元，而是直接调用runtime.mcache.allocLarge分配大片内存
  
  - 收集器(垃圾回收)
    - 垃圾回收为啥会有stw?
      - 如果没有stw，当标记清理的线程和用户程序线程并行执行，会导致用户程序修改已经被标记的对象的引用，导致该对象被误清理或者没有清理。所以需要在标记阶段停止用户程序运行，只做标记的处理。
  
    - golang垃圾回收三色标记法是如何减少stw的呢？
      - 写屏障技术。写屏障技术可以保证用户线程和垃圾回收线程并发执行时保证强三色一致性或者弱三色一致性。强三色一致性是黑色对象不能引用白色对象，弱三色一致性是被黑色对象引用的白色对象都处于灰色对象保护
      写屏障分为插入写屏障和删除写屏障；插入写屏障是当有用户程序修改对象引用时，当黑色对象引用白色对象时，将白色对象染色为灰色。删除写屏障是对象引用被删除时，将白色的对象染成灰色；
      - 混合写屏障技术：
        - GC开始将栈上的对象全部扫描并标记为黑色；
        - GC期间任何栈上新创建的对象都为黑色；
        - 黑色对象引用白色对象时，将白色对象染为灰色；
        - 对象引用被删除时，将白色对象染为灰色；
  
    - 为什么三色标记法可以使用写屏障降低stw，但是标记-清除不能用写屏障降低stw？
      - 三色标记法引入灰色对象这样一个中间态，中间状态会增加一次扫描的机会。在用户程序修改原本已经扫描完成的对象时，可以根据写屏障去对修改引用后的对象染色为灰色，然后对灰色对象及它的下游再做一次扫描保证对象不被误删。
      - 降低stw的机制总结：三色标记+写屏障
  
    - 标记清除法：（启动STW-标记-清除-停止STW）优化（启动STW-标记-停止STW-清除）
      1. 标记阶段：从根对象出发查找并标记堆中所有存活的对象；
      2. 清除阶段：遍历堆中的对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表；
     
    - 三色标记法：(白色对象：被回收对象；灰色对象：中间状态，会扫描灰色对象引用的子对象；黑色对象：可达对象，不被回收；)
      1. 从灰色对象集合中选择一个灰色对象并标记成黑色；
      2. 将黑色对象指向的所有对象都标记为灰色；
      3. 重复上述两个步骤；
    - golang gc哪些阶段需要stw？为什么？
      - 标记
        - 标记准备:开启写屏障，根对象扫描(stw)（该阶段为什么会开启stw?我理解开启写屏障是需要是需要暂停所有运行的goroutine的）
        - 扫描标记:用户程序恢复执行，与gc标记线程并行执行(解除stw)。gc标记线程开始扫描所有根对象，包括goroutine栈（扫描goroutine栈期间会暂停当前处理器P）,全局对象；将其加入标记队列(灰色队列)并循环处理灰色队列对象直到灰色队列为空。
      - 标记终止:重新扫描栈(1.8之前),因为栈上没有写屏障(没有对栈上的对象开启写屏障)。并发扫描阶段用户程序可能会修改栈上对象的指针，所以需要stw重新扫描(暂停程序将所有栈对象标记为灰色并重新扫描(stw));(1.8版本后引入了混合写屏障不再需要对栈进行重新扫描从而大大减少了标记终止阶段的工作量)
      - 清除
      - 清除终止

    - 混合写屏障(减少stw)
      - 为了防止1.8之前标记终止阶段重新扫描栈对象，引入了混合写屏障技术(将创建的所有新对象都标记为黑色,防止新的栈对象和堆对象被错误的回收)
      - gc开始将栈上的对象全部扫描并标记为黑色
      - gc期间任何栈上创建的对象均为黑色
      - 被删除的对象标记为灰色
      - 被添加的对象标记为灰色(垃圾回收阶段只要堆上的一个赋值*slot=ptr就会被hook住，然后把旧值指向的对象( *slot)添加的扫描队列(置灰),把新值(ptr)指向的对象也置灰)

    - stw实现
      - 将goroutine的抢占标识置为true,每个函数入口都会识别该标识来决定是否让该gouroutine被抢占（该gouroutine自己放弃执行权）

    - 黑白灰的状态描述(队列+位图)
      - 白:对象所在span的gcmarkBits中对应的bit为0且不在队列
      - 灰:对象所在span的gcmarkBits中对应的bit为1且对象在扫描队列
      - 黑:对象所在span的gcmarkBits中对应的bit为1且对象已经从扫描队列处理并移除 

- 栈内存管理



## 调度器
1. GPM模型是什么？
  - G:调度器中待执行的任务，占用更小的内存空间，也降低了上下文切换的开销；
  - P:是线程和Goroutine的中间层，它能提供线程运行需要的上下文环境;P结构中持有一个由可运行的goroutine组成的环形运行队列，还持有一个线程。
  - M:操作系统线程
2. GM模型可以实现协程的调度，为什么还需要P?
  - GM模型只维护了一个调度器持有的全局队列，每次M从全局队列中获取G运行时都需要加锁;
  - 线程间需要传递可运行的goroutine，效率低；(为什么？)
  - 每个线程需要处理内存缓存，导致大量内存占用;
  - 系统调用频繁阻塞和解除阻塞正在运行的线程，增加了额外开销;
3. GPM完整的调度流程是什么?(即,go func()经历了什么过程？)
  - 初始化结构体
    1. 获取或者创建新的goroutine结构体；
    2. 将传入的参数移到goroutine的栈上；
    3. 更新goroutine结构体中的参数,包括栈指针，程序计数器并更新其状态到_Grunnable。
  - 加入运行队列
    1. 如果本地队列(P)有剩余空间，将goroutine加入处理器(P)持有的本地队列；如果本地队列没有剩余空间，将goroutine加入调度器持有的全局队列。
  - 调度循环
    1. 调度器会从以下几个地方查找待执行的goroutine:
       1. 当全局队列中有待执行的goroutine时，会有一定几率从全局队列中查找可运行的goroutine；
       2. 从处理器本地队列中查找待执行的goroutine；
       3. 如果上述两种方式都没有获取到goroutine，会通过窃取的方式从其他P本地队列中查找待执行的goroutine；
    2. 执行获取到的goroutine,将goroutine调度到当前线程上(具体是如何调度的？会将协程栈上的内容拷贝到线程栈上执行吗？还是goroutine栈和线程栈共享呢？)
       1. 从gobuf中去取出exit函数的程序计数器和待执行用户函数的程序计数器放到关联线程的执行栈上
       2. 线程开始执行用户函数，用户函数执行完后会去执行exit函数
       3. exit函数会将goroutine转换为dead状态；移除goroutine和线程的关联并将gouroutine加入到空闲列表中
       4. 最后exit函数会重新调用调度方法触发新一轮的goroutine调度
4. 如果运行中的goroutine发生挂起或者发生系统调用，会怎么处理? 
   - 会触发调度(调度器会重新选择goroutine在线程上执行)
      1. 主动挂起触发调度的流程:
         1. 将当前goroutine的状态从运行状态切换到阻塞状态；
         2. 解除goroutine和线程的关联；
         3. 触发新一轮调度；
         4. 当goroutine等待的条件满足后，将goroutine切换到可运行状态并加入P的运行队列;
      2. 系统调用触发调度流程:
         1. 当系统调用时，会进入系统调用前的准备工作和系统调用后的收尾工作；
         2. 准备工作包括：保存当前当前程序的栈信息;更新gouroutine状态到系统调用阻塞状态，M陷入系统调用等待返回;将P和M分离;
         3. 收尾工作：系统调用结束为goroutine重新分配资源，寻找一个用于执行当前goroutine的处理器P
  [*必看go调度器详解](https://mp.weixin.qq.com/mp/homepage?__biz=MzU1OTg5NDkzOA==&hid=1&sn=8fc2b63f53559bc0cee292ce629c4788&scene=25#wechat_redirect)