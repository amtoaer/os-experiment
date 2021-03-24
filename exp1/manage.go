package main

import (
	"container/heap"
	"fmt"
)

// manager接口实现
func (m *Model) Create() {
	var memory, priority int
	fmt.Print("请输入进程内存占用：")
	fmt.Scanln(&memory)
	fmt.Print("请输入进程优先级(int)：")
	fmt.Scanln(&priority)
	clear()
	m.New = append(m.New, ProcessGenerator(memory, priority))
	m.printWithComment("添加进程到New中")
	m.admit()
}

func (m *Model) Dispatch() {

}

func (m *Model) Timeout() {

}

func (m *Model) EventWait() {

}

func (m *Model) EventOccur() {

}

func (m *Model) OutsideMemoryEventOccur() {

}

func (m *Model) acitive() {

}

func (m *Model) admit() {
	clear()
	item := m.New.Pop().(Process)
	// 如果放入模型中会导致内存溢出则放入就绪挂起
	if m.CurrentMemory+item.Memory > m.MaxMemory {
		heap.Push(&m.ReadySuspend, item)
		m.printWithComment("将New中的新进程放入就绪挂起")
	} else {
		// 否则放入就绪
		heap.Push(&m.Ready, item)
		m.CurrentMemory += item.Memory
		m.printWithComment("将New中的新进程放入就绪")
	}
}
