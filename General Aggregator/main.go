package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Occupancy struct {
	Summary Summary   `json:"summary"`
	Details []Details `json:"details"`
}

type Summary struct {
	Label string  `json:"label"`
	Value float32 `json:"value"`
}

type Details struct {
	Label string          `json:"label"`
	Value OccupancyDetail `json:"value"`
}

type OccupancyDetail struct {
	Percentage float32 `json:"percentage"`
	PodsSold   int     `json:"pods_sold"`
	TotalPods  int     `json:"total_pods"`
}

func main() {

	data := generateRawData()
	// data := generateRawData2()
	agg := generalAggregator(data)

	occupancy := calculateOccupancy(agg)

	jsonData, _ := json.Marshal(occupancy)
	fmt.Printf("%s", jsonData)

}

func generateRawData() map[string]map[string]float32 {

	var data = map[string]map[string]float32{}

	for i := 1; i < 9; i++ {

		tanggal := fmt.Sprintf("2021-01-0%d", i)

		data[tanggal] = make(map[string]float32)

		data[tanggal]["percentage"] = 97.5 * float32(i)
		data[tanggal]["total_pods"] = 55 * float32(i)
		data[tanggal]["pods_sold"] = 50 * float32(i)

	}

	for i := 1; i < 9; i++ {

		tanggal := fmt.Sprintf("2021-02-0%d", i)

		data[tanggal] = make(map[string]float32)

		data[tanggal]["percentage"] = 60.5 * float32(i)
		data[tanggal]["total_pods"] = 30 * float32(i)
		data[tanggal]["pods_sold"] = 10 * float32(i)

	}

	for i := 1; i < 9; i++ {

		tanggal := fmt.Sprintf("2021-03-0%d", i)

		data[tanggal] = make(map[string]float32)

		data[tanggal]["percentage"] = 40.5 * float32(i)
		data[tanggal]["total_pods"] = 60 * float32(i)
		data[tanggal]["pods_sold"] = 50 * float32(i)

	}

	return data
}

// Total Occupancy Rate
func generateRawData2() map[int]map[string]float32 {
	var data = map[int]map[string]float32{}

	for i := 1; i < 9; i++ {

		data[i] = make(map[string]float32)

		data[i]["percentage"] = 97.5 * float32(i)
		data[i]["total_pods"] = 55 * float32(i)
		data[i]["pods_sold"] = 50 * float32(i)

	}

	return data
}

func generalAggregator(data map[string]map[string]float32) map[string]map[string][]float32 {

	var (
		aggregatedData = make(map[string]map[string][]float32)
		granularity    string
	)

	for k, v := range data {
		// for i := 0; i < len(aggregatedData); i++ {

		year, month := convertDateToYearMonth(k)
		granularity = fmt.Sprintf(`%s-%d`, month, year)

		if _, ok := aggregatedData[granularity]; !ok {
			aggregatedData[granularity] = make(map[string][]float32)
		}

		for key, val := range v {
			aggregatedData[granularity][key] = append(aggregatedData[granularity][key], val)
		}

	}

	return aggregatedData
}

func calculateOccupancy(aggregatedData map[string]map[string][]float32) Occupancy {

	var (
		finalData            Occupancy
		totalPercentage      float32
		totalAggregatedLabel = len(aggregatedData)
	)

	for label, value := range aggregatedData {
		var (
			percentage float32
			total_pods int
			pods_sold  int
		)

		for k, v := range value {
			if k == "percentage" {
				percentage = calculateAverage(v, len(v))
				totalPercentage += percentage
			} else if k == "total_pods" {
				total_pods = int(calculateSum(v))
			} else if k == "pods_sold" {
				pods_sold = int(calculateSum(v))
			}
		}

		occupancyDetail := OccupancyDetail{
			Percentage: percentage,
			TotalPods:  total_pods,
			PodsSold:   pods_sold,
		}

		detail := Details{
			Label: label,
			Value: occupancyDetail,
		}

		summary := Summary{
			Label: "Total Occupancy Data",
			Value: totalPercentage / float32(totalAggregatedLabel),
		}

		finalData.Details = append(finalData.Details, detail)
		finalData.Summary = summary
	}

	return finalData
}

func calculateAverage(arr []float32, len int) float32 {

	var sum, avg float32

	for _, v := range arr {
		sum += v
	}

	avg = sum / float32(len)

	return avg
}

func calculateSum(arr []float32) float32 {

	var sum float32

	for _, v := range arr {
		sum += v
	}

	return sum
}

func convertDateToYearMonth(strDate string) (int, string) {

	date, _ := time.Parse("2006-01-02", strDate)

	year, month, _ := date.Date()

	return year, month.String()
}
