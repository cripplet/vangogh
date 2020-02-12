package main

import (
    "fmt"
    vpb "github.com/cripplet/vangogh/proto"
    pb "github.com/golang/protobuf/proto"
)


func main() {
    b := &vpb.Blog{
        Posts: []*vpb.Post{
            {
                Metadata: &vpb.PostMetadata{
                    RenderCategory: vpb.RenderCategoryEnum_RENDER_CATEGORY_TEXT,
                },
            },
        },
    }
    fmt.Println(pb.MarshalTextString(b))
}
