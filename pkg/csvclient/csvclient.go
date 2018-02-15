package csvclient

import (
		"log"
		"os"
		"encoding/csv"
		"bufio"
		"strings"
)


type Client struct {
}

func (client *Client) Read(name string) (data []map[string]interface{}, err error){
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	d := make([][]string, i)
	for scanner.Scan(){
		rows := make([]string, 1)
		if err != nil {
			log.Fatal(err)
		}
		rows[0] = scanner.Text()
		i++
		result := strings.Split(rows[0],",")
		d = append(d, result)
	}
	coloums := (len(d[0]))
	rows := (len(d)-1)
	r := 1
	result := []map[string] interface{}{}
	for ; r <= rows; r++ {
		m := make(map[string]interface{} )

		for c := 0; c < coloums; c++ {
			key := d[0][c]
			value := d[r][c]
			m[key] = value
		}
		result = append(result,m)
	}
	return result, nil
}

func (client *Client) Write(data []map[string]interface{}, name string) (err error){
	// Create csv file
	file, err := os.Create(name)
	if err != nil{
		return err
	}
	defer file.Close()

	// define writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// create list of headers
	headers := make([]string, len(data[0]))
	i := 0
	for key , _ := range data[0]{
		headers[i] = key
		i ++
	}

	// write headers to csv file
	err = writer.Write(headers)
	if err != nil{
		return err
	}

	// Create as list from each map
	value1 := make([]string, len(headers))
	for _ , row := range data{
		for i = 0; i < len(headers);i++{
			value1[i] = row[headers[i]].(string)
		}

		// write list to csv file and loop to next
		err = writer.Write(value1)
		if err != nil{
			return err
		}
	}
	log.Printf("csv file written")
	return nil
}