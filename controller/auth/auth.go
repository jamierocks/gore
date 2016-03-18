package auth

import (
    "bytes"
    "net/url"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/jamierocks/gore/modules"
    "github.com/gin-gonic/gin"
)

type GitHubResponse struct {
    AccessToken string `json:"access_token"`
}

func GetLogin(ctx *gin.Context) {
    if isAuthenticated(ctx) {
        ctx.Redirect(308, "/")
    } else {
        ctx.Redirect(307, "https://github.com/login/oauth/authorize?scope=user:email&client_id=" +
            modules.CONFIG.Section("auth").Key("CLIENT_ID").String())
    }
}

func GetCallback(ctx *gin.Context) {
    code := ctx.Query("code")

    data := url.Values{}
    data.Set("client_id", modules.CONFIG.Section("auth").Key("CLIENT_ID").String())
    data.Set("client_secret", modules.CONFIG.Section("auth").Key("CLIENT_SECRET").String())
    data.Set("code", code)

    client := &http.Client{}
    r, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(data.Encode()))
    r.Header.Set("Accept", "json")

    resp, _ := client.Do(r)
    body, _ := ioutil.ReadAll(resp.Body)

    var res GitHubResponse
    json.Unmarshal(body, &res)
    ctx.SetCookie("access_token", res.AccessToken, 999, "/",
        modules.CONFIG.Section("web").Key("DOMAIN").String(), true, true)

    ctx.JSON(200, res)
}

func isAuthenticated(ctx *gin.Context) bool {
    _, err := ctx.Cookie("access_token")
    return err == nil
}
