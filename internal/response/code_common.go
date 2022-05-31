package response

const (
	CommonSuccess      = "common_success"
	CommonBadRequest   = "common_bad_request"
	CommonUnauthorized = "common_unauthorized"
	CommonNotFound     = "common_not_found"

	CommonInvalidChecksum      = "common_invalid_checksum"
	CommonInvalidPagination    = "common_invalid_pagination"
	CommonInvalidName          = "common_invalid_name"
	CommonInvalidPhone         = "common_invalid_phone"
	CommonInvalidEmail         = "common_invalid_email"
	CommonInvalidPassword      = "common_invalid_password"
	CommonInvalidOldPassword   = "common_invalid_old_password"
	CommonInvalidNewPassword   = "common_invalid_new_password"
	CommonInvalidPermissions   = "common_invalid_permissions"
	CommonInvalidRole          = "common_invalid_role"
	CommonInvalidStatus        = "common_invalid_status"
	CommonInvalidPhoto         = "common_invalid_photo"
	CommonInvalidToken         = "common_invalid_token"
	CommonInvalidNote          = "common_invalid_note"
	CommonInvalidUser          = "common_invalid_user"
	CommonInvalidAddress       = "common_invalid_address"
	CommonInvalidProvince      = "common_invalid_province"
	CommonInvalidDistrict      = "common_invalid_district"
	CommonInvalidWard          = "common_invalid_ward"
	CommonInvalidInventory     = "common_invalid_inventory"
	CommonInvalidWeight        = "common_invalid_weight"
	CommonInvalidHeight        = "common_invalid_height"
	CommonInvalidWidth         = "common_invalid_width"
	CommonInvalidLength        = "common_invalid_length"
	CommonInvalidCode          = "common_invalid_code"
	CommonInvalidPaymentMethod = "common_invalid_payment_method"

	CommonExistedName  = "common_existed_name"
	CommonExistedPhone = "common_existed_phone"
	CommonExistedEmail = "common_existed_email"

	CommonNotFoundUser      = "common_not_found_user"
	CommonNotFoundRole      = "common_not_found_role"
	CommonNotFoundStaff     = "common_not_found_staff"
	CommonNotFoundOrder     = "common_not_found_order"
	CommonNotFoundInventory = "common_not_found_inventory"
	CommonNotExistedEmail   = "common_not_existed_email"

	CommonNoPermission       = "common_no_permission"
	CommonCannotChangeStatus = "common_cannot_change_status"
	CommonCannotCancelStatus = "common_cannot_cancel_status"
	CommonCannotWarnStatus   = "common_cannot_warn_status"
	CommonCannotUpdate       = "common_cannot_update"

	CommonResetPasswordFailed = "common_reset_password_failed"
	CommonErrorWhenHandler    = "common_error_when_handler"
	CommonIncorrectPassword   = "common_incorrect_password"
	CommonBannedUser          = "common_banned_user"
)

