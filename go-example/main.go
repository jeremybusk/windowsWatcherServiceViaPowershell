package main

import (
	"bytes"
	"io"
	// "fmt"
	log "github.com/sirupsen/logrus"
	// "os"
	"os/exec"
	"time"
)

func pipeIt1() {
	socket := "0.0.0.0:22"
	first := exec.Command("ss", "-lntp")
	second := exec.Command("grep", socket)
	reader, writer := io.Pipe()
	first.Stdout = writer
	second.Stdin = reader
	var buff bytes.Buffer
	second.Stdout = &buff

	first.Start()
	second.Start()
	first.Wait()
	writer.Close()
	second.Wait()

	out := buff.String()
	if len(out) == 0 {
		log.Warn("Listener on socket %v appears to not be running.", socket)
		cmd := exec.Command("sudo", "systemctl", "restart", "sshd")
		stdout, err := cmd.Output()
		// log.Info(stdout)
		if err != nil {
			log.Error(err.Error())
			return
		} else {
			log.Info("Restarting %v socket service", socket)
			log.Info(stdout)
		}
	}

}

func main() {
	log.Info("Starting socket monitor")
	for {
		pipeIt1()
		time.Sleep(10 * time.Second)
	}
}
