package models

import "time"

type (
	SubscriptionInfo struct {
		SubscriptionID   string `json:"subscriptionId"`
		SubscriptionName string `json:"subscriptionName"`
		OfferID          string `json:"offerId"`
		PlanID           string `json:"planId"`
		IsTest           bool   `json:"isTest"`
		IsFreeTrial      bool   `json:"isFreeTrial"`
		SeatQuantity     int    `json:"seatQuantity"`
		Status           int    `json:"status"`
		Term             struct {
			TermUnit  string    `json:"termUnit"`
			StartDate time.Time `json:"startDate"`
			EndDate   time.Time `json:"endDate"`
		} `json:"term"`
		Beneficiary struct {
			UserID      string `json:"userId"`
			UserEmail   string `json:"userEmail"`
			AadObjectID string `json:"aadObjectId"`
			AadTenantID string `json:"aadTenantId"`
		} `json:"beneficiary"`
		Purchaser struct {
			UserID      string `json:"userId"`
			UserEmail   string `json:"userEmail"`
			AadObjectID string `json:"aadObjectId"`
			AadTenantID string `json:"aadTenantId"`
		} `json:"purchaser"`
	}
)
