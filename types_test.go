package gitlab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compareAccessControl(t *testing.T) {

	valueDisabledAccessControl := AccessControlValue(DisabledAccessControl)
	valueEnabledAccessControl := AccessControlValue(EnabledAccessControl)
	valuePrivateAccessControl := AccessControlValue(PrivateAccessControl)
	valuePublicAccessControl := AccessControlValue(PublicAccessControl)

	require.Equal(t, true, valueDisabledAccessControl.compareAccessControl(DisabledAccessControl))
	require.Equal(t, true, valueEnabledAccessControl.compareAccessControl(EnabledAccessControl))
	require.Equal(t, true, valuePrivateAccessControl.compareAccessControl(PrivateAccessControl))
	require.Equal(t, true, valuePublicAccessControl.compareAccessControl(PublicAccessControl))

	require.Equal(t, true, valuePublicAccessControl.IsPublicAccessControl())
	require.Equal(t, true, valueEnabledAccessControl.IsEnabledAccessControl())
	require.Equal(t, true, valuePrivateAccessControl.IsPrivateAccessControl())
	require.Equal(t, true, valuePublicAccessControl.IsPublicAccessControl())
}
