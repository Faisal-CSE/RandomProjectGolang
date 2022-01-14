package main

import (
	"RandomProjectGolang/CSV_FILES"
	dataChart "RandomProjectGolang/DataVisualization_Chart"
	jsonRead "RandomProjectGolang/JsonFileParse"
	pdfGen "RandomProjectGolang/PDF_Generator"
	"log"
	"net/http"
)

func main() {

	//TODO JSON FILE READ
	jsonRead.ReadJsonFile() // Console print
	log.Println("---------------------------------------")

	//TODO CREATE & WRITE CSV FILE
	CSV_FILES.CreateCSVFileAndWrite()
	log.Println("---------------------------------------")

	//TODO READ FROM CSV FILE
	CSV_FILES.ReadCSVFile() // Console print
	log.Println("---------------------------------------")

	//TODO GENERATE PDF FILE
	pdfErr := pdfGen.GeneratePdf("test.pdf")
	if pdfErr != nil {
		log.Println(pdfErr)
	}

	//TODO WORD CLOUD CHART
	http.HandleFunc("/word-cloud-chart", dataChart.CreateWordCloud)
	//TODO LINE CHART WITH RANDOM DATA
	http.HandleFunc("/line-chart", dataChart.LineChartRandomData)
	//TODO BAR CHART
	http.HandleFunc("/bar-chart", dataChart.CreateBarChart)
	//TODO PIE CHART
	http.HandleFunc("/pie-chart", dataChart.CreatePieChart)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
