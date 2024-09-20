package main

import (
	"flag"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/jszwec/csvutil"
)

var fileFlag = flag.String("file", "tests.csv", "file with your data")

type testData struct {
	Data     DataTypeStruct
	ToAssert bool
}

func parseCSV(filename string, dataVar any) error {
	// Open csv file
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error", err)
	}
	// Parse csv data
	if err := csvutil.Unmarshal(file, dataVar); err != nil {
		log.Fatal("Error", err)
	}
	return nil
}

func setupLogger() *log.Logger {
	// Override the default log level styles
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().SetString("TEST FAIL").Padding(0, 2, 0, 2).Foreground(lipgloss.Color("204")).Bold(true)
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().SetString("TEST SUCCEED").Padding(0, 0, 0, 1).Foreground(lipgloss.Color("86")).Bold(true)
	styles.Keys["isEnoughChange"] = lipgloss.NewStyle().Foreground(lipgloss.Color("134"))
	styles.Keys["assertIsEnoughChange"] = lipgloss.NewStyle().Foreground(lipgloss.Color("134"))
	styles.Values["isEnoughChange"] = lipgloss.NewStyle().Bold(true)
	styles.Values["assertIsEnoughChange"] = lipgloss.NewStyle().Bold(true)
	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	return logger
}

func SetupTestData() ([]testData, *log.Logger) {
	flag.Parse()

	logger := setupLogger()

	// Get the CSV file's data mapped according to the data struct
	var csvData []DataTypeStruct
	if err := parseCSV(*fileFlag, &csvData); err != nil {
		logger.Fatal("Error", err)
	}

	mapData := make([]testData, len(csvData))

	for i, v := range csvData {
		mapData[i].Data = v
		mapData[i].ToAssert = v.AssertIs
	}

	return mapData, logger
}
