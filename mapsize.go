package mapsize

import (
	"reflect"
	"unsafe"
)

// Size returns how many bytes of memory a map occupies,
// which is an approximation, but has reference value
func Size[K comparable, V any](m map[K]V) int64 {
	var zeroK K
	var zeroValue V
	keySize := unsafe.Sizeof(zeroK)
	valueSize := unsafe.Sizeof(zeroValue)
	vo := reflect.ValueOf(m)
	hm := (*hmap)(unsafe.Pointer(vo.Pointer())) //#nosec G103
	bn := 1<<hm.B + uintptr(hm.noverflow)
	bz := unsafe.Sizeof(bmap{}) + (keySize+valueSize)*bucketCnt
	return int64(unsafe.Sizeof(hmap{}) + bz*bn)
}

// PtrValueSize returns how many bytes of memory a map occupies,
// which is an approximation, but has reference value
func PtrValueSize[K comparable, V any](m map[K]*V) int64 {
	var zeroK K
	var zeroValue V
	keySize := unsafe.Sizeof(zeroK)
	valueSize := unsafe.Sizeof(zeroValue)
	vo := reflect.ValueOf(m)
	hm := (*hmap)(unsafe.Pointer(vo.Pointer())) //#nosec G103
	bn := 1<<hm.B + uintptr(hm.noverflow)
	bz := unsafe.Sizeof(bmap{}) + (keySize+valueSize)*bucketCnt
	return int64(unsafe.Sizeof(hmap{})+bz*bn) + int64(len(m)*8)
}
