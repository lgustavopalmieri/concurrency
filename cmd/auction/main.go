package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/configuration/database/mongodb"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/api/web/controller/auction_controller"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/api/web/controller/bid_controller"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/api/web/controller/user_controller"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/database/auction"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/database/bid"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/infra/database/user"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/usecase/auction_usecase"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/usecase/bid_usecase"
	"github.com/lgustavopalmieri/labs-go-expert-auctiont/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	usersController, bidsController, auctionsController := initDependencies(databaseConnection)

	router.POST("/auctions", auctionsController.CreateAuction)
	router.GET("/auctions", auctionsController.FindAuctions)
	router.GET("/auctions/:auctionId", auctionsController.FindAuctionById)
	router.GET("/auction/winner/:auctionId", auctionsController.FindWinningBidByAuctionId)
	router.POST("/bid", bidsController.CreateBid)
	router.GET("/bid/:auctionId", bidsController.FindBidByAuctionId)
	router.GET("/user/:userId", usersController.FindUserById)

	router.Run(":8080")
}

func initDependencies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController,
) {
	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))
	return
}
