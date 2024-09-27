package main

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type iExplorerHandler interface {
	Start(processListener chan bool) error
	Stop() error
	KillProcess() error
}

type explorerHandler struct {
	cmd            *exec.Cmd
	pathToChromium string
	pdvEndpoint    string
}

func newExplorerHandler(pathToChromium string, pdvEndpoint string) iExplorerHandler {
	b := &explorerHandler{}
	b.pathToChromium = pathToChromium
	b.pdvEndpoint = pdvEndpoint

	return b
}

func (b *explorerHandler) Start(processListener chan bool) error {
	endpoint := fmt.Sprintf("--app=%s", b.pdvEndpoint)

	b.cmd = exec.Command(b.pathToChromium, endpoint, "--kiosk")
	err := b.cmd.Run()
	if err != nil {
		log.WithError(err).Error("Error with the browser process")
		processListener <- false
	} else {
		processListener <- true
	}

	return err
}

func (b *explorerHandler) Stop() error {
	err := b.cmd.Process.Release()
	if err != nil {
		log.WithError(err).Fatal("Error shutting down chromium")
	}

	return err
}

func (b *explorerHandler) KillProcess() error {
	err := b.cmd.Process.Kill()
	if err != nil {
		log.WithError(err).Fatal("Error killing chromium process")
	}

	return err
}
