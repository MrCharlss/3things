package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mrcharlss/3thingstoday/dal"
	"github.com/mrcharlss/3thingstoday/types"
	"github.com/mrcharlss/3thingstoday/utils"
)

func CreateThings(ctx *fiber.Ctx) error {
	b := new(types.CreteThingDTO)

	if err := ctx.BodyParser(b); err != nil {

		return err
	}

	// if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
	//     return err
	// }
	id := uint(1)
	idUint := &id

	t := &dal.Thing{
		Win:     b.Win,
		Unhappy: b.Unhappy,
		Lesson:  b.Lesson,
        DayScore: b.DayScore,
        OtherScore: b.OtherScore,
		// User:utils.GetUser(ctx),
		User: idUint,
	}
    fmt.Printf("%v",t)

	if err := dal.CreateThing(t).Error; err != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error())
	}

	return ctx.JSON(&types.ThingCreateResponse{
		Thing: &types.ThingResponse{
			ID:         t.ID,
			Win:        t.Win,
			Unhappy:    t.Unhappy,
			Lesson:     t.Lesson,
			DayScore:   t.DayScore,
			OtherScore: t.OtherScore,
		},
	})

}

func GetThings(ctx *fiber.Ctx)  error {
    d := &[]types.ThingResponse{}

    err := dal.FindThingsByUser(d, utils.GetUser(ctx)).Error
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, err.Error())
    }

    return ctx.JSON(&types.GetThingsResponse{
        Things: d,
    })
}
