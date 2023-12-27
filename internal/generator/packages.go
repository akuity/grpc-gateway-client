package generator

import "google.golang.org/protobuf/compiler/protogen"

// Standard library packages
var (
	pkgContext = protogen.GoImportPath("context")
	pkgFmt     = protogen.GoImportPath("fmt")
	pkgNetURL  = protogen.GoImportPath("net/url")
)

// Akuity packages
var (
	pkgGatewayClient = protogen.GoImportPath("github.com/omrikiei/grpc-gateway-client/pkg/grpc/gateway")
)

// protobuf packages
var (
	pkgAnypb           = protogen.GoImportPath("google.golang.org/protobuf/types/known/anypb")
	pkgApipb           = protogen.GoImportPath("google.golang.org/protobuf/types/known/apipb")
	pkgDurationpb      = protogen.GoImportPath("google.golang.org/protobuf/types/known/durationpb")
	pkgEmptypb         = protogen.GoImportPath("google.golang.org/protobuf/types/known/emptypb")
	pkgFieldmaskpb     = protogen.GoImportPath("google.golang.org/protobuf/types/known/fieldmaskpb")
	pkgSourcecontextpb = protogen.GoImportPath("google.golang.org/protobuf/types/known/sourcecontextpb")
	pkgStructpb        = protogen.GoImportPath("google.golang.org/protobuf/types/known/structpb")
	pkgTimestamppb     = protogen.GoImportPath("google.golang.org/protobuf/types/known/timestamppb")
	pkgTypepb          = protogen.GoImportPath("google.golang.org/protobuf/types/known/typepb")
	pkgWrapperspb      = protogen.GoImportPath("google.golang.org/protobuf/types/known/wrapperspb")
)

func getMessageIdentifier(msg *protogen.Message) protogen.GoIdent {
	switch msg.Location.SourceFile {
	case "google/protobuf/any.proto":
		return pkgAnypb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/api.proto":
		return pkgApipb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/duration.proto":
		return pkgDurationpb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/empty.proto":
		return pkgEmptypb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/field_mask.proto":
		return pkgFieldmaskpb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/source_context.proto":
		return pkgSourcecontextpb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/struct.proto":
		return pkgStructpb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/timestamp.proto":
		return pkgTimestamppb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/type.proto":
		return pkgTypepb.Ident(msg.GoIdent.GoName)
	case "google/protobuf/wrappers.proto":
		return pkgWrapperspb.Ident(msg.GoIdent.GoName)
	default:
		return msg.GoIdent
	}
}
