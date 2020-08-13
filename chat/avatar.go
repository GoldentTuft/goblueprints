package main

import (
	"errors"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです。
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar はユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した倍にはエラーを返します。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar は認証サーバーからプロフィール画像を得ます
type AuthAvatar struct{}

// UseAuthAvatar は認証サーバーからプロフィール画像を得ます
var UseAuthAvatar AuthAvatar

// GetAvatarURL は認証サーバーからプロフィール画像を得ます
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar はGravatarからプロフィール画像を得ます
type GravatarAvatar struct{}

// UseGravatar はGravatarからプロフィール画像を得ます
var UseGravatar GravatarAvatar

// GetAvatarURL はGravatarからプロフィール画像を得ます
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// FileSystemAvatar はアップロードされたファイルからプロフィール画像を得ます
type FileSystemAvatar struct{}

// UseFileSystemAvatar はアップロードされたファイルからプロフィール画像を得ます
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL はアップロードされたファイルからプロフィール画像を得ます
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "/avatars/" + useridStr + ".jpg", nil
		}
	}
	return "", ErrNoAvatarURL
}
