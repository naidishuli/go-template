package common

import (
    "strings"

    "bets/internal/api/apierror"
    "bets/internal/app"
    "github.com/gofiber/fiber/v3"
)

type Middleware struct {
    log app.Logger
    jwt app.JWT

    baseRepo app.BaseRepo
}

func NewMiddleware(dep app.App) *Middleware {
    return &Middleware{
        log:      dep.Log(),
        jwt:      dep.Pkg().JWT,
        baseRepo: dep.Repository().Base,
    }
}

// Authorize validates the jwt token passed and inject the user data to the request context.
func (g *Middleware) Authorize(ctx fiber.Ctx) error {
    authHeader := ctx.Get("authorization")
    bearerParts := strings.Split(authHeader, " ")

    if len(bearerParts) < 2 {
        return apierror.Unauthorized(map[string]any{"reason": "no token present"})
    }

    //userClaims := new(dts.UserJWT)
    //err := g.jwt.ParseToken(bearerParts[1], userClaims)
    //if err != nil {
    //    return apierror.Unauthorized(map[string]any{"reason": "token not valid"})
    //}
    //
    //user := new(model.User)
    //err = g.baseRepo.First(&user, "uuid = ?", userClaims.UID)
    //if err != nil {
    //    return err
    //}
    //
    //appCtx := app.NewCtx(ctx.UserContext())
    //userCtx := appCtx.WithUser(user).Context()
    //
    //ctx.SetUserContext(userCtx)
    return ctx.Next()
}
