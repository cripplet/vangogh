package main

import (
  "io"
  "io/ioutil"

  vau "github.com/cripplet/vangogh/api/api_util"
  vpb "github.com/cripplet/vangogh/api/proto"
  "github.com/cripplet/vangogh/core/render"
  "github.com/golang/protobuf/proto"
)

var directory map[string][]byte = map[string][]byte{}

type vgInterface struct {}
func (v vgInterface) GeneratePages(
    pb vpb.Site) (map[string]io.Reader, error) {
  return vangogh_core_render.VangoghGenerate(pb)
}

func main() {
  data, err := ioutil.ReadFile("lib/api/proto/testdata/example.textpb")
  if err != nil {
    panic(err)
    return
  }

  pb := vpb.Site{}
  if err = proto.UnmarshalText(string(data), &pb); err != nil {
    panic(err)
    return
  }

  s, err := vau.CreateVangoghHTTPServer(vgInterface{}, pb, "0.0.0.0:8000")
  if err != nil {
    panic(err)
    return
  }

  s.ListenAndServe()
}
