package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatar is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatar = errors.New("chat: Unable to get an avatar  URL")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatar
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatar
	}
	return url, nil
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrNoAvatar
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
			return "/avatars/" + file.Name(), nil
		}
	}

	return "", ErrNoAvatar
}
