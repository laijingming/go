## 旧版调度器（已废弃）
符号 |含义
---|---
G  |goroutine协程
M  |thread线程
<img width="770px" height="575.02px" src="https://www.topgoer.com/static/7.1/gmp/10.jpg" alt="">

M想要执行、放回G必须访问全局G队列，并且M有多个，即多线程想要访问同一资源需要加锁保证同步/互斥,所以全局队列有互斥锁进行保护。

老调度器的缺点：
- 执行、销毁、创建G都需要每个M获取锁，这就形成了激烈地锁竞争。
- M 转移 G 会造成延迟和额外的系统负载。比如当 G 中包含创建新协程的时候，M 创建了 G’，为了继续执行 G，需要把 G’交给 M’执行，也造成了很差的局部性，因为 G’和 G 是相关的，最好放在 M 上执行，而不是其他 M’。
- 系统调用 (CPU 在 M 之间的切换) 导致频繁的线程阻塞和取消阻塞操作增加了系统开销。
## [Goroutine调度器的GMP模型的设计思想](https://www.bilibili.com/video/BV1N84y117ii/)
符号 |含义
---|---
G  |goroutine协程
P  |processor处理器
M  |thread线程
面对之前调度器的问题，Go 设计了新的调度器。

在新调度器中，出列 M (thread) 和 G (goroutine)，又引进了 P (Processor)。

<img width="770px" height="575.02px" src="https://www.topgoer.com/static/7.1/gmp/12.jpg" alt="">

- 全局队列（Global Queue），存放等待运行的G。
- P本地队列：同全局队列一样，存放等待运行的G，存储数量有限不超过256个，新建G'优先存放本地队列，当本地队列满时，取一半G放入全局队列。
- P列表：在程序启动时创建，并保存在数组中，最多有GOMAXPROCS（可配置）个。
- M：线程想要运行任务就得获得P，从P队列获取G，P队列为空时，M尝试从全局队列拿一批G放到P的本地队列，或从其他P偷一半放到自己本地队列。M运行G，执行之后会从P获取下一个G，不断重复下去。

### 调度器调度场景
* 场景 1：拥有 G1，M1 获取 P 后开始运行 G1，G1 使用 go func() 创建了 G2，为了局部性 G2 优先加入到 P1 的本地队列。
* 场景 2：G1 运行完成后 (函数：goexit)，M 上运行的 goroutine 切换为 G0，G0 负责调度时协程的切换（函数：schedule）。从 P 的本地队列取 G2，从 G0 切换到 G2，并开始运行 G2 (函数：execute)。实现了线程 M1 的复用。


