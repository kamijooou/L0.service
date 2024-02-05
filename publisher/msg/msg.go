package msg

var Fake string

const Message0 = `
{
	"order_uid": "b563feb7b2b84b6test",
	"track_number": "WBILMTESTTRACK",
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
}
`

const Message1 = `
{
	"order_uid": "a1b2c3d4e5f6",
	"track_number": "WBILMTESTTRACK1",
	"entry": "WBIL",
	"delivery": {
	  "name": "John Doe",
	  "phone": "+9720000000",
	  "zip": "123456",
	  "city": "Example City",
	  "address": "123 Main Street",
	  "region": "Example Region",
	  "email": "john.doe@example.com"
	},
	"payment": {
	  "transaction": "a1b2c3d4e5f6",
	  "request_id": "req123",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1500,
	  "payment_dt": 1643467200,
	  "bank": "alpha",
	  "delivery_cost": 1200,
	  "goods_total": 300,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 123456,
		"track_number": "WBILMTESTTRACK1",
		"price": 150,
		"rid": "abcdef123456",
		"name": "Example Item",
		"sale": 10,
		"size": "M",
		"total_price": 300,
		"nm_id": 789012,
		"brand": "Example Brand",
		"status": 200
	  }
	],
	"locale": "en",
	"internal_signature": "sig123",
	"customer_id": "john_doe",
	"delivery_service": "example_service",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2022-01-30T12:00:00Z",
	"oof_shard": "1"
}
`

const Message2 = `
{
	"order_uid": "b2c3d4e5f6a1",
	"track_number": "WBILMTESTTRACK2",
	"entry": "WBIL",
	"delivery": {
	  "name": "Jane Doe",
	  "phone": "+9720000000",
	  "zip": "654321",
	  "city": "Another City",
	  "address": "456 Oak Street",
	  "region": "Another Region",
	  "email": "jane.doe@example.com"
	},
	"payment": {
	  "transaction": "b2c3d4e5f6a1",
	  "request_id": "req456",
	  "currency": "EUR",
	  "provider": "wbeuro",
	  "amount": 1800,
	  "payment_dt": 1643467200,
	  "bank": "beta",
	  "delivery_cost": 1500,
	  "goods_total": 300,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 789012,
		"track_number": "WBILMTESTTRACK2",
		"price": 150,
		"rid": "123456abcdef",
		"name": "Another Item",
		"sale": 15,
		"size": "L",
		"total_price": 300,
		"nm_id": 123456,
		"brand": "Another Brand",
		"status": 201
	  }
	],
	"locale": "fr",
	"internal_signature": "sig456",
	"customer_id": "jane_doe",
	"delivery_service": "another_service",
	"shardkey": "8",
	"sm_id": 88,
	"date_created": "2022-01-31T14:30:00Z",
	"oof_shard": "2"
}
`

const Message3 = `
{
	"order_uid": "c3d4e5f6a1b2",
	"track_number": "WBILMTESTTRACK3",
	"entry": "WBIL",
	"delivery": {
	  "name": "Alice Johnson",
	  "phone": "+9720000000",
	  "zip": "543210",
	  "city": "Wonderland",
	  "address": "123 Rabbit Hole",
	  "region": "Fictional Region",
	  "email": "alice@example.com"
	},
	"payment": {
	  "transaction": "c3d4e5f6a1b2",
	  "request_id": "req789",
	  "currency": "GBP",
	  "provider": "wbpound",
	  "amount": 2000,
	  "payment_dt": 1643467200,
	  "bank": "gamma",
	  "delivery_cost": 1800,
	  "goods_total": 200,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 987654,
		"track_number": "WBILMTESTTRACK3",
		"price": 100,
		"rid": "abcdef987654",
		"name": "Wonderful Item",
		"sale": 20,
		"size": "S",
		"total_price": 200,
		"nm_id": 654321,
		"brand": "Wonder Brand",
		"status": 203
	  }
	],
	"locale": "es",
	"internal_signature": "sig789",
	"customer_id": "alice",
	"delivery_service": "wonder_service",
	"shardkey": "7",
	"sm_id": 77,
	"date_created": "2022-02-01T16:45:00Z",
	"oof_shard": "3"
}
`

const Message4 = `
{
	"order_uid": "d4e5f6a1b2c3",
	"track_number": "WBILMTESTTRACK4",
	"entry": "WBIL",
	"delivery": {
	  "name": "Bob Smith",
	  "phone": "+9720000000",
	  "zip": "987654",
	  "city": "Exampleville",
	  "address": "789 Main Street",
	  "region": "Test Region",
	  "email": "bob@example.com"
	},
	"payment": {
	  "transaction": "d4e5f6a1b2c3",
	  "request_id": "req012",
	  "currency": "CAD",
	  "provider": "wbdollar",
	  "amount": 2200,
	  "payment_dt": 1643467200,
	  "bank": "delta",
	  "delivery_cost": 2000,
	  "goods_total": 200,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 789012,
		"track_number": "WBILMTESTTRACK4",
		"price": 100,
		"rid": "abcdef789012",
		"name": "Test Item",
		"sale": 25,
		"size": "XL",
		"total_price": 200,
		"nm_id": 123789,
		"brand": "Test Brand",
		"status": 204
	  }
	],
	"locale": "de",
	"internal_signature": "sig012",
	"customer_id": "bob",
	"delivery_service": "test_service",
	"shardkey": "6",
	"sm_id": 66,
	"date_created": "2022-02-02T09:30:00Z",
	"oof_shard": "4"
}
`

