package font

func LoadFromHexString(dataHex string) *Font {
	f := &Font{}

	err := f.Parse(dataHex)

	if err != nil {
		panic(err)
	}

	return f
}
