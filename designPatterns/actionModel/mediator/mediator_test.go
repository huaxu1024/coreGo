package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := GetMediatorInstance()
	mediator.CD = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}

	//Tiggle
	mediator.CD.ReadData()

	assert.Equal(t, "music,image", mediator.CD.Data)
	assert.Equal(t, "music", mediator.CPU.Sound)
	assert.Equal(t, "image", mediator.CPU.Video)
	assert.Equal(t, "image", mediator.Video.Data)
	assert.Equal(t, "music", mediator.Sound.Data)
}