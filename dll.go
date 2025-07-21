//go:build windows

package memdll

//#include<stdlib.h>
// void * MemoryLoadLibrary(const void *, size_t);
// void * MemoryGetProcAddress(void *, char *);
// void MemoryFreeLibrary(void *);
import "C"
import (
	"errors"
	"syscall"
	"unsafe"
)

type DLL struct {
	Name   string
	Handle syscall.Handle
}

func (d *DLL) FindProc(name string) (proc *Proc, err error) {
	var cname = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if addr := C.MemoryGetProcAddress(unsafe.Pointer(d.Handle), cname); addr != nil {
		return &Proc{
			Dll:  d,
			Name: name,
			addr: uintptr(addr),
		}, nil
	} else {
		var e = errors.New("no such function")
		return nil, &syscall.DLLError{
			Err:     e,
			ObjName: name,
			Msg:     "Failed to find " + name + " procedure in " + d.Name + ": " + e.Error(),
		}
	}
}

func (d *DLL) MustFindProc(name string) *Proc {
	if p, e := d.FindProc(name); e != nil {
		panic(e)
	} else {
		return p
	}
}

func (d *DLL) Release() {
	C.MemoryFreeLibrary(unsafe.Pointer(d.Handle))
}

type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

func (p *Proc) Addr() uintptr {
	return p.addr
}

func (p *Proc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return syscall.SyscallN(p.Addr(), a...)
}

func NewDLL(data []byte, dllname string) (*DLL, error) {
	if h := C.MemoryLoadLibrary(
		unsafe.Pointer(&data[0]), C.size_t(len(data)),
	); h != nil {
		return &DLL{
			Name:   dllname,
			Handle: syscall.Handle(h),
		}, nil
	} else {
		var e = errors.New("dll data error")
		return nil, &syscall.DLLError{
			Err:     e,
			ObjName: dllname,
			Msg:     "Failed to load " + dllname + ": " + e.Error(),
		}
	}
}
