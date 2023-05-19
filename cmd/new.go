/*
Copyright © 2023 radLad

*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Composes a new system of linear equations",
	Long: `Reads values from CSV file indicated or default if none provided, constructing and displaying the matrix read`,
	Run: func(cmd *cobra.Command, args []string) {
		var file = "testSystem.csv"

		if len(args) >= 1 && args[0] != "" {
			file = args[0]
		}

		system, err := ConvertCSV(file)
		if err != nil {
			log.Fatal("could not convert csv to matrix format"+"\n", err)
		}

		Show(system)
	},
}

var (
	BITSIZE = 64
)

func Show(f [][]float64) {
	// for r := 0; r < len(f); r++ {
	// 	fmt.Println(f[r])
	// }

	for row := 0; row < len(f); row++ {
		for column := 0; column < len(f[row]); column++ {
			fmt.Printf("%3v", f[row][column])
		}
		fmt.Print("\n")
	}
}

func ConvertCSV(f string) ([][]float64, error) {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		return nil, errors.New("cannot open source file " + f + "\n")
	}
	csvReader := csv.NewReader(file)
	linearEquations := make([][]float64, 0)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("unable to read csv" + "\n")
		}
		record = record[:len(record)-1]

		floats, err := StringsToFloats(record)
		if err != nil {
			return nil, errors.New("unable to parse csv" + "\n")
		}
		linearEquations = append(linearEquations, floats)
	}
	return linearEquations, nil
}

func StringsToFloats(s []string) ([]float64, error) {
	floats := make([]float64, len(s))
	for i, str := range s {
		float, err := strconv.ParseFloat(str, BITSIZE)
		if err != nil {
			return nil, errors.New("could not parse (" + str + ") to float")
		}
		floats[i] = float
	}

	return floats, nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
