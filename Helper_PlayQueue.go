package voxvoid

func PlayQueueGetNumberOfItems() (uint64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueGetNumberOfItems
	response, err := VzxControllerApiCall(request)
	return response.Integer_1, err
}

func PlayQueueClear() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueClear
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueShuffle() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueShuffle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueSetToPack(packHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueSetToPack
	request.String_1 = packHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueAddPack(packHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueAddPack
	request.String_1 = packHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueSetToVisual(packHandle, visualHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueSetToVisual
	request.String_1 = packHandle
	request.String_2 = visualHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueAddVisual(packHandle, visualHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueAddVisual
	request.String_1 = packHandle
	request.String_2 = visualHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueSetToPreset(packHandle, visualHandle, presetHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueSetToPreset
	request.String_1 = packHandle
	request.String_2 = visualHandle
	request.String_3 = presetHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueAddPreset(packHandle, visualHandle, presetHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueAddPreset
	request.String_1 = packHandle
	request.String_2 = visualHandle
	request.String_3 = presetHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueAddPlaylist(playlistHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueAddPlaylist
	request.String_1 = playlistHandle
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueSetToPlaylist(playlistHandle string) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueSetToPlaylist
	request.String_1 = playlistHandle
	_, err = VzxControllerApiCall(request)
	return err
}
