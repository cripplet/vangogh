package vangogh_core_render_post

import (
  "io"
  "html/template"
  "strings"

  vcru "github.com/cripplet/vangogh/core/render_util"
  vct "github.com/cripplet/vangogh/core/type"
)

func RenderPostList(v vct.ViewPostListData, path_prefix string) (string, io.Reader, error) {
  f, err := vcru.GetComponentFiles()
  if err != nil {
    return "", nil, err
  }

  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post_list.gohtml")
  t, err := template.New("").Funcs(
    vcru.GetVangoghCoreTemplateFuncMap(),
  ).ParseFiles(f...)
  if err != nil {
    return "", nil, err
  }

  err = t.ExecuteTemplate(&b, "page", v)
  if err != nil {
    return "", nil, err
  }

  return path_prefix, strings.NewReader(b.String()), nil
}

func RenderPost(v vct.ViewPostData) (string, io.Reader, error) {
  f, err := vcru.GetComponentFiles()
  if err != nil {
    return "", nil, err
  }

  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post.gohtml")
  t, err := template.New("").Funcs(
    vcru.GetVangoghCoreTemplateFuncMap(),
  ).ParseFiles(f...)
  if err != nil {
    return "", nil, err
  }

  err = t.ExecuteTemplate(&b, "page", v)
  if err != nil {
    return "", nil, err
  }

  path, err := vcru.FormatPostPath(v.Content)
  if err != nil {
    return "", nil, err
  }
  return path, strings.NewReader(b.String()), nil
}
