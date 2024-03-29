### 资料
1. 共识算法概览:https://draveness.me/consensus/
2. raft算法细节,待看https://raft.github.io/
3. 深入raft协议https://juejin.cn/post/6907151199141625870

### 问题:
1. 分布式系统节点通信有两种：共享内存和消息传递？分别是怎么实现的？
    消息传递:rpc方式
    
2. 集群节点一致性(日志复制一致性)被划分为四个可以被独立解释并处理的子问题：领导选举，日志复制，安全性，成员变更.
    - 领导选举:
        - 有三种角色:领导者,跟随者,候选人
        - 集群初始化时，所有节点都是跟随者，每个跟随者都会启动随机选举超时时间，最先超时的节点从跟随者变为候选人并将任期号加一开始进行选举(向其他节点发送请求选举rpc)，结果有如下三种情况:
          - 赢到选举成为领导者，后续向其他节点发送心跳信息阻止其他节点选举
          - 其他节点成为领导者，在等待投票期间，跟随者可能收到其他服务器节点声明它是领导人的选举rpc。如果这个节点的任期大于等于我的任期，我会承认领导者的合法性并回到跟随者状态。如果这个节点的任期小于我的任期，我会拒绝这次rpc并继续保持候选人状态。
          - 多个跟随者同时成为候选人，选票可能被瓜分而不能产生领导人。根据超时时间重新发起选举.
    - 日志复制:
      - 每一条日志包含任期号、索引值和状态机指令;
      - 客户端向领导者发送一条指令，领导者会把该指令添加的日志中，然后并行的向其他节点发起日志复制rpc,当一多半的其他节点返回成功时(日志与领导者日志一定相同)，日志被定义为已提交状态，然后领导人会在状态机上运行该指令并返回给客户端。

    - 安全性

    - 成员变更:集群中有节点增加和配置变化，会出现两个领导者的情况。如何解决？
      - 为什么会出现两个领导者？答：假如当前集群只有两个节点A,B;后续集群中又增加了三个节点C,D,E; A,B的配置文件中总节点数还是两个(旧配置),但C,D,E的配置文件中总节点数是五个(新配置),此时A在选举中获取到了两票(满足旧配置多数选票原则);此时C在选举中获取到了三票(满足新配置5/3原则),所以集群中出现了两个领导者A,C;
      - 如何解决？

3. raft是如何保证数据一致的？(日志复制形式)
- 在分布式环境中，如果我们要让一个服务具有容错能力，最常用的方法是让一个服务的多个副本同时运行在不同的节点上。为了保证多个副本在运行时状态都是同步的，即客户端无论将请求发送到哪一个节点中，最后都能得到相同的结果，因此采用状态机复制方法。
- 状态机复制通常使用**日志复制实现**，每个服务器节点存储一个包含一系列命令的日志，其状态机按照顺序执行日志中的命令，每个日志中的命令都相同且顺序也一致，因此每个状态机处理相同的命令序列，这样就能得到相同的状态和输出序列。