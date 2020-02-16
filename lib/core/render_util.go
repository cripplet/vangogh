package vangogh_core_render_util

import (
  "fmt"
  "html/template"
  "net/url"
  "path/filepath"
  "regexp"
  "strings"

  vpb "github.com/cripplet/vangogh/api/proto"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/ptypes/any"
  "github.com/golang/protobuf/ptypes/timestamp"
)

var illegalTitleCharRegex string = "[^A-Za-z0-0\\-]+"
var postPathTimeFormat string = "2006/01/02"
var componentFilePattern string = "lib/core/template/component/*.gohtml"

func GetVangoghCoreTemplateFuncMap() template.FuncMap {
  return template.FuncMap{
      "deserialize": UnmarshalExtension,
      "formatTime": FormatTime,
  }
}

func UnmarshalExtension(pb *any.Any) (proto.Message, error) {
  var p ptypes.DynamicAny
  err := ptypes.UnmarshalAny(pb, &p)
  if err != nil {
    return nil, err
  }

  return p.Message, err
}


func FormatTime(f string, pb timestamp.Timestamp) (string, error) {
  t, err := ptypes.Timestamp(&pb)
  if err != nil {
    return "", err
  }

  return t.Format(f), nil
}

func FormatPostPath(p vpb.Post) (string, error) {
  r, err := regexp.Compile(illegalTitleCharRegex)
  if err != nil {
    return "", err
  }

  t , err := FormatTime(postPathTimeFormat, *p.Metadata.PublishTimestamp)
  if err != nil {
    return "", err
  }

  return fmt.Sprintf(
      "/posts/%s/%s/",
      t,
      url.QueryEscape(
          r.ReplaceAllString(
              strings.ReplaceAll(
                  strings.ToLower(p.Metadata.Title), " ", "-"), ""))), nil
}

func GetComponentFiles() ([]string, error) {
  fs, err := filepath.Glob(componentFilePattern)
  if err != nil {
    return nil, err
  }

  return fs, nil
}
