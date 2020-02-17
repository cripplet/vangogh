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

func generateAllPostList(pb vpb.Site) ([]vapi.RoutingTableRow, error) {
  posts := []vpb.Post{}
  for _, p := range pb.Posts {
    posts = append(posts, *p)
  }

  return vcrp.RenderPostList(
      vct.ViewPostListData{Site: pb, Content: posts}, "/")
}

func generatePhotoPostList(pb vpb.Site) ([]vapi.RoutingTableRow, error) {
  posts := []vpb.Post{}

  for _, p := range pb.Posts {
    if (
        p.Metadata != nil &&
        p.Metadata.Extension != nil &&
        p.Metadata.Extension.Extension != nil) {
      var ext vpbc.PostMetadataExtension
      err := ptypes.UnmarshalAny(p.Metadata.Extension.Extension, &ext)
      if err != nil {
        return nil, err
      }

      if ext.RenderCategory == vpbc.RenderCategoryEnum_RENDER_CATEGORY_PHOTO {
        posts = append(posts, *p)
      }
    }
  }

  return vcrp.RenderPostList(
      vct.ViewPostListData{Site: pb, Content: posts}, "/photography/")
}

func generatePages(pb vpb.Site) (vapi.RoutingTable, error) {
  rt := vapi.RoutingTable{}
  rs := []vapi.RoutingTableRow{}

  for _, p := range pb.Posts {
    r, err := vcrp.RenderPost(
        vct.ViewPostData{Site: pb, Content: *p})
    if err != nil {
      return nil, err
    }
    rs = append(rs, r)
  }

  trs, err := generatePhotoPostList(pb)
  if err != nil {
    return nil, err
  }
  rs = append(rs, trs...)

  trs, err = generateAllPostList(pb)
  if err != nil {
    return nil, err
  }
  rs = append(rs, trs...)

  if rt.AddRoutes(rs) != nil {
    return nil, err
  }

  return rt, nil
}
