package main

func main() {
	// 使用done取消通道时，你无法对通道的行为方式作出判断，也就是说你不知道正在执行读取操作的goroutine现在是什么状态，所以需要用select语句来封装我们的读取操作和done通道
	//  简单形式
	for val:=range myChan{
		// 对val进行处理
	}

	// 展开的形式
	loop:
		for{
			select {
			case <-done:
				break loop
			case maybeVal,ok:=<-myChan:
				if ok==false{
					return // or maybe break from for
				}
				// Do something with val
			}
		}

     // 封装orDone
     orDone:= func(done <-chan interface{},c <-chan interface{})<-chan interface{} {
     	valStream:=make(chan  interface{})
     	go func() {
     		defer close(valStream)
     		for{
				select {
				case <-done:
					return
				case v,ok:=<-c:
					if ok ==false{
						return
					}
					select {
					case valStream<-v:
					case <-done:
					}
				}
			}
		}()
     	return valStream
	 }

     // 回到简单的循环方式
     for val:=range orDone(done,myChan){
     	// Do something with val
	 }

}
