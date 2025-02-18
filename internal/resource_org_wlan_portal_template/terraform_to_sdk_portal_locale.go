package resource_org_wlan_portal_template

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
)

func localeTerraformToSdk(d attr.Value) *models.WlanPortalTemplateSettingLocale {
	data := models.WlanPortalTemplateSettingLocale{}
	if !d.IsNull() && !d.IsUnknown() {
		var vInterface interface{} = d
		plan := vInterface.(LocalesValue)
		// plan, e := NewLocalesValue(LocalesValue{}.AttributeTypes(ctx), d)
		// if e != nil {
		// 	diags.Append(e...)
		// } else {
		if !plan.AuthButtonAmazon.IsNull() && !plan.AuthButtonAmazon.IsUnknown() {
			data.AuthButtonAmazon = plan.AuthButtonAmazon.ValueStringPointer()
		}

		if !plan.AuthButtonAzure.IsNull() && !plan.AuthButtonAzure.IsUnknown() {
			data.AuthButtonAzure = plan.AuthButtonAzure.ValueStringPointer()
		}

		if !plan.AuthButtonEmail.IsNull() && !plan.AuthButtonEmail.IsUnknown() {
			data.AuthButtonEmail = plan.AuthButtonEmail.ValueStringPointer()
		}

		if !plan.AuthButtonFacebook.IsNull() && !plan.AuthButtonFacebook.IsUnknown() {
			data.AuthButtonFacebook = plan.AuthButtonFacebook.ValueStringPointer()
		}

		if !plan.AuthButtonGoogle.IsNull() && !plan.AuthButtonGoogle.IsUnknown() {
			data.AuthButtonGoogle = plan.AuthButtonGoogle.ValueStringPointer()
		}

		if !plan.AuthButtonMicrosoft.IsNull() && !plan.AuthButtonMicrosoft.IsUnknown() {
			data.AuthButtonMicrosoft = plan.AuthButtonMicrosoft.ValueStringPointer()
		}

		if !plan.AuthButtonPassphrase.IsNull() && !plan.AuthButtonPassphrase.IsUnknown() {
			data.AuthButtonPassphrase = plan.AuthButtonPassphrase.ValueStringPointer()
		}

		if !plan.AuthButtonSms.IsNull() && !plan.AuthButtonSms.IsUnknown() {
			data.AuthButtonSms = plan.AuthButtonSms.ValueStringPointer()
		}

		if !plan.AuthButtonSponsor.IsNull() && !plan.AuthButtonSponsor.IsUnknown() {
			data.AuthButtonSponsor = plan.AuthButtonSponsor.ValueStringPointer()
		}

		if !plan.AuthLabel.IsNull() && !plan.AuthLabel.IsUnknown() {
			data.AuthLabel = plan.AuthLabel.ValueStringPointer()
		}

		if !plan.BackLink.IsNull() && !plan.BackLink.IsUnknown() {
			data.BackLink = plan.BackLink.ValueStringPointer()
		}

		if !plan.CompanyError.IsNull() && !plan.CompanyError.IsUnknown() {
			data.CompanyError = plan.CompanyError.ValueStringPointer()
		}

		if !plan.CompanyLabel.IsNull() && !plan.CompanyLabel.IsUnknown() {
			data.CompanyLabel = plan.CompanyLabel.ValueStringPointer()
		}

		if !plan.EmailAccessDomainError.IsNull() && !plan.EmailAccessDomainError.IsUnknown() {
			data.EmailAccessDomainError = plan.EmailAccessDomainError.ValueStringPointer()
		}

		if !plan.EmailCancel.IsNull() && !plan.EmailCancel.IsUnknown() {
			data.EmailCancel = plan.EmailCancel.ValueStringPointer()
		}

		if !plan.EmailCodeError.IsNull() && !plan.EmailCodeError.IsUnknown() {
			data.EmailCodeError = plan.EmailCodeError.ValueStringPointer()
		}

		if !plan.EmailError.IsNull() && !plan.EmailError.IsUnknown() {
			data.EmailError = plan.EmailError.ValueStringPointer()
		}

		if !plan.EmailFieldLabel.IsNull() && !plan.EmailFieldLabel.IsUnknown() {
			data.EmailFieldLabel = plan.EmailFieldLabel.ValueStringPointer()
		}

		if !plan.EmailLabel.IsNull() && !plan.EmailLabel.IsUnknown() {
			data.EmailLabel = plan.EmailLabel.ValueStringPointer()
		}

		if !plan.EmailMessage.IsNull() && !plan.EmailMessage.IsUnknown() {
			data.EmailMessage = plan.EmailMessage.ValueStringPointer()
		}

		if !plan.EmailSubmit.IsNull() && !plan.EmailSubmit.IsUnknown() {
			data.EmailSubmit = plan.EmailSubmit.ValueStringPointer()
		}

		if !plan.EmailTitle.IsNull() && !plan.EmailTitle.IsUnknown() {
			data.EmailTitle = plan.EmailTitle.ValueStringPointer()
		}

		if !plan.Field1error.IsNull() && !plan.Field1error.IsUnknown() {
			data.Field1Error = plan.Field1error.ValueStringPointer()
		}

		if !plan.Field1label.IsNull() && !plan.Field1label.IsUnknown() {
			data.Field1Label = plan.Field1label.ValueStringPointer()
		}

		if !plan.Field2error.IsNull() && !plan.Field2error.IsUnknown() {
			data.Field2Error = plan.Field2error.ValueStringPointer()
		}

		if !plan.Field2label.IsNull() && !plan.Field2label.IsUnknown() {
			data.Field2Label = plan.Field2label.ValueStringPointer()
		}

		if !plan.Field3error.IsNull() && !plan.Field3error.IsUnknown() {
			data.Field3Error = plan.Field3error.ValueStringPointer()
		}

		if !plan.Field3label.IsNull() && !plan.Field3label.IsUnknown() {
			data.Field3Label = plan.Field3label.ValueStringPointer()
		}

		if !plan.Field4error.IsNull() && !plan.Field4error.IsUnknown() {
			data.Field4Error = plan.Field4error.ValueStringPointer()
		}

		if !plan.Field4label.IsNull() && !plan.Field4label.IsUnknown() {
			data.Field4Label = plan.Field4label.ValueStringPointer()
		}

		if !plan.Message.IsNull() && !plan.Message.IsUnknown() {
			data.Message = plan.Message.ValueStringPointer()
		}

		if !plan.NameError.IsNull() && !plan.NameError.IsUnknown() {
			data.NameError = plan.NameError.ValueStringPointer()
		}

		if !plan.NameLabel.IsNull() && !plan.NameLabel.IsUnknown() {
			data.NameLabel = plan.NameLabel.ValueStringPointer()
		}

		if !plan.OptoutLabel.IsNull() && !plan.OptoutLabel.IsUnknown() {
			data.OptoutLabel = plan.OptoutLabel.ValueStringPointer()
		}

		if !plan.PassphraseCancel.IsNull() && !plan.PassphraseCancel.IsUnknown() {
			data.PassphraseCancel = plan.PassphraseCancel.ValueStringPointer()
		}

		if !plan.PassphraseError.IsNull() && !plan.PassphraseError.IsUnknown() {
			data.PassphraseError = plan.PassphraseError.ValueStringPointer()
		}

		if !plan.PassphraseLabel.IsNull() && !plan.PassphraseLabel.IsUnknown() {
			data.PassphraseLabel = plan.PassphraseLabel.ValueStringPointer()
		}

		if !plan.PassphraseMessage.IsNull() && !plan.PassphraseMessage.IsUnknown() {
			data.PassphraseMessage = plan.PassphraseMessage.ValueStringPointer()
		}

		if !plan.PassphraseSubmit.IsNull() && !plan.PassphraseSubmit.IsUnknown() {
			data.PassphraseSubmit = plan.PassphraseSubmit.ValueStringPointer()
		}

		if !plan.PassphraseTitle.IsNull() && !plan.PassphraseTitle.IsUnknown() {
			data.PassphraseTitle = plan.PassphraseTitle.ValueStringPointer()
		}

		if !plan.PrivacyPolicyAcceptLabel.IsNull() && !plan.PrivacyPolicyAcceptLabel.IsUnknown() {
			data.PrivacyPolicyAcceptLabel = plan.PrivacyPolicyAcceptLabel.ValueStringPointer()
		}

		if !plan.PrivacyPolicyError.IsNull() && !plan.PrivacyPolicyError.IsUnknown() {
			data.PrivacyPolicyError = plan.PrivacyPolicyError.ValueStringPointer()
		}

		if !plan.PrivacyPolicyLink.IsNull() && !plan.PrivacyPolicyLink.IsUnknown() {
			data.PrivacyPolicyLink = plan.PrivacyPolicyLink.ValueStringPointer()
		}

		if !plan.PrivacyPolicyText.IsNull() && !plan.PrivacyPolicyText.IsUnknown() {
			data.PrivacyPolicyText = plan.PrivacyPolicyText.ValueStringPointer()
		}

		if !plan.RequiredFieldLabel.IsNull() && !plan.RequiredFieldLabel.IsUnknown() {
			data.RequiredFieldLabel = plan.RequiredFieldLabel.ValueStringPointer()
		}

		if !plan.SignInLabel.IsNull() && !plan.SignInLabel.IsUnknown() {
			data.SignInLabel = plan.SignInLabel.ValueStringPointer()
		}

		if !plan.SmsCarrierDefault.IsNull() && !plan.SmsCarrierDefault.IsUnknown() {
			data.SmsCarrierDefault = plan.SmsCarrierDefault.ValueStringPointer()
		}

		if !plan.SmsCarrierError.IsNull() && !plan.SmsCarrierError.IsUnknown() {
			data.SmsCarrierError = plan.SmsCarrierError.ValueStringPointer()
		}

		if !plan.SmsCarrierFieldLabel.IsNull() && !plan.SmsCarrierFieldLabel.IsUnknown() {
			data.SmsCarrierFieldLabel = plan.SmsCarrierFieldLabel.ValueStringPointer()
		}

		if !plan.SmsCodeCancel.IsNull() && !plan.SmsCodeCancel.IsUnknown() {
			data.SmsCodeCancel = plan.SmsCodeCancel.ValueStringPointer()
		}

		if !plan.SmsCodeError.IsNull() && !plan.SmsCodeError.IsUnknown() {
			data.SmsCodeError = plan.SmsCodeError.ValueStringPointer()
		}

		if !plan.SmsCodeFieldLabel.IsNull() && !plan.SmsCodeFieldLabel.IsUnknown() {
			data.SmsCodeFieldLabel = plan.SmsCodeFieldLabel.ValueStringPointer()
		}

		if !plan.SmsCodeMessage.IsNull() && !plan.SmsCodeMessage.IsUnknown() {
			data.SmsCodeMessage = plan.SmsCodeMessage.ValueStringPointer()
		}

		if !plan.SmsCodeSubmit.IsNull() && !plan.SmsCodeSubmit.IsUnknown() {
			data.SmsCodeSubmit = plan.SmsCodeSubmit.ValueStringPointer()
		}

		if !plan.SmsCodeTitle.IsNull() && !plan.SmsCodeTitle.IsUnknown() {
			data.SmsCodeTitle = plan.SmsCodeTitle.ValueStringPointer()
		}

		if !plan.SmsCountryFieldLabel.IsNull() && !plan.SmsCountryFieldLabel.IsUnknown() {
			data.SmsCountryFieldLabel = plan.SmsCountryFieldLabel.ValueStringPointer()
		}

		if !plan.SmsCountryFormat.IsNull() && !plan.SmsCountryFormat.IsUnknown() {
			data.SmsCountryFormat = plan.SmsCountryFormat.ValueStringPointer()
		}

		if !plan.SmsHaveAccessCode.IsNull() && !plan.SmsHaveAccessCode.IsUnknown() {
			data.SmsHaveAccessCode = plan.SmsHaveAccessCode.ValueStringPointer()
		}

		if !plan.SmsMessageFormat.IsNull() && !plan.SmsMessageFormat.IsUnknown() {
			data.SmsMessageFormat = plan.SmsMessageFormat.ValueStringPointer()
		}

		if !plan.SmsNumberCancel.IsNull() && !plan.SmsNumberCancel.IsUnknown() {
			data.SmsNumberCancel = plan.SmsNumberCancel.ValueStringPointer()
		}

		if !plan.SmsNumberError.IsNull() && !plan.SmsNumberError.IsUnknown() {
			data.SmsNumberError = plan.SmsNumberError.ValueStringPointer()
		}

		if !plan.SmsNumberFieldLabel.IsNull() && !plan.SmsNumberFieldLabel.IsUnknown() {
			data.SmsNumberFieldLabel = plan.SmsNumberFieldLabel.ValueStringPointer()
		}

		if !plan.SmsNumberFormat.IsNull() && !plan.SmsNumberFormat.IsUnknown() {
			data.SmsNumberFormat = plan.SmsNumberFormat.ValueStringPointer()
		}

		if !plan.SmsNumberMessage.IsNull() && !plan.SmsNumberMessage.IsUnknown() {
			data.SmsNumberMessage = plan.SmsNumberMessage.ValueStringPointer()
		}

		if !plan.SmsNumberSubmit.IsNull() && !plan.SmsNumberSubmit.IsUnknown() {
			data.SmsNumberSubmit = plan.SmsNumberSubmit.ValueStringPointer()
		}

		if !plan.SmsNumberTitle.IsNull() && !plan.SmsNumberTitle.IsUnknown() {
			data.SmsNumberTitle = plan.SmsNumberTitle.ValueStringPointer()
		}

		if !plan.SmsUsernameFormat.IsNull() && !plan.SmsUsernameFormat.IsUnknown() {
			data.SmsUsernameFormat = plan.SmsUsernameFormat.ValueStringPointer()
		}

		if !plan.SponsorBackLink.IsNull() && !plan.SponsorBackLink.IsUnknown() {
			data.SponsorBackLink = plan.SponsorBackLink.ValueStringPointer()
		}

		if !plan.SponsorCancel.IsNull() && !plan.SponsorCancel.IsUnknown() {
			data.SponsorCancel = plan.SponsorCancel.ValueStringPointer()
		}

		if !plan.SponsorEmail.IsNull() && !plan.SponsorEmail.IsUnknown() {
			data.SponsorEmail = plan.SponsorEmail.ValueStringPointer()
		}

		if !plan.SponsorEmailError.IsNull() && !plan.SponsorEmailError.IsUnknown() {
			data.SponsorEmailError = plan.SponsorEmailError.ValueStringPointer()
		}

		if !plan.SponsorInfoApproved.IsNull() && !plan.SponsorInfoApproved.IsUnknown() {
			data.SponsorInfoApproved = plan.SponsorInfoApproved.ValueStringPointer()
		}

		if !plan.SponsorInfoDenied.IsNull() && !plan.SponsorInfoDenied.IsUnknown() {
			data.SponsorInfoDenied = plan.SponsorInfoDenied.ValueStringPointer()
		}

		if !plan.SponsorInfoPending.IsNull() && !plan.SponsorInfoPending.IsUnknown() {
			data.SponsorInfoPending = plan.SponsorInfoPending.ValueStringPointer()
		}

		if !plan.SponsorName.IsNull() && !plan.SponsorName.IsUnknown() {
			data.SponsorName = plan.SponsorName.ValueStringPointer()
		}

		if !plan.SponsorNameError.IsNull() && !plan.SponsorNameError.IsUnknown() {
			data.SponsorNameError = plan.SponsorNameError.ValueStringPointer()
		}

		if !plan.SponsorNotePending.IsNull() && !plan.SponsorNotePending.IsUnknown() {
			data.SponsorNotePending = plan.SponsorNotePending.ValueStringPointer()
		}

		if !plan.SponsorStatusApproved.IsNull() && !plan.SponsorStatusApproved.IsUnknown() {
			data.SponsorStatusApproved = plan.SponsorStatusApproved.ValueStringPointer()
		}

		if !plan.SponsorStatusDenied.IsNull() && !plan.SponsorStatusDenied.IsUnknown() {
			data.SponsorStatusDenied = plan.SponsorStatusDenied.ValueStringPointer()
		}

		if !plan.SponsorStatusPending.IsNull() && !plan.SponsorStatusPending.IsUnknown() {
			data.SponsorStatusPending = plan.SponsorStatusPending.ValueStringPointer()
		}

		if !plan.SponsorSubmit.IsNull() && !plan.SponsorSubmit.IsUnknown() {
			data.SponsorSubmit = plan.SponsorSubmit.ValueStringPointer()
		}

		if !plan.SponsorsError.IsNull() && !plan.SponsorsError.IsUnknown() {
			data.SponsorsError = plan.SponsorsError.ValueStringPointer()
		}

		if !plan.TosAcceptLabel.IsNull() && !plan.TosAcceptLabel.IsUnknown() {
			data.TosAcceptLabel = plan.TosAcceptLabel.ValueStringPointer()
		}

		if !plan.TosError.IsNull() && !plan.TosError.IsUnknown() {
			data.TosError = plan.TosError.ValueStringPointer()
		}

		if !plan.TosLink.IsNull() && !plan.TosLink.IsUnknown() {
			data.TosLink = plan.TosLink.ValueStringPointer()
		}

		if !plan.TosText.IsNull() && !plan.TosText.IsUnknown() {
			data.TosText = plan.TosText.ValueStringPointer()
		}
		// }
	}

	return &data
}

