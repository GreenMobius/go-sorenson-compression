package sorensen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/bits"

	"github.com/pierrec/xxHash/xxHash64"
)

var bufferSize int = 65536

const CompressedSize = 24

type CompressionFile struct {
	Hash uint64
	Bits uint64
	Size uint64
}

func Ones(data []byte) uint64 {
	var ones uint64 = 0
	for _, byt := range data {
		ones += uint64(bits.OnesCount8(byt))
	}
	return ones
}

func Compress(in io.Reader, seed uint64) ([]byte, error) {
	compressedData := CompressionFile{}
	hash := xxHash64.New(seed)
	dataBuf := make([]byte, bufferSize)

	for {
		count, err := in.Read(dataBuf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read error: %w", err)
		}

		hash.Write(dataBuf[:count])

		compressedData.Size += uint64(count)
		compressedData.Bits += Ones(dataBuf[:count])
	}

	compressedData.Hash = hash.Sum64()

	compressedBytes := bytes.Buffer{}
	if err := binary.Write(&compressedBytes, binary.LittleEndian, compressedData); err != nil {
		return nil, fmt.Errorf("write error: %w", err)
	}

	return compressedBytes.Bytes(), nil
}
