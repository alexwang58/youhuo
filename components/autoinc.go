// Generate auto increated id
// user for mongodb uid
// This struct is newd befor http server is started
// To get the auto increased id, just call func(ai *AutoInc)Id()int
// features:
// 	-number auto generate from chanel(Queen)
//	-LastGenerateId will flush to disk automaticly at a specific interval time
//
package components

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// Size of the number queen
	AUTOINC_QUEEN_SIZE = 100
	AUTOINC_PSN_TYP    = "PSN"
	AUTOINC_PRD_TYP    = "PRD"
	// Interval time which flush the last id to disk,
	// the id which had been taken from client user
	// We call the number LastGenerateId
	AUTOINC_FLUSH_INTERVAL_SEC = 30
	// Store the LastGenerateId
	AUTOINC_PSN_DAT_FILE = "./dat/psn_autoinc.dat"
	AUTOINC_PRD_DAT_FILE = "./dat/prd_autoinc.dat"
)

type AutoInc struct {
	start, step, curid int
	queue              chan int
	running            bool
	dfile              string
	timer              *time.Ticker
}

func NewAutoInc(step int, typ string) (ai *AutoInc) {
	dfile := AUTOINC_PSN_TYP
	if typ == AUTOINC_PRD_TYP {
		dfile = AUTOINC_PRD_TYP
	}
	_curid := loadDat(dfile)
	ai = &AutoInc{
		start:   _curid + 1,
		step:    step,
		curid:   _curid,
		running: true,
		queue:   make(chan int, AUTOINC_QUEEN_SIZE),
		dfile:   dfile,
	}
	fmt.Println("[Autoinc Strat From] :" + strconv.Itoa(ai.start))
	ai.setTimer().flushProcess()
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	ai.curid = <-ai.queue
	fmt.Printf("Current : %v\n", ai.curid)
	return ai.curid
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

func (ai *AutoInc) setTimer() *AutoInc {
	ai.timer = time.NewTicker(time.Second * AUTOINC_FLUSH_INTERVAL_SEC)
	return ai
}

func (ai *AutoInc) flushProcess() {
	go func() {
		for {
			select {
			case <-ai.timer.C:
				ai.FlushDat()
			}
		}
	}()
}

func loadDat(dfile string) int {
	file, err := os.OpenFile(dfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	str := ""
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		str += string(buf[:n])
	}
	str = strings.TrimSpace(str)
	fmt.Println("[Autoinc Load]:" + str)
	if str == "" {
		return 0
	}
	var rint int
	rint, err = strconv.Atoi(str)
	if err != nil {
		panic("Autoinc data file [" + dfile + "] irregular: " + strconv.Itoa(rint))
	}

	return rint
}

func (ai *AutoInc) FlushDat() error {
	//fmt.Printf("AutoicFlushDat to file: %v\n", ai.dfile)
	if ai.curid < ai.start {
		//fmt.Printf("AutoicNoNeedFlush  %v(start)-%v(current)\n", ai.start, ai.curid)
		return nil
	}
	file, err := os.OpenFile(ai.dfile, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	//fmt.Printf("AutoicFlushDat : %v @ %v/sec\n", ai.curid, AUTOINC_FLUSH_INTERVAL_SEC)
	if _, err = file.WriteString(strconv.Itoa(ai.curid)); err != nil {
		return err
	}

	return nil
}