// 1 - 199
var common = []Code{
	{
		Key:     CommonSuccess,
		Message: "thành công",
		Code:    1,
	},
	{
		Key:     CommonBadRequest,
		Message: "dữ liệu không hợp lệ",
		Code:    2,
	},
	{
		Key:     CommonUnauthorized,
		Message: "bạn không có quyền thực hiện hành động này",
		Code:    3,
	},
	{
		Key:     CommonNotFound,
		Message: "dữ liệu không tìm thấy",
		Code:    4,
	},
	{
		Key:     CommonInvalidChecksum,
		Message: "xác thực dữ liệu thất bại",
		Code:    5,
	},
	{
		Key:     CommonInvalidPagination,
		Message: "phân trang không hợp lệ",
		Code:    6,
	},
	{
		Key:     CommonInvalidName,
		Message: "tên không hợp lệ",
		Code:    7,
	},
	{
		Key:     CommonInvalidPhone,
		Message: "số điện thoại không hợp lệ",
		Code:    8,
	},
	{
		Key:     CommonInvalidEmail,
		Message: "email không hợp lệ",
		Code:    9,
	},
	{
		Key:     CommonInvalidPassword,
		Message: "mật khẩu hợp lệ",
		Code:    10,
	},
	{
		Key:     CommonInvalidPermissions,
		Message: "danh sách quyền không hợp lệ",
		Code:    11,
	},
	{
		Key:     CommonExistedName,
		Message: "tên đã tồn tại trong hệ thống",
		Code:    12,
	},
	{
		Key:     CommonExistedPhone,
		Message: "số điện thoại đã tồn tại trong hệ thống",
		Code:    13,
	},
	{
		Key:     CommonExistedEmail,
		Message: "email đã tồn tại trong hệ thống",
		Code:    14,
	},
	{
		Key:     CommonInvalidRole,
		Message: "vai trò không hợp lệ",
		Code:    15,
	},
	{
		Key:     CommonInvalidUser,
		Message: "người dùng không hợp lệ",
		Code:    16,
	},
	{
		Key:     CommonNotFoundUser,
		Message: "người dùng không tồn tại",
		Code:    18,
	},
	{
		Key:     CommonNotFoundRole,
		Message: "vai trò không tồn tại",
		Code:    19,
	},
	{
		Key:     CommonNotFoundOrder,
		Message: "đơn hàng không tồn tại",
		Code:    20,
	},
	{
		Key:     CommonInvalidStatus,
		Message: "trạng thái không hợp lệ",
		Code:    21,
	},
	{
		Key:     CommonNotFoundStaff,
		Message: "nhân viên không tồn tại",
		Code:    22,
	},
	{
		Key:     CommonNoPermission,
		Message: "bạn không có quyền thực hiện hành động này",
		Code:    23,
	},
	{
		Key:     CommonInvalidOldPassword,
		Message: "mật khẩu cũ không hợp lệ",
		Code:    24,
	},
	{
		Key:     CommonInvalidNewPassword,
		Message: "mật khẩu mới không hợp lệ",
		Code:    25,
	},
	{
		Key:     CommonInvalidNewPassword,
		Message: "ảnh không hợp lệ",
		Code:    26,
	},
	{
		Key:     CommonBannedUser,
		Message: "tài khoản của bạn đã bị khóa",
		Code:    27,
	},
	{
		Key:     CommonCannotChangeStatus,
		Message: "không thể thay đổi trạng thái",
		Code:    28,
	},
	{
		Key:     CommonCannotCancelStatus,
		Message: "không thể hủy đơn hàng này",
		Code:    29,
	},
	{
		Key:     CommonCannotWarnStatus,
		Message: "không thể cảnh báo đơn hàng này",
		Code:    30,
	},
	{
		Key:     CommonInvalidNote,
		Message: "ghi chú không hợp lệ",
		Code:    31,
	},
	{
		Key:     CommonResetPasswordFailed,
		Message: "đặt lại mật khẩu không thành công",
		Code:    32,
	},
	{
		Key:     CommonInvalidToken,
		Message: "token không hợp lệ",
		Code:    33,
	},
	{
		Key:     CommonNotExistedEmail,
		Message: "email không tồn tại",
		Code:    34,
	},
	{
		Key:     CommonErrorWhenHandler,
		Message: "có lỗi trong quá trình xử lý",
		Code:    35,
	},
	{
		Key:     CommonInvalidAddress,
		Message: "địa chỉ không hợp lệ",
		Code:    36,
	},
	{
		Key:     CommonInvalidProvince,
		Message: "tỉnh, thành phố không hợp lệ",
		Code:    37,
	},
	{
		Key:     CommonInvalidDistrict,
		Message: "quận, huyện không hợp lệ",
		Code:    38,
	},
	{
		Key:     CommonInvalidWard,
		Message: "phường, xã không hợp lệ",
		Code:    39,
	},
	{
		Key:     CommonNotFoundInventory,
		Message: "kho hàng không tồn tại",
		Code:    40,
	},
	{
		Key:     CommonInvalidInventory,
		Message: "kho hàng không hợp lệ",
		Code:    41,
	},
	{
		Key:     CommonCannotUpdate,
		Message: "không thể cập nhật",
		Code:    42,
	},
}
