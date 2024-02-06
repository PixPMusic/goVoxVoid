package voxvoid

func PalletteGetNumberOfPallettes() (uint64, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_PaletteGetNumberOfPalettes

	resp, err := VzxControllerApiCall(request)
	return resp.Integer_1, err
}

func PalletteGetPalletteInfo(palletteIndex int) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_PaletteGetPaletteInfo
	request.Integer_1 = uint64(palletteIndex)

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func PalletteCurrentPalletteSet(palletteHandle string) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_PaletteCurrentPaletteSet
	request.String_1 = palletteHandle

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func PalletteGetColorInformation(palletteHandle string, colorIndex int) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_PaletteGetColorInfo
	request.String_1 = palletteHandle
	request.Integer_1 = uint64(colorIndex)

	resp, err := VzxControllerApiCall(request)

	return resp, err
}

func PalletteSetColor(palletteHandle string, colorHandle string, colorH float64, colorS float64, colorV float64) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_PaletteSetColor
	request.String_1 = palletteHandle
	request.String_2 = colorHandle
	request.Double_1 = colorH
	request.Double_2 = colorS
	request.Double_3 = colorV

	resp, err := VzxControllerApiCall(request)

	return resp, err
}
