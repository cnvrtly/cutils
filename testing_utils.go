package cutils

import (
	"testing"
	"net/http"
	"io"
	"google.golang.org/appengine/aetest"
)

func Testing_InitAppengineRequest(t *testing.T) ( *http.Request, io.Closer) {
	inst, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	req1, err := inst.NewRequest("POST", "/gophers", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	return req1,inst
}

