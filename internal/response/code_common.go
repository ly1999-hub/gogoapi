package response

const (
	CommonSuccess      = "common_success"
	CommonBadRequest   = "common_bad_request"
	CommonUnauthorized = "common_unauthorized"
	CommonNotFound     = "common_not_found"

	CommonInvalidChecksum    = "common_invalid_checksum"
	CommonInvalidPagination  = "common_invalid_pagination"
	CommonInvalidName        = "common_invalid_name"
	CommonInvalidPhone       = "common_invalid_phone"
	CommonInvalidEmail       = "common_invalid_email"
	CommonInvalidPassword    = "common_invalid_password"
	CommonInvalidOldPassword = "common_invalid_old_password"
	CommonInvalidNewPassword = "common_invalid_new_password"
	CommonInvalidPermissions = "common_invalid_permissions"
	CommonInvalidRole        = "common_invalid_role"
	CommonInvalidStatus      = "common_invalid_status"
	CommonInvalidToken       = "common_invalid_token"
	CommonInvalidNote        = "common_invalid_note"
	CommonInvalidUser        = "common_invalid_user"
	CommonInvalidAddress     = "common_invalid_address"
	CommonInvalidProvince    = "common_invalid_province"
	CommonInvalidDistrict    = "common_invalid_district"
	CommonInvalidWard        = "common_invalid_ward"
	CommonInvalidInventory   = "common_invalid_inventory"

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
		Message: "th??nh c??ng",
		Code:    1,
	},
	{
		Key:     CommonBadRequest,
		Message: "d??? li???u kh??ng h???p l???",
		Code:    2,
	},
	{
		Key:     CommonUnauthorized,
		Message: "b???n kh??ng c?? quy???n th???c hi???n h??nh ?????ng n??y",
		Code:    3,
	},
	{
		Key:     CommonNotFound,
		Message: "d??? li???u kh??ng t??m th???y",
		Code:    4,
	},
	{
		Key:     CommonInvalidChecksum,
		Message: "x??c th???c d??? li???u th???t b???i",
		Code:    5,
	},
	{
		Key:     CommonInvalidPagination,
		Message: "ph??n trang kh??ng h???p l???",
		Code:    6,
	},
	{
		Key:     CommonInvalidName,
		Message: "t??n kh??ng h???p l???",
		Code:    7,
	},
	{
		Key:     CommonInvalidPhone,
		Message: "s??? ??i???n tho???i kh??ng h???p l???",
		Code:    8,
	},
	{
		Key:     CommonInvalidEmail,
		Message: "email kh??ng h???p l???",
		Code:    9,
	},
	{
		Key:     CommonInvalidPassword,
		Message: "m???t kh???u h???p l???",
		Code:    10,
	},
	{
		Key:     CommonInvalidPermissions,
		Message: "danh s??ch quy???n kh??ng h???p l???",
		Code:    11,
	},
	{
		Key:     CommonExistedName,
		Message: "t??n ???? t???n t???i trong h??? th???ng",
		Code:    12,
	},
	{
		Key:     CommonExistedPhone,
		Message: "s??? ??i???n tho???i ???? t???n t???i trong h??? th???ng",
		Code:    13,
	},
	{
		Key:     CommonExistedEmail,
		Message: "email ???? t???n t???i trong h??? th???ng",
		Code:    14,
	},
	{
		Key:     CommonInvalidRole,
		Message: "vai tr?? kh??ng h???p l???",
		Code:    15,
	},
	{
		Key:     CommonInvalidUser,
		Message: "ng?????i d??ng kh??ng h???p l???",
		Code:    16,
	},
	{
		Key:     CommonNotFoundUser,
		Message: "ng?????i d??ng kh??ng t???n t???i",
		Code:    18,
	},
	{
		Key:     CommonNotFoundRole,
		Message: "vai tr?? kh??ng t???n t???i",
		Code:    19,
	},
	{
		Key:     CommonNotFoundOrder,
		Message: "????n h??ng kh??ng t???n t???i",
		Code:    20,
	},
	{
		Key:     CommonInvalidStatus,
		Message: "tr???ng th??i kh??ng h???p l???",
		Code:    21,
	},
	{
		Key:     CommonNotFoundStaff,
		Message: "nh??n vi??n kh??ng t???n t???i",
		Code:    22,
	},
	{
		Key:     CommonNoPermission,
		Message: "b???n kh??ng c?? quy???n th???c hi???n h??nh ?????ng n??y",
		Code:    23,
	},
	{
		Key:     CommonInvalidOldPassword,
		Message: "m???t kh???u c?? kh??ng h???p l???",
		Code:    24,
	},
	{
		Key:     CommonInvalidNewPassword,
		Message: "m???t kh???u m???i kh??ng h???p l???",
		Code:    25,
	},
	{
		Key:     CommonInvalidNewPassword,
		Message: "???nh kh??ng h???p l???",
		Code:    26,
	},
	{
		Key:     CommonBannedUser,
		Message: "t??i kho???n c???a b???n ???? b??? kh??a",
		Code:    27,
	},
	{
		Key:     CommonCannotChangeStatus,
		Message: "kh??ng th??? thay ?????i tr???ng th??i",
		Code:    28,
	},
	{
		Key:     CommonCannotCancelStatus,
		Message: "kh??ng th??? h???y ????n h??ng n??y",
		Code:    29,
	},
	{
		Key:     CommonCannotWarnStatus,
		Message: "kh??ng th??? c???nh b??o ????n h??ng n??y",
		Code:    30,
	},
	{
		Key:     CommonInvalidNote,
		Message: "ghi ch?? kh??ng h???p l???",
		Code:    31,
	},
	{
		Key:     CommonResetPasswordFailed,
		Message: "?????t l???i m???t kh???u kh??ng th??nh c??ng",
		Code:    32,
	},
	{
		Key:     CommonInvalidToken,
		Message: "token kh??ng h???p l???",
		Code:    33,
	},
	{
		Key:     CommonNotExistedEmail,
		Message: "email kh??ng t???n t???i",
		Code:    34,
	},
	{
		Key:     CommonErrorWhenHandler,
		Message: "c?? l???i trong qu?? tr??nh x??? l??",
		Code:    35,
	},
	{
		Key:     CommonInvalidAddress,
		Message: "?????a ch??? kh??ng h???p l???",
		Code:    36,
	},
	{
		Key:     CommonInvalidProvince,
		Message: "t???nh, th??nh ph??? kh??ng h???p l???",
		Code:    37,
	},
	{
		Key:     CommonInvalidDistrict,
		Message: "qu???n, huy???n kh??ng h???p l???",
		Code:    38,
	},
	{
		Key:     CommonInvalidWard,
		Message: "ph?????ng, x?? kh??ng h???p l???",
		Code:    39,
	},
	{
		Key:     CommonNotFoundInventory,
		Message: "kho h??ng kh??ng t???n t???i",
		Code:    40,
	},
	{
		Key:     CommonInvalidInventory,
		Message: "kho h??ng kh??ng h???p l???",
		Code:    41,
	},
	{
		Key:     CommonCannotUpdate,
		Message: "kh??ng th??? c???p nh???t",
		Code:    42,
	},
}
