package cmd

import (
	"context"
	"io"

	"sync"

	"github.com/siriushq/midio/cmd/logger"
)

// Writes in parallel to writers
type parallelWriter struct {
	writers     []io.Writer
	writeQuorum int
	errs        []error
}

// Write writes data to writers in parallel.
func (p *parallelWriter) Write(ctx context.Context, blocks [][]byte) error {
	var wg sync.WaitGroup

	for i := range p.writers {
		if p.writers[i] == nil {
			p.errs[i] = errDiskNotFound
			continue
		}

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, p.errs[i] = p.writers[i].Write(blocks[i])
			if p.errs[i] != nil {
				p.writers[i] = nil
			}
		}(i)
	}
	wg.Wait()

	// If nilCount >= p.writeQuorum, we return nil. This is because HealFile() uses
	// CreateFile with p.writeQuorum=1 to accommodate healing of single disk.
	// i.e if we do no return here in such a case, reduceWriteQuorumErrs() would
	// return a quorum error to HealFile().
	nilCount := 0
	for _, err := range p.errs {
		if err == nil {
			nilCount++
		}
	}
	if nilCount >= p.writeQuorum {
		return nil
	}
	return reduceWriteQuorumErrs(ctx, p.errs, objectOpIgnoredErrs, p.writeQuorum)
}

// Encode reads from the reader, erasure-encodes the data and writes to the writers.
func (e *Erasure) Encode(ctx context.Context, src io.Reader, writers []io.Writer, buf []byte, quorum int) (total int64, err error) {
	writer := &parallelWriter{
		writers:     writers,
		writeQuorum: quorum,
		errs:        make([]error, len(writers)),
	}

	for {
		var blocks [][]byte
		n, err := io.ReadFull(src, buf)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			logger.LogIf(ctx, err)
			return 0, err
		}
		eof := err == io.EOF || err == io.ErrUnexpectedEOF
		if n == 0 && total != 0 {
			// Reached EOF, nothing more to be done.
			break
		}
		// We take care of the situation where if n == 0 and total == 0 by creating empty data and parity files.
		blocks, err = e.EncodeData(ctx, buf[:n])
		if err != nil {
			logger.LogIf(ctx, err)
			return 0, err
		}

		if err = writer.Write(ctx, blocks); err != nil {
			logger.LogIf(ctx, err)
			return 0, err
		}
		total += int64(n)
		if eof {
			break
		}
	}
	return total, nil
}
