package main

// #cgo CXXFLAGS: -I/usr/lib/ -std=c++11 -I${SRCDIR}/c/include -I${SRCDIR}/c/include
// #cgo LDFLAGS: -L/usr/lib/  -lstdc++ -L${SRCDIR}/c/lib -lflatbuffers
// #ifdef __cplusplus
//   extern "C" {
// #endif
// #include "fbcgo.hpp"
// #ifdef __cplusplus
//   }
// #endif
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Parser struct {
	c C.Parser
}

func NewParser() *Parser {
	p := &Parser{}
	p.create()
	return p
}

func (p *Parser) create() {
	C.Create(&p.c)
}

func (p *Parser) AddBuffer(b []byte) {
	cb := C.CString(string(b))
	defer C.free(unsafe.Pointer(cb))
	C.AddBuffer(&p.c, cb)
}

// Read is an io.Reader compatible for the last added buffer
func (p *Parser) Read(b []byte) (n int, err error) {
	return 0, fmt.Errorf("not implemented")
}

// Len returns the number of bytes used by the last added buffer
func (p *Parser) Len() (n int) {
	sz := C.GetSize(&p.c)
	return int(sz)
}

func (p *Parser) Finish() {
	C.Finish(&p.c)
}

func (p *Parser) GetBytes() []byte {
	cb := C.GetBuffer(&p.c)
	i := p.Len()
	if i == 0 {
		return nil
	}
	return C.GoBytes(unsafe.Pointer(cb), C.int(i))
}

func (p *Parser) Destroy() {
	C.Destroy(&p.c)
}

func main() {
	s := []byte(`
	table CreateRequest {
	  display_name:string;
	  audience:string;
	  scopes:string;
	}
	root_type CreateRequest;
	`)
	j := []byte(`{"display_name": "Bob", "audience": "builders", "scopes": "one two three"}`)

	parser := NewParser()
	parser.AddBuffer(s)
	parser.AddBuffer(j)

	// parser.Finish()
	b := parser.GetBytes()
	// fmt.Println(hex.EncodeToString(b))
	fmt.Printf("%c", b)
	parser.Destroy()

	// parser = NewParser()
	// parser.AddBuffer(s)
	// parser.AddBuffer(b)
	// parser.Finish()
	// bb := parser.GetBytes()
	// fmt.Println(bb)
	// parser.Destroy()
}
