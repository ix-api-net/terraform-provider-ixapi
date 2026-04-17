package resources

import "testing"

func TestMacRequestFromResourceData(t *testing.T) {
	resource := NewMACResource()
	res := resource.Data(nil)
	_ = res.Set("managing_account", "managing:1")
	_ = res.Set("consuming_account", "consuming:2")
	_ = res.Set("address", "23:42:12:aa:bb:cc")
	_ = res.Set("valid_not_before", "2022-10-11T12:23:42Z")

	req, err := macRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Address != "23:42:12:aa:bb:cc" {
		t.Error("unexpected request data:", req)
	}
}
