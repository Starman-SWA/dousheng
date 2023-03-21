package util

import "io"

type ReadFunc func([]byte) (int, error)

func ReadAll(r ReadFunc) ([]byte, error) {
	// 创建一个 512 字节的 buf
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// 如果 buf 满了，则追加一个元素，使其重新分配内存
			b = append(b, 0)[:len(b)]
		}
		// 读取内容到 buf
		n, err := r(b[len(b):cap(b)])
		b = b[:len(b)+n]
		// 遇到结尾或者报错则返回
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}
