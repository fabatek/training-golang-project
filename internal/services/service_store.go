package services

// Container contains all service interfaces
type Container struct {
	User    IUser
	Product IProduct
	Order   IOrder
}
