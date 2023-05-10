// Author: Amar Singh
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Vehicle struct {
	Year  int
	Make  string
	Model string
	Color string
}

type Customer struct {
	ID                        int
	Name                      string
	PhoneNumber               string
	Vehicle                   *Vehicle
	ServiceHistory            []Service
	CouponNotificationHistory []*CouponNotification
}

var Customers = []*Customer{}

func NewCustomer(id int, name string, phoneNumber string, year int, make string, model string, color string) Customer {
	vehicle := Vehicle{
		Year:  year,
		Make:  make,
		Model: model,
		Color: color,
	}
	return Customer{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
		Vehicle:     &vehicle,
	}
}

func FindCustomerByID(customerID int) *Customer {
	for _, c := range Customers {
		if c.ID == customerID {
			return c
		}
	}
	return nil
}

func UpdateCustomer(id int, name, phoneNumber string, vehicle Vehicle) error {
	c := FindCustomerByID(id)
	if c == nil {
		return fmt.Errorf("customer with ID %d not found", id)
	}

	if name != "" {
		c.Name = name
	}

	if phoneNumber != "" {
		c.PhoneNumber = phoneNumber
	}

	if vehicle.Year != 0 {
		c.Vehicle.Year = vehicle.Year
	}

	if vehicle.Make != "" {
		c.Vehicle.Make = vehicle.Make
	}

	if vehicle.Model != "" {
		c.Vehicle.Model = vehicle.Model
	}

	if vehicle.Color != "" {
		c.Vehicle.Color = vehicle.Color
	}

	return nil
}

type ServiceType string

const (
	OilChange = "Oil Change"
	CarWash   = "Car Wash"
)

type Service struct {
	CustomerID                int
	ServiceType               ServiceType
	Date                      time.Time
	CouponNotificationHistory []CouponNotification
	CouponApplied             bool // added flag indicating whether a coupon has been applied
	CouponType                CouponType
}

// ServiceHistory is a slice that keeps track of all services provided
var Services = []*Service{}

func AddService(customerID int, serviceType ServiceType, date time.Time) *Service {
	c := FindCustomerByID(customerID)
	if c == nil {
		fmt.Printf("customer with id %d not found\n", customerID)
		return nil
	}

	newService := Service{ServiceType: serviceType, Date: date}

	c.ServiceHistory = append(c.ServiceHistory, newService)

	fmt.Printf("Added new service with type %v and date %v for customer %s\n", serviceType, date.Format("01/02/2006"), c.Name)

	return &newService
}

func GenerateServiceReport(date time.Time) {
	fmt.Printf("Service Report for %v\n\n", date.Format("01/02/2006"))
	for _, customer := range Customers {
		for _, service := range customer.ServiceHistory {
			if service.Date.Before(date) || service.Date.Equal(date) {
				fmt.Printf("Customer: %s \nCar: %s %d %s %s \nService Type: %s\nLast Service Date: %s\n\n",
					customer.Name, customer.Vehicle.Color, customer.Vehicle.Year, customer.Vehicle.Make, customer.Vehicle.Model, service.ServiceType, service.Date.Format("01/02/2006"))
				break
			}
		}
	}
}

type ServiceTransactionHistory struct {
	Services             []Service
	CustomerIDToServices map[int][]Service
}

func (sth *ServiceTransactionHistory) AddService(service Service) {
	sth.Services = append(sth.Services, service)
	if _, ok := sth.CustomerIDToServices[service.CustomerID]; !ok {
		sth.CustomerIDToServices[service.CustomerID] = []Service{}
	}
	sth.CustomerIDToServices[service.CustomerID] = append(sth.CustomerIDToServices[service.CustomerID], service)
}

func (sth *ServiceTransactionHistory) GetServicesForCustomer(customerID int) []Service {
	if services, ok := sth.CustomerIDToServices[customerID]; ok {
		return services
	}
	return []Service{}
}
func GetLastServiceDate(c *Customer, serviceType ServiceType) (time.Time, error) {
	var lastService time.Time
	found := false
	for _, service := range c.ServiceHistory {
		if service.ServiceType == serviceType {
			if !found {
				lastService = service.Date
				found = true
			} else {
				if service.Date.After(lastService) {
					lastService = service.Date
				}
			}
		}
	}
	if !found {
		return time.Time{}, fmt.Errorf("No %s service found for customer %s", serviceType, c.Name)
	}
	return lastService, nil
}

type CouponType string

const (
	PercentOff  CouponType = "50% Off"
	FreeService CouponType = "Free Service"
)

type CouponNotification struct {
	CustomerID int
	CouponType CouponType
	Date       time.Time
}

func (cn CouponNotification) String() string {
	return fmt.Sprintf("Coupon Type: %s\nDate: %s\n", cn.CouponType, cn.Date.Format("01/02/2006"))
}

var CouponNotificationHistory []CouponNotification

func AddCouponNotification(customerID int, couponType CouponType, date time.Time, service *Service) error {
	c := FindCustomerByID(customerID)
	if c == nil {
		return fmt.Errorf("customer with id %d not found", customerID)
	}

	newCouponNotification := CouponNotification{CouponType: couponType, Date: date} // pointer created here
	service.CouponNotificationHistory = append(service.CouponNotificationHistory, newCouponNotification)
	FindCustomerByID(customerID).CouponNotificationHistory = append(FindCustomerByID(customerID).CouponNotificationHistory, &newCouponNotification)
	service.CouponApplied = false

	fmt.Printf("Customer %d recieved new coupon notification with type %v and date %v to service with type %v\n", customerID, couponType, date.Format("01/02/2006"), service.ServiceType)

	return nil
}

