package gerencianet

import "strconv"

type endpoints struct {
	requester interface {
		request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error)
	}
}

func (endpoints endpoints) CreateCharge(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/charge", "POST", nil, body)
}

func (endpoints endpoints) DetailCharge(chargeID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id", "GET", params, nil)
}

func (endpoints endpoints) UpdateChargeMetadata(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/metadata", "PUT", params, body)
}

func (endpoints endpoints) UpdateBillet(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/billet", "PUT", params, body)
}

func (endpoints endpoints) PayCharge(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/pay", "POST", params, body)
}

func (endpoints endpoints) CancelCharge(chargeID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/cancel", "PUT", params, nil)
}

func (endpoints endpoints) CreateCarnet(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/carnet", "POST", nil, body)
}

func (endpoints endpoints) DetailCarnet(carnetID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(carnetID) }
	return endpoints.requester.request("/carnet/:id", "GET", params, nil)
}

func (endpoints endpoints) UpdateParcel(carnetID int, parcel int, body map[string]interface{}) (string, error) {
	params := map[string]string{ 
		"id": strconv.Itoa(carnetID),
		"parcel": strconv.Itoa(parcel),
	}
	return endpoints.requester.request("/carnet/:id/parcel/:parcel", "PUT", params, body)
}

func (endpoints endpoints) UpdateCarnetMetadata(carnetID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(carnetID) }
	return endpoints.requester.request("/carnet/:id/metadata", "PUT", params, body)
}

func (endpoints endpoints) GetNotification(token string) (string, error) {
	params := map[string]string{ "token": token }
	return endpoints.requester.request("/notification/:token", "GET", params, nil)
}

func (endpoints endpoints) GetPlans(limit int, offset int) (string, error) {
	params := map[string]string{ 
		"limit": strconv.Itoa(limit),
		"offset": strconv.Itoa(offset),
	}
	return endpoints.requester.request("/plans", "GET", params, nil)
}

func (endpoints endpoints) CreatePlan(body map[string]interface{}) (string, error) {
	return endpoints.requester.request("/plan", "POST", nil, body)
}

func (endpoints endpoints) DeletePlan(planID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(planID) }
	return endpoints.requester.request("/plan/:id", "DELETE", params, nil)
}

func (endpoints endpoints) CreateSubscription(planID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(planID) }
	return endpoints.requester.request("/plan/:id/subscription", "POST", params, body)
}

func (endpoints endpoints) DetailSubscription(subscriptionID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(subscriptionID) }
	return endpoints.requester.request("/subscription/:id", "GET", params, nil)
}

func (endpoints endpoints) PaySubscription(subscriptionID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(subscriptionID) }
	return endpoints.requester.request("/subscription/:id/pay", "POST", params, body)
}

func (endpoints endpoints) CancelSubscription(subscriptionID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(subscriptionID) }
	return endpoints.requester.request("/subscription/:id/cancel", "PUT", params, nil)
}

func (endpoints endpoints) UpdateSubscriptionMetadata(subscriptionID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(subscriptionID) }
	return endpoints.requester.request("/subscription/:id/metadata", "PUT", params, body)
}

func (endpoints endpoints) GetInstallments(total int, brand string) (string, error) {
	params := map[string]string{ 
		"total": strconv.Itoa(total),
		"brand": brand,
	}
	return endpoints.requester.request("/installments", "GET", params, nil)
}

func (endpoints endpoints) ResendBillet(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/billet/resend", "POST", params, body)
}

func (endpoints endpoints) CreateChargeHistory(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/history", "POST", params, body)
}

func (endpoints endpoints) ResendCarnet(carnetID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(carnetID) }
	return endpoints.requester.request("/carnet/:id/resend", "POST", params, body)
}

func (endpoints endpoints) ResendParcel(carnetID int, parcel int, body map[string]interface{}) (string, error) {
	params := map[string]string{ 
		"id": strconv.Itoa(carnetID),
		"parcel": strconv.Itoa(parcel),
	}
	return endpoints.requester.request("/carnet/:id/parcel/:parcel/resend", "POST", params, body)
}

func (endpoints endpoints) CreateCarnetHistory(carnetID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(carnetID) }
	return endpoints.requester.request("/carnet/:id/history", "POST", params, body)
}

func (endpoints endpoints) CancelCarnet(carnetID int) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(carnetID) }
	return endpoints.requester.request("/carnet/:id/cancel", "PUT", params, nil)
}

func (endpoints endpoints) CancelParcel(carnetID int, parcel int) (string, error) {
	params := map[string]string{ 
		"id": strconv.Itoa(carnetID),
		"parcel": strconv.Itoa(parcel),
	}
	return endpoints.requester.request("/carnet/:id/parcel/:parcel/cancel", "PUT", params, nil)
}

func (endpoints endpoints) ChargeLink(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/link", "POST", params, body)
}

func (endpoints endpoints) UpdateChargeLink(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/link", "PUT", params, body)
}

func (endpoints endpoints) CreateSubscriptionHistory(subscriptionID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(subscriptionID) }
	return endpoints.requester.request("/subscription/:id/history", "POST", params, body)
}

func (endpoints endpoints) UpdatePlan(planID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(planID) }
	return endpoints.requester.request("/plan/:id", "PUT", params, body)
}

func (endpoints endpoints) CreateChargeBalanceSheet(chargeID int, body map[string]interface{}) (string, error) {
	params := map[string]string{ "id": strconv.Itoa(chargeID) }
	return endpoints.requester.request("/charge/:id/balance-sheet", "POST", params, body)
}