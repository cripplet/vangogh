package vangogh_core_render_post

import (
  "html/template"
  "math"
  "strings"

  vapi "github.com/cripplet/vangogh/api"
  vpb "github.com/cripplet/vangogh/api/proto"
  vcru "github.com/cripplet/vangogh/core/render_util"
  vct "github.com/cripplet/vangogh/core/type"
)

func getPageNumber(sliceFloor int, batchSize int) int {
  return sliceFloor / batchSize + 1
}

// Function RenderPostList will generate the paginated index of a list of
// vangogh_api_proto.Posts.
func RenderPostList(
    pb vpb.Site,
    ps []vpb.Post,
    pathPrefix string) ([]vapi.RoutingTableRow, error) {
  pages := []vapi.RoutingTableRow{}

  batchSize := 1  // TODO(minkezhang): Add this as a SiteMetadata property.

  f, err := vcru.GetComponentFiles()
  if err != nil {
    return nil, err
  }

  f = append(f, "lib/core/template/view/post_list.gohtml")
  t, err := template.New("").Funcs(
    vcru.GetVangoghCoreTemplateFuncMap(),
  ).ParseFiles(f...)
  if err != nil {
    return nil, err
  }

  maxPageNum := (len(ps) / batchSize)
  for p := 0; p < maxPageNum; p++ {
    pageNum := p + 1
    sliceFloor := p * batchSize
    sliceCeil := int(math.Min(float64(sliceFloor + batchSize), float64(len(ps))))

    b := strings.Builder{}
    err = t.ExecuteTemplate(&b, "page", vct.ViewPostListData{
      Site: pb,
      Content: vct.ViewPostListDataContent{
        Posts: ps[sliceFloor:sliceCeil],
        PageInfo: vct.PaginatePageInfo{
            TotalPages: maxPageNum,
            CurrentPage: pageNum,
            PathPrefix: pathPrefix,
        },
      },
    })
    if err != nil {
      return nil, err
    }

    pages = append(pages, vapi.RoutingTableRow{
        Path: vcru.FormatPaginateURL(pathPrefix, pageNum),
        Reader: strings.NewReader(b.String()),
    })
    // First page of an index shouldn't need to explicitly specify the page number.
    if pageNum == 1 {
      pages = append(pages, vapi.RoutingTableRow{
          Path: pathPrefix,
          Reader: strings.NewReader(b.String()),
      })
    }
  }

  return pages, nil
}

func RenderPost(v vct.ViewPostData) (vapi.RoutingTableRow, error) {
  f, err := vcru.GetComponentFiles()
  if err != nil {
    return vapi.RoutingTableRow{}, err
  }

  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post.gohtml")
  t, err := template.New("").Funcs(
    vcru.GetVangoghCoreTemplateFuncMap(),
  ).ParseFiles(f...)
  if err != nil {
    return vapi.RoutingTableRow{}, err
  }

  err = t.ExecuteTemplate(&b, "page", v)
  if err != nil {
    return vapi.RoutingTableRow{}, err
  }

  path, err := vcru.FormatPostPath(v.Content)
  if err != nil {
    return vapi.RoutingTableRow{}, err
  }

  return vapi.RoutingTableRow{
      Path: path,
      Reader: strings.NewReader(b.String()),
  }, nil
}
