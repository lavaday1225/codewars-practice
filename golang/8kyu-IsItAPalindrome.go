package kata

import "strings"

func IsPalindrome(str string) bool {
  str = strings.ToLower(str)
  n := len(str)
  for i := 0; i < n; i++ {
    if str[i] != str[n-i-1] {
     return false
    }
  }
  return true
}
