package controllers

import (
	"fakewebserver/fakewebserver/src/database"
	"fakewebserver/fakewebserver/src/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	userName := c.PostForm("login")
	userPassword := c.PostForm("password")
	user, ok := database.MyDatabase.Users[userName]
	if ok {
		if user.UserPassword == userPassword {
			c.Header("Cache-Control", "no-cache")
			c.Header("Content-Security-Policy", "default-src 'none'; base-uri 'self'; block-all-mixed-content; child-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/; connect-src 'self' uploads.github.com objects-origin.githubusercontent.com www.githubstatus.com collector.githubapp.com api.github.com github-cloud.s3.amazonaws.com github-production-repository-file-5c1aeb.s3.amazonaws.com github-production-upload-manifest-file-7fdce7.s3.amazonaws.com github-production-user-asset-6210df.s3.amazonaws.com cdn.optimizely.com logx.optimizely.com/v1/events translator.github.com wss://alive.github.com; font-src github.githubassets.com; form-action 'self' github.com gist.github.com objects-origin.githubusercontent.com; frame-ancestors 'none'; frame-src render.githubusercontent.com viewscreen.githubusercontent.com notebooks.githubusercontent.com; img-src 'self' data: github.githubassets.com identicons.github.com collector.githubapp.com github-cloud.s3.amazonaws.com secured-user-images.githubusercontent.com/ *.githubusercontent.com; manifest-src 'self'; media-src github.com user-images.githubusercontent.com/; script-src github.githubassets.com; style-src 'unsafe-inline' github.githubassets.com; worker-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/")
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.Header("Expect-Ct", "max-age=2592000, report-uri=\"https://api.github.com/_private/browser/errors\"")
			c.Header("Location", "https://github.com/login")
			c.Header("Permissions-Policy", "interest-cohort=()")
			c.Header("Referrer-Policy", "origin-when-cross-origin, strict-origin-when-cross-origin")
			c.Header("Server", "GitHub.com")
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "has_recent_activity",
				Value:    "1",
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 2,
			})
			userSession := util.GetSession()
			session.Set("usersession", userSession)
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "user_session",
				Value:    userSession,
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 2,
			})
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "__Host-user_session_same_site",
				Value:    userSession,
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 3,
			})
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "tz",
				Value:    "Asia%2FShanghai",
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 2,
			})
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "dotcom_user",
				Value:    userName,
				Domain:   ".github.com",
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 2,
			})
			jumpSession := util.GetSession()
			session.Set("jumpsession", jumpSession)
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "_gh_sess",
				Value:    jumpSession,
				Path:     "/",
				Secure:   true,
				HttpOnly: true,
				SameSite: 2,
			})
			session.Save()
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubdomains; preload")
			c.String(302, "")
		}
	}
	//TODO 没找到用户的情况
}

func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	cookie := c.PostForm("cookie")
	userSessionIndex := strings.Index(cookie, "user_session=")
	userSession := cookie[userSessionIndex+13 : userSessionIndex+13+10]
	jumpSessionIndex := strings.Index(cookie, "_gh_sess=")
	jumpSession := cookie[jumpSessionIndex+9 : jumpSessionIndex+9+10]
	if session.Get("usersession") == userSession && session.Get("jumpsession") == jumpSession {
		c.Header("Cache-Control", "no-cache")
		c.Header("Content-Security-Policy", "default-src 'none'; base-uri 'self'; block-all-mixed-content; child-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/; connect-src 'self' uploads.github.com objects-origin.githubusercontent.com www.githubstatus.com collector.githubapp.com api.github.com github-cloud.s3.amazonaws.com github-production-repository-file-5c1aeb.s3.amazonaws.com github-production-upload-manifest-file-7fdce7.s3.amazonaws.com github-production-user-asset-6210df.s3.amazonaws.com cdn.optimizely.com logx.optimizely.com/v1/events translator.github.com wss://alive.github.com; font-src github.githubassets.com; form-action 'self' github.com gist.github.com objects-origin.githubusercontent.com; frame-ancestors 'none'; frame-src render.githubusercontent.com viewscreen.githubusercontent.com notebooks.githubusercontent.com; img-src 'self' data: github.githubassets.com identicons.github.com collector.githubapp.com github-cloud.s3.amazonaws.com secured-user-images.githubusercontent.com/ *.githubusercontent.com; manifest-src 'self'; media-src github.com user-images.githubusercontent.com/; script-src github.githubassets.com; style-src 'unsafe-inline' github.githubassets.com; worker-src github.com/assets-cdn/worker/ gist.github.com/assets-cdn/worker/")
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Header("Expect-Ct", "max-age=2592000, report-uri=\"https://api.github.com/_private/browser/errors\"")
		c.Header("Location", "https://github.com/")
		c.Header("Permissions-Policy", "interest-cohort=()")
		c.Header("Referrer-Policy", "origin-when-cross-origin, strict-origin-when-cross-origin")
		c.Header("Server", "Github.com")
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "has_recent_activity",
			Value:    "1",
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: 2,
		})
		jumpSession := util.GetSession()
		session.Set("jumpsession", jumpSession)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "_gh_sess",
			Value:    jumpSession,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			SameSite: 2,
		})
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubdomains; preload")
		session.Save()
		c.String(302, "")
	}
}

func GithubHandler(c *gin.Context) {

}
