package main

import (
	"net/http"

	"time"

	"github.com/labstack/echo/v4"
)

var wallets []Wallet

//var coins []Coin
type ResponseCoin struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Amount  float64 `json:"amount"`
	Rate    float64 `json:"rate"`
	Code    int     `json:"code"`
	Message string  `json:"message"`
}
type ResponseWalletGet struct {
	Size    int      `json:"size"`
	Wallets []Wallet `json:"wallets"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
}

type ResponseWallet struct {
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	Coins       []Coin    `json:"coins"`
	LastUpdated time.Time `json:"last_updated"`
	Code        int       `json:"code"`
	Message     string    `json:"message"`
}
type Wallet struct {
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	Coins       []Coin    `json:"coins"`
	LastUpdated time.Time `json:"last_updated"`
}

type Coin struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
}

type Name struct {
	Name string `json:"name"`
}

//remove by swapping element with last element and returning slice - 1 elements
func deleteWallet(index int) Wallet {
	wallets[index] = wallets[len(wallets)-1]
	wal := wallets[len(wallets)-1]
	wallets = wallets[:len(wallets)-1]
	return wal
}

func getWallet(name string) int {
	for key, val := range wallets {
		if name == val.Name {
			return key
		}
	}
	return -1
}

func getCoin(wallet Wallet, symbol string) int {
	for key, val := range wallet.Coins {
		if symbol == val.Symbol {
			return key
		}
	}
	return -1
}

func getBalance(index int) float64 {
	var sum float64 = 0
	for _, val := range wallets[index].Coins {
		sum += val.Rate * val.Amount
	}
	return sum
}

func deleteCoin(walletIndex int, coinIndex int) {
	wallets[walletIndex].Coins[coinIndex] = wallets[walletIndex].Coins[len(wallets[walletIndex].Coins)-1]
	wallets[walletIndex].Coins = wallets[walletIndex].Coins[:len(wallets[walletIndex].Coins)-1]
}
func main() {
	e := echo.New()

	// handle wallet creation
	e.POST("/wallets", func(c echo.Context) error {
		var wal Wallet
		if err := c.Bind(&wal); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if getWallet(wal.Name) != -1 {
			return c.JSON(http.StatusBadRequest, "wallet with such a name already exists")
		}
		wal.LastUpdated = time.Now()
		wallets = append(wallets, wal)
		resp := ResponseWallet{
			Name:        wal.Name,
			Balance:     wal.Balance,
			Coins:       wal.Coins,
			LastUpdated: wal.LastUpdated,
			Code:        200,
			Message:     "Wallet added successfully!",
		}
		return c.JSON(http.StatusOK, resp)
	})

	// get all wallets
	e.GET("/wallets", func(c echo.Context) error {
		resp := ResponseWalletGet{
			Size:    len(wallets),
			Wallets: wallets,
			Code:    200,
			Message: "All wallets received successfully!",
		}
		return c.JSON(http.StatusAccepted, resp)
	})

	e.PUT("/wallets/:wname", func(c echo.Context) error {
		name := c.Param("wname")
		index := getWallet(name)
		if index == -1 {
			return c.JSON(http.StatusBadRequest, "no such name")
		}
		//update wallet
		var wal Wallet
		if err := c.Bind(&wal); err != nil {
			return c.JSON(http.StatusBadRequest, "incorrect input format")
		}
		wallets[index].Name = wal.Name
		wallets[index].LastUpdated = time.Now()
		// return new wallet
		resp := ResponseWallet{
			Name:        wallets[index].Name,
			Balance:     wallets[index].Balance,
			Coins:       wallets[index].Coins,
			LastUpdated: wallets[index].LastUpdated,
			Code:        200,
			Message:     "Wallet name changed successfully!",
		}
		return c.JSON(http.StatusAccepted, resp)
	})

	e.DELETE("/wallets/:wname", func(c echo.Context) error {
		name := c.Param("wname")
		index := getWallet(name)
		if index == -1 {
			return c.JSON(http.StatusNotFound, "selected item does not exist!")
		}
		wal := deleteWallet(index)
		resp := ResponseWallet{
			Name:        wal.Name,
			Balance:     wal.Balance,
			Coins:       wal.Coins,
			LastUpdated: wal.LastUpdated,
			Code:        200,
			Message:     "Wallet deleted (logged out) successfully!",
		}
		return c.JSON(http.StatusAccepted, resp)

	})

	e.POST("/:wname/coins", func(c echo.Context) error {
		var coin Coin
		if err := c.Bind(&coin); err != nil {
			return c.JSON(http.StatusBadRequest, "incorect input format")
		}
		name := c.Param("wname")
		index := getWallet(name)
		if index == -1 {
			return c.JSON(http.StatusNotFound, "name does not exist!")
		}
		if getCoin(wallets[index], coin.Symbol) != -1 {
			return c.JSON(http.StatusBadRequest, "coin with such symbol already exists")
		}
		wallets[index].Coins = append(wallets[index].Coins, coin)
		resp := ResponseCoin{
			Name:    coin.Name,
			Symbol:  coin.Symbol,
			Amount:  coin.Amount,
			Rate:    coin.Rate,
			Code:    200,
			Message: "Coin added successfully!",
		}
		wallets[index].Balance = getBalance(index)
		wallets[index].LastUpdated = time.Now()
		return c.JSON(http.StatusAccepted, resp)
	})

	e.GET("/:wname", func(c echo.Context) error {
		name := c.Param("wname")
		index := getWallet(name)
		if index == -1 {
			return c.JSON(http.StatusNotFound, "name dosnt exist")
		}
		wal := wallets[index]
		resp := ResponseWallet{
			Name:        wal.Name,
			Balance:     wal.Balance,
			Coins:       wal.Coins,
			LastUpdated: wal.LastUpdated,
			Code:        200,
			Message:     "All coins received successfully!",
		}
		return c.JSON(http.StatusAccepted, resp)
	})

	e.PUT("/:wname/:symbol", func(c echo.Context) error {
		name := c.Param("wname")
		symbol := c.Param("symbol")
		index := getWallet(name)

		if index == -1 {
			return c.JSON(http.StatusNotFound, "name dosnt exist")
		}
		coinIndex := getCoin(wallets[index], symbol)

		if coinIndex == -1 {
			return c.JSON(http.StatusNotFound, "coin dosnt exist")
		}
		coin := wallets[index].Coins[coinIndex]
		if err := c.Bind(&coin); err != nil {
			return c.JSON(http.StatusBadRequest, "input is incorrect format")
		}
		wallets[index].Coins[coinIndex] = coin
		resp := ResponseCoin{
			Name:    coin.Name,
			Symbol:  coin.Symbol,
			Amount:  coin.Amount,
			Rate:    coin.Rate,
			Code:    200,
			Message: "Coin updated successfully!",
		}
		wallets[index].Balance = getBalance(index)
		wallets[index].LastUpdated = time.Now()
		return c.JSON(http.StatusAccepted, resp)
	})

	e.DELETE("/:wallet/:symbol", func(c echo.Context) error {
		name := c.Param("wname")
		symbol := c.Param("symbol")
		index := getWallet(name)

		if index == -1 {
			return c.JSON(http.StatusNotFound, "wallet dosnt exist")
		}
		coinIndex := getCoin(wallets[index], symbol)
		if coinIndex == -1 {
			return c.JSON(http.StatusNotFound, "coin dosnt exist")
		}
		coin := wallets[index].Coins[coinIndex]
		deleteCoin(index, coinIndex)
		wallets[index].Balance = getBalance(index)
		resp := ResponseCoin{
			Name:    coin.Name,
			Symbol:  coin.Symbol,
			Amount:  coin.Amount,
			Rate:    coin.Rate,
			Message: "Coin deleted successfully!",
		}
		return c.JSON(http.StatusAccepted, resp)
	})
	e.Start("0.0.0.0:8080")

}
