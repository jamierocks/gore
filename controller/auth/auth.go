package auth

import (
    "bytes"
    "net/url"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/jamierocks/gore/models"
    "github.com/jamierocks/gore/modules"
    "github.com/google/go-github/github"
    "gopkg.in/macaron.v1"
)

type AccessTokenResponse struct {
    AccessToken string `json:"access_token"`
}

func GetLogin(ctx *macaron.Context) {
    if isAuthenticated(ctx) {
        data := url.Values{}
        data.Set("access_token", ctx.GetCookie("access_token"))

        client := &http.Client{}
        r, _ := http.NewRequest("POST", "https://api.github.com/user", bytes.NewBufferString(data.Encode()))
        r.Header.Set("Accept", "application/json")

        resp, _ := client.Do(r)
        body, _ := ioutil.ReadAll(resp.Body)

        var res github.User
        json.Unmarshal(body, &res)

        if !models.DoesUserExist(*res.Login) {
            modules.DB.Create(&models.User{
                Username: *res.Login,
                FullName: *res.Name,
                Type: 0,
            })
        }

        ctx.Redirect("/", 308)
    } else {
        ctx.Redirect("https://github.com/login/oauth/authorize?scope=user:email&client_id=" +
            modules.CONFIG.Section("auth").Key("CLIENT_ID").String(), 307)
    }
}

func GetCallback(ctx *macaron.Context) {
    code := ctx.Query("code")

    data := url.Values{}
    data.Set("client_id", modules.CONFIG.Section("auth").Key("CLIENT_ID").String())
    data.Set("client_secret", modules.CONFIG.Section("auth").Key("CLIENT_SECRET").String())
    data.Set("code", code)

    client := &http.Client{}
    r, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(data.Encode()))
    r.Header.Set("Accept", "application/json")

    resp, _ := client.Do(r)
    body, _ := ioutil.ReadAll(resp.Body)

    var res AccessTokenResponse
    json.Unmarshal(body, &res)
    ctx.SetCookie("access_token", res.AccessToken)

    ctx.Redirect("/login", 308)
}

func isAuthenticated(ctx *macaron.Context) bool {
    return ctx.GetCookie("access_token") != ""
}
