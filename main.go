package main

/*
#cgo LDFLAGS: -lasound
#include <alsa/asoundlib.h>
#include <stdbool.h>
extern long set_volume(int volume);
extern long get_volume();
extern long set_mute(bool mute);
extern long get_mute();

*/
import "C"
import "fmt"

func SetVolume(volume int) error {
	errCode := C.set_volume(C.int(volume))
	if errCode < 0 {
		return fmt.Errorf("Failed to set volume: ALSA error code %d\n", errCode)
	}

	return nil
}
func GetVolume() (int, error) {
	volume := int(C.get_volume())
	if volume < 0 {
		return 0, fmt.Errorf("Failed to get volume: ALSA error code %d\n", volume)
	}
	return volume, nil
}

func GetMute() (bool, error) {
	switch res := int(C.get_mute()); {
	case res < 0:
		return false, nil
	case res > 0:
		return true, nil
	default:
		return false, fmt.Errorf("Failed to get mute status: ALSA error code %d\n", res)
	}
}

func SetMute(mute bool) error {
	errCode := C.set_mute(C.bool(mute))
	if errCode < 0 {
		return fmt.Errorf("Failed to set mute: ALSA error code %d\n", int(errCode))
	}
	return nil
}

func main() {
	fmt.Println(GetVolume()) // Prints: 0
	SetVolume(100)
	fmt.Println(GetVolume()) // Prints: 0
	muted, err := GetMute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("muted", muted)
	SetMute(!muted)

}
