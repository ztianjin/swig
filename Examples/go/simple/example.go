/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 2.0.1
 * 
 * This file is not intended to be easily readable and contains a number of 
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG 
 * interface file instead. 
 * ----------------------------------------------------------------------------- */

package example


type _swig_fnptr *byte
type _swig_memberptr *byte


func _swig_allocatememory(int) *byte
func _swig_internal_allocate(len int) *byte {
	return _swig_allocatememory(len)
}

func _swig_allocatestring(*byte, int) string
func _swig_internal_makegostring(p *byte, l int) string {
	return _swig_allocatestring(p, l)
}

func _swig_internal_gopanic(p *byte, l int) {
	panic(_swig_allocatestring(p, l))
}

func Gcd(int, int) int
func _swig_wrap_Foo_set(float64)

func SetFoo(arg1 float64) {
	_swig_wrap_Foo_set(arg1)
}

func GetFoo() float64

type SwigcptrVoid uintptr
type Void interface {
	Swigcptr() uintptr;
}
func (p SwigcptrVoid) Swigcptr() uintptr {
	return uintptr(p)
}

