package cmd

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func testBitrotReaderWriterAlgo(t *testing.T, bitrotAlgo BitrotAlgorithm) {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	volume := "testvol"
	filePath := "testfile"

	disk, err := newLocalXLStorage(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	disk.MakeVol(context.Background(), volume)

	writer := newBitrotWriter(disk, volume, filePath, 35, bitrotAlgo, 10, false)

	_, err = writer.Write([]byte("aaaaaaaaaa"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = writer.Write([]byte("aaaaaaaaaa"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = writer.Write([]byte("aaaaaaaaaa"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = writer.Write([]byte("aaaaa"))
	if err != nil {
		t.Fatal(err)
	}
	writer.(io.Closer).Close()

	reader := newBitrotReader(disk, nil, volume, filePath, 35, bitrotAlgo, bitrotWriterSum(writer), 10)
	b := make([]byte, 10)
	if _, err = reader.ReadAt(b, 0); err != nil {
		t.Fatal(err)
	}
	if _, err = reader.ReadAt(b, 10); err != nil {
		t.Fatal(err)
	}
	if _, err = reader.ReadAt(b, 20); err != nil {
		t.Fatal(err)
	}
	if _, err = reader.ReadAt(b[:5], 30); err != nil {
		t.Fatal(err)
	}
}

func TestAllBitrotAlgorithms(t *testing.T) {
	for bitrotAlgo := range bitrotAlgorithms {
		testBitrotReaderWriterAlgo(t, bitrotAlgo)
	}
}
