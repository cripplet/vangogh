package vangogh_core_type

import (
  vpb "github.com/cripplet/vangogh/api/proto"
)

type ViewPostData struct {
  Blog vpb.Blog
  // Page-specific data to be used in generating the content template.
  Content vpb.Post
}
