package constants

type StoreTypeEnum string

const (
	CENTRAL StoreTypeEnum = "CENTRAL"
	BRANCH  StoreTypeEnum = "BRANCH"
)

type RoleTypeEnum string

const (
	OWNER      RoleTypeEnum = "OWNER"
	SUPERVISOR RoleTypeEnum = "SUPERVISOR"
	ADMIN      RoleTypeEnum = "ADMIN"
)

type AccountStatusEnum string

const (
	ACTIVE    AccountStatusEnum = "ACTIVE"
	SUSPENDED AccountStatusEnum = "SUSPENDED"
)

type OrderStatusEnum string

const (
	PENDING   OrderStatusEnum = "PENDING"
	APPROVED  OrderStatusEnum = "APPROVED"
	SHIPPED   OrderStatusEnum = "SHIPPED"
	RECEIVED  OrderStatusEnum = "RECEIVED"
	CANCELLED OrderStatusEnum = "CANCELLED"
)

type DiscountTypeEnum string

const (
	PERCENTAGE DiscountTypeEnum = "PERCENTAGE"
	FIXED      DiscountTypeEnum = "FIXED"
)
