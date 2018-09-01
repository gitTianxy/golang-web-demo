package base

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErr2Bool(err error) bool  {
	return err == nil
}