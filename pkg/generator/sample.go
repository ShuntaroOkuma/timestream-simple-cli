package generator

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type HomeData struct {
	DeviceID    string
	Location    string
	Timestamp   time.Time
	Temperature string // degree Celsius
	Humidity    string
	Pressure    string // hPa
}

func GenerateHomeData(deviceID string, location string, startTime time.Time, count int) []HomeData {
	data := make([]HomeData, count)
	for i := 0; i < count; i++ {
		data[i] = HomeData{
			DeviceID:    deviceID,
			Location:    location,
			Timestamp:   startTime.Add(time.Duration(i) * time.Minute),
			Temperature: fmt.Sprintf("%.2f", 15+rand.Float64()*10),   // 15 - 25 degree Celsius
			Humidity:    fmt.Sprintf("%.2f", 30+rand.Float64()*20),   // 30 - 50 %
			Pressure:    fmt.Sprintf("%.2f", 1000+rand.Float64()*50), // 1000 - 1050 hPa
		}
	}
	return data
}

type SmartBuildingData struct {
	BuildingID       string
	FloorNumber      string
	Timestamp        time.Time
	PowerConsumption string // kWh
	CO2Level         string // ppm(co2 concentration)
	LightLevel       string // lux
}

func GenerateSmartBuildingData(buildingID string, floorNumber string, startTime time.Time, count int) []SmartBuildingData {
	data := make([]SmartBuildingData, count)
	for i := 0; i < count; i++ {
		data[i] = SmartBuildingData{
			BuildingID:       buildingID,
			FloorNumber:      floorNumber,
			Timestamp:        startTime.Add(time.Duration(i) * time.Minute),
			PowerConsumption: fmt.Sprintf("%.2f", rand.Float64()*50),      // 0 - 50 kWh
			CO2Level:         fmt.Sprintf("%.2f", 400+rand.Float64()*200), // 400 - 600 ppm(co2 concentration)
			LightLevel:       fmt.Sprintf("%.2f", 100+rand.Float64()*300), // 100 - 400 lux
		}
	}
	return data
}
