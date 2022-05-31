package constant

const (
	OrderStatusWaitingApproved string = "waiting_approved"
	OrderStatusApproved        string = "approved"
	OrderStatusDelivering      string = "delivering"
	OrderStatusDelivered       string = "delivered"
	OrderStatusCompleted       string = "completed"
	OrderStatusCanceled        string = "canceled"
)

const (
	PaymentMethodSenderPay   = "sender_pay"
	PaymentMethodReceiverPay = "receiver_pay"
)
