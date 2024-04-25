package bid_usecase

import (
	"time"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/bid_entity"
)

type BidOutputDTO struct {
	BidId     string    `json:"bid_id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	bidRepository bid_entity.BidEntityRepositoryInterface
}
