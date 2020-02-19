package vangogh_core_render_util

import (
  "errors"
  "fmt"
  "html/template"
  "net/url"
  "path/filepath"
  "regexp"
  "strings"

  vpb "github.com/cripplet/vangogh/api/proto"
  vpbc "github.com/cripplet/vangogh/core/proto"
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
      "formatPaginate": FormatPaginateURL,
      "addInt": AddInt,
      "subInt": SubInt,
      "getSocialMediaIconClass": GetSocialMediaIconClass,
      "formatTagPath": FormatTagPath,
      "formatPostPath": FormatPostPath,
  }
}

func GetSocialMediaIconClass(p vpbc.SocialMediaEnum) (string, error) {
  l := map[vpbc.SocialMediaEnum]string{
      vpbc.SocialMediaEnum_SOCIAL_MEDIA_FACEBOOK: "fab fa-facebook-f",
      vpbc.SocialMediaEnum_SOCIAL_MEDIA_INSTAGRAM: "fab fa-instagram",
      vpbc.SocialMediaEnum_SOCIAL_MEDIA_YOUTUBE: "fab fa-youtube",
      vpbc.SocialMediaEnum_SOCIAL_MEDIA_GITHUB: "fab fa-github",
  }
  if c, ok := l[p]; !ok {
    return "", errors.New("Cannot find specified social media icon.")
  } else {
    return c, nil
  }
}

func FormatPaginateURL(pathPrefix string, pageNumber int) string {
  return pathPrefix + fmt.Sprintf("page/%d/", pageNumber)
}

func AddInt(i, j int) int {
  return i + j
}

func SubInt(i, j int) int {
  return i - j
}

// Function UnmarshalExtension is used in the templates as a quick way
// to surface extensions. Normal library code should do this manually,
// e.g. by explicitly declaring the proto type.
//
// var p MyExtension
// ...
//
// TODO(minkezhang): Figure out why invoking this in the library case
// results in the proto.Message type being strictly enforced, causing us
// to be unable to resolve extension fields.
func UnmarshalExtension(pb *any.Any) (proto.Message, error) {
  var p ptypes.DynamicAny
  if err := ptypes.UnmarshalAny(pb, &p); err != nil {
    return nil, err
  }

  return p.Message, nil
}

func FormatTime(f string, pb timestamp.Timestamp) (string, error) {
  t, err := ptypes.Timestamp(&pb)
  if err != nil {
    return "", err
  }

  return t.Format(f), nil
}

func formatURLSafeText(s string) (string, error) {
  r, err := regexp.Compile(illegalTitleCharRegex)
  if err != nil {
    return "", err
  }

  return url.QueryEscape(
      r.ReplaceAllString(
          strings.ReplaceAll(
              strings.ToLower(s), " ", "-"), "")), nil
}

func FormatTagPath(t string) (string, error) {
  p, err := formatURLSafeText(t)
  if err != nil {
    return "", err
  }

  return fmt.Sprintf("/tags/%s/", p), nil
}

func FormatPostPath(p vpb.Post) (string, error) {
  t , err := FormatTime(postPathTimeFormat, *p.Metadata.PublishTimestamp)
  if err != nil {
    return "", err
  }

  pt, err := formatURLSafeText(p.Metadata.Title)
  if err != nil {
    return "", err
  }

  return fmt.Sprintf("/posts/%s/%s/", t, pt), nil
}

func GetComponentFiles() ([]string, error) {
  fs, err := filepath.Glob(componentFilePattern)
  if err != nil {
    return nil, err
  }

  return fs, nil
}