const Message5 = `
{
	"order_uid": "e5f6a1b2c3d4",
	"track_number": "WBILMTESTTRACK5",
	"entry": "WBIL",
	"delivery": {
	  "name": "Eva Green",
	  "phone": "+9720000000",
	  "zip": "135790",
	  "city": "Green City",
	  "address": "42 Green Street",
	  "region": "Green Region",
	  "email": "eva.green@example.com"
	},
	"payment": {
	  "transaction": "e5f6a1b2c3d4",
	  "request_id": "req135",
	  "currency": "EUR",
	  "provider": "wbeuro",
	  "amount": 2500,
	  "payment_dt": 1643467200,
	  "bank": "epsilon",
	  "delivery_cost": 2200,
	  "goods_total": 300,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 135792,
		"track_number": "WBILMTESTTRACK5",
		"price": 200,
		"rid": "abcdef135792",
		"name": "Green Item",
		"sale": 30,
		"size": "XS",
		"total_price": 300,
		"nm_id": 246810,
		"brand": "Green Brand",
		"status": 205
	  }
	],
	"locale": "it",
	"internal_signature": "sig135",
	"customer_id": "eva_green",
	"delivery_service": "green_service",
	"shardkey": "5",
	"sm_id": 55,
	"date_created": "2022-02-03T14:15:00Z",
	"oof_shard": "5"
}
`

const Message6 = `
{
	"order_uid": "f6a1b2c3d4e5",
	"track_number": "WBILMTESTTRACK6",
	"entry": "WBIL",
	"delivery": {
	  "name": "Frank Black",
	  "phone": "+9720000000",
	  "zip": "111111",
	  "city": "Black City",
	  "address": "1 Dark Street",
	  "region": "Dark Region",
	  "email": "frank.black@example.com"
	},
	"payment": {
	  "transaction": "f6a1b2c3d4e5",
	  "request_id": "req777",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 2700,
	  "payment_dt": 1643467200,
	  "bank": "phi",
	  "delivery_cost": 2400,
	  "goods_total": 300,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 777888,
		"track_number": "WBILMTESTTRACK6",
		"price": 200,
		"rid": "abcdef777888",
		"name": "Black Item",
		"sale": 35,
		"size": "XXL",
		"total_price": 300,
		"nm_id": 888999,
		"brand": "Black Brand",
		"status": 206
	  }
	],
	"locale": "es",
	"internal_signature": "sig777",
	"customer_id": "frank_black",
	"delivery_service": "black_service",
	"shardkey": "4",
	"sm_id": 44,
	"date_created": "2022-02-04T16:00:00Z",
	"oof_shard": "6"
}
`
const Message7 = `
{
	"order_uid": "g1b2c3d4e5f6a",
	"track_number": "WBILMTESTTRACK7",
	"entry": "WBIL",
	"delivery": {
	  "name": "Grace Taylor",
	  "phone": "+9720000000",
	  "zip": "222222",
	  "city": "Graceful City",
	  "address": "7 Serene Street",
	  "region": "Calm Region",
	  "email": "grace.taylor@example.com"
	},
	"payment": {
	  "transaction": "g1b2c3d4e5f6a",
	  "request_id": "req111",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 3000,
	  "payment_dt": 1643467200,
	  "bank": "gamma",
	  "delivery_cost": 2700,
	  "goods_total": 300,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 111222,
		"track_number": "WBILMTESTTRACK7",
		"price": 150,
		"rid": "abcdef111222",
		"name": "Graceful Item 1",
		"sale": 40,
		"size": "S",
		"total_price": 150,
		"nm_id": 222333,
		"brand": "Graceful Brand",
		"status": 207
	  },
	  {
		"chrt_id": 333444,
		"track_number": "WBILMTESTTRACK7",
		"price": 150,
		"rid": "abcdef333444",
		"name": "Graceful Item 2",
		"sale": 40,
		"size": "M",
		"total_price": 150,
		"nm_id": 444555,
		"brand": "Graceful Brand",
		"status": 208
	  }
	],
	"locale": "fr",
	"internal_signature": "sig111",
	"customer_id": "grace_taylor",
	"delivery_service": "graceful_service",
	"shardkey": "3",
	"sm_id": 33,
	"date_created": "2022-02-05T10:30:00Z",
	"oof_shard": "7"
}
`

