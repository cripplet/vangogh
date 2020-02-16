package vangogh_core_render_post

import (
  "strings"

  vcru "github.com/cripplet/vangogh/core/render_util"
  vct "github.com/cripplet/vangogh/core/type"
)

func RenderPost(v vct.ViwePostData) (string, io.Reader, error) {
  f, err := vcru.GetComponentFiles()
  if err != nil {
    return "", nil, err
  }

  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post.gohtml")
  t, err := template.New("").Funcs(
    template.FuncMap{
      "deserialize": deserializeAnyProto,
      "formatTime": vcru.FormatTime,
    },
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
  fmt.Println(path)
  return path, strings.NewReader(b.String()), nil
} 

