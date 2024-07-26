package encryption

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	res, err := AESEncrypt("hello", "12345678")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	t.Log("---------------------")
	res, err = AESDecrypt(res, "12345678")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
