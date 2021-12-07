package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "os/exec"
    "io/ioutil"
    "encoding/json"
)
const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}

type Host struct {
	Hostname string `json: "hostname"`
}


func main() {
    var file, _ = ioutil.ReadFile("hostFile.json")
    data := Host{}
    _ = json.Unmarshal([]byte(file), &data)
    hostname := data.Hostname
    fmt.Println(hostname)

    _, checkProcessRunning, _  := Shellout("pgrep pyxbackup > /dev/null 2>&1 && echo 1 || echo 0")
    _, checkProcessOK, _       := Shellout("/opt/pyxbackup/pyxbackup status -o nagios | grep 'OK' > /dev/null 2>&1 && echo 1 || echo 0")
    _, checkProcessDONE, _     := Shellout("/opt/pyxbackup/pyxbackup status -o nagios | grep 'Last backup' > /dev/null 2>&1 && echo 1 || echo 0")


    fmt.Println(checkProcessRunning)
    fmt.Println(checkProcessOK)
    fmt.Println(checkProcessDONE)

    var backupStatus string = "0"
    if string(checkProcessRunning) == "1" {
        backupStatus = "1"
    } else if string (checkProcessOK) == "1" || string (checkProcessDONE) == "1" {
        backupStatus = "0"
    } else {
        backupStatus = "-1"
    }
    fmt.Println(backupStatus)


    RESULT := fmt.Sprintf("xtrabackup_status %v\n", backupStatus)
    BASE_DIR := fmt.Sprintf("/backups/%v/stor", hostname)
    fmt.Println(BASE_DIR)
    

    lastFullBackupName := ""
    _, lastFullBackupName, _ = Shellout(fmt.Sprintf("cd %v/full && ls -1 | tail -1", BASE_DIR))

    lastIncrBackupName := ""
    if os.Chdir(fmt.Sprintf("%s/incr/%s", BASE_DIR, lastFullBackupName)) == nil {
        _, lastIncrBackupName, _ = Shellout("ls -1 | tail -1")
    }

    sizeFull := ""
    _, sizeFull, _ = Shellout(fmt.Sprintf("du -s %s/full/%s | awk '{print $1}'", BASE_DIR, lastFullBackupName))
    RESULT += fmt.Sprintf("xtrabackup_size{type=\"full\",name=\"%s\"} %s\n", lastFullBackupName, sizeFull)

    sizeIncr := ""
    _, sizeIncr, _ = Shellout(fmt.Sprintf("du -s %s/incr/%s/%s | awk '{print $1}'", BASE_DIR, lastFullBackupName, lastIncrBackupName))
    RESULT += fmt.Sprintf("xtrabackup_size{type=\"incr\",name=\"%s\"} %s\n", lastIncrBackupName, sizeIncr)

    // Save data in file
    err1 := os.WriteFile("./xtrabackup.prom", []byte(RESULT), 0666) 
    if err1 != nil {
		  log.Fatal(err1)
    }
    
    fmt.Println(RESULT)
}
