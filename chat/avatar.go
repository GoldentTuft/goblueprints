package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです。
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar はユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した倍にはエラーを返します。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	GetAvatarURL(u ChatUser) (string, error)
}

// TryAvatars は
type TryAvatars []Avatar

// GetAvatarURL はどれかからプロフィール画像を得ようとします
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar は認証サーバーからプロフィール画像を得ます
type AuthAvatar struct{}

// UseAuthAvatar は認証サーバーからプロフィール画像を得ます
var UseAuthAvatar AuthAvatar

// GetAvatarURL は認証サーバーからプロフィール画像を得ます
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar はGravatarからプロフィール画像を得ます
type GravatarAvatar struct{}

// UseGravatar はGravatarからプロフィール画像を得ます
var UseGravatar GravatarAvatar

// GetAvatarURL はGravatarからプロフィール画像を得ます
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	// なければデフォルト画像が返されるっぽい
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar はアップロードされたファイルからプロフィール画像を得ます
type FileSystemAvatar struct{}

// UseFileSystemAvatar はアップロードされたファイルからプロフィール画像を得ます
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL はアップロードされたファイルからプロフィール画像を得ます
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}

	return "", ErrNoAvatarURL
}
