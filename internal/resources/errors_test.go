package resources

import (
	"strings"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestErrUnexpectedPolymorphic(t *testing.T) {
	err := ErrUnexpectedPolymorphic(
		&ixapi.MP2MPNetworkService{},
		ixapi.P2PNetworkServiceType,
	)
	s := err.Error()
	if !strings.Contains(s, "p2p_vc") {
		t.Fatal("unexpected err:", err)
	}
	if !strings.Contains(s, "mp2mp_vc") {
		t.Fatal("unexpected err:", err)
	}
}
