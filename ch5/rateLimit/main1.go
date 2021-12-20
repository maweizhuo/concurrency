package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
)

// Limit定义了事件的最大频率
// Limit被表示为每秒事件的数量
// 值为0的Limit不允许任何事件
type Limit float64

// NewLimiter 返回一个新的Limiter实例
// 发生率为r，并允许至多b个令牌爆发
//func NewLimiter(r Limit,b int) *Limiter

func Open1() *APIConnection1 {
	return &APIConnection1{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type APIConnection1 struct {
	rateLimiter *rate.Limiter
}

func (a *APIConnection1) ReadFile(ctx context.Context)error  {
   if err:=a.rateLimiter.Wait(ctx);err!=nil{
	   return err
   }
   // Pretend we do work hre
	return nil
}

func (a *APIConnection1)ResolveAddress(ctx context.Context)error  {
	if err:=a.rateLimiter.Wait(ctx);err!=nil{
		return err
	}
	// Pretend we do work hre
	return nil
}

func main() {
	defer log.Println("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime|log.LUTC)

	apiConnection:=Open1()
	var wg sync.WaitGroup
	wg.Add(20)

	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			err:=apiConnection.ReadFile(context.Background())
			if err !=nil{
				log.Printf("cannot Readfile:%v",err)
			}
		}()
	}
	log.Printf("ReadFile")

	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			err :=apiConnection.ResolveAddress(context.Background())
			if err !=nil{
				log.Printf("cannnot Resolve Address: %v",err)
			}
		}()
	}
	log.Printf("ResolveAddress")
	wg.Wait()
}
