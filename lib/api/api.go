package vangogh_api

import (
    "io"
    "github.com/cripplet/vangogh/proto"
)

type RenderOutput struct {
    path string
    data io.Reader
}

type RenderInterface interface {
    Generate(vangogh_proto_base.Blog) ([]RenderOutput, error)
    Render([]RenderOutput) error
}
