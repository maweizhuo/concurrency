package main

import "fmt"

func main() {
	// 管道  批处理
	//multiply:= func(values []int,multiplier int)[]int {
	//	multipliedValues:=make([]int,len(values))
	//	for i,v:=range values{
	//		multipliedValues[i]=v*multiplier
	//	}
	//	return multipliedValues
	//}
	//add:= func(values []int,additive int)[]int {
	//	addedValues:=make([]int,len(values))
	//	for i,v:=range values{
	//		addedValues[i]=v+additive
	//	}
	//	return addedValues
	//}
	//// 合并上面的乘法加法
	//ints:=[]int{1,2,3,4}
	////for _,v:=range add(multiply(ints,2),1){
	////	fmt.Println(v)
	////}
	//// 在加个阶段，*2
	//for _,v:=range multiply(add(multiply(ints,2),1),2){
	//	fmt.Println(v)
	//}

	// 粗暴方式
	//ints:=[]int{1,2,3,4}
	//for _,v:=range ints{
	//	fmt.Println(2*(v*2+1))
	//}
	//

   // 流操作
   multiply:= func(value,multiplier int)int {
   	 return value*multiplier
   }
   add:= func(value,additive int)int{
   	return value+additive
   }
   ints:=[]int{1,2,3,4}
   for _,v:=range ints{
   	fmt.Println(multiply(add(multiply(v,2),1),2))
   }

}
