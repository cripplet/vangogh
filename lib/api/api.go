package vangogh_api

import (
    "io"
    "github.com/cripplet/vangogh/api/proto"
)

type RenderOutput struct {
    path string
    data io.Reader
}

type RenderInterface interface {
    Generate(vangogh_api_proto_base.Blog) ([]RenderOutput, error)
    Render([]RenderOutput) error
}

type VangoghURLMap struct {
    m[string]string
}

type VangoghPageWriter interface {
    WritePage(vangogh_api_proto_base.BlogMetadata
}

type VangoghRenderer interface {
    
    GeneratePages(vangogh_api_proto_base.Blog) ([]VangoghPageWriter, error)
}
