package main

/*func main() {
	var x = 1
	var ch = make(chan int, 1)
	ch <- x          // 将变量x发送给管道ch
	var y int = <-ch // 变量y接受管道ch发送过来的值
	//<-ch // 接受管道ch发送过来的值，但不使用
	fmt.Println(y)
}*/

/*func main() {
	var x = 1
	var ch = make(chan int, 1)
	ch <- x // 将变量x发送给管道ch
	//var y int = <-ch // 变量x接受管道ch发送过来的值
	<-ch // 接受管道ch发送过来的值，但不使用
	var y int
	fmt.Println(y)
}*/

/*func main() {
	var channel = make(chan int)
	var send = 6666
	// 关键字go后跟的是需要被并发执行的代码块，它由一个匿名函数代表
	// 在这里，我们只要知道在花括号中的就是将要并发执行的代码就可以
	go func() {
		channel <- send
		fmt.Println("数据已发送 ") // 这个语句在这个例子里能不能执行，为什么呢？
	}()
	// 这里让主线程休眠1秒钟
	// 以使上面的goroutine有机会被执行
	time.Sleep(1 * time.Second)
}*/

/*func main() {
var channel = make(chan int)
var send = 6666
var receive int
go func() {
	// 向channel中传递sent的值*/

/*		for i := 1; i <= 100; i++ {
			send++
			channel <- send
			fmt.Println("数据以传输")
		}

	}()
	// 使用receive接受从channel中传来的值
	receive = <-channel
	receive = <-channel
	fmt.Println(receive)
}*/
