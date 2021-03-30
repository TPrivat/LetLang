// Author: Taylor Privat

package main

import (
    "fmt"
)

func convertType(t int) string {
	switch t {
	case 10:
		return "ConstExp"
	case 11:
		return "Var"
	case 16:
		return "LetExp"
	case 17:  // IN
		return "Body"
	case 18:
		return "IfExp"
	case 19:
		return "Then"
	case 20:
		return "Else"
	case 21:
		return "Equals"
	case 22:
		return "Comma"
	case 23:
		return "DiffExp"
	case 24:
		return "ZeroExp"
	default:
		return "Paren"
	}
}

func initTreeNode(node *astNode, isterm bool) {
	node.termsym = isterm
	node.children = make([]*astNode, 0, 5)
	node.ttype = convertType(tokenQueue[0].tokenType)
	node.contents = tokenQueue[0].tokenValue
	tokenQueue = tokenQueue[1:]
}

func checkToken(tok string) {
	if (tokenQueue[0].tokenValue != tok) {
		fmt.Println("Grammer error! Expected: ", tok, " but got ", tokenQueue[0].tokenValue)
	} else { tokenQueue = tokenQueue[1:] }
}

func parseExp() astNode {
	root := astNode{}
	switch tokenQueue[0].tokenType {
	case LET:
		initTreeNode(&root, false)
		child1 := parseExp()
		checkToken("=")
		child2 := parseExp()
		checkToken("in")
		child3 := parseExp()
		root.children = append(root.children, &child1)
                root.children = append(root.children, &child2)
                root.children = append(root.children, &child3)
	case IF:
		initTreeNode(&root, false)
		child1 := parseExp()
		checkToken("then")
		child2 := parseExp()
		checkToken("else")
		child3 := parseExp()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)
	case ISZERO:
		initTreeNode(&root, false)
		checkToken("(")
		child1 := parseExp()
		checkToken(",")
		child2 := parseExp()
		checkToken(")")
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
	case MINUS:
		initTreeNode(&root, false)	// returns next token after init root
		checkToken("(")			// check for (
		child1 := parseExp()		// pops tokenQueue
		checkToken(",")			// check for ,
		child2 := parseExp()		// pops tokenQueue
		checkToken(")")			// check for )
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		break
	case INT_LIT:
		initTreeNode(&root, true)
		break
	case IDENT:
		initTreeNode(&root, true)
		break
	}
	return root
}

func printTree(node *astNode, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("   ")
	}
	fmt.Println(node.ttype, "(")

	// if leaf print the contents too
	if (node.termsym) {
		for i:= 0; i<depth+4; i++ {
			fmt.Print("  ")
		}
		fmt.Println(node.contents)
		for i:= 0; i<depth; i++ {
			fmt.Print("   ")
		}
		fmt.Println("),")
	}

	// print children too
	for (!node.termsym) {
		if (len(node.children) >= 1) {
			for _, child := range node.children {
				printTree(child, depth+1)
			}
		}
		for i:=0;i<depth;i++ {
			fmt.Print("   ")
		}
		fmt.Println("),")
		node.termsym = true
	}
}
