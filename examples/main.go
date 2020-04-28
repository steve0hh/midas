package main

import (
	"encoding/csv"
	"os"
	"log"
	"bufio"
	"io"
	"strconv"
	"github.com/steve0hh/midas"
	"fmt"
	"flag"
)

func main() {
	norelations := flag.Bool("norelations", false, "Run MIDAS instead of MIDAS-R")
	flag.Parse()

	f, err:= os.Open("./darpa_midas.csv")
	if err != nil {
		log.Fatal(err)
	}

	buff := bufio.NewReader(f)
	r := csv.NewReader(buff)

	src := []int{}
	dst := []int{}
	times := []int{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		s, err := strconv.ParseInt(record[0], 10, 0)
		if err != nil {
			log.Fatal(err)
		}

		d, err := strconv.ParseInt(record[1],10, 0)
		if err != nil {
			log.Fatal(err)
		}

		t, err := strconv.ParseInt(record[2], 10,0)
		if err != nil {
			log.Fatal(err)
		}

		src = append(src, int(s))
		dst = append(dst, int(d))
		times = append(times, int(t))
	}

	var anormScore []float64
	if *norelations {
		anormScore = midas.Midas(src, dst, times, 2, 769)
	}else{
		anormScore = midas.MidasR(src, dst, times, 2, 769, 0.6)
	}

	for _, v := range anormScore {
		fmt.Println(v)
	}
}
