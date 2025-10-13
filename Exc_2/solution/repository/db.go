package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data		DONE
	// drinks := ...
	drinks := []model.Drink{
		{ID: 1, Name: "CocaCola", Price: 2.5, Description: "Soda with caffeine"},
		{ID: 2, Name: "Water", Price: 1.0, Description: "still Water"},
		{ID: 3, Name: "Beer", Price: 3.0, Description: "Alcohol"},
	}

	// Init orders slice with some test data	DONE
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now().Add(-10 * time.Minute), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now().Add(-5 * time.Minute), Amount: 1},
		{DrinkID: 1, CreatedAt: time.Now().Add(-2 * time.Minute), Amount: 3},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)

	for i := 0; i < len(db.drinks); i++ {
		var amount uint64 = 0
		for j := 0; j < len(db.orders); j++ {
			if j == i {
				amount += db.orders[j].Amount
			}
		}
		totalledOrders[db.drinks[i].ID] = uint64(amount)
	}
	// calculate total orders		DONE
	// key = DrinkID, value = Amount of orders		DONE
	// totalledOrders map[uint64]uint64		DONE
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	db.orders = append(db.orders, *order)
	// add order to db.orders slice
}
