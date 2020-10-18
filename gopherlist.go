package grmonitor

import (
	"fmt"
	"log"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)
// GRMonitor runs a routine that checks the goroutines every poll seconds for any
// additions or changes if the showDiffs flag is true
func Run(match string, poll int, showDiffs bool) {
	// Goroutine name -> instance count
	var routines map[string]int64
	go func() {
		for {
			var w strings.Builder
			r := make(map[string]int64)
			pprof.Lookup("goroutine").WriteTo(&w, 1)
			var title, last, lastMatch string
			s := strings.Split(w.String(), "\n")
			for _, a := range s {
				l := strings.TrimSpace(a)

				if l == "" {
					if title != "" && (last != "" || lastMatch != "") {
						if last == "" {
							last = lastMatch
						}
						x := strings.Split(last, "\t")
						last = x[2]
						plus := strings.Index(last, "+0x")
						last = last[:plus]

						at := strings.Index(title, "@")
						i, _ := strconv.ParseInt(title[:at-1], 10, 64)
						r[last] = i
					}
					last = ""
					title = ""
					continue
				}

				if title != "" {
					p := strings.Index(l, match)
					if p > 0 {
						lastMatch = l
						//fmt.Printf("GOT %s --> %s\n", z, title)
					}
					last = l
				}

				if l[0] != '#' {
					title = l
					continue
				}
			}
			if showDiffs {
				if routines != nil {
					// Anything gone?
					for a := range routines {
						if _, ok := r[a]; !ok {
							log.Printf("Routine %s ended", a)
						}
					}
					// Anything added
					for a := range r {
						if _, ok := routines[a]; !ok {
							log.Printf("Routine %s added %d", a, r[a])
						} else if routines[a] != r[a] {
							log.Printf("Routine %s count changed from %d to %d", a, routines[a], r[a])
						}
					}
				}
			}
			routines = r
			fmt.Println("\nGoRoutine Dump")
			for a, b := range routines {
				fmt.Printf("\t %d ==> %s\n", b, a)
			}
			time.Sleep(time.Duration(poll) * time.Second)
		}
	}()
}
