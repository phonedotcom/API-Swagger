package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/waiyuen/Phone.com-API-SDK-go"
	"io/ioutil"
	"strings"
	"time"
)

type SampleJsonCreator struct{}

func (s *SampleJsonCreator) createSampleInOutIfNeeded(param CliParams) {

	inputType := "json"

	if strings.EqualFold(param.inputFormat, "xml") {
		inputType = "xml"
	}

	if strings.EqualFold(param.inputFormat, "csv") {
		inputType = "csv"
	}

	if param.samplein != "" {

		stringEmailSlice := make([]string, 0)
		stringEmailSlice = append(stringEmailSlice, "asd@asd.com")

		switch param.samplein {

		case createCall:
			createCallParamsSample := swagger.CreateCallParams{"+18189640647", 1767963, "+19109485024", true, "+19109485024", 1750618, "+18189640647", true}
			s.marshalInput(createCallParamsSample, "createCall", inputType)

		case createDevice:
			createDeviceParamsSample := swagger.CreateDeviceParams{randomString(12), nil}
			s.marshalInput(createDeviceParamsSample, "createDevice", inputType)

		case createExtension:
			createExtensionParamsSample := swagger.CreateExtensionParams{"+12019570328", "unlimited", true, int32(randomNumber(10, 9999999)), true, "The name", "The full name", "America/Los_Angeles", swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, 619, true, false, true, false, 12345, "standard", swagger.MediaSummary{int32(randomNumber(10, 99999)), randomString(12)}, "automated", stringEmailSlice, "+18587741111", stringEmailSlice, "+18587748888"}
			s.marshalInput(createExtensionParamsSample, "createExtension", inputType)

		case createContact:
			createContactParamsSample := swagger.CreateContactParams{"Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", nil, nil, nil, nil}
			s.marshalInput(createContactParamsSample, "createContact", inputType)

		case createGroup:
			createGroupParamsSample := swagger.CreateGroupParams{"Ferengi Traders"}
			s.marshalInput(createGroupParamsSample, "createGroup", inputType)

		case createMedia:
			createMediaParamsSample := swagger.CreateMediaParams{randomString(12), "tts", "hold_music", "allison", randomString(100), "Y", 900, 100, "Notes aboute the media", "N"}
			s.marshalInput(createMediaParamsSample, "createMedia", inputType)

		case createMenu:
			createMenuParamsSample := swagger.CreateMenuParams{randomString(12), nil, nil, true, 3, nil, nil}
			s.marshalInput(createMenuParamsSample, "createMenu", inputType)

		case createPhoneNumber:
			createPhoneNumberParamsSample := swagger.CreatePhoneNumberParams{"+12546551377", swagger.RouteSummary{123, randomString(12)}, "Phone Name Now", true, true, "Phone N", "business", "extension", swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, stringEmailSlice, "+18587740222"}
			s.marshalInput(createPhoneNumberParamsSample, "createPhoneNumber", inputType)

		case createQueue:
			createQueueParamsSample := swagger.CreateQueueParams{randomString(12), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, 60, "called_number", 10, nil}
			s.marshalInput(createQueueParamsSample, "createQueue", inputType)

		case createRoute:
			createRouteJsonSample := CreateRouteJson{randomString(12), []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{int32(22035), "ntud62prqbl7"}}}}}}
			s.marshalInput(createRouteJsonSample, "createRoute", inputType)

		case createSms:
			createSmsParamsSample := swagger.CreateSmsParams{"+16309624775", "+12019570328", "Another message for create", 1767963}
			s.marshalInput(createSmsParamsSample, "createSms", inputType)

		case createSubaccount:
			createSubaccountJsonSample := CreateSubaccountJson{randomString(12), randomString(12), ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"},
				ContactJson{"Bobby", AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}, "+18585553333", "asd@sd.co"}}
			s.marshalInput(createSubaccountJsonSample, "createSubaccount", inputType)

		case createTrunk:
			createTrunkParamsSample := swagger.CreateTrunkParams{randomString(12), "SIP/1234@phone.com:5060", int32(60), int32(800), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, nil}
			s.marshalInput(createTrunkParamsSample, "createTrunk", inputType)

		case replaceDevice:
			s.marshalInput(createDeviceParams, "replaceDevice", inputType)

		case replaceExtension:
			replaceExtensionParamsSample := swagger.ReplaceExtensionParams{nil, nil, randomString(12), "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, nil, "standard", "private", 619, true, true, "automated", nil, "+18587741111", nil, "+18587748888", nil}
			s.marshalInput(replaceExtensionParamsSample, "replaceExtension", inputType)

		case replaceMenu:
			replaceMenuParamsSample := swagger.ReplaceMenuParams{randomString(12), nil, nil, false, 5, nil, nil}
			s.marshalInput(replaceMenuParamsSample, "replaceMenu", inputType)

		case replacePhoneNumber:
			replacePhoneNumberParamsSample := swagger.ReplacePhoneNumberParams{swagger.RouteSummary{123, randomString(12)}, "Robert", true, true, "Phone N", "business", "extension", swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, nil, stringEmailSlice, "+18587740222"}
			s.marshalInput(replacePhoneNumberParamsSample, "replacePhoneNumber", inputType)

		case replaceQueue:
			createQueueParamsSample := swagger.CreateQueueParams{randomString(12), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, 60, "called_number", 10, nil}
			s.marshalInput(createQueueParamsSample, "replaceQueue", inputType)

		case replaceRoute:
			replaceRouteJsonSample := ReplaceRouteJson{int32(4705073), randomString(12), []RulesJson{RulesJson{[]ActionsJson{ActionsJson{"queue", QueueJson{22026, "61kkjklmin74"}}}}}}
			s.marshalInput(replaceRouteJsonSample, "replaceRoute", inputType)

		case replaceTrunk:
			createTrunkParamsSample := swagger.CreateTrunkParams{randomString(12), "SIP/1234@phone.com:5060", int32(60), int32(800), swagger.MediaSummary{123, randomString(12)}, swagger.MediaSummary{123, randomString(12)}, nil}
			s.marshalInput(createTrunkParamsSample, "replaceTrunk", inputType)

		default:
			fmt.Printf("Could not create sample input file for command: %s\n", param.samplein)
		}

	}

	if param.sampleout != "" {
		stringEmailSlice := make([]string, 0)
		stringEmailSlice = append(stringEmailSlice, "asd@asd.com")

		switch param.sampleout {
		case getAccount:
			accountFullSample := swagger.AccountFull{int32(randomNumber(10, 9999)), randomString(12), randomString(12), randomString(12), swagger.AccountSummary{int32(randomNumber(10, 9999)), randomString(12)}, swagger.ContactAccount{randomString(12), randomString(12), swagger.Address{randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), strings.ToUpper(randomAlphaString(2))}, randomString(12), randomString(12), "primary@email.com", "alternate@email.com"}, swagger.ContactAccount{randomString(12), randomString(12), swagger.Address{randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), strings.ToUpper(randomAlphaString(2))}, randomString(12), randomString(12), "primary@email.com", "alternate@email.com"}}
			s.marshalInput(accountFullSample, "getAccount", inputType)

		case getApplication:
			applicationFullSample := swagger.ApplicationFull{int32(randomNumber(10, 9999)), randomString(12)}
			s.marshalInput(applicationFullSample, "getApplication", inputType)

		case getCallLog:
			callLogFullSample := swagger.CallFull{"45ab77e8-05ac-11e7-9437-26f1abd8fe53"}
			s.marshalInput(callLogFullSample, "getCallLog", inputType)

		case getDevice:
			deviceFullSample := swagger.DeviceFull{int32(randomNumber(10, 9999)), randomString(12), swagger.SipAuthentication{randomString(12), int32(randomNumber(10, 9999)), randomString(12), randomString(12)}, nil /*[]swagger.Line{int32(randomNumber(1,9999)), swagger.ExtensionSummary{int32(randomNumber(1,9999)), randomString(12), int32(randomNumber(1,9999))}}*/}
			s.marshalInput(deviceFullSample, "getDevice", inputType)

		case getExpressServiceCode:
			expressServiceCodeFullSample := swagger.ExpressServiceCodeFull{int32(randomNumber(1, 9999)), randomNumericString(8), int32(randomNumber(9999, 9999999999))}
			s.marshalInput(expressServiceCodeFullSample, "getExpressServiceCode", inputType)

		case getExtension:
			extensionFullSample := swagger.ExtensionFull{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999)), randomString(12), randomString(12), swagger.DeviceMembership{int32(randomNumber(1, 9999)), swagger.DeviceSummary{int32(randomNumber(1, 9999)), randomString(12)}}, randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, true, "+454654564534", randomString(12), true, true, swagger.Voicemail{true, randomString(12), swagger.Greeting{randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, true}, randomString(12), swagger.Notification{stringEmailSlice, randomString(12)}, randomString(12)}, swagger.Notification{stringEmailSlice, randomString(12)}, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}}
			s.marshalInput(extensionFullSample, "getExtension", inputType)

		case getContact:
			emailSlice := make([]swagger.Email, 0)
			emailSlice = append(emailSlice, swagger.Email{"primary", "asd@asd.com"})
			phoneNumberContactslice := make([]swagger.PhoneNumberContact, 0)
			phoneNumberContactslice = append(phoneNumberContactslice, swagger.PhoneNumberContact{"home", "+45454684545", "+45454684545"})
			addressListContacts := make([]swagger.AddressListContacts, 0)
			addressListContacts = append(addressListContacts, swagger.AddressListContacts{"home", randomString(12), randomString(12), randomString(12), randomNumericString(5), strings.ToUpper(randomAlphaString(2))})
			contactFullSample := swagger.ContactFull{int32(randomNumber(1, 9999)), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), randomString(12), emailSlice, phoneNumberContactslice, addressListContacts, swagger.GroupListContacts{int32(randomNumber(1, 9999)), randomString(12)}, int32(randomNumber(9999, 9999999999)), int32(randomNumber(9999, 9999999999))}
			s.marshalInput(contactFullSample, "getContact", inputType)

		case getGroup:
			groupFullSample := swagger.GroupFull{int32(randomNumber(1, 9999)), randomString(12)}
			s.marshalInput(groupFullSample, "getGroup", inputType)

		case getMedia:
			mediaFullSample := swagger.MediaFull{int32(randomNumber(1, 9999)), randomString(12), "hold_music"}
			s.marshalInput(mediaFullSample, "getRecording", inputType)

		case getMenu:
			optionSlice := make([]swagger.Option, 0)
			optionSlice = append(optionSlice, swagger.Option{randomNumericString(1), swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}})
			menuFullSample := swagger.MenuFull{int32(randomNumber(1, 9999)), randomString(12), true, int32(randomNumber(1, 9999)), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}, optionSlice}
			s.marshalInput(menuFullSample, "getMenu", inputType)

		case getPhoneNumber:
			phoneNumberFullSample := swagger.PhoneNumberFull{int32(randomNumber(1, 9999)), randomString(12), "+54654612511", true, true, swagger.RouteSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.CallerIdPhoneNumber{randomString(12), "bussiness"}, swagger.SmsForwarding{"extension", swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, swagger.ApplicationSummary{int32(randomNumber(1, 9999)), randomString(12)}}, swagger.CallNotifications{stringEmailSlice, "+45456486464"}}
			s.marshalInput(phoneNumberFullSample, "getPhoneNumber", inputType)

		case getQueue:
			memberSlice := make([]swagger.Member, 0)
			memberSlice = append(memberSlice, swagger.Member{swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, randomString(12)})
			queueFullSample := swagger.QueueFull{int32(randomNumber(1, 9999)), randomString(12), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.HoldMusic{int32(randomNumber(1, 9999)), randomString(12)}, 300, randomString(12), 20, memberSlice}
			s.marshalInput(queueFullSample, "getQueue", inputType)

		case getRoute:
			ruleSetForwardItemSlice := make([]swagger.RuleSetForwardItem, 0)
			ruleSetForwardItemSlice = append(ruleSetForwardItemSlice, swagger.RuleSetForwardItem{"+154514214546", swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, "+453484845122", true, "calling_number", randomString(12), "DEFAULT"})
			ruleSetActionSlice := make([]swagger.RuleSetAction, 0)
			ruleSetActionSlice = append(ruleSetActionSlice, swagger.RuleSetAction{randomString(12), swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, ruleSetForwardItemSlice, int32(randomNumber(5, 90)), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, int32(randomNumber(1, 60)), swagger.MenuSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.QueueSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.TrunkSummary{int32(randomNumber(1, 9999)), randomString(12)}})
			rulesetSlice := make([]swagger.RuleSet, 0)
			rulesetSlice = append(rulesetSlice, swagger.RuleSet{swagger.RuleSetFilter{"schedule", swagger.ScheduleSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.ContactSummary{int32(randomNumber(1, 9999)), "Mr", randomString(12), randomString(12), randomString(12), "Jr", randomString(12), randomString(12)}, swagger.GroupSummary{int32(randomNumber(1, 9999)), randomString(12)}}, ruleSetActionSlice})
			routeFullSample := swagger.RouteFull{int32(randomNumber(1, 9999)), randomString(12), swagger.ExtensionSummary{int32(randomNumber(1, 9999)), randomString(12), int32(randomNumber(1, 9999))}, rulesetSlice}
			s.marshalInput(routeFullSample, "getRoute", inputType)

		case getSchedule:
			scheduleFullSample := swagger.ScheduleFull{int32(randomNumber(1, 9999)), randomString(12)}
			s.marshalInput(scheduleFullSample, "getSchedule", inputType)

		case getSms:
			recipientSlice := make([]swagger.Recipient, 0)
			recipientSlice = append(recipientSlice, swagger.Recipient{"+12454354513", "sent"})
			smsFullSample := swagger.SmsFull{randomSmsId(), "+5646517686", recipientSlice, "in", int32(randomNumber(9999, 9999999999)), time.Time{}, randomString(12)}
			s.marshalInput(smsFullSample, "getSms", inputType)

		case getTrunk:
			stringCodeSlice := make([]string, 0)
			stringCodeSlice = append(stringCodeSlice, "g711u 64k")
			trunkFullSample := swagger.TrunkFull{int32(randomNumber(1, 9999)), randomString(12), "SIP/01%e@243.1.45.52:5060", int32(randomNumber(1, 100)), int32(randomNumber(500, 2000)), swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, swagger.MediaSummary{int32(randomNumber(1, 9999)), randomString(12)}, stringCodeSlice}
			s.marshalInput(trunkFullSample, "getTrunk", inputType)

		default:
			fmt.Printf("Could not create sample output file for command: %s\n", param.samplein)
		}
	}
}

func (s *SampleJsonCreator) marshalInput(param interface{}, fileName string, outputType string) {

	var marshalled []byte
	var err error

	if outputType == "json" {
		marshalled, err = json.MarshalIndent(param, "", "  ")
	} else if outputType == "xml" {
		marshalled, err = xml.MarshalIndent(param, "", "  ")
	}

	if err == nil {
		err = ioutil.WriteFile(fileName+"."+outputType, marshalled, 0644)
	}

	if err != nil {

		fmt.Printf("Could not create sample %s\n", outputType)
	} else {
		fmt.Println("Sample " + outputType + " created successfully")
	}

}
