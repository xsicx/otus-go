package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const Chunk int64 = 1024

func Copy(fromPath, toPath string, offset, limit int64) error {
	destFileSize, err := calcDestinationFileSize(fromPath, offset, limit)
	if err != nil {
		return err
	}

	bar := pb.Start64(destFileSize)

	sourceFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0o666)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	buf := make([]byte, Chunk)
	isCopied := false
	for {
		n, err := sourceFile.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 || isCopied {
			break
		}

		bufLimit := int64(n)
		if limit != 0 && bufLimit > limit {
			bufLimit = limit
			isCopied = true
		}

		if _, err := destinationFile.Write(buf[:bufLimit]); err != nil {
			return err
		}
		bar.Add64(bufLimit)

		offset += bufLimit
	}

	bar.Finish()

	return nil
}

func calcDestinationFileSize(sourceFilePath string, offset int64, limit int64) (int64, error) {
	fileInfo, err := os.Stat(sourceFilePath)
	if err != nil || fileInfo.Size() == 0 {
		return 0, ErrUnsupportedFile
	}

	if offset > fileInfo.Size() {
		return 0, ErrOffsetExceedsFileSize
	}

	fileSize := fileInfo.Size() - offset

	if limit != 0 && limit < fileSize {
		fileSize = limit
	}

	return fileSize, nil
}
