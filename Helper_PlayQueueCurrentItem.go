package voxvoid

func PlayQueueCurrentItemFxLevelGet() (float64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueCurrentItemFxLevelGet
	response, err := VzxControllerApiCall(request)
	return response.Double_1, err
}

func PlayQueueCurrentItemFxLevelSet(newFxLevel float64) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueCurrentItemFxLevelSet
	request.Double_1 = newFxLevel
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueCurrentItemSpeedGet() (float64, uint64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueCurrentItemSpeedGet
	response, err := VzxControllerApiCall(request)
	return response.Double_1, response.Integer_1, err
}

func PlayQueueCurrentItemSpeedSet(newSpeedFactor float64) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueCurrentItemSpeedSet
	request.Double_1 = newSpeedFactor
	_, err = VzxControllerApiCall(request)
	return err
}
