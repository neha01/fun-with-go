package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFT struct {
	ID       string `json:"id"`
	NFTType  string `json:"nftType"`
	BuyPrice int    `json:"buyPrice"` // in dollars
}

var nfts = []NFT{
	{
		ID:       "1",
		NFTType:  "ERC1155",
		BuyPrice: 30,
	},
	{
		ID:       "2",
		NFTType:  "ERC721",
		BuyPrice: 28,
	},
}

func getNFTs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nfts)
}

func main() {
	router := gin.Default()
	router.GET("/nfts", getNFTs)
	router.Run("localhost:8080")
}
