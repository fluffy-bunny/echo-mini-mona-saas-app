package models

import "time"

type (
	UnsubscribeSubscriptionInfo struct {
		SubscriptionID   string `json:"subscriptionId"`
		SubscriptionName string `json:"subscriptionName"`
		OfferID          string `json:"offerId"`
		PlanID           string `json:"planId"`
		IsTest           bool   `json:"isTest"`
		IsFreeTrial      bool   `json:"isFreeTrial"`
		SeatQuantity     int    `json:"seatQuantity"`
		Status           string `json:"status"`
		Beneficiary      struct {
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
	UnsubscribeInfo struct {
		OperationId  string                      `json:"operationId"`
		Subscription UnsubscribeSubscriptionInfo `json:"subscription"`
	}
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
	ChangePlan struct {
		ID                     string    `json:"id"`
		ActivityID             string    `json:"activityId"`
		OperationRequestSource string    `json:"operationRequestSource"`
		SubscriptionID         string    `json:"subscriptionId"`
		TimeStamp              time.Time `json:"timeStamp"`
		Action                 string    `json:"action"`
	}
	ChangeQuantity struct {
		ID                     string    `json:"id"`
		ActivityID             string    `json:"activityId"`
		PublisherID            string    `json:"publisherId"`
		OfferID                string    `json:"offerId"`
		PlanID                 string    `json:"planId"`
		Quantity               int       `json:"quantity"`
		SubscriptionID         string    `json:"subscriptionId"`
		TimeStamp              time.Time `json:"timeStamp"`
		Action                 string    `json:"action"`
		Status                 string    `json:"status"`
		OperationRequestSource string    `json:"operationRequestSource"`
		Subscription           struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			PublisherID string `json:"publisherId"`
			OfferID     string `json:"offerId"`
			PlanID      string `json:"planId"`
			Quantity    int    `json:"quantity"`
			Beneficiary struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"beneficiary"`
			Purchaser struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"purchaser"`
			AllowedCustomerOperations []string `json:"allowedCustomerOperations"`
			SessionMode               string   `json:"sessionMode"`
			IsFreeTrial               bool     `json:"isFreeTrial"`
			IsTest                    bool     `json:"isTest"`
			SandboxType               string   `json:"sandboxType"`
			SaasSubscriptionStatus    string   `json:"saasSubscriptionStatus"`
			Term                      struct {
				StartDate      time.Time `json:"startDate"`
				EndDate        time.Time `json:"endDate"`
				TermUnit       string    `json:"termUnit"`
				ChargeDuration string    `json:"chargeDuration"`
			} `json:"term"`
			AutoRenew    bool      `json:"autoRenew"`
			Created      time.Time `json:"created"`
			LastModified time.Time `json:"lastModified"`
		} `json:"subscription"`
		PurchaseToken string `json:"purchaseToken"`
	}
	Reinstate struct {
		ID             string    `json:"id"`
		ActivityID     string    `json:"activityId"`
		SubscriptionID string    `json:"subscriptionId"`
		PublisherID    string    `json:"publisherId"`
		OfferID        string    `json:"offerId"`
		PlanID         string    `json:"planId"`
		Quantity       int       `json:"quantity"`
		TimeStamp      time.Time `json:"timeStamp"`
		Action         string    `json:"action"`
		Status         string    `json:"status"`
	}
	Renew struct {
		ID                     string    `json:"id"`
		ActivityID             string    `json:"activityId"`
		PublisherID            string    `json:"publisherId"`
		OfferID                string    `json:"offerId"`
		PlanID                 string    `json:"planId"`
		Quantity               int       `json:"quantity"`
		SubscriptionID         string    `json:"subscriptionId"`
		TimeStamp              time.Time `json:"timeStamp"`
		Action                 string    `json:"action"`
		Status                 string    `json:"status"`
		OperationRequestSource string    `json:"operationRequestSource"`
		Subscription           struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			PublisherID string `json:"publisherId"`
			OfferID     string `json:"offerId"`
			PlanID      string `json:"planId"`
			Quantity    int    `json:"quantity"`
			Beneficiary struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"beneficiary"`
			Purchaser struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"purchaser"`
			AllowedCustomerOperations []string `json:"allowedCustomerOperations"`
			SessionMode               string   `json:"sessionMode"`
			IsFreeTrial               bool     `json:"isFreeTrial"`
			IsTest                    bool     `json:"isTest"`
			SandboxType               string   `json:"sandboxType"`
			SaasSubscriptionStatus    string   `json:"saasSubscriptionStatus"`
			Term                      struct {
				StartDate      time.Time `json:"startDate"`
				EndDate        time.Time `json:"endDate"`
				TermUnit       string    `json:"termUnit"`
				ChargeDuration string    `json:"chargeDuration"`
			} `json:"term"`
			AutoRenew    bool      `json:"autoRenew"`
			Created      time.Time `json:"created"`
			LastModified time.Time `json:"lastModified"`
		} `json:"subscription"`
		PurchaseToken string `json:"purchaseToken"`
	}
	Suspend struct {
		ID                     string    `json:"id"`
		ActivityID             string    `json:"activityId"`
		PublisherID            string    `json:"publisherId"`
		OfferID                string    `json:"offerId"`
		PlanID                 string    `json:"planId"`
		Quantity               int       `json:"quantity"`
		SubscriptionID         string    `json:"subscriptionId"`
		TimeStamp              time.Time `json:"timeStamp"`
		Action                 string    `json:"action"`
		Status                 string    `json:"status"`
		OperationRequestSource string    `json:"operationRequestSource"`
		Subscription           struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			PublisherID string `json:"publisherId"`
			OfferID     string `json:"offerId"`
			PlanID      string `json:"planId"`
			Quantity    int    `json:"quantity"`
			Beneficiary struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"beneficiary"`
			Purchaser struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"purchaser"`
			AllowedCustomerOperations []string `json:"allowedCustomerOperations"`
			SessionMode               string   `json:"sessionMode"`
			IsFreeTrial               bool     `json:"isFreeTrial"`
			IsTest                    bool     `json:"isTest"`
			SandboxType               string   `json:"sandboxType"`
			SaasSubscriptionStatus    string   `json:"saasSubscriptionStatus"`
			Term                      struct {
				StartDate      time.Time `json:"startDate"`
				EndDate        time.Time `json:"endDate"`
				TermUnit       string    `json:"termUnit"`
				ChargeDuration string    `json:"chargeDuration"`
			} `json:"term"`
			AutoRenew    bool      `json:"autoRenew"`
			Created      time.Time `json:"created"`
			LastModified time.Time `json:"lastModified"`
		} `json:"subscription"`
		PurchaseToken string `json:"purchaseToken"`
	}
	Unsubscribe struct {
		ID                     string    `json:"id"`
		ActivityID             string    `json:"activityId"`
		PublisherID            string    `json:"publisherId"`
		OfferID                string    `json:"offerId"`
		PlanID                 string    `json:"planId"`
		Quantity               int       `json:"quantity"`
		SubscriptionID         string    `json:"subscriptionId"`
		TimeStamp              time.Time `json:"timeStamp"`
		Action                 string    `json:"action"`
		Status                 string    `json:"status"`
		OperationRequestSource string    `json:"operationRequestSource"`
		Subscription           struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			PublisherID string `json:"publisherId"`
			OfferID     string `json:"offerId"`
			PlanID      string `json:"planId"`
			Quantity    int    `json:"quantity"`
			Beneficiary struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"beneficiary"`
			Purchaser struct {
				EmailID  string `json:"emailId"`
				ObjectID string `json:"objectId"`
				TenantID string `json:"tenantId"`
				Puid     string `json:"puid"`
			} `json:"purchaser"`
			AllowedCustomerOperations []string `json:"allowedCustomerOperations"`
			SessionMode               string   `json:"sessionMode"`
			IsFreeTrial               bool     `json:"isFreeTrial"`
			IsTest                    bool     `json:"isTest"`
			SandboxType               string   `json:"sandboxType"`
			SaasSubscriptionStatus    string   `json:"saasSubscriptionStatus"`
			Term                      struct {
				StartDate      time.Time `json:"startDate"`
				EndDate        time.Time `json:"endDate"`
				TermUnit       string    `json:"termUnit"`
				ChargeDuration string    `json:"chargeDuration"`
			} `json:"term"`
			AutoRenew    bool      `json:"autoRenew"`
			Created      time.Time `json:"created"`
			LastModified time.Time `json:"lastModified"`
		} `json:"subscription"`
		PurchaseToken string `json:"purchaseToken"`
	}
	Subscription struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		PublisherID string `json:"publisherId"`
		OfferID     string `json:"offerId"`
		PlanID      string `json:"planId"`
		Quantity    int    `json:"quantity"`
		Beneficiary struct {
			EmailID  string `json:"emailId"`
			ObjectID string `json:"objectId"`
			TenantID string `json:"tenantId"`
			Puid     string `json:"puid"`
		} `json:"beneficiary"`
		Purchaser struct {
			EmailID  string `json:"emailId"`
			ObjectID string `json:"objectId"`
			TenantID string `json:"tenantId"`
			Puid     string `json:"puid"`
		} `json:"purchaser"`
		AllowedCustomerOperations []string  `json:"allowedCustomerOperations"`
		SessionMode               string    `json:"sessionMode"`
		IsFreeTrial               bool      `json:"isFreeTrial"`
		AutoRenew                 bool      `json:"autoRenew"`
		IsTest                    bool      `json:"isTest"`
		SandboxType               string    `json:"sandboxType"`
		Created                   time.Time `json:"created"`
		LastModified              string    `json:"lastModified"`
		SaasSubscriptionStatus    string    `json:"saasSubscriptionStatus"`
		Term                      struct {
			StartDate time.Time `json:"startDate"`
			EndDate   time.Time `json:"endDate"`
			TermUnit  string    `json:"termUnit"`
		} `json:"term"`
	}
	Subscriptions struct {
		Subscriptions []Subscription `json:"subscriptions"`
	}
	// https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#post-httpsmarketplaceapimicrosoftcomapisaassubscriptionssubscriptionidactivateapi-versionapiversion
	ActivateRequest struct {
		PlanID   string  `json:"planId"`
		Quantity *string `json:"quantity"`
	}
)
