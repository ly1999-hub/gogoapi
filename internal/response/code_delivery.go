package response

const (
	DeliveryLocationNotMatch = "delivery_location_not_match"
	DeliveryServiceNotFound = "delivery_service_not_found"
	DeliveryCreateDeliveryFailed = "delivery_create_delivery_failed"
)

// 200 - 299
var delivery = []Code{
	{
		Key:     DeliveryLocationNotMatch,
		Message: "Dữ liệu địa điểm giao hàng chưa được cập nhật, vui lòng liên hệ với chúng tôi để được hỗ trợ.",
		Code:    200,
	},
	{
		Key:     DeliveryServiceNotFound,
		Message: "Không tìm thấy dịch vụ vận chuyển",
		Code:    201,
	},
	{
		Key:     DeliveryCreateDeliveryFailed,
		Message: "Tạo đơn vị vận chuyển thất bại",
		Code:    202,
	},
}
