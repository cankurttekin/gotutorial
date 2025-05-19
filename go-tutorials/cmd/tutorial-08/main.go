package main
import (
	"fmt"
	//"math/rand"
	"time"
	"sync"
)

//var m = sync.Mutex{} // Mutex to protect the shared resource
var m = sync.RWMutex{} // we also have read lock and unlock
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

/*
func main() {
	t0 := time.Now()
	for i := 0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\ntotal execution time: %v\n", time.Since(t0))
	fmt.Printf("the results are: %v\n", results)
}


func dbCall(i int) {
	// Simulate a database call delay	
	//var delay float32 = rand.Float32()*2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Printf("the result from db is: %s\n", dbData[i])
	//m.Lock() // Lock the mutex before accessing the shared resource
	//results = append(results, dbData[i])
	//m.Unlock() // Unlock the mutex after accessing the shared resource
	//wg.Done()

	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("the current results are %v\n", results)
	m.RUnlock()
}
*/


func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\ntotal execution time: %v\n", time.Since(t0))
}

func count() {
	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	wg.Done()
}























