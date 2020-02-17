package vangogh_core_render_post

import (
  "html/template"
  "strings"

  vapi "github.com/cripplet/vangogh/api"
  vcru "github.com/cripplet/vangogh/core/render_util"
  vct "github.com/cripplet/vangogh/core/type"
)

// Function RenderPostList will generate the paginated index of a list of
// vangogh_api_proto.Posts.
func RenderPostList(v vct.ViewPostListData, path_prefix string) ([]vapi.RoutingTableRow, error) {
  pages := []vapi.RoutingTableRow{}

  f, err := vcru.GetComponentFiles()
  if err != nil {
    return nil, err
  }

  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post_list.gohtml")
  t, err := template.New("").Funcs(
    vcru.GetVangoghCoreTemplateFuncMap(),
  ).ParseFiles(f...)
  if err != nil {
    return nil, err
  }

  err = t.ExecuteTemplate(&b, "page", v)
  if err != nil {
    return nil, err
  }

  pages = append(pages, vapi.RoutingTableRow{
      Path: path_prefix,
      Reader: strings.NewReader(b.String()),
  })
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
