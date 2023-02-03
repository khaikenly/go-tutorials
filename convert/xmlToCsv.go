package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	excelize "github.com/xuri/excelize/v2"
)

type Services struct {
	ServiceCredentials []ServiceCredential `xml:"service-credential"`
}

type ServiceCredential struct {
	Id  string `xml:"service-credential-id,attr"`
	Url string `xml:"url"`
}

func main() {
	dat, err := os.Open("./services.xml")
	if err != nil {
		fmt.Println("ERROR !!!")
	}
	defer dat.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(dat)

	// we initialize our Users array
	var services Services
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &services)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Service Credential ID")
	f.SetCellValue("Sheet1", "B1", "URL")

	col := 1
	for _, v := range services.ServiceCredentials {
		col += 1
		str := fmt.Sprint(col)
		f.SetCellValue("Sheet1", "A"+str, v.Id)
		f.SetCellValue("Sheet1", "B"+str, v.Url)
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs("Services.xlsx"); err != nil {
		fmt.Println(err)
	}
}
