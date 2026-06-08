package model

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Name      string    `json:"name" db:"name"`
	Role      string    `json:"role" db:"role"` // admin, guide, visitor, staff
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ScenicSpot struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Code        string    `json:"code" db:"code"`
	Description string    `json:"description" db:"description"`
	Category    string    `json:"category" db:"category"` // natural, cultural, theme_park
	Address     string    `json:"address" db:"address"`
	Latitude    float64   `json:"latitude" db:"latitude"`
	Longitude   float64   `json:"longitude" db:"longitude"`
	Images      string    `json:"images" db:"images"` // JSON array
	Rating      float64   `json:"rating" db:"rating"`
	OpenTime    string    `json:"open_time" db:"open_time"`
	CloseTime   string    `json:"close_time" db:"close_time"`
	MaxCapacity int       `json:"max_capacity" db:"max_capacity"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type TicketType struct {
	ID          int64     `json:"id" db:"id"`
	SpotID      int64     `json:"spot_id" db:"spot_id"`
	Name        string    `json:"name" db:"name"` // adult, child, senior, student, vip
	Price       float64   `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	StartTime   string    `json:"start_time" db:"start_time"`
	EndTime     string    `json:"end_time" db:"end_time"`
	Quota       int       `json:"quota" db:"quota"`
	Sold        int       `json:"sold" db:"sold"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type TicketOrder struct {
	ID            int64     `json:"id" db:"id"`
	OrderNo       string    `json:"order_no" db:"order_no"`
	UserID        int64     `json:"user_id" db:"user_id"`
	SpotID        int64     `json:"spot_id" db:"spot_id"`
	TicketTypeID  int64     `json:"ticket_type_id" db:"ticket_type_id"`
	Quantity      int       `json:"quantity" db:"quantity"`
	TotalAmount   float64   `json:"total_amount" db:"total_amount"`
	VisitDate     time.Time `json:"visit_date" db:"visit_date"`
	QRCode        string    `json:"qr_code" db:"qr_code"`
	Status        string    `json:"status" db:"status"` // paid, used, expired, refunded
	PaidAt        *time.Time `json:"paid_at" db:"paid_at"`
	UsedAt        *time.Time `json:"used_at" db:"used_at"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type VisitorFlow struct {
	ID        int64     `json:"id" db:"id"`
	SpotID    int64     `json:"spot_id" db:"spot_id"`
	Date      time.Time `json:"date" db:"date"`
	Hour      int       `json:"hour" db:"hour"`
	EnterCount int      `json:"enter_count" db:"enter_count"`
	ExitCount int       `json:"exit_count" db:"exit_count"`
	CurrentCount int    `json:"current_count" db:"current_count"`
}

type TourGuide struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Phone       string    `json:"phone" db:"phone"`
	Language    string    `json:"language" db:"language"`
	Rating      float64   `json:"rating" db:"rating"`
	TourCount   int       `json:"tour_count" db:"tour_count"`
	LicenseNo   string    `json:"license_no" db:"license_no"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type GuideSchedule struct {
	ID        int64     `json:"id" db:"id"`
	GuideID   int64     `json:"guide_id" db:"guide_id"`
	SpotID    int64     `json:"spot_id" db:"spot_id"`
	Date      time.Time `json:"date" db:"date"`
	StartTime string    `json:"start_time" db:"start_time"`
	EndTime   string    `json:"end_time" db:"end_time"`
	GroupSize int       `json:"group_size" db:"group_size"`
	Status    string    `json:"status" db:"status"`
}

type Hotel struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Address     string    `json:"address" db:"address"`
	Star        int       `json:"star" db:"star"`
	Phone       string    `json:"phone" db:"phone"`
	Description string    `json:"description" db:"description"`
	Rating      float64   `json:"rating" db:"rating"`
	Images      string    `json:"images" db:"images"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Room struct {
	ID         int64     `json:"id" db:"id"`
	HotelID    int64     `json:"hotel_id" db:"hotel_id"`
	Type       string    `json:"type" db:"type"` // single, double, suite, family
	Price      float64   `json:"price" db:"price"`
	Capacity   int       `json:"capacity" db:"capacity"`
	TotalRooms int       `json:"total_rooms" db:"total_rooms"`
	Available  int       `json:"available" db:"available"`
	Status     string    `json:"status" db:"status"`
}

type RoomBooking struct {
	ID         int64      `json:"id" db:"id"`
	OrderNo    string     `json:"order_no" db:"order_no"`
	UserID     int64      `json:"user_id" db:"user_id"`
	RoomID     int64      `json:"room_id" db:"room_id"`
	CheckIn    time.Time  `json:"check_in" db:"check_in"`
	CheckOut   time.Time  `json:"check_out" db:"check_out"`
	Guests     int        `json:"guests" db:"guests"`
	TotalPrice float64    `json:"total_price" db:"total_price"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}

type Route struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Duration    string    `json:"duration" db:"duration"` // "2 hours", "half day"
	Distance    float64   `json:"distance" db:"distance"`
	SpotIDs     string    `json:"spot_ids" db:"spot_ids"` // JSON array
	Difficulty  string    `json:"difficulty" db:"difficulty"` // easy, moderate, hard
	Rating      float64   `json:"rating" db:"rating"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Event struct {
	ID          int64      `json:"id" db:"id"`
	SpotID      int64      `json:"spot_id" db:"spot_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	StartDate   time.Time  `json:"start_date" db:"start_date"`
	EndDate     time.Time  `json:"end_date" db:"end_date"`
	Location    string     `json:"location" db:"location"`
	MaxParticipants int   `json:"max_participants" db:"max_participants"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

type Feedback struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	SpotID    int64     `json:"spot_id" db:"spot_id"`
	Rating    int       `json:"rating" db:"rating"`
	Content   string    `json:"content" db:"content"`
	Images    string    `json:"images" db:"images"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Complaint struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	SpotID      *int64    `json:"spot_id" db:"spot_id"`
	Type        string    `json:"type" db:"type"`
	Content     string    `json:"content" db:"content"`
	Status      string    `json:"status" db:"status"` // pending, processing, resolved
	HandlerID   *int64    `json:"handler_id" db:"handler_id"`
	Result      string    `json:"result" db:"result"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Emergency struct {
	ID          int64     `json:"id" db:"id"`
	SpotID      int64     `json:"spot_id" db:"spot_id"`
	Type        string    `json:"type" db:"type"` // medical, fire, lost, weather
	Description string    `json:"description" db:"description"`
	Level       string    `json:"level" db:"level"` // low, medium, high, critical
	Location    string    `json:"location" db:"location"`
	Status      string    `json:"status" db:"status"` // reported, responding, resolved
	ReporterID  int64     `json:"reporter_id" db:"reporter_id"`
	HandlerID   *int64    `json:"handler_id" db:"handler_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ResolvedAt  *time.Time `json:"resolved_at" db:"resolved_at"`
}
