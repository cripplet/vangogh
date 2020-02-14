package main

import (
  "bytes"
  "io"
  "io/ioutil"
  "net/http"
  "github.com/cripplet/vangogh/core/render"
  vpb "github.com/cripplet/vangogh/api/proto"
  "github.com/golang/protobuf/proto"
)

var directory map[string][]byte = map[string][]byte{}

func main() {
  var data []byte
  var err error

  var directory_readers map[string]io.Reader

  data, err = ioutil.ReadFile("lib/api/proto/testdata/example.textpb")
  if err != nil {
    panic(err)
    return
  }

  pb := vpb.Blog{}
  if err = proto.UnmarshalText(string(data), &pb); err != nil {
    panic(err)
    return
  }

  directory_readers, err = vangogh_core_render.VangoghGenerate(pb)
  if err != nil {
    panic(err)
    return
  }

  for u, r := range directory_readers {
    b := bytes.Buffer{}
    b.ReadFrom(r)
    directory[u] = b.Bytes()
  }

  http.HandleFunc("/", handler)
  http.ListenAndServe("0.0.0.0:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")

  b, is_found := directory[r.URL.Path]
  // TODO(minkezhang): Make this configurable as a page.
  if !is_found {
    http.Error(w, "404 NOT FOUND", http.StatusNotFound)
    return
  }

  _, err := w.Write(b)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
