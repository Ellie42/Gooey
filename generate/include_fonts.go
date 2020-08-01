package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type FontFile struct {
	Path      string
	HexString string
	Name      string
}

func main() {
	files, _ := ioutil.ReadDir("./fonts")

	fileData := make([]FontFile, 0)

	for _, file := range files {
		if !strings.Contains(file.Name(), ".bff") {
			continue
		}

		fontFile := FontFile{}
		data, err := ioutil.ReadFile(fmt.Sprintf("./fonts/%s", file.Name()))

		if err != nil {
			panic(err)
		}

		fontFile.Path = "./fonts/" + file.Name()
		fontFile.Name = strings.SplitN(file.Name(), ".", 2)[0]

		rawBuffer := make([]byte, 0, len(data))
		buffer := bytes.NewBuffer(rawBuffer)

		writer := gzip.NewWriter(buffer)

		_, err = writer.Write(data)

		if err != nil {
			panic(err)
		}

		err = writer.Flush()

		if err != nil {
			panic(err)
		}

		err = writer.Close()

		if err != nil {
			panic(err)
		}

		fontFile.HexString = hex.EncodeToString(buffer.Bytes())

		fileData = append(fileData, fontFile)
	}

	tmpString := `package fonts

const (
	{{ range $_, $value := . }} 
		{{ $value.Name }} = "{{ $value.HexString }}"
	{{ end }}
)
`

	tmp := template.New("fonts")

	var err error

	tmp, err = tmp.Parse(tmpString)

	if err != nil {
		panic(err)
	}

	fileHandle, err := os.OpenFile("./fonts/fonts_generated.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		panic(err)
	}

	err = tmp.Execute(fileHandle, fileData)

	if err != nil {
		panic(err)
	}
}
