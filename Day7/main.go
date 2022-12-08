package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jethridge13/AdventOfCode2022/util"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name     string
	size     int
	parent   *Dir
	children []*Dir
	files    []File
}

func addSizesToTree(root *Dir) int {
	size := 0
	for _, dir := range root.children {
		size += addSizesToTree(dir)
	}
	for _, file := range root.files {
		size += file.size
	}
	root.size = size
	return size
}

func getSizeBelowThreshold(root *Dir, threshold int) int {
	size := 0
	if root.size < threshold {
		size += root.size
	}
	for _, dir := range root.children {
		size += getSizeBelowThreshold(dir, threshold)
	}
	return size
}

func findSmallestDirToDelete(root *Dir, space int) int {
	smallestSize := 70000000
	if root.size < smallestSize && root.size >= space {
		smallestSize = root.size
	}
	for _, dir := range root.children {
		size := findSmallestDirToDelete(dir, space)
		if size < smallestSize && size >= space {
			smallestSize = size
		}
	}
	return smallestSize
}

func buildTree(path string) Dir {
	scanner := util.GetFileScanner(path)
	root := Dir{name: "/"}
	var cwd *Dir
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ cd /" {
			cwd = &root
			continue
		}
		parts := strings.Fields(line)
		if parts[0] == "$" {
			// Command
			if parts[1] == "cd" {
				if parts[2] == ".." {
					cwd = cwd.parent
				} else {
					for _, d := range cwd.children {
						if parts[2] == d.name {
							cwd = d
							break
						}
					}
				}
			}
			// Ignore ls commands, other conditionals will add files and dirs
		} else if parts[0] == "dir" {
			// Directory
			d := Dir{name: parts[1], parent: cwd}
			cwd.children = append(cwd.children, &d)
		} else {
			// File
			size, _ := strconv.Atoi(parts[0])
			cwd.files = append(cwd.files, File{size: size, name: parts[1]})
		}
	}
	return root
}

func part1(path string) int {
	root := buildTree(path)
	addSizesToTree(&root)
	return getSizeBelowThreshold(&root, 100000)
}

func part2(path string) int {
	root := buildTree(path)
	addSizesToTree(&root)
	maxDiskSpace := 70000000
	freeSpaceNeeded := 30000000
	freeSpace := maxDiskSpace - root.size
	spaceToFree := freeSpaceNeeded - freeSpace
	return findSmallestDirToDelete(&root, spaceToFree)
}

func main() {
	file := "input.txt"
	// Part 1: 1232307
	fmt.Println(part1(file))
	// Part 2: 7268994
	fmt.Println(part2(file))
}
