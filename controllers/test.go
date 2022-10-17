package controllers

func AddBizInfo(c *fiber.Ctx) error {
	biz := new(entities.Business)
	if err := c.BodyParser(&biz); err != nil {
		return utils.ResponseHandler(
			c,
			500,
			false,
			err.Error(),
			nil,
			0,
		)
	}
	apiKey := security.GenerateApiKey()
	biz.ApiKey = &apiKey
	createBiz := database.Database.Create(&biz)
	if createBiz.Error != nil {
		return utils.ResponseHandler(
			c,
			500,
			false,
			createBiz.Error.Error(),
			nil,
			0,
		)
	}
	return utils.ResponseHandler(
		c,
		201,
		true,
		"Your Business registered successfuly",
		&biz,
		0,
	)
}