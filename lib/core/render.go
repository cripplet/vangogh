package vangogh_core_render

import (
    "github.com/cripplet/vangogh/api"
    "github.com/cripplet/vangogh/proto"
)

type CoreRenderer struct {
}

func (r CoreRenderer) Generate(vangogh_proto_base.Blog) ([]vangogh_api.RenderOutput, error) {
    return []vangogh_api.RenderOutput{}, nil
}

func (r CoreRenderer) Render([]vangogh_api.RenderOutput) error {
    return nil
}
