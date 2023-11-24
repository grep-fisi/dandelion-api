package routes

import (
	"context"
	"dandelion_api/pkg/utils"
	"io"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func sessionGetJSON(key string, _ string, client *redis.Client) (string, error) {
    val, getErr := client.Get(context.TODO(), key).Result()
    if getErr != nil {
        return "", getErr
    }

    return val, nil
}

func sessionSetJSON(key string, value string, client *redis.Client) (string, error) {
    val, setErr := client.Set(context.TODO(), key, value, 0).Result()
    if setErr != nil {
        return "", setErr
    }

    return val, nil
}

func deleteKey(key string, _ string, client *redis.Client) (string, error) {
    deleteErr := client.Unlink(context.TODO(), key).Err()
    if deleteErr != nil {
        return "", deleteErr
    }

    return "", nil
}

var methodHandlerMap = map[string]func(key string, value string ,client *redis.Client) (string, error) {
    http.MethodGet:     sessionGetJSON,
    http.MethodPost:    sessionSetJSON,
    http.MethodDelete:  deleteKey,
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
    client := redis.NewClient(&redis.Options{
        Addr: "1.178.38.70:6379",
        Password: "Dandelion_4",
        DB: 0,
    })

    if client == nil {
        utils.NetThrowError(w, "broken database connection", http.StatusInternalServerError)
        return
    }

    var params = r.URL.Query()
    var idParam string

    if !params.Has("id") {
        if r.Method == http.MethodGet {
            utils.NetThrowError(w, "no id was given", http.StatusBadRequest)
            return
        }

        randomKey, randomErr := genRandomKey(client)
        idParam = randomKey
        log.Println(randomKey)

        if randomErr != nil {
            log.Println(randomErr)
            utils.NetThrowError(w, "could not generate random key", http.StatusInternalServerError)
            return
        }
    } else if idParam = params.Get("id"); idParam == "" {
        log.Println("empty id param")
        utils.NetThrowError(w, "empty id", http.StatusBadRequest)
        return
    }

    bodyBytes, readErr := io.ReadAll(r.Body)
    if readErr != nil {
        log.Println(readErr)
        utils.NetThrowError(w, "invalid json schema as body", http.StatusBadRequest)
        return
    }

    var body string = string(bodyBytes)
    strResponse, respErr := methodHandlerMap[r.Method](idParam, body, client)
    if respErr == redis.Nil {
        log.Println(respErr)
        utils.NetThrowError(w, "key does not exist", http.StatusNotFound)
        return
    }

    w.Write([]byte(strResponse))
    return
}




