package domain

const (
    TestByteString = "{\n  \"status\": true,\n  \"data\": {\n    \"sms\": [\n      [\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"mms\": [\n      [\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ],\n      [\n        {\n          \"country\": \"Canada\",\n          \"bandwidth\": \"12\",\n          \"response_time\": \"67\",\n          \"provider\": \"Rond\"\n        },\n        {\n          \"country\": \"Great Britain\",\n          \"bandwidth\": \"98\",\n          \"response_time\": \"593\",\n          \"provider\": \"Kildy\"\n        },\n        {\n          \"country\": \"Russian Federation\",\n          \"bandwidth\": \"77\",\n          \"response_time\": \"1734\",\n          \"provider\": \"Topolo\"\n        }\n      ]\n    ],\n    \"voice_call\": [\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"TransparentCalls\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      },\n      {\n        \"country\": \"US\",\n        \"bandwidth\": \"53\",\n        \"response_time\": \"321\",\n        \"provider\": \"E-Voice\",\n        \"connection_stability\": 0.72,\n        \"ttfb\": 442,\n        \"voice_purity\": 20,\n        \"median_of_call_time\": 5\n      }\n    ],\n    \"email\": [\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 195\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ],\n      [\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        },\n        {\n          \"country\": \"RU\",\n          \"provider\": \"Gmail\",\n          \"delivery_time\": 393\n        }\n      ]\n    ],\n    \"billing\": {\n      \"create_customer\": true,\n      \"purchase\": true,\n      \"payout\": true,\n      \"recurring\": false,\n      \"fraud_control\": true,\n      \"checkout_page\": false\n    },\n    \"support\": [\n      3,\n      62\n    ],\n    \"incident\": [\n      {\"topic\":  \"Topic 1\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 2\", \"status\": \"active\"},\n      {\"topic\":  \"Topic 3\", \"status\": \"closed\"},\n      {\"topic\":  \"Topic 4\", \"status\": \"closed\"}\n    ]\n  },\n  \"error\": \"\"\n}"
)

const MinResponseTime = 30
const MaxResponseTime = 2000

const MinConnectionStability = 600
const MaxConnectionStability = 1000

const MinVoicePurity = 0
const MaxVoicePurity = 92

const MinVoiceCallMedian = 3
const MaxVoiceCallMedian = 60

const MinTTFB = 2
const MaxTTFB = 980

const MinBandwidth = 0
const MaxBandwidth = 100

const MinEmailDeliveryTime = 0
const MaxEmailDeliveryTime = 600

const SmsFilename = "sms.data"
const MmsApiUrl = "http://localhost:8282/mms" // to params
const VoiceFilename = "voice.data"
const EmailFilename = "email.data"
const BillingFilename = "billing.data"
const SupportApiUrl = "http://localhost:8282/support"
const AccendentListFilename = "accendents.data"