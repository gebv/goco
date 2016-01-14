package main

import "testing"

func TestToLower(t *testing.T) {
	if v := toLower("FirstName", "_"); "first_name" != v {
		t.Errorf("not expected value: %s", v)
	}

	if v := toLower("First    Name", "_"); "first_name" != v {
		t.Errorf("not expected value: %s", v)
	}

	if v := toLower("First  a  Name", "_"); "first_a_name" != v {
		t.Errorf("not expected value: %s", v)
	}

	if v := toLower("FirST  a  Name", "_"); "fir_s_t_a_name" != v {
		t.Errorf("not expected value: %s", v)
	}
}
