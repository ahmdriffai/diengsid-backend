package mailview

import "fmt"

func RegisterOtpMailView(code string) string {

	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8" />
		</head>
		<style>
		.body {
			margin: 0;
			background: #f2f2f2;
			font-family: Arial, sans-serif;
		}

		.wrapper {
			width: 100%%;
		}

		.wrapper-cell {
			padding: 40px 0;
		}

		.container {
			width: 600px;
			background: #ffffff;
			padding: 40px;
			border-radius: 16px;
			border: 1px solid #e5e5e5;
		}

		/* Logo */
		.logo {
			padding-bottom: 24px;
		}

		/* Title */
		.title {
			font-size: 32px;
			font-weight: 700;
			color: #222;
			line-height: 1.3;
		}

		/* Description */
		.description {
			padding-top: 16px;
			font-size: 16px;
			color: #444;
			line-height: 1.6;
		}

		/* Code */
		.code {
			padding-top: 28px;
			font-size: 40px;
			font-weight: 700;
			color: #222;
			letter-spacing: 2px;
		}

		/* Time */
		.time {
			padding-top: 8px;
			font-size: 14px;
			color: #777;
		}

		/* Help */
		.help-title {
			padding-top: 28px;
			font-size: 18px;
			font-weight: 600;
			color: #222;
		}

		.help-text {
			padding-top: 12px;
			font-size: 16px;
			color: #444;
			line-height: 1.6;
		}

		/* Link */
		.link {
			color: #222;
			font-weight: 600;
			text-decoration: underline;
		}

		/* Divider */
		.divider {
			padding: 32px 0;
		}

		.divider hr {
			border: none;
			border-top: 1px solid #eee;
		}

		/* Footer */
		.footer-logo {
			padding-bottom: 16px;
		}

		.address {
			font-size: 14px;
			color: #555;
			line-height: 1.6;
		}
		</style>
		<body style="margin:0; background:#eee; font-family:Arial, sans-serif;">
			<table width="100%%" cellpadding="0" cellspacing="0">
				<tr>
					<td align="center" style="padding:40px 0;">
						<table width="600" style="background:#ffffff; padding:40px; border-radius:16px; border:1px solid #e5e5e5;">
							
							<!-- Logo Airbnb -->
							<tr>
								<td style="padding-bottom:24px;">
									<img src="https://www.image2url.com/r2/default/images/1776225307615-4af4606b-ecc0-476c-a8b9-9f965397be27.png" alt="image" width="32" />
								</td>
							</tr>

							<!-- Title -->
							<tr>
								<td style="font-size:32px; font-weight:700; color:#222; line-height:1.3;">
									Kode pengaman Diengs.id Anda
								</td>
							</tr>

							<!-- Description -->
							<tr>
								<td style="padding-top:16px; font-size:16px; color:#444; line-height:1.6;">
									Jangan pernah membagikan kode Anda dengan siapa pun--karyawan Diengs.id tidak akan pernah memintanya.
								</td>
							</tr>

							<!-- Code -->
							<tr>
								<td style="padding-top:28px; font-size:40px; font-weight:700; color:#222; letter-spacing:2px;">
									%s
								</td>
							</tr>

							<!-- Time -->
							<tr>
								<td style="padding-top:8px; font-size:14px; color:#777;">
									2026-04-05 15:39:15
								</td>
							</tr>

							<!-- Help -->
							<tr>
								<td style="padding-top:28px; font-size:18px; font-weight:600; color:#222;">
									Tidak mengenali ini?
								</td>
							</tr>

							<tr>
								<td style="padding-top:12px; font-size:16px; color:#444; line-height:1.6;">
									<a href="#" style="color:#222; font-weight:600; text-decoration:underline;">Beri tahu kami</a>
									— kami akan membantu mengamankan dan meninjau akun Anda. Bila tidak, Anda tidak perlu melakukan apa-apa.
								</td>
							</tr>

							<!-- Divider -->
							<tr>
								<td style="padding:32px 0;">
									<hr style="border:none; border-top:1px solid #eee;" />
								</td>
							</tr>

							<!-- Footer Logo -->
							<tr>
								<td style="padding-bottom:16px;">
									<img src="https://www.image2url.com/r2/default/images/1776225307615-4af4606b-ecc0-476c-a8b9-9f965397be27.png" alt="image" width="32" />
								</td>
							</tr>

							<!-- Address -->
							<tr>
								<td style="font-size:14px; color:#555; line-height:1.6;">
									Diengsid<br/>
									8 Hanover Quay<br/>
									Dublin 2, Ireland
								</td>
							</tr>

						</table>
					</td>
				</tr>
			</table>
		</body>
		</html>

	`, code)
}
