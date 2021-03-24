package main

import (
	"container/heap"
	"fmt"
)

// 创建进程
func (m *Model) Create() {
	var memory, priority int
	fmt.Print("请输入进程内存占用（单位 MB ）：")
	fmt.Scanln(&memory)
	fmt.Print("请输入进程优先级(int)：")
	fmt.Scanln(&priority)
	m.New = append(m.New, ProcessGenerator(memory, priority))
	m.printWithComment("添加进程到“新建”中")
	m.admit()
}

// 调度
func (m *Model) dispatch() {
	// 如果当前无运行任务且就绪状态不为空，从就绪状态中取出一个
	if len(m.Running) == 0 && len(m.Ready) > 0 {
		m.Running = append(m.Running, heap.Pop(&m.Ready).(Process))
		m.printWithComment("将“就绪”中优先级最高的进程放入“运行”")
	}
}

// 时间片超时
func (m *Model) Timeout() {
	// 如果当前有运行任务
	if len(m.Running) == 1 {
		currentProcess := m.Running[0]
		m.Running = m.Running[:0]
		heap.Push(&m.Ready, currentProcess)
		m.printWithComment("将“运行”进程放回“就绪”")
		m.dispatch()
	} else {
		getError("当前无运行进程")
	}
}

func (m *Model) EventWait() {
	// 如果当前有运行任务
	if len(m.Running) == 1 {
		currentProcess := m.Running[0]
		m.Running = m.Running[:0]
		heap.Push(&m.Blocked, currentProcess)
		m.printWithComment("将“运行”进程放到“阻塞”")
		// 尝试调度
		m.dispatch()
	} else {
		getError("当前无运行进程")
	}
}

func (m *Model) EventOccur() {
	// 为了简单，默认将阻塞进程中的最高优先级任务放到就绪中
	if len(m.Blocked) > 0 {
		heap.Push(&m.Ready, heap.Pop(&m.Blocked))
		m.printWithComment("将“阻塞”中最高优先级的进程放入“就绪”中")
		// 尝试调度
		m.dispatch()
	} else {
		getError("阻塞进程列表为空")
	}
}

func (m *Model) Release() {
	if len(m.Running) == 1 {
		toRelease := m.Running[0]
		m.Exit = append(m.Exit, toRelease)
		m.Running = m.Running[:0]
		m.CurrentMemory -= toRelease.Memory
		m.printWithComment("将“运行”进程放到“退出”中")
		// 尝试调度
		m.dispatch()
	} else {
		getError("当前无运行进程")
	}
}

func (m *Model) admit() {
	item := m.New.Pop().(Process)
	// 如果放入模型中会导致内存溢出则放入就绪挂起
	if m.CurrentMemory+item.Memory > m.MaxMemory {
		heap.Push(&m.ReadySuspend, item)
		m.printWithComment("将“新建”中的新进程放入“就绪挂起”")
	} else {
		// 否则放入就绪
		heap.Push(&m.Ready, item)
		m.CurrentMemory += item.Memory
		m.printWithComment("将“新建”中的新进程放入“就绪”")
		m.dispatch()
	}
}
