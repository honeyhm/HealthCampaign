package usecase

import (
	"HealthCampaign/BackEnd/model"
)


type MedicalCenterUsecase interface {
	SaveMedicalCenter(user *model.MedicalCenter) (*model.MedicalCenter, error)
	UpdateMedicalCenter(string, *model.MedicalCenter) (*model.MedicalCenter, error)
	GetByID(string) (*model.MedicalCenter, error)
	GetAll() (model.MedicalCenters, error)
}


type DiseaseUsecase interface {
	SaveDisease(user *model.Disease) (*model.Disease, error)
	UpdateDisease(string, *model.Disease) (*model.Disease, error)
	GetByID(string) (*model.Disease, error)
	GetAll() (model.Diseases, error)

}


type CampaignUsecase interface {
	SaveCampaign(user *model.Campaign) (*model.Campaign, error)
	UpdateCampaign(string, *model.Campaign) (*model.Campaign, error)
	GetByID(string) (*model.Campaign, error)
	GetAll() (model.Campaigns, error)

}



type AlterationUsecase interface {
	SaveAlteration(user *model.Alteration) (*model.Alteration, error)
	UpdateAlteration(string, *model.Alteration) (*model.Alteration, error)
	GetByID(string) (*model.Alteration, error)
	GetAll() (model.Alterations, error)
}


type PatientUsecase interface {
	SavePatient(user *model.Patient) (*model.Patient, error)
	UpdatePatient(string, *model.Patient) (*model.Patient, error)
	GetByID(string) (*model.Patient, error)
	GetByUserPass( string ,  string) (*model.Patient, error)
	GetAll() (model.Patients, error)
}

type MedicalStaffUsecase interface {
	SaveMedicalStaff(user *model.MedicalStaff) (*model.MedicalStaff, error)
	UpdateMedicalStaff(string, *model.MedicalStaff) (*model.MedicalStaff, error)
	GetByID(string) (*model.MedicalStaff, error)
	GetByUserPass( string ,  string) (*model.MedicalStaff, error)
	GetAll() (model.MedicalStaffs, error)
}


type ProgressUsecase interface {
	SaveProgress(user *model.Progress) (*model.Progress, error)
	UpdateProgress(string, *model.Progress) (*model.Progress, error)
	GetByID(string) (*model.Progress, error)
	GetAll() (model.Progresses, error)
}


type GroupUsecase interface {
	SaveGroup(user *model.Group) (*model.Group, error)
	UpdateGroup(string, *model.Group) (*model.Group, error)
	GetByID(string) (*model.Group, error)
	GetAll() (model.Groups, error)
}

type MessageUsecase interface {
	SaveMessage(user *model.Message) (*model.Message, error)
	UpdateMessage(string, *model.Message) (*model.Message, error)
	GetByID(string) (*model.Message, error)
	GetAll() (model.Messages, error)
}
