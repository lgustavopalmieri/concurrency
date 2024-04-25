package bid_entity

import "time"

type Bid struct {
	BidId     string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}