const Message8 = `
{
	"order_uid": "h2c3d4e5f6a1b",
	"track_number": "WBILMTESTTRACK8",
	"entry": "WBIL",
	"delivery": {
	  "name": "Harry Brown",
	  "phone": "+9720000000",
	  "zip": "333333",
	  "city": "Magical City",
	  "address": "8 Wizard Street",
	  "region": "Enchanted Region",
	  "email": "harry.brown@example.com"
	},
	"payment": {
	  "transaction": "h2c3d4e5f6a1b",
	  "request_id": "req222",
	  "currency": "EUR",
	  "provider": "wbeuro",
	  "amount": 3200,
	  "payment_dt": 1643467200,
	  "bank": "phi",
	  "delivery_cost": 2900,
	  "goods_total": 400,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 555666,
		"track_number": "WBILMTESTTRACK8",
		"price": 200,
		"rid": "abcdef555666",
		"name": "Magical Item 1",
		"sale": 45,
		"size": "L",
		"total_price": 200,
		"nm_id": 666777,
		"brand": "Magical Brand",
		"status": 209
	  },
	  {
		"chrt_id": 777888,
		"track_number": "WBILMTESTTRACK8",
		"price": 200,
		"rid": "abcdef777888",
		"name": "Magical Item 2",
		"sale": 45,
		"size": "XL",
		"total_price": 200,
		"nm_id": 888999,
		"brand": "Magical Brand",
		"status": 210
	  }
	],
	"locale": "de",
	"internal_signature": "sig222",
	"customer_id": "harry_brown",
	"delivery_service": "magical_service",
	"shardkey": "2",
	"sm_id": 22,
	"date_created": "2022-02-06T13:45:00Z",
	"oof_shard": "8"
}
`

const Message9 = `
{
	"order_uid": "i3d4e5f6a1b2c",
	"track_number": "WBILMTESTTRACK9",
	"entry": "WBIL",
	"delivery": {
	  "name": "Ivy White",
	  "phone": "+9720000000",
	  "zip": "444444",
	  "city": "Whiteland",
	  "address": "9 Snow Street",
	  "region": "Snowy Region",
	  "email": "ivy.white@example.com"
	},
	"payment": {
	  "transaction": "i3d4e5f6a1b2c",
	  "request_id": "req333",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 3500,
	  "payment_dt": 1643467200,
	  "bank": "gamma",
	  "delivery_cost": 3200,
	  "goods_total": 400,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 999000,
		"track_number": "WBILMTESTTRACK9",
		"price": 250,
		"rid": "abcdef999000",
		"name": "Snowy Item 1",
		"sale": 50,
		"size": "S",
		"total_price": 250,
		"nm_id": 111222,
		"brand": "Snowy Brand",
		"status": 211
	  },
	  {
		"chrt_id": 111222,
		"track_number": "WBILMTESTTRACK9",
		"price": 250,
		"rid": "abcdef111222",
		"name": "Snowy Item 2",
		"sale": 50,
		"size": "M",
		"total_price": 250,
		"nm_id": 222333,
		"brand": "Snowy Brand",
		"status": 212
	  }
	],
	"locale": "es",
	"internal_signature": "sig333",
	"customer_id": "ivy_white",
	"delivery_service": "snowy_service",
	"shardkey": "1",
	"sm_id": 11,
	"date_created": "2022-02-07T17:00:00Z",
	"oof_shard": "9"
}
`

const Message10 = `
{
	"order_uid": "j4e5f6a1b2c3d",
	"track_number": "WBILMTESTTRACK10",
	"entry": "WBIL",
	"delivery": {
	  "name": "Jack Black",
	  "phone": "+9720000000",
	  "zip": "555555",
	  "city": "Blacktown",
	  "address": "10 Dark Alley",
	  "region": "Dark Region",
	  "email": "jack.black@example.com"
	},
	"payment": {
	  "transaction": "j4e5f6a1b2c3d",
	  "request_id": "req444",
	  "currency": "GBP",
	  "provider": "wbpound",
	  "amount": 3800,
	  "payment_dt": 1643467200,
	  "bank": "phi",
	  "delivery_cost": 3500,
	  "goods_total": 500,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 333444,
		"track_number": "WBILMTESTTRACK10",
		"price": 300,
		"rid": "abcdef333444",
		"name": "Dark Item 1",
		"sale": 55,
		"size": "L",
		"total_price": 300,
		"nm_id": 444555,
		"brand": "Dark Brand",
		"status": 213
	  },
	  {
		"chrt_id": 555666,
		"track_number": "WBILMTESTTRACK10",
		"price": 300,
		"rid": "abcdef555666",
		"name": "Dark Item 2",
		"sale": 55,
		"size": "XL",
		"total_price": 300,
		"nm_id": 666777,
		"brand": "Dark Brand",
		"status": 214
	  }
	],
	"locale": "fr",
	"internal_signature": "sig444",
	"customer_id": "jack_black",
	"delivery_service": "dark_service",
	"shardkey": "10",
	"sm_id": 10,
	"date_created": "2022-02-08T10:15:00Z",
	"oof_shard": "10"
}
`
