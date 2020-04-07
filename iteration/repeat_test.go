package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertRepeat := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("repeat 'a' five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeat(t, expected, repeated)
	})

	t.Run("repeat 'b' three times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"
		assertRepeat(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
