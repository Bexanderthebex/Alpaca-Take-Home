package main

import "os"

type LotteryBetsReader struct {
	bufferOffset int64
	file         *os.File
}

func NewBetsReader(initialBufferOffset int64, file *os.File) *LotteryBetsReader {
	return &LotteryBetsReader{
		bufferOffset: initialBufferOffset,
		file:         file,
	}
}

func (l *LotteryBetsReader) Read(buffer *[]byte) (int, error) {
	length, fileReadError := l.file.ReadAt(*buffer, l.bufferOffset)

	for i := len(*buffer) - 1; i > 0; i -= 1 {
		if string((*buffer)[i]) == "\n" {
			*buffer = (*buffer)[:i]
			break
		}
	}
	l.updateBufferOffset(int64(len(*buffer) + 1))

	return length, fileReadError
}

func (l *LotteryBetsReader) updateBufferOffset(additionalOffset int64) int64 {
	l.bufferOffset += additionalOffset
	return l.bufferOffset
}
