package pin

import(
	"os"
	"fmt"
	"strconv"
	"log"
	"flag"
)

var(
	Lflag bool
	lval int
	arg0 string
	delim rune = '\n'
	status int = 0
)


func
usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [n_utf8_chars] [len]\n", arg0)
	os.Exit(1)
}

func
Pow(x, p int) int {
	ret := 1 
	for i:=0 ; i<p ; i++ {
		ret *= x
	}
	return ret
}

func
GetPin(s []rune, l int, i int) string {
	ret := ""
	slen := len(s)
	for j:=0 ; j<l ; j++ {
		ret = string(s[ ( i/Pow(slen, j) ) % slen] ) + ret
	}
	return ret
}

func
printPins(s []rune, l int) {
	n := Pow(len(s), l)
	for i:=0 ; i<n ; i++ {
		fmt.Println(GetPin(s, l, i))
	}
}

func
Run(args []string) int {
	arg0 = args[0]
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.IntVar(&lval, "l", 0,
		"Add combinations with less number of chars starting with arg.")
	flagSet.BoolVar(&Lflag, "L", false,
		"Set less number to 1. Overrides l flag.")
	flagSet.Parse(args[1:])
	args = flagSet.Args()
	if len(args) < 2 {
		usage()
	}

	chrs := []rune(args[0])
	n, err := strconv.Atoi(args[1])
	if err!=nil {
		log.Fatal(err)
	}

	if Lflag {
		lval = 1
	}

	if lval != 0 {
		if lval > n {
			usage()
		}
		for i := lval ; i<=n ; i++ {
			printPins(chrs, i)
		}
	} else {
		printPins(chrs, n)
	}

	return status
}
