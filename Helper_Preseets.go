package voxvoid

func PresetLoadForPack(packHandle string) (Response, error) {
	var request Request
	request.Action_t = RequestAction_PresetLoadForPackAndVisual
	request.String_1 = packHandle
	return VzxControllerApiCall(request)
}

func PresetLoadForPackAndVisual(packHandle, visualHandle string) (Response, error) {
	var request Request
	request.Action_t = RequestAction_PresetLoadForPackAndVisual
	request.String_1 = packHandle
	request.String_2 = visualHandle
	return VzxControllerApiCall(request)
}

func PresetGetNumberOfPresetsForVisual(packHandle, visualHandle string) (uint64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PresetGetNumberOfPresetsForVisual
	request.String_1 = packHandle
	request.String_2 = visualHandle
	response, err := VzxControllerApiCall(request)
	return response.Integer_1, err
}

func PresetGetPresetHandleByVisualAndIndex(packHandle, visualHandle string, index uint64) (Response, error) {
	var request Request
	request.Action_t = RequestAction_PresetGetPresetHandleByVisualAndIndex
	request.String_1 = packHandle
	request.String_2 = visualHandle
	request.Integer_1 = index
	return VzxControllerApiCall(request)
}
