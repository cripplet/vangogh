package vangogh_core_type

import (
  vpb "github.com/cripplet/vangogh/api/proto"
)

type ViewPostData struct {
  Site vpb.Site
  // Page-specific data to be used in generating the content template.
  Content vpb.Post
}

type ViewPostListDataContent struct {
  Posts []vpb.Post
  NextPageLink string
  PrevPageLink string
}
type ViewPostListData struct {
  Site vpb.Site
  Content ViewPostListDataContent
}
