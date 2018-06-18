package keys

import (
  "os"
  "reflect"
  "testing"
)

func TestGetKeys(t *testing.T) {
  os.Setenv("PANGO_SIGN_KEY", "foo")
  os.Setenv("PANGO_VERIFY_KEY", "bar")
  if !reflect.DeepEqual(GetVerifyKey(), []byte("bar")) {
    t.Errorf("GetVerifyKey not working, expected bar got", GetVerifyKey())
  }
  if !reflect.DeepEqual(GetSignKey(), []byte("foo")) {
    t.Errorf("GetSignKey expected foo, got", GetSignKey())
  }
}
