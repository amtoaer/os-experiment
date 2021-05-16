package main

import "fmt"

func (b *BufferImpl) checkEmpty() {
	if b.BufEmpty <= 0 && len(b.Full) > 0 {
		val := b.Full[0]
		b.Full = b.Full[1:]
		b.produce(val)
	}
}

func (b *BufferImpl) checkFull() {
	if b.BufFull <= 0 && len(b.Empty) > 0 {
		b.Empty = b.Empty[1:]
		b.consume()
	}
}

func (b *BufferImpl) produce(produceVal int) {
	b.Content[b.ppos] = produceVal
	b.ppos = (b.ppos + 1) % b.Size
	b.BufFull++
}

func (b *BufferImpl) Produce() {
	produceVal := b.autoIncre.Get()
	b.BufEmpty--
	if b.BufEmpty < 0 {
		b.Full = append(b.Full, produceVal)
	} else {
		b.produce(produceVal)
		b.checkFull()
	}
}
func (b *BufferImpl) consume() {
	b.Content[b.cpos] = 0
	b.cpos = (b.cpos + 1) % b.Size
	b.BufEmpty++
}

func (b *BufferImpl) Consume() {
	b.BufFull--
	if b.BufFull < 0 {
		b.Empty = append(b.Empty, 0)
	} else {
		b.consume()
		b.checkEmpty()
	}
}

func (b *BufferImpl) String() string {
	return fmt.Sprintf(`
BufferEmpty : %d		BufferFull : %d
Buffer : %v
Empty : %v
Full : %v
	`, b.BufEmpty, b.BufFull, b.Content, b.Empty, b.Full)
}
