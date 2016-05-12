package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LifeInsurance struct {
	Id                                     int    `orm:"column(id);auto"`
	BatchId                                int    `orm:"column(batch_id);null"`
	Gender                                 string `orm:"column(gender);size(2);null"`
	Smoker                                 int    `orm:"column(smoker);null"`
	LastSmoked                             int    `orm:"column(last_smoked);null"`
	Age                                    int    `orm:"column(age);null"`
	CoverAmount                            int    `orm:"column(cover_amount);null"`
	CountryResidence                       string `orm:"column(country_residence);size(45);null"`
	ResidencyStatus                        string `orm:"column(residency_status);size(45);null"`
	Height                                 int    `orm:"column(height);null"`
	Weight                                 int    `orm:"column(weight);null"`
	Cancer                                 int    `orm:"column(cancer);null"`
	Diabetes                               int    `orm:"column(diabetes);null"`
	BloodDisorder                          int    `orm:"column(blood_disorder);null"`
	CancerType                             string `orm:"column(cancer_type);size(255);null"`
	CancerSpread                           int    `orm:"column(cancer_spread);null"`
	CancerClear10                          int    `orm:"column(cancer_clear_10);null"`
	DiabetesComplications                  int    `orm:"column(diabetes_complications);null"`
	BloodDisorderType                      string `orm:"column(blood_disorder_type);size(255);null"`
	LifeInsurancecol                       string `orm:"column(life_insurancecol);size(45);null"`
	BloodPressure                          int    `orm:"column(blood_pressure);null"`
	Cholesterol                            int    `orm:"column(cholesterol);null"`
	HeartProblem                           int    `orm:"column(heart_problem);null"`
	BloodPressureTreatment                 string `orm:"column(blood_pressure_treatment);size(45);null"`
	KidneyProblem                          int    `orm:"column(kidney_problem);null"`
	CholesterolTreatment                   string `orm:"column(cholesterol_treatment);size(45);null"`
	CholesterolCheck12                     int    `orm:"column(cholesterol_check_12);null"`
	HeartProblemDetails                    string `orm:"column(heart_problem_details);size(255);null"`
	GastroIntestinal                       int    `orm:"column(gastro_intestinal);null"`
	BladderProblem                         int    `orm:"column(bladder_problem);null"`
	LungProblem                            int    `orm:"column(lung_problem);null"`
	GastroProblemDetails                   string `orm:"column(gastro_problem_details);size(255);null"`
	HerniaFurtherProblems                  int    `orm:"column(hernia_further_problems);null"`
	GastritisEpisodes                      int    `orm:"column(gastritis_episodes);null"`
	GallstonesRemoved                      int    `orm:"column(gallstones_removed);null"`
	GallstonesOngoing                      int    `orm:"column(gallstones_ongoing);null"`
	UlcerHospitialised                     int    `orm:"column(ulcer_hospitialised);null"`
	GastroOtherTreatment                   int    `orm:"column(gastro_other_treatment);null"`
	GastroProblems5y                       int    `orm:"column(gastro_problems_5y);null"`
	KydneyBladderProblemDetails            string `orm:"column(kydney_bladder_problem_details);size(255);null"`
	LungProblemDetails                     string `orm:"column(lung_problem_details);size(255);null"`
	Neurological                           int    `orm:"column(neurological);null"`
	MuscularSkeletal                       int    `orm:"column(muscular_skeletal);null"`
	MentalHealth5y                         int    `orm:"column(mental_health_5y);null"`
	NeurologicalProblemDetails             string `orm:"column(neurological_problem_details);size(255);null"`
	MuscularSkeletalProblemDetails         string `orm:"column(muscular_skeletal_problem_details);size(255);null"`
	MuscularSkeletalNeckBack               int    `orm:"column(muscular_skeletal_neck_back);null"`
	MuscleInjuryStrainSymptoms2y           int    `orm:"column(muscle_injury_strain_symptoms_2y);null"`
	FractureBackNeckSkull                  int    `orm:"column(fracture_back_neck_skull);null"`
	FractureAnyProblems2y                  int    `orm:"column(fracture_any_problems_2y);null"`
	ArthritisDetails                       string `orm:"column(arthritis_details);size(255);null"`
	GoutTendonitisTenosynovitis2y          int    `orm:"column(gout_tendonitis_tenosynovitis_2y);null"`
	MuscularSkeletalSportsInjury           int    `orm:"column(muscular_skeletal_sports_injury);null"`
	MuscularSkeletalJointNeckBack          int    `orm:"column(muscular_skeletal_joint_neck_back);null"`
	MentalHealthDetails                    string `orm:"column(mental_health_details);size(255);null"`
	Alcohol28Week                          int    `orm:"column(alcohol_28_week);null"`
	IllegalDrugs5y                         int    `orm:"column(illegal_drugs_5y);null"`
	HivPositive                            int    `orm:"column(hiv_positive);null"`
	AlcoholProfessionalHelp                int    `orm:"column(alcohol_professional_help);null"`
	AlcoholDetoxHospital                   int    `orm:"column(alcohol_detox_hospital);null"`
	IllegalDrugUseDetails                  string `orm:"column(illegal_drug_use_details);size(255);null"`
	OtherMedicalConditions                 string `orm:"column(other_medical_conditions);size(255);null"`
	FamilyMemberCancer60                   int    `orm:"column(family_member_cancer_60);null"`
	SeekingMedicalAdviceCurrentlyDetails   string `orm:"column(seeking_medical_advice_currently_details);size(255);null"`
	ThyroidProblem2y                       int    `orm:"column(thyroid_problem_2y);null"`
	ThryoidSurgeryRadio6m                  int    `orm:"column(thryoid_surgery_radio_6m);null"`
	CurrentMedicalTreatmentProfessions     string `orm:"column(current_medical_treatment_professions);size(255);null"`
	FamilyMemberIllnessDetails             string `orm:"column(family_member_illness_details);size(255);null"`
	Family2OrMoreHeartDisease              int    `orm:"column(family_2_or_more_heart_disease);null"`
	Family2OrMoreStroke                    int    `orm:"column(family_2_or_more_stroke);null"`
	KidneyDiseasePolycystic                int    `orm:"column(kidney_disease_polycystic);null"`
	Family2OrMoreDiabetes                  int    `orm:"column(family_2_or_more_diabetes);null"`
	CancerTypeDiagnosed                    string `orm:"column(cancer_type_diagnosed);size(255);null"`
	FamilyCountColonRectal60               int    `orm:"column(family_count_colon_rectal_60);null"`
	ColonRectalBefore50                    int    `orm:"column(colon_rectal_before_50);null"`
	FamilyCountProstateCancer60            int    `orm:"column(family_count_prostate_cancer_60);null"`
	ProstateCancerBefore50                 int    `orm:"column(prostate_cancer_before_50);null"`
	Family2MoreSameCancer60                int    `orm:"column(family_2_more_same_cancer_60);null"`
	CancerBefore50                         int    `orm:"column(cancer_before_50);null"`
	FamilyCountTesticularCancer            int    `orm:"column(family_count_testicular_cancer_);null"`
	TesticularCancerBefore50               int    `orm:"column(testicular_cancer_before_50);null"`
	FamilyCountMs60                        int    `orm:"column(family_count_ms_60);null"`
	FamilyCountParkinsons60                int    `orm:"column(family_count_parkinsons_60);null"`
	FamilyCountMotorNeurone60              int    `orm:"column(family_count_motor_neurone_60);null"`
	RiskyOccupation                        string `orm:"column(risky_occupation);size(255);null"`
	RecreationalActivities                 string `orm:"column(recreational_activities);size(255);null"`
	TypeOfCancerWhenDiagnosed              string `orm:"column(type_of_cancer_when_diagnosed);size(255);null"`
	CancertMedicatedTreated                string `orm:"column(cancert_medicated_treated);size(255);null"`
	CancerStateRemission                   string `orm:"column(cancer_state_remission);size(255);null"`
	DiabetesFirstDiagnosed                 string `orm:"column(diabetes_first_diagnosed);size(255);null"`
	DiabetesMedicationControl              string `orm:"column(diabetes_medication_control);size(255);null"`
	DiabetesTests12m                       string `orm:"column(diabetes_tests_12m);size(255);null"`
	BloodDisorderDiagnosed                 string `orm:"column(blood_disorder_diagnosed);size(255);null"`
	BloodDisorderMedication                string `orm:"column(blood_disorder_medication);size(255);null"`
	BloodDisorderTests12m                  string `orm:"column(blood_disorder_tests_12m);size(255);null"`
	BloodPressureHighDiagnosed             string `orm:"column(blood_pressure_high_diagnosed);size(255);null"`
	BloodPressureMedication                string `orm:"column(blood_pressure_medication);size(255);null"`
	BloodPressureTest                      string `orm:"column(blood_pressure_test);size(255);null"`
	CholesterolDiagnosed                   string `orm:"column(cholesterol_diagnosed);size(255);null"`
	CholestrolMedication                   string `orm:"column(cholestrol_medication);size(255);null"`
	CholesterolTest                        string `orm:"column(cholesterol_test);size(255);null"`
	HeartConditionDiagnosed                string `orm:"column(heart_condition_diagnosed);size(255);null"`
	HeartConditionMedication               string `orm:"column(heart_condition_medication);size(255);null"`
	HeartConditionTest12m                  string `orm:"column(heart_condition_test_12m);size(255);null"`
	GastroIntestinalDiagnosed              string `orm:"column(gastro_intestinal_diagnosed);size(255);null"`
	GastroIntestinalMedication             string `orm:"column(gastro_intestinal_medication);size(255);null"`
	GastroIntestinalCurrentState           string `orm:"column(gastro_intestinal_current_state);size(255);null"`
	KidneyBladderDiagnosed                 string `orm:"column(kidney_bladder_diagnosed);size(255);null"`
	KidneyBladderMedication                string `orm:"column(kidney_bladder_medication);size(255);null"`
	KidneyBladderCurrentState              string `orm:"column(kidney_bladder_current_state);size(255);null"`
	LungProblemDiagnosed                   string `orm:"column(lung_problem_diagnosed);size(255);null"`
	LungProblemMedication                  string `orm:"column(lung_problem_medication);size(255);null"`
	LungProblemCurrentState                string `orm:"column(lung_problem_current_state);size(255);null"`
	NeurologicalDiagnosed                  string `orm:"column(neurological_diagnosed);size(255);null"`
	NeurologicalMedication                 string `orm:"column(neurological_medication);size(255);null"`
	NeurologicalCurrentState               string `orm:"column(neurological_current_state);size(255);null"`
	MuscularSkeletalDiagnosed              string `orm:"column(muscular_skeletal_diagnosed);size(255);null"`
	MuscularSkeletalMedication             string `orm:"column(muscular_skeletal_medication);size(255);null"`
	MuscularSkeletalCurrentStatus          string `orm:"column(muscular_skeletal_current_status);size(255);null"`
	MentalHealthDiagnosed                  string `orm:"column(mental_health_diagnosed);size(255);null"`
	MentalHealthCausedEventIllness         string `orm:"column(mental_health_caused_event_illness);size(255);null"`
	MentalHealthMedication                 string `orm:"column(mental_health_medication);size(255);null"`
	AlcoholTypicalUse                      string `orm:"column(alcohol_typical_use);size(255);null"`
	AlcoholConsultedProfessional           string `orm:"column(alcohol_consulted_professional);size(255);null"`
	AlcoholAlcoholismOrCounselling         string `orm:"column(alcohol_alcoholism_or_counselling);size(255);null"`
	IllegalDrugsUse                        string `orm:"column(illegal_drugs_use);size(255);null"`
	IllegalDrugsConsultedProfessional      string `orm:"column(illegal_drugs_consulted_professional);size(255);null"`
	IllegalDrugsHaveStoppedPeriod          string `orm:"column(illegal_drugs_have_stopped_period);size(255);null"`
	CurrentMedicalAdviceSymptoms           string `orm:"column(current_medical_advice_symptoms);size(255);null"`
	CurrentMedicalAdviceProfession         string `orm:"column(current_medical_advice_profession);size(255);null"`
	CurrentMedicalAdviceTests              string `orm:"column(current_medical_advice_tests);size(255);null"`
	FamilyMembersIllnessList               string `orm:"column(family_members_illness_list);null"`
	FamilyMembersMedicalTestsYouHadRelated string `orm:"column(family_members_medical_tests_you_had_related);size(255);null"`
}

func (t *LifeInsurance) TableName() string {
	return "life_insurance"
}

func init() {
	orm.RegisterModel(new(LifeInsurance))
}

// AddLifeInsurance insert a new LifeInsurance into database and returns
// last inserted Id on success.
func AddLifeInsurance(m *LifeInsurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLifeInsuranceById retrieves LifeInsurance by Id. Returns error if
// Id doesn't exist
func GetLifeInsuranceById(id int) (v *LifeInsurance, err error) {
	o := orm.NewOrm()
	v = &LifeInsurance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLifeInsurance retrieves all LifeInsurance matches certain condition. Returns empty list if
// no records exist
func GetAllLifeInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LifeInsurance))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []LifeInsurance
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLifeInsurance updates LifeInsurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateLifeInsuranceById(m *LifeInsurance) (err error) {
	o := orm.NewOrm()
	v := LifeInsurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLifeInsurance deletes LifeInsurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLifeInsurance(id int) (err error) {
	o := orm.NewOrm()
	v := LifeInsurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LifeInsurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
