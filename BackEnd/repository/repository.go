package repository

import (
	"HealthCampaign/BackEnd/model"
	_"time"
)


type PatientRepository interface {
	Save(*model.Patient) error
	Update(string,*model.Patient) error
	Delete(string) error
	FindByID(string)(*model.Patient, error)
	FindAll()(model.Patients , error)
	FindByUserPass(string , string) bool
}


type MedicalStaffRepository interface {
	Save(*model.MedicalStaff) error
	Update(string,*model.MedicalStaff) error
	Delete(string) error
	FindByID(string)(*model.MedicalStaff, error)
	FindAll()(model.MedicalStaffs , error)
	FindByUserPass(string , string) bool
}



type CampaignRepository interface {
	Save(*model.Campaign) error
	Update(string, *model.Campaign) error
	Delete(string) error
	FindByID(string) (*model.Campaign, error)
	FindAll() (model.Campaigns, error)

}


type DiseaseRepository interface {
	Save(*model.Disease) error
	Update(string, *model.Disease) error
	Delete(string) error
	FindByID(string) (*model.Disease, error)
	FindAll() (model.Diseases, error)

}

type MedicalCenterRepository interface {
	Save(*model.MedicalCenter) error
	Update(string, *model.MedicalCenter) error
	Delete(string) error
	FindByID(string) (*model.MedicalCenter, error)
	FindAll() (model.MedicalCenters, error)

}


type ProgressRepository interface {
	Save(*model.Progress) error
	Update(string, *model.Progress) error
	Delete(string) error
	FindByID(string) (*model.Progress, error)
	FindAll() (model.Progresses, error)

}


type AlterationRepository interface {
	Save(*model.Alteration) error
	Update(string, *model.Alteration) error
	Delete(string) error
	FindByID(string) (*model.Alteration, error)
	FindAll() (model.Alterations, error)

}


type GroupRepository interface {
	Save(*model.Group) error
	Update(string, *model.Group) error
	Delete(string) error
	FindByID(string) (*model.Group, error)
	FindAll() (model.Groups, error)
}

type MessageRepository interface {
	Save(*model.Message) error
	Update(string, *model.Message) error
	Delete(string) error
	FindByID(string) (*model.Message, error)
	FindAll() (model.Messages, error)
	FindAllByGroupId(string) (model.Messages, error)
}