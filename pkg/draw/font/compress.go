package font

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io/ioutil"
)

func decompressFontHex(fontHex string) ([]byte, error) {
	decoded, err := hex.DecodeString(fontHex)
	readBuffer := bytes.NewBuffer(decoded)
	reader, err := gzip.NewReader(readBuffer)

	if err != nil {
		return nil, err
	}

	decompressed, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
