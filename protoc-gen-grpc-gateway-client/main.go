package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/omrikiei/grpc-gateway-client/internal/generator"
)

func main() {
	flag.Parse()
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(p *protogen.Plugin) error {
		p.SupportedFeatures = generator.SupportedFeatures

		for _, f := range p.Files {
			if f.Generate {
				if _, err := generator.Generate(p, f); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
