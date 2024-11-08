package Export

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type DataWriterInterface interface {
	WriteData(writer io.Writer) error
}

func WriteToFile(fileName string, dataWriter DataWriterInterface) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file: %v\n", err)
		}
	}()

	writer := bufio.NewWriter(file)
	if err := dataWriter.WriteData(writer); err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w\n", err)
	}
	return nil
}
