package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)

  dp := make([]string, n+1)
  dp[1] = "SK"
  if n >= 2 {
    dp[2] = "CY"
  }
  if n >= 3 {
    dp[3] = "SK"
  }
  if n >= 4 {
    dp[4] = "SK"
  }

  for i := 5; i <= n; i++ {
    if dp[i-1] == "CY" || dp[i-3] == "CY" || dp[i-4] == "CY" {
      dp[i] = "SK"
    } else {
      dp[i] = "CY"
    }
  }
  fmt.Println(dp[n])
}