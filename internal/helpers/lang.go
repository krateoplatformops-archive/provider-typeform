package helpers

// IntValue converts the supplied int pointer to an int, returning zero if
// the pointer is nil.
func IntValue(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// IntPtr converts the supplied int to a pointer to that int.
func IntPtr(p int) *int { return &p }

// BoolPtr converts the supplied bool to a pointer to that bool
func BoolPtr(p bool) *bool { return &p }

// BoolValue converts the supplied bool pointer to an bool, returning false if
// the pointer is nil.
func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// StringPtr converts the supplied string to a pointer to that string.
func StringPtr(p string) *string { return &p }

// StringValue converts the supplied string pointer to a string, returning the
// empty string if the pointer is nil.
func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// LateInitializeBool implements late initialization for bool type.
func LateInitializeBool(b *bool, from bool) *bool {
	if b != nil || !from {
		return b
	}
	return &from
}

// LateInitializeInt implements late initialization for int type.
func LateInitializeInt(i *int, from int) *int {
	if i != nil || from == 0 {
		return i
	}
	return &from
}

// LateInitializeString implements late initialization for string type.
func LateInitializeString(s *string, from string) *string {
	if s != nil || from == "" {
		return s
	}
	return &from
}
