package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"math"
	"strconv"
	"strings"
)

type CliParams struct {
	slice          []string
	filtersId      []string
	limit          int32
	offset         int32
	id             int32
	idString       string
	dryRun         bool
	verbose        bool
	input          string
	command        string
	idSecondary    int32
	accountId      int32
	fields         string
	contact        string
	billingContact string
	filterType     string
	filterValue    string
	fullList       bool
	inputFormat    string
	outputFormat   string
	apiKey         string
	apiKeyPrefix   string
	sortType       string
	sortValue      string
	samplein       string
	sampleout      string

	filterParams FilterParams
	sortParams   SortParams
	otherParams  OtherParams

	from string
	to   string
	text string

	mediaFilePath string

	trunkName            string
	trunkUri             string
	trunkConcurrentCalls int32
	trunkMaxMinutes      int32
}

func createCliParams(context *cli.Context) (CliParams, error) {

	slice := make([]string, 0)
	limit := int32(context.Int(limitLong))
	offset := int32(context.Int(offsetLong))
	idString := context.String(idLong)
	var id int32 = 0

	_, err := strconv.Atoi(idString)

	var par CliParams

	if err == nil {
		idInt := 0
		idInt, err = strconv.Atoi(idString)
		id = int32(idInt)
	}

	verbose := context.Bool(verboseLong)
	dryRun := context.Bool(dryrunLong)
	input := context.String(inputLong)
	command := context.String(commandLong)
	idSecondary := int32(context.Int(idSecondaryLong))
	accountId := int32(context.Int(accountIdLong))
	contact := context.String(contactLong)
	billingContact := context.String(billingContactLong)
	fields := ""
	filtersType := context.String(filtersTypeLong)
	filtersValue := context.String(filtersValueLong)
	sortType := context.String(sortTypeLong)
	sortValue := context.String(sortValueLong)
	samplein := context.String(sampleInLong)
	sampleout := context.String(sampleOutLong)
	fullList := context.Bool(fullListLong)
	inputFormat := context.String(inputFormatLong)
	outputFormat := context.String(outputFormatLong)
	apiKey := context.String(apiKeyLong)
	apiKeyPrefix := context.String(apiKeyPrefixLong)

	var filtersId []string

	from := context.String(fromLong)
	to := context.String(toLong)
	text := context.String(textLong)

	mediaFilePath := context.String(mediaFileLong)

	trunkName := context.String(trunkNameLong)
	trunkUri := context.String(trunkUriLong)
	trunkConcurrentCalls := int32(context.Int(maxConcurrentCallsLong))
	trunkMaxMinutes := int32(context.Int(maxMinutesPerMonthLong))

	var filterParams FilterParams
	var sortParams SortParams
	var otherParams OtherParams

	if input != "" {
		var err error
		var listParams ListParams
		err, listParams = getListParams(input)
		if err != nil {
			return par, err
		}

		accountId = listParams.accountId
		limit = listParams.limit
		offset = listParams.offset
		fields = listParams.fields

		err, filterParams = getFiltersParams(input)

		if err != nil {
			return par, err
		}

		filtersId = filterParams.filtersId

		err, sortParams = getSortParams(input)

		if err != nil {
			return par, err
		}

		err, otherParams = getOtherParams(input)

		if err != nil {
			return par, err
		}
	}

	par.verbose = verbose
	par.dryRun = dryRun
	par.input = input
	par.command = command
	par.idSecondary = idSecondary
	par.accountId = accountId
	par.contact = contact
	par.billingContact = billingContact
	par.fields = fields
	par.filterType = filtersType
	par.filterValue = filtersValue
	par.sortType = sortType
	par.sortValue = sortValue
	par.samplein = samplein
	par.sampleout = sampleout
	par.fullList = fullList
	par.slice = slice
	par.limit = limit
	par.offset = offset
	par.id = id
	par.idString = idString
	par.apiKey = apiKey
	par.apiKeyPrefix = apiKeyPrefix
	par.inputFormat = inputFormat
	par.outputFormat = outputFormat

	par.filterParams = filterParams
	par.sortParams = sortParams
	par.otherParams = otherParams

	par.filtersId = filtersId

	par.from = from
	par.to = to
	par.text = text

	par.mediaFilePath = mediaFilePath

	par.trunkName = trunkName
	par.trunkUri = trunkUri
	par.trunkConcurrentCalls = trunkConcurrentCalls
	par.trunkMaxMinutes = trunkMaxMinutes

	if par.sortType != "" && par.sortValue != "" {

		sortTypes := strings.Split(par.sortType, ";")
		sortValues := strings.Split(par.sortValue, ";")
		sortTypesLength := float64(len(sortTypes))
		sortValuesLength := float64(len(sortValues))
		min := int(math.Min(float64(len(sortTypes)), float64(len(sortValues))))
		if sortTypesLength != sortValuesLength {
			fmt.Println("Warning: The number of sort types is not the same with the number of sort values!")
		}

		for counter := 0; counter < min; counter++ {
			switch sortTypes[counter] {
			case "id":
				par.sortParams.sortId = sortValues[counter]
			case "name":
				par.sortParams.sortName = sortValues[counter]
			case "start_time":
				par.sortParams.sortStartTime = sortValues[counter]
			case "created_at":
				par.sortParams.sortCreatedAt = sortValues[counter]
			case "extension":
				par.sortParams.sortExtension = sortValues[counter]
			case "number":
				par.sortParams.sortNumber = sortValues[counter]
			case "updated_at":
				par.sortParams.sortUpdatedAt = sortValues[counter]
			case "phone_number":
				par.sortParams.sortPhoneNumber = sortValues[counter]
			case "internal":
				par.sortParams.sortInternal = sortValues[counter]
			case "price":
				par.sortParams.sortPrice = sortValues[counter]
			case "npa":
				par.sortParams.sortNpa = sortValues[counter]
			case "nxx":
				par.sortParams.sortNxx = sortValues[counter]
			case "is_toll_free":
				par.sortParams.sortIsTollFree = sortValues[counter]
			case "city":
				par.sortParams.sortCity = sortValues[counter]
			case "province_postal_code":
				par.sortParams.sortProvincePostalCode = sortValues[counter]
			case "country_postal_code":
				par.sortParams.sortCountryPostalCode = sortValues[counter]
			case "country_code":
				par.sortParams.sortCountryCode = sortValues[counter]
			}
		}
	}

	if par.filterType != "" && par.filterValue != "" {

		filtersTypes := strings.Split(par.filterType, ";")
		filtersValues := strings.Split(par.filterValue, ";")
		filtersTypesLength := float64(len(filtersTypes))
		filtersValuesLength := float64(len(filtersValues))
		min := int(math.Min(filtersTypesLength, filtersValuesLength))
		if filtersTypesLength != filtersValuesLength {
			fmt.Println("Warning: The number of filters types is not the same with the number of filters values!")
		}

		for counter := 0; counter < min; counter++ {
			filter := []string{filtersValues[counter]}

			switch filtersTypes[counter] {
			case "id":
				par.filtersId = filter

			case "name":
				par.filterParams.filtersName = filter

			case "start_time":
				par.filterParams.filtersStartTime = filter

			case "created_at":
				par.filterParams.filtersCreatedAt = filter[0]

			case "direction":
				par.filterParams.filtersDirection = filter[0]

			case "called_number":
				par.filterParams.filtersCalledNumber = filter[0]

			case "type":
				par.filterParams.filtersType = filter[0]

			case "extension":
				par.filterParams.filtersExtension = filter

			case "number":
				par.filterParams.filtersNumber = filter

			case "group_id":
				par.filterParams.filtersGroupId = filter

			case "updated_at":
				par.filterParams.filtersUpdatedAt = filter

			case "phone_number":
				par.filterParams.filtersPhoneNumber = filter

			case "from":
				par.filterParams.filtersFrom = filter[0]

			case "to":
				par.filterParams.filtersTo = filter

			case "country_code":
				par.filterParams.filtersCountryCode = filter

			case "npa":
				par.filterParams.filtersNpa = filter

			case "nxx":
				par.filterParams.filtersNxx = filter

			case "xxxx":
				par.filterParams.filtersXxxx = filter

			case "city":
				par.filterParams.filtersCity = filter

			case "province":
				par.filterParams.filtersProvince = filter

			case "country":
				par.filterParams.filtersCountry = filter

			case "price":
				par.filterParams.filtersPrice = filter

			case "category":
				par.filterParams.filtersCategory = filter

			case "is_toll_free":
				par.filterParams.filtersIsTollFree = filter

			case "province_postal_code":
				par.filterParams.filtersProvincePostalCode = filter

			case "country_postal_code":
				par.filterParams.filtersCountryPostalCode = filter

			default:
				return par, errors.New(fmt.Sprintf(msgFilterTypeNotRecognized, par.filterType))
			}
		}

	}

	return par, nil
}
