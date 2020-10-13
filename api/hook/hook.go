package hook

import "github.com/eiprilvinesipopit/yandex-kassa/api/info"

const (
	TYPE_NOTIFICATION                 = "notification"
	EVENT_PAYMENT_WAITING_FOR_CAPTURE = "payment.waiting_for_capture"
	EVENT_PAYMENT_SUCCEEDED           = "payment.succeeded"
	EVENT_PAYMENT_CANCELED            = "payment.canceled"
	EVENT_REFUND_SUCCEEDED            = "refund.succeeded"
)

type Hook struct {
	Type   string        `json:"type" bson:"type"`
	Event  string        `json:"event" bson:"event"`
	Object *info.Payment `json:"object" bson:"object"`
}
