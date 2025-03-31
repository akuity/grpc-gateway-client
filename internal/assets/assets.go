package assets

import _ "embed"

var (
	//go:embed large_file.txt
	LargeFile string
)
