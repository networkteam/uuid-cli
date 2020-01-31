package uuidcli

type ExitCodeErr struct {
	ExitCode int
	err      error
}

func (e ExitCodeErr) Error() string {
	return e.err.Error()
}

func (e ExitCodeErr) Unwrap() error {
	return e.err
}

func NewExitCodeErr(err error, exitCode int) error {
	return ExitCodeErr{
		ExitCode: exitCode,
		err:      err,
	}
}
