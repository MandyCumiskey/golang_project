package main 


import (
	"fmt"
	"time"
)

/*
Go 语言的协程实现被称之为 goroutine，由 Go 运行时管理，在 Go 语言中通过协程实现并发编程非常简单：
我们可以在一个处理进程中通过关键字 go 启用多个协程，然后在不同的协程中完成不同的子任务，
这些用户在代码中创建和维护的协程本质上是用户级线程，Go 语言运行时会在底层通过调度器将用户级线程交给操作系统的系统级线程去处理，
如果在运行过程中遇到某个 IO 操作而暂停运行，调度器会将用户级线程和系统级线程分离，以便让系统级线程去处理其他用户级线程，
而当 IO 操作完成，需要恢复运行，调度器又会调度空闲的系统级线程来处理这个用户级线程，从而达到并发处理多个协程的目的。
此外，调度器还会在系统级线程不够用时向操作系统申请创建新的系统级线程，而在系统级线程过多的情况下销毁一些空闲的线程，
这个过程和 PHP-FPM 的工作机制有点类似，实际上这也是很多进程/线程池管理器的工作机制，这样一来，可以保证对系统资源的高效利用，
避免系统资源的浪费。

以上，就是 Go 语言并发编程的独特实现模型。
*/

/*
嗯，就是这么简单，在这段代码中包含了两个协程，一个是显式的，通过 go 关键字声明的这条语句，表示启用一个新的协程来处理加法运算，
另一个是隐式的，即 main 函数本身也是运行在一个主协程中，该协程和调用 add 函数的子协程是并发运行的两个协程，
就好比从 go 关键字开始，从主协程中叉出一条新路。和之前不使用协程的方式相比，由此也引入了不确定性：我们不知道子协程什么时候执行完毕，
运行到了什么状态。在主协程中启动子协程后，程序就退出运行了，这就意味着包含这两个协程的处理进程退出了，所以，我们运行这段代码，
不会看到子协程里运行的打印结果，因为还没来得及执行它们，进程就已经退出了。另外，我们也不要试图从 add 函数返回处理结果，因为在主协程中，
根本获取不到子协程的返回值，从子协程开始执行起就已经和主协程没有任何关系了，返回值会被丢弃。

如果要显示出子协程的打印结果，一种方式是在主协程中等待足够长的时间再退出，以便保证子协程中的所有代码执行完毕：

*/
func main (){
   go add (1,2)
   time.Sleep(2*time.Second)
}



func add (a,b int) {
	var c =a+b
	fmt.Printf("%d + %d = %d \n",a,b,c)
}
