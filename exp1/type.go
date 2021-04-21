package main

import (
	"container/heap"
	"fmt"
)

type Process struct {
	// 进程ID（自动分配）
	PID int
	// 占用内存
	Memory int
	// 优先级（数字越小优先级越高）
	Priority int
}

// 函数闭包
func processGenerator() func(int, int) Process {
	id := 0
	return func(Memory, Priority int) Process {
		id += 1
		return Process{
			PID:      id,
			Memory:   Memory,
			Priority: Priority,
		}
	}
}

// 自动编号的进程生成器
var ProcessGenerator func(int, int) Process = processGenerator()

// 覆写进程结构体的输出方法
func (p Process) String() string {
	return fmt.Sprintf("(ID:%d %dMB %d)", p.PID, p.Memory, p.Priority)
}

// 优先队列
type PQueue []Process

// 检验优先队列是否实现小顶堆接口
var _ heap.Interface = (*PQueue)(nil)

// 实现优先队列所需的接口
func (p PQueue) Len() int {
	return len(p)
}

func (p PQueue) Less(i, j int) bool {
	return p[i].Priority < p[j].Priority
}

func (p PQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQueue) Push(item interface{}) {
	*p = append(*p, item.(Process))
}

func (p *PQueue) Pop() interface{} {
	item := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return item
}

// 调度模型
type Model struct {
	// 五状态优先队列
	New, Ready, Running, Blocked, Exit PQueue
	// 最大内存
	MaxMemory int
	// 当前内存
	CurrentMemory int
}

// 覆写状态模型的输出方法
func (m Model) String() string {
	return fmt.Sprintf(`注：操作后，可敲击若干次回车查看详细过程。
最大内存：%vMB
当前内存：%vMB
新建：%v
就绪：%v
阻塞：%v
运行：%v
退出：%v`, m.MaxMemory, m.CurrentMemory, m.New, m.Ready, m.Blocked, m.Running, m.Exit)
}

// 为状态模型绑定带注释的输出方法
func (m Model) printWithComment(comment string) {
	clear()
	fmt.Println(m)
	fmt.Println(comment)
	fmt.Scanln()
}

// 管理接口
type Manager interface {
	// 主动调用方法
	Timeout()
	EventWait()
	EventOccur()
	Create()
	Release()
	// 被动调用方法
	dispatch()
	admit()
}

func NewManager(MaxMemory int) Manager {
	return &Model{
		MaxMemory: MaxMemory,
	}
}
