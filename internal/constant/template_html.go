package constant

const TemplateHTMLForgotPassword = `
<!DOCTYPE html>
<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="vi-VN">

<head>
	<title></title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]-->
	<!--[if !mso]><!-->
	<link href="https://fonts.googleapis.com/css?family=Lato" rel="stylesheet" type="text/css">
	<link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet" type="text/css">
	<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" type="text/css">
	<!--<![endif]-->
	<style>
		* {
			box-sizing: border-box;
		}

		body {
			margin: 0;
			padding: 0;
		}

		a[x-apple-data-detectors] {
			color: inherit !important;
			text-decoration: inherit !important;
		}

		#MessageViewBody a {
			color: inherit;
			text-decoration: none;
		}

		p {
			line-height: inherit
		}

		@media (max-width:520px) {
			.row-content {
				width: 100% !important;
			}
		}
	</style>
</head>

<body style="background-color: #FFFFFF; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;">
	<table class="nl-container" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF;">
		<tbody>
			<tr>
				<td>
					<table class="row row-1" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
						<tbody>
							<tr>
								<td>
									<table class="row-content" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;" width="500">
										<tbody>
											<tr>
												<td class="column" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-right: 20px; padding-left: 20px; padding-top: 20px; padding-bottom: 20px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;">
													<table class="image_block" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
														<tr>
															<td style="width:100%;padding-right:0px;padding-left:0px;">
																<div style="line-height:10px"><a href="https://www.gogo.vn/" target="_blank" style="outline:none" tabindex="-1"><img src="https://go-go.s3.ap-southeast-1.amazonaws.com/gogo-logo.png" style="display: block; height: auto; border: 0; width: 92px; max-width: 100%;" width="80" alt="GoGo logo" title="GoGo logo"></a></div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td style="padding-top:30px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 16px; text-align: center;"><strong>Đổi mật khẩu</strong></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td style="padding-top:30px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left;">Bạn có thể đặt lại mật khẩu GoGo bằng cách click vào link bên dưới:</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left;"><a href="{{.URI}}" target="_blank" style="text-decoration: none; color: #cc3366;" rel="noopener"><strong>Thay đổi mật khẩu</strong></a></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left;">Nếu đây không phải là yêu cầu đặt lại mật khẩu của bạn, đừng ngại xóa email này!</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left; letter-spacing: normal;"><strong>GoGo</strong></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
												</td>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table><!-- End -->
</body>

</html>
`

const TemplateHTMLNewUser = `
<!DOCTYPE html>
<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" lang="vi-VN">

<head>
	<title></title>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]-->
	<!--[if !mso]><!-->
	<link href="https://fonts.googleapis.com/css?family=Lato" rel="stylesheet" type="text/css">
	<link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet" type="text/css">
	<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" type="text/css">
	<!--<![endif]-->
	<style>
		* {
			box-sizing: border-box;
		}

		body {
			margin: 0;
			padding: 0;
		}

		a[x-apple-data-detectors] {
			color: inherit !important;
			text-decoration: inherit !important;
		}

		#MessageViewBody a {
			color: inherit;
			text-decoration: none;
		}

		p {
			line-height: inherit
		}

		@media (max-width:520px) {
			.row-content {
				width: 100% !important;
			}
		}
	</style>
</head>

<body style="background-color: #FFFFFF; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;">
	<table class="nl-container" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF;">
		<tbody>
			<tr>
				<td>
					<table class="row row-1" align="center" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
						<tbody>
							<tr>
								<td>
									<table class="row-content" align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;" width="500">
										<tbody>
											<tr>
												<td class="column" width="100%" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-right: 20px; padding-left: 20px; padding-top: 20px; padding-bottom: 20px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;">
													<table class="image_block" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;">
														<tr>
															<td style="width:100%;padding-right:0px;padding-left:0px;">
																<div style="line-height:10px"><a href="https://www.gogo.vn/" target="_blank" style="outline:none" tabindex="-1"><img src="https://go-go.s3.ap-southeast-1.amazonaws.com/gogo-logo.png" style="display: block; height: auto; border: 0; width: 80px; max-width: 100%;" width="92" alt="GoGo logo" title="GoGo logo"></a></div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="0" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td style="padding-top:30px;padding-right:10px;padding-bottom:10px;padding-left:10px;">
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 16px; text-align: center;"><strong>Tạo tài khoản thành công</strong></p>
																	</div>
																</div>
															</td>
														</tr>
                                                    <table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left; color: #cc3366;"><strong>Tên đăng nhập: {{.Account}}</strong></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left; color: #cc3366;"><strong>Mật khẩu: {{.Password}}</strong></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left;">Sử dụng email và mật khẩu trên để đăng nhập vào ứng dụng</p>
																		<p style="margin: 0; font-size: 14px; text-align: left;">GoGo!</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table class="text_block" width="100%" border="0" cellpadding="10" cellspacing="0" role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div style="font-size: 12px; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; mso-line-height-alt: 14.399999999999999px; color: #555555; line-height: 1.2;">
																		<p style="margin: 0; font-size: 14px; text-align: left; letter-spacing: normal;"><strong>GoGo</strong></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
												</td>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table><!-- End -->
</body>

</html>
`
