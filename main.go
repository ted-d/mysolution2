package main
import ("fmt";"strings")
func brackets(str string)(bool){
    var brcount = 0
    str1 := "{(["
    str2 := ")}]"
    
    
    for i:=0;i<len(str);i++{
        char := string(str[i])
     res1 := strings.Contains(str1, char)
     res2 := strings.Contains(str2, char)
  
   
        if res1{
             brcount += 1
         }
     if res2{
             brcount -= 1
             if brcount < 0{
                 return false
                
             }
         }
        
    }
return brcount == 0
}
func main(){
    
    var s string
    fmt.Scanln(&s)
    
    if brackets(s){
        fmt.Println("Скобки сбалансированы")
    }else{
        fmt.Println("Скобки несбалансированы")
    }
}
