// Package vangogh_api_util provides QoL functions for Vangogh.
// Implementations may freely call these functions to help surface
// generated data.
package vangogh_api_util

import (
  "bytes"
  "fmt"
  "net/http"

  "github.com/cripplet/vangogh/api"
  vpb "github.com/cripplet/vangogh/api/proto"
)

// Function vangoghHTTPServerHandler returns a mux for emulating serving
// static files. The input routing table is a map of partial paths,
// e.g. "/posts/2010/12/31/happy-new-year" to the HTML content of the page.
func vangoghHTTPServerHandler(
  rs map[string][]byte) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    body, is_found := rs[r.URL.Path]
    // TODO(minkezhang): Make this configurable as a page.
    if !is_found {
      http.Error(w, "404 NOT FOUND", http.StatusNotFound)
      return
    }

    _, err := w.Write(body)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }
}

// Function CreateVangoghHTTPServer creates a valid http.Server object
// which emulates a static file server (without actually committing to
// creating the files).
func CreateVangoghHTTPServer(
    r vangogh_api.VangoghRenderer,
    pb vpb.Site,
    address string) (http.Server, error) {
  var routes map[string][]byte = map[string][]byte{}

  m, err := r.GeneratePages(pb)
  if err != nil {
    return http.Server{}, err
  }

  for p, reader := range m {
    fmt.Println(p)

    b := bytes.Buffer{}
    b.ReadFrom(reader)
    routes[p] = b.Bytes()
  }

  mux := http.NewServeMux()
  mux.HandleFunc("/", vangoghHTTPServerHandler(routes))

  return http.Server{
    Addr: address,
    Handler: mux,
  }, nil
}
