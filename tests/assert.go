package tests

import "testing"

func AssertEqualCode(t testing.TB, got int, want int) {
	t.Helper()
	if want != got {
		t.Errorf("got %d, want %d \n", got, want)
	}
}

func AssertEqual(t testing.TB, got string, want string) {
	t.Helper()
	if want != got {
		t.Errorf("got %s, want %s \n", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("didn't expect error, but got one, err:%+v \n", err)
	}
}
