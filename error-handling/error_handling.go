package erratum

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	resource, err = o()
	for err != nil {
		if _, valid := err.(TransientError); !valid {
			return err
		}
		resource, err = o()
	}

	defer func() {
		if r := recover(); r != nil {
			if err, valid := r.(FrobError); valid {
				resource.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)

	return err
}
