package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Location struct {
	X, Y float64
}

// Scan implements the sql.Scanner interface
//func (loc *Location) Scan(v interface{}) error {
// Scan a value into struct from database driver
//}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", loc.X, loc.Y)},
	}
}

type BillDetail struct {
	BillId              uint64
	RestaurantProductID string
	amount              uint
}

type BillTax struct {
	BillId uint
	TaxId  uint64
	Value  float64
}

type PayBill struct {
	BillId        uint64    `gorm:"type:integer"`
	PayBillAmount float64   `gorm:"type:double precision"`
	PayBillDate   time.Time `gorm:"type:timestamptz"`
}

type Reservation struct {
	ReservationId             uint64     `gorm:"primaryKey;autoIncrement:true"`
	ReservationTable          string     `gorm:"type:varchar(100)"`
	ReservationNumberOfPeople uint       `gorm:"type:integer"`
	ReservationDate           *time.Time `gorm:"type:timestamptz"`
	ReservationStatus         string     `gorm:"type:varchar(9)"`
	ReservationToPickUp       bool       `gorm:"type:boolean"`
	RestaurantId              uint       `gorm:"foreignKey:Restaurant"`
	ClientId                  uint64     `gorm:"foreignKey:Client"`
}

type RestaurantSchedule struct {
	RestaurantId uint64 `gorm:"type:integer"`
	ScheduleId   uint64 `gorm:"type:integer"`
	Status       bool   `gorm:"type:boolean"`
}
type Schedule struct {
	ScheduleDay         string       `gorm:"type:varchar(15)"`
	ScheduleId          uint64       `gorm:"primaryKey;autoIncrement:true"`
	ScheduleOpeningHour string       `gorm:"type:varchar(15)"`
	ScheduleClosingHour string       `gorm:"type:varchar(15)"`
	Restaurants         []Restaurant `gorm:"many2many:restaurant_schedule"`
}

type Tax struct {
	TaxID     uint64  `gorm:"primaryKey;autoIncrement:true"`
	TaxValue  float64 `gorm:"type:double precision"`
	TaxYear   uint    `gorm:"type:integer"`
	TaxStatus bool    `gorm:"type:boolean"`
	Bills     []Bill  `gorm:"many2many:bill_tax"`
}
type DocumentType struct {
	DocumentTypeID     string           `gorm:"type:varchar(100);primaryKey"`
	DocumentTypeStatus bool             `gorm:"type:boolean"`
	Clients            []Client         `gorm:"foreignKey:ClientID"`
	RestaurantUsers    []RestaurantUser `gorm:"foreignKey:RestaurantUserID"`
}
type RestaurantUser struct {
	RestaurantUserID           uint64     `gorm:"primaryKey;column:restaurant_user_id;integer;autoIncrement:true;"`
	RestaurantUserName         string     `gorm:"column:restaurant_user_name;type:varchar(100);"`
	RestaurantUserLastname     string     `gorm:"column:restaurant_user_lastname;type:varchar(100);"`
	RestaurantUserLatitude     float64    `gorm:"type:double precision;"`
	RestaurantUserLongitude    float64    `gorm:"type:double precision;"`
	RestaurantUserBirthdate    *time.Time `gorm:"column:restaurant_user_birthdate;type:timestamptz;"`
	RestaurantUserPass         string     `gorm:"column:restaurant_user_pass;type:text;"`
	RestaurantUserCreationDate time.Time  `gorm:"column:restaurant_user_creation_date;type:timestamptz;"`
	RestaurantUserModdate      time.Time  `gorm:"column:restaurant_user_moddate;type:timestamptz;"`
	RestaurantUserPicture      string     `gorm:"column:restaurant_user_picture;type:text;"`
	RestaurantUserStatus       bool       `gorm:"column:restaurant_user_status;type:boolean;"`
	DocumentTypeID             string     `gorm:"column:document_type_id;foreignKey:DocumentTypeID;type:varchar(100);"`
	Profiles                   []Profile  `gorm:"many2many:user_profile"`
}

type Profile struct {
	ProfileID           uint64           `gorm:"primaryKey;autoIncrement:true;column:profile_id;"`
	ProfileName         string           `gorm:"column:profile_name;type:varchar(100)"`
	ProfileStatus       bool             `gorm:"column:profile_status;type:boolean"`
	ProfileCreationDate time.Time        `gorm:"column:profile_creation_date;type:timestamptz"`
	RestaurantUsers     []RestaurantUser `gorm:"many2many:user_profile"`
}

type UserProfile struct {
	RestaurantUserID        uint64    `gorm:"type:integer;"`
	ProfileID               uint64    `gorm:"type:integer;"`
	UserProfileStatus       bool      `gorm:"column:user_profile_status;type:boolean"`
	UserProfileCreationDate time.Time `gorm:"column:user_profile_creation_date;type:timestamptz"`
}

