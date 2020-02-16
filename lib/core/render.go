package vangogh_core_render

import (
  "fmt"
  "html/template"
  "io"
  "strings"

  vpb "github.com/cripplet/vangogh/api/proto"
  vpbc "github.com/cripplet/vangogh/core/proto"
  vcru "github.com/cripplet/vangogh/core/render_util"
  vct "github.com/cripplet/vangogh/core/type"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
)

func tempGenerateAnyProto() string {
  p := vpbc.PostMetadataExtension{
    RenderCategory: vpbc.RenderCategoryEnum_RENDER_CATEGORY_PHOTO,
    PhotoUrl: "https://storage.blogzhang.com/photography/processed/0037-086.jpg",
    Camera: "Minolta X-700",
    Lens: "17mm f4 MD W.Rokkor",
    Filters: []string{
        "Nisi 3 Stop Hard Edge GND",
    },
    Film: "Velvia 50",
    Location: "Seattle",
  }

  a, err := ptypes.MarshalAny(&p)
  if err != nil {
    panic(err)
  }
  fmt.Println(a)
  fmt.Println(vcru.UnmarshalExtension(a))

  return proto.MarshalTextString(a)
}

func VangoghGenerate(pb vpb.Blog) (map[string]io.Reader, error) {
  directory := map[string]io.Reader{}

  for _, p := range pb.Posts {
    path, r, err := generatePost(vct.ViewPostData{Blog: pb, Content: *p})
    if err != nil {
      return nil, err
    }
    directory[path] = r
  }

  return directory, nil
}

func generatePost(v vct.ViewPostData) (string, io.Reader, error) {
  f, err := vcru.GetComponentFiles()
  if err != nil {
    return "", nil, err
  }
  b := strings.Builder{}
  f = append(f, "lib/core/template/view/post.gohtml")
  t, err := template.New("").Funcs(
    template.FuncMap{
      "deserialize": vcru.UnmarshalExtension,
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
