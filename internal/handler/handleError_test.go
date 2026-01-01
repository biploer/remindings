package handler

import (
	"net/http"
	"testing"
)

func TestHandle(t *testing.T) {
	pre := func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
	after := handle(pre)

	_ = after
	// if after != http.HandlerFunc {
	// 	t.Errorf("handle func can`t")
	// }
}
