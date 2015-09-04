package test_components

import (
	"github.com/alexwang58/youhuo/components"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	apd := components.NewAutoInc(1)
	var nucant []int
	var wg sync.WaitGroup
	TestJobnum := 5
	wg.Add(TestJobnum)
	for i := 0; i < TestJobnum; i++ {
		go func() {
			s := rand.Intn(1)
			time.Sleep(time.Millisecond * time.Duration(s))
			id := apd.Id()
			//t.Logf("%v -> %v\n", i, id)
			nucant = append(nucant, id)
			wg.Done()
		}()
	}
	wg.Wait()

	defer func() {
		//		t.Log(nucant)
		sort.Ints(nucant)
		for j := 0; j < TestJobnum; j++ {
			if j+1 != nucant[j] {
				//t.Fatalf("%v nq %v\n", j+1, nucant[j])
				time.Sleep(time.Second * 1)
			}
		}
		t.Log("Ready to flush\n")
		if err := apd.FlushDat(); err != nil {
			panic(err)
		}
		//		t.Log(nucant)
	}()
}
