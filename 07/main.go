package main

import (
	"aoc_2022/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Some constants we need
const totalSpace = 70000000
const spaceNeeded = 30000000
const smallDirectoryLimit = 100000

func main() {
	start := time.Now()

	tree := parseTree()

	freeSpace := totalSpace - tree.Root.Value
	neededSpace := spaceNeeded - freeSpace

	part1, part2 := findSum(tree.Root, neededSpace, tree.Root.Value)

	fmt.Println(part1, part2)

	tree.Print()

	fmt.Println("Binomial took: ", time.Since(start))
}

func findSum(node *utils.Node, spaceNeeded int, spaceToDelete int) (int, int) {
	// define our collecting variables
	var sum int
	var space = spaceToDelete

	// if the node value is less than the limit we want to add this Value to our tree
	if node.Value <= smallDirectoryLimit {
		sum += node.Value
	}

	// if the node is more than the space needed and less than the spaceToDelete we add the node.Value to space
	if node.Value >= spaceNeeded && node.Value < space {
		space = node.Value
	}

	for _, d := range node.Children {
		// we dont need to go over any files because the folders already have the correct value for each folder
		if d.Kind == "dir" {
			x, y := findSum(d, spaceNeeded, space)

			sum += x
			space = y
		}

	}

	return sum, space
}

func parseTree() *utils.Tree {
	root := &utils.Node{Value: 0, Parent: nil, Name: "root", Kind: "dir", Children: make(map[string]*utils.Node)}

	tree := &utils.Tree{Root: root}

	scanner := utils.ReadFileAsScanner("input")

	currentNode := root

	for scanner.Scan() {
		if scanner.Text() == "$ cd /" {
			// This is an edge case where we always want to return to the root ('/') of the file system
			currentNode = tree.Root
			continue
		}

		if scanner.Text() == "$ ls" {
			// '$ ls' can be ignored by our parser
			continue
		}

		data := strings.Split(scanner.Text(), " ")

		if data[0] == "$" && data[1] == "cd" {
			// Because we check all edge cases we know any command will be a 'cd'
			if data[2] == ".." {
				// .. means we want to go up 1 directory so we replace the currentNode with the parent of the current node
				currentNode = currentNode.Parent
			} else {
				// If its not .. we simply change the current node for the child of current node with the correct name
				currentNode = currentNode.Children[data[2]]
			}

			// after this we can quit this itteration
			continue
		}

		if data[0] == "dir" {
			// If the file starts with dir we add a new node of Kind dir
			newNode := &utils.Node{Value: 0, Parent: currentNode, Name: data[1], Kind: "dir", Children: make(map[string]*utils.Node)}
			currentNode.Insert(newNode)
		} else {
			// If anyting else we add it as type file
			val, _ := strconv.Atoi(data[0])
			newNode := &utils.Node{Value: val, Parent: currentNode, Name: data[1], Kind: "file", Children: nil}
			currentNode.Insert(newNode)
		}
	}

	return tree
}
