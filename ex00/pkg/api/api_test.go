// package api
package api

import (
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"
	"omikuji/pkg/omikuji"
)

func TestSetupServer(t *testing.T) {

	r := SetupServer()

	// テスト用のサーバーを立てる
	ts := httptest.NewServer(r)
	defer ts.Close()

	// テスト用のサーバーにリクエストを投げる
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

	// レスポンスボディの内容を検証するには、
	// 応答をデコードして適切なフィールドをチェックします。
	var result omikuji.OmikujiResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to parse the response body: %v", err)
	}

	if result.Fortune == "" || result.Health == "" || result.Love == "" || result.Residence == "" || result.Study == "" || result.Travel == "" {
		t.Errorf("Expected the fortune; got empty string")
	} else {
		t.Logf("\nFortune: %s", result.Fortune)
	}
}
