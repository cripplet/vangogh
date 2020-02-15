package vangogh_core_render

import (
  "fmt"
  "html/template"
  "io"
  "net/url"
  "path/filepath"
  "regexp"
  "strings"
  "time"

  vpb "github.com/cripplet/vangogh/api/proto"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/ptypes/any"
  "github.com/golang/protobuf/ptypes/timestamp"
)

type ViewPostData struct {
  Blog vpb.Blog
  // Page-specific data to be used in generating the content template.
  Content vpb.Post
}

func deserializeAnyProto(pb *any.Any) proto.Message {
  var p ptypes.DynamicAny
  ptypes.UnmarshalAny(pb, &p)
  return p.Message
}

func formatTime(f string, pb timestamp.Timestamp) string {
  t := time.Unix(pb.Seconds, int64(pb.Nanos))
  return t.Format(f)
}

func VangoghGenerate(pb vpb.Blog) (map[string]io.Reader, error) {
  directory := map[string]io.Reader{}

  for _, p := range pb.Posts {
    path, r, err := generatePost(ViewPostData{Blog: pb, Content: *p})
    if err != nil {
      return nil, err
    }
    directory[path] = r
  }

  return directory, nil
}

func formatTitlePath(t timestamp.Timestamp, title string) (string, error) {
  r, err := regexp.Compile("[^a-zA-Z0-9\\-]+")
  if err != nil {
    return "", err
  }

  return fmt.Sprintf(
      "/posts/%s/%s/",
      formatTime("2006/01/02", t),
      url.QueryEscape(
          r.ReplaceAllString(
              strings.ReplaceAll(
                  strings.ToLower(title), " ", "-"), ""))), nil
}

func generatePost(v ViewPostData) (string, io.Reader, error) {
  f, err := getComponentFiles()
  if err != nil {
    return "", nil, err
  }
  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post.gohtml")
  t, err := template.New("").Funcs(
    template.FuncMap{
      "deserialize": deserializeAnyProto,
      "formatTime": formatTime,
    },
  ).ParseFiles(f...)
  if err != nil {
    return "", nil, err
  }


  err = t.ExecuteTemplate(&b, "page", v)
  if err != nil {
    return "", nil, err
  }

  path, err := formatTitlePath(*v.Content.Metadata.PublishTimestamp, v.Content.Metadata.Title)
  if err != nil {
    return "", nil, err
  }
  fmt.Println(path)
  return path, strings.NewReader(b.String()), nil
}

func getComponentFiles() ([]string, error) {
  files, err := filepath.Glob("lib/core/template/component/*.gohtml")
  if err != nil {
    return nil, err
  }
  return files, nil
}
