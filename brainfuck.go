package main

import (
	"fmt"
)

//表达brainfuck使用的机器模型，连续字节内存块
type tape struct {
	mem []byte
	pos int
}

func TapeNew() *tape {
	t := new(tape)
	t.mem = make([]byte, 4096)
	return t
}

//.操作， putchar(*ptr)
func (t *tape) get() byte { return t.mem[t.pos] }

//,操作， *ptr = getchar()
func (t *tape) set(val byte) { t.mem[t.pos] = val }

//+操作， ++*ptr
func (t *tape) inc() { t.mem[t.pos]++ }

//-操作， --*ptr
func (t *tape) dec() { t.mem[t.pos]-- }

//>操作,  ++ptr
func (t *tape) forward() {
	t.pos++
	if len(t.mem) <= t.pos {
		t.mem = append(t.mem, 0)
	}
}

//<操作，--ptr
func (t *tape) backward() { t.pos-- }

func interpret(prog string, whilemap map[int]int) {
	pc := 0
	tape := TapeNew()
	var tmp byte
	for pc < len(prog) {
		switch prog[pc] {
		case '>':
			tape.forward()
		case '<':
			tape.backward()
		case '+':
			tape.inc()
		case '-':
			tape.dec()
		case '.':
			c := tape.get()
			fmt.Printf("%c", c)
		case ',':
			fmt.Scanf("%c", &tmp)
			tape.set(tmp)
		case '[':
			if tape.get() == 0 {
				pc = whilemap[pc]
			}
		case ']':
			if tape.get() != 0 {
				pc = whilemap[pc]
			}

		}
		pc++
	}
}

func parse(prog string) (string, map[int]int) {
	parsed := make([]byte, 0)
	pcstack := make([]int, 0)
	//记录[,对应的],索引(指令)位置
	whilemap := make(map[int]int, 128)
	pc := 0
	for _, char := range prog {
		//fmt.Printf("got char: %c\n", char)
		switch char {
		case '>', '<', '+', '-', '.', ',', '[', ']':
			parsed = append(parsed, byte(char))
			if char == '[' {
				pcstack = append(pcstack, pc)
			} else if char == ']' {
				last := len(pcstack) - 1
				left := pcstack[last]
				pcstack = pcstack[:last]
				right := pc
				whilemap[right] = left
				whilemap[left] = right
			}
			pc++
		}
	}
	return string(parsed), whilemap
}
