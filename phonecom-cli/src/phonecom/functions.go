package main

import (
  "time"
  "math/rand"
)

func randomString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}

func randomNumber(min int, max int) int {

  rand.Seed(time.Now().UTC().UnixNano())
  return int(rand.Float64()*(float64(max) - float64(min)) + float64(min))
}

func randomNumericString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}

func randomAlphaString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}

func randomSmsId() string {
  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, 36)
  for i := 0; i < 8; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  result[8] = '-'
  for i := 9; i < 13; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  result[13] = '-'
  for i := 14; i < 18; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  result[18] = '-'
  for i := 19; i < 23; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  result[23] = '-'
  for i := 24; i < 36; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}