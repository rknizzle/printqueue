package printqueue

import (
	"errors"
)

// PrinterClient handles communication to a 3D printer client device.
type PrinterClient interface {
	SendStartCommand(file string) error
	SendPauseCommand() error
	SendResumeCommand() error
	SendCancelCommand() error
}

// Printer represents functions of a 3D printer.
type Printer interface {
	Start(file string) error
	Pause() error
	Resume() error
	Cancel() error
}

// A PrinterFFF represents a fused filament fabrication 3D printer.
type PrinterFFF struct {
	ID     string
	Name   string
	Client PrinterClient
	X      int
	Y      int
	Z      int
	Status string
}

// NewPrinterFFFWithClient returns a fused filament fabrication 3D printer with a method of interacting with
// the printers client device.
func NewPrinterFFFWithClient(p PrinterClient) *PrinterFFF {
	// TODO: remove test data
	return &PrinterFFF{
		ID:     "xxx",
		Name:   "test",
		Client: p,
		Status: "Ready",
		X:      200,
		Y:      200,
		Z:      200,
	}
}

func (p *PrinterFFF) Start(file string) error {
	if p.Status == "Ready" {
		err := p.Client.SendStartCommand(file)
		if err != nil {
			return err
		}

		p.Status = "Running"
		return nil
	} else {
		return errors.New("not in a ready state")
	}
}

func (p *PrinterFFF) Pause() error {
	if p.Status == "Running" {
		err := p.Client.SendPauseCommand()
		if err != nil {
			return err
		}

		p.Status = "Paused"
		return nil
	} else {
		return errors.New("not in a running state")
	}
}

func (p *PrinterFFF) Resume() error {
	if p.Status == "Paused" {
		err := p.Client.SendResumeCommand()
		if err != nil {
			return err
		}

		p.Status = "Running"
		return nil
	} else {
		return errors.New("not in a paused state")
	}
}

func (p *PrinterFFF) Cancel() error {
	if p.Status == "Running" {
		err := p.Client.SendCancelCommand()
		if err != nil {
			return err
		}

		p.Status = "Cancelled"
		return nil
	} else {
		return errors.New("not in a running state")
	}
}
