package kata

import( 
"strings"
"strconv"
)

func factorial(v int) int {  
    var factVal = 1
    for i:=1; i<=v; i++ {
        factVal *= i 
    }    
    return factVal  
}

func Strong(n int) string {
  s:= strconv.Itoa(n)
  str := strings.Split(s, "")
  strLen := len(s)
  var fac int
  var result int
  for i := 0; i < strLen ; i++{
    x, _ := strconv.Atoi(str[i])
    fac = factorial(x)
    result += fac
  }
  if result != n {
    return "Not Strong !!"
  }
  return "STRONG!!!!"
}
