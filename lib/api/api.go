package vangogh_api

import (
    "io"
    "github.com/cripplet/vangogh/api/proto"
)

type VangoghRenderer interface {
  GeneratePages(vangogh_api_proto.Blog) (map[string]io.Reader, error)
}
