package models

import (
	"encoding/json"
	"net"
)

type Record struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	City     string `json:"city"`
	Protocol string `json:"protocol"`
	Domain   string `json:"domain"`
}

func UnmarshalRecords(data []byte) ([]Record, error) {
	var records []Record
	err := json.Unmarshal(data, &records)
	return records, err
}

// IsValidIP - Exports the IP validation function to be accessible from other packages
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
