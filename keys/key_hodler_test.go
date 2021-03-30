package keys

import "testing"

func TestTIsAccount(t *testing.T) {
	addr1 := "10123123123"
	addr2 := "0126acfaed3d827dd696c7f15cef0c6f247860cabaec133250240601bbc3d2929510123123123"
	addr3 := "t126acfaed3d827dd696c7f15cef0c6f247860cabaec133250240601bbc3d29295"

	addr4 := "0126acfaed3d827dd696c7f15cef0c6f247860cabaec133250240601bbc3d29295"
	addr5 := "0x0126acfaed3d827dd696c7f15cef0c6f247860cabaec133250240601bbc3d29295"

	addr6 := "0x0776acfaed3d827dd696c7f15cef0c6f247860cabaec133250240601bbc3d29295"

	addr7 := "0203447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1"

	if IsAccount(addr1) {
		t.Fatal("failed to test: addr len too small")
	}
	if IsAccount(addr2) {
		t.Fatal("failed to test: addr len too long")
	}
	if IsAccount(addr3) {
		t.Fatal("failed to test: addr not all hex")
	}
	if !IsAccount(addr4) {
		t.Fatal("failed to test: addr(66) without 0x")
	}
	if !IsAccount(addr5) {
		t.Fatal("failed to test: addr(68) with 0x")
	}
	if IsAccount(addr6) {
		t.Fatal("failed to test: addr(68) with 0x but invalid prefix")
	}
	if !IsAccount(addr7) {
		t.Fatal("failed to test: secp256k1 addr")
	}
}
