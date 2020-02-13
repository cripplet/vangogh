package main

import (
    "fmt"
    vpb "github.com/cripplet/vangogh/api/proto"
    pb "github.com/golang/protobuf/proto"
)


func main() {
    b := &vpb.Blog{
        Posts: []*vpb.Post{
            {
                Metadata: &vpb.PostMetadata{
                },
            },
        },
    }
    fmt.Println(pb.MarshalTextString(b))
}
