package constant

// DeliverySource ...
type DeliverySource int

// DeliverySourceName ...
var DeliverySourceName = []string{
	"ghn",
	"ghtk",
}

// String ...
func (d DeliverySource) String() string {
	if d < 0 || int(d) > len(DeliverySourceName) {
		return ""
	}
	return DeliverySourceName[d-1]
}

const (
	DeliverySourceGHN DeliverySource = iota + 1
	DeliverySourceBoxme
)

// Delivery status
const (
	DeliveryStatusWaitingToConfirm = "waiting_to_confirm"
	DeliveryStatusWaitingToPick    = "waiting_to_pick"
	DeliveryStatusPicking          = "picking"
	DeliveryStatusPickupFailed     = "pickup_failed"
	DeliveryStatusPicked           = "picked"
	DeliveryStatusDelayPickup      = "delay_pickup"

	DeliveryStatusDelivering     = "delivering"
	DeliveryStatusDelivered      = "delivered"
	DeliveryStatusDeliveryFailed = "delivery_failed"
	DeliveryStatusDelayDelivery  = "delay_delivery"

	DeliveryStatusWaitingToReturn = "waiting_to_return"
	DeliveryStatusReturning       = "returning"
	DeliveryStatusReturned        = "returned"
	DeliveryStatusReturnFailed    = "return_failed"
	DeliveryStatusDelayReturn     = "delay_return"

	DeliveryStatusCancelled    = "cancelled"
	DeliveryStatusCompensation = "compensation"
)

var (
	// DeliveryStatusNameMapping ...
	DeliveryStatusNameMapping = map[string]string{
		DeliveryStatusWaitingToPick:   "Chờ lấy hàng",
		DeliveryStatusPicking:         "Đang lấy hàng",
		DeliveryStatusPickupFailed:    "Lấy hàng thất bại",
		DeliveryStatusPicked:          "Đã lấy hàng",
		DeliveryStatusDelayPickup:     "Delay lấy hàng",
		DeliveryStatusDelivering:      "Đang giao hàng",
		DeliveryStatusDelivered:       "Đã giao hàng",
		DeliveryStatusDeliveryFailed:  "Giao hàng thất bại",
		DeliveryStatusDelayDelivery:   "Delay giao hàng",
		DeliveryStatusWaitingToReturn: "Chờ trả hàng",
		DeliveryStatusReturning:       "Đang trả hàng",
		DeliveryStatusReturned:        "Đã trả hàng",
		DeliveryStatusReturnFailed:    "Trả hàng thất bại",
		DeliveryStatusDelayReturn:     "Delay trả hàng",
		DeliveryStatusCancelled:       "Hủy",
		DeliveryStatusCompensation:    "Thất lạc",
	}
)
