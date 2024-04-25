package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/logger"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/entity/auction_entity"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_auction_id": id}

	var auctionEntityMongo AuctionEntityMongo
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by id %s", id), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find auction by id %s", id))
	}
	return &auction_entity.Auction{
		AuctionId:   auctionEntityMongo.AuctionId,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}
	if status != 0 {
		filter["status"] = status
	}
	if category != "" {
		filter["category"] = category
	}
	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find auctions")
	}
	defer cursor.Close(ctx)

	var auctionsEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionsEntityMongo); err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find auctions")
	}

	var auctionsEntity []auction_entity.Auction
	for _, auctionMongo := range auctionsEntityMongo {
		auctionsEntity = append(auctionsEntity, auction_entity.Auction{
			AuctionId:   auctionMongo.AuctionId,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}
	return auctionsEntity, nil
}
