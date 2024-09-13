package generator

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/visibility"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	loopKeyAccessor   = "k"
	loopValueAccessor = "v"

	mapKeyVarName = "key"
)

func getClientInterfaceName(svc *protogen.Service) string {
	return fmt.Sprintf("%sGatewayClient", svc.GoName)
}

func getClientStructName(svc *protogen.Service) string {
	return unexport(getClientInterfaceName(svc))
}

type HTTPRule struct {
	Method  string
	Pattern string
	Body    string
}

func getHTTPRule(m *protogen.Method) (HTTPRule, bool) {
	options, ok := m.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return HTTPRule{}, false
	}

	rule, ok := proto.GetExtension(options, annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return HTTPRule{}, false
	}

	switch rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		return HTTPRule{
			Method:  http.MethodGet,
			Pattern: rule.GetGet(),
		}, true
	case *annotations.HttpRule_Post:
		return HTTPRule{
			Method:  http.MethodPost,
			Pattern: rule.GetPost(),
			Body:    rule.GetBody(),
		}, true
	case *annotations.HttpRule_Put:
		return HTTPRule{
			Method:  http.MethodPut,
			Pattern: rule.GetPut(),
			Body:    rule.GetBody(),
		}, true
	case *annotations.HttpRule_Patch:
		return HTTPRule{
			Method:  http.MethodPatch,
			Pattern: rule.GetPatch(),
			Body:    rule.GetBody(),
		}, true
	case *annotations.HttpRule_Delete:
		return HTTPRule{
			Method:  http.MethodDelete,
			Pattern: rule.GetDelete(),
			Body:    rule.GetBody(),
		}, true
	default:
		return HTTPRule{}, false
	}
}

func hasGatewayCompatibleMethods(file *protogen.File) bool {
	for _, srv := range file.Services {
		for _, method := range srv.Methods {
			if !isGatewayCompatibleMethod(method) {
				continue
			}
			return true
		}
	}
	return false
}

func isGatewayCompatibleMethod(m *protogen.Method) bool {
	_, ok := getHTTPRule(m)
	return ok && !m.Desc.IsStreamingClient()
}

func generateQueryParam(
	g *protogen.GeneratedFile,
	field *protogen.Field,
	structFields []string,
	isMapKeyDefined bool,
	queryKeyFields ...string,
) {
	isOptional := field.Desc.HasOptionalKeyword()
	isMap := field.Desc.IsMap()
	isRepeated := field.Desc.Cardinality() == protoreflect.Repeated

	queryKeyName := newStructAccessor(queryKeyFields, field.Desc.JSONName())
	queryValueAccessor := newStructAccessor(structFields, field.GoName)

	if field.Oneof != nil {
		queryValueAccessor = newStructAccessor(structFields, field.Oneof.GoName)
	}

	var restrictions []string
	if proto.HasExtension(field.Desc.Options(), visibility.E_FieldVisibility) {
		ext := proto.GetExtension(field.Desc.Options(), visibility.E_FieldVisibility)
		if opts, ok := ext.(*visibility.VisibilityRule); ok {
			restrictions = strings.Split(strings.TrimSpace(opts.Restriction), ",")
		}
	}

	if len(restrictions) > 0 {
		return
	}

	// If current field is inside the repeated message, ignore intermediate fields
	// since the loopValueAccessor directs the current field itself.
	if len(structFields) > 1 && structFields[0] == loopValueAccessor {
		queryValueAccessor = newStructAccessor(structFields[:1], field.GoName)
	}

	if isMap {
		g.P("for ", loopKeyAccessor, ", ", loopValueAccessor, " := range ", queryValueAccessor, " {")
		g.P(mapKeyVarName, " := ", pkgFmt.Ident("Sprintf"), `("`, queryKeyName, `[%v]", `, loopKeyAccessor, ")")
		queryKeyName = mapKeyVarName
		queryValueAccessor = loopValueAccessor
		structFields = []string{loopValueAccessor}
		defer g.P("}")
	} else {
		queryKeyName = fmt.Sprintf("%q", queryKeyName)
		if isRepeated {
			g.P("for _, ", loopValueAccessor, " := range ", queryValueAccessor, " {")
			queryValueAccessor = loopValueAccessor
			structFields = []string{loopValueAccessor}
			defer g.P("}")
		} else if isOptional {
			g.P("if ", queryValueAccessor, " != nil {")
			defer g.P("}")
		}
	}

	switch {
	case !isMap && field.Desc.Message() != nil:
		for _, f := range field.Message.Fields {
			generateQueryParam(g, f, append(structFields, field.GoName), isMapKeyDefined, append(queryKeyFields, field.Desc.JSONName())...)
		}
		return
	case field.Desc.Enum() != nil:
		g.P(`q.Add(`, queryKeyName, `, `, queryValueAccessor, ".String())")
	case isOptional:
		g.P(`q.Add(`, queryKeyName, `, `, pkgFmt.Ident("Sprintf"), `("%v", *`, queryValueAccessor, "))")
	default:
		g.P(`q.Add(`, queryKeyName, `, `, pkgFmt.Ident("Sprintf"), `("%v", `, queryValueAccessor, "))")
	}
}

func generateParamValues(g *protogen.GeneratedFile, m *protogen.Method) {
	rule, ok := getHTTPRule(m)
	if !ok {
		return
	}

	g.P(`gwReq := c.gwc.NewRequest("`, rule.Method, `", "`, rule.Pattern, `")`)
	fieldsByName := make(map[string]*protogen.Field)

	pathFields := make(map[string]bool)
	for _, field := range m.Input.Fields {
		fieldsByName[field.Desc.TextName()] = field
		fieldName := field.Desc.TextName()
		if strings.Contains(rule.Pattern, fmt.Sprintf("{%s}", fieldName)) {
			pathFields[fieldName] = true
			valueAccessor := newStructAccessor([]string{"req"}, field.GoName)
			if field.Oneof != nil {
				valueAccessor = newStructAccessor([]string{"req"}, field.Oneof.GoName)
			}
			if field.Desc.Enum() != nil {
				g.P(`gwReq.SetPathParam("`, fieldName, `", `, valueAccessor, ".String())")
			} else {
				g.P(`gwReq.SetPathParam("`, fieldName, `", `, pkgFmt.Ident("Sprintf"), `("%v", `, valueAccessor, "))")
			}
		}
	}

	switch rule.Method {
	case http.MethodGet:
		isQueryDefined := false
		for _, field := range m.Input.Fields {
			if _, ok := pathFields[field.Desc.TextName()]; !ok {
				if !isQueryDefined {
					g.P("q := ", pkgNetURL.Ident("Values"), "{}")
					isQueryDefined = true
				}
				generateQueryParam(g, field, []string{"req"}, false)
			}
		}
		if isQueryDefined {
			g.P("gwReq.SetQueryParamsFromValues(q)")
		}
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete:
		field := "req"
		if rule.Body != "*" {
			if bodyField, ok := fieldsByName[rule.Body]; ok {
				field = newStructAccessor([]string{field}, bodyField.GoName)
			}
		}
		g.P("gwReq.SetBody(", field, ")")
	}
}
