package goid

import (
	"unsafe"
)

var offset uintptr

func init() {
	gg := (*g)(nil)
	offset = unsafe.Offsetof(gg.goid)
}

func getG() uintptr

func Gid() int64 {
	gPtr := getG()
	return *(*int64)(unsafe.Pointer(gPtr + offset))
}

type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic         uintptr
	_defer         uintptr
	m              uintptr
	sched          gobuf
	syscallsp      uintptr
	syscallpc      uintptr
	stktopsp       uintptr
	param          unsafe.Pointer
	atomicstatus   uint32
	stackLock      uint32
	goid           int64
	waitsince      int64
	waitreason     string
	schedlink      uintptr
	preempt        bool
	paniconfault   bool
	preemptscan    bool
	gcscandone     bool
	gcscanvalid    bool
	throwsplit     bool
	raceignore     int8
	sysblocktraced bool
	sysexitticks   int64
	traceseq       uint64
	tracelastp     uintptr
	lockedm        uintptr
	sig            uint32
	writebuf       []byte
	sigcode0       uintptr
	sigcode1       uintptr
	sigpc          uintptr
	gopc           uintptr
	startpc        uintptr
	racectx        uintptr
	waiting        uintptr
	cgoCtxt        []uintptr
	labels         unsafe.Pointer
	timer          uintptr
	selectDone     uint32
	gcAssistBytes  int64
}

type stack struct {
	lo uintptr
	hi uintptr
}

type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    uintptr
	ctxt unsafe.Pointer
	ret  uint64
	lr   uintptr
	bp   uintptr
}
