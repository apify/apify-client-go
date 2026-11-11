package apify

import (
	"context"
	"strings"
	"testing"
)

// Charge must always send an idempotency-key header so a transport retry cannot double-charge.
func TestChargeSendsIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "my-event"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	key := backend.lastHeaders.Get("idempotency-key")
	if key == "" {
		t.Fatal("expected an idempotency-key header to be sent")
	}
	if !strings.HasPrefix(key, "run123-my-event-") {
		t.Fatalf("idempotency key should embed run id and event name, got %q", key)
	}
	if !strings.Contains(backend.lastBody, `"eventName":"my-event"`) {
		t.Fatalf("unexpected charge body: %s", backend.lastBody)
	}
	if !strings.Contains(backend.lastURL, "/actor-runs/run123/charge") {
		t.Fatalf("expected charge request against /actor-runs/run123/charge, got %q", backend.lastURL)
	}
}

// A caller-supplied idempotency key must be used verbatim.
func TestChargeHonorsExplicitIdempotencyKey(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "e", IdempotencyKey: "fixed-key"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	if got := backend.lastHeaders.Get("idempotency-key"); got != "fixed-key" {
		t.Fatalf("expected explicit idempotency key, got %q", got)
	}
}

// RequestID must be sent as requestId in the charge body, and omitted when empty.
func TestChargeSendsRequestID(t *testing.T) {
	backend := &mockBackend{responses: okData()}
	client := testClient(backend, 0)

	err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "e", RequestID: "req-1"})
	if err != nil {
		t.Fatalf("charge: %v", err)
	}
	if !strings.Contains(backend.lastBody, `"requestId":"req-1"`) || !strings.Contains(backend.lastBody, `"count":1`) {
		t.Fatalf("unexpected charge body: %s", backend.lastBody)
	}

	if err := client.Run("run123").Charge(context.Background(), RunChargeOptions{EventName: "e"}); err != nil {
		t.Fatalf("charge: %v", err)
	}
	if strings.Contains(backend.lastBody, "requestId") {
		t.Fatalf("requestId must be omitted when empty, got %s", backend.lastBody)
	}
}

// Run.GetWithWait must forward the waitForFinish query parameter.
func TestRunGetWithWaitForwardsParam(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"r1"}}`)}
	client := testClient(backend, 0)

	secs := int64(30)
	if _, _, err := client.Run("r1").GetWithWait(context.Background(), &secs); err != nil {
		t.Fatalf("get with wait: %v", err)
	}
	if !strings.Contains(backend.lastURL, "waitForFinish=30") {
		t.Fatalf("expected waitForFinish=30 in URL, got %q", backend.lastURL)
	}
}

// GetStreamedLog must open a raw, streaming log request (stream=1 & raw=1).
func TestGetStreamedLogParams(t *testing.T) {
	backend := &mockBackend{responses: constant(200, "live log bytes")}
	client := testClient(backend, 0)

	stream, err := client.Run("r1").GetStreamedLog(context.Background())
	if err != nil {
		t.Fatalf("get streamed log: %v", err)
	}
	defer func() { _ = stream.Close() }()
	if !strings.Contains(backend.lastURL, "stream=1") {
		t.Fatalf("expected stream=1 in URL, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "raw=1") {
		t.Fatalf("expected raw=1 in URL, got %q", backend.lastURL)
	}
}

// ActorClient.StartRaw sends the given bytes verbatim (no JSON serialization), defaulting to
// application/octet-stream when the caller does not set options.ContentType.
func TestStartRawSendsBytesVerbatimWithDefaultContentType(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"run1"}}`)}
	client := testClient(backend, 0)

	input := []byte("not json, just bytes: \x00\x01\x02")
	if _, err := client.Actor("me/actor").StartRaw(context.Background(), input, ActorStartOptions{}); err != nil {
		t.Fatalf("start raw: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Type"); got != "application/octet-stream" {
		t.Fatalf("expected default content type application/octet-stream, got %q", got)
	}
	if backend.lastBody != string(input) {
		t.Fatalf("expected the raw bytes to be sent verbatim, got %q", backend.lastBody)
	}
}

// RunClient.MetamorphRaw sends the given bytes verbatim and still forwards targetActorId/build
// as query parameters like Metamorph.
func TestMetamorphRawSendsBytesVerbatim(t *testing.T) {
	backend := &mockBackend{responses: constant(200, `{"data":{"id":"run1"}}`)}
	client := testClient(backend, 0)

	input := []byte("raw metamorph input")
	_, err := client.Run("run1").MetamorphRaw(context.Background(), "other/actor", input, MetamorphOptions{Build: "1.0"})
	if err != nil {
		t.Fatalf("metamorph raw: %v", err)
	}
	if got := backend.lastHeaders.Get("Content-Type"); got != "application/octet-stream" {
		t.Fatalf("expected default content type application/octet-stream, got %q", got)
	}
	if backend.lastBody != string(input) {
		t.Fatalf("expected the raw bytes to be sent verbatim, got %q", backend.lastBody)
	}
	if !strings.Contains(backend.lastURL, "targetActorId=other~actor") {
		t.Fatalf("expected targetActorId in URL, got %q", backend.lastURL)
	}
	if !strings.Contains(backend.lastURL, "build=1.0") {
		t.Fatalf("expected build in URL, got %q", backend.lastURL)
	}
}
