package types

// DefaultAdmin represents the Admin default value.
// TODO: Determine the default value.
var DefaultAdmin string = "admin"

// NewParams creates a new Params instance.
func NewParams(
	admin string,
) Params {
	return Params{
		Admin: admin,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultAdmin,
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {
	if err := validateAdmin(p.Admin); err != nil {
		return err
	}

	return nil
}

// validateAdmin validates the Admin parameter.
func validateAdmin(v string) error {
	// TODO implement validation
	return nil
}
