package filesystem

import (
	"github.com/derhabicht/eagle-rock-cli/internal/documents/model/document"
	"github.com/derhabicht/eagle-rock-cli/internal/documents/repository"
	"github.com/derhabicht/eagle-rock-cli/pkg/documents"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFileRepository_IsIRepository(t *testing.T) {
	var _ repository.IRepository = (*FileRepository)(nil)
	assert.True(t, true)
}

func TestMemoParser_ParseMr(t *testing.T) {
	input := `
---
logo:           logos/barnards-star/barnards_star-500x500.png
address: |-     
    HEADQUARTERS, CRAZYLAND
    1600 Pennsylvania Ave
    Washington, DC 20003
tlp:            TLP:GREEN
control_number: MR-20-001
date:           2020-08-18
track:          PROJECT
subject:        'Test Memorandum for Record'
attachments:
    - 'Three Step Plan to World Domination'
cc:
    - Tweedle Dee
    - Tweedle Dum
status_history:
    registered:
        date: 2020-08-18
    published:
        date: 2020-08-18
...

####
Test paragraph 1.

####
Test paragraph 2.

####
Test paragraph 3.

####
Test paragraph 4. Oh look! A Table!

|         Date | Event                             |
|-------------:|-----------------------------------|
| 1903--12--17 | The most important day in history |
| 1969--07--20 | A close second                    |

---
name:      Tweedle Dee
signature: signatures/tweedle_dee.jpg
...
`
	expectedTlp, err := documents.ParseTlp("TLP:GREEN")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedDate, err := documents.ParseDate("2020-08-18")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedHeader := document.MrHeader{
		MemoHeader: document.MemoHeader{
			Logo: "logos/barnards-star/barnards_star-500x500.png",
			Address: `HEADQUARTERS, CRAZYLAND
1600 Pennsylvania Ave
Washington, DC 20003`,
			Tlp: expectedTlp,
			ControlNumber: documents.ControlNumber{
				Class:        documents.MR,
				Year:         2020,
				MainSequence: 1,
			},
			Date: &expectedDate,
			Attachments: []string{
				"Three Step Plan to World Domination",
			},
			Cc: []string{
				"Tweedle Dee",
				"Tweedle Dum",
			},
		},
		Track:   document.PROJECT,
		Subject: "Test Memorandum for Record",
		StatusHistory: document.MrStatusHistory{
			Registered: document.MrStatus{
				Date: expectedDate,
			},
			Published: &document.MrStatus{
				Date: expectedDate,
			},
		},
	}

	expectedBody := `
####
Test paragraph 1.

####
Test paragraph 2.

####
Test paragraph 3.

####
Test paragraph 4. Oh look! A Table!

|         Date | Event                             |
|-------------:|-----------------------------------|
| 1903--12--17 | The most important day in history |
| 1969--07--20 | A close second                    |
`

	expectedSignature := document.MemoSignature{
		Name:      "Tweedle Dee",
		Signature: "signatures/tweedle_dee.jpg",
	}

	result, err := memoParser(documents.MR, []byte(input))
	if err != nil {
		assert.FailNow(t, "%s", err)
	}

	resultHeader := result.Header().(document.MrHeader)
	resultBody := result.Body()
	resultSignature := result.Signature().(document.MemoSignature)

	assert.Equal(t, expectedHeader.Logo, resultHeader.Logo)
	assert.Equal(t, expectedHeader.Address, resultHeader.Address)
	assert.Equal(t, expectedHeader.Tlp, resultHeader.Tlp)
	assert.Equal(t, expectedHeader.ControlNumber, resultHeader.ControlNumber)
	assert.Equal(t, *expectedHeader.Date, *resultHeader.Date)
	assert.Equal(t, expectedHeader.Attachments, resultHeader.Attachments)
	assert.Equal(t, expectedHeader.Cc, resultHeader.Cc)
	assert.Equal(t, expectedHeader.Track, resultHeader.Track)
	assert.Equal(t, expectedHeader.Subject, resultHeader.Subject)
	assert.Equal(t, expectedHeader.StatusHistory, resultHeader.StatusHistory)

	assert.Equal(t, strings.Trim(expectedBody, "\n"), strings.Trim(resultBody, "\n"))

	assert.Equal(t, expectedSignature.Name, resultSignature.Name)
	assert.Equal(t, expectedSignature.Signature, resultSignature.Signature)

}

func TestMemoParser_ParseOpord(t *testing.T) {
	input := `
---
logo: logos/barnards-star/barnards_star-500x500.png
address: |-     
    HEADQUARTERS, CRAZYLAND
    1600 Pennsylvania Ave
    Washington, DC 20003
tlp:                TLP:GREEN
control_number:     OPORD-20-001
date:               2020-08-18
mission_number:     HERP-20-001
timezone:           UTC-05
ic:                 Tweedle Dee
references:
    - 'Three Step Plan to World Domination'
cc:
    - Tweedle Dee
    - Tweedle Dum
...

####
Test paragraph 1.

####
Test paragraph 2.

####
Test paragraph 3.

####
Test paragraph 4. Oh look! A Table!

|         Date | Event                             |
|-------------:|-----------------------------------|
| 1903--12--17 | The most important day in history |
| 1969--07--20 | A close second                    |

---
name:      Tweedle Dee
signature: signatures/tweedle_dee.jpg
...
`
	expectedTlp, err := documents.ParseTlp("TLP:GREEN")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedDate, err := documents.ParseDate("2020-08-18")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedHeader := document.OpordHeader{
		MemoHeader: document.MemoHeader{
			Logo: "logos/barnards-star/barnards_star-500x500.png",
			Address: `HEADQUARTERS, CRAZYLAND
1600 Pennsylvania Ave
Washington, DC 20003`,
			Tlp: expectedTlp,
			ControlNumber: documents.ControlNumber{
				Class:        documents.OPORD,
				Year:         2020,
				MainSequence: 1,
			},
			Date: &expectedDate,
			Cc: []string{
				"Tweedle Dee",
				"Tweedle Dum",
			},
		},
		MissionNumber:     "HERP-20-001",
		TimeZone:          "UTC-05",
		IncidentCommander: "Tweedle Dee",
		References: []string{
			"Three Step Plan to World Domination",
		},
	}

	expectedBody := `
####
Test paragraph 1.

####
Test paragraph 2.

####
Test paragraph 3.

####
Test paragraph 4. Oh look! A Table!

|         Date | Event                             |
|-------------:|-----------------------------------|
| 1903--12--17 | The most important day in history |
| 1969--07--20 | A close second                    |
`

	expectedSignature := document.MemoSignature{
		Name:      "Tweedle Dee",
		Signature: "signatures/tweedle_dee.jpg",
	}

	result, err := memoParser(documents.OPORD, []byte(input))
	if err != nil {
		assert.FailNow(t, "%s", err)
	}

	resultHeader := result.Header().(document.OpordHeader)
	resultBody := result.Body()
	resultSignature := result.Signature().(document.MemoSignature)

	assert.Equal(t, expectedHeader.Logo, resultHeader.Logo)
	assert.Equal(t, expectedHeader.Address, resultHeader.Address)
	assert.Equal(t, expectedHeader.Tlp, resultHeader.Tlp)
	assert.Equal(t, expectedHeader.ControlNumber, resultHeader.ControlNumber)
	assert.Equal(t, *expectedHeader.Date, *resultHeader.Date)
	assert.Equal(t, expectedHeader.Attachments, resultHeader.Attachments)
	assert.Equal(t, expectedHeader.Cc, resultHeader.Cc)
	assert.Equal(t, expectedHeader.MissionNumber, resultHeader.MissionNumber)
	assert.Equal(t, expectedHeader.TimeZone, resultHeader.TimeZone)
	assert.Equal(t, expectedHeader.IncidentCommander, resultHeader.IncidentCommander)
	assert.Equal(t, expectedHeader.References, resultHeader.References)

	assert.Equal(t, strings.Trim(expectedBody, "\n"), strings.Trim(resultBody, "\n"))

	assert.Equal(t, expectedSignature.Name, resultSignature.Name)
	assert.Equal(t, expectedSignature.Signature, resultSignature.Signature)

}

func TestWarnoParser(t *testing.T) {
	input := `
TLP:GREEN
20200209T2130Z HQ, CRAZYLAND
WARNO-20-001 HERP-20-001
TZ UTC-0500 IC TWEEDLE DEE

LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT. NULLAM VITAE IACULIS
LIBERO. DONEC PORTA IPSUM EU MOLLIS MOLESTIE. NULLA FAUCIBUS SODALES LIBERO AT
LOBORTIS. SUSPENDISSE VEL FINIBUS NUNC, NEC ULTRICIES MAURIS. PRAESENT SED ODIO
EUISMOD, TEMPOR QUAM QUIS, ELEIFEND NISL. AENEAN QUIS LIGULA IACULIS, EGESTAS
AUGUE AC, LAOREET MAURIS. ETIAM SED VEHICULA VELIT.

ETIAM CONSECTETUR PURUS QUIS LIGULA MALESUADA, NON CONGUE NUNC LOBORTIS. MAURIS
VULPUTATE, LECTUS QUIS IACULIS FAUCIBUS, MAURIS PURUS TINCIDUNT NEQUE, VEL
EUISMOD ARCU RISUS VEL DIAM. SED TEMPOR SIT AMET NISL A ELEIFEND. MAURIS TORTOR
ANTE, SEMPER VEL ORCI SED, LUCTUS VARIUS DOLOR. PRAESENT A FERMENTUM ARCU.
PHASELLUS IN PRETIUM MI, NON IACULIS ODIO. QUISQUE EGET CONDIMENTUM EROS.
ALIQUAM ERAT VOLUTPAT. QUISQUE IN VULPUTATE EROS.

SED LACUS NEQUE, VOLUTPAT UT PURUS SIT AMET, TINCIDUNT LACINIA SEM. FUSCE
CONSEQUAT LUCTUS ODIO, NEC DICTUM PURUS LACINIA QUIS. QUISQUE LAOREET LUCTUS
NISI, VITAE LAOREET EX ALIQUET NEC. NUNC VEL IPSUM VULPUTATE, LOBORTIS NISL UT,
MATTIS METUS. VIVAMUS IN ELIT BIBENDUM, MAXIMUS URNA VITAE, ELEMENTUM MAURIS.
NULLA NEC PLACERAT SEM, RHONCUS LOBORTIS NIBH. VESTIBULUM SEMPER, EROS EU
SCELERISQUE PORTTITOR, EST LECTUS RUTRUM EROS, EGET LACINIA LECTUS EX NON ANTE.
DUIS POSUERE ELEIFEND EST, SED PELLENTESQUE NULLA HENDRERIT EU. NULLAM EGESTAS
DICTUM ALIQUET. SED EU ELEMENTUM ELIT, AC DIGNISSIM TORTOR. PRAESENT PORTTITOR
LECTUS VEL DIAM DICTUM CONVALLIS. NAM MASSA ELIT, TINCIDUNT SOLLICITUDIN MI
EGET, POSUERE AUCTOR FELIS.

NULLA VITAE NIBH ET DIAM MOLESTIE AUCTOR EU A MAGNA. PELLENTESQUE VITAE LECTUS
CONDIMENTUM, POSUERE SEM A, ULTRICES ANTE. ETIAM ULTRICIES SEM NIBH, AC SEMPER
MAURIS TEMPUS AC. FUSCE SED SAPIEN SED ODIO FEUGIAT FEUGIAT. CRAS AT ELEIFEND
QUAM. DUIS SIT AMET PORTTITOR MI, ID IACULIS METUS. CRAS LEO NISI, ORNARE A
VELIT VITAE, MALESUADA COMMODO DUI. PHASELLUS EGET NULLA PELLENTESQUE,
ELEMENTUM DIAM EU, PRETIUM ERAT. ETIAM EGET LACUS URNA. CRAS AC VENENATIS ARCU.
VESTIBULUM CONDIMENTUM EX TEMPOR LOREM MOLESTIE, ID VENENATIS RISUS TEMPUS.
CRAS CONDIMENTUM DIAM NISL, EGET RUTRUM PURUS PRETIUM ID. SUSPENDISSE SAGITTIS
TORTOR EU NIBH COMMODO PULVINAR.

ALIQUAM UT LIGULA VESTIBULUM, VOLUTPAT ENIM A, DICTUM MI. NULLAM A IPSUM SIT
AMET ERAT DIGNISSIM POSUERE NON EU URNA. UT AT DAPIBUS VELIT. MAURIS FINIBUS
TURPIS FEUGIAT MAGNA ULTRICIES, AC EGESTAS MI VARIUS. FUSCE CONVALLIS ERAT
DIAM, UT POSUERE NISI LUCTUS UT. SUSPENDISSE VOLUTPAT NIBH ET TURPIS FERMENTUM,
QUIS POSUERE AUGUE SCELERISQUE. ALIQUAM A SEM NEC NISL ACCUMSAN CONGUE. NULLA
SIT AMET ACCUMSAN VELIT. NAM TINCIDUNT LECTUS VARIUS PURUS DIGNISSIM EGESTAS. 

// TWEEDLE DEE //
`

	expeectedTlp, err := documents.ParseTlp("TLP:GREEN")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedDtg, err := documents.ParseShortDtg("20200209T2130Z")
	if err != nil {
		assert.FailNow(t, "%s", err)
	}
	expectedHeader := document.WarnoHeader{
		Tlp:      expeectedTlp,
		DateTime: expectedDtg,
		Issuer:   "HQ, CRAZYLAND",
		ControlNumber: documents.ControlNumber{
			Class:        documents.WARNO,
			Year:         2020,
			MainSequence: 1,
		},
		MissionNumber:     "HERP-20-001",
		TimeZone:          "UTC-0500",
		IncidentCommander: "TWEEDLE DEE",
	}

	expectedBody := `
LOREM IPSUM DOLOR SIT AMET, CONSECTETUR ADIPISCING ELIT. NULLAM VITAE IACULIS
LIBERO. DONEC PORTA IPSUM EU MOLLIS MOLESTIE. NULLA FAUCIBUS SODALES LIBERO AT
LOBORTIS. SUSPENDISSE VEL FINIBUS NUNC, NEC ULTRICIES MAURIS. PRAESENT SED ODIO
EUISMOD, TEMPOR QUAM QUIS, ELEIFEND NISL. AENEAN QUIS LIGULA IACULIS, EGESTAS
AUGUE AC, LAOREET MAURIS. ETIAM SED VEHICULA VELIT.

ETIAM CONSECTETUR PURUS QUIS LIGULA MALESUADA, NON CONGUE NUNC LOBORTIS. MAURIS
VULPUTATE, LECTUS QUIS IACULIS FAUCIBUS, MAURIS PURUS TINCIDUNT NEQUE, VEL
EUISMOD ARCU RISUS VEL DIAM. SED TEMPOR SIT AMET NISL A ELEIFEND. MAURIS TORTOR
ANTE, SEMPER VEL ORCI SED, LUCTUS VARIUS DOLOR. PRAESENT A FERMENTUM ARCU.
PHASELLUS IN PRETIUM MI, NON IACULIS ODIO. QUISQUE EGET CONDIMENTUM EROS.
ALIQUAM ERAT VOLUTPAT. QUISQUE IN VULPUTATE EROS.

SED LACUS NEQUE, VOLUTPAT UT PURUS SIT AMET, TINCIDUNT LACINIA SEM. FUSCE
CONSEQUAT LUCTUS ODIO, NEC DICTUM PURUS LACINIA QUIS. QUISQUE LAOREET LUCTUS
NISI, VITAE LAOREET EX ALIQUET NEC. NUNC VEL IPSUM VULPUTATE, LOBORTIS NISL UT,
MATTIS METUS. VIVAMUS IN ELIT BIBENDUM, MAXIMUS URNA VITAE, ELEMENTUM MAURIS.
NULLA NEC PLACERAT SEM, RHONCUS LOBORTIS NIBH. VESTIBULUM SEMPER, EROS EU
SCELERISQUE PORTTITOR, EST LECTUS RUTRUM EROS, EGET LACINIA LECTUS EX NON ANTE.
DUIS POSUERE ELEIFEND EST, SED PELLENTESQUE NULLA HENDRERIT EU. NULLAM EGESTAS
DICTUM ALIQUET. SED EU ELEMENTUM ELIT, AC DIGNISSIM TORTOR. PRAESENT PORTTITOR
LECTUS VEL DIAM DICTUM CONVALLIS. NAM MASSA ELIT, TINCIDUNT SOLLICITUDIN MI
EGET, POSUERE AUCTOR FELIS.

NULLA VITAE NIBH ET DIAM MOLESTIE AUCTOR EU A MAGNA. PELLENTESQUE VITAE LECTUS
CONDIMENTUM, POSUERE SEM A, ULTRICES ANTE. ETIAM ULTRICIES SEM NIBH, AC SEMPER
MAURIS TEMPUS AC. FUSCE SED SAPIEN SED ODIO FEUGIAT FEUGIAT. CRAS AT ELEIFEND
QUAM. DUIS SIT AMET PORTTITOR MI, ID IACULIS METUS. CRAS LEO NISI, ORNARE A
VELIT VITAE, MALESUADA COMMODO DUI. PHASELLUS EGET NULLA PELLENTESQUE,
ELEMENTUM DIAM EU, PRETIUM ERAT. ETIAM EGET LACUS URNA. CRAS AC VENENATIS ARCU.
VESTIBULUM CONDIMENTUM EX TEMPOR LOREM MOLESTIE, ID VENENATIS RISUS TEMPUS.
CRAS CONDIMENTUM DIAM NISL, EGET RUTRUM PURUS PRETIUM ID. SUSPENDISSE SAGITTIS
TORTOR EU NIBH COMMODO PULVINAR.

ALIQUAM UT LIGULA VESTIBULUM, VOLUTPAT ENIM A, DICTUM MI. NULLAM A IPSUM SIT
AMET ERAT DIGNISSIM POSUERE NON EU URNA. UT AT DAPIBUS VELIT. MAURIS FINIBUS
TURPIS FEUGIAT MAGNA ULTRICIES, AC EGESTAS MI VARIUS. FUSCE CONVALLIS ERAT
DIAM, UT POSUERE NISI LUCTUS UT. SUSPENDISSE VOLUTPAT NIBH ET TURPIS FERMENTUM,
QUIS POSUERE AUGUE SCELERISQUE. ALIQUAM A SEM NEC NISL ACCUMSAN CONGUE. NULLA
SIT AMET ACCUMSAN VELIT. NAM TINCIDUNT LECTUS VARIUS PURUS DIGNISSIM EGESTAS. 
`
	expectedSignature := document.WarnoSignature{
		Name: "TWEEDLE DEE",
	}

	result, err := warnoParser([]byte(input))

	assert.NoError(t, err)
	assert.Equal(t, expectedHeader, result.Header().(document.WarnoHeader))
	assert.Equal(t, strings.Trim(expectedBody, "\n"), strings.Trim(result.Body(), "\n"))
	assert.Equal(t, expectedSignature, result.Signature().(document.WarnoSignature))
}
