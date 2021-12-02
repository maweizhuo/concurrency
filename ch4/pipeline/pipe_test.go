package main

import "testing"

// go test -bench=. ch4/pipeline/pipe_test.go
func BenchmarkGeneric(b *testing.B)  {
	take:= func(done <-chan interface{},valueStream <-chan interface{},num int)<-chan interface{}{
		takeStream:=make(chan interface{})
		go func() {
			defer close(takeStream)
			for i:=0;i<num;i++{
				select {
				case <-done:
					return
				case takeStream<-<-valueStream:
				}
			}
		}()
		return takeStream
	}
	// 断言
	toString:= func(done <-chan interface{},valueStream <-chan interface{})<-chan string {
		stringStream:=make(chan string)
		go func() {
			defer close(stringStream)
			for v:=range valueStream{
				select {
				case <-done:
					return
				case stringStream<-v.(string):
				}
			}
		}()
		return stringStream
	}
	repeat:= func(done <-chan interface{},values ...interface{})<-chan interface{} {
		valueStream:=make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _,v:=range values{
					select {
					case <-done:
						return
					case valueStream<-v:
					}
				}
			}
		}()
		return  valueStream
	}

	done:=make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range toString(done,take(done,repeat(done,"a"),b.N)){
		
	}
}

// 已知类型的比默认的多两倍
func BenchmarkTyped(b *testing.B)  {
	repeat:= func(done <-chan interface{},values ...string)<-chan string {
      valueStream:=make(chan string)
      go func() {
		  defer close(valueStream)
		  for{
			  for _,v:=range values{
				  select {
				  case <-done:
					  return
				  case valueStream<-v:
				  }
			  }
		  }
	  }()
      return valueStream
	}
	take:= func(done <-chan interface{},valueStream <-chan string,num int)<-chan string {
       takeStream:=make(chan string)
       go func() {
       	defer close(takeStream)
       	for i:=num;i>0|| i==-1;{
       		if i!=-1{
       			i--
			}
			select {
       		 case <-done:
				 return
			case takeStream<-<-valueStream:
			}
		}
	   }()
       return takeStream
	}
	done:=make(chan interface{})
	defer close(done)
	b.ResetTimer()
	for range take(done,repeat(done,"a"),b.N){

	}
}