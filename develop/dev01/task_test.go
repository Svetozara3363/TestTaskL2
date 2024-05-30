package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// Проверка, что функция не возвращает пустую строку
func TestOsTimeNotEmpty(t *testing.T) {
	sysTime := osTime()
	require.NotEmpty(t, sysTime)
}

// Проверка, что возразаемое знаечение времени не является нулевым
func TestOsTimeNotZero(t *testing.T) {
	sysTime := osTime()
	require.NotEqual(t, sysTime, time.Time{})
}

// Проверка, что функция не возвращает пустую строку
func TestNetTimeNotEmpty(t *testing.T) {
	netTime, _ := netTime()
	require.NotEmpty(t, netTime)
}

// Проверка, что возразаемое знаечение времени не является нулевым
func TestNetTimeNotZero(t *testing.T) {
	netTime, _ := netTime()
	require.NotEqual(t, netTime, time.Time{})
}

// Проверка, что функция не возвращает пустую строку
func TestNetTimeDurationTimeNotEmpty(t *testing.T) {
	netDTime, _ := netTimeWithDuration()
	require.NotEmpty(t, netDTime)
}

// Проверка, что возразаемое знаечение времени не является нулевым
func TestNetTimeDurationTimeNotZero(t *testing.T) {
	netDTime, _ := netTimeWithDuration()
	require.NotEqual(t, netDTime, time.Time{})
}
