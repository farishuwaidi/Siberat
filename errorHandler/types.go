package errorhandler

type NotFoundError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type UnAuthorizedError struct {
	Message string
}

func (c *NotFoundError) Error() string {
	return c.Message
}

func (c *BadRequestError) Error() string {
	return c.Message
}

func (c *InternalServerError) Error() string {
	return c.Message
}

func (c *UnAuthorizedError) Error() string {
	return c.Message
}