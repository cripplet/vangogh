package main

import (
  "io"
  "io/ioutil"
  "net/http"
  "github.com/cripplet/vangogh/core/render"
  vpb "github.com/cripplet/vangogh/api/proto"
  "github.com/golang/protobuf/proto"
)

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe("0.0.0.0:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  data, err := ioutil.ReadFile("lib/api/proto/testdata/example.textpb")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  pb := vpb.Blog{}
  err = proto.UnmarshalText(string(data), &pb)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  directory, err := vangogh_core_render.VangoghGenerate(pb)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "text/html")
  _, err = io.Copy(w, directory[""])
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
