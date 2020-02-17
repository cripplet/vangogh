// Package vangogh_core_render makes some basic assumptions about the
// structure of a website and implements a rendering engine with
// minimal extensions. The output here is compatible with the
// vangogh_api_util helper functions.
package vangogh_core_render

import (
  "fmt"

  vapi "github.com/cripplet/vangogh/api"
  vpb "github.com/cripplet/vangogh/api/proto"
  vpbc "github.com/cripplet/vangogh/core/proto"
  vcrp "github.com/cripplet/vangogh/core/render_post"
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

type CoreRenderInterface struct {}
func (c CoreRenderInterface) GeneratePages(
    pb vpb.Site) (vapi.RoutingTable, error) {
  return generatePages(pb)
}

func generatePages(pb vpb.Site) (vapi.RoutingTable, error) {
  rt := vapi.RoutingTable{}

  all_posts := []vpb.Post{}
  text_posts := []vpb.Post{}
  photo_posts := []vpb.Post{}

  for _, p := range pb.Posts {
    all_posts = append(all_posts, *p)
    if p.Metadata != nil && p.Metadata.Extension != nil && p.Metadata.Extension.Extension != nil {
      var ext vpbc.PostMetadataExtension
      err := ptypes.UnmarshalAny(p.Metadata.Extension.Extension, &ext)
      if err != nil {
        return nil, err
      }
      switch rc := ext.RenderCategory; rc {
      case vpbc.RenderCategoryEnum_RENDER_CATEGORY_TEXT:
        text_posts = append(text_posts, *p)
      case vpbc.RenderCategoryEnum_RENDER_CATEGORY_PHOTO:
        photo_posts = append(photo_posts, *p)
      }
    }

    path, r, err := vcrp.RenderPost(
        vct.ViewPostData{Site: pb, Content: *p})
    if err != nil {
      return nil, err
    }
    rt[path] = r
  }

  path, r, err := vcrp.RenderPostList(
      vct.ViewPostListData{Site: pb, Content: text_posts}, "/category/text/")
  if err != nil {
    return nil, err
  }
  rt[path] = r

  path, r, err = vcrp.RenderPostList(
      vct.ViewPostListData{Site: pb, Content: photo_posts}, "/category/photography/")
  if err != nil {
    return nil, err
  }
  rt[path] = r

  path, r, err = vcrp.RenderPostList(
      vct.ViewPostListData{Site: pb, Content: all_posts}, "/")
  if err != nil {
    return nil, err
  }
  rt[path] = r

  return rt, nil
}
