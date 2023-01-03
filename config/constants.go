package config

const (
	// Config Constant
	DEFAULT_DATE_LAYOUT       = "2006-01-02"
	DEFAULT_DATETIME_LAYOUT   = "2006-01-02 15:04:05"
	DEFAULT_DATETIME_LAYOUT_Z = "2006-01-02 15:04:05 -0700"
	DEFAULT_DATETIME_LOCATION = "Asia/Jakarta"

	// Partner verification constant
	PARTNER_VERIFICATION_STATUS_NOT_VERIFIED = "Not Verified"
	PARTNER_VERIFICATION_STATUS_REVISION     = "Revision"
	PARTNER_VERIFICATION_STATUS_VERIFIED     = "Verified"

	// Order Constant
	ORDER_STATUS_WAITING_FOR_PAYMENT  = "Waiting For Payment"
	ORDER_STATUS_WAITING_CONFIRMATION = "Waiting Confirmation"
	ORDER_STATUS_ORDER_CONFIRMED      = "Order Confirmed"
	ORDER_STATUS_COMPLETE_ORDER       = "Complete Order"
	ORDER_STATUS_PAID_OFF             = "Paid Off"

	//Show log
	SHOW_LOGS = true

	DEFAULT_IMAGE_URL = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/partner/EEz06AIRAiyJe4ghKfU5-default_image.jpg"

	// Upload file path
	CLIENT_FOLDER     = "client"
	CLIENT_IMAGE_FILE = "client_image_file"

	PARTNER_FOLDER     = "partner"
	COMPANY_IMAGE_FILE = "company_image_file"
	NIB_IMAGE_FILE     = "nib_image_file"
	SIUP_IMAGE_FILE    = "siup_image_file"
	EVENT1_IMAGE_FILE  = "event1_image_file"
	EVENT2_IMAGE_FILE  = "event2_image_file"
	EVENT3_IMAGE_FILE  = "event3_image_file"

	SERVICE_FOLDER     = "service"
	SERVICE_IMAGE_FILE = "service_image_file"

	ORDER_FOLDER              = "order"
	ORDER_PAYOUT_RECEIPT_FILE = "payout_receipt_file"
)

// Midtrans
type PaymentMethod string

const (
	//VABni : bni
	VABNI PaymentMethod = "va bni"

	//VAMandiri : mandiri
	VAMandiri PaymentMethod = "va mandiri"

	//VACimb : cimb
	VACimb PaymentMethod = "va cimb"

	//VABca : bca
	VABca PaymentMethod = "va bca"

	//VABri : bri
	VABri PaymentMethod = "va bri"

	//VAMaybank : maybank
	VAMaybank PaymentMethod = "va maybank"

	//VAPermata : permata
	VAPermata PaymentMethod = "va permata"

	//VAMega : mega
	VAMega PaymentMethod = "va mega"
)

const PAYMENT_EXPIRED_DURATION = 1
const PAYMENT_EXPIRED_UNIT = "day"
