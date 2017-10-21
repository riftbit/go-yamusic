package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_GetStatus(t *testing.T) {
	setup()
	defer teardown()

	want := &AccountStatus{}
	want.InvocationInfo.ReqID = "Account.GetStatus"

	mux.HandleFunc("/account/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Account().GetStatus(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
