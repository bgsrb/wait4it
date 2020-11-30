package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"
	"wait4it/checker"
	"wait4it/config"
)

const invalidUsageStatus = 2

func findCheckModule(ct string) (interface{}, error) {
	m, ok := modules[ct]
	if !ok {
		return nil, errors.New("unsupported check type")
	}

	return m, nil
}

func ticker(cs checker.Checker, t *time.Ticker, d chan bool) {
	for {
		select {
		case <-d:
			return
		case <-t.C:
			check(cs)
		}
	}
}

func check(cs checker.Checker) {
	r, eor, err := cs.Check()
	if err != nil && eor {
		wStdErr(err.Error())
		os.Exit(invalidUsageStatus)
	}

	wStdOut(r)
}

func wStdOut(r bool) {
	if r {
		_, _ = fmt.Fprintln(os.Stdout, "succeed")
		os.Exit(0)
	} else {
		_, _ = fmt.Fprint(os.Stdout, ".")
	}
}

func wStdErr(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
}

//Run ...
func Run(c config.CheckContext) {
	m, err := findCheckModule(c.Config.CheckType)
	if err != nil {
		wStdErr(err)
		os.Exit(invalidUsageStatus)
	}

	cx := m.(checker.Checker)

	cx.BuildContext(c)
	err = cx.Validate()
	if err != nil {
		wStdErr(err)
		os.Exit(invalidUsageStatus)
	}

	fmt.Print("Wait4it...")

	t := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go ticker(cx, t, done)

	time.Sleep(time.Duration(c.Config.Timeout) * time.Second)
	done <- true

	fmt.Println("failed")
	os.Exit(1)
}
