package arkeocli

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arkeonetwork/arkeo/common/cosmos"
)

func TestParseBondAmount(t *testing.T) {
	amount, err := parseBondAmount("100uuram")
	require.NoError(t, err)
	require.Equal(t, cosmos.NewInt(100), amount)

	amount, err = parseBondAmount("-100uuram")
	require.NoError(t, err)
	require.Equal(t, cosmos.NewInt(-100), amount)

	_, err = parseBondAmount("100arkeo")
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("-100arkeo")
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("100")
	require.Error(t, err)
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("-100")
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("")
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("uuram")
	require.ErrorContains(t, err, "bad bond amount")

	_, err = parseBondAmount("uuram")
	require.ErrorContains(t, err, "bad bond amount")

	_, err = parseBondAmount("Paul Revere")
	require.ErrorContains(t, err, "bad bond denom")

	_, err = parseBondAmount("100uuram 100uuram")
	require.ErrorContains(t, err, "bad bond denom")
}
