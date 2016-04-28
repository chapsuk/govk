# Vkontakte API server client

[![Build Status](https://travis-ci.org/chapsuk/govk.svg)](https://travis-ci.org/chapsuk/govk)

## API Methods

* [authorization](https://new.vk.com/dev/secure_how_to)
* [orders.get](https://new.vk.com/dev/orders.get)

## Run

```
go install gthub.com/chapsuk/govk/cmd/govk
``` 
## Example

orders.get

```
 $ govk -c CLIENT_ID -s CLIENT_SECRET -i 20 -t 1
2016/04/28 14:58:15
Gotten access_token: ACCESS_TOKEN
2016/04/28 14:58:15
Result orders.get method
{ID:957030 AppOrderID:0 Status:charged UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:5 Date:1448969223}
{ID:957012 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_three#45189849 Amount:0 Date:1448967342}
{ID:957007 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:0 Date:1448967017}
{ID:957006 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_two#45189849 Amount:0 Date:1448966974}
{ID:957005 AppOrderID:0 Status:declined UserID:1000 ReceiverID:1000 Item:coin_one#45189849 Amount:0 Date:1448966971}
{ID:948881 AppOrderID:0 Status:charged UserID:10872 ReceiverID:10872 Item:coins_30#9256941 Amount:3 Date:1447853021}
{ID:838963 AppOrderID:0 Status:charged UserID:87715133 ReceiverID:87715133 Item:premium_7#80706626 Amount:25 Date:1434620746}
{ID:838959 AppOrderID:0 Status:charged UserID:87715133 ReceiverID:87715133 Item:premium_7#80706626 Amount:25 Date:1434620662}
{ID:767480 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1426177226}
{ID:761985 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:like_1#149855 Amount:5 Date:1425558347}
{ID:761089 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1425470860}
{ID:760405 AppOrderID:0 Status:charged UserID:4718705 ReceiverID:4718705 Item:premium_7#149855 Amount:25 Date:1425394741}
{ID:378144 AppOrderID:0 Status:charged UserID:8182 ReceiverID:8182 Item:premium_1#4528709 Amount:10 Date:1386677299}
{ID:41867 AppOrderID:0 Status:charged UserID:5499135 ReceiverID:5499135 Item:premium_7d Amount:10 Date:1351602715}
```