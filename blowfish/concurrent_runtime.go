package main

import (
	"io"
	"os"
	"sync"
)

type Block struct {
	Data  []byte
	Index uint
}

func loadBlock(file *os.File, blockChan chan<- Block, encrypt bool, errChan chan<- error) {
	defer close(blockChan)
	buffer := make([]byte, blockSize)
	var index uint = 0

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			errChan <- err
			return
		}

		if n == 0 {
			break
		}

		if n < blockSize && encrypt {
			buffer = applyPadding(buffer[:n], blockSize)
		} else {
			buffer = buffer[:n]
		}

		blockChan <- Block{Data: append([]byte{}, buffer...), Index: index}
		index++
		if err != nil && err == io.EOF {
			break
		}
	}
}

func processBlockWorker(block Block, resultChan chan<- Block, localP p, localS s, encrypt bool, wg *sync.WaitGroup) {
	defer wg.Done()
	processedData := processBlock(block.Data, localP, localS, encrypt)
	resultChan <- Block{Data: processedData, Index: block.Index}
}

func writeBlock(outFile *os.File, resultChan <-chan Block, errChan chan<- error) {
	buffer := make(map[uint][]byte)
	var expectedIndex uint = 0

	for block := range resultChan {
		buffer[block.Index] = block.Data

		for {
			data, exists := buffer[expectedIndex]
			if !exists {
				break
			}

			_, err := outFile.Write(data)
			if err != nil {
				errChan <- err
				return
			}

			delete(buffer, expectedIndex)
			expectedIndex++
		}
	}
}

func ConcurrentStream(file *os.File, outFile *os.File, key []byte, encrypt bool) error {
	localP, localS := initializeBlowfishKey(key)

	blockChan := make(chan Block, workerCount)
	resultChan := make(chan Block, workerCount)
	errChan := make(chan error, 1)

	var wg sync.WaitGroup

	go loadBlock(file, blockChan, encrypt, errChan)

	go func() {
		for block := range blockChan {
			wg.Add(1)
			go processBlockWorker(block, resultChan, localP, localS, encrypt, &wg)
		}
		wg.Wait()
		close(resultChan)
	}()

	writerDone := make(chan struct{})
	go func() {
		writeBlock(outFile, resultChan, errChan)
		close(writerDone)
	}()

	select {
	case <-writerDone:
	case err := <-errChan:
		return err
	}

	return nil
}
