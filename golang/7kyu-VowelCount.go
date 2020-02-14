package kata

import "strings"

func GetCount(str string) (count int) {
  spl := strings.Split(str, "")
  count = 0
  for _, s := range spl {
    if s == "a" || s == "e" || s == "u" || s == "i"|| s == "o" {
      count = count + 1
    }
  }
  return count
}

//cod3x, aytrack, kroppt, siliconbrain, SlavaDaun, postannihilation
package kata

func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}

// use switch-case can be more efficient to judge "or"
