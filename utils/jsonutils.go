package utils

import (
	"encoding/json"
	"fmt"
	"fofa-derp/models"
	"io/ioutil"
	"strconv"
)

func ProcessRecords(records []models.Record) []byte {
	regions := make(map[string]interface{})
	regionID := 901

	for _, record := range records {
		if models.IsValidIP(record.IP) && record.Domain == "" && record.Port != "443" { // Updated to use the exported function
			port, err := strconv.Atoi(record.Port)
			if err != nil {
				continue
			}

			regionCode := fmt.Sprintf("other%d", regionID-900)
			regionName := record.City

			node := map[string]interface{}{
				"Name":             fmt.Sprintf("%d-%s", regionID, record.IP),
				"RegionID":         regionID,
				"DERPPort":         port,
				"IPv4":             record.IP,
				"InsecureForTests": true,
			}

			region := map[string]interface{}{
				"RegionID":   regionID,
				"RegionCode": regionCode,
				"RegionName": regionName,
				"Nodes":      []interface{}{node},
			}

			regions[fmt.Sprintf("%d", regionID)] = region
			regionID++
		}
	}

	result := map[string]interface{}{
		"Regions": regions,
	}
	outputData, _ := json.MarshalIndent(result, "", "    ")
	return outputData
}

func WriteJSONToFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}
