package main

import (
	"log"
	"net/http"

	"github.com/pbnjay/memory"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func memoriaLivre () float64 {
	memoria_livre := memory.FreeMemory()
	return float64(memoria_livre)
}

func totalMemoria() float64 {
	memoria_total := memory.TotalMemory()
	return memoria_total
}

var (
	memoriaLivreBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_bytes",
		Help: "Quantidade de memoria livre em bytes",
	})

	memoriaLivreMegasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memoria_livre_megas",
		Help: "Quantidade de memoria livre em megas",
	})

	totalMemoryBytesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_memoria_bytes",
		Help: "Total de memoria livre em bytes",
	})

	totalMemoryGigasGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_memoria_gigas",
		Help: "Total de memoria livre em Gigas",
	})
)

func init() {
	prometheus.MustRegister(memoriaLivreBytesGauge)
	prometheus.MustRegister(memoriaLivreMegasGauge)
	prometheus.MustRegister(totalMemoryBytesGauge)
	prometheus.MustRegister(totalMemoryGigasGauge)
}

func main() {
	memoriaLivreBytesGauge.Set(memoria_livre())
	memoriaLivreBytesGauge.Set(memoria_livre() / 1024/ 1024)
	totalMemoryBytesGauge.Set(totalMemoria())
	totalMemoryGigasGauge.Set(totalMemoria() / 1024/ 1024/ 1024)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":7788", nil))
}