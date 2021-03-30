// Author: Taylor Privat

package main

import (
	"fmt"
	"strconv"
)

type Binding struct {
	name string
	value rune
}

func empty_env() []Binding {
	var newEnv = []Binding{}
	return newEnv
}

func extend_env(varn string, val rune, env []Binding) []Binding {
	var newEnv []Binding = []Binding{Binding{varn, val}}
	env = append(newEnv, env...)
	return env
}

func apply_env(env []Binding, search_var string) rune {
	if len(env) == 0 {
		fmt.Println("Error in apply_env(): Bad environment, ", env)
	}

	for i := range env {
		if env[i].name == search_var {
			return env[i].value
		}
	}

	fmt.Println("Error in apply_env(): No binding for ", search_var)
	return 0
}

func evaluate(root astNode, env []Binding) rune {
	var final_answer rune
	switch {
	case root.ttype == "LetExp":
		vname := root.children[0].contents
		e1 := evaluate(*root.children[1], env)
		var newEnv []Binding = []Binding{Binding{vname, e1}}
		env = append(newEnv, env...)
		final_answer = evaluate(*root.children[2], env)
		//return evaluate(*root.children[2], env)
		break
	case root.ttype == "IfExp":
		e1 := evaluate(*root.children[0], env)
		e2 := evaluate(*root.children[1], env)
		e3 := evaluate(*root.children[2], env)
		if e1 > 0 {
			final_answer = e2
		} else {
			final_answer = e3
		}
		break
	case root.ttype == "ZeroExp":
		// Returns 1 if true, 0 otherwise
		e := evaluate(*root.children[0], env)
		if e == 0 {
			final_answer = 1
		} else {
			final_answer = 0
		}
		break
	case root.ttype == "DiffExp":
		e1 := evaluate(*root.children[0], env)
		e2 := evaluate(*root.children[1], env)
		// May need to use strconv.Atoi() here
		final_answer = e1 - e2
		break
	case root.ttype == "ConstExp":
		c, _ := strconv.Atoi(root.contents)
		final_answer = int32(c)
		break
	case root.ttype == "Var":
		var name = root.contents
		value := apply_env(env, name)
		final_answer = value
		break
	default:
		fmt.Println("Error in evaluate(): Unrecongnized language construct -- ", root.ttype)
		final_answer = 0
	}
	return final_answer
}
