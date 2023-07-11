package volume

/*
#cgo LDFLAGS: -lasound
#include <alsa/asoundlib.h>
#include <stdbool.h>

extern long set_volume(snd_mixer_elem_t *elem, int volume);
extern long get_volume(snd_mixer_elem_t *elem);
extern long set_mute(snd_mixer_elem_t *elem,bool mute);
extern long get_mute(snd_mixer_elem_t *elem);
extern long open_mixer(snd_mixer_t **handle, char *name);
extern long close_mixer(snd_mixer_t *handle);
extern snd_mixer_elem_t *get_elem(snd_mixer_t *handle, char *name);
*/
import "C"
import (
	"fmt"
	"io"
)

var _ io.Closer = &Handler{}

type Handler struct {
	handle *C.snd_mixer_t
	elem   *C.snd_mixer_elem_t
}

func (v *Handler) Close() error {
	errCode := C.close_mixer(v.handle)
	if errCode < 0 {
		return fmt.Errorf("Error closing mixer: ALSA error code %d\n", int(errCode))
	}
	return nil
}

func NewVolumeHandler(mixerName string, elemName string) (*Handler, error) {
	var handle *C.snd_mixer_t // Open the mixer
	errCode := C.open_mixer(&handle, C.CString(mixerName))
	if errCode != 0 {
		return nil, fmt.Errorf("Error opening mixer: ALSA error code %d\n", int(errCode))
	}

	var elem *C.snd_mixer_elem_t
	// Get the element
	elem = C.get_elem(handle, C.CString(elemName))
	if elem == nil {
		return nil, fmt.Errorf("Error getting element: ALSA error code %d\n", int(errCode))
	}
	return &Handler{
		handle: handle,
		elem:   elem,
	}, nil

}

func (v *Handler) SetVolume(volume int) error {
	//vol := max(0, min(100, volume))             // Clamp the volume between 0 and 100
	errCode := C.set_volume(v.elem, C.int(volume)) //
	if errCode < 0 {
		return fmt.Errorf("Failed to set volume: ALSA error code %d\n", errCode)
	}

	return nil
}

func (v *Handler) GetVolume() (int, error) {
	volume := C.get_volume(v.elem)
	if volume < 0 {
		return 0, fmt.Errorf("Failed to get volume: ALSA error code %d\n", volume)
	}

	return int(volume), nil
}

func (v *Handler) SetMute(mute bool) error {
	errCode := C.set_mute(v.elem, C.bool(mute))
	if errCode < 0 {
		return fmt.Errorf("Failed to set mute: ALSA error code %d\n", int(errCode))
	}

	return nil
}

func (v *Handler) IsMuted() (bool, error) {
	mute := C.get_mute(v.elem)
	if mute < 0 {
		return false, fmt.Errorf("Failed to get mute: ALSA error code %d\n", int(mute))
	}

	return mute == 0, nil
}
