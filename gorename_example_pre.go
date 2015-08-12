package main

var documentInfo string

func buildDocumentInfo(document string, hasAddress bool) string {
	documentInfo = document

	documentInfo += "Applicant Name: Regis Philbin"

	if hasAddress {
		documentInfo += "\n 123 Test Street"
		documentInfo += "\n Atlanta, GA 30045"
	}

	return documentInfo
}
