package listeners

import (
	"errors"

	"github.com/golang/protobuf/proto"

	"github.com/totoval/framework/config"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/hub"

	"totoval/app/events"
	pbs "totoval/app/events/protocol_buffers"
	"totoval/app/models"
)

func init() {
	hub.Register(&AddUserAffiliation{})
}

type AddUserAffiliation struct {
	user                models.User
	affiliationFromCode *string
	hub.Listen
}

func (auaff *AddUserAffiliation) Name() hub.ListenerName {
	return "add-user-affiliation"
}

func (auaff *AddUserAffiliation) Subscribe() (eventPtrList []hub.Eventer) {
	return []hub.Eventer{
		&events.UserRegistered{},
	}
}

func (auaff *AddUserAffiliation) Construct(paramPtr proto.Message) error {
	// event type assertion
	param, ok := paramPtr.(*pbs.UserRegistered)
	if !ok {
		return errors.New("listener param is invalid")
	}

	auaff.affiliationFromCode = nil
	if param.AffiliationFromCode != "" && checkFromCode(param.AffiliationFromCode) {
		auaff.affiliationFromCode = &param.AffiliationFromCode
	}

	uid := uint(param.GetUserId())
	auaff.user = models.User{ID: &uid}
	if err := m.H().First(&auaff.user, false); err != nil {
		return err
	}

	return nil
}

func (auaff *AddUserAffiliation) Handle() error {
	// add user affiliation
	if config.GetBool("user_affiliation.enable") {
		uaffPtr := &models.UserAffiliation{
			UserID: auaff.user.ID,
		}
		var err error
		if auaff.affiliationFromCode != nil {
			err = uaffPtr.InsertNode(&auaff.user, *auaff.affiliationFromCode)
		} else {
			err = uaffPtr.InsertNode(&auaff.user)
		}
		if err != nil {
			return errors.New("user affiliation insert failed")
		}
	}

	return nil
}

// check affiliationFromCode is valid
func checkFromCode(affiliationFromCode string) bool {
	uaff := models.UserAffiliation{
		Code: &affiliationFromCode,
	}
	return m.H().Exist(&uaff, false)
}
