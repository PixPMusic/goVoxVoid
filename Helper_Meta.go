package voxvoid

func MetaGetNumberOfPacks() (uint64, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_MetaGetNumberOfPacks

	resp, err := VzxControllerApiCall(request)
	return resp.Integer_1, err
}

func MetaGetPackInfo(packIndex int) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_MetaGetPackInfo
	request.Integer_1 = uint64(packIndex)

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func MetaLoadVisualsForPack(packHandle string) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_MetaLoadVisualsForPack
	request.String_1 = packHandle

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func MetaGetNumberOfVisualsInPack(packHandle string) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_MetaGetNumberOfVisualsInPack
	request.String_1 = packHandle

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func MetaGetVisualInfo(packHandle string, visualIndex int) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_MetaGetVisualInfo
	request.Integer_1 = uint64(visualIndex)
	request.String_1 = packHandle

	resp, err := VzxControllerApiCall(request)

	return resp, err
}
