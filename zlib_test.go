//go:build llgo

package zlib_test

import (
	"bytes"
	"testing"
	"unsafe"

	"github.com/goplus/lib/c"
	"github.com/llarhub/zlib"
)

const sampleText = "zlib is a software library used for data compression. It was created by Jean-loup Gailly and Mark Adler and first released in 1995. zlib is designed to be a free, legally unencumbered—that is, not covered by any patents—alternative to the proprietary DEFLATE compression algorithm, which is often used in software applications for data compression."

func bytePtr(data []byte) *zlib.Bytef {
	if len(data) == 0 {
		return nil
	}
	return (*zlib.Bytef)(unsafe.Pointer(unsafe.SliceData(data)))
}

func TestCRC32(t *testing.T) {
	data := []byte("Hello world")
	got := zlib.Crc32Z(0, bytePtr(data), zlib.SizeT(len(data)))
	if got != zlib.ULong(0x8bd69e52) {
		t.Fatalf("Crc32Z = %#x, want %#x", got, zlib.ULong(0x8bd69e52))
	}
}

func TestCompression(t *testing.T) {
	source := []byte(sampleText)
	sourceLen := zlib.ULong(len(source))
	compressedLen := zlib.ULongf(zlib.CompressBound(sourceLen))
	compressed := make([]byte, int(compressedLen))

	result := zlib.Compress(bytePtr(compressed), &compressedLen, bytePtr(source), sourceLen)
	if result != zlib.OK {
		t.Fatalf("Compress = %d, want %d", result, zlib.OK)
	}

	decompressedLen := zlib.ULongf(sourceLen)
	decompressed := make([]byte, int(decompressedLen))
	result = zlib.Uncompress(bytePtr(decompressed), &decompressedLen, bytePtr(compressed), zlib.ULong(compressedLen))
	if result != zlib.OK {
		t.Fatalf("Uncompress = %d, want %d", result, zlib.OK)
	}
	if !bytes.Equal(decompressed[:int(decompressedLen)], source) {
		t.Fatal("decompressed data does not match the source")
	}
}

func TestCompressionLevels(t *testing.T) {
	source := []byte(sampleText)
	sourceLen := zlib.ULong(len(source))

	for level := 0; level <= 9; level++ {
		t.Run(string(rune('0'+level)), func(t *testing.T) {
			compressedLen := zlib.ULongf(zlib.CompressBound(sourceLen))
			compressed := make([]byte, int(compressedLen))
			result := zlib.Compress2(bytePtr(compressed), &compressedLen, bytePtr(source), sourceLen, c.Int(level))
			if result != zlib.OK {
				t.Fatalf("Compress2(level=%d) = %d, want %d", level, result, zlib.OK)
			}

			decompressedLen := zlib.ULongf(sourceLen)
			decompressed := make([]byte, int(decompressedLen))
			result = zlib.Uncompress(bytePtr(decompressed), &decompressedLen, bytePtr(compressed), zlib.ULong(compressedLen))
			if result != zlib.OK {
				t.Fatalf("Uncompress(level=%d) = %d, want %d", level, result, zlib.OK)
			}
			if !bytes.Equal(decompressed[:int(decompressedLen)], source) {
				t.Fatalf("decompressed data at level %d does not match the source", level)
			}
		})
	}
}
