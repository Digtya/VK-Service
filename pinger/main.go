package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type ContainerStatus struct {
	IPAddress string  `json:"ip_address"`
	PingTime  float64 `json:"ping_time"`
}

//выполняем пинг и возвращаем время пинга

func ping(ip string) float64 {
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return -1 // Ошибка пинга
	}

	// парсим время
	parts := string(output)
	start := "time="
	end := " ms"
	indexStart := len(parts) - len(end) - len(start)
	if indexStart < 0 {
		return -1
	}
	pingTime, _ := strconv.ParseFloat(parts[indexStart:len(parts)-len(end)], 64)
	return pingTime
}

// отправляем статус
func sendStatus(status ContainerStatus) {
	jsonData, _ := json.Marshal(status)
	resp, err := http.Post("http://backend:8080/add", "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != 201 {
		log.Println("Failed to send status:", err)
	}
}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	// Запускаем бесконечный цикл, который каждые 10 секунд будет проверять контейнеры
	for range ticker.C {
		// получаем список IP-адресов контейнеров (список жестко закодирован)
		containers := []string{"192.168.1.1", "192.168.1.2"}

		for _, ip := range containers {
			pingTime := ping(ip)
			sendStatus(ContainerStatus{IPAddress: ip, PingTime: pingTime})
		}
	}
}
