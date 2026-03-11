package subscribes

import (
	"fmt"
)

type List struct {
	subscribes []Subscribe
}

func NewList() *List {
	return &List{
		subscribes: []Subscribe{},
	}
}

func (l *List) CreateNoteSubscribe(subscribe Subscribe) error {
	l.subscribes = append(l.subscribes, subscribe)
	fmt.Println("1111")

	return nil
}

func (l *List) GetSubscribes() []Subscribe {
	return l.subscribes
}

func (l *List) SortServiceNameSubscribe(serviceName string) ([]Subscribe, int) {
	arrSortServiceNameSubscribe := []Subscribe{}
	priceAllSubscribes := 0
	subscribes := l.subscribes

	for i := 0; i < len(subscribes); i++ {
		if subscribes[i].ServiceName == serviceName {
			arrSortServiceNameSubscribe = append(arrSortServiceNameSubscribe, subscribes[i])
			priceAllSubscribes += subscribes[i].Price
		}
	}

	return arrSortServiceNameSubscribe, priceAllSubscribes
}

func (l *List) SortServiceUserIdSubscribe(userId string) ([]Subscribe, int) {
	arrSortServiceUserIdSubscribe := []Subscribe{}
	priceAllSubscribes := 0
	subscribes := l.subscribes

	for i := 0; i < len(subscribes); i++ {
		if subscribes[i].UserId.String() == userId {
			arrSortServiceUserIdSubscribe = append(arrSortServiceUserIdSubscribe, subscribes[i])
			priceAllSubscribes += subscribes[i].Price
		}
	}

	return arrSortServiceUserIdSubscribe, priceAllSubscribes
}

// func (l *List) PriceAllSubscribes(serviceName string, userId string) int {
// 	priceAllSubscribes := 0
// 	subscribes := l.subscribes

// 	for i := 0; i < len(subscribes); i++ {
// 		if subscribes[i].ServiceName == serviceName {
// 			priceAllSubscribes += subscribes[i].Price
// 		} else if subscribes[i].UserId.String() == userId {
// 			priceAllSubscribes += subscribes[i].Price
// 		}
// 	}

// 	return priceAllSubscribes
// }
