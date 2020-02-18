package vangogh_core_type

import (
  vpb "github.com/cripplet/vangogh/api/proto"
)

type ViewPostData struct {
  Site vpb.Site
  // Page-specific data to be used in generating the content template.
  Content vpb.Post
}

type PaginatePageInfo struct {
  TotalPages int
  CurrentPage int
  PathPrefix string
}
type ViewPostListDataContent struct {
  Posts []vpb.Post
  PageInfo PaginatePageInfo
}
type ViewPostListData struct {
  Site vpb.Site
  Content ViewPostListDataContent
}
