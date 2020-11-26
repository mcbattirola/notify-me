# Notify-me
 
Make sending push notifications to you and your team easy, by sending a simple HTTP Request.
 
To achieve this, we have a server and a mobile app (Android only for now). They are independents and the app can be replaced by your own. As long as it gets the push notification permission token and sends it to the server through a push notification.
 
On the server side, you can register a sender. Each sender will have its IP persisted. When someone sends a push notification, we send a FCM push notification request to all users with a subscription to that sender.
 
Registered devices can subscribe to a sender via a POST request and be notified every time a sender pushes a notification.
 
This project's main objective is to learn Golang and make something useful along the way, it is just a POC.
It is *not* production ready. Above all, it lacks any kind of authentication and authorization.