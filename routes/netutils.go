package routes

import (
	"context"
	"dandelion_api/pkg/utils"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func genRandomKey(client *redis.Client) (string, error) {
    newKey, genErr := client.RandomKey(context.TODO()).Result()
    if genErr != nil {
        return "", genErr
    }

    return newKey, nil
}

func UtilsHandler(w http.ResponseWriter, r *http.Request) {
    client := redis.NewClient(&redis.Options{
        Addr: "1.178.38.70:6379",
        Password: "Dandelion_4",
        DB: 0,
    })

    RanKey, RandomErr := genRandomKey(client)
    if RandomErr != nil {
        utils.NetThrowError(w, "no hay pues", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(RanKey))
    return
}








