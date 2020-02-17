// Package vangogh_api documents the core API that rendering packages need
// to implement for consistency. The included vangogh_api_util package
// expects inerfaces and structures defined here.
package vangogh_api

import (
    "io"
    "github.com/cripplet/vangogh/api/proto"
)

// Interface VangoghRenderer is tasked with transforming the contents
// encapsulated in the proto into a context-dependent consumable form.
// The return value is designed to be easily-digestable by the
// vangough_api_util package.
type VangoghRenderer interface {
  GeneratePages(vangogh_api_proto.Site) (map[string]io.Reader, error)
}
