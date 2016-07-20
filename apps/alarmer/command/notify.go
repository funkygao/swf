// +build darwin
package command

import (
	"github.com/ciarand/notify"
)

func displayNotify(title, txt string, snd *string) {
	var sound string
	if snd == nil {
		sound = "Glass"
	} else {
		sound = *snd
	}
	n := notify.NewNotificationWithSound(title, txt, sound)
	n.Display()
}
