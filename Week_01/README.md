学习笔记

- 分析 Queue 和 Priority Queue 的源码

看了下go的Queue和 Priority Queue的源码，链接如下：

官方版：
https://golang.org/pkg/container/heap/

使用堆来实现


非官方版：
[go-datastructures/blob/master/queue/queue.go](https://github.com/Workiva/go-datastructures/blob/master/queue/queue.go)

实现了并发安全的queue，内部使用channel和waitGroup来实现