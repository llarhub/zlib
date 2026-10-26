package zlib

import (
	"unsafe"

	"github.com/goplus/lib/c"
)

func Crc32ZBytes(crc c.Ulong, buf []byte) c.Ulong {
	return Crc32Z(crc, unsafe.SliceData(buf), uintptr(len(buf)))
}

func Crc32ZString(crc c.Ulong, buf string) c.Ulong {
	return Crc32Z(crc, unsafe.StringData(buf), uintptr(len(buf)))
}
