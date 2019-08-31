// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"testing"
)

func TestHttp(t *testing.T) {
	{
		rsp, _ := NewRequest().Post("https://github.com", nil)
		time := rsp.Cost()
		t.Log(time)
		rsp.Body()
		rsp.StatusCode()
	}

	{
		rsp, _ := NewRequest().SetTimeout(10).SetHeaders("k", "v").Post("https://github.com", nil)
		var i int
		rsp.Json(&i)
		t.Log(i)
	}

	{
		rsp, _ := NewRequest().Get("https://github.com")
		time := rsp.Cost()
		t.Log(time)
	}
}