type Client struct {
	ClientID           uint64        `gorm:"column:client_id;primaryKey;type:bigint;"`
	ClientName         string        `gorm:"column:client_name;type:varchar(100)"`
	ClientLastName     string        `gorm:"column:client_lastname;type:varchar(100)"`
	ClientAddress      Location      `gorm:"column:client_address;type:point;"`
	ClientPhone        string        `gorm:"column:client_phone;type:varchar(15);"`
	ClientBirthDate    *time.Time    `gorm:"column:client_birthdate;type:timestamptz;"`
	ClientEmail        string        `gorm:"column:client_email;type:varchar(200);";`
	ClientPass         string        `gorm:"column:client_pass;type:text;";`
	ClientCreationDate time.Time     `gorm:"column:client_creation_date;type:timestamptz;";`
	ClientStatus       bool          `gorm:"column:client_status;type:boolean;";`
	DocumentTypeID     string        `gorm:"column:document_type_id;type:varchar(100);""`
	DocumentType       DocumentType  `gorm:"foreignKey:DocumentTypeID"`
	Banks              []Bank        `gorm:"many2many:pay_method"`
	Reservations       []Reservation `gorm:"foreignKey:ReservationId"`
}
type Bank struct {
	BankID     string   `gorm:"column:bank_id;primaryKey;type:varchar(100);"`
	BankStatus bool     `gorm:"column:bank_status;type:boolean;"`
	Clients    []Client `gorm:"many2many:pay_method"`
}
type PayMethod struct {
	CardNumber            string    `gorm:"column:card_number;primaryKey;type:varchar(16);"`
	BankID                string    `gorm:"column:bank_id;primaryKey;type:varchar(100);"`
	PayMethodCreationDate time.Time `gorm:"column:pay_method_creation_date;type:timestamptz;"`
	clientID              uint64    `gorm:"column:client_id;type:bigint;"`
}
type Restaurant struct {
	RestaurantID           uint64     `gorm:"column:restaurant_id;primaryKey;autoIncrement:true"`
	RestaurantName         string     `gorm:"column:restaurant_name;type:varchar(150)"`
	RestaurantAddress      Location   `gorm:"column:restaurant_address;type:point"`
	RestaurantCapacity     uint       `gorm:"column:restaurant_capacity;type:integer"`
	RestaurantStatus       bool       `gorm:"column:restaurant_status;type:boolean"`
	RestaurantCreationDate time.Time  `gorm:"column:restaurant_creation_date;type:timestamptz"`
	Products               []Product  `gorm:"many2many:restaurant_product"`
	Schedule               []Schedule `gorm:"many2many:restaurant_schedule"`
}
type Product struct {
	ProductID           string       `gorm:"column:product_id;type:varchar(150);primaryKey"`
	ProductDescription  string       `gorm:"column:product_description;type:text;"`
	ProductPrice        float64      `gorm:"column:product_price;type:float;"`
	ProductStatus       bool         `gorm:"column:product_status;type:boolean;"`
	ProductCreationDate time.Time    `gorm:"column:product_creation_date;type:timestamptz;"`
	Restaurants         []Restaurant `gorm:"many2many:restaurant_product"`
}
type RestaurantProduct struct {
	RestaurantProductID     uint64     `gorm:"column:restaurant_product_id;primaryKey;autoIncrement:true"`
	RestaurantID            uint64     `gorm:"column:restaurant_id;type:integer;"`
	ProductID               string     `gorm:"column:product_id;type:varchar(150);"`
	RestaurantProductAmount uint       `gorm:"column:restaurant_product_amount;type:integer;"`
	RestaurantProductDate   time.Time  `gorm:"column:restaurant_product_date;type:timestamptz;"`
	RestaurantProductState  bool       `gorm:"column:restaurant_product_state;type:boolean;"`
	Discounts               []Discount `gorm:"many2many:product_discount"`
	Bills                   []Bill     `gorm:"many2many:bill_detail"`
}

type Discount struct {
	DiscountID           uint64    `gorm:"column:discount_id;primaryKey;autoIncrement:true"`
	DiscountName         string    `gorm:"column:discount_name;type:varchar(150)"`
	DiscountDescription  string    `gorm:"column:discount_description;type:text"`
	DiscountPercentage   float64   `gorm:"column:discount_percentage;type:float"`
	DiscountStatus       bool      `gorm:"column:discount_status;type:boolean"`
	DiscountCreationDate time.Time `gorm:"column:discount_creation_date;type:timestamptz"`
	Products             []Product `gorm:"many2many:product_discount"`
}

type ProductDiscount struct {
	DiscountID            uint64    `gorm:"column:discount_id;type:bigint"`
	RestaurantProductID   uint64    `gorm:"column:restaurant_product_id;type:bigint"`
	ProductDiscountStatus bool      `gorm:"column:product_discount_status;type:boolean"`
	ProductDiscountDate   time.Time `gorm:"column:product_discount_date;type:timestamptz"`
}
type Bill struct {
	BillID             uint64              `gorm:"column:bill_id;primaryKey;autoIncrement:true"`
	BillBeforeTax      float64             `gorm:"column:bill_before_tax;type:double precision"`
	BillTotal          float64             `gorm:"column:bill_total;type:double precision"`
	BillStatus         string              `gorm:"column:bill_status;type:varchar(9)"`
	BillCreationDate   time.Time           `gorm:"column:bill_creation_date;type:timestamptz"`
	ReservationID      uint64              `gorm:"column:reservation_id;type:bigint;"`
	Reservation        Reservation         `gorm:"foreignKey:ReservationID"`
	RestaurantProducts []RestaurantProduct `gorm:"many2many:bill_detail"`
	Taxes              []Tax               `gorm:"many2many:bill_tax"`
	PayBills           []PayBill           `gorm:"foreignKey:BillId"`
}
