package main

import (
	"bems-api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type FutureEnergyConsumptionRateRequest struct {
	Min  int       `json:"min"`
	Max  int       `json:"max"`
	Time time.Time `json:"time"`
}

type FutureEnergyConsumptionRateResponse struct {
	Decision int       `json:"min"`
	Time     time.Time `json:"time"`
}

type EnergyConsumptionRateResponse struct {
	Consumed int `json:"consumed"`
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/energy-consumption-rate", EnergyConsumptionRate)
		v1.GET("/future-energy-consumption-rate", FutureEnergyConsumptionRate)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}

// @BasePath /api/v1

// EnergyConsumptionRate godoc
// @Summary The Building exposes the metrics of electrical energy consumption (server+cooling) related to the entire Building in kW
// @Schemes
// @Description The Building exposes the metrics of electrical energy consumption (server+cooling) related to the entire Building in kW
// @Accept json
// @Produce json
// @Success 200 {object} EnergyConsumptionRateResponse
// @Router /energy-consumption-rate [get]
func EnergyConsumptionRate(g *gin.Context) {
	g.JSON(http.StatusOK, "energy-consumption-rate")
}

// FutureEnergyConsumptionRate godoc
// @Summary A possible range for the optimisation, e.g.: "10-25kW" for the Building
// @Schemes
// @Description Accepts a possible energy range for the optimisation, e.g.: "10-25kW". Returns a value in kW contained in the range representing
// the decision of how much energy the building should consume.
// @Accept json
// @Produce json
// @Param energy-consumption-rate-request body FutureEnergyConsumptionRateRequest true "Request providing a possible range of energy consumption"
// @Success 200 {object} FutureEnergyConsumptionRateResponse
// @Router /future-energy-consumption-rate [get]
func FutureEnergyConsumptionRate(g *gin.Context) {
	g.JSON(http.StatusOK, "future-energy-consumption-rate")
}
