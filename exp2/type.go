package main

// 自增计数器，用于产生缓存区内容
type AutoNumber struct {
	value int
}

func (a *AutoNumber) Get() int {
	a.value += 1
	return a.value
}

type Buffer interface {
	String() string
	Produce()
	Consume()
}

// 缓存实现
type BufferImpl struct {
	Content           []int
	Size              int
	BufEmpty, BufFull int
	ppos, cpos        int
	Empty, Full       []int
	autoIncre         AutoNumber
}

func NewBuffer(bufferSize int) Buffer {
	if bufferSize <= 0 {
		return nil
	}
	return &BufferImpl{
		Content:  make([]int, bufferSize),
		BufEmpty: bufferSize,
		Size:     bufferSize,
	}
}
