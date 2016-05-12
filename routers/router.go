// @APIVersion 1.0.0
// @Title NoHassls Customer Records API
// @Description All the tools you need to access customer information from NoHassls
// @Contact dazmiller@gmail.com
// @TermsOfServiceUrl http://nohassls.co.nz/

package routers

import (
	"nohasslsapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		//		beego.NSNamespace("/activations",
		//			beego.NSInclude(
		//				&controllers.ActivationsController{},
		//			),
		//		),

		beego.NSNamespace("/boat_insurance",
			beego.NSInclude(
				&controllers.BoatInsuranceController{},
			),
		),

		beego.NSNamespace("/boat_insurance_quotes",
			beego.NSInclude(
				&controllers.BoatInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/broad_band",
		//			beego.NSInclude(
		//				&controllers.BroadBandController{},
		//			),
		//		),

		//		beego.NSNamespace("/broad_band_quotes",
		//			beego.NSInclude(
		//				&controllers.BroadBandQuotesController{},
		//			),
		//		),

		//		beego.NSNamespace("/buy_vehicle",
		//			beego.NSInclude(
		//				&controllers.BuyVehicleController{},
		//			),
		//		),

		//		beego.NSNamespace("/buy_vehicle_quotes",
		//			beego.NSInclude(
		//				&controllers.BuyVehicleQuotesController{},
		//			),
		//		),

		beego.NSNamespace("/car_insurance",
			beego.NSInclude(
				&controllers.CarInsuranceController{},
			),
		),

		beego.NSNamespace("/car_insurance_quotes",
			beego.NSInclude(
				&controllers.CarInsuranceQuotesController{},
			),
		),

		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.CommentsController{},
			),
		),

		//		beego.NSNamespace("/content_insurance",
		//			beego.NSInclude(
		//				&controllers.ContentInsuranceController{},
		//			),
		//		),

		//		beego.NSNamespace("/content_insurance_quotes",
		//			beego.NSInclude(
		//				&controllers.ContentInsuranceQuotesController{},
		//			),
		//		),

		beego.NSNamespace("/contents_insurance",
			beego.NSInclude(
				&controllers.ContentsInsuranceController{},
			),
		),

		beego.NSNamespace("/contents_insurance_quotes",
			beego.NSInclude(
				&controllers.ContentsInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/credit_cards",
		//			beego.NSInclude(
		//				&controllers.CreditCardsController{},
		//			),
		//		),

		//		beego.NSNamespace("/credit_cards_quotes",
		//			beego.NSInclude(
		//				&controllers.CreditCardsQuotesController{},
		//			),
		//		),

		beego.NSNamespace("/health_insurance",
			beego.NSInclude(
				&controllers.HealthInsuranceController{},
			),
		),

		beego.NSNamespace("/health_insurance_quotes",
			beego.NSInclude(
				&controllers.HealthInsuranceQuotesController{},
			),
		),

		beego.NSNamespace("/home_insurance",
			beego.NSInclude(
				&controllers.HomeInsuranceController{},
			),
		),

		//		beego.NSNamespace("/home_insurance_profile",
		//			beego.NSInclude(
		//				&controllers.HomeInsuranceProfileController{},
		//			),
		//		),

		beego.NSNamespace("/home_insurance_quotes",
			beego.NSInclude(
				&controllers.HomeInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/investment",
		//			beego.NSInclude(
		//				&controllers.InvestmentController{},
		//			),
		//		),

		//		beego.NSNamespace("/investment_quotes",
		//			beego.NSInclude(
		//				&controllers.InvestmentQuotesController{},
		//			),
		//		),

		beego.NSNamespace("/landlord_insurance",
			beego.NSInclude(
				&controllers.LandlordInsuranceController{},
			),
		),

		beego.NSNamespace("/landlord_insurance_quotes",
			beego.NSInclude(
				&controllers.LandlordInsuranceQuotesController{},
			),
		),

		beego.NSNamespace("/life_insurance",
			beego.NSInclude(
				&controllers.LifeInsuranceController{},
			),
		),

		beego.NSNamespace("/life_insurance_quotes",
			beego.NSInclude(
				&controllers.LifeInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/migration",
		//			beego.NSInclude(
		//				&controllers.MigrationController{},
		//			),
		//		),

		//		beego.NSNamespace("/mortgage",
		//			beego.NSInclude(
		//				&controllers.MortgageController{},
		//			),
		//		),

		//		beego.NSNamespace("/mortgage_quotes",
		//			beego.NSInclude(
		//				&controllers.MortgageQuotesController{},
		//			),
		//		),

		beego.NSNamespace("/motorbike_insurance",
			beego.NSInclude(
				&controllers.MotorbikeInsuranceController{},
			),
		),

		beego.NSNamespace("/motorbike_insurance_quotes",
			beego.NSInclude(
				&controllers.MotorbikeInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/no_table_supplied",
		//			beego.NSInclude(
		//				&controllers.NoTableSuppliedController{},
		//			),
		//		),

		//		beego.NSNamespace("/persistences",
		//			beego.NSInclude(
		//				&controllers.PersistencesController{},
		//			),
		//		),

		//		beego.NSNamespace("/personal_loan",
		//			beego.NSInclude(
		//				&controllers.PersonalLoanController{},
		//			),
		//		),

		//		beego.NSNamespace("/personal_loan_quotes",
		//			beego.NSInclude(
		//				&controllers.PersonalLoanQuotesController{},
		//			),
		//		),

		//		beego.NSNamespace("/posts",
		//			beego.NSInclude(
		//				&controllers.PostsController{},
		//			),
		//		),

		beego.NSNamespace("/profile",
			beego.NSInclude(
				&controllers.ProfileController{},
			),
		),

		//		beego.NSNamespace("/quote_messages",
		//			beego.NSInclude(
		//				&controllers.QuoteMessagesController{},
		//			),
		//		),

		//		beego.NSNamespace("/real_estate",
		//			beego.NSInclude(
		//				&controllers.RealEstateController{},
		//			),
		//		),

		//		beego.NSNamespace("/real_estate_quotes",
		//			beego.NSInclude(
		//				&controllers.RealEstateQuotesController{},
		//			),
		//		),

		//		beego.NSNamespace("/reminders",
		//			beego.NSInclude(
		//				&controllers.RemindersController{},
		//			),
		//		),

		//		beego.NSNamespace("/roles",
		//			beego.NSInclude(
		//				&controllers.RolesController{},
		//			),
		//		),

		//		beego.NSNamespace("/service_providers",
		//			beego.NSInclude(
		//				&controllers.ServiceProvidersController{},
		//			),
		//		),

		//		beego.NSNamespace("/superannuation",
		//			beego.NSInclude(
		//				&controllers.SuperannuationController{},
		//			),
		//		),

		//		beego.NSNamespace("/superannuation_quotes",
		//			beego.NSInclude(
		//				&controllers.SuperannuationQuotesController{},
		//			),
		//		),

		//		beego.NSNamespace("/throttle",
		//			beego.NSInclude(
		//				&controllers.ThrottleController{},
		//			),
		//		),

		beego.NSNamespace("/travel_insurance",
			beego.NSInclude(
				&controllers.TravelInsuranceController{},
			),
		),

		beego.NSNamespace("/travel_insurance_quotes",
			beego.NSInclude(
				&controllers.TravelInsuranceQuotesController{},
			),
		),

		//		beego.NSNamespace("/user",
		//			beego.NSInclude(
		//				&controllers.UserController{},
		//			),
		//		),

		//		beego.NSNamespace("/users_old",
		//			beego.NSInclude(
		//				&controllers.UsersOldController{},
		//			),
		//		),
	)
	beego.AddNamespace(ns)
}
