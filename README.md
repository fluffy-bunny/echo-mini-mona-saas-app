# echo-mini-mona-saas-app

An attempt to make a simple app that honors the Mona integration.  This app does nothing but receive the right webhooks, exposes landing pages, and make calls back to the marketplace when necessary.  

Its a fake app.  

## References

[mona-saas FAQ](https://github.com/microsoft/mona-saas/tree/main/docs)  
[puchase landing page FAQ](https://github.com/microsoft/mona-saas/tree/main/docs#can-i-retrieve-subscription-details-from-the-purchase-confirmation-page)  
[configuration landing page FAQ](https://github.com/microsoft/mona-saas/tree/main/docs#what-is-the-subscription-configuration-page)  
[pc-saas-fulfillment-subscription-api](https://docs.microsoft.com/en-us/azure/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api)  

## Credentials

I got these OAuth2 credentials from the mona application.  They need to be in your environment.  

```bash
"AZURE_AD__CREDENTIALS__CLIENT_ID":"97ce31f4-**REDACTED**",
"AZURE_AD__CREDENTIALS__CLIENT_SECRET":"ALJ8Q~**REDACTED**",
"AZURE_AD__TENANT":"5905aed6-**REDACTED**",
"MONA_STORAGE_ACCOUNT_NAME": "monastorage**REDACTED**",
```

## NGROK
I use [ngrok](https://ngrok.com/) to expose an public internet address that proxies to my localhost app.  

```bash
ngrok http 1323
```
When ngrok starts up it tells you what public urls you can now use.  
```
Account                       ghstahl@gmail.com (Plan: Free)  
Version                       2.3.40  
Region                        United States (us)  
Web Interface                 http://127.0.0.1:4040  
Forwarding                    http://74c5-104-173-245-55.ngrok.io -> http://localhost:1323  
Forwarding                    https://74c5-104-173-245-55.ngrok.io -> http://localhost:1323 
```

I use the ```https://74c5-104-173-245-55.ngrok.io``` url when populating the [mono configuation](https://mona-web-monamapped.azurewebsites.net/setup)  
I use the same endpoint when populating the LogicApps HTTP POST urls.    
My [mona resource group](https://portal.azure.com/#@mapped.com/resource/subscriptions/948d7b66-1e6f-4856-8b1d-41a767b5b4fa/resourceGroups/mona-mapped-marketplace/overview)  

In my [Cancel Subscription LogicApp](https://portal.azure.com/#@mapped.com/resource/subscriptions/948d7b66-1e6f-4856-8b1d-41a767b5b4fa/resourceGroups/mona-mapped-marketplace/providers/Microsoft.Logic/workflows/mona-on-subscription-canceled-monamapped/designer) I use the following;

```bash
https://74c5-104-173-245-55.ngrok.io/api/v1/saas/unsubscribe
```
