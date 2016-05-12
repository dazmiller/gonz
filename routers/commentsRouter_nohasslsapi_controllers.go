package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:BoatInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CarInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:CommentsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ContentsInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HealthInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:HomeInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LandlordInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:LifeInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:MotorbikeInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:ProfileController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:QuoteMessagesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"] = append(beego.GlobalControllerRouter["nohasslsapi/controllers:TravelInsuranceQuotesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
