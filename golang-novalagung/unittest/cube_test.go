package main

import "testing"
import "github.com/stretchr/testify/assert"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "Perhitungan Volume Salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
