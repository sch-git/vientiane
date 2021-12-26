package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vientiane/server/utils"
)

func TestAES(t *testing.T) {
	psd := "1234"
	encodePsd, err := utils.AESEncrypt(psd)
	if nil != err {
		t.Log(err)
		return
	}
	t.Logf("encode psd: %s", encodePsd)

	decodePsd, err := utils.AESDecrypt(encodePsd)
	if nil != err {
		t.Log(err)
		return
	}

	t.Log(assert.Equal(t, psd, decodePsd))
}
