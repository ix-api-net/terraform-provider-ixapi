package resources

import (
	"fmt"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

// ErrUnexpectedPolymorphic creates an error, informing the
// user, that the API returned an unexpected polymorphic type.
func ErrUnexpectedPolymorphic(
	resp ixapi.Polymorphic,
	expected string,
) error {
	return fmt.Errorf(
		"the API returned an unexpected `%s` instead of `%s`",
		resp.PolymorphicType(),
		expected,
	)
}
