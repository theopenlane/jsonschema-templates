package invoice

import "time"

// Config represents the configuration for the invoice
type Config struct {
	Customer        Customer        `json:"customer"`
	Products        []Product       `json:"products"`
	PaymentTerm     PaymentTerm     `json:"paymentTerm"`
	Organization    Organization    `json:"organization"`
	Location        Location        `json:"location"`
	Address         CustomerAddress `json:"address"`
	BusinessPartner BusinessPartner `json:"businessPartner"`
}

// Seller includes the seller details for invoicing
type Seller struct {
	// Name of the customer
	Name string `json:"name"`
	// FirstName of the customer
	FirstName string `json:"firstName"`
	// LastName of the customer
	LastName string `json:"lastName"`
	// Email contact of the customer
	Email string `json:"email"`
}

// Customer includes the customer details for invoicing
type Customer struct {
	// Name of the customer
	Name string `json:"name"`
	// FirstName of the customer
	FirstName string `json:"firstName"`
	// LastName of the customer
	LastName string `json:"lastName"`
	// Email contact of the customer
	Email string `json:"email"`
	// CustomerID is the unique identifier for the customer
	CustomerID int64 `json:"customerId"`
}

// Product the invoice is for
type Product struct {
	// PartNumber is the unique identifier for the product
	PartNumber string `json:"partNumber"`
	// Name of the product
	Name string `json:"name"`
	// Quantity of the product
	Quantity int64 `json:"quantity"`
	// UnitPrice is the price of the product
	UnitPrice float64 `json:"unitPrice"`
	// UnitOfMeasure is the unit of measure for the product
	UnitOfMeasure string `json:"unitOfMeasure"`
	// BillingCategory  service billing category
	BillingCategory string `json:"billingCategory"`
	// Product category
	ProductCategory string `json:"productCategory"`
}

// PaymentTerm includes the payment term details
type PaymentTerm struct {
	// Name of the payment term
	Name string `json:"name"`
	// Value of the payment term
	Value string `json:"value"`
	// Description of the payment term
	Description string `json:"description"`
	// IsActive is true if the payment term is active
	IsActive bool `json:"isActive"`
	// TimeCreated is the time the payment term was created
	TimeCreated time.Time `json:"timeCreated"`
	// CreatedBy is the user that created the payment term
	CreatedBy string `json:"createdBy"`
	// TimeUpdated is the time the payment term was updated
	TimeUpdated time.Time `json:"timeUpdated"`
	// UpdatedBy is the user that updated the payment term
	UpdatedBy string `json:"updatedBy"`
}

// Organization details of the organization
type Organization struct {
	// Name of the organization
	Name string `json:"name"`
	// ID of the organization
	ID string `json:"id"`
}

// Location details for the invoice
type Location struct {
	// AddAddress1 is the first line of the address
	Address1 string `json:"address1"`
	// Address2 is the second line of the address
	Address2 string `json:"address2"`
	// PostalCode is the postal code of the address
	PostalCode string `json:"postalCode"`
	// City is the city of the address
	City string `json:"city"`
	// Country is the country of the address
	Country string `json:"country"`
	// Region is the region of the address
	Region string `json:"region"`
}

// Currency details
type Currency struct {
	// IsoCode is the currency code
	IsoCode string `json:"isoCode"`
	// Name of the currency
	Name string `json:"name"`
	// StdPrecision is the standard precision of the currency
	StdPrecision int64 `json:"stdPrecision"`
}

// BusinessPartner is the business partner (customer) details
type BusinessPartner struct {
	// Name also called customer name
	Name string `json:"name"`
	// AccountNumber of the business partner
	AccountNumber string `json:"accountNumber"`
}

// CustomerAddress is the address details for the customer
type CustomerAddress struct {
	Location Location `json:"location"`
	// Name  of the customer
	Name string `json:"name"`
	// Phone number of the customer
	Phone string `json:"phone"`
	// IsBillTo is used to identify as the customer's billing address.
	IsBillTo bool `json:"isBillTo"`
	// IsShipTo is used to identify as the customer's shipping address.
	IsShipTo bool `json:"isShipTo"`
}
