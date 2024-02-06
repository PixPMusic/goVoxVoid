package voxvoid

func EngineExecuteCommand(command string) (Response, error) {
	request := Request{}
	var err error

	request.Action_t = RequestAction_EngineExecuteCommand

	// Divide the command into 4 strings by 64 byte barriers

	len := len(command)
	if len > 64 {
		request.String_1 = command[:64]
		if len > 128 {
			request.String_2 = command[64:128]
			if len > 192 {
				request.String_3 = command[128:192]
				if len > 256 {
					request.String_4 = command[192:256]
				} else {
					request.String_4 = command[192:]
				}
			} else {
				request.String_3 = command[128:]
			}
		} else {
			request.String_2 = command[64:]
		}
	} else {
		request.String_1 = command
	}

	resp, err := VzxControllerApiCall(request)
	return resp, err //resp.Integer_1, err
}
