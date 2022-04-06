package command

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/GreenMobius/sorensen/pkg/sorensen"
)

func runCompression(inputFilePath string, outputFilePath string) error {
	inFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("opening input file: %w", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer outFile.Close()

	data, err := sorensen.Compress(inFile, 0)
	if err != nil {
		return fmt.Errorf("compressing file: %w", err)
	}

	if err := binary.Write(outFile, binary.LittleEndian, data); err != nil {
		return fmt.Errorf("writing compressed file: %w", err)
	}

	return nil
}
