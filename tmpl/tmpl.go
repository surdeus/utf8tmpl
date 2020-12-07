package tmpl

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

var(
	arg0 string
	delim rune
	status int
)


func
usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [n_utf8_chars] [n_strings]\n", arg0)
	os.Exit(1)
}

func
Run(args []string) int {
	status = 0
	delim = '\n'
	arg0 = args[0]
	
	if len(args)<3 || len(args[1]) != len(args)-2 {
		usage()
	}

	chrs := []rune(args[1])
	args = args[2:]

	tmpl := make(map[rune] string)
	for i, s := range args {
		tmpl[rune(chrs[i])] = s
	}

	r := bufio.NewReader(os.Stdin)
	for{
		s, e := r.ReadString(byte(delim))
		if e==io.EOF {
			break
		}
		s = s[:len(s)-1]
		for _, c := range []rune(s) {
			s, ok := tmpl[c]
			if !ok {
				/*fmt.Fprintf(os.Stderr, "%s: '%s': no such character in template string\n",
					arg0, string(c) )*/
				s = string(c)	
			}
			fmt.Printf("%s", s)
		}
		fmt.Printf("%s", string(delim))
	}


	return status
}
