package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/CodingFervor/smart-tourism-management/internal/database"
)

func main() {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().Format(time.RFC3339)})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/auth/login", Login)
		auth := api.Group("/")
		auth.Use(AuthMiddleware())
		{
			// Scenic spots
			auth.GET("/spots", ListSpots)
			auth.POST("/spots", CreateSpot)
			auth.GET("/spots/:id", GetSpot)
			auth.PUT("/spots/:id", UpdateSpot)

			// Tickets
			auth.GET("/tickets/types", ListTicketTypes)
			auth.POST("/tickets/types", CreateTicketType)
			auth.POST("/tickets/buy", BuyTicket)
			auth.GET("/tickets/orders", ListTicketOrders)
			auth.GET("/tickets/validate/:code", ValidateTicket)

			// Visitor flow
			auth.GET("/flow/realtime", RealtimeFlow)
			auth.GET("/flow/history", FlowHistory)
			auth.GET("/flow/forecast", FlowForecast)
			auth.PUT("/flow/limit", SetFlowLimit)

			// Tour guides
			auth.GET("/guides", ListGuides)
			auth.POST("/guides", RegisterGuide)
			auth.GET("/guides/schedule", GuideSchedule)
			auth.POST("/guides/assign", AssignGuide)

			// Hotels
			auth.GET("/hotels", ListHotels)
			auth.POST("/hotels", CreateHotel)
			auth.GET("/hotels/:id/rooms", ListRooms)
			auth.POST("/hotels/book", BookRoom)

			// Routes
			auth.GET("/routes", ListRoutes)
			auth.POST("/routes", CreateRoute)
			auth.GET("/routes/recommend", RecommendRoutes)

			// Events
			auth.GET("/events", ListEvents)
			auth.POST("/events", CreateEvent)
			auth.PUT("/events/:id", UpdateEvent)

			// Feedback
			auth.GET("/feedbacks", ListFeedbacks)
			auth.POST("/feedbacks", CreateFeedback)
			auth.GET("/complaints", ListComplaints)
			auth.POST("/complaints", CreateComplaint)
			auth.PUT("/complaints/:id", HandleComplaint)

			// Emergency
			auth.GET("/emergencies", ListEmergencies)
			auth.POST("/emergencies", ReportEmergency)
			auth.PUT("/emergencies/:id", HandleEmergency)

			// Reports
			auth.GET("/reports/revenue", RevenueReport)
			auth.GET("/reports/visitors", VisitorReport)
			auth.GET("/reports/satisfaction", SatisfactionReport)
			auth.GET("/dashboard", DashboardOverview)
		}
	}
	log.Println("Smart Tourism Management starting on :8080")
	addr := ":" + strconv.Itoa(8080)
	srv := &http.Server{Addr: addr, Handler: r}
	go func() {
		logger.Info("server listening", "port", 8080)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("forced shutdown", "error", err)
	}
	logger.Info("server exited")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" { c.AbortWithStatus(http.StatusNoContent); return }
		c.Next()
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" { c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"}); return }
		c.Next()
	}
}

func Login(c *gin.Context)              { c.JSON(http.StatusOK, gin.H{"message": "login"}) }
func ListSpots(c *gin.Context)          { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateSpot(c *gin.Context)         { c.JSON(http.StatusCreated, gin.H{"message": "spot created"}) }
func GetSpot(c *gin.Context)            { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func UpdateSpot(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"message": "spot updated"}) }
func ListTicketTypes(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateTicketType(c *gin.Context)   { c.JSON(http.StatusCreated, gin.H{"message": "ticket type created"}) }
func BuyTicket(c *gin.Context)          { c.JSON(http.StatusCreated, gin.H{"message": "ticket bought"}) }
func ListTicketOrders(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func ValidateTicket(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": gin.H{"valid": true}}) }
func RealtimeFlow(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func FlowHistory(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func FlowForecast(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func SetFlowLimit(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"message": "flow limit set"}) }
func ListGuides(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func RegisterGuide(c *gin.Context)      { c.JSON(http.StatusCreated, gin.H{"message": "guide registered"}) }
func GuideSchedule(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func AssignGuide(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"message": "guide assigned"}) }
func ListHotels(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateHotel(c *gin.Context)        { c.JSON(http.StatusCreated, gin.H{"message": "hotel created"}) }
func ListRooms(c *gin.Context)          { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func BookRoom(c *gin.Context)           { c.JSON(http.StatusCreated, gin.H{"message": "room booked"}) }
func ListRoutes(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateRoute(c *gin.Context)        { c.JSON(http.StatusCreated, gin.H{"message": "route created"}) }
func RecommendRoutes(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func ListEvents(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateEvent(c *gin.Context)        { c.JSON(http.StatusCreated, gin.H{"message": "event created"}) }
func UpdateEvent(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"message": "event updated"}) }
func ListFeedbacks(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateFeedback(c *gin.Context)     { c.JSON(http.StatusCreated, gin.H{"message": "feedback created"}) }
func ListComplaints(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func CreateComplaint(c *gin.Context)    { c.JSON(http.StatusCreated, gin.H{"message": "complaint created"}) }
func HandleComplaint(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "complaint handled"}) }
func ListEmergencies(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func ReportEmergency(c *gin.Context)    { c.JSON(http.StatusCreated, gin.H{"message": "emergency reported"}) }
func HandleEmergency(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"message": "emergency handled"}) }
func RevenueReport(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func VisitorReport(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func SatisfactionReport(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func DashboardOverview(c *gin.Context)  { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
