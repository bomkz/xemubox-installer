package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rivo/tview"
)

var Disks []string
var GraphicsDrivers = []string{"All open-source", "AMD / ATI (open-source)", "Nividia (open kernel module for newer GPUs, Turing+)", "Nvidia (open-source nouveau driver)", "Nvidia (proprietary)"}

func main() {
	app = tview.NewApplication()

	root = tview.NewPages()

	root.SetBorder(false).SetTitle("XemuBOX Installer Utility")

	buildInstallForm()

	runApp()

}


func runApp() {
	go keepAlive()

	if err := app.SetRoot(root, true).Run(); err != nil {
		log.Panic(err)
	}

}

func keepAlive() {
	for {
		time.Sleep(50)
		app.Draw()
	}
}
func calculateGridSize(grid string) (height int, width int, err error) {

	lines := strings.Split(grid, "\n")

	if len(lines) == 0 {
		err = fmt.Errorf("Empty String Value provided.")
		return
	}
	width = len(lines[0])

	height = len(lines) + 1

	return
}

var (
	rootchanged bool
	rootfs      bool
	root        *tview.Pages
	app         *tview.Application
)

func buildInstallForm() {
	buildBlockdevList()
	blkdevgrid := prepareBlkdevGrid()
	blkheight, blkwidth, err := calculateGridSize(blkdevgrid)
	if err != nil {
		log.Panic(err)
	}

	buildGpudevList()
	gpudevgrid := prepareGpudevGrid()
	gpuheight, gpuwidth, err := calculateGridSize(gpudevgrid)
	if err != nil {
		log.Panic(err)
	}
	Disks = append(Disks, "None")
	for _, x := range blkdevs {
		Disks = append(Disks, x.name)
	}

	form := tview.NewForm()
	form.AddTextView("Disks", blkdevgrid, blkwidth, blkheight, false, false).
		AddTextView("GPUs", gpudevgrid, gpuwidth, gpuheight, false, false).
		AddDropDown("Disk Select", Disks, 0, dskVal).
		AddDropDown("Graphics Drivers", GraphicsDrivers, 0, gfxVal).
		AddPasswordField("Root Account Password", "", 0, '*', rtpVal).
		AddPasswordField("Repeat Root Account Password", "", 0, '*', rtp2Val).
		AddPasswordField("User Account Password", "", 0, '*', usrVal).
		AddPasswordField("Repeat User Account Password", "", 0, '*', usr2Val).
		AddButton("Install", func() {
			if err := checkinstall(); err != nil {
				modal := tview.NewModal()
				modal.SetText(err.Error()).SetBorder(false)
				modal.AddButtons([]string{"OK"})
				modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "OK" {
						root.SwitchToPage("installform")
					}
				})
				root.AddPage("errormodal", modal, true, true)
				return
			}

			if err := populateConfigJson(); err != nil {
				modal := tview.NewModal()
				modal.SetText(err.Error()).SetBorder(false)
				modal.AddButtons([]string{"OK"})
				modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "OK" {
						root.SwitchToPage("installform")
					}
				})
				root.AddPage("errormodal", modal, true, true)
				return
			}

			modal := tview.NewModal()
			modal.SetText("XemuBOX Configured successfully, beginning install...").SetBorder(false)
			root.AddPage("successmodal", modal, true, true)
			go quitSuccess()
		}).
		AddButton("Quit", cancel)

	form.SetBorder(false)
	root.AddAndSwitchToPage("installform", form, true)

}
func cancel() {
	app.Stop()
}

func quitSuccess() {
	time.Sleep(5 * time.Second)
	app.Stop()
}

func checkinstall() error {

	if usr != usr2 {
		return fmt.Errorf("User Password Mismatch!")
	}

	if rtp != rtp2 {
		return fmt.Errorf("Root Password Mismatch!")
	}

	if rtp == "" {
		return fmt.Errorf("Root Password Empty!")
	}

	if usr == "" {
		return fmt.Errorf("User Password Empty!")
	}

	if dsk == "" {
		return fmt.Errorf("Disk Option Empty!")
	}

	if dsk == "Choose Disk" {
		return fmt.Errorf("Disk Option Empty!")
	}

	if gfx == "" {
		return fmt.Errorf("Graphics Driver Option Empty!")
	}

	if gfx == "Choose Driver" {
		return fmt.Errorf("Graphics Driver Option Empty!")
	}
	return nil
}

func dskVal(option string, optionIndex int) {
	dsk = option
}

func gfxVal(option string, optionIndex int) {
	gfx = option
}

func rtpVal(option string) {
	rtp = option
}

func rtp2Val(option string) {
	rtp2 = option
}

var rtp, rtp2, usr, usr2, dsk, gfx string

func usrVal(option string) {
	usr = option
}

func usr2Val(option string) {
	usr2 = option
}
