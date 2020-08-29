package pin

import(
	"os"
	"fmt"
	"strconv"
	"log"
)

var(
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
		ret += string(s[ ( i/Pow(slen, j) ) % slen] )
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

	if len(args) < 3 {
		usage()
	}

	chrs := []rune(args[1])
	n, err := strconv.Atoi(args[2])
	if err!=nil {
		log.Fatal(err)
	}

	printPins(chrs, n)

	return status
}
