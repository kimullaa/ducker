package main

import(
    "fmt"
    "strconv"
    "strings"
    "os"
    "flag"
    "io"
    "os/exec"
    "time"
    "log"
)

const (
    RUN = "run"
    PS = "ps"
    PULL = "pull"
)

var out io.Writer = os.Stdout

// MAP[COMMAND , AsciiArt]
var m  = map[string] []string {
RUN : strings.Split(
`        
          
  ___   
 / "" \\  I'm a highway star
<  o   \     
 \     |___     
 /    _____\<      3
 |    ''''" \  ＝＝ 3
 \\_________/      3
     \ \ 
     ` , "\n"),
PULL : strings.Split(
`        
             
         ___    someone is pulling me
        / "" \
<<<<<<<<  o   \     
        \     |___     
        /    _____\<   
        |    ''''" ~\  
        \\~~~~~~~~~~/ 
        
     ` , "\n"),
PS : strings.Split(
`        
        
  ___    
 / "" \\  p.s. my holiday has gone to waste
<  o   \   
 \     |___     
 /    _____\<
 |    ''''" \ 
 \\_________/
 ～～～～～～～～～～～
     ` , "\n"),
}


// return current terminal window (width,height)
func getTerminalSize() (width,height int) {
    cmd := exec.Command("stty","size")
    cmd.Stdin = os.Stdin
    out,err := cmd.Output()

    if err != nil {
	log.Fatal(err)
	//default size
        width  = 70
        height = 10
	return 
    } else {
	// command's result pattern => [width height\n]
        tmp := strings.Split(strings.Trim(string(out),"\n")," ");
        height,err  = strconv.Atoi(tmp[0])
        width,err = strconv.Atoi(tmp[1])
        return 
    }
}


func decideCursolPosition() (width,height int) {
    width,height = getTerminalSize()
    height = 5
    return 
}


type Ascii struct {
    Texts []string
    StopTimeInMillis time.Duration
}

func NewAscii(command string) *Ascii {

    ascii := &Ascii{}

    switch  command{
    case RUN: 
        ascii.Texts = m[RUN]
        ascii.StopTimeInMillis = 10 * time.Millisecond
    case PULL:
        ascii.Texts = m[PULL]
        ascii.StopTimeInMillis = 20 * time.Millisecond
    case PS:
        ascii.Texts = m[PS]
        ascii.StopTimeInMillis = 20 * time.Millisecond
    default:
        ascii.Texts = m[PS]
        ascii.StopTimeInMillis = 20 * time.Millisecond
    }

    return ascii
}



func main(){

    flag.Parse()
    ascii := NewAscii(flag.Arg(0))

    // clear a window
    fmt.Fprintln(out,"\x1b[2J")

    // print AA
    start_width,start_height := decideCursolPosition()
    for width := start_width; width > 0; width-- {
	for i := 0; i < len(ascii.Texts); i++{
            fmt.Fprintf(out,"\x1b[%d;%dH\x1b[2K",start_height+i,width)
	    fmt.Print(ascii.Texts[i])
	}
	time.Sleep(ascii.StopTimeInMillis)
    }

}
