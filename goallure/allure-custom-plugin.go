package goallure

/*
	Purpose:	custom plugin for allure-reports
	Developer:  Muhammad Haris Shafiq
	Date:		04-03-2020
*/

import(
	"fmt"
	//"encoding/json"
	//"os"
	//"io/ioutil"
)

const(
	//used for description types
	TEXT			= "TEXT"
	MARKDOWN		= "MARKDOWN"
	HTML			= "HTML"
	//used for testcase status
	PASSED			= "PASSED"
	FAILED			= "FAILED"
	SKIPPED			= "SKIPPED"
	BROKEN			= "BROKEN"
	PENDING			= "PENDING"
	CANCELED		= "CANCELED"
	//used for testcase severity
	NORMAL			= "NORMAL"
	CRITICAL		= "CRITICAL"
	TRIVIAL			= "TRIVIAL"
	BLOCKER			= "BLOCKER"
	MINOR			= "MINOR" 
)

//file-attachment 
type Attachment struct{
	Title		string		`json:"title"`
	// for now omitting size as it allure detects it by default
	// Size		string		`json:"size"`
	Type		string		`json:"type"`
	Source		string		`json:"source"`
}

//description of testcase
type Description struct{
	Type	string		`json:"type"`
	Value	string		`json:"value"`
}

//label for test-cases,suit,step
type Label struct{
	Name		string		`json:"name"`
	Value		string		`json:"value"`
}

//failure obj for test-case
type Failure struct{
	Message		string		`json:"message"`
	StackTrace	string		`json:"stackTrace"`
}

//step object for test-cases
type Step struct{
	Name	string		`json:"name"`
	Title	string		`json:"title"`
	Start   int64		`json:"start"`
	Stop	int64		`json:"stop"`
	Status	string		`json:"status"`
	//for now leaving nested steps
	//Steps	[]Step		`json:"steps"`
	Attachments	[]Attachment	`json:"attachments"`
}

//testcase
type TestCase struct{
	Name		string		`json:"name"`
	Title		string		`json:"title"`
	Description	Description	`json:"description"`
	Start		int64		`json:"start"`
	Stop		int64		`json:"stop"`
	Severity 	string		`json:"severity"`
	Status		string		`json:"status"`
	Failure		Failure		`json:"failure"`
	Attachments	[]Attachment`json:"attachments"`
	Steps		[]Step		`json:"steps"`
	Labels		[]Label		`json:"labels"`
}

//suit
type Suit 	struct{
	Name		string		`json:"name"`
	Title		string		`json:"title"`
	Start		int64		`json:"start"`
	Stop		int64		`json:"stop"`
	Version		string		`json:"version"`
	TestCases	[]TestCase	`json:"testCases"`
	Labels		[]Label		`json:"Labels"`
}

//create label object
func CreateLabel(name string, value string) (Label,error) {
	var label Label

	if (name != "" && value != ""){
		label.Name = name
		label.Value = value
		return label,nil
	}
	
	//empty values for label creation,return error
	return label,fmt.Errorf("Error! Invalid or  Empty Values for Label: provided values(%s,%s)\n",name,value)
}

//create description
func CreateDescription(dtype string,value string)(Description,error){
	var description Description 
	if (dtype == MARKDOWN || dtype == HTML || dtype == TEXT){
		description.Type = dtype
		description.Value = value
		return description,nil
	}

	//wrong description type
	return description,fmt.Errorf("Error! Invalid or  Empty Values for descrition type: provided values(%s,%s)\n",dtype,value)
}

//create attachment
func CreateAttachment(title string,atype string,src string)(Attachment,error){
	var attach Attachment
	if (title != "" && atype != "" && src !=""){
		attach.Title = title
		attach.Type  = atype
		attach.Source = src
		return attach,nil
	}

	//empty values
	return attach,fmt.Errorf("Error! Invalid or  Empty Values for Attachment object: provided values(%s,%s,%s)\n",title,atype,src)
}

//step
func CreateStep(name string,title string, start int64,stop int64, status string, attach []Attachment)(Step,error){
	var step Step
	if (name != "" && title != "" && start != 0 && stop != 0 && status != ""){
		step.Name=name
		step.Title=title
		step.Start=start
		step.Stop=stop
		step.Status=status
		step.Attachments=attach
		
		return step,nil
	}
	return step,fmt.Errorf("Error! Invalid or  Empty Values for step object.")
}

//testcase
func CreateTestCase(name string, title string, des Description, start int64, stop int64, severity string, status string, fail Failure, attachments []Attachment, steps []Step, labels []Label) (TestCase,error) {
	var testCase TestCase
	if(name != "" && start != 0 && stop != 0 && status != ""){
		testCase.Name	= name
		testCase.Title	= title
		testCase.Start	= start
		testCase.Stop	= stop
		testCase.Status	= status
		testCase.Description	= des
		testCase.Severity	= severity
		
		if(status == FAILED){
			testCase.Failure = fail
		}

		testCase.Attachments = attachments
		testCase.Steps	= steps
		testCase.Labels	= labels

		return testCase,nil
	}
	return testCase,fmt.Errorf("Error! Invalid or  Empty Values for testCase object.")
}

//suit
func CreateSuit(name string, title string, start int64, stop int64, version string, testCases []TestCase, labels []Label)(Suit,error){
	var suit Suit
	if (name != "" && start != 0 && stop != 0){
		suit.Name	= name
		suit.Title	= title
		suit.Start	= start
		suit.Stop	= stop
		suit.Version= version
		suit.TestCases = testCases
		suit.Labels	= labels
		return suit,nil
	}

	return suit, fmt.Errorf("Error! invalid or empty Values for suit object.")
}