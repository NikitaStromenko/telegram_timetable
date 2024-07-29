package main

import "strings"

func readArgs(args []string) map[string]string {
	if args != nil && cap(args) > 1 {
		argsMap := make(map[string]string)

		for index, arg := range args {
			if strings.Contains(arg, "--") {
				argKeyVal := strings.Split(arg, "=")

				if len(argKeyVal) == 2 {
					argsMap[strings.Replace(argKeyVal[0], "--", "", 1)] = argKeyVal[1]
				} else if len(argKeyVal) > 2 {
					var sb strings.Builder
					for _, str := range argKeyVal[1:] {
						if sb.Len() > 0 {
							sb.WriteString("=")
						}
						sb.WriteString(str)
					}
					argsMap[strings.Replace(argKeyVal[0], "--", "", 1)] = sb.String()
				} else {
					if len(args) > index+1 {
						nextArg := args[index+1]
						if !strings.Contains(nextArg, "--") {
							argsMap[strings.Replace(arg, "--", "", 1)] = nextArg
						}
					}
				}
			}
		}

		return argsMap
	}
	return nil
}
