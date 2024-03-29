1. mysql索引原理
2. mysql事务
3. mysql锁 
4. 三层b+树索引可以存多少条数据？
5. 联合索引的技巧
  - 覆盖索引
  - 最左匹配
  - 索引下推
6. mysql知识体系
  - 单机mysql
    - 事务
    - 索引
    - 锁
    - 日志
  - 多机mysql
    - 主从一致
      - 主从复制延迟
    - 高可用 
    - 主从切换
    - 读写分离
7. next-key lock的作用










1. 三层b+树索引可以存多少条数据？
   答：
    背景
      Innodb引擎数据是索引记录和用户记录都存在"页"中，页大小是16kb;
      B+索引的存储结构是:叶子结点存用户记录，非叶子结点存索引记录;
      索引记录主要存用户记录的主键和指向其他页的指针;
      用户记录存用户建立的行数据;
    数据
      页大小是16kb
      非叶子结点中主键大小是8b;
      非叶子结点中页指针是6b;
      假设叶子结点中每行记录是1kb  
    计算
      第一层:
        最少2条
        最多16*1024/(6+8)=1170条
      第二层:
        最少1170+1=1171条
        最多1170*1170=1368900条
      第三层: 
        最少1170*16+1=1w+
        最多1368900*16=2kw+

        备注：b+树每一个节点都是一个页，如果页满了(16kb),就会分裂出新的页，此时树的节点就会增加可能会引起树高的变化

7. next-key lock含义及作用
   - next-key lock是行锁和间隙锁的合称


## MySql实战45讲      
  - 一条sql查询语句是如何执行的?
    - mysql架构可以分为Sever层和存储引擎层两部分
      - Sever层:
      - 存储引擎层:

  - 一条sql更新语句是如何执行的?
    - redolog(存储引擎层InnoDB):数据库崩溃后的数据恢复
    - binlog (Server层):
    - 两阶段提交:为了使redolog和binlog逻辑一致
    - 为什么binlog不能做数据库崩溃恢复?

  - 事务隔离(保证事务正确可靠: 原子性,一致性,隔离性,持久性)
    - 读未提交: 一个事务还没提交它做的变更就能被别的事务看到
    - 读提交:一个事务提交之后它做的变更才能被其他事务看到
    - 可重复读:一个事务执行过程中可看到数据总是跟这个事务在启动时看到的数据一致
    - 串行化:对于同一行记录写会加写锁,读会加读锁，当出现读写冲突时，后访问的事务等前一个事务执行完成后才能执行

  - 事务隔离实现(可重复读，读提交，读未提交，串行化)
    - 数据库会创建一个视图，访问的时候以视图的逻辑结果为准。在"可重复读"隔离级别下，这个视图在事务启动时创建的，整个事务存在期间都用到这个视图。在"读提交"隔离级别下，这个视图是在每条sql语句开始执行的时候创建的。在"读未提交"隔离级别下直接返回记录上的最新值，没有视图的概念。在"串行化"隔离级别下直接使用加锁的方式避免并行访问。

  - 事务隔离的具体实现(以"可重复读"为例)
    - 同一条记录在系统中可以存在多个版本(即mvcc:数据库的多版本并发控制),每条记录在更新的时候都会同时记录一条回滚操作(记录在undolog中)，记录上最新值通过回滚操作可以得到前一个状态的值(当前版本的值指向上一个版本(可以使用事务号查找))
    - 事务启动时会生成事务id(递增的)并且每次事务更新数据时，都会生成一个新的数据版本并且把事务id赋值给这个数据版本的事务ID，所以表中的一行数据有多个版本，每个版本都有自己的事务id
    - Innodb为每个事务构造了一个数组，用来保存这个事务启动瞬间当前正在"活跃"的所有事务ID(活跃:启动但未提交),数组里面事务id最小值记为低水位，系统中的最大事务id+1记为高水位.Innodb就是根据数据的多版本特性创建"快照"的。
      - 当select的最新数据版本中的事务id<低水位，说明这个版本是已提交的,可以直接查看;
      - 当select的最新数据版本中的事务id(高水位>id>低水位),分两种情况：
        - 如果该事务id在数组中，该版本数据不可见；通过undolog回退前一个版本对比事务id。
        - 如果该事务id不在数组中，说明是已提交的版本，该数据版本可见;
      - 当select的最新数据版本中的事务id>高水位，说明是后续新事务修改了该数据，不可见;通过undolog回退前一个版本对比事务id
  - 事务的可重复读能力是怎么实现的？
    - 可重复读核心是一致性读，一致性读是根据视图(快照)实现的，视图是依据数据多版本特性实现的。
    - 事务更新数据(先读再更新)时，只能用当前读(这个数据的最新版本)。如果当前的记录的行锁被其他事务占用就需要进入锁等待了。
  - 事务两阶段锁协议
      - 事务中某条语句执行时才加锁但是等到事务提交后才释放锁


  - 索引 
    - InnoDB索引模型 
      - b+树:每个结点都是一个数据页，每个数据页16kb，非叶子结点存放索引和页指针，叶子节点存放行数据(主键索引)或者关联的主键id(普通索引)。叶子结点存放全量数据，数据按照索引从小到达排列，同层每个结点都通过双指针链接（方便范围查找）。
    - 主键索引: 按照主键建立的b+树
    - 普通索引: 按照非主键字段建立的b+树
    - 覆盖索引: 查找的数据在索引中
    - 联合索引: 使用多个字段建立索引，依次按照每个字段排序建立b+树
    - 最左前缀原则: 可以利用索引的最左前缀定位记录 
    - 索引下推: 通过联合索引中的字段做过滤

  - 全局锁(全库只读不能写)
    - 使用场景
      - 全库备份
  - 表级锁
  - 行级锁

  - 更新操作的事务隔离


  - 幻读:一个事务在前后两次查询同一个范围的时候，后一次查询看到了前一次查询没有看到的行
    - [*参考](https://cloud.tencent.com/developer/article/2136022)


  -[*资料](https://xiaolincoding.com/)