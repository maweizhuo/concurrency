package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// context 包连个主要目的
	// 1. 提供取消操作 2. 提供用于通过调用传输请求附加数据的数据包

   var wg sync.WaitGroup
   done:=make(chan interface{})
   defer close(done)
   wg.Add(1)
   go func() {
   	defer wg.Done()
   	if err:=printGreeting(done);err !=nil{
   		fmt.Printf("%v",err)
		return
	}
   }()
   wg.Add(1)
  go func() {
  	defer wg.Done()
  	if err:=printFarewell(done);err!=nil{
  		fmt.Printf("%v",err)
		return
	}
  }()
 wg.Wait()

}

func printGreeting(done <-chan interface{}) error {
	greeting,err:=genGreeting(done)
	if err !=nil{
		return err
	}
	fmt.Printf("%s world!\n",greeting)
    return nil
}

func printFarewell(done <-chan interface{})error  {
	farewell,err:=genFarewell(done)
	if err !=nil{
		return err
	}
	fmt.Printf("%s world!\n",farewell)
	return nil
}

func genGreeting(done <-chan interface{})(string,error)  {
	switch locale,err:=locale(done);{
	  case err!=nil:
	  	return "",err
	case locale=="EN/US":
		return "hello",nil
	}
	return "",fmt.Errorf("unsupported locale")
}

func genFarewell(done <-chan interface{})(string,error)  {
	switch locale,err:=locale(done); {
	case err !=nil:
		return "",err
	case locale =="EN/US":
       return "goodbye",nil
	}
	return "",fmt.Errorf("unsuppored locale")
}

func locale(done <-chan interface{})(string,error)  {
	select {
	case <-done:
		return "",fmt.Errorf("canceled")
	case <-time.After(5*time.Second):
	}
	return "EN/US",nil
}