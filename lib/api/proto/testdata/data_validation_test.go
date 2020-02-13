package vangogh_api_proto_testdata

import (
    "io/ioutil"
    "testing"
    "github.com/cripplet/vangogh/api/proto"
    "github.com/golang/protobuf/proto"
)

var example_textpb_files []string = []string{
    "example.textpb",
}

func TestDataValidation(t *testing.T) {
    for _, fn := range example_textpb_files {
        data, err := ioutil.ReadFile(fn)
        if err != nil {
            t.Errorf("Could not read file \"%s\": %s", fn, err)
        }
        pb := vangogh_api_proto_base.Blog{}
        err = proto.UnmarshalText(string(data), &pb)
        if err != nil {
            t.Errorf("Could not load textproto, exited with error: %s", err)
        }
    }
}
