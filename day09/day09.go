package day09

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/ruegerj/aoc-2024/pkg/common"
	"github.com/ruegerj/aoc-2024/pkg/util"
)

type Day09 struct{}

func (d Day09) Part1(input string) *common.Solution {
	fs, _ := parseFs(input)
	cur := fs.Front()

	for cur != nil {
		next := cur.Next()

		if !cur.Value.(*fsBlock).isEmpty() {
			cur = next
			continue
		}

		curBlock := cur.Value.(*fsBlock)
		for curBlock.size > 0 {
			tail := findPrev(fs.Back(), func(e *list.Element) bool { return !e.Value.(*fsBlock).isEmpty() })
			tailBlock := tail.Value.(*fsBlock)

			if tailBlock.size > curBlock.size {
				tailBlock.size -= curBlock.size
				fs.InsertBefore(&fsBlock{id: tailBlock.id, size: curBlock.size}, cur)
				curBlock.size = 0
				continue
			}

			curBlock.size -= tailBlock.size
			removedTail := fs.Remove(tail)
			fs.InsertBefore(removedTail, cur)
		}

		fs.Remove(cur)
		cur = next
	}

	fsChecksum := calcChecksum(fs)

	return common.NewSolution(1, fsChecksum)
}

func (d Day09) Part2(input string) *common.Solution {
	fs, fileLookup := parseFs(input)
	fileIdsDesc := util.SortedKeys(fileLookup, true)

	for _, fileId := range fileIdsDesc {
		fileElem := fileLookup[fileId]
		blockToMove := fileElem.Value.(*fsBlock)
		cur := fs.Front()

		for cur != nil {
			next := cur.Next()
			block := cur.Value.(*fsBlock)

			if block.id == blockToMove.id {
				break
			}

			if !block.isEmpty() {
				cur = next
				continue
			}

			if block.size >= blockToMove.size {
				block.size -= blockToMove.size
				fs.InsertBefore(blockToMove, cur)
				fs.InsertBefore(&fsBlock{id: -1, size: blockToMove.size}, fileElem)
				fs.Remove(fileElem)
				if block.size == 0 {
					fs.Remove(cur)
				}
				break
			}

			cur = next
		}

	}

	fsChecksum := calcChecksum(fs)

	return common.NewSolution(2, fsChecksum)
}

type filesystem *list.List

type fsBlock struct {
	id   int
	size int
}

func (fsb *fsBlock) isEmpty() bool {
	return fsb.id < 0
}

func (fsb fsBlock) String() string {
	return fmt.Sprintf("block{id: %d, value: %d}", fsb.id, fsb.size)
}

func swapBlocks(fs *list.List, a, b *list.Element) {
	fs.InsertAfter(b.Value, a)
	fs.InsertAfter(a.Value, b)
	fs.Remove(a)
	fs.Remove(b)
}

func findPrev(elem *list.Element, predicate func(*list.Element) bool) *list.Element {
	cur := elem
	for cur != nil && !predicate(cur) {
		cur = cur.Prev()
	}
	return cur
}

func calcChecksum(fs *list.List) int {
	pos := 0
	fsChecksum := 0
	cur := fs.Front()
	for cur != nil {
		block := cur.Value.(*fsBlock)
		next := cur.Next()
		if block.isEmpty() {
			pos += block.size
			cur = next
			continue
		}

		for i := 0; i < block.size; i++ {
			fsChecksum += block.id * (pos + i)
		}

		pos += block.size
		cur = next
	}
	return fsChecksum
}

func printFs(fs *list.List) {
	cur := fs.Front()
	for cur != nil {
		block := cur.Value.(*fsBlock)
		char := "."
		if !block.isEmpty() {
			char = fmt.Sprint(block.id)
		}

		for i := 0; i < block.size; i++ {
			fmt.Print(char)
		}

		cur = cur.Next()
	}
	fmt.Println()
}

func parseFs(input string) (*list.List, map[int]*list.Element) {
	fs := list.New()
	isEmpty := false
	fileLookup := make(map[int]*list.Element)
	id := 0

	for _, digit := range strings.Split(input, "") {
		space := util.MustParseInt(digit)

		block := fsBlock{id: -1, size: space}
		elem := fs.PushBack(&block)
		if !isEmpty {
			block.id = id
			fileLookup[id] = elem
			id++
		}
		isEmpty = !isEmpty
	}

	return fs, fileLookup
}
