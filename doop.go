package main

import (
	"os"
)

var op = []string{"+", "-", "*", "/", "%"}

//dont copy unless you can explain it and know its flaws
func main() {
	args := os.Args[1:]
	var result int
	switch {
	case len(args) < 3 || len(args) > 3:
		return
	case opCheck(args[0]) || opCheck(args[2]):
		return
	case !opCheck(args[1]):
		return
	case !validcharCheck(args[0]) || !validcharCheck(args[2]):
		return
	}
	arg1 := sToI(args[0])
	arg3 := sToI(args[2])
	arg2 := args[1]
	switch {
	case arg2 == op[4] && arg3 == 0:
		os.Stdout.WriteString("No modulo by 0\n")
	case arg2 == op[3] && arg3 == 0:
		os.Stdout.WriteString("No division by 0\n")
	}
	switch arg2 {
	case op[0]:
		result = arg1 + arg3
		printcalc(result)
	case op[1]:
		result = arg1 - arg3
		printcalc(result)
	case op[2]:
		result = arg1 * arg3
		printcalc(result)
	case op[3]:
		result = arg1 / arg3
		printcalc(result)
	default:
		result = arg1 % arg3
		printcalc(result)
	}
}

func printcalc(sum int) {
	switch {
	case overFlow(sum):
		return
	case sum == 0:
		os.Stdout.WriteString("0\n")
	default:
		os.Stdout.WriteString(iToS(sum))
		os.Stdout.WriteString("\n")
	}
}

func opCheck(arg string) bool {
	for _, s := range op {
		if s == arg {
			return true
		}
	}
	return false
}

//shitty overflow
func overFlow(sum int) bool {
	if (sum > 999999999999999999) || (sum < -999999999999999999) {
		return true
	}
	return false
}

func validcharCheck(s string) bool {
	runes := []rune(s)

	for i := 0; i <= len(runes)-1; i++ {
		switch {
		case runes[i] >= 33 && runes[i] <= 36:
			return false
		case runes[i] >= 38 && runes[i] <= 41:
			return false
		case runes[i] == 44:
			return false
		case runes[i] == 46:
			return false
		case runes[i] >= 58 && runes[i] <= 127:
			return false
		}
	}
	return true
}

func sToI(s string) int {
	result := 0
	bytes := []byte(s)
	for _, v := range bytes {
		if v == '-' {
			continue
		}
		v -= '0'
		result = result*10 + int(v)
	}
	if s[0] == '-' {
		result = -result
	}
	return result
}

func iToS(n int) string {
	result := ""
	isneg := false
	if n < 0 {
		isneg = true
		n = -n
	}
	for n >= 1 {
		result = string(n%10+'0') + result
		n /= 10
	}
	if isneg {
		result = string("-") + result
	}
	return result
}
