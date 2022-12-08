package models

import "time"

type (
	/*
		When the SaaS subscription is in Subscribed status:
			ChangePlan
			ChangeQuantity
			Renew (notify only, no ACK needed)
			Suspend (notify only, no ACK needed)
			Unsubscribe (notify only, no ACK needed)
		When SaaS subscription is in Suspended status:
			Reinstate
			Unsubscribe (notify only, no ACK needed)
	*/
	isWebhook_Webhook interface {
		isWebhook_Webhook()
	}
	WebhookMetadata struct {
		ID             string    `json:"id"`
		ActivityID     string    `json:"activityId"`
		SubscriptionID string    `json:"subscriptionId"`
		TimeStamp      time.Time `json:"timeStamp"`
		Action         string    `json:"action"`
	}

	Webhook struct {
		Metadata *WebhookMetadata  `json:"metadata"`
		Webhook  isWebhook_Webhook `json:"webhook"`
	}
	// https://learn.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#resolve-a-purchased-subscription
	ResolvedSubsription struct {
		ID               string `json:"id"`
		SubscriptionName string `json:"subscriptionName"`
		OfferID          string `json:"offerId"`
		PlanID           string `json:"planId"`
		Subscription     struct {
			ID                     string `json:"id"`
			PublisherID            string `json:"publisherId"`
			OfferID                string `json:"offerId"`
			Name                   string `json:"name"`
			SaasSubscriptionStatus string `json:"saasSubscriptionStatus"`
			Beneficiary            struct {
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
			PlanID string `json:"planId"`
			Term   struct {
				TermUnit string `json:"termUnit"`
			} `json:"term"`
			AutoRenew                 bool      `json:"autoRenew"`
			IsTest                    bool      `json:"isTest"`
			IsFreeTrial               bool      `json:"isFreeTrial"`
			AllowedCustomerOperations []string  `json:"allowedCustomerOperations"`
			SandboxType               string    `json:"sandboxType"`
			Created                   time.Time `json:"created"`
			LastModified              string    `json:"lastModified"`
			SessionMode               string    `json:"sessionMode"`
		} `json:"subscription"`
	}

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

	ChangePlanWebhook struct {
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
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			PublisherID string      `json:"publisherId"`
			OfferID     string      `json:"offerId"`
			PlanID      string      `json:"planId"`
			Quantity    interface{} `json:"quantity"`
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
				StartDate      time.Time   `json:"startDate"`
				EndDate        time.Time   `json:"endDate"`
				TermUnit       string      `json:"termUnit"`
				ChargeDuration interface{} `json:"chargeDuration"`
			} `json:"term"`
			AutoRenew    bool      `json:"autoRenew"`
			Created      time.Time `json:"created"`
			LastModified time.Time `json:"lastModified"`
		} `json:"subscription"`
		PurchaseToken interface{} `json:"purchaseToken"`
	}
	Webhook_ChangePlanWebhook struct {
		ChangePlanWebhook *ChangePlanWebhook `json:"changePlan"`
	}
	ChangeQuantityWebhook struct {
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
	Webhook_ChangeQuantityWebhook struct {
		ChangeQuantityWebhook *ChangeQuantityWebhook `json:"changeQuantity"`
	}
	ReinstateWebhook struct {
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
	Webhook_ReinstateWebhook struct {
		ReinstateWebhook *ReinstateWebhook `json:"reinstate"`
	}
	RenewWebhook struct {
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
	Webhook_RenewWebhook struct {
		RenewWebhook *RenewWebhook `json:"renew"`
	}
	SuspendWebhook struct {
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
	Webhook_SuspendWebhook struct {
		SuspendWebhook *SuspendWebhook `json:"suspend"`
	}
	UnsubscribeWebhook struct {
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
	Webhook_UnsubscribeWebhook struct {
		UnsubscribeWebhook *UnsubscribeWebhook `json:"unsubscribe"`
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

func (*Webhook_UnsubscribeWebhook) isWebhook_Webhook()    {}
func (*Webhook_ReinstateWebhook) isWebhook_Webhook()      {}
func (*Webhook_ChangePlanWebhook) isWebhook_Webhook()     {}
func (*Webhook_ChangeQuantityWebhook) isWebhook_Webhook() {}
func (*Webhook_SuspendWebhook) isWebhook_Webhook()        {}
func (*Webhook_RenewWebhook) isWebhook_Webhook()          {}
func (m *Webhook) GetWebhook() isWebhook_Webhook {
	if m != nil {
		return m.Webhook
	}
	return nil
}
func (x *Webhook) SetUnsubscribeWebhook(webhook *UnsubscribeWebhook) {
	x.Webhook = &Webhook_UnsubscribeWebhook{UnsubscribeWebhook: webhook}
}
func (x *Webhook) GetUnsubscribeWebhook() *UnsubscribeWebhook {
	if x, ok := x.GetWebhook().(*Webhook_UnsubscribeWebhook); ok {
		return x.UnsubscribeWebhook
	}
	return nil
}
func (x *Webhook) SetReinstateWebhook(webhook *ReinstateWebhook) {
	x.Webhook = &Webhook_ReinstateWebhook{ReinstateWebhook: webhook}
}
func (x *Webhook) GetReinstateWebhook() *ReinstateWebhook {
	if x, ok := x.GetWebhook().(*Webhook_ReinstateWebhook); ok {
		return x.ReinstateWebhook
	}
	return nil
}
func (x *Webhook) SetChangePlanWebhook(webhook *ChangePlanWebhook) {
	x.Webhook = &Webhook_ChangePlanWebhook{ChangePlanWebhook: webhook}
}
func (x *Webhook) GetChangePlanWebhook() *ChangePlanWebhook {
	if x, ok := x.GetWebhook().(*Webhook_ChangePlanWebhook); ok {
		return x.ChangePlanWebhook
	}
	return nil
}
func (x *Webhook) SetChangeQuantityWebhook(webhook *ChangeQuantityWebhook) {
	x.Webhook = &Webhook_ChangeQuantityWebhook{ChangeQuantityWebhook: webhook}
}
func (x *Webhook) GetChangeQuantityWebhook() *ChangeQuantityWebhook {
	if x, ok := x.GetWebhook().(*Webhook_ChangeQuantityWebhook); ok {
		return x.ChangeQuantityWebhook
	}
	return nil
}
func (x *Webhook) SetSuspendWebhook(webhook *SuspendWebhook) {
	x.Webhook = &Webhook_SuspendWebhook{SuspendWebhook: webhook}
}
func (x *Webhook) GetSuspendWebhook() *SuspendWebhook {
	if x, ok := x.GetWebhook().(*Webhook_SuspendWebhook); ok {
		return x.SuspendWebhook
	}
	return nil
}
func (x *Webhook) SetRenewWebhook(webhook *RenewWebhook) {
	x.Webhook = &Webhook_RenewWebhook{RenewWebhook: webhook}
}
func (x *Webhook) GetRenewWebhook() *RenewWebhook {
	if x, ok := x.GetWebhook().(*Webhook_RenewWebhook); ok {
		return x.RenewWebhook
	}
	return nil
}
