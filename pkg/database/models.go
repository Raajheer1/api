package database

import (
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dbTypes "github.com/kzdv/api/pkg/database/types"
	"github.com/kzdv/api/pkg/logger"
)

var log = logger.Logger.WithField("component", "database")

func AddRoleToUser(user *dbTypes.User, role *dbTypes.Role) error {
	if err := DB.Model(user).Association("Roles").Append(role); err != nil {
		return err
	}

	return nil
}

func RemoveRoleFromUser(user *dbTypes.User, role *dbTypes.Role) error {
	if err := DB.Model(user).Association("Roles").Delete(role); err != nil {
		return err
	}

	return nil
}

func FindEmailTemplate(name string) (*dbTypes.EmailTemplate, error) {
	template := &dbTypes.EmailTemplate{}
	if err := DB.Where(dbTypes.EmailTemplate{Name: name}).First(template).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return template, nil
}

func FindRole(name string) (*dbTypes.Role, error) {
	role := &dbTypes.Role{}
	if err := DB.Where(dbTypes.Role{Name: name}).First(role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			role = &dbTypes.Role{
				Name: name,
			}
			if err := DB.Create(role).Error; err != nil {
				return nil, err
			}
			return role, nil
		}
		return nil, err
	}

	return role, nil
}

func FindUsersWithRole(role string) ([]dbTypes.User, error) {
	var users []dbTypes.User

	r, err := FindRole(role)
	if err != nil {
		return nil, err
	}

	if err := DB.Model(r).Association("Users").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserByCID(cid string) (*dbTypes.User, error) {
	user := &dbTypes.User{}
	log.Tracef("Finding user by CID: %s -- atou()=%+v", cid, atou(cid))
	if err := DB.Preload(clause.Associations).Where(dbTypes.User{CID: atou(cid)}).First(&user).Error; err != nil {
		log.Tracef("Error finding user by CID (%s): %v", cid, err)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func FindOI(user *dbTypes.User) (string, error) {
	oi := string(user.FirstName[0]) + string(user.LastName[0])
	if err := DB.Where(dbTypes.User{OperatingInitials: oi}).First(&dbTypes.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return oi, nil
		}
		return "", err
	}

	return "", nil
}

func FindRatingByShort(short string) (*dbTypes.Rating, error) {
	rating := &dbTypes.Rating{}
	if err := DB.Where(dbTypes.Rating{Short: short}).First(rating).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return rating, nil
}

func AddDelayedJob(queue, body string, duration time.Duration) error {
	djob := &dbTypes.DelayedJob{
		Queue:     queue,
		Body:      body,
		NotBefore: time.Now().Add(duration),
	}
	if err := DB.Create(djob).Error; err != nil {
		log.Errorf("Error creating delayed lob %+v: %v", djob, err)
		return err
	}

	return nil
}

func atou(a string) uint {
	i, err := strconv.ParseUint(a, 10, 0)
	if err != nil {
		log.Warnf("Error converting string (%s) to uint: %v", a, err)
		return 0
	}
	return uint(i)
}
