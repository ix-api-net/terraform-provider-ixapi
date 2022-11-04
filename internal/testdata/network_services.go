package testdata

import "github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"

// NewP2PNetworkService makes a testing p2p network service
func NewP2PNetworkService() *ixapi.P2PNetworkService {
	return &ixapi.P2PNetworkService{
		Type:                 ixapi.P2PNetworkServiceType,
		ManagingAccount:      "managing:123",
		ConsumingAccount:     "consuming:123",
		BillingAccount:       "billing:123",
		ProductOffering:      "product:123",
		JoiningMemberAccount: "joining:123",
		DisplayName:          NewOptString("p2p network service"),
	}
}
