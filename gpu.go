package main

import (
	"bytes"
	"fmt"
	"log"
	"text/tabwriter"

	"github.com/jaypipes/ghw"
)

type gpudev struct {
	vendor string
	model  string
}

var gpudevs []gpudev

func buildGpudevList() {

	gpudevices, err := ghw.GPU(ghw.WithDisableWarnings())
	if err != nil {
		log.Panic(err)
	}

	for _, gpudevice := range gpudevices.GraphicsCards {
		newgpudev := gpudev{
			vendor: gpudevice.DeviceInfo.Vendor.Name,
			model:  gpudevice.DeviceInfo.Product.Name,
		}

		gpudevs = append(gpudevs, newgpudev)
	}

}

func prepareGpudevGrid() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Vendor\t Model\t")

	for _, v := range gpudevs {
		fmt.Fprintf(w, "%s\t %s\t\n", v.vendor, v.model)
	}

	w.Flush()

	return buf.String()
}
