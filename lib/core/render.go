package vangogh_core_render

import (
  "html/template"
  "io"
  "path/filepath"
  "strings"
  vpb "github.com/cripplet/vangogh/api/proto"
)

func VangoghGenerate(pb vpb.Blog) (map[string]io.Reader, error) {
  directory := map[string]io.Reader{}

  for _, p := range pb.Posts {
    path, r, err := generatePost(*p)
    if err != nil {
      return nil, err
    }
    directory[path] = r
  }

  return directory, nil
}

func generatePost(pb vpb.Post) (string, io.Reader, error) {
  f, err := getComponentFiles()
  if err != nil {
    return "", nil, err
  }
  b := strings.Builder{}
  f = append(f, "lib/core/template/view/text_post.gohtml")
  t, err := template.ParseFiles(f...)
  if err != nil {
    return "", nil, err
  }
  err = t.Execute(&b, pb)
  if err != nil {
    return "", nil, err
  }
  return "/", strings.NewReader(b.String()), nil
}

func getComponentFiles() ([]string, error) {
  files, err := filepath.Glob("lib/core/template/component/*.gohtml")
  if err != nil {
    return nil, err
  }
  return files, nil
}
