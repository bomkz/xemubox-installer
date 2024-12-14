package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/tabwriter"

	"github.com/jaypipes/ghw"
)

type blkdev struct {
	name     string
	capacity uint64
	vendor   string
	model    string
}

var blkdevs []blkdev

func buildBlockdevList() {

	blockdevs, err := ghw.Block()
	if err != nil {
		log.Panic(err)
	}

	for _, blockdev := range blockdevs.Disks {
		newblkdev := blkdev{
			name:     blockdev.Name,
			capacity: blockdev.SizeBytes,
			vendor:   blockdev.Vendor,
			model:    blockdev.Model,
		}

		if newblkdev.name == "ATA" || newblkdev.name == "ata" {
			newblkdev.name = "Generic"
		}

		newblkdev.model = strings.ReplaceAll(newblkdev.model, "_", " ")

		blkdevs = append(blkdevs, newblkdev)
	}

}

func prepareBlkdevGrid() string {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "ID\t Capacity\t Vendor\t Model\t")

	for _, v := range blkdevs {
		fmt.Fprintf(w, "%s\t %d\t %s\t %s\t\n", v.name, (v.capacity / 1000000000), v.vendor, v.model)
	}

	w.Flush()

	return buf.String()
}
