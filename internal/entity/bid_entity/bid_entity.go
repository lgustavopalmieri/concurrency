package bid_entity

import (
	"context"
	"time"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/internal_error"
)

type Bid struct {
	BidId     string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
