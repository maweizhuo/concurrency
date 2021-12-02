package main

import "fmt"

func main() {
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
	
	// 有时候会需要一系列通道的值    <-chan <-chan interface{}
	bridge:= func(done <-chan interface{},chanStream <-chan  <-chan interface{})<-chan interface{} {
		valStream:=make(chan interface{})
		go func() {
			defer close(valStream)
			for{
				var stream <-chan interface{}
				select {
				case maybeStream,ok:=<-chanStream:
					if ok==false{
						return
					}
					stream=maybeStream
				case <-done:
					return
				}
				for val:=range orDone(done,stream){
					select {
					case valStream<-val:
					case <-done:
					}
				}

			}
		}()
		return valStream
	}

	// 使用它， 创建10个通道，每个通道写入一个元素，并将这些通道传入给bridge
	genVals:= func()<-chan <-chan interface{} {
         chanStream:=make(chan (<-chan interface{}))
         go func() {
         	defer close(chanStream)
         	for i:=0;i<10;i++{
         		stream:=make(chan interface{},1)
         		stream<-i
         		close(stream)
         		chanStream<-stream
			}
		 }()
         return chanStream
	}
	for val:=range bridge(nil,genVals()){
		fmt.Printf("%v ",val)
	}

}
