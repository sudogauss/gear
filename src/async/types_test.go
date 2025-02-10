package async

import (
	"reflect"
	"testing"
)

// ========== Explicit type Unfolding tests ==========
func TestOkDataUnfoldingSuccess(t *testing.T) {
	res := &ok[string]{data: "ok"}
	_, d := res.unfold()
	if d != nil {
		if *d != "ok" {
			t.Fatalf("Incorrect data: expected ok, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestOkDataFullUnfoldingSuccess(t *testing.T) {
	res := &ok[string]{data: "ok"}
	tp, d := res.unfold()

	if tp != OK {
		t.Fatalf("Expected Response to be of type OK, found %s", tp)
	}

	if d != nil {
		if *d != "ok" {
			t.Fatalf("Incorrect data: expected ok, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestErrDataUnfoldingSuccess(t *testing.T) {
	res := &err[string]{data: "err"}
	_, d := res.unfold()
	if d != nil {
		if *d != "err" {
			t.Fatalf("Incorrect data: expected err, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestErrFullUnfoldingSuccess(t *testing.T) {
	res := &err[string]{data: "err"}
	tp, d := res.unfold()

	if tp != ERR {
		t.Fatalf("Expect Response to be of type ERR. found %s", tp)
	}

	if d != nil {
		if *d != "err" {
			t.Fatalf("Incorrect data: expected err, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

// ========== Response type Unfolding tests ==========

func TestOkResponseDataUnfoldingSuccess(t *testing.T) {
	res := func() response[string] {
		return &ok[string]{data: "ok"}
	}()

	var sample_res *ok[string]

	_, d := res.unfold()

	// check correctness of the res type
	if resType := reflect.TypeOf(res); resType != reflect.TypeOf(sample_res) {
		t.Fatalf("Underlying data expected to be of type Ok[string], type %v found", resType)
	}

	// check underlying data
	if d != nil {
		if *d != "ok" {
			t.Fatalf("Incorrect data: expected ok, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestOkResponseFullUnfoldingSuccess(t *testing.T) {
	res := func() response[string] {
		return &ok[string]{data: "ok"}
	}()

	tp, d := res.unfold()

	if tp != OK {
		t.Fatalf("Expected Response to be of type OK, found %s", tp)
	}

	// check underlying data
	if d != nil {
		if *d != "ok" {
			t.Fatalf("Incorrect data: expected ok, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestErrResponseDataUnfoldingSuccess(t *testing.T) {
	res := func() response[string] {
		return &err[string]{data: "err"}
	}()

	var sample_res *err[string]

	_, d := res.unfold()

	// check correctness of the res type
	if resType := reflect.TypeOf(res); resType != reflect.TypeOf(sample_res) {
		t.Fatalf("Underlying data expected to be of type Err[string], type %v found", resType)
	}

	// check underlying data
	if d != nil {
		if *d != "err" {
			t.Fatalf("Incorrect data: expected err, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

func TestErrResponseFullUnfoldingSuccess(t *testing.T) {
	res := func() response[string] {
		return &err[string]{data: "err"}
	}()

	tp, d := res.unfold()

	if tp != ERR {
		t.Fatalf("Expected Response to be of type ERR, found %s", tp)
	}

	// check underlying data
	if d != nil {
		if *d != "err" {
			t.Fatalf("Incorrect data: expected err, got %s", *d)
		}
	} else {
		t.Fatalf("Data must not be nil")
	}
}

// ========== Unknown Response type Unfolding tests ==========

type mockR struct {
	data bool
}

func (a *mockR) unfold() (responseType, *bool) {
	return UNKNOWN, nil
}

func TestUnknownResponseFullUnfoldingFailuret(t *testing.T) {
	res := func() response[bool] {
		return &mockR{data: true}
	}()

	tp, d := res.unfold()

	if tp != UNKNOWN {
		t.Fatalf("Expected Response to be of type UNKNOWN, found %s", tp)
	}

	// check underlying data
	if d != nil {
		t.Fatalf("Expected underlying data to be nil, found %v", *d)
	}
}
