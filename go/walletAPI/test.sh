curl --request POST \
  --url http://127.0.0.1:8080/wallets \
  --header 'Content-Type: application/json' \
  --data '{
	"name" : "d2"
}'

curl --request GET \
  --url http://127.0.0.1:8080/d2


curl --request POST \
  --url http://127.0.0.1:8080/d2/coins \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Go Coin",
    "symbol": "GOC",
    "amount": 2.1,
    "rate": 21.4
}'

curl --request GET \
  --url http://127.0.0.1:8080/d2



curl --request PUT \
  --url http://127.0.0.1:8080/d2/GOC \
  --header 'Content-Type: application/json' \
  --data '{
						"name": "Bitcoin",
            "symbol": "BTC",
            "amount": 2.15
}'

curl --request POST \
  --url http://127.0.0.1:8080/d2/coins \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Go Coin",
    "symbol": "GOC",
    "amount": 2.1,
    "rate": 21.4
}'


curl --request GET \
  --url http://127.0.0.1:8080/d2



curl --request DELETE \
  --url http://127.0.0.1:8080/d2/GOC

curl --request GET \
  --url http://127.0.0.1:8080/d2

sleep 120s