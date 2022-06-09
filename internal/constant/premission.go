package constant

// Permissions
const (
	PermissionRoleView = "role_view"
	PermissionRoleEdit = "role_edit"

	PermissionOrderView = "order_view"
	PermissionOrderEdit = "order_edit"

	PermissionWarehouseView = "warehouse_view"
	PermissionWarehouseEdit = "warehouse_edit"

	PermissionUserView = "user_view"
	PermissionUserEdit = "user_edit"
)

type Permission struct {
	Name string `bson:"name" json:"name"`
	Code string `bson:"code" json:"code"`
}

var (
	Permissions = []Permission{
		// Account
		{
			Name: "Role view",
			Code: PermissionRoleView,
		},
		{
			Name: "Role edit",
			Code: PermissionRoleEdit,
		},

		// Order
		{
			Name: "Order view",
			Code: PermissionOrderView,
		},
		{
			Name: "Order edit",
			Code: PermissionOrderEdit,
		},

		// Warehouse
		{
			Name: "Warehouse view",
			Code: PermissionWarehouseView,
		},
		{
			Name: "Warehouse edit",
			Code: PermissionWarehouseEdit,
		},

		// User
		{
			Name: "User view",
			Code: PermissionUserView,
		},
		{
			Name: "User edit",
			Code: PermissionUserEdit,
		},
	}
)
