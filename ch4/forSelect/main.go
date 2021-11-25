package main

func main() {

	for  { // 无限循环或遍历
		select {
          // 对通道进行操作
        }
	}

   // 在通道上发送变量
   for _,s:=range []string{"a","b","c"}{
	   select {
	   case <-done:
		   return
	   case stringStream<-s:
	 }
   }

   // 无限循环等待停止
   for {
	   select {
	   case <-done:
		   return
	   default:
	}
	   // 执行非抢占任务
   }

   // 进入select语句时，如果done通道尚未关闭，我们将执行default子句。
   for{
	   select {
   	  case <-done:
		  return
	   default:
		   // 执行非抢占任务
	}
   }

}
