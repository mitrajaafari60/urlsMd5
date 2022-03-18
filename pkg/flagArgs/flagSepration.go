package flagArgs

func FlagsCount(args []string, flagNaming string) int {
	//first param is
	flagParamNo := 1

	// no flag
	if len(args) <= 1 {
		return flagParamNo
	}

	//flag with -flag value
	if args[1] == "-"+flagNaming {
		flagParamNo = flagParamNo + 1
		if len(args) > 2 {
			flagParamNo = flagParamNo + 1
		}
	}

	//flag with -flag=value
	if len(args[1]) > len(flagNaming)+2 {
		if args[1][:len(flagNaming)+2] == "-"+flagNaming+"=" {
			flagParamNo = flagParamNo + 1
		}
	}

	return flagParamNo
}
