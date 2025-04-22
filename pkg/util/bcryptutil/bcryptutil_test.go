package bcryptutil

import "testing"

func TestHashPassword(t *testing.T) {
	hashedPassword := "$2a$10$GcRzupsOIrPgqOoOLtsZH.spUY0uoiM8NWmWX4o39WHqwmMdaJ2.6"
	password := "validpasswd"

	t.Errorf("%v", ComparePassword(hashedPassword, password))
}
