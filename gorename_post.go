package main

var resumeInfo string

func buildDocumentInfo(document string, hasAddress bool) string { // HL
	resumeInfo = document

	resumeInfo += "Applicant Name: Regis Philbin"

	if hasAddress {
		resumeInfo += "\n 123 Test Street"
		resumeInfo += "\n Atlanta, GA 30045"
	}

	return resumeInfo
}
