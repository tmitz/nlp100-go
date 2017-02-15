package british

import "os"
import "compress/gzip"
import "bufio"
import "io"
import "encoding/json"
import "strings"

type Page struct {
	Title, Text string
}

func Parse(file string) (body string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer fz.Close()

	var p Page
	reader := bufio.NewReaderSize(fz, 4096)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		dec := json.NewDecoder(strings.NewReader(line))
		if err = dec.Decode(&p); err != nil {
			panic(err)
		}
		if p.Title == "イギリス" {
			body = p.Text
			break
		}
	}
	return
}
