package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/rivo/tview"
)

var Disks []string
var NvidiaDrivers = []string{"None", "Nvidia Open Source DKMS (Turing+)", "Nividia DKMS (Maxwell to Lovelace)", "Nouveau"}
var AMDDrivers = []string{"None", "Generic Driver 1", "Generic Driver 2", "Generic Driver 3"}
var IntelDrivers = []string{"None", "Generic Driver 1", "Generic Driver 2", "Generic Driver 3"}

func main() {
	app := tview.NewApplication()

	form := buildInstallForm()

	if err := app.SetRoot(form, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		log.Panic(err)
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

func buildInstallForm() (form *tview.Form) {
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
	form = tview.NewForm().
		AddTextView("Disks", blkdevgrid, blkwidth, blkheight, false, false).
		AddDropDown("Disk Select", Disks, 0, dskVal).
		AddTextView("GPUs", gpudevgrid, gpuwidth, gpuheight, false, false).
		AddDropDown("Nvidia Drivers", NvidiaDrivers, 0, nvdVal).
		AddDropDown("AMD Drivers", AMDDrivers, 0, amdVal).
		AddDropDown("Intel Drivers", IntelDrivers, 0, intVal).
		AddPasswordField("Root Account Password", "", 32, '*', rtpVal)

	form.SetBorder(true).SetTitle("XemuBOX").SetTitleAlign(tview.AlignCenter)
	return
}

func dskVal(option string, optionIndex int) {

}

func nvdVal(option string, optionIndex int) {

}

func amdVal(option string, optionIndex int) {

}

func intVal(option string, optionIndex int) {

}

func rtpVal(option string) {}
