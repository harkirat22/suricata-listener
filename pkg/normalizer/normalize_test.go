package normalizer_test

import (
	"os"
	"strings"

	"github.com/harkirat22/suricata-listener/pkg/normalizer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Normalizer", func() {
	Context("When reading log entries", func() {
		It("Should parse log entries correctly", func() {
			// Test data in NDJSON format.
			input := `{"event_type": "alert", "timestamp": "2023-09-20T12:34:56.789", "src_ip": "192.168.0.1", "alert": {"signature": "Suspicious connection to port 13666"}}`

			// Create a temporary file and write the test data to it.
			tempFile, err := os.CreateTemp("", "test*.json")
			Expect(err).ToNot(HaveOccurred())
			defer os.Remove(tempFile.Name()) // Cleanup after test.

			_, err = tempFile.WriteString(strings.TrimSpace(input))
			Expect(err).ToNot(HaveOccurred())

			// Close the file to ensure that all buffered data gets written.
			err = tempFile.Close()
			Expect(err).ToNot(HaveOccurred())

			// Re-open the file for reading.
			tempFile, err = os.Open(tempFile.Name())
			Expect(err).ToNot(HaveOccurred())

			// Use the temporary file with the ReadLogEntries function.
			entries, _, err := normalizer.ReadLogEntries(tempFile, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(entries).To(HaveLen(1))
			Expect(entries[0].Type).To(Equal("alert"))
			Expect(entries[0].SrcIP).To(Equal("192.168.0.1"))
			Expect(entries[0].Alert.Signature).To(Equal("Suspicious connection to port 13666"))
			// ... more assertions ...

			// Close the file.
			Expect(tempFile.Close()).ToNot(HaveOccurred())
		})
	})
})
