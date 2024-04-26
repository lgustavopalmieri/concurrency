package bid

import (
	"context"
	"sync"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/logger"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/auction_entity"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/bid_entity"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/database/auction"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	BidId     string  `bson:"_bid_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository{
	return &BidRepository{
		Collection: database.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)
		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)
			if err != nil {
				logger.Error("Error trying to find auction by id for bid", err)
				return 
			}

			if auctionEntity.Status != auction_entity.Active{
				return
			}

			bidEntityMongo := &BidEntityMongo{
				BidId: bidValue.BidId,
				UserId: bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount: bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert bid", err)
				return 
			}
		}(bid)
	}

	wg.Wait()
	return nil
}
