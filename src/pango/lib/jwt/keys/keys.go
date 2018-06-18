package keys

import "os"

func GetVerifyKey() []byte {
  return []byte(os.Getenv("PANGO_VERIFY_KEY"))
}

func GetSignKey() []byte {
  return []byte(os.Getenv("PANGO_SIGN_KEY"))
}
