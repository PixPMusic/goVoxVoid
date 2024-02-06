package voxvoid

func PlayQueueProgressionPlay() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionPlay
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionStop() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionStop
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionSelectPreviousVisual() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionPreviousVisual
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionSelectNextVisual() error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionNextVisual
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionSetCurrentIndex(index uint64) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionSetCurrentIndex
	request.Integer_1 = index
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionGetCurrentIndex() (uint64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionGetCurrentIndex
	response, err := VzxControllerApiCall(request)
	return response.Integer_1, err
}

func PlayQueueProgressionBaseTimeGet() (float64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionBaseTimeGet
	response, err := VzxControllerApiCall(request)
	return response.Double_1, err
}

func PlayQueueProgressionBaseTimeSet(newBaseTimeInSeconds float64) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionBaseTimeSet
	request.Double_1 = newBaseTimeInSeconds
	_, err = VzxControllerApiCall(request)
	return err
}

func PlayQueueProgressionRandomFactorGet() (float64, error) {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionRandomFactorGet
	response, err := VzxControllerApiCall(request)
	return response.Double_1, err
}

func PlayQueueProgressionRandomFactorSet(newRandomFactor float64) error {
	var request Request
	var err error
	request.Action_t = RequestAction_PlayQueueProgressionRandomFactorSet
	request.Double_1 = newRandomFactor
	_, err = VzxControllerApiCall(request)
	return err
}
