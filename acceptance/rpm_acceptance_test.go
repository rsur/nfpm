package acceptance

import "testing"

func TestSimpleRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		accept(t, "simple_rpm", "simple.yaml", "rpm", "rpm.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		accept(t, "simple_rpm_386", "simple.386.yaml", "rpm", "rpm.386.dockerfile")
	})
}

func TestComplexRPM(t *testing.T) {
	t.Run("amd64", func(t *testing.T) {
		accept(t, "complex_rpm", "complex.yaml", "rpm", "rpm.dockerfile")
	})
	t.Run("i386", func(t *testing.T) {
		accept(t, "complex_rpm_386", "complex.386.yaml", "rpm", "rpm.386.dockerfile")
	})
}
