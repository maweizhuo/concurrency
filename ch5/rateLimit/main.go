package main

import (
	"context"
	"log"
	"os"
	"sync"
)

func Open() *APIConnection {
	return &APIConnection{}
}

type APIConnection struct {}

func (a *APIConnection)ReadFile(ctx context.Context)error  {
  // Pretended we do work here
  return nil
}

func (a *APIConnection)ResolveAddress(ctx context.Context)error  {
  // pretended we do work here
  return nil
}

func main() {
	defer log.Println("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime|log.LUTC)

	apiConnection:=Open()
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
