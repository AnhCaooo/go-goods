package helpers

import (
	"testing"
)

func BenchmarkSetTime(b *testing.B) {
	for b.Loop() {
		_, err := SetTime(10, 30)
		if err != nil {
			b.Fatalf("SetTime error: %v", err)
		}
	}
}

func BenchmarkLoadHelsinkiLocation(b *testing.B) {
	for b.Loop() {
		_, err := LoadHelsinkiLocation()
		if err != nil {
			b.Fatalf("LoadHelsinkiLocation error: %v", err)
		}
	}
}

func BenchmarkGetTodayDate(b *testing.B) {
	for b.Loop() {
		_ = GetTodayDate()
	}
}

func BenchmarkGetTomorrowDate(b *testing.B) {
	for b.Loop() {
		_ = GetTomorrowDate()
	}
}

func BenchmarkGetYesterdayDate(b *testing.B) {
	for b.Loop() {
		_ = GetYesterdayDate()
	}
}

func BenchmarkParseHour(b *testing.B) {
	for b.Loop() {
		_, _ = ParseHour("12:03")
	}
}

func BenchmarkParseHourMinute(b *testing.B) {
	for b.Loop() {
		_, _, _ = ParseHourMinute("12:03")
	}
}
