

//
Семафор до запуска горутины или внутри нее ?
Посмотреть эталонный вариант



### Задача 1.
// Задача 1 Что нужно сделать если нам нужно запускать всего по две за раз горутины вместо всех сразу

package main  

import (  
    "fmt"  
    "math/rand"
    "sync"    
    "time"
    )

type storage struct {  
    result map[int32]string  
    mux sync.Mutex  

}  

func LongWorker(res *storage, wg *sync.WaitGroup) {  
    time.Sleep(time.Duration(rand.Int31n(10)) * time.Second)  
    var key = rand.Int31n(100)  
    var value = fmt.Sprintf("some result: %d", rand.Int())  
    res.mux.Lock()  
    res.result[key] = value  
    res.mux.Unlock()  
}  

func main() {  
    s := &storage{result: make(map[int32]string)}  
    WorkersCount := rand.Int31n(12) + 1  

    var wg sync.WaitGroup  
    sem :=  make(chan struct{}, 2)

    for i := int32(0); i < WorkersCount; i++ {  
       sem <- struct{}{} // inside
       wg.Add(1)
       go func(){
        defer wg.Done
        LongWorker(s)
        <- sem // exit
       }()
    }  

    wg.Wait()  
    fmt.Println(s.result)  
}