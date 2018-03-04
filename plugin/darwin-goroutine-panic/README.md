# Go plugin issue with goroutines on darwin

```
$ go version
go version go1.10 darwin/amd64

$ ./test.sh
[plugin] NumGoroutine: 2
runtime: bad pointer in frame runtime.mallocgc at 0xc420050bc0: 0x7
fatal error: invalid pointer found on stack

runtime stack:
runtime.throw(0x45270e8, 0x1e)
	/usr/local/go/src/runtime/panic.go:619 +0x81 fp=0x70000fbdf6f8 sp=0x70000fbdf6d8 pc=0x4467891
runtime.adjustpointers(0xc420050b98, 0x70000fbdf7f0, 0x70000fbdfbc0, 0x454a458, 0x459e080)
	/usr/local/go/src/runtime/stack.go:592 +0x23e fp=0x70000fbdf768 sp=0x70000fbdf6f8 pc=0x447b30e
runtime.adjustframe(0x70000fbdfad0, 0x70000fbdfbc0, 0x459e080)
	/usr/local/go/src/runtime/stack.go:663 +0x32c fp=0x70000fbdf820 sp=0x70000fbdf768 pc=0x447b65c
runtime.gentraceback(0xffffffffffffffff, 0xffffffffffffffff, 0x0, 0xc420076780, 0x0, 0x0, 0x7fffffff, 0x452a250, 0x70000fbdfbc0, 0x0, ...)
	/usr/local/go/src/runtime/traceback.go:355 +0x136c fp=0x70000fbdfb38 sp=0x70000fbdf820 pc=0x44843bc
runtime.copystack(0xc420076780, 0x1000, 0x70000fbdfd01)
	/usr/local/go/src/runtime/stack.go:891 +0x26e fp=0x70000fbdfcf0 sp=0x70000fbdfb38 pc=0x447c14e
runtime.newstack()
	/usr/local/go/src/runtime/stack.go:1063 +0x310 fp=0x70000fbdfe80 sp=0x70000fbdfcf0 pc=0x447c560
runtime: unexpected return pc for runtime.morestack called from 0x0
stack: frame={sp:0x70000fbdfe80, fp:0x70000fbdfe88} stack=[0x70000fb60290,0x70000fbdfe90)
000070000fbdfd80:  000000c420076180  000070000fbdfdd0
000070000fbdfd90:  00000000044534ac <runtime.(*mcentral).grow+236>  000000c41fff92ff
000070000fbdfda0:  000000c400000000  000000000416e3a0
000070000fbdfdb0:  000000000416e3a0  0000000000000100
000070000fbdfdc0:  000000c420076780  000000c42003eb50
000070000fbdfdd0:  000000000444f0c5 <runtime.mallocgc+2021>  000000c420076780
000070000fbdfde0:  0000000000000000  0000000000000000
000070000fbdfdf0:  0000000000000000  0000000000000000
000070000fbdfe00:  0000000000002000  000000000416e3a0
000070000fbdfe10:  000000c420076780  000070000fbdfe48
000070000fbdfe20:  0000000004452b2c <runtime.(*mcache).refill+156>  00000000041065f0
000070000fbdfe30:  000000000416e3a0  0000000000000007
000070000fbdfe40:  000000c420076180  000070000fbdfe68
000070000fbdfe50:  0000000004487ab2 <runtime.(*mcache).nextFree.func1+50>  0000000004168d90
000070000fbdfe60:  000070000fbdfe07  000000c42003ebc8
000070000fbdfe70:  000000c42003ebe0  000000000448a349 <runtime.morestack+137>
000070000fbdfe80: <0000000000000000 >0100000004200000
runtime.morestack()
	/usr/local/go/src/runtime/asm_amd64.s:480 +0x89 fp=0x70000fbdfe88 sp=0x70000fbdfe80 pc=0x448a349

goroutine 19 [copystack]:
runtime.(*mcache).nextFree(0x4168d90, 0x10e, 0x0, 0x0, 0xc42003ebc8)
	/usr/local/go/src/runtime/malloc.go:545 +0x254 fp=0xc420050b50 sp=0xc420050b48 pc=0x444e8d4
runtime.mallocgc(0x60, 0x40a1460, 0x101000000000001, 0x0)
	/usr/local/go/src/runtime/malloc.go:710 +0x7e5 fp=0xc420050bf0 sp=0xc420050b50 pc=0x444f0c5
runtime.newobject(0x40a1460, 0xc420072000)
	/usr/local/go/src/runtime/malloc.go:839 +0x38 fp=0xc420050c20 sp=0xc420050bf0 pc=0x444f478
runtime.acquireSudog(0xc4200ae05c)
	/usr/local/go/src/runtime/proc.go:330 +0x2b4 fp=0xc420050c90 sp=0xc420050c20 pc=0x4469604
runtime.semacquire1(0xc4200ae05c, 0x0, 0x1)
	/usr/local/go/src/runtime/sema.go:115 +0x5e fp=0xc420050d00 sp=0xc420050c90 pc=0x447617e
internal/poll.runtime_Semacquire(0xc4200ae05c)
	/usr/local/go/src/runtime/sema.go:61 +0x39 fp=0xc420050d28 sp=0xc420050d00 pc=0x4475f69
internal/poll.(*fdMutex).rwlock(0xc4200ae050, 0x4525800, 0x16)
	/usr/local/go/src/internal/poll/fd_mutex.go:152 +0xad fp=0xc420050d68 sp=0xc420050d28 pc=0x44b555d
internal/poll.(*FD).writeLock(0xc4200ae050, 0xc4200d8040, 0x2)
	/usr/local/go/src/internal/poll/fd_mutex.go:237 +0x36 fp=0xc420050d90 sp=0xc420050d68 pc=0x44b59b6
internal/poll.(*FD).Write(0xc4200ae050, 0xc4200da000, 0x19, 0x20, 0x0, 0x0, 0x0)
	/usr/local/go/src/internal/poll/fd_unix.go:243 +0x46 fp=0xc420050df8 sp=0xc420050d90 pc=0x44b76b6
os.(*File).write(0xc4200ac008, 0xc4200da000, 0x19, 0x20, 0xc4200d6000, 0x76, 0x0)
	/usr/local/go/src/os/file_unix.go:243 +0x4e fp=0xc420050e40 sp=0xc420050df8 pc=0x44bdbce
os.(*File).Write(0xc4200ac008, 0xc4200da000, 0x19, 0x20, 0x8, 0xc4200d6000, 0x0)
	/usr/local/go/src/os/file.go:144 +0x6f fp=0xc420050ec0 sp=0xc420050e40 pc=0x44bb51f
fmt.Fprintln(0x4536da0, 0xc4200ac008, 0xc42003efb0, 0x2, 0x2, 0x408b600, 0x4445900, 0xc4200d6000)
	/usr/local/go/src/fmt/print.go:255 +0x8b fp=0xc420050f28 sp=0xc420050ec0 pc=0x44e216b
fmt.Println(0xc42003efb0, 0x2, 0x2, 0xc4200d6000, 0x0, 0x0)
	/usr/local/go/src/fmt/print.go:264 +0x5a fp=0xc420050f78 sp=0xc420050f28 pc=0x44e223a
github.com/tiborvass/goissues/plugin/darwin-goroutine-panic/plugin.Main()
	/Users/tiborvass/go/src/github.com/tiborvass/goissues/plugin/darwin-goroutine-panic/plugin/plugin.go:10 +0xa2 fp=0xc420050fe0 sp=0xc420050f78 pc=0x44e8572
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420050fe8 sp=0xc420050fe0 pc=0x40505e1
created by main.main
	/Users/tiborvass/go/src/github.com/tiborvass/goissues/plugin/darwin-goroutine-panic/plugin-runner.go:21 +0xdd

goroutine 1 [sleep]:
runtime.gopark(0x40abc30, 0x4103160, 0x40a61b0, 0x5, 0x13, 0x2)
	/usr/local/go/src/runtime/proc.go:291 +0x11a fp=0xc420067e60 sp=0xc420067e40 pc=0x402b55a
runtime.goparkunlock(0x4103160, 0x40a61b0, 0x5, 0x13, 0x2)
	/usr/local/go/src/runtime/proc.go:297 +0x5e fp=0xc420067ea0 sp=0xc420067e60 pc=0x402b60e
time.Sleep(0x3b9aca00)
	/usr/local/go/src/runtime/time.go:102 +0x166 fp=0xc420067f00 sp=0xc420067ea0 pc=0x40426e6
main.main()
	/Users/tiborvass/go/src/github.com/tiborvass/goissues/plugin/darwin-goroutine-panic/plugin-runner.go:24 +0xf4 fp=0xc420067f88 sp=0xc420067f00 pc=0x407b494
runtime.main()
	/usr/local/go/src/runtime/proc.go:198 +0x212 fp=0xc420067fe0 sp=0xc420067f88 pc=0x402b112
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420067fe8 sp=0xc420067fe0 pc=0x40505e1

goroutine 2 [force gc (idle)]:
runtime.gopark(0x40abc30, 0x41004d0, 0x40a6ee6, 0xf, 0x40abb14, 0x1)
	/usr/local/go/src/runtime/proc.go:291 +0x11a fp=0xc420042768 sp=0xc420042748 pc=0x402b55a
runtime.goparkunlock(0x41004d0, 0x40a6ee6, 0xf, 0x14, 0x1)
	/usr/local/go/src/runtime/proc.go:297 +0x5e fp=0xc4200427a8 sp=0xc420042768 pc=0x402b60e
runtime.forcegchelper()
	/usr/local/go/src/runtime/proc.go:248 +0xcc fp=0xc4200427e0 sp=0xc4200427a8 pc=0x402b39c
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc4200427e8 sp=0xc4200427e0 pc=0x40505e1
created by runtime.init.4
	/usr/local/go/src/runtime/proc.go:237 +0x35

goroutine 3 [GC sweep wait]:
runtime.gopark(0x40abc30, 0x41005c0, 0x40a6b81, 0xd, 0x401d714, 0x1)
	/usr/local/go/src/runtime/proc.go:291 +0x11a fp=0xc420042f60 sp=0xc420042f40 pc=0x402b55a
runtime.goparkunlock(0x41005c0, 0x40a6b81, 0xd, 0x14, 0x1)
	/usr/local/go/src/runtime/proc.go:297 +0x5e fp=0xc420042fa0 sp=0xc420042f60 pc=0x402b60e
runtime.bgsweep(0xc420074000)
	/usr/local/go/src/runtime/mgcsweep.go:52 +0xa3 fp=0xc420042fd8 sp=0xc420042fa0 pc=0x401d793
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420042fe0 sp=0xc420042fd8 pc=0x40505e1
created by runtime.gcenable
	/usr/local/go/src/runtime/mgc.go:216 +0x58

goroutine 18 [finalizer wait]:
runtime.gopark(0x40abc30, 0x411c7e0, 0x40a6d74, 0xe, 0x14, 0x1)
	/usr/local/go/src/runtime/proc.go:291 +0x11a fp=0xc42003e718 sp=0xc42003e6f8 pc=0x402b55a
runtime.goparkunlock(0x411c7e0, 0x40a6d74, 0xe, 0x14, 0x1)
	/usr/local/go/src/runtime/proc.go:297 +0x5e fp=0xc42003e758 sp=0xc42003e718 pc=0x402b60e
runtime.runfinq()
	/usr/local/go/src/runtime/mfinal.go:175 +0xad fp=0xc42003e7e0 sp=0xc42003e758 pc=0x4014a9d
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc42003e7e8 sp=0xc42003e7e0 pc=0x40505e1
created by runtime.createfing
	/usr/local/go/src/runtime/mfinal.go:156 +0x62

goroutine 20 [syscall]:
runtime.notetsleepg(0x4103180, 0x3b9aa4e1, 0x0)
	/usr/local/go/src/runtime/lock_sema.go:280 +0x4b fp=0xc42003f760 sp=0xc42003f720 pc=0x400f05b
runtime.timerproc(0x4103160)
	/usr/local/go/src/runtime/time.go:261 +0x2e7 fp=0xc42003f7d8 sp=0xc42003f760 pc=0x4042f67
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc42003f7e0 sp=0xc42003f7d8 pc=0x40505e1
created by runtime.(*timersBucket).addtimerLocked
	/usr/local/go/src/runtime/time.go:160 +0x107

$ ./test.sh
[plugin] NumGoroutine: 2
[plugin] NumGoroutine: 2
```

The second run is to show that the panic is not 100% reproducible.

## Comparison with Linux

```
$ go version
go version go1.10 linux/amd64

$ ./test.sh
+ go build -o plugin/plugin.so -buildmode=plugin ./plugin
+ go build
+ ./darwin-goroutine-panic ./plugin/plugin.so
[plugin] NumGoroutine: 2
[plugin] NumGoroutine: 2

$ ./test.sh
+ go build -o plugin/plugin.so -buildmode=plugin ./plugin
+ go build
+ ./darwin-goroutine-panic ./plugin/plugin.so
[plugin] NumGoroutine: 2
[plugin] NumGoroutine: 2
```

The second run is to show that it's never panicking.
