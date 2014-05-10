package main

import "testing"

func chop(s string) string {
	return s[0 : len(s)-1]
}

func TestRootHandler(t *testing.T) {
	recorder := newTestRequest("GET", "/", "")

	func() {
		actual := recorder.Code
		expected := 200
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()
}

func TestListHandler(t *testing.T) {
	recorder := newTestRequest("GET", "/list", "")

	func() {
		actual := recorder.Code
		expected := 200
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()

	func() {
		actual := chop(string(recorder.Body.Bytes()))
		expected := `{"success":true,"message":"","result":[]}`
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()
}

func TestCreateHandler(t *testing.T) {
	recorder := newTestRequest("POST", "/create", "title=TODO-1")

	func() {
		actual := recorder.Code
		expected := 200
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()

	func() {
		actual := chop(string(recorder.Body.Bytes()))
		expected := `{"success":true,"message":"","result":[{"id":1,"title":"TODO-1","completed":false}]}`
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()
}

func TestSwitchHandler(t *testing.T) {
	recorder := newTestRequest("POST", "/switch/1", "")

	func() {
		actual := recorder.Code
		expected := 200
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()

	func() {
		actual := chop(string(recorder.Body.Bytes()))
		expected := `{"success":true,"message":"","result":[{"id":1,"title":"TODO-1","completed":true}]}`
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()
}

func TestDeleteHandler(t *testing.T) {
	recorder := newTestRequest("POST", "/delete/1", "")

	func() {
		actual := recorder.Code
		expected := 200
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()

	func() {
		actual := chop(string(recorder.Body.Bytes()))
		expected := `{"success":true,"message":"","result":[]}`
		if actual != expected {
			t.Errorf("Expect %q but %q", expected, actual)
		}
	}()
}
