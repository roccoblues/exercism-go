package erratum

// Use opens the resource o, and calls Frob(input) on it.
// The opened resource is guaranteed to be closed.
func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	resource, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}

	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(FrobError); ok {
				resource.Defrob(err.defrobTag)
			}
			err = r.(error)
		}
		resource.Close()
	}()

	resource.Frob(input)

	return err
}
