package subscriptions

import (
	"encoding/json"
	"gitlab.com/chertokdmitry/surfproto/src/domain/subscriptions"
	"gitlab.com/chertokdmitry/surfproto/src/services"
	"log"
)

func GetByChatId(chatId int64) []subscriptions.Sub {
	subs, err := services.SubscriptionService.GetByChatId(chatId)

	if err != nil {
		log.Println(err)
	}

	return subs
}

func Create(requestSub string) {
	var sub subscriptions.Sub

	err := json.Unmarshal([]byte(requestSub), &sub)

	if err != nil {
		log.Println(err)
	}

	saveErr := services.SubscriptionService.Insert(&sub)

	if saveErr != nil {
		log.Println(saveErr)
	}
}

func Delete(subId int64) {
	if err := services.SubscriptionService.Delete(subId); err != nil {
		log.Println(err)
	}
}

