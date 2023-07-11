#include <alsa/asoundlib.h>
#include <stdbool.h>

int open_mixer(snd_mixer_t **handle, const char *card) {
    int err;
    if ((err = snd_mixer_open(handle, 0)) < 0) {
        return err;
    }
    if ((err = snd_mixer_attach(*handle, card)) < 0) {
        return err;
    }
    if ((err = snd_mixer_selem_register(*handle, NULL, NULL)) < 0) {
        return err;
    }
    if ((err = snd_mixer_load(*handle)) < 0) {
        return err;
    }
    return 0;
}

snd_mixer_elem_t* get_elem(snd_mixer_t *handle, const char *selem_name) {
    snd_mixer_selem_id_t *sid;
    snd_mixer_selem_id_alloca(&sid);
    snd_mixer_selem_id_set_index(sid, 0);
    snd_mixer_selem_id_set_name(sid, selem_name);
    return snd_mixer_find_selem(handle, sid);
}

long set_volume(int volume) {
    snd_mixer_t *handle;
    const char *card = "default";
    const char *selem_name = "Master";

    int err;
    if (err = open_mixer(&handle, card) < 0) {
        return err;
    }

    snd_mixer_elem_t* elem = get_elem(handle, selem_name);
    if (!elem) {
        return -1;
    }

    long minv, maxv;
    snd_mixer_selem_get_playback_volume_range(elem, &minv, &maxv);
    long volume_to_set = (volume * maxv + 50) / 100;
    snd_mixer_selem_set_playback_volume_all(elem, volume_to_set);

    snd_mixer_close(handle);

    return volume_to_set;
}

long get_volume() {
    snd_mixer_t *handle;
    const char *card = "default";
    const char *selem_name = "Master";

    int err;
    if (err = open_mixer(&handle, card) < 0) {
        return err;
    }

    snd_mixer_elem_t* elem = get_elem(handle, selem_name);
    if (!elem) {
        return -1;
    }

    long minv, maxv, volume;
    snd_mixer_selem_get_playback_volume_range(elem, &minv, &maxv);
    snd_mixer_selem_get_playback_volume(elem, SND_MIXER_SCHN_MONO, &volume);

    snd_mixer_close(handle);

    return (volume * 100 + maxv / 2) / maxv;
}

long get_mute(){
    snd_mixer_t *handle;
    const char *card = "default";
    const char *selem_name = "Master";

    int err;
    if (err = open_mixer(&handle, card) < 0) {
        return err;
    }

    snd_mixer_elem_t* elem = get_elem(handle, selem_name);
    if (!elem) {
        return -1;
    }

    int current_mute;
    snd_mixer_selem_get_playback_switch(elem, SND_MIXER_SCHN_MONO, &current_mute);

    return current_mute;
}

long set_mute(bool  mute) {
    snd_mixer_t *handle;
    const char *card = "default";
    const char *selem_name = "Master";

    int err;
    if (err = open_mixer(&handle, card) < 0) {
        return err;
    }


    snd_mixer_elem_t* elem = get_elem(handle, selem_name);
    if (!elem) {
        return err;
    }
    if (mute) {
        err = snd_mixer_selem_set_playback_switch_all(elem, 0);
    } else {
        err = snd_mixer_selem_set_playback_switch_all(elem, 1);
    }
    if (err < 0) {
        return err;
    }

    snd_mixer_close(handle);

    return 0;
}