package common

import (
	"encoding/json"
	"testing"
)

func TestVersionAPIError(t *testing.T) {
	responseRaw := []byte("{\"Response\":{\"Error\":{\"Code\":\"InternalError\",\"Message\":\"An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the Tencent cloud forums.\"},\"RequestId\":\"request-id-mock\"}}")

	versionErrorResponse := VersionAPIError{}

	err := json.Unmarshal(responseRaw, &versionErrorResponse)
	if err != nil {
		t.Fatal(err)
	}

	if (versionErrorResponse.Response.Error.Code != "") != true {
		t.Fatal("unable to detect versioned api error.")
	}
}