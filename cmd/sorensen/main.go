package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var usage string = `Sorensen compression is a constant-size compression algorithm that solves all your compression needs.

Usage:

	sorensen <command> [flags] <file>

The commands are:

        compress  compress a file
        extract   decompress a compressed file
		info      view information on a compressed file

Use "sorensen <command> -h" for more information about a command.`

func main() {
	compressCmd := flag.NewFlagSet("compress", flag.ExitOnError)
	compressOutFile := compressCmd.String("o", "", "output file")

	decompressCmd := flag.NewFlagSet("decompress", flag.ExitOnError)
	decompressOutFile := decompressCmd.String("o", "", "output file")

	flag.Parse()

	args := os.Args
	if len(args) < 3 {
		log.Println(usage)
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]

	switch command {
	case "compress":
		if *compressOutFile == "" {
			*compressOutFile = strings.Trim(inputFile, filepath.Ext(inputFile)) + ".sor"
		}
		log.Println("Sorensen compression!")
		log.Printf("  outfile: %s", *compressOutFile)
	case "decompress":
		if *decompressOutFile == "" {
			*decompressOutFile = strings.Trim(inputFile, filepath.Ext(inputFile)) + ".sor"
		}
		log.Println("Sorensen decompression!")
		log.Printf("  outfile: %s", *decompressOutFile)
	case "info":
		log.Println("Sorensen info!")
	default:
		log.Println(usage)
	}
}
