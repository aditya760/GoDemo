package main  
import "fmt"  

func main() {  
   fmt.Printf("Result::",check(1))
 } 

func check(a int) bool{
    if( a % 2==0 ) {    
        return false
    }else{
       return true
    }
}
 

