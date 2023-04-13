package main

import (
	"bufio"
	_ "bufio"
	_ "bytes"
	"io"
	_ "io"
	"os"
	_ "os"
	"reflect"
	"strings"
	_ "strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		res, msg := isPrime(e.testNum)
		if e.expected && !res {
			t.Errorf("'%s': '%t' expected, but got '%t'", e.name, res, !res)
		}

		if !e.expected && res {
			t.Errorf("'%s': '%t' expected, but got '%t'", e.name, !res, res)
		}

		if e.msg != msg {
			t.Errorf("'%s': '%s' expected, but got '%s'", e.name, e.msg, msg)
		}
	}
}

func TestPrompt(t *testing.T) {
	expectedValue := "-> "

	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()
	w.Close()
	os.Stdout = temp

	in, _ := io.ReadAll(r)
	out := string(in)

	if !(out == expectedValue) {
		t.Errorf("'%s' expected, but got '%s'", expectedValue, out)
	}
}

func TestIntro(t *testing.T) {
	expectedValue := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "
	temp := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()
	w.Close()
	os.Stdout = temp

	in, _ := io.ReadAll(r)
	out := string(in)

	if out != expectedValue {
		t.Errorf("'%s' expected, but got '%s'", expectedValue, out)
	}
}

func TestCheckNumbers(t *testing.T) {
	values := "q\ncb\n6\n7\n"
	sc := bufio.NewScanner(strings.NewReader(values))

	out, res := checkNumbers(sc)
	if out != "" || res == false {
		t.Errorf("Unexpected result for 'q' input: output='%q', result='%v'", out, res)
	}

	out, res = checkNumbers(sc)
	if out != "Please enter a whole number!" || res == true {
		t.Errorf("Unexptected result for numeric number: output='%q', result='%v'", out, res)
	}

	out, res = checkNumbers(sc)
	if strings.Contains(out, "not a prime number") == false || res == true {
		t.Errorf("Unexpected result for non prime numbers: output:'%q', result='%v'", out, res)
	}

	out, res = checkNumbers(sc)
	if strings.Contains(out, "is a prime number") == false || res == true {
		t.Errorf("Unexpected result for non prime numbers: output:'%q', result='%v'", out, res)
	}
}

func TestReadUserInput(t *testing.T) {
	values := "5\nq\n"
	sc := bufio.NewReader(strings.NewReader(values))
	curChan := make(chan bool)

	go readUserInput(sc, curChan)

	var output []string

	for res := range curChan {
		if res {
			return
		}
		str, _, err := sc.ReadLine()
		if err != io.EOF {
			return
		}
		output = append(output, string(str))
	}

	expectedValue := []string{"7 is a prime number", "Goodbye."}
	if !reflect.DeepEqual(expectedValue, output) {
		t.Errorf("Expected '%v', but got '%v'", expectedValue, output)
	}
}
