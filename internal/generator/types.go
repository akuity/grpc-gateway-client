package generator

import "strings"

const (
	rpcUnaryReturnType     = "*"
	rpcStreamingReturnType = "<-chan *"
)

func newStructAccessor(parentFields []string, field string) string {
	return strings.Join(append(parentFields, field), ".")
}

func unexport(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}
