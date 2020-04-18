package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"git.sr.ht/~mna/sendkeys"
)

func main() {
	var (
		flagBytes = flag.Bool("-bytes", false, "Send bytes instead of runes.")
		flagTTY   = flag.String("-tty", "", "Path of the target tty.")
		flagDelay = flag.Duration("-delay", 100*time.Millisecond, "Delay between each string to send.")
	)
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 || *flagTTY == "" {
		flag.Usage()
		return
	}

	args := flag.Args()
	for i, s := range args {
		v, err := strconv.Unquote(`"` + s + `"`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid argument when treated as a Go double-quoted string: %s", s)
			os.Exit(1)
		}
		args[i] = v
	}

	t, err := sendkeys.Open(*flagTTY, *flagDelay)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer t.Close()

	if *flagBytes {
		_, err = t.SendBytes(args...)
	} else {
		_, err = t.SendRunes(args...)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() {
	const msg = `usage: sendkeys -tty PATH [-delay DUR] [-bytes] STRING...
Each STRING is treated as if it was a Go double-quoted string, so that
e.g. "\x1b" is treated as an escape sequence.
`
	fmt.Println(msg)
}
