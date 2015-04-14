package backend

import (
	"bufio"
	"log"
	"net"
//	"regexp"
    "os"
    "os/user"
)


var active_ng_file string = "/News/groups/ng.active"
var new_ng_file    string = "/News/groups/ng.new"

// initializes everything


func init() {
    var user_home = GetHomeDir()
    active_ng_file  = user_home + active_ng_file
    new_ng_file = user_home + new_ng_file
}


// gets the active NG and sends them to the given sockets


func Trasmit_Active_NG(conn net.Conn)  (error) {
  file, err := os.Open(active_ng_file)
  if err != nil {
      log.Printf("[WTF] can't open file %s", active_ng_file )  
    return err
  }
  defer file.Close()

  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line string =  scanner.Text() 
    conn.Write([]byte(line +"\n"))
    log.Printf("[INFO] NNTP print: %s ", line)  
  }
    file.Close()
  return scanner.Err()
}

// transmits NEW newgroups (here "local") to the given socket


func Trasmit_New_NG(conn net.Conn)  (error) {
  file, err := os.Open(new_ng_file)
  if err != nil {
      log.Printf("[WTF] can't open file %s", new_ng_file )  
    return err
  }
  defer file.Close()

  
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line string =  scanner.Text() 
    conn.Write([]byte(line+"\n"))
    log.Printf("[INFO] NNTP print: %s ", line)  
  }
    file.Close()
  return scanner.Err()
}




// just gets the home directory. to be moved in "tools"

        

func GetHomeDir() (string) {
    
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
        log.Printf("[WTF] can't get homedir for user! SYSADMIIIN!"  )
    }
    return usr.HomeDir
}