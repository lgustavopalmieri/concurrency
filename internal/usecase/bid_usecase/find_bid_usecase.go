package bid_usecase

import (
	"context"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/internal_error"
)

func (bu *BidUseCase) FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidList, err := bu.bidRepository.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	var bidOutputList []BidOutputDTO

	for _, bid := range bidList {
		bidOutputList = append(bidOutputList, BidOutputDTO{
			BidId:     bid.BidId,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}
	return bidOutputList, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError) {
	bidEntity, err := bu.bidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	bidOutput := &BidOutputDTO{
		BidId:     bidEntity.BidId,
		UserId:    bidEntity.UserId,
		AuctionId: bidEntity.AuctionId,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}

	return bidOutput, nil
}
