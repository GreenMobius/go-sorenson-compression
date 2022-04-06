package sorensen

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/pierrec/xxHash/xxHash64"
)

func TestCompressSmall(t *testing.T) {
	var inputData = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}

	seed := uint64(0xDEADBEEF)
	hash := xxHash64.New(seed)
	hash.Write(inputData)
	expected := CompressionFile{
		Hash: hash.Sum64(),
		Bits: 15,
		Size: uint64(len(inputData)),
	}

	bufferSize = 16
	compressed, err := Compress(bytes.NewReader(inputData), seed)
	if err != nil {
		t.Fatalf("Compression returned error: %v", err)
	}

	compressedData := CompressionFile{}
	if err := binary.Read(bytes.NewReader(compressed), binary.LittleEndian, &compressedData); err != nil {
		t.Fatalf("Reading compressed data returned error: %v", err)
	}

	if expected.Hash != compressedData.Hash {
		t.Fatalf("Expected hash %016X, Compress returned hash %016X", expected.Hash, compressedData.Hash)
	}

	if expected.Bits != compressedData.Bits {
		t.Fatalf("Expected %v bits, Compress returned %d bits", expected.Bits, compressedData.Bits)
	}

	if expected.Size != compressedData.Size {
		t.Fatalf("Expected size %v, Compress returned size %d", expected.Size, compressedData.Size)
	}
}

func TestCompressLarge(t *testing.T) {
	var inputData = []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}

	seed := uint64(0xD0D0CAFE)
	hash := xxHash64.New(seed)
	hash.Write(inputData)
	expected := CompressionFile{
		Hash: hash.Sum64(),
		Bits: 80,
		Size: uint64(len(inputData)),
	}

	bufferSize = 8
	compressed, err := Compress(bytes.NewReader(inputData), seed)
	if err != nil {
		t.Fatalf("Compression returned error: %v", err)
	}

	compressedData := CompressionFile{}
	if err := binary.Read(bytes.NewReader(compressed), binary.LittleEndian, &compressedData); err != nil {
		t.Fatalf("Reading compressed data returned error: %v", err)
	}

	if expected.Hash != compressedData.Hash {
		t.Fatalf("Expected hash %016X, Compress returned hash %016X", expected.Hash, compressedData.Hash)
	}

	if expected.Bits != compressedData.Bits {
		t.Fatalf("Expected %v bits, Compress returned %d bits", expected.Bits, compressedData.Bits)
	}

	if expected.Size != compressedData.Size {
		t.Fatalf("Expected size %v, Compress returned size %d", expected.Size, compressedData.Size)
	}
}
