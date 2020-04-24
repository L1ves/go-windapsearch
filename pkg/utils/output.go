package utils

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"io"
)

func WriteSearchResults(result *ldap.SearchResult, writer io.Writer) {
	if result == nil || len(result.Entries) == 0 {
		io.WriteString(writer, "[-] No results\n")
		return
	}
	for _, entry := range result.Entries {
		for _, attribute := range entry.Attributes {
			for _, value := range attribute.ByteValues {
				valueString := HandleLDAPBytes(attribute.Name, value)
				io.WriteString(writer, fmt.Sprintf("%s: %v\n", attribute.Name, valueString))
			}
		}
		io.WriteString(writer, "\n")

	}
	return
}
