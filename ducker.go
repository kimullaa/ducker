package main

import(
    "fmt"
    "strconv"
    "strings"
    "os"
    "io"
    "os/exec"
    "time"
    "log"
)

var out io.Writer = os.Stdout

var duck_ascii_target = []string{
`
              o 
  ___       oo  
 /    \\   o    
<  o   \\       
 \\    |___     
 /    _____\\_  
 |    ''''  \\  
 \\__________|  
 
`}

func getTerminalSize() (width,height int) {
    cmd := exec.Command("stty","size")
    cmd.Stdin = os.Stdin
    out,err := cmd.Output()

    if err != nil {
	log.Fatal(err)
	//default 
        width  = 70
        height = 10
	return 
    }else {
    tmp := strings.Split(string(out)," ");
    fmt.Println(tmp)
    width,err  = strconv.Atoi(tmp[0])
    height,err = strconv.Atoi(tmp[1])
    return 

    }
}


func main(){
    //画面をリフレッシュ
    fmt.Fprintln(out,"\x1b[2J")

    w_width,w_height := getTerminalSize()
    fmt.Println(w_width,w_height)

    duck_ascii := strings.Split(duck_ascii_target[0], "\n")
    for width := w_width; width > 0; width-- {
	for height := 5; height < 5 + len(duck_ascii); height++{
	    //カーソル位置ズラしながら描画
            fmt.Fprintf(out,"\x1b[%d;%dH\x1b[K",height,width)
	    fmt.Print(duck_ascii[height-5])
	}
	time.Sleep(100 * time.Millisecond)
    }

}