func portalTemplateLocalesTerraformToSdk(plan *PortalTemplateValue, data *models.WlanPortalTemplateSetting) {

	for k, v := range plan.Locales.Elements() {
		var vInterface interface{} = v
		o := vInterface.(LocalesValue)
		locale := localeTerraformToSdk(o)
		switch k {
		case "ar":
			data.Ar = locale
		case "ca-ES":
			data.CaES = locale
		case "cs-CZ":
			data.CsCZ = locale
		case "da-DK":
			data.DaDK = locale
		case "de-DE":
			data.DeDE = locale
		case "el-GR":
			data.ElGR = locale
		case "en-GB":
			data.EnGB = locale
		case "en-US":
			data.EnUS = locale
		case "es-ES":
			data.EsES = locale
		case "fi-FI":
			data.FiFI = locale
		case "fr-FR":
			data.FrFR = locale
		case "he-IL":
			data.HeIL = locale
		case "hi-IN":
			data.HiIN = locale
		case "hr-HR":
			data.HrHR = locale
		case "hu-HU":
			data.HuHU = locale
		case "id-ID":
			data.IdID = locale
		case "it-IT":
			data.ItIT = locale
		case "ja-JP":
			data.JaJP = locale
		case "ko-KR":
			data.KoKR = locale
		case "ms-MY":
			data.MsMY = locale
		case "nb-NO":
			data.NbNO = locale
		case "nl-NL":
			data.NlNL = locale
		case "pl-PL":
			data.PlPL = locale
		case "pt-BR":
			data.PtBR = locale
		case "pt-PT":
			data.PtPT = locale
		case "ro-RO":
			data.RoRO = locale
		case "ru-RU":
			data.RuRU = locale
		case "sk-SK":
			data.SkSK = locale
		case "sv-SE":
			data.SvSE = locale
		case "th-TH":
			data.ThTH = locale
		case "tr-TR":
			data.TrTR = locale
		case "uk-UA":
			data.UkUA = locale
		case "vi-VN":
			data.ViVN = locale
		case "zh-Hans":
			data.ZhHans = locale
		case "zh-Hant":
			data.ZhHans = locale
		}
	}

}
