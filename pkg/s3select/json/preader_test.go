package json

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/siriushq/midio/pkg/s3select/sql"
)

func TestNewPReader(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Run(file.Name(), func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", file.Name()))
			if err != nil {
				t.Fatal(err)
			}
			r := NewPReader(f, &ReaderArgs{})
			var record sql.Record
			for {
				record, err = r.Read(record)
				if err != nil {
					break
				}
			}
			r.Close()
			if err != io.EOF {
				t.Fatalf("Reading failed with %s, %s", err, file.Name())
			}
		})

		t.Run(file.Name()+"-close", func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", file.Name()))
			if err != nil {
				t.Fatal(err)
			}
			r := NewPReader(f, &ReaderArgs{})
			r.Close()
			var record sql.Record
			for {
				record, err = r.Read(record)
				if err != nil {
					break
				}
			}
			if err != io.EOF {
				t.Fatalf("Reading failed with %s, %s", err, file.Name())
			}
		})
	}
}

func BenchmarkPReader(b *testing.B) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		b.Fatal(err)
	}
	for _, file := range files {
		b.Run(file.Name(), func(b *testing.B) {
			f, err := ioutil.ReadFile(filepath.Join("testdata", file.Name()))
			if err != nil {
				b.Fatal(err)
			}
			b.SetBytes(int64(len(f)))
			b.ReportAllocs()
			b.ResetTimer()
			var record sql.Record
			for i := 0; i < b.N; i++ {
				r := NewPReader(ioutil.NopCloser(bytes.NewBuffer(f)), &ReaderArgs{})
				for {
					record, err = r.Read(record)
					if err != nil {
						break
					}
				}
				r.Close()
				if err != io.EOF {
					b.Fatalf("Reading failed with %s, %s", err, file.Name())
				}
			}
		})
	}
}
