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
)

var out io.Writer = os.Stdout

// return current terminal window (width,height)
func getTerminalSize() (width,height int) {
    cmd := exec.Command("stty","size")
    cmd.Stdin = os.Stdin
    out,err := cmd.Output()

    if err != nil {
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


type Command struct {
    AsciiArt []string
    StopTimeInMillis time.Duration
}

func NewCommand(inputText string) *Command {

    command := &Command{}

    switch inputText{
    case "run": 
        command.StopTimeInMillis = 10 * time.Millisecond
        command.AsciiArt = strings.Split(
`        
  ___   
 / "" \\  I'm a highway star
<  o   \     
 \     |___     
 /    _____\<      3
 |    ''''" \  ＝＝ 3
 \\_________/      3
     \ \ 
` , "\n")
    case "pull":
        command.StopTimeInMillis = 20 * time.Millisecond
        command.AsciiArt = strings.Split(
`        
         ___    
        / "" \  someone is pulling me
<<<<<<<<  o   \     
        \     |___     
        /    _____\<   
        |    ''''" ~\  
        \\~~~~~~~~~~/ 
        
` , "\n")
    case "ps":
        command.StopTimeInMillis = 20 * time.Millisecond
        command.AsciiArt = strings.Split(
`        
  ___    
 / "" \\  p.s. my holiday has gone to waste
<  o   \   
 \     |___     
 /    _____\<
 |    ''''" \ 
 \\_________/
 ～～～～～～～～～～～
` , "\n")
    default:
        command.StopTimeInMillis = 20 * time.Millisecond
        command.AsciiArt =  strings.Split(
`        
  ___    
 / "" \\  p.s. my holiday has gone to waste
<  o   \   
 \     |___     
 /    _____\<
 |    ''''" \ 
 \\_________/
 ～～～～～～～～～～～
` , "\n")
    }

    return command
}


func main(){

    flag.Parse()
    command := NewCommand(flag.Arg(0))

    // clear a window
    fmt.Fprintln(out,"\x1b[2J")

    // print AsciiArt
    start_width,start_height := decideCursolPosition()
    for width := start_width; width > 0; width-- {
	for i := 0; i < len(command.AsciiArt); i++{
            fmt.Fprintf(out,"\x1b[%d;%dH\x1b[2K",start_height+i,width)
	    fmt.Print(command.AsciiArt[i])
	}
	time.Sleep(command.StopTimeInMillis)
    }

}
