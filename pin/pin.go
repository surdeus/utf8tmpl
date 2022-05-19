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
	nVal int
	chrs []rune
)

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
		pin := GetPin(s, l, i)
		if Fits([]rune(pin)) {
			fmt.Println(GetPin(s, l, i))
		}
	}
}

func
Fits(s []rune) bool {
	a := make([]int, len(chrs))
	for i, v1 := range chrs {
		for _, v2 := range s {
			if v1 == v2 {
				a[i]++
				if a[i] > nVal {
					return false
				}
			}
		}
	}
	return true
}

func
Run(args []string) int {
	arg0 = args[0]
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.IntVar(&lval, "l", 0,
		"Add combinations with less number of chars starting with arg.")
	flagSet.BoolVar(&Lflag, "L", false,
		"Set less number to 1. Overrides l flag.")
	flagSet.IntVar(&nVal, "n", 1, "Max repeats of the rune." )
	flagSet.Parse(args[1:])
	args = flagSet.Args()
	flagSet.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [options] <tmpl_chars> [len]\n", arg0)
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	var(
		err error
		n int
	)

	if len(args) == 0 {
		flagSet.Usage()
	}

	chrs = []rune(args[0])
	if len(args) == 2 {
		n, err = strconv.Atoi(args[1])
	} else if len(args) == 1 {
		n = len(chrs)
	}

	if err!=nil {
		log.Fatal(err)
	}

	if Lflag {
		lval = 1
	}

	if lval != 0 {
		if lval > n {
			flagSet.Usage()
		}
		for i := lval ; i<=n ; i++ {
			printPins(chrs, i)
		}
	} else {
		printPins(chrs, n)
	}

	return status
}