func SendOilChangeReminder() {
	for _, customer := range Customers {
		lastServiceDate, err := GetLastServiceDate(customer, OilChange)

		if err != nil {
			fmt.Println(err)
			continue
		}
		monthsSinceLastService := int(time.Since(lastServiceDate).Hours() / 24 / 30)
		if monthsSinceLastService >= 6 {
			fmt.Printf("Sending oil change reminder to %s's Number: %s for their %s %d %s %s.\n",
				customer.Name, customer.PhoneNumber, customer.Vehicle.Color, customer.Vehicle.Year, customer.Vehicle.Make, customer.Vehicle.Model)
		}
	}
}

func UpdateCouponNotification(cn CouponNotification) error {
	customer := FindCustomerByID(cn.CustomerID)
	if customer == nil {
		return fmt.Errorf("customer with ID %d not found", cn.CustomerID)
	}

	found := false
	for i, c := range customer.CouponNotificationHistory {
		if c.CouponType == cn.CouponType {
			customer.CouponNotificationHistory[i] = &cn
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("coupon notification with type %s not found for customer with ID %d", cn.CouponType, cn.CustomerID)
	}

	return nil
}

func DeleteCouponNotification(customerID int, couponType CouponType) error {
	for i, customer := range Customers {
		if customer.ID == customerID {
			for j, notification := range customer.CouponNotificationHistory {
				if notification.CouponType == couponType {
					CouponNotificationHistory = append(CouponNotificationHistory[:j], CouponNotificationHistory[j+1:]...)
					Customers[i] = customer
					return nil
				}
			}
			return fmt.Errorf("customer with ID %d does not have a %s coupon notification", customerID, couponType)
		}
	}
	return fmt.Errorf("could not find customer with ID %d", customerID)
}

func GenerateCouponNotificationReport(customerID int) {
	for _, cust := range Customers {
		if cust.ID == customerID {
			fmt.Printf("Coupon Notification Report for Customer %s\n\n", cust.Name)
			serviceTypeToLastCoupon := make(map[ServiceType]*CouponNotification)
			for _, serv := range cust.ServiceHistory {
				lastCouponNotification := getLastCouponNotification(serv)
				if lastCouponNotification != nil {
					if prevLastCoupon, ok := serviceTypeToLastCoupon[serv.ServiceType]; !ok || lastCouponNotification.Date.After(prevLastCoupon.Date) {
						serviceTypeToLastCoupon[serv.ServiceType] = lastCouponNotification
					}
				}
			}
			for serviceType, lastCouponNotification := range serviceTypeToLastCoupon {
				fmt.Printf("Service Type: %s\n", serviceType)
				fmt.Printf("Coupon Type: %s\n\n", lastCouponNotification.CouponType)
			}
		}
	}
}

func getLastCouponNotification(serv Service) *CouponNotification {
	if len(serv.CouponNotificationHistory) == 0 {
		return nil
	}
	lastCouponNotification := serv.CouponNotificationHistory[len(serv.CouponNotificationHistory)-1]
	return &lastCouponNotification
}

func main() {
	customer1 := NewCustomer(10000, "John Doe", "555-1234", 2010, "Honda", "Accord", "Silver")
	customer2 := NewCustomer(22222, "Jane Smith", "555-5678", 2015, "Toyota", "Corolla", "White")
	customer3 := NewCustomer(30303, "Bob Johnson", "555-9012", 2012, "Ford", "Focus", "Black")
	customer4 := NewCustomer(12345, "Adam Ford", "513-8283", 2001, "Jeep", "Grand Cherokee", "Green")

	Customers = append(Customers, &customer1, &customer2, &customer3, &customer4)

	service4 := AddService(10000, OilChange, time.Now().AddDate(0, -1, 0))
	service5 := AddService(22222, CarWash, time.Now().AddDate(0, 0, -7))
	service1 := AddService(30303, OilChange, time.Now().AddDate(0, -7, 1))
	service2 := AddService(30303, CarWash, time.Now().AddDate(0, -6, -6))
	service3 := AddService(30303, CarWash, time.Now().AddDate(-1, -6, -6))
	service6 := AddService(12345, OilChange, time.Now().AddDate(0, -5, -30))

	Services = append(Services, service1, service2, service3, service4, service5, service6)

	GenerateServiceReport(time.Now())

	AddCouponNotification(30303, FreeService, time.Now().AddDate(1, 1, 3), service1)
	AddCouponNotification(30303, PercentOff, time.Now().AddDate(1, 1, 3), service2)
	AddCouponNotification(30303, PercentOff, time.Now().AddDate(0, -2, 0), service3)

	GenerateCouponNotificationReport(30303)

	SendOilChangeReminder()

	//create CSV file to store info
	cData, err := os.Create("customerdata.csv")
	if err != nil {
		panic(err)
	}

	// use defer to close csv file
	defer cData.Close()

	//call new csv writer vehicle service history and couponhistory
	writer := csv.NewWriter(cData)
	for _, c := range Customers {
		row := []string{fmt.Sprintf("ID: %d", c.ID), fmt.Sprintf("Name: %s", c.Name), fmt.Sprintf("Phone #:%s", c.PhoneNumber),
			fmt.Sprintf("Vehicle Year: %d", c.Vehicle.Year), fmt.Sprintf("Color: %s", c.Vehicle.Color), fmt.Sprintf("Make: %s", c.Vehicle.Make),
			fmt.Sprintf("Model: %s", c.Vehicle.Model), fmt.Sprintf("Sevice History: %v", c.ServiceHistory), fmt.Sprintf("Coupon History %v", c.CouponNotificationHistory)}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
	writer.Flush()
}
