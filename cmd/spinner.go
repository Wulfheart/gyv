package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

func spin(processMsg string, doneMsg string, f func() error) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprint("  ", processMsg)
	s.Start()
	err := f()
	s.Stop()
	if err != nil {
		return err
	}
	fmt.Println("✅ ", doneMsg)
	return nil
}
